package network

import (
	"errors"
	"strings"

	"time"

	"github.com/colefan/logg"
)

type MsgHandlerModeType int

const (
	MsgHandleModeEvent MsgHandlerModeType = iota
	MsgHandleModeHandler
)

//Server server class
type Server struct {
	log             *logg.BaseLogger
	networkType     string
	address         string
	protocol        Protocol
	sendChannelSize int
	mode            MsgHandlerModeType
	logicLoopFunc   func(PackInf)
	dispatcher      PackDispatcherInf
	bQos            bool //是否开启qos
	crtFilePath     string
	keyFilePath     string
	agentHandler    SessionHandler
	*TCPServer
}

// SetCrtFilePath set file path
func (s *Server) SetCrtFilePath(path string) {
	s.crtFilePath = path
}

// SetKeyFilePath set file path
func (s *Server) SetKeyFilePath(path string) {
	s.keyFilePath = path
}

//SetQos setter
func (s *Server) SetQos(qos bool) {
	s.bQos = qos
}

//SetLogicLoopFunc setter
func (s *Server) SetLogicLoopFunc(f func(PackInf)) {
	s.logicLoopFunc = f
}

//SetDispatcher d
func (s *Server) SetDispatcher(d PackDispatcherInf) {
	s.dispatcher = d
}

//SetMode setter
func (s *Server) SetMode(mode MsgHandlerModeType) {
	s.mode = mode
}

//SetSendChannelSize size
func (s *Server) SetSendChannelSize(size int) {
	s.sendChannelSize = size
}

//SetProtocol setter
func (s *Server) SetProtocol(p Protocol) {
	s.protocol = p
}

//SetLogger setter
func (s *Server) SetLogger(log *logg.BaseLogger) {
	s.log = log
}

//SetNetworkType setter
func (s *Server) SetNetworkType(network string) {
	s.networkType = network
}

//SetAddress address
func (s *Server) SetAddress(address string) {
	s.address = address
}

// SetSessionHandler agent handler
func (s *Server) SetSessionHandler(agentHandler SessionHandler) {
	s.agentHandler = agentHandler
}

//Init init server
func (s *Server) init() error {
	if s.log == nil {
		return errors.New("log is nil")
	}

	if s.networkType != NetWorkTypeTCP && s.networkType != NetWorkTypeWS && s.networkType != NetWorkTypeWSS {
		return errors.New("unkowned network type " + s.networkType)
	}

	if s.protocol == nil {
		return errors.New("protocol is nil")
	}

	if len(s.address) == 0 {
		return errors.New("address is invalid:" + s.address)
	}
	if strings.Index(s.address, ":") < 0 {
		return errors.New("adress is invalid:" + s.address)
	}

	if s.mode == MsgHandleModeEvent {
		if s.dispatcher == nil {
			return errors.New("dispatcher is nil")
		}
		if s.logicLoopFunc == nil {
			return errors.New("logicLoopFunc is nil")
		}
	} else {
		if s.agentHandler == nil {
			return errors.New("sessionHandler is nil")
		}
	}

	if s.networkType == NetWorkTypeWSS {
		if s.crtFilePath == "" || s.keyFilePath == "" {
			return errors.New("wss : needs ssl files")
		}
	}

	return nil
}

//Run run
func (s *Server) Run() error {
	var err error
	err = s.init()
	if err != nil {
		return err
	}
	s.TCPServer, err = newServer(s.networkType, s.address, s.protocol, s.sendChannelSize, s.mode, s.dispatcher, s.agentHandler)
	if err != nil {
		return err
	}

	if s.bQos {
		GetTCPServerQos().Stat()
	}

	if s.mode == MsgHandleModeHandler {
		netDebug("modeHandler")
		go s.TCPServer.Serve(s.handler)
	} else {
		netDebug("modeEvent")
		go s.TCPServer.Serve(nil)
		go s.eventLoop()
	}

	return nil
}

func (s *Server) eventLoop() {
	isChannel := s.GetPackDispatcher().IsChannelDispatcher()
	if isChannel {
		for {
			item := s.GetPackDispatcher().FetchData()
			if item != nil {
				s.logicLoopFunc(item.Pack)
			}
		}

	} else {
		emptyCount := 0
		for {
			packList := s.GetPackDispatcher().FetchAllData()
			if packList != nil {
				emptyCount = 0
				for _, v := range packList {
					if v != nil {
						s.logicLoopFunc(v.Pack)
					}
				}
			} else {
				emptyCount++
			}

			if emptyCount >= 100 {
				time.Sleep(10 * time.Microsecond)
			}
		}

	}

}

func (s *Server) handler(session *TCPSession) {
	for {
		pack, err := session.ReadMsg()
		if err != nil {
			session.Close()
			return
		}
		pack.SetTCPSession(session)
		s.logicLoopFunc(pack)
	}

}

//ShutDown shutdown
func (s *Server) ShutDown() {
	if s.TCPServer != nil {
		s.TCPServer.Stop()
	}
}
