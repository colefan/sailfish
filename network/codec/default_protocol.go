package codec

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
	"net"
	"strconv"

	"github.com/colefan/sailfish/network"
	"github.com/gorilla/websocket"
)

// DefaultProtocol 默认消息协议
type DefaultProtocol struct {
	useLittleEndian bool
	headWord        int
	maxReadSize     int
	maxWriteSize    int
}

// NewDefaultProtocol new default protocol
func NewDefaultProtocol(useLittleEndian bool, headWord int) *DefaultProtocol {
	p := &DefaultProtocol{
		useLittleEndian: useLittleEndian,
		headWord:        headWord,
	}
	if headWord < 1 {
		headWord = DefaultFixHeadWord
	}
	switch headWord {
	case 1:
		p.maxReadSize = math.MaxUint8
	case 2:
		p.maxReadSize = math.MaxUint16
	case 4:
		p.maxReadSize = math.MaxUint32
	}
	p.maxWriteSize = p.maxReadSize
	return p
}

// DefaultSocketCodec default codec
type DefaultSocketCodec struct {
	p       *DefaultProtocol
	conn    net.Conn
	head    [4]byte
	headBuf []byte
}

// DefaultWebsocketCodec websocket codec
type DefaultWebsocketCodec struct {
	p       *DefaultProtocol
	wsConn  *websocket.Conn
	head    [4]byte
	headBuf []byte
}

// ReceiveMsg 收到消息进行解码操作
func (c *DefaultSocketCodec) ReceiveMsg() (network.PackInf, error) {
	if c.conn == nil {
		return nil, ErrFixHeadConnIsNull
	}

	if _, err := io.ReadFull(c.conn, c.headBuf); err != nil {
		return nil, err
	}
	nLen := FixHeadDecoder(c.p.useLittleEndian, c.p.headWord, c.headBuf)
	if nLen > c.p.maxReadSize {
		return nil, ErrFixHeadTooLargePacket
	}

	if nLen < network.HeaderSize {
		return nil, errors.New("packLen is less than header size :" + strconv.Itoa(network.HeaderSize))
	}

	buf := make([]byte, nLen, nLen)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return nil, err
	}

	pack := network.GetPooledPack()
	msg := pack.(*network.Message)
	msg.Head.Magic = buf[0]
	msg.Head.Version = buf[1]
	msg.Head.TargetType = buf[2]
	msg.Head.CompressType = buf[3]
	if c.p.useLittleEndian {
		msg.Head.Cmd = int32(binary.LittleEndian.Uint32(buf[4:8]))
		msg.Head.CheckCode = binary.LittleEndian.Uint32(buf[8:12])
		msg.Head.UID = binary.LittleEndian.Uint64(buf[12:20])
	} else {
		msg.Head.Cmd = int32(binary.BigEndian.Uint32(buf[4:8]))
		msg.Head.CheckCode = binary.BigEndian.Uint32(buf[8:12])
		msg.Head.UID = binary.BigEndian.Uint64(buf[12:20])

	}
	msg.SetData(buf[20:])
	return pack, nil

}

// SendMsg send msg
func (c *DefaultSocketCodec) SendMsg(msg network.PackInf) (int, error) {
	if c.conn == nil {
		return 0, ErrFixHeadConnIsNull
	}

	nLen := network.HeaderSize
	bodyBytes := msg.GetData()
	if bodyBytes != nil {
		nLen += len(bodyBytes)
	}
	msgHeadData := make([]byte, network.HeaderSize, network.HeaderSize)
	msgHeadData[0] = msg.GetMagic()
	msgHeadData[1] = msg.GetVersion()
	msgHeadData[2] = msg.GetTargetType()
	msgHeadData[3] = msg.GetCompressType()
	if c.p.useLittleEndian {
		binary.LittleEndian.PutUint32(msgHeadData[4:], uint32(msg.GetCmd()))
		binary.LittleEndian.PutUint32(msgHeadData[8:], uint32(msg.GetCheckCode()))
		binary.LittleEndian.PutUint64(msgHeadData[12:], msg.GetUID())
	} else {
		binary.BigEndian.PutUint32(msgHeadData[4:], uint32(msg.GetCmd()))
		binary.BigEndian.PutUint32(msgHeadData[8:], uint32(msg.GetCheckCode()))
		binary.BigEndian.PutUint64(msgHeadData[12:], msg.GetUID())
	}
	if nLen > c.p.maxWriteSize {
		return 0, ErrFixHeadTooLargePacket
	}
	encodedData := make([]byte, c.p.headWord, c.p.headWord+nLen)
	FixHeadEncoder(c.p.useLittleEndian, c.p.headWord, encodedData, nLen)
	encodedData = append(encodedData, msgHeadData...)
	if bodyBytes != nil {
		encodedData = append(encodedData, bodyBytes...)
	}
	if n, err := c.conn.Write(encodedData); err != nil {
		return 0, err
	} else {
		return n, nil
	}

	return 0, ErrFixHeadUnknownMsgInterface

}

