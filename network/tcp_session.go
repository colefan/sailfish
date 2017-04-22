package network

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

//ErrorSessionClosed error_closed
var ErrorSessionClosed = errors.New("session is closed")

var ErrorSessionWriteBlocked = errors.New("session write blocked")

var sTCPSessionID uint64

//TCPSession cls
type TCPSession struct {
	sync.Mutex
	id             uint64
	conn           net.Conn
	lastActiveTime time.Time
	codec          Codec //编解码器
	closeFlag      int32
	closeCallback  SessionCloseCallBack
	sendChan       chan PackInf
	userData       interface{} //用户数据
	dispatcher     PackDispatcherInf
	status         int
	withHandler    bool
}

//newTCPSession private interface
func newTCPSession(conn net.Conn, codec Codec, sendChanBuff int, dispather PackDispatcherInf) *TCPSession {
	session := &TCPSession{
		id:         atomic.AddUint64(&sTCPSessionID, 1),
		conn:       conn,
		codec:      codec,
		dispatcher: dispather,
	}
	if sendChanBuff <= 0 {
		sendChanBuff = 100
	}
	session.sendChan = make(chan PackInf, sendChanBuff)
	//	go session.writeMsgLoop()
	//	go session.readMsgLoop()
	return session
}

//Start start session reading and writing.
//bWithHandlerRoutine will read outside otherwise read inside
func (session *TCPSession) Start(bWithHandlerRoutine bool) {
	go session.writeMsgLoop()
	session.withHandler = true
	if !bWithHandlerRoutine {
		session.withHandler = false
		go session.readMsgLoop()
	}
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
	if session.conn != nil {
		return session.conn.RemoteAddr().String()
	}
	return ""
}

//LocalAddress return local address
func (session *TCPSession) LocalAddress() string {
	if session.conn != nil {
		return session.conn.LocalAddr().String()
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
		return ErrorSessionClosed
	}
	select {
	case session.sendChan <- msg:
		return nil
	default:
		return ErrorSessionWriteBlocked
	}
}

//ReadMsg interface
func (session *TCPSession) ReadMsg() (PackInf, error) {

	pack, err := session.codec.ReceiveMsg()
	if pack != nil {
		pack.SetTCPSession(session)
		if !session.withHandler {
			session.dispatcher.PostData(pack, session)
		}
		GetTCPServerQos().AddReadPacket(pack.GetPackLen())
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
	session.conn.Close()
	atomic.StoreInt32(&session.closeFlag, 1)
	if session.closeCallback != nil {
		session.closeCallback(session)
	}
}

//ForceClose 强迫关闭
func (session *TCPSession) ForceClose() {
	session.conn.Close()
}

func (session *TCPSession) writeMsgLoop() {
	//TODO
	for {
		select {
		case msg, ok := <-session.sendChan:

			if !ok {
				fmt.Println("channel error")
				return
			}
			//fmt.Println("write msg")
			wLen, err := session.codec.SendMsg(msg)
			if err != nil {
				fmt.Println("write error")
				session.Close()
				return
			}
			GetTCPServerQos().AddWritePacket(wLen)

		}

	}
}

func (session *TCPSession) readMsgLoop() {
	if session.codec == nil {
		return
	}

	for {
		msg, err := session.codec.ReceiveMsg()
		if err != nil {
			session.Close()
			return
		}
		if session.dispatcher != nil {
			GetTCPServerQos().AddReadPacket(msg.GetPackLen())
			msg.SetTCPSession(session)
			session.dispatcher.PostData(msg, session)
		}
	}
}

//SetCloseCallback setter
func (session *TCPSession) SetCloseCallback(f SessionCloseCallBack) {
	session.closeCallback = f
}

//SessionCloseCallBack close callback func
type SessionCloseCallBack func(session *TCPSession)
