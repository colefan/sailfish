package network

//PackInf interface of pack
type PackInf interface {
	GetMagic() byte
	GetVersion() byte
	GetTargetType() byte
	GetCompressType() byte
	GetCmd() int32
	SetCmd(cmd int32)
	GetCheckCode() uint32
	GetUID() uint64
	SetTCPSession(session *TCPSession)
	GetTCPSession() *TCPSession
	SetUID(uid uint64)
	SetMagic(m byte)
	Reset()
	CheckMessageValidate() bool
	GetPackLen() int
	GetData() []byte
	SetData([]byte)
}
