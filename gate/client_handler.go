package gate

import (
	"strings"
	"time"

	"github.com/colefan/sailfish/log"
	"github.com/colefan/sailfish/network/codec"

	"github.com/colefan/sailfish/gate/gatemsg"
	"github.com/colefan/sailfish/network"
)

//ClientHandler handler for real client session
type ClientHandler struct {
	pGate *Gate
}

// NewClientHandler new client Handler
func NewClientHandler(pGate *Gate) *ClientHandler {
	h := &ClientHandler{pGate: pGate}
	return h
}

// SessionOpen 打开一个会话
func (h *ClientHandler) SessionOpen(session *network.TCPSession) {
	userData := NewClientUserData()
	remoteAddress := session.RemoteAddress()
	if strings.Index(remoteAddress, ":") > 0 {
		userData.UserIP = remoteAddress[0:strings.Index(remoteAddress, ":")]
	} else {
		userData.UserIP = remoteAddress
	}
	log.Debugf("client session opened,session_id = %d,user_ip = %v", session.ID(), userData.UserIP)
	session.SetUserData(userData)
	GetClientSessionMntInst().AddSession(session)

}

// SessionClose 关闭一个会话
func (h *ClientHandler) SessionClose(session *network.TCPSession) {
	log.Debugf("client session closed, session_id = %d ,session_status = %d", session.ID(), session.Status())
	//通知相关的代理服务下线
	if user, ok := session.UserData().(*ClientUserData); ok {
		if user.ProxyNodeList != nil && user.Status > 0 {
			if false == user.KickOffWhenRepeatedLogin {
				log.Debugf("client session closed, session_id = %d ,user.Status = %d", session.ID(), user.Status)
				// log.Debugf("ProxyNodeList.Len = %d, values = %v", len(user.ProxyNodeList), user.ProxyNodeList)
				for k, v := range user.ProxyNodeList {
					if k == UIDTypeUser {
						if v != nil {
							notifyClientCloseReq(k, v, user, session.ID())
						}
					}
				}
			}
		}

	}
	// }

	GetClientSessionMntInst().DelSession(session)
}

// HandleMsg 处理消息请求
func (h *ClientHandler) HandleMsg(pack network.PackInf) {
	log.Debugf("gate recv client cmd = 0x%0x", pack.GetCmd())
	switch pack.GetCmd() {
	case int32(gatemsg.MsgTypeGateClient_ClientHandShakeReq):
		h.HandleClientShakeReq(pack)
	case int32(gatemsg.MsgTypeGateClient_ClientBeatReq):
		h.HandleClientBeatReq(pack)
	default:
		log.Debugf("forward to proxyNode = %d, cmd = 0x%0x", pack.GetTargetType(), pack.GetCmd())
		h.ForwordToProxyNode(pack)
	}

}

// HandleClientShakeReq logic
func (h *ClientHandler) HandleClientShakeReq(pack network.PackInf) {
	var reqMsg gatemsg.ClientHandShakeReq
	err := codec.ProtobufDecoder(pack, &reqMsg)
	if err != nil {
		log.Error("ClientHandShakeReq decode failed:", err)
		ntpack := ErrorClientPack(pack.GetCmd(), int32(gatemsg.ErrorCode_UnmarshalFailed))
		pack.GetTCPSession().WriteMsg(ntpack)
		network.FreePack(pack)
		return
	}
	if user, ok := pack.GetTCPSession().UserData().(*ClientUserData); ok {
		if user.Status == ClientStatusInit {
			user.ChannleID = reqMsg.ClientChannel
			user.ClientType = reqMsg.ClientType
			user.Version = reqMsg.ClientVersion
			user.ClientAuth = reqMsg.ClientAuth
			user.Status = ClientStatusHandShaked
			var respMsg gatemsg.ClientHandShakeResp
			respMsg.ServerTimeSecond = int32(time.Now().Unix())
			respMsg.Token = ""
			respPack := codec.ProtobufEncoder(int32(gatemsg.MsgTypeGateClient_ClientHandShakeResp), &respMsg)
			// just for debug ~
			// message := respPack.(*network.Message)
			// message.SetMagic(0x08)
			// message.SetTargetType(0x01)
			// message.SetUID(1111)
			// message.Head.CompressType = 2
			// message.Head.Version = 2
			// message.Head.CheckCode = 88
			// ~
			pack.GetTCPSession().WriteMsg(respPack)
			network.FreePack(pack)
			return
		}
	}

	ntpack := ErrorClientPack(pack.GetCmd(), int32(gatemsg.ErrorCode_HandShakeFailed))
	pack.GetTCPSession().WriteMsg(ntpack)
	network.FreePack(pack)
}

