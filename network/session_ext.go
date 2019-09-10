package network

// SessionHandler session handler interface
type SessionHandler interface {
	SessionOpen(session *TCPSession)
	SessionClose(session *TCPSession)
	HandleMsg(msg PackInf)
}

// DefaultSessionHandler default
type DefaultSessionHandler struct {
	msgMapper map[int32]func(PackInf)
}

// NewDefaultSessionHandler new handler
func NewDefaultSessionHandler() *DefaultSessionHandler {
	h := &DefaultSessionHandler{}
	h.msgMapper = make(map[int32]func(PackInf))
	return h
}

// SessionOpen when session open will be call
func (h *DefaultSessionHandler) SessionOpen(session *TCPSession) {

}

// SessionClose when session close will be call
func (h *DefaultSessionHandler) SessionClose(session *TCPSession) {

}

// RegisterCmdHandler register
func (h *DefaultSessionHandler) RegisterCmdHandler(cmdID int32, f func(PackInf)) {
	h.msgMapper[cmdID] = f
}

// HandleMsg handle msg
func (h *DefaultSessionHandler) HandleMsg(msg PackInf) {
	cmdID := msg.GetCmd()
	if f, ok := h.msgMapper[cmdID]; ok {
		f(msg)
	} else {
		netWarn("no handler find:cmd_id = ", cmdID)
	}

}
