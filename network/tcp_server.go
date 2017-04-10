package network

import (
	"fmt"
	"net"
	"time"
)

const (
	defaultSendChannelBufferSize = 100
)

//accept begin to accept net connection
func accept(listener net.Listener) (net.Conn, error) {
	var tempDelay time.Duration
	for {
		conn, err := listener.Accept()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if tempDelay > 1*time.Second {
					tempDelay = 1 * time.Second
				}
				time.Sleep(tempDelay)
				continue
			}
			return nil, err
		}
		return conn, err
	}

}

//TCPServer cls
type TCPServer struct {
	sessMgr             *SessionMgr
	dispatcher          *PackDispatcher
	listener            net.Listener
	sessionSendChanBuff int
	protocol            Protocol
	id                  uint64
	serverID            string
}

func newTCPServer(l net.Listener, p Protocol, sendChanSize int) *TCPServer {
	s := &TCPServer{
		listener:            l,
		protocol:            p,
		sessionSendChanBuff: sendChanSize,
	}
	if s.sessionSendChanBuff <= 0 {
		s.sessionSendChanBuff = defaultSendChannelBufferSize
	}

	s.sessMgr = newManager()
	s.dispatcher = newPackDispatcher()
	return s
}

//GetSessionMgr getter
func (s *TCPServer) GetSessionMgr() *SessionMgr {
	return s.sessMgr
}

//GetPackDispatcher getter
func (s *TCPServer) GetPackDispatcher() *PackDispatcher {
	return s.dispatcher
}

//SetPackDispatcher setter
func (s *TCPServer) SetPackDispatcher(dispatcher *PackDispatcher) {
	s.dispatcher = dispatcher
}

//ID getter
func (s *TCPServer) ID() uint64 {
	return s.id
}

//SetID setter
func (s *TCPServer) SetID(id uint64) {
	s.id = id
}

//SetServerID setter
func (s *TCPServer) SetServerID(sid string) {
	s.serverID = sid
}

//GetServerID getter
func (s *TCPServer) GetServerID() string {
	return s.serverID
}

//Serve server loop
func (s *TCPServer) Serve(handler HandleFunc) {
	for {
		conn, err := accept(s.listener)
		if err != nil {
			//TODO LOG
			fmt.Println("[ERROR] accept conn error :" + err.Error())
			return
		}
		GetTCPServerQos().AddAcceptedSessions()

		codec, err := s.protocol.NewCodec(conn)
		if err != nil {
			//TODO LOG
			fmt.Println("[ERROR] new codec error :" + err.Error())
			conn.Close()
			continue
		}
		session := newTCPSession(conn, codec, s.sessionSendChanBuff, s.dispatcher)
		s.GetSessionMgr().AddSession(session)
		if handler != nil {
			session.Start(true)
			go handler(session)
		} else {
			session.Start(false)
		}

	}

}

//Stop stop and distroy server
func (s *TCPServer) Stop() {
	if s.listener != nil {
		s.listener.Close()
		s.sessMgr.Dispose()
		s.dispatcher.Dispose()
	}

}

//HandleFunc  handler
type HandleFunc func(session *TCPSession)
