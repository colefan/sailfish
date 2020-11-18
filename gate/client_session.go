package gate

import "time"

const (
	ClientStatusInit = iota
	ClientStatusHandShaked
)

// ProxyNodeTypeMapper node type mapper
type ProxyNodeTypeMapper map[int32]*ProxyServerNode

//ClientUserData client data
type ClientUserData struct {
	ChannleID                string
	Version                  string
	UID                      uint64
	Status                   int
	ClientType               string
	UserIP                   string
	Account                  string
	AccountType              int
	Token                    string
	ClientAuth               string
	LastBeatTime             int64
	KickOffWhenRepeatedLogin bool
	AccountID                uint64

	ProxyNodeList ProxyNodeTypeMapper
}

// NewClientUserData new client user data
func NewClientUserData() *ClientUserData {
	client := new(ClientUserData)
	client.LastBeatTime = time.Now().Unix()
	client.ProxyNodeList = make(map[int32]*ProxyServerNode)
	return client
}
