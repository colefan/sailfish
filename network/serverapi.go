package network

const (
	NetWorkTypeTCP = "tcp"
	NetWorkTypeWS  = "ws"
	NetWorkTypeWSS = "wss"
)

// NewTCPServer  新建一个TcpServer
func NewTCPServer(address string, msgHandlerMode MsgHandlerModeType, msgProtocol Protocol) *Server {
	//s := newServer()
	s := &Server{}
	s.address = address
	s.networkType = NetWorkTypeTCP
	s.mode = msgHandlerMode
	s.protocol = msgProtocol
	return s
}

// NewWSServer 新建一个WebSocketServer
func NewWSServer(address string, msgHandlerMode MsgHandlerModeType, msgProtocol Protocol) *Server {
	s := &Server{}
	s.address = address
	s.networkType = NetWorkTypeWS
	s.mode = msgHandlerMode
	s.protocol = msgProtocol
	return s
}

// NewWSSServer 新建一个安全websocket
func NewWSSServer(address string, msgHandlerMode MsgHandlerModeType, msgProtocol Protocol, crtFile string, keyFile string) *Server {
	s := &Server{}
	s.address = address
	s.networkType = NetWorkTypeWSS
	s.mode = msgHandlerMode
	s.protocol = msgProtocol
	return s
}
