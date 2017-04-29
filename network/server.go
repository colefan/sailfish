package network

import (
	"errors"

	"time"

	"github.com/colefan/logg"
)

//Server server class
type Server struct {
	log             *logg.BaseLogger
	networkType     string
	address         string
	protocol        Protocol
	sendChannelSize int
	mode            int
	logicLoopFunc   func(PackInf)
	dispatcher      PackDispatcherInf
	bQos            bool //是否开启qos
	*TCPServer
}

//NewServer new server
func NewServer(network string, address string, p Protocol, sendBufferSize int, mode int, dispatcher PackDispatcherInf) *Server {
	s := &Server{
		networkType:     network,
		address:         address,
		protocol:        p,
		sendChannelSize: sendBufferSize,
		mode:            mode,
		dispatcher:      dispatcher,
	}
	return s
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
func (s *Server) SetMode(mode int) {
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

//Init init server
func (s *Server) Init() error {
	if s.log == nil {
		return errors.New("log is nil")
	}

	if s.networkType != "tcp" {
		return errors.New("unkowned network type " + s.networkType)
	}

	if s.protocol == nil {
		return errors.New("protocol is nil")
	}

	if len(s.address) == 0 {
		return errors.New("address is invalid")
	}

	if s.mode == ModeEvent {
		if s.dispatcher == nil {
			return errors.New("dispatcher is nil")
		}
	}
	return nil
}

//Run run
func (s *Server) Run() error {
	var err error
	s.TCPServer, err = NewTCPServer(s.networkType, s.address, s.protocol, s.sendChannelSize, s.mode, s.dispatcher)
	if err != nil {
		return err
	}
	if s.bQos {
		GetTCPServerQos().Stat()
	}
	if s.mode == ModeHandler {
		go s.TCPServer.Serve(s.handler)
	} else {
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
		pack.SetTCPSession(session)
		if err != nil {
			session.Close()
			return
		}
		s.logicLoopFunc(pack)
	}

}

//ShutDown shutdown
func (s *Server) ShutDown() {
	if s.TCPServer != nil {
		s.TCPServer.Stop()
	}
}
