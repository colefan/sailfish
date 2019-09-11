package main

import (
	"fmt"

	"github.com/colefan/sailfish/gate"
	"github.com/colefan/sailfish/log"
)

func main() {
	g := new(gate.Gate)
	err := g.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	log.Info("hello")
	g.Run()
	g.Daemon()
	g.ShutDown()
	log.Close()
	//time.Sleep(1 * time.Minute)
}
