package main

import (
	"encoding/binary"
	"math/rand"

	"fmt"

	"encoding/json"

	"time"

	"github.com/colefan/sailfish/network"
	"github.com/colefan/sailfish/network/codec"
)

func main() {
	p := codec.NewFixHeadProtocol(4)
	server, err := network.NewServer("tcp", "0.0.0.0:11111", p, 100)
	if err != nil {
		return
	}
	go func() {
		err2 := server.Serve(network.HandlerFunc(serverLoop))
		if err2 != nil {
			return
		}
	}()

	client, err := network.NewClientConnect("tcp", "127.0.0.1:8000", p, 100)
	if err != nil {
		fmt.Println("client can't be created.")
		return
	}
	clientLoop(client, nil, nil)
	server.Stop()

}

func serverLoop(session *network.Session, ctx network.Context, starterror error) {
	for {
		msg, err := session.ReadMsg()
		if err != nil {
			session.Close()
			fmt.Println("session closed from server loop")
			return
		}

		if msg != nil {
			if b, ok := msg.([]byte); ok {
				head := make([]byte, 4, 4)
				binary.LittleEndian.PutUint32(head, uint32(len(b)))
				req := Request{}
				err := json.Unmarshal(b, &req)
				if err != nil {
					fmt.Printf("server loop read msg from socket error %v\n", err)
				} else {
					fmt.Printf("server loop read msg = %v", req)
				}

				head = append(head, b...)
				session.SendMsg(head)
				fmt.Println("server response head ")
			}
		}
	}

}

func clientLoop(session *network.Session, ctx network.Context, starterror error) {
	go func() {
		for {
			msg, err := session.ReadMsg()
			if err != nil {
				fmt.Println("client read msg error")
				return
			}
			fmt.Printf("client read msg %v\n", msg)
		}

	}()

	for {
		req := Request{}
		req.A = rand.Intn(100)
		req.B = rand.Intn(100)
		b, err := json.Marshal(req)
		if err != nil {
			fmt.Println("client Marshal error")
			return
		}
		bLen := len(b)
		head := make([]byte, 4, 4+bLen)
		binary.LittleEndian.PutUint32(head, uint32(bLen))
		head = append(head, b...)
		err = session.SendMsg(head)
		if err != nil {
			fmt.Printf("client send msg error %v\n", err)
		}
		time.Sleep(time.Second * 1.0)

	}

}

//Request sample message
type Request struct {
	A int `json:"a"`
	B int `json:"b"`
}
