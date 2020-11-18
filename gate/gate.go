package gate

import (
	"errors"
	"os"

	"os/signal"

	"fmt"

	"github.com/colefan/config"
	"github.com/colefan/sailfish/log"
	"github.com/colefan/sailfish/network"
)

const (
	// DefaultMaxNum 默认最大用户数
	DefaultMaxNum int = 5000
)

const (
	TagProd = "prod"
	TagDev  = "dev"
	TagTest = "test"
)

//Gate a gate for routing and load balance
type Gate struct {
	MaxClientConnNum int
	LittleEndian     bool //大小字节
	proxyServer      *network.Server
	certFile         string
	keyFile          string
	clientServer     *network.Server
	bInit            bool
	clientHandler    network.SessionHandler //面向客户端的handler
	serverHandler    network.SessionHandler //面向服务端的handler
	clientProtocol   network.Protocol
	serverProtocol   network.Protocol
	conf             *config.IniConfig
	Tag              string
}

//GetClientServerSessionMgr get session manager
func (g *Gate) GetClientServerSessionMgr() *network.SessionMgr {
	if g.clientServer != nil {
		return g.clientServer.GetSessionMgr()
	}
	return nil
}

//GetProxyServerSessionMgr get proxy server manager
func (g *Gate) GetProxyServerSessionMgr() *network.SessionMgr {
	if g.proxyServer != nil {
		return g.proxyServer.GetSessionMgr()
	}
	return nil
}

//RegisterClientHandler register client handler
func (g *Gate) RegisterClientHandler(handler network.SessionHandler) {
	g.clientHandler = handler
}

//RegisterProxyServerHandler register server handler
func (g *Gate) RegisterProxyServerHandler(handler network.SessionHandler) {
	g.serverHandler = handler
}

//SetClientProtocol setter
func (g *Gate) SetClientProtocol(p network.Protocol) {
	g.clientProtocol = p
}

//SetServerProtocol setter
func (g *Gate) SetServerProtocol(p network.Protocol) {
	g.serverProtocol = p
}

//SetConfig setter
func (g *Gate) SetConfig(cnf *config.IniConfig) {
	g.conf = cnf
}

//Config getter
func (g *Gate) Config() *config.IniConfig {
	return g.conf
}

//Init init gate
func (g *Gate) Init() error {

	log.Info("log starting...")
	//读取配置
	if g.conf == nil {
		g.conf = config.NewIniConfig()
	}
	err := g.conf.Parse("./config/gate.ini")
	if err != nil {
		log.Error("parse ./gate.ini error " + err.Error())
		return err
	}
	g.Tag = g.conf.String("tag")
	log.Info("gate config parsing success.")

	g.clientHandler = new(ClientHandler)
	g.serverHandler = new(ProxyInnerHandler)

	g.bInit = true
	return nil
}

//Run gate
func (g *Gate) Run() error {
	//启动服务
	if !g.bInit {
		return errors.New("gate has not inited")
	}
	log.Info("gate run ...")

	if g.clientHandler == nil || g.serverHandler == nil {
		return errors.New("gate client handler or server handler not registed")
	}
	//g.logger.Info("%v", g.clientProtocol)
	if g.clientProtocol == nil || g.serverProtocol == nil {
		//g.logger.Error("gate client protocol or server protocol is nill")
		return errors.New("gate client protocol or server protocol is nill")

	}
	var err error
	log.Info("proxy listener:" + g.conf.String("serverListener"))
	g.proxyServer = network.NewTCPServer(g.conf.String("serverListener"), network.MsgHandleModeHandler, g.serverProtocol)
	g.proxyServer.SetSendChannelSize(10240)
	g.proxyServer.SetSessionHandler(g.serverHandler)

	clientServerNetworkType := g.conf.String("clientNetwork")
	switch clientServerNetworkType {
	case network.NetWorkTypeTCP:
		g.clientServer = network.NewTCPServer(g.conf.String("clientListener"), network.MsgHandleModeHandler, g.clientProtocol)
	case network.NetWorkTypeWS:
		g.clientServer = network.NewWSServer(g.conf.String("clientListener"), network.MsgHandleModeHandler, g.clientProtocol)
	case network.NetWorkTypeWSS:
		g.clientServer = network.NewWSServer(g.conf.String("clientListener"), network.MsgHandleModeHandler, g.clientProtocol)
		g.keyFile = g.conf.String("keyFile")
		g.certFile = g.conf.String("certFile")
		g.clientServer.SetKeyFilePath(g.keyFile)
		g.clientServer.SetCrtFilePath(g.certFile)
	default:
		err = errors.New("unknow client network type:" + clientServerNetworkType)
	}
	if err != nil {
		log.Error("clientListener create error " + err.Error())
		return err
	}
	log.Info("client listener:" + g.conf.String("clientListener"))

	g.clientServer.SetSessionHandler(g.clientHandler)
	g.clientServer.SetSendChannelSize(256)
	g.clientServer.SetCheckPerSecond(true)

	if err := g.clientServer.Run(); err != nil {
		return err
	}

	if err := g.proxyServer.Run(); err != nil {
		return err
	}

	GetBroadCast().Run()

	return nil

}

//Daemon daemon
func (g *Gate) Daemon() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan)
	//closec := make(chan int)
	for {
		select {
		case s := <-signalChan:
			fmt.Printf("receive sigal:%v", s)
			if s == os.Kill || s == os.Interrupt {
				g.ShutDown()
				return
			}
		}

	}

	log.Info("system quit")
	//	select {}
}

//ShutDown shutdown gate server
func (g *Gate) ShutDown() {
	log.Info("shutdown route server")
	if g.proxyServer != nil {
		g.proxyServer.Stop()
	}
	log.Info("shutdown client server ")

	if g.clientServer != nil {
		g.clientServer.Stop()
	}
	log.Info("finish")
}

// HandleMsgFunc handle msg func
type HandleMsgFunc func(pack network.PackInf)
