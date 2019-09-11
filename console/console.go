package console

import (
	"os"
	"os/signal"
)

var CloseChan chan string

func Deamon() {

	select {
	case <-CloseChan:
		return
	}

}

func Parse() {

}

func init() {
	CloseChan = make(chan string)
}

func handleSignal(signalType os.Signal, handleFunc func(*chan os.Signal)) {
	ch := make(chan os.Signal)
	signal.Notify(ch, signalType)
	go handleFunc(&ch)
}

func handlePIPE(ch *chan os.Signal) {

}
