package gate

import (
	"errors"
	"strconv"
	"sync"
)

type proxyServerListMap map[int32]*ProxyServerNode

// ProxyServerStore 代理服务器信息存储
type ProxyServerStore struct {
	serverMaps     proxyServerListMap
	serverTypeMaps map[int32]proxyServerListMap
	sync.RWMutex
}

var proxyServerStoreInst *ProxyServerStore

func newProxyServerStore() *ProxyServerStore {
	s := new(ProxyServerStore)
	s.serverMaps = make(map[int32]*ProxyServerNode)
	s.serverTypeMaps = make(map[int32]proxyServerListMap)
	return s
}

// MatchProxyServer match proxy server node by server type
func (s *ProxyServerStore) MatchProxyServer(sType int32) *ProxyServerNode {
	s.RLock()
	defer s.RUnlock()
	if v, ok := s.serverTypeMaps[sType]; ok {
		var selID int32
		var selServer *ProxyServerNode
		var selBusy bool
		for key, val := range v {
			if selID == 0 {
				if val.GetStatus() == 0 {
					selID = key
					selServer = val
					selBusy = val.IsBusyLimit()
				}
			} else {
				// if val.CurrentUsers
				if selBusy {
					//找负载不重的
					if val.IsBusyLimit() == false {
						selBusy = false
						selID = key
						selServer = val
					}
				} else {
					if val.IsBusyLimit() == false {
						if val.GetCurrentUsers() < selServer.GetCurrentUsers() {
							selServer = val
							selID = key
						}
					}
				}
			}
		}
		return selServer
	}

	return nil
}

// RegisterProxyNode 节点内存
func (s *ProxyServerStore) RegisterProxyNode(info *ProxyServerNode) error {
	if info == nil {
		return nil
	}
	s.Lock()
	defer s.Unlock()
	if _, ok := s.serverMaps[info.ServerID]; ok {
		return errors.New("repeated server id :" + strconv.Itoa(int(info.ServerID)))
	}
	s.serverMaps[info.ServerID] = info
	if v, ok := s.serverTypeMaps[info.ServerType]; ok {
		v[info.ServerID] = info
	} else {
		tmpV := make(map[int32]*ProxyServerNode)
		tmpV[info.ServerID] = info
		s.serverTypeMaps[info.ServerType] = tmpV
	}

	return nil
}

// UnRegisterProxyNode unregister proxy node
func (s *ProxyServerStore) UnRegisterProxyNode(info *ProxyServerNode) error {
	if info == nil {
		return nil
	}
	s.Lock()
	defer s.Unlock()

	delete(s.serverMaps, info.ServerID)

	if v, ok := s.serverTypeMaps[info.ServerType]; ok {
		delete(v, info.ServerID)
	}

	return nil
}

func (s *ProxyServerStore) Find(sid int32, serverType int32) *ProxyServerNode {
	s.RLock()
	defer s.RUnlock()
	if v, ok := s.serverMaps[sid]; ok {
		return v
	}
	return nil
}

// GetProxyServerStoreInst get instance
func GetProxyServerStoreInst() *ProxyServerStore {
	return proxyServerStoreInst
}

func init() {
	proxyServerStoreInst = newProxyServerStore()
}
