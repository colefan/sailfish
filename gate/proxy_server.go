package gate

import (
	"sync/atomic"

	"github.com/colefan/sailfish/network"
)

//SERVER_TYPE
const (
	ServerTypeUnknow = 0
	ServerTypeLogin  = 1
	ServerTypeHall   = 2
	ServerTypeGame   = 10
	ServerTypeCenter = 20
)

const (
	UIDTypeUser    = 0
	UIDTypeSession = 101
)

const (
	ProxyNodeStatusInit = iota
	ProxyNodeStatusNormal
	ProxyNodeStatusOffline
)

// ProxyServerNode 内部服务节点
type ProxyServerNode struct {
	ServerType        int32
	ServerID          int32
	status            int32
	IP                string
	Port              int
	Load              int //负载信息
	currentUsers      int
	MaxUsers          int
	Session           *network.TCPSession
	LastHeartBeatTime int32
}

// NewInnerServerNode 内部服务节点
func NewProxyServerNode() *ProxyServerNode {
	s := new(ProxyServerNode)
	return s
}

// GetCurrentUsers
func (proxy *ProxyServerNode) GetCurrentUsers() int {
	return proxy.currentUsers
}

func (proxy *ProxyServerNode) IsBusyLimit() bool {
	return false
}

// GetStatus status
func (proxy *ProxyServerNode) GetStatus() int32 {
	return atomic.LoadInt32(&proxy.status)
}

func (proxy *ProxyServerNode) SetStatus(state int32) (old int32) {
	return atomic.SwapInt32(&proxy.status, state)
}

func (proxy *ProxyServerNode) Swap(old, new int32) bool {
	return atomic.CompareAndSwapInt32(&proxy.status, old, new)
}
