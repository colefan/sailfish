package network

//BasePack cls
type BasePack struct {
	cmdType int
	cmdID   int
	data    []byte
	packLen int
	msg     interface{}
	from    uint32
	to      uint32
	uid     uint64
	session *TCPSession
}

//SetPackType interface of PackInf
func (pack *BasePack) SetPackType(packType int) {
	pack.cmdType = packType
}

//GetPackType interface of PackInf
func (pack *BasePack) GetPackType() int {
	return pack.cmdType
}

//SetPackID interface of PackInf
func (pack *BasePack) SetPackID(id int) {
	pack.cmdID = id
}

//GetPackID interface of PackInf
func (pack *BasePack) GetPackID() int {
	return pack.cmdID
}

//GetPackData interface of PackInf
func (pack *BasePack) GetPackData() []byte {
	return pack.data
}

//GetPackLen interface of PackInf
func (pack *BasePack) GetPackLen() int {
	if pack.packLen == 0 && len(pack.data) > 0 {
		return len(pack.data)
	}
	return pack.packLen
}

//SetPackData setter
func (pack *BasePack) SetPackData(data []byte) {
	if pack.data == nil {
		pack.data = make([]byte, 0, len(data))
	}
	pack.data = append(pack.data[0:0], data...)
}

//GetPackBody interface of PackInf
func (pack *BasePack) GetPackBody() interface{} {
	return pack.msg
}

//SetPackBody interface of PackInf
func (pack *BasePack) SetPackBody(msg interface{}) {
	pack.msg = msg
}

//GetFrom getter route from
func (pack *BasePack) GetFrom() uint32 {
	return pack.from
}

//SetFrom setter route from
func (pack *BasePack) SetFrom(from uint32) {
	pack.from = from
}

//GetTO getter route to
func (pack *BasePack) GetTO() uint32 {
	return pack.to
}

//SetTO setter route to
func (pack *BasePack) SetTO(to uint32) {
	pack.to = to
}

func (pack *BasePack) SetTCPSession(session *TCPSession) {
	pack.session = session
}

func (pack *BasePack) GetTCPSession() *TCPSession {
	return pack.session
}
