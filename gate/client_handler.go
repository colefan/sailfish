package gate

import "github.com/colefan/sailfish/network"

//ClientHandler handler for real client session
type ClientHandler struct {
}

func (h *ClientHandler) SessionOpen(session *network.TCPSession) {
	//TODO
}

func (h *ClientHandler) SessionClose(session *network.TCPSession) {
	//TODO
}

func (h *ClientHandler) HandleMessage(pack network.PackInf, session *network.TCPSession) {
	//TODO
}
