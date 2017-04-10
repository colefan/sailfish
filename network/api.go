package network

import "net"
import "time"

//PackInf interface of pack
type PackInf interface {
	GetPackType() int
	SetPackType(packType int)
	GetPackID() int
	SetPackID(id int)
	GetPackData() []byte
	SetPackData(data []byte)
	GetPackBody() interface{}
	SetPackBody(msg interface{})
	GetPackLen() int
	SetTCPSession(session *TCPSession)
	GetTCPSession() *TCPSession
}

//Codec 编解码接口
type Codec interface {
	//关闭编解码器
	//ReceiveMsg
	ReceiveMsg() (PackInf, error)
	//SendMsg return len,error
	SendMsg(msg PackInf) (int, error)
	//SendRawMsg return len,error
	SendRawMsg(data []byte) (int, error)
}

//Protocol 服务器遵循的协议接口
type Protocol interface {
	NewCodec(conn net.Conn) (Codec, error)
}

//NewTCPServer create a server with params
func NewTCPServer(network, address string, protocol Protocol, sendChanSize int) (*TCPServer, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}
	return newTCPServer(listener, protocol, sendChanSize), nil
}

//NewTCPClient create a client session
func NewTCPClient(network, address string, protocol Protocol, sendChanSize int) (*TCPClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	codec, err := protocol.NewCodec(conn)
	if err != nil {
		conn.Close()
		return nil, err
	}
	return newTCPClient(conn, codec, sendChanSize), nil
}

//NewClientConntectTimeout create a client session with timeout
func NewTCPClientTimeout(network, address string, timeout time.Duration, protocol Protocol, sendChanSize int) (*TCPClient, error) {
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return nil, err
	}
	codec, err := protocol.NewCodec(conn)
	if err != nil {
		conn.Close()
		return nil, err
	}
	return newTCPClient(conn, codec, sendChanSize), nil
}
