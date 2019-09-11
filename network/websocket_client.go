package network

import (
	"net/url"
	"strings"

	"github.com/colefan/sailfish/log"
	"github.com/gorilla/websocket"
)

func newWebSocketClient(network, address string, p Protocol, sendChanSize int, mode MsgHandlerModeType, dispatcher PackDispatcherInf, agent SessionHandler) (*TCPClient, error) {
	client := &TCPClient{
		dispatcher:   dispatcher,
		mode:         mode,
		isWebSocket:  true,
		agentHandler: agent,
	}
	var u url.URL
	path := "/ws"
	if idx := strings.LastIndex(address, ":"); idx >= 0 {
		tmp := address[idx:]
		if idx2 := strings.Index(tmp, "/"); idx2 >= 0 {
			tmp2 := tmp[idx2:]
			path = tmp2
			address = address[0 : len(address)-len(tmp2)]
		}
	}

	if network == "ws" {
		u = url.URL{Scheme: "ws", Host: address, Path: path}
	} else {
		u = url.URL{Scheme: "wss", Host: address, Path: path}
	}

	log.Infof("connecting to %s", u.String())
	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Errorf("websocket client dail error :%v", err)
		return nil, err
	}
	codec, err := p.NewWsSocketCodec(ws)
	if err != nil {
		log.Errorf("websocket client new codec error :%v", err)
		ws.Close()
		return nil, err
	}

	session := newWsSession(ws, codec, sendChanSize, client.dispatcher, client.agentHandler)
	client.session = session
	return client, nil
}
