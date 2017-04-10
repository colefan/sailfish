package main

import (
	"github.com/colefan/sailfish/network"
	"github.com/colefan/sailfish/network/codec"
)

func main() {
	p := codec.NewFixHeadProtocol(4)
	server, err := network.NewTCPServer("tcp", "0.0.0.0:11111", p, 100)
	if err != nil {
		return
	}
	network.GetTCPServerQos().Stat()
	go serverLoop(server)
	server.Serve(nil)
	server.Stop()
}

func loop(session *network.TCPSession) {
	for {
		msg, err := session.ReadMsg()
		if err != nil {
			session.Close()
		}
		session.WriteMsg(msg)
	}

}

func serverLoop(server *network.TCPServer) {
	for {
		items := server.GetPackDispatcher().FetchDataList()
		for _, item := range items {
			item.Session.WriteMsg(item.Pack)
		}
		//time.Sleep(time.Millisecond * 1)
	}

}

//Request sample message
type Request struct {
	A int `json:"a"`
	B int `json:"b"`
}
