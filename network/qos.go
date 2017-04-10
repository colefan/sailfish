package network

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var byteUnitList = []string{"", "K", "M", "G", "T", "PB", "EB", "ZB", "YB", "BB", "NB", "DB"}

//QosInf interface of qos
type QosInf interface {
	IsEnable() bool
	Enabled(enable bool)
	ShowQos()
	Stat()
	Interval() int
	AddAcceptedSessions()
	AddClosedSessions()
	AddReadPacket(packBytes int)
	AddWritePacket(packBytes int)
}

//TCPServerQos server's qos
type TCPServerQos struct {
	isEnable          int32
	timeSleep         int //stat interval time second
	startStatTime     time.Time
	acceptedSessions  int32
	closedSessions    int32
	readPacketCounts  int32
	writePacketCounts int32
	readBytes         int64
	writeBytes        int64
	statOnce          sync.Once
}

//IsEnable return ture or false
func (q *TCPServerQos) IsEnable() bool {
	val := atomic.LoadInt32(&q.isEnable)
	if val == 1 {
		return true
	}
	return false
}

//Enabled setter enable
func (q *TCPServerQos) Enabled(enable bool) {
	if enable {
		atomic.CompareAndSwapInt32(&q.isEnable, 0, 1)
	}
}

func formatCounts(count int32) string {
	if count < 1000 {
		return strconv.FormatInt(int64(count), 10)

	} else if count < 1000000 {
		return strconv.FormatFloat(float64(count)/1000.0, 'f', 2, 64) + "k"
	} else {
		return strconv.FormatFloat(float64(count)/1000000.0, 'f', 2, 64) + "m"
	}
}

func formatBytes(bytes int64) string {
	if bytes < 1024 {
		return strconv.FormatInt(bytes, 10)
	}
	tmpBytes := float64(bytes) / 1024.0
	unitIndex := 1
	for tmpBytes > 1024.0 {
		tmpBytes = tmpBytes / 1024.0
		unitIndex++
	}

	if unitIndex > len(byteUnitList) {
		return strconv.FormatFloat(tmpBytes, 'f', 2, 64) + "MAXLIMIT"

	}

	return strconv.FormatFloat(tmpBytes, 'f', 2, 64) + byteUnitList[unitIndex]

}

//ShowQos print the qos infomation
func (q *TCPServerQos) ShowQos() {
	if q.IsEnable() {
		t := time.Now().Sub(q.startStatTime)
		fmt.Println("QOS:t[", t.String(),
			"]\tconns[all:",
			formatCounts(atomic.LoadInt32(&q.acceptedSessions)),
			"\talive:",
			formatCounts(atomic.LoadInt32(&q.acceptedSessions)-atomic.LoadInt32(&q.closedSessions)),
			"\tclosed:",
			formatCounts(atomic.LoadInt32(&q.closedSessions)),
			"\t]\treadio[p:",
			formatCounts(atomic.LoadInt32(&q.readPacketCounts)),
			"\tb:",
			formatBytes(atomic.LoadInt64(&q.readBytes)),
			"\t]\twriteio[p:",
			formatCounts(atomic.LoadInt32(&q.writePacketCounts)),
			"\tb:",
			formatBytes(atomic.LoadInt64(&q.writeBytes)), "\t]")
	}
}

//Stat begin to stat the qos
func (q *TCPServerQos) Stat() {
	q.statOnce.Do(func() {
		q.startStatTime = time.Now()
		go func() {
			timer1 := time.NewTicker(time.Duration(q.Interval()) * time.Second)
			for {
				select {
				case <-timer1.C:
					go q.ShowQos()
				}
			}
		}()
	})
}

//Interval return the interval time seconds
func (q *TCPServerQos) Interval() int {
	return q.timeSleep
}

//AddAcceptedSessions all the accepts
func (q *TCPServerQos) AddAcceptedSessions() {
	atomic.AddInt32(&q.acceptedSessions, 1)
}

//AddClosedSessions all the closed sessions
func (q *TCPServerQos) AddClosedSessions() {
	atomic.AddInt32(&q.closedSessions, 1)
}

//AddReadPacket record all the read packets
func (q *TCPServerQos) AddReadPacket(packBytes int) {
	atomic.AddInt32(&q.readPacketCounts, 1)
	atomic.AddInt64(&q.readBytes, int64(packBytes))
}

//AddWritePacket record all the write packets
func (q *TCPServerQos) AddWritePacket(packBytes int) {
	atomic.AddInt32(&q.writePacketCounts, 1)
	atomic.AddInt64(&q.writeBytes, int64(packBytes))
}

var tcpServerQosInst *TCPServerQos

//GetTCPServerQos get tcpserverqos instance
func GetTCPServerQos() *TCPServerQos {
	return tcpServerQosInst
}

func init() {
	tcpServerQosInst = &TCPServerQos{timeSleep: 10}
	tcpServerQosInst.Enabled(true)
}
