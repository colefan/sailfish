package codec

import (
	"encoding/json"
	"errors"
	"net"
	"reflect"

	"github.com/colefan/sailfish/network"
)

//JSONProtocol json协议
type JSONProtocol struct {
	types map[string]reflect.Type
	names map[reflect.Type]string
}

//NewJSONProtocol create json protocol
func NewJSONProtocol() *JSONProtocol {
	return &JSONProtocol{
		types: make(map[string]reflect.Type),
		names: make(map[reflect.Type]string),
	}
}

//Register register  msg to the protocol
func (p *JSONProtocol) Register(msg interface{}) {
	t := reflect.TypeOf(msg)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	name := t.PkgPath() + "/" + t.Name()
	p.types[name] = t
	p.names[t] = name
}

//RegisterWithName register msg to protocol with msgname
func (p *JSONProtocol) RegisterWithName(msg interface{}, msgName string) {
	t := reflect.TypeOf(msg)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if msgName == "" {
		msgName = t.PkgPath() + "/" + t.Name()
	}
	p.types[msgName] = t
	p.names[t] = msgName
}

//JSONMsg json msg pack
type jsonMsg struct {
	Head string
	Body interface{}
}

type jsonIn struct {
	Head string
	Body json.RawMessage
}

//NewCodec 生成JSON编解码
func (p *JSONProtocol) NewCodec(conn net.Conn) (network.Codec, error) {
	c := newJSONCodec(p, conn)
	return c, nil
}

func newJSONCodec(p *JSONProtocol, conn net.Conn) network.Codec {
	c := &JSONCodec{
		p:    p,
		conn: conn,
	}
	c.decoder = json.NewDecoder(conn)
	c.encoder = json.NewEncoder(conn)
	return c

}

//JSONCodec JSON编解码器
type JSONCodec struct {
	p       *JSONProtocol
	conn    net.Conn
	decoder *json.Decoder
	encoder *json.Encoder
}

//ReceiveMsg JSONCodec Receive interface
func (c *JSONCodec) ReceiveMsg() (network.PackInf, error) {
	var msg jsonIn
	err := c.decoder.Decode(&msg)
	if err != nil {
		return nil, err
	}
	var body interface{}
	if msg.Head != "" {
		if t, ok := c.p.types[msg.Head]; ok {
			body = reflect.New(t).Interface()
		}
	}
	json.Unmarshal(msg.Body, body)
	pack := new(network.BasePack)
	pack.SetPackBody(body)
	return pack, nil
}

//SendMsg interface of codec
func (c *JSONCodec) SendMsg(pack network.PackInf) (int, error) {
	msg := pack.GetPackBody()
	t := reflect.TypeOf(msg)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	var out jsonMsg
	if name, ok := c.p.names[t]; ok {
		out.Head = name
	} else {
		return 0, errors.New("unregisted json msg")
	}
	out.Body = msg
	err := c.encoder.Encode(out)

	return 0, err
}

//SendRawMsg interface of codec
func (c *JSONCodec) SendRawMsg(data []byte) (int, error) {
	if c.conn == nil {
		return 0, ErrFixHeadConnIsNull
	}
	if data != nil {
		_, err := c.conn.Write(data)
		return len(data), err
	}
	return 0, ErrFixHeadUnknownMsgInterface

}
