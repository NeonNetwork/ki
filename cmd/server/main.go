package main

import (
	"github.com/heartbytenet/bblib/objects"
	"github.com/neonnetwork/ki/pkg/ki"
	"log"
)

var (
	SERVER *ki.Server
)

func init() {
	SERVER = objects.Init[ki.Server](&ki.Server{})
}

func main() {
	var (
		err error
	)

	err = SERVER.Start()
	if err != nil {
		log.Println("failed at starting server:", err)
		return
	}

	err = SERVER.Close()
	if err != nil {
		log.Println("failed at closing server:", err)
		return
	}
}
