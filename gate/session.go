package gate

import "github.com/colefan/sailfish/network"

const (
	StatusInit      = 0
	StatusConnected = 1
	StatusLogined   = 2
	StatusSet       = 3
	StatusClosed    = 10
)

//ClientUserData client data
type ClientUserData struct {
	UID                uint64
	StrUID             string
	CurServerID        uint64
	OpenID             string
	ValidCode          string
	CurServerSession   *network.TCPSession
	LoginServerSession *network.TCPSession
}

const (
	ServerInit      = 0
	ServerConnected = 1
	ServerRegisted  = 2
	ServerOffline   = 3
	ServerClosed    = 4
)

//ServerUserData server user data
type ServerUserData struct {
	Type         int    //服务器类型
	SID          uint64 //服务器ID
	StrSID       string //服务器ID
	userCount    int    //用户数
	warningCount int    //警告用户数
	maxCount     int    //最大用户数
}

//NewServerUserData creator
func NewServerUserData() *ServerUserData {
	return &ServerUserData{maxCount: 3000, warningCount: 2500}
}

//AddUserCount ++
func (user *ServerUserData) AddUserCount() {
	user.userCount++
}

//GetUserCount getter
func (user *ServerUserData) GetUserCount() int {
	return user.userCount
}

//SetMaxCount setter
func (user *ServerUserData) SetMaxCount(count int) {
	user.maxCount = count
}

//GetMaxCount getter
func (user *ServerUserData) GetMaxCount() int {
	return user.maxCount
}

//WarningCaps setter
func (user *ServerUserData) WarningCaps(less int) {
	user.warningCount = user.maxCount - less
	if user.warningCount <= 100 {
		user.warningCount = user.maxCount
	}

}

//IsWarning getter
func (user *ServerUserData) IsWarning() bool {
	if user.userCount >= user.maxCount {
		return true
	}
	return false

}
