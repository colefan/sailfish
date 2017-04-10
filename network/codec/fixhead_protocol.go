package codec

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
	"net"

	"github.com/colefan/sailfish/network"
)

//ErrFixHeadConnIsNull self define error
var ErrFixHeadConnIsNull = errors.New("FixHeadProtocol:conn is nil")

//ErrFixHeadTooLargePacket self define error
var ErrFixHeadTooLargePacket = errors.New("FixHeadProtocol:too large packet")

//ErrFixHeadUnknownMsgInterface self define error
var ErrFixHeadUnknownMsgInterface = errors.New("FixHeadProtocol:unknown msg interface")

// defaultFixHeadSize page const
var defaultFixHeadWord = 4

//FixHeadProtocol fixhead for a protocol
type FixHeadProtocol struct {
	headWord     int
	maxReadSize  int
	maxWriteSize int
	headDecoder  func([]byte) int
	headEncode   func([]byte, int)
}

// FixHeadCodec cls
type FixHeadCodec struct {
	conn         net.Conn
	head         [8]byte
	headBuf      []byte
	headWord     int
	maxReadSize  int
	maxWriteSize int
	p            *FixHeadProtocol
}

// NewFixHeadProtocol interface for server
func NewFixHeadProtocol(headWord int) *FixHeadProtocol {
	p := &FixHeadProtocol{}
	if headWord < 0 {
		headWord = defaultFixHeadWord
	}
	maxMsgSize := 0
	switch headWord {
	case 1:
		maxMsgSize = math.MaxUint8
		p.headDecoder = func(b []byte) int {
			return int(b[0])
		}
		p.headEncode = func(b []byte, size int) {
			b[0] = byte(size)
		}

	case 2:
		maxMsgSize = math.MaxUint16
		p.headDecoder = func(b []byte) int {
			return int(binary.LittleEndian.Uint16(b))
		}
		p.headEncode = func(b []byte, size int) {
			binary.LittleEndian.PutUint16(b, uint16(size))
		}
	case 4:
		maxMsgSize = math.MaxUint32
		p.headDecoder = func(b []byte) int {
			return int(binary.LittleEndian.Uint32(b))
		}
		p.headEncode = func(b []byte, size int) {
			binary.LittleEndian.PutUint32(b, uint32(size))
		}
	default:
		panic("FixHeadProtocol: unsupported head size")
	}
	p.headWord = headWord
	p.maxReadSize = maxMsgSize
	p.maxWriteSize = maxMsgSize
	return p
}

// NewCodec interface for protocol
func (p *FixHeadProtocol) NewCodec(conn net.Conn) (network.Codec, error) {
	c := &FixHeadCodec{}
	c.conn = conn
	c.headWord = p.headWord
	c.maxReadSize = p.maxReadSize
	c.maxWriteSize = p.maxReadSize
	c.headBuf = c.head[:p.headWord]
	c.p = p
	return c, nil
}

//Receive interface of Codec
func (c *FixHeadCodec) ReceiveMsg() (network.PackInf, error) {
	if c.conn == nil {
		return nil, ErrFixHeadConnIsNull
	}
	if _, err := io.ReadFull(c.conn, c.headBuf); err != nil {
		return nil, err
	}
	msgLen := c.p.headDecoder(c.headBuf)
	if msgLen > c.maxReadSize {
		return nil, ErrFixHeadTooLargePacket
	}
	buf := make([]byte, msgLen, msgLen)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return nil, err
	}
	pack := new(network.BasePack)
	pack.SetPackData(buf)
	return pack, nil
}

//Send interface of Send
func (c *FixHeadCodec) SendMsg(msg network.PackInf) (int, error) {
	if c.conn == nil {
		return 0, ErrFixHeadConnIsNull
	}
	data := msg.GetPackData()
	wLen := len(data)

	encodedData := make([]byte, c.headWord, c.headWord+wLen)
	c.p.headEncode(encodedData, wLen)
	encodedData = append(encodedData, data...)
	//fmt.Printf("write encodedData %d\n", wLen)
	if data != nil {
		_, err := c.conn.Write(encodedData)
		return wLen + c.headWord, err
	}

	return 0, ErrFixHeadUnknownMsgInterface
}

func (c *FixHeadCodec) SendRawMsg(data []byte) (int, error) {
	if c.conn == nil {
		return 0, ErrFixHeadConnIsNull
	}
	wLen := len(data)
	encodedData := make([]byte, c.headWord, c.headWord+wLen)
	c.p.headEncode(encodedData, wLen)
	encodedData = append(encodedData, data...)

	if data != nil {
		_, err := c.conn.Write(encodedData)
		return len(encodedData), err
	}
	return 0, ErrFixHeadUnknownMsgInterface

}
