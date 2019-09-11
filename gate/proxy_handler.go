package gate

import (
	"time"

	"github.com/colefan/sailfish/gate/gatemsg"
	"github.com/colefan/sailfish/log"
	"github.com/colefan/sailfish/network"
	"github.com/colefan/sailfish/network/codec"
)

// ProxyInnerHandler Server request
type ProxyInnerHandler struct {
	pGate *Gate
}

// NewProxyInnerHandler create instance
func NewProxyInnerHandler(s *Gate) *ProxyInnerHandler {
	return &ProxyInnerHandler{
		pGate: s,
	}
}

// SessionOpen session
func (h *ProxyInnerHandler) SessionOpen(session *network.TCPSession) {
	log.Debugf("proxy server session opened, session_id = %d", session.ID())
	proxyNode := NewProxyServerNode()
	proxyNode.Session = session
	proxyNode.SetStatus(ProxyNodeStatusInit)
	session.SetUserData(proxyNode)

}

// SessionClose proxy server node closed
func (h *ProxyInnerHandler) SessionClose(session *network.TCPSession) {
	log.Debugf("proxy server session closed, session_id = %d,status = %d", session.ID(), session.Status())

}

// HandleMsg handleMsg
func (h *ProxyInnerHandler) HandleMsg(pack network.PackInf) {
	cmd := pack.GetCmd()
	switch cmd {
	case int32(gatemsg.MsgTypeGateInnerNode_RegServerReq):
		h.HandleRegisterReq(pack)
	case int32(gatemsg.MsgTypeGateInnerNode_ServerBeatReq):
		h.HandleHeatBeatReq(pack)
	default:
		h.ForwordToClient(pack)
	}

}

// ForwordToClient forword to client
func (h *ProxyInnerHandler) ForwordToClient(pack network.PackInf) {
	uid := pack.GetUID()
	uidType := pack.GetTargetType()
	var clientSession *network.TCPSession
	if uidType == UIDTypeSession {
		clientSession = GetClientSessionMntInst().FindBySessionID(uid)
	} else {
		clientSession = GetClientSessionMntInst().FindByUID(uid)
	}
	if clientSession == nil {
		log.Error("client session is nil forword to client")
		network.FreePack(pack)
		return
	}
	clientSession.WriteMsg(pack)
}

// HandleRegisterReq handle proxy req
func (h *ProxyInnerHandler) HandleRegisterReq(pack network.PackInf) {
	var reqMsg gatemsg.RegServerReq
	if err := codec.ProtobufDecoder(pack, &reqMsg); err != nil {
		log.Error("RegServerReq decode failed:", err)
		network.FreePack(pack)
		return
	}
	sessionData := pack.GetTCPSession().UserData().(*ProxyServerNode)
	if sessionData == nil {
		network.FreePack(pack)
		return
	}

	proxyNode := GetProxyServerStoreInst().Find(reqMsg.ServerId, reqMsg.ServerType)
	if proxyNode != nil {
		if proxyNode.GetStatus() != ProxyNodeStatusNormal {
			proxyNode.Session = pack.GetTCPSession()
			proxyNode.SetStatus(ProxyNodeStatusNormal)
		} else {
			log.Errorf("RegServerReq repeated,server_id = %d, status = %d", reqMsg.ServerId, proxyNode.GetStatus())
			network.FreePack(pack)
			return
		}
	} else {
		sessionData.ServerID = reqMsg.ServerId
		sessionData.ServerType = reqMsg.ServerType
		sessionData.MaxUsers = int(reqMsg.MaxLimit)
		sessionData.IP = reqMsg.Ip
		sessionData.Port = int(reqMsg.Port)
		sessionData.Load = 80
		sessionData.Session = pack.GetTCPSession()
		GetProxyServerStoreInst().RegisterProxyNode(sessionData)
	}

	var respMsg gatemsg.RegServerResp
	respMsg.Time = int32(time.Now().Unix())
	respPack := codec.ProtobufEncoder(int32(gatemsg.MsgTypeGateInnerNode_RegServerResp), &respMsg)
	pack.GetTCPSession().WriteMsg(respPack)
	network.FreePack(pack)
}

// HandleHeatBeatReq handleHeatbeat req
func (h *ProxyInnerHandler) HandleHeatBeatReq(pack network.PackInf) {
	var reqMsg gatemsg.ServerBeatReq
	codec.ProtobufDecoder(pack, &reqMsg)
	if sessionData, ok := pack.GetTCPSession().UserData().(*ProxyServerNode); ok {
		sessionData.LastHeartBeatTime = int32(time.Now().Unix())
	}

	var respMsg gatemsg.ServerBeatResp
	respMsg.ReqTime = reqMsg.ReqTime
	respMsg.ReplyTime = int32(time.Now().Unix())
	respPack := codec.ProtobufEncoder(int32(gatemsg.MsgTypeGateInnerNode_ServerBeatResp), &respMsg)
	pack.GetTCPSession().WriteMsg(respPack)
	network.FreePack(pack)
}
