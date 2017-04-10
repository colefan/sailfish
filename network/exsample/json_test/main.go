package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/colefan/sailfish/network"
	"github.com/colefan/sailfish/network/codec"
)

type AddReq struct {
	Param1 int
	Param2 int
}

type AddResp struct {
	Val int
}

func main() {
	p := codec.NewJSONProtocol()
	p.Register(AddReq{})
	p.Register(AddResp{})
	server, err := network.NewServer("tcp", "0.0.0.0:7777", p, 100)
	if err != nil {
		return
	}
	go func() {
		err1 := server.Serve(network.HandlerFunc(serverLoop))
		fmt.Printf("server loop finiished %v\n", err1)
	}()

	//server.Listener(
	client, err2 := network.NewClientConnect("tcp", "0.0.0.0:7777", p, 100)
	if err2 != nil {
		fmt.Printf("client loop finished %v\n", err2)
		server.Stop()
	}
	clientLoop(client)

}

func serverLoop(session *network.Session, ctx network.Context, err error) {
	if session == nil {
		fmt.Printf("serverLoop session is nil\n")
		return
	}
	for {
		req, err2 := session.ReadMsg()
		if req == nil {
			fmt.Printf("serverLoop ReadMsg is nil\n")
			continue
		}
		if err2 != nil {
			fmt.Printf(err2.Error())
			session.Close()
			return
		}
		fmt.Printf("serverLoop readMsg is %v \n", req)
		var resp AddResp
		resp.Val = req.(*AddReq).Param1 + req.(*AddReq).Param2

		err2 = session.SendMsg(&resp)
		if err2 != nil {
			fmt.Printf("serverLoop send resp error %v\n", err2.Error())
			session.Close()
			return
		}
	}

}

func clientLoop(session *network.Session) {
	fmt.Println("Begin to start to send msg to Server")
	var clientCount int
	clientCount = 0
	go func() {
		for {
			var req AddReq
			req.Param1 = rand.Intn(100)
			req.Param2 = rand.Intn(100)
			fmt.Printf("req send %v\n", req)

			err := session.SendMsg(&req)

			if err != nil {
				fmt.Println(err.Error())
				session.Close()
				return
			}
			req.Param1++
			req.Param2++
			fmt.Printf("req send %v\n", req)

			err = session.SendMsg(&req)
			if err != nil {
				session.Close()
				return
			}

			time.Sleep(time.Microsecond * 500)
		}

	}()

	for {
		resp, err := session.ReadMsg()
		if err != nil {
			fmt.Println("client loop ReadMsg error " + err.Error())
			session.Close()
			return
		}
		clientCount++
		fmt.Printf("clientLoop resp %v\n", resp)
		fmt.Printf("clientLoop loop = %d \n ", clientCount)
		fmt.Printf("clientLoop read msg AddResp.Val = %d \n", (resp.(*AddResp).Val))

	}

}
