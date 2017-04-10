package network

import (
	"net"
)

//TCPClient cls
type TCPClient struct {
	session    *TCPSession
	dispatcher *PackDispatcher
}

func newTCPClient(conn net.Conn, codec Codec, sendChanSize int) *TCPClient {
	client := &TCPClient{
		dispatcher: newPackDispatcher(),
	}
	session := newTCPSession(conn, codec, sendChanSize, client.dispatcher)
	client.session = session

	return client
}

//SetPackDispatcher setter
func (client *TCPClient) SetPackDispatcher(dispatcher *PackDispatcher) {
	client.dispatcher = dispatcher
}

//GetPackDispatcher getter
func (client *TCPClient) GetPackDispatcher() *PackDispatcher {
	return client.dispatcher
}

//WriteMsg write msg
func (client *TCPClient) WriteMsg(msg PackInf) {
	client.session.WriteMsg(msg)
}

//Close close
func (client *TCPClient) Close() {
	client.session.Close()
}

//Loop logic loop
func (client *TCPClient) Loop(handler HandleFunc) {
	//TODO
	handler(client.session)
}
