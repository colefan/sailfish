package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/colefan/logg"
	"github.com/colefan/sailfish/console"
	"github.com/colefan/sailfish/network"
	"github.com/colefan/sailfish/network/codec"
)

func main() {

	log := logg.NewLogger(1024)
	log.SetAppender("console", `{}`)
	log.Async()

	for i := 0; i < 2000; i++ {

		protocol := codec.NewDefaultProtocol(true, 4)
		handler := &ClientHandler{}

		client := network.NewTCPClient("192.168.20.127:8000", protocol)
		client.SetLogger(log)
		client.SetHandler(handler)
		client.SetAutoReconnect(true)
		client.SetSendChannelSize(1024)
		err := client.Run()
		if err != nil {
			fmt.Printf("client run failed :%v\n", err)
			os.Exit(0)
		}
	}
	console.Deamon()

}

type PingPangMsg struct {
	SendTime  int64
	ReplyTime int64
}

type ClientHandler struct {
	sendCount    int
	serverTimeMs int64
	dltTimeMs    int64
}

func (h *ClientHandler) SessionOpen(session *network.TCPSession) {
	fmt.Printf(" session open : %v\n", session.ID())
	pack := network.GetPooledPack()
	msg := pack.(*network.Message)
	msg.Head.Magic = 0x08
	msg.Head.Cmd = 0x01
	dataMsg := PingPangMsg{}
	dataMsg.SendTime = time.Now().UnixNano() / 1e6
	data, _ := json.Marshal(&dataMsg)
	msg.SetData(data)

	session.WriteMsg(pack)

}

func (h *ClientHandler) SessionClose(session *network.TCPSession) {
	fmt.Printf(" session close : %v\n", session.ID())

}

func (h *ClientHandler) HandleMsg(pack network.PackInf) {
	if pack.GetCmd() == 0x01 {

		tmpDataMsg := PingPangMsg{}
		json.Unmarshal(pack.GetData(), &tmpDataMsg)
		h.serverTimeMs = tmpDataMsg.ReplyTime
		h.dltTimeMs = h.serverTimeMs - time.Now().UnixNano()/1e6
		tmpDataMsg.SendTime = time.Now().UnixNano()/1e6 + h.dltTimeMs
		tmpDataMsg.ReplyTime = 0

		tmpData, _ := json.Marshal(&tmpDataMsg)
		pack.SetData(tmpData)

		pack.GetTCPSession().WriteMsg(pack)
		return

	}

	// fmt.Printf(" recv msg : %v\n", pack)
	h.sendCount++
	if h.sendCount%10 == 0 {
		time.Sleep(time.Millisecond * 100)
	}
	s := pack.GetTCPSession()
	// s.WriteMsg(pack)
	datamsg := PingPangMsg{}
	json.Unmarshal(pack.GetData(), &datamsg)
	if datamsg.ReplyTime-datamsg.SendTime <= 100 && datamsg.ReplyTime >= datamsg.SendTime {
		//ok
	} else {
		fmt.Printf("reply msg is delayed: %v ms,reply = %v,send = %v\n", datamsg.ReplyTime-datamsg.SendTime, datamsg.ReplyTime, datamsg.SendTime)
	}

	datamsg.SendTime = time.Now().UnixNano()/1e6 + (h.dltTimeMs)
	datamsg.ReplyTime = 0
	msg := pack.(*network.Message)
	msg.Head.Magic = 0x08
	msg.Head.Cmd = 0x02
	data, _ := json.Marshal(datamsg)
	msg.SetData(data)
	// pack.SetPackBody(msg)
	s.WriteMsg(pack)

}
