package rpc

import (
	"net"
	netrpc "net/rpc"
	"net/rpc/jsonrpc"

	"github.com/colefan/sailfish/log"
)

type RPCServer struct {
	*netrpc.Server
	address string
}

func NewServer(address string) *RPCServer {
	s := new(RPCServer)
	s.Server = netrpc.NewServer()
	s.address = address
	return s
}

func (s *RPCServer) Serve() {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalf("rpc server listen fatal :%v", err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			log.Info("rpc server new client in coming ")
			s.ServeCodec(jsonrpc.NewServerCodec(conn))
		}(conn)
	}
}

func (s *RPCServer) RegisterService(service interface{}) {
	s.Register(service)
}

func (s *RPCServer) RegisterServiceName(name string, service interface{}) {
	s.RegisterName(name, service)
}
