package console

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
