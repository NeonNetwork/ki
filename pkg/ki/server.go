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
	case "BINANCE_ORDERS":
		{
			err = server.client.HttpGetJson(
				"https://api.binance.com/api/v3/depth?symbol=ETHUSDT",
				&result)
			if err != nil {
				res.ToError("failed at fetching depth json: " + err.Error())
				return
			}

			// Go from [][]string -> [][]float64
			process := func(src []any) (dst []any, err error) {
				dst = make([]any, 0)

				for i, v := range src {
					a := v.([]any)
					s := make([]string, len(a))
					d := make([]float64, len(a))

					s[0] = a[0].(string)
					s[1] = a[1].(string)

					d[0], err = strconv.ParseFloat(s[0], 64)
					if err != nil {
						return
					}

					d[1], err = strconv.ParseFloat(s[1], 64)
					if err != nil {
						return
					}

					dst = append(dst, d)

					_ = i
				}

				return
			}

			var (
				asks any
				bids any
			)

			asks, err = process(result["asks"].([]any))
			if err != nil {
				return
			}

			bids, err = process(result["bids"].([]any))
			if err != nil {
				return
			}

			result = map[string]any{
				"value": map[string]any{
					"asks": asks,
					"bids": bids,
				},
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
