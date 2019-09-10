package network

import (
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

//ErrorSessionClosed error_closed
var ErrorSessionClosed = errors.New("session is closed")

//ErrorSessionWriteBlocked error_write blocked
var ErrorSessionWriteBlocked = errors.New("session write blocked")

var sTCPSessionID uint64

var (
	MaxReadCountPerSecond int = 100
)

var (
	defaultSendChannelBuffSize int = 100
)

//TCPSession cls
type TCPSession struct {
	sync.Mutex
	id                          uint64
	conn                        net.Conn
	lastActiveTime              int64
	codec                       Codec //编解码器
	closeFlag                   int32
	closeCallback               SessionCloseCallBack
	sendChan                    chan PackInf
	userData                    interface{} //用户数据
	dispatcher                  PackDispatcherInf
	status                      int
	withHandler                 bool
	needCheckPerSecond          bool
	perSecondCount              int
	closeChan                   chan int
	timeout                     int
	wsConn                      *websocket.Conn
	isWebSocket                 bool
	maxReadMsgCountPerSecond    int
	withOutHandlerCloseCallback SessionCloseCallBack
	token                       string
	handler                     SessionHandler
	handleMode                  MsgHandlerModeType
}

//newTCPSession private interface
func newTCPSession(conn net.Conn, codec Codec, sendChanBuff int, dispather PackDispatcherInf, agent SessionHandler) *TCPSession {
	session := &TCPSession{
		id:                 atomic.AddUint64(&sTCPSessionID, 1),
		conn:               conn,
		codec:              codec,
		dispatcher:         dispather,
		perSecondCount:     0,
		needCheckPerSecond: false,
		closeChan:          make(chan int),
		isWebSocket:        false,
		handler:            agent,
	}
	if sendChanBuff <= 0 {
		sendChanBuff = 100
	}
	session.sendChan = make(chan PackInf, sendChanBuff)
	session.maxReadMsgCountPerSecond = MaxReadCountPerSecond
	//	go session.writeMsgLoop()
	//	go session.readMsgLoop()
	return session
}

//newWsSession private interface
func newWsSession(wsconn *websocket.Conn, codec Codec, sendChanBuff int, dispather PackDispatcherInf, agent SessionHandler) *TCPSession {
	session := &TCPSession{
		id:                 atomic.AddUint64(&sTCPSessionID, 1),
		conn:               nil,
		codec:              codec,
		dispatcher:         dispather,
		perSecondCount:     0,
		needCheckPerSecond: false,
		closeChan:          make(chan int),
		wsConn:             wsconn,
		isWebSocket:        true,
		handler:            agent,
	}
	if sendChanBuff <= 0 {
		sendChanBuff = defaultSendChannelBuffSize
	}
	session.sendChan = make(chan PackInf, sendChanBuff)
	session.maxReadMsgCountPerSecond = MaxReadCountPerSecond
	//	go session.writeMsgLoop()
	//	go session.readMsgLoop()
	return session
}

//Start start session reading and writing.
//bWithHandlerRoutine will read outside otherwise read inside
func (session *TCPSession) Start(bWithHandlerRoutine bool) {
	session.withHandler = bWithHandlerRoutine
	if bWithHandlerRoutine {
		if session.handler != nil {
			session.handler.SessionOpen(session)
		}
	}
	go session.writeMsgLoop()
	//	session.withHandler = true
	go session.readMsgLoop()

	// if !bWithHandlerRoutine {
	// 	session.withHandler = false
	// 	go session.readMsgLoop()
	// }
}

func (session *TCPSession) SetCheckPerSecond(b bool) {
	session.needCheckPerSecond = b
}

func (session *TCPSession) SetTimeOut(t int) {
	session.timeout = t
}

//Status getter
func (session *TCPSession) Status() int {
	return session.status
}

//SetStatus setter
func (session *TCPSession) SetStatus(status int) {
	session.status = status
}

//ID return session's id
func (session *TCPSession) ID() uint64 {
	return session.id
}

//RemoteAddress return remote address
func (session *TCPSession) RemoteAddress() string {
	if session.isWebSocket {
		if session.wsConn != nil {
			return session.wsConn.RemoteAddr().String()
		}
	} else {
		if session.conn != nil {
			return session.conn.RemoteAddr().String()
		}
	}
	return ""
}

//LocalAddress return local address
func (session *TCPSession) LocalAddress() string {
	if session.isWebSocket {
		if session.wsConn != nil {
			return session.wsConn.LocalAddr().String()
		}
	} else {
		if session.conn != nil {
			return session.conn.LocalAddr().String()
		}
	}
	return ""
}

//IsClosed return
func (session *TCPSession) IsClosed() bool {
	return atomic.LoadInt32(&session.closeFlag) == 1

}

//Codec return
func (session *TCPSession) Codec() Codec {
	return session.codec
}

//SetUserData setter
func (session *TCPSession) SetUserData(data interface{}) {
	session.userData = data
}

//UserData return userdata
func (session *TCPSession) UserData() interface{} {
	return session.userData
}

//WriteMsg send msg to io
func (session *TCPSession) WriteMsg(msg PackInf) error {
	if session.IsClosed() {
		FreePack(msg)
		return ErrorSessionClosed
	}
	select {
	case session.sendChan <- msg:
		session.lastActiveTime = time.Now().Unix()
		return nil
	default:
		return ErrorSessionWriteBlocked
	}
}

func (session *TCPSession) checkPerSecond() bool {
	if session.needCheckPerSecond {
		if time.Now().Unix() == session.lastActiveTime {
			session.perSecondCount++
		} else {
			session.perSecondCount = 1
		}

		if session.perSecondCount > session.maxReadMsgCountPerSecond {
			netWarn("per second larger than ", session.maxReadMsgCountPerSecond)
			return false
		}

	}

	return true
}

//ReadMsg interface
func (session *TCPSession) ReadMsg() (PackInf, error) {
	if session.isWebSocket {
		return session.readMsgFromWebSocket()
	} else {
		return session.readMsgFromSocket()
	}
}

func (session *TCPSession) readMsgFromSocket() (PackInf, error) {
	if session.timeout > 0 {
		if session.conn != nil {
			session.conn.SetReadDeadline(time.Now().Add(time.Duration(session.timeout) * time.Second))
		}
	}
	pack, err := session.codec.ReceiveMsg()
	if pack != nil {
		if false == session.checkPerSecond() {
			return nil, errors.New("too many request per second ")
		}
		pack.SetTCPSession(session)
		session.lastActiveTime = time.Now().Unix()
		if !session.withHandler {
			session.dispatcher.PostData(pack, session)
		}
		GetTCPServerQos().AddReadPacket(pack.GetPackLen())
	}
	return pack, err

}

func (session *TCPSession) readMsgFromWebSocket() (PackInf, error) {
	if session.timeout > 0 {
		if session.wsConn != nil {
			session.wsConn.SetReadDeadline(time.Now().Add(time.Duration(session.timeout) * time.Second))
		}
	}
	pack, err := session.codec.ReceiveMsg()
	if pack != nil {
		if false == session.checkPerSecond() {
			return nil, errors.New("too many request per second ")
		}
		pack.SetTCPSession(session)
		session.lastActiveTime = time.Now().Unix()
		if !session.withHandler {
			session.dispatcher.PostData(pack, session)
		}
		GetTCPServerQos().AddReadPacket(pack.GetPackLen())
	}
	if err != nil {
		netError("readMsgFromWebSocket error :" + err.Error())
	}
	return pack, err

}

//GetDispatcher getter
func (session *TCPSession) GetDispatcher() PackDispatcherInf {
	return session.dispatcher
}

//Close close session
func (session *TCPSession) Close() {
	//TODO
	if atomic.CompareAndSwapInt32(&session.closeFlag, 0, 1) {
		if session.isWebSocket {
			netDebug("websocket session closed")
			session.wsConn.Close()
		} else {
			netDebug("tcp session closed")
			session.conn.Close()
		}
		//	atomic.StoreInt32(&session.closeFlag, 1)
		//fmt.Println("close id = %d", session.id)
		close(session.closeChan)
		//close(session.sendChan)
		if session.closeCallback != nil {
			session.closeCallback(session)
		}
		//新增事件模式中的回调函数
		if !session.withHandler {
			if session.withOutHandlerCloseCallback != nil {
				session.withOutHandlerCloseCallback(session)
			}
		} else {
			if session.handler != nil {
				session.handler.SessionClose(session)
			}
		}
	}

}

//ForceClose 强迫关闭
func (session *TCPSession) ForceClose() {
	if atomic.CompareAndSwapInt32(&session.closeFlag, 0, 1) {
		if session.isWebSocket {
			netDebug("websocket session closed")
			session.wsConn.Close()
		} else {
			netDebug("tcp session closed")
			session.conn.Close()
		}
		//fmt.Println("force id = %d", session.id)
		//atomic.StoreInt32(&session.closeFlag, 1)
		close(session.closeChan)
	}

}

func (session *TCPSession) writeMsgLoop() {
	//TODO
	for {
		select {
		case msg, ok := <-session.sendChan:

			if !ok {
				netError("session send channel error ")
				return
			}
			//fmt.Println("write msg")
			wLen, err := session.codec.SendMsg(msg)
			FreePack(msg)
			if err != nil {
				netError("session codec write error: %v", err)

				session.Close()
				return
			}
			//session.conn.SetDeadline(time.Now().Add(time.Duration(60) * time.Second))
			GetTCPServerQos().AddWritePacket(wLen)
		case <-session.closeChan:
			return
		}

	}
}

func (session *TCPSession) readMsgLoop() {
	if session.codec == nil {
		return
	}

	// if session.dispatcher != nil {
	// 	GetTCPServerQos().AddReadPacket(msg.GetPackLen())
	// 	msg.SetTCPSession(session)
	// 	session.dispatcher.PostData(msg, session)
	// 	session.lastActiveTime = time.Now().Unix()
	// }

	// sessionOpenMsg := &BasePack{}
	// sessionOpenMsg.SetPackID()
	// if session.

	for {
		if session.timeout > 0 {
			if session.isWebSocket {
				if session.wsConn != nil {
					session.wsConn.SetReadDeadline(time.Now().Add(time.Duration(session.timeout) * time.Second))
				}
				// session.wsConn.SetReadDeadline(time.Now().Add())
			} else {
				if session.conn != nil {
					session.conn.SetReadDeadline(time.Now().Add(time.Duration(session.timeout) * time.Second))
				}
			}
		}
		msg, err := session.codec.ReceiveMsg()
		if err != nil {
			netError("readMsgLoop read error :" + err.Error())
			session.Close()
			return
		}
		GetTCPServerQos().AddReadPacket(msg.GetPackLen())
		msg.SetTCPSession(session)
		if session.withHandler {
			if session.handler != nil {
				session.handler.HandleMsg(msg)
			}
		} else {
			if session.dispatcher != nil {
				session.dispatcher.PostData(msg, session)
				session.lastActiveTime = time.Now().Unix()
			}
		}
	}
}

//SetCloseCallback setter
func (session *TCPSession) SetCloseCallback(f SessionCloseCallBack) {
	session.closeCallback = f
}

//GetLastActiveTime lastactive time
func (session *TCPSession) GetLastActiveTime() int64 {
	return session.lastActiveTime
}

//SetLastActiveTime set lastactive time
func (session *TCPSession) SetLastActiveTime(t int64) {
	session.lastActiveTime = t
}

func (session *TCPSession) SetCloseCallbackForModeEvent(f SessionCloseCallBack) {
	session.withOutHandlerCloseCallback = f
}

//SetToken setter
func (session *TCPSession) SetToken(token string) {
	session.token = token
}

//GetToken getter
func (session *TCPSession) GetToken() string {
	return session.token
}

//SessionCloseCallBack close callback func
type SessionCloseCallBack func(session *TCPSession)
