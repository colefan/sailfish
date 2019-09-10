package network

import (
	"errors"
	"net"
)

// newServer 新增内部接口
func newServer(network string, address string, p Protocol, sendBufferSize int, mode MsgHandlerModeType, dispatcher PackDispatcherInf, agentHandler SessionHandler) (*TCPServer, error) {
	var err error
	var s *TCPServer
	if network == "ws" || network == "wss" {
		s, err = createWebSocketServer(network, address, p, sendBufferSize, mode, dispatcher, agentHandler)

	} else if network == "tcp" {
		s, err = createTCPServer(network, address, p, sendBufferSize, mode, dispatcher, agentHandler)
	} else {
		return nil, errors.New("unsupport network :" + network)
	}
	return s, err
}

//NewTCPServer create a server with params
func createTCPServer(network, address string, protocol Protocol, sendChanSize int, mode MsgHandlerModeType, dispatcher PackDispatcherInf, agentHandler SessionHandler) (*TCPServer, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	return newTCPServer(network, listener, protocol, sendChanSize, mode, dispatcher, agentHandler), nil
}

//NewClient new client
func newClient(network string, address string, protocol Protocol, sendBufferSize int, mode MsgHandlerModeType) *Client {
	c := &Client{
		networkType:           network,
		address:               address,
		protocol:              protocol,
		sendChannelBufferSize: sendBufferSize,
		mode:                  mode,
	}
	return c

}

//createTCPClient create a client session
func createTCPClient(network, address string, protocol Protocol, sendChanSize int,
	mode MsgHandlerModeType, dispatcher PackDispatcherInf, agent SessionHandler) (*TCPClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	codec, err := protocol.NewSocketCodec(conn)
	if err != nil {
		conn.Close()
		return nil, err
	}
	return newTCPClient(conn, codec, sendChanSize, mode, dispatcher, agent), nil
}

func createWebSocketClient(network, address string, p Protocol, sendChanSize int, mode MsgHandlerModeType, d PackDispatcherInf, agent SessionHandler) (*TCPClient, error) {
	return newWebSocketClient(network, address, p, sendChanSize, mode, d, agent)

}
