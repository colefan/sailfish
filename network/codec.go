package network

import (
	"net"

	"github.com/gorilla/websocket"
)

//Codec 编解码接口
type Codec interface {
	//关闭编解码器
	//ReceiveMsg
	ReceiveMsg() (PackInf, error)
	//SendMsg return len,error
	SendMsg(msg PackInf) (int, error)
}

//Protocol 服务器遵循的协议接口
type Protocol interface {
	NewSocketCodec(conn net.Conn) (Codec, error)
	NewWsSocketCodec(wsconn *websocket.Conn) (Codec, error)
}
