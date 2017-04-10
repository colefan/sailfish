package gate

const (
	statusInit   = 0
	statusSet    = 1
	statusClosed = 10
)

//GateClientData client data
type GateClientData struct {
	nUID   uint64
	strUID string
	status int
}

func (user *GateClientData) SetUID(id uint64) {
	user.nUID = id
}

func (user *GateClientData) SetStringUID(id string) {
	user.strUID = id
}

func (user *GateClientData) UID() (uint64, string) {
	return user.nUID, user.strUID
}

func (user *GateClientData) SetStatus(status int) {
	user.status = status
}

func (user *GateClientData) Status() int {
	return user.status
}

const (
	serverInit     = 0
	serverRegisted = 1
	serverOffline  = 2
	serverClosed   = 3
)

type serverData struct {
	serverType      int
	serverStatus    int
	serverID        uint64
	strServerID     string
	serverUserCount int
	serverMaxCount  int
}

func newServerData() *serverData {
	return &serverData{serverMaxCount: 3000}
}

func (user *serverData) SetServerType(sType int) {
	user.serverType = sType
}

func (user *serverData) GetServerType() int {
	return user.serverType
}

func (user *serverData) SetServerStatus(status int) {
	user.serverStatus = status
}

func (user *serverData) GetServerStatus() int {
	return user.serverStatus
}

func (user *serverData) SetID(id uint64) {
	user.serverID = id
}

func (user *serverData) SetStringID(strID string) {
	user.strServerID = strID
}

func (user *serverData) GetID() (uint64, string) {
	return user.serverID, user.strServerID
}

func (user *serverData) AddUserCount() {
	user.serverUserCount++
}

func (user *serverData) GetUserCount() int {
	return user.serverUserCount
}

func (user *serverData) SetMaxCount(count int) {
	user.serverMaxCount = count
}

func (user *serverData) GetMaxCount() int {
	return user.serverMaxCount
}
