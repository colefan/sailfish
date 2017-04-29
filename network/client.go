package network

import (
	"errors"

	"time"

	"github.com/colefan/logg"
)

const (
	ModeEvent   = 0 //事件机制，用于投递消息
	ModeHandler = 1 //每个session就是独立一个goroutine
)

//Client client
type Client struct {
	networkType           string
	address               string
	mode                  int
	log                   *logg.BaseLogger
	protocol              Protocol
	sendChannelBufferSize int
	dispatcher            PackDispatcherInf
	autoReconnect         bool
	closeChannel          chan string
	bReconnect            bool
	*TCPClient
	logicLoop             func(PackInf)
	sessionOpenedCallback func(session *TCPSession)
	sessionClosedCallback func(session *TCPSession)
}

//NewClient new client
func NewClient(network string, address string, protocol Protocol, sendBufferSize int, mode int) *Client {
	c := &Client{
		networkType:           network,
		address:               address,
		protocol:              protocol,
		sendChannelBufferSize: sendBufferSize,
		mode: mode,
	}
	return c

}

//SetSessionOpenedCallback setter
func (c *Client) SetSessionOpenedCallback(f func(session *TCPSession)) {
	c.sessionOpenedCallback = f
}

//SetSessionClosedCallback setter
func (c *Client) SetSessionClosedCallback(f func(session *TCPSession)) {
	c.sessionClosedCallback = f
}

//SetLogicLoopCall setter
func (c *Client) SetLogicLoopCall(f func(PackInf)) {
	c.logicLoop = f
}

//SetAutoReconnect setter
func (c *Client) SetAutoReconnect(auto bool) {
	c.autoReconnect = auto
}

//SetDispatcher setter
func (c *Client) SetDispatcher(dispatcher PackDispatcherInf) {
	c.dispatcher = dispatcher
}

//SetSendChannelSize setter
func (c *Client) SetSendChannelSize(size int) {
	c.sendChannelBufferSize = size
}

//SetLogger setter
func (c *Client) SetLogger(log *logg.BaseLogger) {
	c.log = log
}

//SetNetworkType setter
func (c *Client) SetNetworkType(networkType string) {
	c.networkType = networkType
}

//SetAddress setter
func (c *Client) SetAddress(address string) {
	c.address = address
}

//SetMode setter
func (c *Client) SetMode(mode int) {
	c.mode = mode
}

//SetProtocol setter
func (c *Client) SetProtocol(p Protocol) {
	c.protocol = p
}

//WriteMsg write
func (c *Client) WriteMsg(pack PackInf) error {
	return c.TCPClient.session.WriteMsg(pack)
}

//Init init
func (c *Client) Init() error {
	c.closeChannel = make(chan string)
	if c.log == nil {
		return errors.New("log is nil")
	}
	if c.networkType != "tcp" {
		return errors.New("unknow network tyep for client : " + c.networkType)
	}
	if len(c.address) == 0 {
		return errors.New("unkown address for client : " + c.address)
	}
	if c.protocol == nil {
		return errors.New("protocol is nil ")
	}
	if c.mode == ModeEvent {
		if c.dispatcher == nil {
			return errors.New("dispatcher is nil")
		}
	}

	return nil
}

//Run run
func (c *Client) Run() error {
	var err error
	c.TCPClient, err = NewTCPClient(c.networkType, c.address, c.protocol, 100, c.mode, c.dispatcher)
	if err != nil {
		return err
	}
	c.TCPClient.session.SetCloseCallback(c.closeCallback)

	c.TCPClient.Start()
	if c.sessionOpenedCallback != nil {
		c.sessionOpenedCallback(c.TCPClient.session)
	}
	go c.loop()

	return nil
}

func (c *Client) closeCallback(session *TCPSession) {
	if c.sessionClosedCallback != nil {
		c.sessionClosedCallback(session)
	}
	if c.autoReconnect {
		go c.Reconnect()
	}

}

//ShutDown shut down
func (c *Client) ShutDown() {
	c.TCPClient.session.ForceClose()
}

//Reconnect reconnect
func (c *Client) Reconnect() {
	count := time.Duration(1)
	for {
		time.Sleep(time.Second * 5 * count)
		err := c.Run()
		if err == nil {
			c.bReconnect = true
			c.closeChannel <- "close"
			return
		}
		count++
		if count >= 10 {
			count = 1
		}

	}

}

func (c *Client) loop() {
	if c.mode == ModeEvent {
		c.loopEvent()
	} else {
		c.loopHandler()
	}
}

func (c *Client) loopEvent() {
	isChannel := c.GetPackDispatcher().IsChannelDispatcher()
	if isChannel {
		for {
			item := c.GetPackDispatcher().FetchData()
			if item != nil {
				c.logicLoop(item.Pack)
			}
			select {
			case <-c.closeChannel:
				return
			default:
			}

		}

	} else {
		for {
			emptyCount := 0
			packList := c.GetPackDispatcher().FetchAllData()
			if packList != nil {
				emptyCount = 0
				for _, v := range packList {
					if v != nil {
						c.logicLoop(v.Pack)
					}
				}
			} else {
				emptyCount++
			}

			select {
			case <-c.closeChannel:
				return
			default:
			}

			if emptyCount >= 100 {
				time.Sleep(10 * time.Microsecond)
			}

		}
	}

}

func (c *Client) loopHandler() {
	for {
		pack, err := c.TCPClient.session.ReadMsg()
		if err != nil {
			c.TCPClient.session.Close()
			return
		}
		c.logicLoop(pack)

	}

}

//GetSession return tcp session
func (c *Client) GetSession() *TCPSession {
	if c.TCPClient == nil {
		return nil
	}

	return c.TCPClient.session
}
