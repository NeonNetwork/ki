package ki

import (
	"github.com/heartbytenet/bblib/objects"
	"github.com/heartbytenet/go-lerpc/pkg/lerpc"
	"github.com/heartbytenet/go-lerpc/pkg/proto"
	"log"
	"os"
	"strconv"
)

var (
	ServerDefaultPort = 12000
)

type Server struct {
	client *Client

	lerpc.Server
}

func (server *Server) Init() *Server {
	var (
		port int
		err  error
	)

	server.client = objects.Init[Client](&Client{})

	port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Println("failed at parsing $PORT, using default port:", ServerDefaultPort)
		port = ServerDefaultPort
	}

	server.Server.Init(&lerpc.ServerSettings{
		Port: port,
	})

	return server
}

func (server *Server) Start() (err error) {
	server.RegisterHandler("data::pull", server.HandleDataPull)

	err = server.Server.Start()
	if err != nil {
		return
	}

	return
}

func (server *Server) Close() (err error) {
	return
}

func (server *Server) HandleDataPull(cmd *proto.ExecuteCommand, res *proto.ExecuteResult) {
	var (
		result map[string]any
		value  any
		key    string
		err    error
	)

	// init
	result = map[string]any{}

	key, err = proto.GetCommandParam[string](cmd, "key").GetTry()
	if err != nil {
		res.ToError("invalid `key` param")
		return
	}

	switch key {
	case "BINANCE_PRICE":
		{
			err = server.client.HttpGetJson(
				"https://api.binance.com/api/v3/ticker/price?symbol=ETHUSDT",
				&result)
			if err != nil {
				res.ToError("failed at fetching price json: " + err.Error())
				return
			}

			value, err = strconv.ParseFloat(result["price"].(string), 64)
			if err != nil {
				res.ToError("failed at parsing price value: " + err.Error())
				return
			}

			result = map[string]any{
				"value": value,
			}
		}
	default:
		{
			break
		}
	}

	res.ToPayload(result)
	return
}
