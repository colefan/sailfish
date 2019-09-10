package main

import (
	"fmt"

	"sailfish/gate"
)

func main() {
	g := new(gate.Gate)
	err := g.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	g.Logger().Info("hello")
	g.Run()
	g.Daemon()
	//time.Sleep(1 * time.Minute)
}
