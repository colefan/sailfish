package network

import (
	"net"
)

//TCPClient cls
type TCPClient struct {
	session    *TCPSession
	dispatcher PackDispatcherInf
	mode       int
}

//Start  start client
func (c *TCPClient) Start() {
	if c.mode == ModeEvent {
		c.session.Start(false)
	} else {
		c.session.Start(true)
	}
}

func newTCPClient(conn net.Conn, codec Codec, sendChanSize int, mode int, dispatcher PackDispatcherInf) *TCPClient {
	client := &TCPClient{
		dispatcher: dispatcher,
		mode:       mode,
	}
	session := newTCPSession(conn, codec, sendChanSize, client.dispatcher)
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