// NewSocketCodec new socket codec
func (p *DefaultProtocol) NewSocketCodec(conn net.Conn) (network.Codec, error) {
	c := &DefaultSocketCodec{
		p:    p,
		conn: conn,
	}
	c.headBuf = c.head[0:c.p.headWord]
	return c, nil
}

// ReceiveMsg receive msg
func (c *DefaultWebsocketCodec) ReceiveMsg() (network.PackInf, error) {
	if c.wsConn == nil {
		return nil, ErrFixHeadConnIsNull
	}
	_, p, err := c.wsConn.ReadMessage()
	if err != nil {
		return nil, err
	}
	if len(p) < c.p.headWord {
		return nil, ErrFixHeadSizeIsNotEnough
	}
	// fmt.Printf("recv p = %v,len(p) = %v \n", p, len(p))
	// realLen := len(p) - c.p.headWord
	copy(c.headBuf, p[0:c.p.headWord])
	nLen := FixHeadDecoder(c.p.useLittleEndian, c.p.headWord, c.headBuf)

	// fmt.Printf(" recv p = %v,nLen = %v\n", p, nLen)

	if nLen < network.HeaderSize {
		return nil, errors.New("head word size is less than HeaderSize")
	}

	if nLen > c.p.maxReadSize {
		return nil, ErrFixHeadTooLargePacket
	}
	p = p[c.p.headWord:]
	// fmt.Printf("cut headword  recv p = %v,nLen = %v\n", p, nLen)

	if len(p) < nLen {
		return nil, errors.New("msg data is less than protocol length")
	}

	if len(p) < network.HeaderSize {
		return nil, errors.New("msg data is less than protocol msg headSize")
	}

	p = p[0:nLen]
	// fmt.Printf("after p = %v,nLen = %d \n", p, nLen)
	pack := network.GetPooledPack()
	msg := pack.(*network.Message)
	msg.Head.Magic = p[0]
	msg.Head.Version = p[1]
	msg.Head.TargetType = p[2]
	msg.Head.CompressType = p[3]
	if c.p.useLittleEndian {
		msg.Head.Cmd = int32(binary.LittleEndian.Uint32(p[4:]))
		msg.Head.CheckCode = binary.LittleEndian.Uint32(p[8:])
		msg.Head.UID = binary.LittleEndian.Uint64(p[12:])
	} else {
		msg.Head.Cmd = int32(binary.BigEndian.Uint32(p[4:]))
		msg.Head.CheckCode = binary.BigEndian.Uint32(p[8:])
		msg.Head.UID = binary.BigEndian.Uint64(p[12:])
	}
	msg.SetData(p[20:])
	// pack.SetPackBody(string(p))

	return pack, nil
}

// SendMsg send msg
func (c *DefaultWebsocketCodec) SendMsg(msg network.PackInf) (int, error) {
	if c.wsConn == nil {
		return 0, ErrFixHeadConnIsNull
	}

	msgHeadData := make([]byte, network.HeaderSize, network.HeaderSize)
	msgHeadData[0] = msg.GetMagic()
	msgHeadData[1] = msg.GetVersion()
	msgHeadData[2] = msg.GetTargetType()
	msgHeadData[3] = msg.GetCompressType()
	if c.p.useLittleEndian {
		binary.LittleEndian.PutUint32(msgHeadData[4:], uint32(msg.GetCmd()))
		binary.LittleEndian.PutUint32(msgHeadData[8:], msg.GetCheckCode())
		binary.LittleEndian.PutUint64(msgHeadData[12:], msg.GetUID())

	} else {
		binary.BigEndian.PutUint32(msgHeadData[4:], uint32(msg.GetCmd()))
		binary.BigEndian.PutUint32(msgHeadData[8:], msg.GetCheckCode())
		binary.BigEndian.PutUint64(msgHeadData[12:], msg.GetUID())
	}
	nLen := network.HeaderSize
	bodyData := msg.GetData()
	if bodyData != nil && len(bodyData) > 0 {
		nLen += len(bodyData)
	}

	if nLen > c.p.maxWriteSize {
		return 0, ErrFixHeadTooLargePacket
	}
	encodedData := make([]byte, c.p.headWord, c.p.headWord+nLen)
	FixHeadEncoder(c.p.useLittleEndian, c.p.headWord, encodedData, nLen)
	encodedData = append(encodedData, msgHeadData...)
	if bodyData != nil {
		encodedData = append(encodedData, bodyData...)
	}
	// fmt.Printf("cmd_id = 0x%0x, data = %v\n", msg.GetCmd(), encodedData)
	if err := c.wsConn.WriteMessage(websocket.BinaryMessage, encodedData); err != nil {
		return 0, err
	} else {
		return nLen + c.p.headWord, nil
	}

	return 0, nil
}

// NewWsSocketCodec new socket codec
func (p *DefaultProtocol) NewWsSocketCodec(wsconn *websocket.Conn) (network.Codec, error) {
	c := &DefaultWebsocketCodec{
		p:      p,
		wsConn: wsconn,
	}
	c.headBuf = c.head[0:c.p.headWord]
	return c, nil
}
