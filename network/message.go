package network

// Message must be impl of PackInf
var (
	MagicNumberDefault     byte = 0x08
	MagicNumberClientInner byte = 0x09
)

const (
	HeaderSize = 30
)

// Header 消息头
type Header struct {
	Magic        byte   //预留信息，默认填写0x08
	Version      byte   //协议版本
	TargetType   byte   //目标类型
	CompressType byte   //压缩类型默认不压缩
	Cmd          int32  //命令号
	CheckCode    uint32 //校验码，预留不填
	UID          uint64 //用户服务端转发时包装
	MsgSeq       uint32 //消息序列号
	SessionID    uint64 //物理会话ID
}

// Message 消息
type Message struct {
	Head      *Header //消息头
	data      []byte  //消息内容
	session   *TCPSession
	timestamp int64 //消息生成时间戳,辅助作用，传输过程丢失
}

// NewMessage create an empty message
func NewMessage() *Message {
	msg := Message{}
	msg.Head = &Header{}
	msg.Head.Version = 0
	msg.Head.Magic = MagicNumberDefault
	msg.Head.TargetType = 0
	msg.Head.CompressType = byte(None)
	msg.Head.Cmd = 0
	msg.Head.CheckCode = 0
	msg.Head.UID = 0
	msg.Head.MsgSeq = 0
	msg.Head.SessionID = 0
	return &msg
}

func (m *Message) resetHeader() {
	m.Head.Magic = 0
	m.Head.TargetType = 0
	m.Head.CompressType = 0
	m.Head.Version = 0
	m.Head.Cmd = 0
	m.Head.CheckCode = 0
	m.Head.UID = 0
	m.Head.MsgSeq = 0
	m.Head.SessionID = 0
}

// Reset 重置消息
func (m *Message) Reset() {
	m.resetHeader()
	m.data = nil
	// if m.data != nil {
	// 	m.data = m.data[:0]
	// }
}

// CheckMessageValidate check the message
func (m *Message) CheckMessageValidate() bool {
	return true
}

// SetData set data
func (c *Message) SetData(data []byte) {
	c.data = data
}

// GetData get data
func (c *Message) GetData() []byte {
	return c.data
}

// GetTCPSession get session
func (c *Message) GetTCPSession() *TCPSession {
	return c.session
}

// SetTCPSession set session
func (c *Message) SetTCPSession(session *TCPSession) {
	c.session = session
}

// GetCheckCode return checkCode
func (c *Message) GetCheckCode() uint32 {
	return c.Head.CheckCode
}

func (c *Message) SetCmd(cmd int32) {
	c.Head.Cmd = cmd
}

func (c *Message) GetCmd() int32 {
	return c.Head.Cmd
}

func (c *Message) GetCompressType() byte {
	return c.Head.CompressType
}

func (c *Message) GetMagic() byte {
	return c.Head.Magic
}

func (c *Message) SetMagic(m byte) {
	c.Head.Magic = m
}

func (c *Message) GetPackLen() int {
	if c.data != nil {
		return HeaderSize + len(c.data)
	}
	return HeaderSize
}

func (c *Message) GetTargetType() byte {
	return c.Head.TargetType
}

func (c *Message) SetTargetType(t byte) {
	c.Head.TargetType = t
}

func (c *Message) SetUID(u uint64) {
	c.Head.UID = u
}

func (c *Message) GetUID() uint64 {
	return c.Head.UID
}

func (c *Message) GetVersion() byte {
	return c.Head.Version
}

func (c *Message) SetVersion(v byte) {
	c.Head.Version = v
}

func (c *Message) GetMsgSeq() uint32 {
	return c.Head.MsgSeq
}

func (c *Message) SetMsgSeq(seq uint32) {
	c.Head.MsgSeq = seq
}

func (c *Message) SetTimestamp(s int64) {
	c.timestamp = s
}

func (c *Message) GetTimestamp() int64 {
	return c.timestamp
}

func (c *Message) SetCompressType(cTyp byte) {
	c.Head.CompressType = cTyp
}

func (c *Message) SetSessionID(sessionID uint64) {
	c.Head.SessionID = sessionID
}

func (c *Message) GetSessionID() uint64 {
	return c.Head.SessionID
}
