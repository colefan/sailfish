package gate

import (
	"sync"

	"github.com/colefan/sailfish/network"
)

const (
	userSessionGroup = 64
)

// ClientSessionMap user session map
type clientSessionMap struct {
	sync.RWMutex
	sessions map[uint64]*network.TCPSession
}

// ClientSessionMnt client session  mnt
type ClientSessionMnt struct {
	sessionMap     clientSessionMap
	userSessionMap [userSessionGroup]clientSessionMap
}

// FindBySessionID find by sessionID
func (mnt *ClientSessionMnt) FindBySessionID(id uint64) *network.TCPSession {
	smap := &mnt.sessionMap
	smap.RLock()
	defer smap.RUnlock()
	if session, ok := smap.sessions[id]; ok {
		return session
	}
	return nil
}

// AddSessionByUID ~
func (mnt *ClientSessionMnt) AddSessionByUID(uid uint64, session *network.TCPSession) {
	groupID := uid % userSessionGroup
	smap := &mnt.userSessionMap[groupID]
	smap.Lock()
	defer smap.Unlock()
	if smap.sessions == nil {
		smap.sessions = make(map[uint64]*network.TCPSession)
	}
	smap.sessions[uid] = session
}

// FindByUID find by uid
func (mnt *ClientSessionMnt) FindByUID(id uint64) *network.TCPSession {
	groupID := id % userSessionGroup
	smap := &mnt.userSessionMap[groupID]
	smap.RLock()
	defer smap.RUnlock()
	if session, ok := smap.sessions[id]; ok {
		return session
	}
	return nil
}

// AddSession add session
func (mnt *ClientSessionMnt) AddSession(session *network.TCPSession) {
	smap := &mnt.sessionMap
	smap.Lock()
	defer smap.Unlock()
	smap.sessions[session.ID()] = session
}

// DelSession del session
func (mnt *ClientSessionMnt) DelSession(session *network.TCPSession) {
	smap := &mnt.sessionMap
	smap.Lock()
	delete(smap.sessions, session.ID())
	smap.Unlock()
	if user, ok := session.UserData().(*ClientUserData); ok {
		if user.UID != 0 {
			userMap := &mnt.userSessionMap[user.UID%userSessionGroup]
			userMap.Lock()
			defer userMap.Unlock()
			delete(userMap.sessions, user.UID)
		}
	}

}

var clientSessionMntInst *ClientSessionMnt

// GetClientSessionMntInst get instance
func GetClientSessionMntInst() *ClientSessionMnt {
	return clientSessionMntInst
}

func init() {
	clientSessionMntInst = new(ClientSessionMnt)
	clientSessionMntInst.sessionMap.sessions = make(map[uint64]*network.TCPSession)
	for i := 0; i < len(clientSessionMntInst.userSessionMap); i++ {
		clientSessionMntInst.userSessionMap[i].sessions = make(map[uint64]*network.TCPSession)
	}
}
