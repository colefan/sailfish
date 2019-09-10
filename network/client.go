package network

import (
	"errors"
	"strings"
	"sync"
	"sync/atomic"

	"time"

	"github.com/colefan/logg"
)

// const (
// 	ModeEvent   = 0 //事件机制，用于投递消息
// 	ModeHandler = 1 //每个session就是独立一个goroutine
// )

//Client client
type Client struct {
	networkType           string
	address               string
	mode                  MsgHandlerModeType
	log                   *logg.BaseLogger
	protocol              Protocol
	sendChannelBufferSize int
	dispatcher            PackDispatcherInf
	autoReconnect         bool
	//closeChannel          chan string
	clientClosed int32
	bReconnect   bool
	*TCPClient
	logicLoop             func(PackInf)
	sessionOpenedCallback func(session *TCPSession)
	sessionClosedCallback func(session *TCPSession)

	disposeOnce   sync.Once
	reconnectWait sync.WaitGroup
	agentHandler  SessionHandler
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
func (c *Client) SetMode(mode MsgHandlerModeType) {
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
func (c *Client) init() error {
	c.clientClosed = 0
	//c.closeChannel = make(chan string)
	if c.log == nil {
		return errors.New("log is nil")
	}
	if c.networkType != NetWorkTypeTCP && c.networkType != NetWorkTypeWS && c.networkType != NetWorkTypeWSS {
		return errors.New("unknow network tyep for client : " + c.networkType)
	}
	if len(c.address) == 0 {
		return errors.New("unkown address for client : " + c.address)
	}
	if strings.Index(c.address, ":") < 0 {
		return errors.New("unkown address for client ：" + c.address)
	}
	if c.protocol == nil {
		return errors.New("protocol is nil ")
	}
	if c.mode == MsgHandleModeEvent {
		if c.dispatcher == nil {
			return errors.New("dispatcher is nil")
		}
	} else {
		if c.agentHandler == nil {
			return errors.New("sessionHandler is nil ")
		}
	}

	return nil
}

//Run run
func (c *Client) Run() error {
	var err error
	err = c.init()
	if err != nil {
		return err
	}
	if c.networkType == NetWorkTypeWS || c.networkType == NetWorkTypeWSS {
		c.TCPClient, err = createWebSocketClient(c.networkType, c.address, c.protocol, c.sendChannelBufferSize, c.mode, c.dispatcher, c.agentHandler)
	} else {
		c.TCPClient, err = createTCPClient(c.networkType, c.address, c.protocol, c.sendChannelBufferSize, c.mode, c.dispatcher, c.agentHandler)
	}
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
	//atomic.CompareAndSwapInt32(&c.clientClosed, 0, 1)

	if c.autoReconnect {
		if c.mode == MsgHandleModeEvent {
			netInfo("reconnect swap")
			atomic.CompareAndSwapInt32(&c.clientClosed, 0, 1)
			netInfo("reconnect add")
			c.reconnectWait.Add(1)
			closePack := GetPooledPack() //new(BasePack)
			closePack.SetCmd(9999)
			closePack.SetMagic(MagicNumberClientInner)
			c.dispatcher.PostData(closePack, nil)
			netInfo("reconnect wait")
			c.reconnectWait.Wait()
			go c.Reconnect()
		} else {
			go c.Reconnect()
		}
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
		netInfo("reconnected count = %d ", count)
		time.Sleep(time.Second * 5 * count)
		err := c.Run()
		if err == nil {
			c.bReconnect = true
			//atomic.SwapInt32(&c.clientClosed, 0)
			return
		}
		count++
		if count >= 10 {
			count = 1
		}

	}

}

func (c *Client) loop() {
	if c.mode == MsgHandleModeEvent {
		c.loopEvent()
	} else {
		// c.agentHandler.HandleMsg(msg PackInf)
		// c.loopHandler()
	}
}

func (c *Client) loopEvent() {
	isChannel := c.GetPackDispatcher().IsChannelDispatcher()
	if isChannel {
		for {

			item := c.GetPackDispatcher().FetchData()
			if item.Pack.GetCmd() == 99999 && item.Pack.GetMagic() == MagicNumberClientInner {
				if atomic.CompareAndSwapInt32(&c.clientClosed, 1, 0) {
					c.reconnectWait.Done()
					return
				}

			}
			if item != nil {
				c.logicLoop(item.Pack)
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

			if emptyCount >= 100 {
				time.Sleep(10 * time.Microsecond)
			}

			if atomic.CompareAndSwapInt32(&c.clientClosed, 1, 0) {
				c.reconnectWait.Done()
				return
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

// SetHandler setter
func (c *Client) SetHandler(agent SessionHandler) {
	c.agentHandler = agent
}

// GetHandler getter handler
func (c *Client) GetHandler() SessionHandler {
	return c.agentHandler
}
