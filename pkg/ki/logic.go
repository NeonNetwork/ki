package ki

import (
	"fmt"
	"github.com/heartbytenet/go-lerpc/pkg/proto"
	"log"
)

type Logic struct {
	engine *Engine
}

func (logic *Logic) Engine() *Engine {
	return logic.engine
}

func (logic *Logic) Init() *Logic {
	if logic.engine == nil {
		log.Fatalln("engine is nil")
	}

	return logic
}

func (logic *Logic) Start() (err error) {
	return
}

func (logic *Logic) Close() (err error) {
	return
}

func (logic *Logic) RpcDataPull(key string) (result any, err error) {
	var (
		cmd *proto.ExecuteCommand
		res *proto.ExecuteResult
	)

	// init
	cmd = &proto.ExecuteCommand{}
	res = &proto.ExecuteResult{}

	cmd.
		SetNamespace("data").
		SetMethod("pull").
		SetParam("key", key)

	err = logic.Engine().
		RpcClient().
		Execute(cmd, res)
	if err != nil {
		return
	}

	if !res.Success {
		err = fmt.Errorf("failed at executing rpc: %v", res.Error)
		return
	}

	result = res.Payload["value"]
	return
}
