package main

import (
	"fmt"
	"time"

	"github.com/daveilers/hs1xxplug"
)

func main() {

	//	HS100.adserv.osd.wednet.edu

	plug := hs1xxplug.Hs1xxPlug{
		IPAddress: "HS100.adserv.osd.wednet.edu",
		Timeout:   time.Second,
	}
	err := plug.TurnOff()
	if err != nil {
		fmt.Println("err:", err)
	}
	time.Sleep(3 * time.Second)
	err = plug.TurnOn()
	if err != nil {
		fmt.Println("err:", err)
	}

}
