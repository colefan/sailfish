package sailfish

import (
	"github.com/colefan/sailfish/rpc"
)

func CreateRPCServer(address string) *rpc.RPCServer {
	return rpc.NewServer(address)

}

func CreateRPCClient(address string) (*rpc.RPCClient, error) {
	return rpc.NewRPCClient(address)
}
