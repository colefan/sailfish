package main

import (
	"math/rand"

	"fmt"

	"encoding/json"

	"time"

	"sync"

	"github.com/colefan/sailfish/network"
	"github.com/colefan/sailfish/network/codec"
)

var clientWaitList sync.WaitGroup

func main() {
	p := codec.NewFixHeadProtocol(4)

	for i := 0; i < 4000; i++ {
		client, err := network.NewTCPClient("tcp", "127.0.0.1:11111", p, 50)
		if err != nil {
			fmt.Println("client can't be created.")
			continue
		}
		clientWaitList.Add(1)
		//go client.Loop(nil)
		go clientLoop(client)
		time.Sleep(time.Millisecond * 5)
	}
	clientWaitList.Wait()

}

func loop(session *network.TCPSession) {
	go func() {
		for {
			_, err := session.ReadMsg()
			if err != nil {
				return
			}
		}
	}()

	for {
		req := Request{}
		req.A = rand.Intn(100)
		req.B = rand.Intn(100)
		b, err := json.Marshal(req)
		if err != nil {
			session.Close()
			return
		}
		pack := new(network.BasePack)
		pack.SetPackData(b)
		session.WriteMsg(pack)
		time.Sleep(time.Millisecond * 50)
	}
}

func clientLoop(client *network.TCPClient) {

	for {
		client.GetPackDispatcher().FetchDataList()
		req := Request{}
		req.A = rand.Intn(100)
		req.B = rand.Intn(100)
		b, err := json.Marshal(req)
		if err != nil {
			client.Close()
			return
		}
		pack := new(network.BasePack)
		pack.SetPackData(b)
		client.WriteMsg(pack)
		time.Sleep(time.Millisecond * 50)

	}

}

//Request sample message
type Request struct {
	A int `json:"a"`
	B int `json:"b"`
}