// HandleClientBeatReq logic
func (h *ClientHandler) HandleClientBeatReq(pack network.PackInf) {
	var reqMsg gatemsg.ClientBeatReq
	if err := codec.ProtobufDecoder(pack, &reqMsg); err != nil {
		ntpack := ErrorClientPack(pack.GetCmd(), int32(gatemsg.ErrorCode_UnmarshalFailed))
		pack.GetTCPSession().WriteMsg(ntpack)
		network.FreePack(pack)
		return
	}
	if user, ok := pack.GetTCPSession().UserData().(*ClientUserData); ok {
		user.LastBeatTime = time.Now().Unix()
	}

	var respMsg gatemsg.ClientBeatResp
	respMsg.ReqTime = reqMsg.ReqTime
	respMsg.ReplyTime = int32(time.Now().Unix())
	respPack := codec.ProtobufEncoder(int32(gatemsg.MsgTypeGateClient_ClientBeatResp), &respMsg)
	pack.GetTCPSession().WriteMsg(respPack)
	network.FreePack(pack)
}

// ForwordToProxyNode logic
func (h *ClientHandler) ForwordToProxyNode(pack network.PackInf) {
	session := pack.GetTCPSession()
	if user, ok := session.UserData().(*ClientUserData); ok {
		if user.Status < ClientStatusHandShaked {
			errPack := ErrorClientPack(pack.GetCmd(), int32(gatemsg.ErrorCode_HandShakeFailed))
			pack.GetTCPSession().WriteMsg(errPack)
			network.FreePack(pack)
			return
		}

		targetServer := int32(pack.GetTargetType())
		if pack.GetUID() != 0 {
			if user.UID == 0 {
				GetClientSessionMntInst().AddSessionByUID(pack.GetUID(), pack.GetTCPSession())
			}
			user.UID = pack.GetUID()
		}

		if user.UID != 0 {
			pack.SetUID(user.UID)
			// pack.SetMagic(0x08 + 0x01)
			pack.SetTargetType(UIDTypeUser)
		} else {
			// pack.SetMagic(0x08)
			pack.SetTargetType(UIDTypeSession)
			pack.SetUID(session.ID())
		}
		if node, ok2 := user.ProxyNodeList[targetServer]; ok2 {
			// node.Session
			if node.Session != nil {
				node.Session.WriteMsg(pack)
				return
			}
		} else {
			tmpNode := GetProxyServerStoreInst().MatchProxyServer(targetServer)
			if tmpNode != nil {
				user.ProxyNodeList[int32(pack.GetTargetType())] = tmpNode
				if tmpNode.Session != nil {

					tmpNode.Session.WriteMsg(pack)
					return
				}
			} else {
				errPack := ErrorClientPack(pack.GetCmd(), int32(gatemsg.ErrorCode_NoTargetServer))
				pack.GetTCPSession().WriteMsg(errPack)
			}
		}
	}

	network.FreePack(pack)

}

func notifyClientCloseReq(serverType int32, serverInfo *ProxyServerNode, user *ClientUserData, clientSessionId uint64) {
	if serverInfo == nil {
		return
	}

	if serverInfo.Session != nil {

		var reqMsg gatemsg.ClientCloseReq
		reqMsg.SessionId = clientSessionId
		reqMsg.Uid = user.UID

		reqMsgPack := codec.ProtobufEncoder(int32(gatemsg.MsgTypeGateInnerNode_ClientCloseReq), &reqMsg)

		if user.UID != 0 {
			reqMsgPack.SetUID(user.UID)
			reqMsgPack.SetTargetType(UIDTypeUser)
		} else {
			reqMsgPack.SetTargetType(UIDTypeSession)
			reqMsgPack.SetUID(serverInfo.Session.ID())
		}
		log.Debugf("notifyClientCloseReq sever_id = %d,user_id = %d", serverInfo.ServerID, user.UID)
		err := serverInfo.Session.WriteMsg(reqMsgPack)
		if err != nil {
			log.Debugf("notifyClientCloseReq error :%v", err)
			tmpServerInfo := GetProxyServerStoreInst().Find(serverInfo.ServerID, serverInfo.ServerType)
			if tmpServerInfo != nil && tmpServerInfo.Session != nil {
				var reqMsg2 gatemsg.ClientCloseReq
				reqMsg2.SessionId = clientSessionId
				reqMsg2.Uid = user.UID

				pack2 := codec.ProtobufEncoder(int32(gatemsg.MsgTypeGateInnerNode_ClientCloseReq), &reqMsg2)

				log.Debugf("notifyClientCloseReq 2 sever_id = %d,user_id = %d", tmpServerInfo.ServerID, user.UID)
				if user.UID != 0 {
					pack2.SetUID(user.UID)
					pack2.SetTargetType(UIDTypeUser)
				} else {
					pack2.SetTargetType(UIDTypeSession)
					pack2.SetUID(tmpServerInfo.Session.ID())
				}
				tmpServerInfo.Session.WriteMsg(pack2)
			}
		}
	}

}
