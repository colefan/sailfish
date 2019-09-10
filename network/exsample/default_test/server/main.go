package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	"time"

	"github.com/colefan/logg"
	"github.com/colefan/sailfish/console"
	"github.com/colefan/sailfish/network"
	"github.com/colefan/sailfish/network/codec"
)

var msgHandler *MyHandler

func main() {
	go func() {

	}()

	log := logg.NewLogger(100)
	log.SetAppender("console", `{}`)
	log.Async()

	protocol := codec.NewDefaultProtocol(true, 4)
	// handler := network.NewDefaultSessionHandler()

	s := network.NewTCPServer("0.0.0.0:8000", network.MsgHandleModeHandler, protocol)
	s.SetSendChannelSize(10240)
	s.SetLogger(log)
	s.SetSessionHandler(msgHandler)
	s.SetQos(true)

	err := s.Run()
	if err != nil {
		fmt.Printf("server run failed:%v\n", err)
		os.Exit(0)
	}

	http.ListenAndServe("0.0.0.0:6060", nil)

	console.Deamon()

}

type MyHandler struct {
	*network.DefaultSessionHandler
}

func (h *MyHandler) SessionOpen(session *network.TCPSession) {
	fmt.Printf("session open :%v\n", session.ID())
}

func (h *MyHandler) SessionClose(session *network.TCPSession) {
	fmt.Printf("session close :%v\n", session.ID())
}

// func (h *MyHandler) HandleMsg(pack network.PackInf) {
// 	fmt.Printf("recv msg = %v,data = %v\n", pack, string(pack.GetPackBody().GetData()))
// }

func logicLoop(pack network.PackInf) {
	// fmt.Printf("1 recv msg = %v,data = %v\n", pack, string(pack.GetPackBody().GetData()))
	s := pack.GetTCPSession()
	// msg := pack.GetPackBody()
	datamsg := PingPangMsg{}
	json.Unmarshal(pack.GetData(), &datamsg)
	datamsg.ReplyTime = time.Now().UnixNano() / 1e6
	data, _ := json.Marshal(&datamsg)
	pack.SetData(data)
	s.WriteMsg(pack)

}

type PingPangMsg struct {
	SendTime  int64
	ReplyTime int64
}

func init() {
	msgHandler = &MyHandler{}
	msgHandler.DefaultSessionHandler = network.NewDefaultSessionHandler()
	msgHandler.RegisterCmdHandler(0x01, logicLoop)
	msgHandler.RegisterCmdHandler(0x02, logicLoop)

}
