package network

// NewTCPClient new tcp client
func NewTCPClient(address string, msgProtocol Protocol) *Client {
	c := newClient(NetWorkTypeTCP, address, msgProtocol, 100, MsgHandleModeHandler)
	return c
}

// NewWSClient new websocket client
func NewWSClient(address string, msgProtocol Protocol) *Client {
	c := newClient(NetWorkTypeWS, address, msgProtocol, 100, MsgHandleModeHandler)
	return c
}

// NewWSSClient new websocket client support ssl
func NewWSSClient(address string, msgProtocol Protocol) *Client {
	c := newClient(NetWorkTypeWSS, address, msgProtocol, 100, MsgHandleModeHandler)
	return c
}
