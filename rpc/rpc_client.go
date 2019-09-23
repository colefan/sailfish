package rpc

import (
	netrpc "net/rpc"
	"net/rpc/jsonrpc"
)

type RPCClient struct {
	*netrpc.Client
	address string
}

func NewRPCClient(address string) (*RPCClient, error) {
	c := new(RPCClient)
	c.address = address
	var err error
	c.Client, err = jsonrpc.Dial("tcp", address)
	return c, err
}
