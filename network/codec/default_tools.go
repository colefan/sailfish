package codec

import (
	"sailfish/network"

	"github.com/golang/protobuf/proto"
)

// ProtobufDecoder protobufDecoder
func ProtobufDecoder(pack network.PackInf, pb proto.Message) error {
	return proto.Unmarshal(pack.GetData(), pb)
}

// ProtobufEncoder protobuf Encoder
func ProtobufEncoder(cmd int32, pb proto.Message) network.PackInf {
	pack := network.GetPooledPack()
	pack.SetCmd(cmd)
	data, _ := proto.Marshal(pb)
	pack.SetData(data)
	return pack
}
