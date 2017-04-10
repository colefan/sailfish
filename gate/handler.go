package gate

import "github.com/colefan/sailfish/network"

type SessionHandler interface {
	SessionOpen(session *network.TCPSession)
	HandleMessage(msg network.PackInf, session *network.TCPSession)
	SessionClose(session *network.TCPSession)
}
