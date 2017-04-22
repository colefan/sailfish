package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/colefan/logg"
	"github.com/colefan/sailfish/console"
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
	log := logg.NewLogger(100)
	log.EnableFuncCallDepath(true)
	log.SetAppender("console", "")
	p := codec.NewJSONProtocol()
	p.Register(AddReq{})
	p.Register(AddResp{})
	server := network.NewServer("tcp", "0.0.0.0:7777", p, 100, 0, network.NewPackDispatcher())
	server.SetLogger(log)
	server.SetQos(true)
	server.SetLogicLoopFunc(serverHandler)
	err := server.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = server.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client := network.NewClient("tcp", "0.0.0.0:7777", p, 100, 0)
	client.SetLogger(log)
	client.SetDispatcher(network.NewPackDispatcher())
	client.SetLogicLoopCall(clientLogic)
	err = client.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = client.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	go clientReq(client)
	fmt.Println("okle")
	console.Deamon()

}

func serverHandler(pack network.PackInf) {
	if pack != nil {
		req := pack.GetPackBody()
		var resp AddResp
		resp.Val = req.(*AddReq).Param1 + req.(*AddReq).Param2
		respPack := &network.BasePack{}
		respPack.SetPackBody(&resp)
		fmt.Println(pack.GetTCPSession())
		err := pack.GetTCPSession().WriteMsg(respPack)
		if err != nil {
			fmt.Println(err)
		}

	}

}

func clientReq(client *network.Client) {
	for {
		var req AddReq
		req.Param1 = rand.Intn(100)
		req.Param2 = rand.Intn(100)
		fmt.Printf("req send %v\n", req)
		pack := &network.BasePack{}
		pack.SetPackBody(&req)

		err := client.WriteMsg(pack)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		req.Param1++
		req.Param2++
		fmt.Printf("req send %v\n", req)

		pack = &network.BasePack{}
		pack.SetPackBody(&req)
		err = client.WriteMsg(pack)
		if err != nil {
			return
		}

		time.Sleep(time.Microsecond * 500)
	}

}

var clientCount int

func clientLogic(pack network.PackInf) {
	if pack != nil {
		resp := pack.GetPackBody()
		clientCount++
		fmt.Printf("clientLoop resp %v\n", resp)
		fmt.Printf("clientLoop loop = %d \n ", clientCount)
		fmt.Printf("clientLoop read msg AddResp.Val = %d \n", (resp.(*AddResp).Val))

	}
}
