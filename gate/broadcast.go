package gate

import (
	"github.com/colefan/sailfish/log"
	"github.com/colefan/sailfish/network"
	"runtime/debug"
	"sync/atomic"
)

type BroadCast struct {
	unSendMsgListChan chan network.PackInf
	started           int32
}

var _instBroadCast *BroadCast

func GetBroadCast() *BroadCast {
	return _instBroadCast
}

func (b *BroadCast) Run() {
	if atomic.CompareAndSwapInt32(&b.started, 0, 1) {
		go b.execute()

	}

}

func (b *BroadCast) postMsg(pack network.PackInf) {
	b.unSendMsgListChan <- pack
}

func (b *BroadCast) execute() {
	defer func() {
		if e := recover(); e != nil {
			log.Errorf("recover = %v\n,%v", e, string(debug.Stack()))
		}
		log.Errorf("broadcast panic")
	}()
	for {
		select {
		case pack, ok := <-b.unSendMsgListChan:
			if ok {
				b.broadCastAll(pack)
			}
		}
	}

}

func (b *BroadCast) broadCastAll(pack network.PackInf) {
	groupSize := GetClientSessionMntInst().GetGroupSize()
	for i := 0; i < groupSize; i++ {
		sessionList := GetClientSessionMntInst().CopyUserSessions(i)
		if sessionList == nil {
			continue
		}
		for _, session := range sessionList {
			session.WriteMsg(pack)
		}
	}

}

func BroadCastMsgToClient(pack network.PackInf) {
	GetBroadCast().postMsg(pack)
}

func init() {
	_instBroadCast = &BroadCast{}
	_instBroadCast.unSendMsgListChan = make(chan network.PackInf, 1024)
}
