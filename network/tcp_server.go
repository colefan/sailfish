package network

import (
	"net"
	"time"

	"github.com/colefan/sailfish/log"
	"github.com/gorilla/websocket"
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
			log.Errorf("listenr accept error %v", err)
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
	sessMgr                          *SessionMgr
	dispatcher                       PackDispatcherInf
	listener                         net.Listener
	sessionSendChanBuff              int
	protocol                         Protocol
	id                               uint64
	serverID                         string
	mode                             MsgHandlerModeType
	needCheckPerSecond               bool
	isWebSocket                      bool
	wsUpgrader                       websocket.Upgrader
	handler                          HandleFunc
	netWorkType                      string
	address                          string
	sessionIdleTimeout               int //读超时设置
	webSocketPath                    string
	sessionCloseCallbackForModeEvent func(session *TCPSession)
	agentHandler                     SessionHandler
}

func newTCPServer(network string, l net.Listener, p Protocol, sendChanSize int, mode MsgHandlerModeType, dispacher PackDispatcherInf, agentHandler SessionHandler) *TCPServer {
	s := &TCPServer{
		listener:            l,
		protocol:            p,
		sessionSendChanBuff: sendChanSize,
		mode:                mode,
		dispatcher:          dispacher,
		needCheckPerSecond:  false,
		isWebSocket:         false,
		netWorkType:         network,
		agentHandler:        agentHandler,
	}
	if s.sessionSendChanBuff <= 0 {
		s.sessionSendChanBuff = defaultSendChannelBufferSize
	}

	s.sessMgr = newManager()
	//s.dispatcher = newPackDispatcher()
	return s
}

func (s *TCPServer) SetCheckPerSecond(b bool) {
	s.needCheckPerSecond = b
}

//GetSessionMgr getter
func (s *TCPServer) GetSessionMgr() *SessionMgr {
	return s.sessMgr
}

//GetPackDispatcher getter
func (s *TCPServer) GetPackDispatcher() PackDispatcherInf {
	return s.dispatcher
}

//SetPackDispatcher setter
func (s *TCPServer) SetPackDispatcher(dispatcher PackDispatcherInf) {
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
	if s.netWorkType == "ws" || s.netWorkType == "wss" {
		s.serveWs(handler)
	} else {
		s.serveSocket(handler)
	}

}

func (s *TCPServer) serveSocket(handler HandleFunc) {
	log.Info("TcpServer Listen " + s.listener.Addr().String())

	for {
		//	fmt.Println("wait next connection ")
		conn, err := accept(s.listener)
		//conn.SetDeadline
		log.Info(">>>>>>tcp server accept a new connection")
		if err != nil {
			log.Error("tcp server accept conn error :" + err.Error())
			return
		}

		codec, err := s.protocol.NewSocketCodec(conn)
		if err != nil {
			log.Error("new socket codec error :" + err.Error())
			conn.Close()
			continue
		}
		session := newTCPSession(conn, codec, s.sessionSendChanBuff, s.dispatcher, s.agentHandler)
		session.SetTimeOut(s.sessionIdleTimeout)
		session.SetCloseCallback(s.afterSessionClosed)
		session.SetCloseCallbackForModeEvent(s.sessionCloseCallbackForModeEvent)
		GetTCPServerQos().AddAcceptedSessions()
		if s.needCheckPerSecond {
			session.SetCheckPerSecond(true)
		}

		s.GetSessionMgr().AddSession(session)
		if handler != nil {
			session.Start(true)
			//go handler(session)
		} else {
			session.Start(false)
		}

	}
}

func (s *TCPServer) afterSessionClosed(session *TCPSession) {
	s.GetSessionMgr().RemoveSession(session)
	GetTCPServerQos().AddClosedSessions()
}

//Stop stop and distroy server
func (s *TCPServer) Stop() {
	if s.listener != nil {
		log.Info("before listener close")
		s.listener.Close()
		log.Info("after listener closed")
		if s.sessMgr != nil {
			log.Info("before session manager closed")
			s.sessMgr.Dispose()
			log.Info("after session manager closed")
		}
		if s.dispatcher != nil {
			log.Info("before dispatcher closed")
			s.dispatcher.Dispose()
			log.Info("after dispatcher closed")
		}
	}

}

// SetSessionHandler setter only ModeHandle valid
func (s *TCPServer) SetSessionHandler(h HandleFunc) {
	s.handler = h
}

func (s *TCPServer) SetSessionIdleTimeOut(t int) {
	s.sessionIdleTimeout = t
}

func (s *TCPServer) GetWebSocketPath() string {
	if s.webSocketPath == "" {
		return "ws"
	} else {
		return s.webSocketPath
	}
}

func (s *TCPServer) SetWebSocketPath(path string) {
	s.webSocketPath = path
}

func (s *TCPServer) SetSessionCloseCallbackForModeEvent(f func(s *TCPSession)) {
	s.sessionCloseCallbackForModeEvent = f
}

//HandleFunc  handler
type HandleFunc func(session *TCPSession)
