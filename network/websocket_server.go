package network

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// createWebSocketServer create a websocket server with params
func createWebSocketServer(network, address string, p Protocol, sendChanSize int, mode MsgHandlerModeType, d PackDispatcherInf, agentHandler SessionHandler) (*TCPServer, error) {

	s := &TCPServer{
		listener:            nil,
		protocol:            p,
		sessionSendChanBuff: sendChanSize,
		mode:                mode,
		dispatcher:          d,
		needCheckPerSecond:  false,
		isWebSocket:         true,
		netWorkType:         network,
		address:             address,
		agentHandler:        agentHandler,
	}
	if s.sessionSendChanBuff <= 0 {
		s.sessionSendChanBuff = defaultSendChannelBufferSize
	}
	s.wsUpgrader = websocket.Upgrader{ReadBufferSize: 1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	s.sessMgr = newManager()

	return s, nil
}

func (s *TCPServer) serveWs(handle HandleFunc) {
	s.SetSessionHandler(handle)
	netInfo("listen path /" + s.GetWebSocketPath())
	http.HandleFunc("/"+s.GetWebSocketPath(), s.wsHandShake)
	netInfo(s.netWorkType + "://" + s.address)
	var err error
	if s.netWorkType == "ws" {
		err = http.ListenAndServe(s.address, nil)
	} else {
		err = http.ListenAndServeTLS(s.address, "./config/domain.crt", "./config/domain.key", nil)
	}
	if err != nil {
		fmt.Printf("websocket listen error :%v\n", err)
		netError("websocket listen error :%v", err)
		panic("websocket error")
	}

}

func (s *TCPServer) wsHandShake(w http.ResponseWriter, r *http.Request) {
	ws, err := s.wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		netError("ws upgrade error :%v", err)
		return
	}

	codec, err := s.protocol.NewWsSocketCodec(ws)
	if err != nil {
		netError("new websocket codec error :" + err.Error())
		ws.Close()
		return
	}

	session := newWsSession(ws, codec, s.sessionSendChanBuff, s.dispatcher, s.agentHandler)
	netInfo(">>>>>>websocket server accept a new connection")
	GetTCPServerQos().AddAcceptedSessions()
	if s.needCheckPerSecond {
		session.SetCheckPerSecond(true)
	}
	session.SetTimeOut(s.sessionIdleTimeout)
	session.SetCloseCallback(s.afterSessionClosed)
	session.SetCloseCallbackForModeEvent(s.sessionCloseCallbackForModeEvent)
	s.GetSessionMgr().AddSession(session)
	if s.handler != nil {
		session.Start(true)
		// go s.handler(session)
	} else {
		session.Start(false)
	}

}
