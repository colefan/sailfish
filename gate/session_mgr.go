package gate

import "github.com/colefan/sailfish/network"
import "sync"

const (
	defaultServerListSize = 32
)

//ServerSessionMgr mgr server sessions
type ServerSessionMgr struct {
	sync.RWMutex
	sessionMap    map[uint64]*ServerSession
	serverTypeMap map[int][]*ServerSession
}

//NewServerSessionMgr create serverSessionMgr
func NewServerSessionMgr() *ServerSessionMgr {
	mgr := &ServerSessionMgr{
		sessionMap:    make(map[uint64]*ServerSession),
		serverTypeMap: make(map[int][]*ServerSession),
	}
	return mgr
}

//GetSession find session
func (mgr *ServerSessionMgr) GetSession(id uint64) *ServerSession {
	mgr.RLock()
	defer mgr.RUnlock()
	session, _ := mgr.sessionMap[id]
	return session
}

//AddSession add session
func (mgr *ServerSessionMgr) AddSession(session *ServerSession) {
	id := session.Session.ID()
	mgr.Lock()
	defer mgr.Unlock()
	if _, ok := mgr.sessionMap[id]; ok {
		gLog.Warn("session has exits.")
	}
	mgr.sessionMap[id] = session

}

//DeleteSession delete session
func (mgr *ServerSessionMgr) DeleteSession(session *ServerSession) {
	if session == nil {
		return
	}
	mgr.Lock()
	defer mgr.Unlock()

	delete(mgr.sessionMap, session.Session.ID())
}

//AddServerList add server list by server type
func (mgr *ServerSessionMgr) AddServerList(serverType int, session *ServerSession) {
	mgr.Lock()
	defer mgr.Unlock()
	serverList, ok := mgr.serverTypeMap[serverType]
	if !ok {
		serverList = make([]*ServerSession, 0, defaultServerListSize)
		mgr.serverTypeMap[serverType] = serverList
	}
	bFound := false
	for _, val := range serverList {
		if val.Session.ID() == session.Session.ID() {
			bFound = true
			break
		}
	}
	if bFound {
		serverList = append(serverList, session)
	}
	mgr.serverTypeMap[serverType] = serverList
}

//RemoveFromServerList remove from server list
func (mgr *ServerSessionMgr) RemoveFromServerList(serverType int, session *ServerSession) {
	mgr.Lock()
	defer mgr.Unlock()
	if serverList, ok := mgr.serverTypeMap[serverType]; ok {
		for idx, v := range serverList {
			if v.Session.ID() == session.Session.ID() {
				serverList = append(serverList[:idx], serverList[idx+1:]...) //TODO
			}
		}

	}
}

//ServerSession class
type ServerSession struct {
	Session    *network.TCPSession
	serverType int
	userCount  int
}

//NewServerSession creator
func NewServerSession(session *network.TCPSession) *ServerSession {
	s := &ServerSession{
		Session:    session,
		serverType: 0,
		userCount:  0,
	}
	return s
}
