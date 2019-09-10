package network

import (
	"net"
)

//TCPClient cls
type TCPClient struct {
	session      *TCPSession
	dispatcher   PackDispatcherInf
	mode         MsgHandlerModeType
	isWebSocket  bool
	agentHandler SessionHandler
}

//Start  start client
func (c *TCPClient) Start() {
	if c.mode == MsgHandleModeEvent {
		c.session.Start(false)
	} else {
		c.session.Start(true)
	}
}

func newTCPClient(conn net.Conn, codec Codec, sendChanSize int, mode MsgHandlerModeType, dispatcher PackDispatcherInf, agentHandler SessionHandler) *TCPClient {
	client := &TCPClient{
		dispatcher:   dispatcher,
		mode:         mode,
		isWebSocket:  false,
		agentHandler: agentHandler,
	}
	session := newTCPSession(conn, codec, sendChanSize, client.dispatcher, client.agentHandler)
	client.session = session

	return client
}

//SetPackDispatcher setter
func (client *TCPClient) SetPackDispatcher(dispatcher PackDispatcherInf) {
	client.dispatcher = dispatcher
}

//GetPackDispatcher getter
func (client *TCPClient) GetPackDispatcher() PackDispatcherInf {
	return client.dispatcher
}

//WriteMsg write msg
func (client *TCPClient) WriteMsg(msg PackInf) {
	client.session.WriteMsg(msg)
}

//ReadMsg read msg
func (client *TCPClient) ReadMsg() (PackInf, error) {
	return client.session.ReadMsg()
}

//Close close
func (client *TCPClient) Close() {
	client.session.Close()
}

// SetHandler set handler
func (client *TCPClient) SetHandler(agent SessionHandler) {
	client.agentHandler = agent
}

// GetHandler getter
func (client *TCPClient) GetHandler() SessionHandler {
	return client.agentHandler
}
