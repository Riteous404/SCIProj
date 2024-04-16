package core

import (
	"SCIProj/initialize"
	"log"
)

func RunWindowsServer(Addr string) {
	initialize.Redis()
	initialize.JWT()
	r := initialize.Routers()
	if err := r.Run(Addr); err != nil {
		log.Fatal("Server start failed!")
	}
}
