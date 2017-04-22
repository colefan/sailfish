package main

import (
	"fmt"

	"time"

	"sync"

	"github.com/colefan/sailfish/network"
	"github.com/gogo/protobuf/proto"
	"qiyayule.com/gameproto/qpmsg"
	"qiyayule.com/qipai/qpcodec"
)

var clientWaitList sync.WaitGroup

func main() {
	p := qpcodec.NewDateClientRespProtocol(true, 4)

	for i := 0; i < 1; i++ {
		client, err := network.NewTCPClient("tcp", "127.0.0.1:9999", p, 50)
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
		bodyMsg := &qpmsg.ClientOpenRequest{}
		bodyMsg.ChannelId = "jiuyou.com"
		bodyMsg.ClientVersion = "111"

		rpcMsg := &qpmsg.RpcMessage{}
		rpcMsg.MsgCmdId = uint32(qpmsg.MSG_GATE_MSG_GATE_CLIENT_OPEN_REQ)

		rpcMsg.MsgBody, _ = proto.Marshal(bodyMsg)

		sendMsg := &qpmsg.ClientRequest{}
		sendMsg.CmdId = rpcMsg.MsgCmdId
		sendMsg.Msg = rpcMsg
		sendMsg.To = 1

		pack := &network.BasePack{}

		b, _ := proto.Marshal(sendMsg)
		pack.SetPackData(b)
		session.WriteMsg(pack)
		time.Sleep(time.Millisecond * 1000)
	}
}

func clientLoop(client *network.TCPClient) {
	go func() {
		for {
			msg, err := client.ReadMsg()

			if err != nil {
				return
			}
			if req, ok := msg.GetPackBody().(*qpmsg.ClientResponse); ok {
				rpcMsg := req.GetMsg()
				if rpcMsg.GetMsgCmdId() != req.CmdId {
					fmt.Println("message valid error cmd_id not equal")
					return
				}
				realMsg := &qpmsg.ClientOpenResponse{}
				err := proto.Unmarshal(rpcMsg.GetMsgBody(), realMsg)
				if err != nil {
					fmt.Printf("parse msg error cmd_id = %d \n", req.GetCmdId())
					return
				}
				fmt.Printf("%v\n", realMsg) //for test
			}

			fmt.Printf("read msg %v", msg)
		}
	}()

	for {
		client.GetPackDispatcher().FetchDataList()
		bodyMsg := &qpmsg.ClientOpenRequest{}
		bodyMsg.ChannelId = "jiuyou.com"
		bodyMsg.ClientVersion = "111"

		rpcMsg := &qpmsg.RpcMessage{}
		rpcMsg.MsgCmdId = uint32(qpmsg.MSG_GATE_MSG_GATE_CLIENT_OPEN_REQ)

		rpcMsg.MsgBody, _ = proto.Marshal(bodyMsg)

		sendMsg := &qpmsg.ClientRequest{}
		sendMsg.CmdId = rpcMsg.MsgCmdId
		sendMsg.Msg = rpcMsg
		sendMsg.To = 1

		pack := &network.BasePack{}

		b, _ := proto.Marshal(sendMsg)
		pack.SetPackData(b)
		client.WriteMsg(pack)
		time.Sleep(time.Millisecond * 1000)

	}

}

//Request sample message
type Request struct {
	A int `json:"a"`
	B int `json:"b"`
}
