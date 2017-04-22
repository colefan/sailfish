package network

import "sync"

const sessionGroupNum = 32

//SessionMgr manager session
type SessionMgr struct {
	sessionMaps [sessionGroupNum]sessionMap
	disposeFlag bool
	disposeOnce sync.Once
	disposeWait sync.WaitGroup
}

func newManager() *SessionMgr {
	manager := &SessionMgr{}
	for i := 0; i < len(manager.sessionMaps); i++ {
		manager.sessionMaps[i].sessions = make(map[uint64]*TCPSession)
	}
	return manager
}

type sessionMap struct {
	sync.RWMutex
	sessions map[uint64]*TCPSession
}

//RemoveSession remove session from manager
func (manager *SessionMgr) RemoveSession(session *TCPSession) {
	manager.delSession(session)
}

func (manager *SessionMgr) delSession(session *TCPSession) {
	if manager.disposeFlag {
		//fmt.Println("delete in flag")
		manager.disposeWait.Done()
		return
	}

	smap := &manager.sessionMaps[session.ID()%sessionGroupNum]
	smap.Lock()
	defer smap.Unlock()
	if _, ok := smap.sessions[session.id]; ok {
		delete(smap.sessions, session.id)
		//fmt.Println("delete out flag")
		manager.disposeWait.Done()
	}

}

//GetSession get session by sessionID
func (manager *SessionMgr) GetSession(sessionID uint64) *TCPSession {
	smap := &manager.sessionMaps[sessionID%sessionGroupNum]
	smap.RLock()
	defer smap.RUnlock()
	session, _ := smap.sessions[sessionID]
	return session
}

//AddSession add session in sessionMgr
func (manager *SessionMgr) AddSession(session *TCPSession) {
	manager.putSession(session)
}

func (manager *SessionMgr) putSession(session *TCPSession) {
	smap := &manager.sessionMaps[session.id%sessionGroupNum]
	smap.Lock()
	defer smap.Unlock()
	smap.sessions[session.id] = session
	//fmt.Println("add session")
	manager.disposeWait.Add(1)
}

//Dispose dispose session manager
func (manager *SessionMgr) Dispose() {
	manager.disposeOnce.Do(func() {
		manager.disposeFlag = true
		for i := 0; i < sessionGroupNum; i++ {
			smap := &manager.sessionMaps[i]
			smap.Lock()
			for _, session := range smap.sessions {
				session.ForceClose()
				manager.delSession(session)
			}
			smap.Unlock()
		}
		manager.disposeWait.Wait()

	})

}
