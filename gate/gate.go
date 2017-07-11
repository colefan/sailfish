package gate

import (
	"os"

	"os/signal"

	"fmt"

	"github.com/colefan/config"
	"github.com/colefan/logg"
	"github.com/colefan/sailfish/network"
)

const (
	MaxNum int = 3000
)

//Gate a gate for routing and load balance
type Gate struct {
	MaxClientConnNum int
	LittleEndian     bool  //大小字节
	route            Route //路由设置
	routeServer      *network.TCPServer
	clientServer     *network.TCPServer
	bInit            bool
	clientHandler    SessionHandler //面向客户端的handler
	serverHandler    SessionHandler //面向服务端的handler
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

//GetRouteServerSessionMgr get route server manager
func (g *Gate) GetRouteServerSessionMgr() *network.SessionMgr {
	if g.routeServer != nil {
		return g.routeServer.GetSessionMgr()
	}
	return nil
}

//RegisterClientHandler register client handler
func (g *Gate) RegisterClientHandler(handler SessionHandler) {
	g.clientHandler = handler
}

//RegisterServerHandler register server handler
func (g *Gate) RegisterServerHandler(handler SessionHandler) {
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

	//g.clientHandler = new(ClientHandler)
	//g.serverHandler = new(ServerHandler)
	//设置路由规则
	g.route = new(BaseRoute)
	//TODO
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
	g.routeServer, err = network.NewTCPServer("tcp", g.conf.String("serverListner"), g.serverProtocol, 10240, 1, nil)
	if err != nil {
		g.logger.Error("serverListner create error " + err.Error())
		return
	}

	g.clientServer, err = network.NewTCPServer("tcp", g.conf.String("clientListner"), g.clientProtocol, 256, 1, nil)

	if err != nil {
		g.logger.Error("clientListener create error " + err.Error())
		return
	}
	g.clientServer.SetCheckPerSecond(true)

	go g.routeServer.Serve(g.routeServerLoop)
	go g.clientServer.Serve(g.clientServerLoop)

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
	if g.routeServer != nil {
		g.routeServer.Stop()
	}
	fmt.Println("shutdown client server ")

	if g.clientServer != nil {
		g.clientServer.Stop()
	}
	fmt.Println("finish")
}

func (g *Gate) routeServerLoop(session *network.TCPSession) {
	for {
		if session.Status() == ServerInit {
			g.serverHandler.SessionOpen(session)
			session.SetStatus(ServerConnected)
		} else {
			msg, err := session.ReadMsg()
			if err != nil {
				session.Close()
				//g.serverHandler.SessionClose(session)
				return
			}
			g.serverHandler.HandleMessage(msg, session)
		}
	}

}

func (g *Gate) clientServerLoop(session *network.TCPSession) {
	for {
		if session.Status() == StatusInit {
			session.SetCheckPerSecond(true)
			session.SetTimeOut(180)
			g.clientHandler.SessionOpen(session)
			session.SetStatus(StatusConnected)
		} else {
			msg, err := session.ReadMsg()
			if err != nil {
				session.Close()
				//g.clientHandler.SessionClose(session)
				//fmt.Println("test for session close error = %s", err.Error())
				return
			}
			g.clientHandler.HandleMessage(msg, session)

		}
	}

}
