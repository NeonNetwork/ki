package main

import (
	"github.com/heartbytenet/bblib/objects"
	"github.com/neonnetwork/ki/pkg/ki"
	"log"
)

var (
	ENGINE *ki.Engine
)

func init() {
	ENGINE = objects.Init[ki.Engine](&ki.Engine{})
}

func main() {
	var (
		err error
	)

	err = ENGINE.Start()
	if err != nil {
		log.Println("failed at starting engine:", err)
		return
	}

	ENGINE.Wait()

	err = ENGINE.Close()
	if err != nil {
		log.Println("failed at closing engine:", err)
		return
	}
}
