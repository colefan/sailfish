package gate

import (
	"errors"
	"os"

	"os/signal"

	"fmt"

	"github.com/colefan/config"
	"github.com/colefan/logg"
	"github.com/colefan/sailfish/network"
)

const (
	// DefaultMaxNum 默认最大用户数
	DefaultMaxNum int = 5000
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
	logger           *logg.BaseLogger
	conf             *config.IniConfig
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

//SetLogger setter
func (g *Gate) SetLogger(logger *logg.BaseLogger) {
	g.logger = logger
}

//Logger getter
func (g *Gate) Logger() *logg.BaseLogger {
	return g.logger
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
	if g.logger == nil {
		g.logger = logg.NewLogger(256)
		g.logger.SetAppender("console", "")
		g.logger.Async()
	}
	//日志设置

	g.logger.Info("log starting...")
	gLog = g.logger
	//读取配置
	if g.conf == nil {
		g.conf = config.NewIniConfig()
	}
	err := g.conf.Parse("./config/gate.ini")
	if err != nil {
		g.logger.Error("parse ./gate.ini error " + err.Error())
		return err
	}
	g.logger.Info("gate config parsing success.")

	g.clientHandler = new(ClientHandler)
	g.serverHandler = new(ProxyInnerHandler)

	g.bInit = true
	return nil
}

//Run gate
func (g *Gate) Run() {
	//启动服务
	if !g.bInit {
		g.logger.Error("gate not inited.")
		return
	}
	g.logger.Info("gate run ...")

	if g.clientHandler == nil || g.serverHandler == nil {
		g.logger.Error("gate client handler or server handler not registed")
		return
	}
	//g.logger.Info("%v", g.clientProtocol)
	if g.clientProtocol == nil || g.serverProtocol == nil {
		g.logger.Error("gate client protocol or server protocol is nill")
		return
	}
	var err error
	gLog.Info(g.conf.String("serverListner"))
	g.proxyServer = network.NewTCPServer(g.conf.String("serverListner"), network.MsgHandleModeHandler, g.serverProtocol)
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
		g.logger.Error("clientListener create error " + err.Error())
		return
	}
	g.clientServer.SetSendChannelSize(256)
	g.clientServer.SetSessionHandler(g.clientHandler)
	g.clientServer.SetCheckPerSecond(true)

}

//Daemon daemon
func (g *Gate) Daemon() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan)
	//closec := make(chan int)
	for {
		select {
		case s := <-signalChan:
			if s == os.Kill {
				g.ShutDown()
			}
		}

	}

	fmt.Println("exit")
	//	select {}
}

//ShutDown shutdown gate server
func (g *Gate) ShutDown() {
	fmt.Println("shutdown route server")
	if g.proxyServer != nil {
		g.proxyServer.Stop()
	}
	fmt.Println("shutdown client server ")

	if g.clientServer != nil {
		g.clientServer.Stop()
	}
	fmt.Println("finish")
}

// HandleMsgFunc handle msg func
type HandleMsgFunc func(pack network.PackInf)
