package gate

import (
	"github.com/colefan/sailfish/gate/gatemsg"
	"github.com/colefan/sailfish/network"
	"github.com/colefan/sailfish/network/codec"
)

// ErrorClientPack error gate client msg
func ErrorClientPack(cmd int32, errorCode int32) network.PackInf {
	var msg gatemsg.GateErrorNt
	msg.CmdId = cmd
	msg.ErrorCode = errorCode
	pack := codec.ProtobufEncoder(int32(gatemsg.MsgTypeGateClient_GateErrorNt), &msg)
	return pack
}
