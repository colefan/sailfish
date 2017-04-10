package gate

import "github.com/colefan/sailfish/network"

//Handle Server request
type ServerHandler struct {
}

func (h *ServerHandler) SessionOpen(session *network.TCPSession) {
	//	gLog.Debug("session opend %d", session.ID())
	//TODO
}

func (h *ServerHandler) SessionClose(session *network.TCPSession) {
	// gLog.Debug("session closed %d", session.ID())
	//TODO
}

func (h *ServerHandler) HandleMessage(pack network.PackInf, session *network.TCPSession) {
	//TODO
	// gLog.Debug("Session received a mesesage %d", session.ID())
}
