package ki

import (
	"fmt"
	"github.com/heartbytenet/bblib/containers/optionals"
	"github.com/heartbytenet/bblib/containers/sync"
	"github.com/neonnetwork/ki/pkg/structure"
	"log"
	"reflect"
)

var (
	POOL *sync.Mutex[*Pool] = nil
)

type Pool struct {
	engine *Engine

	data map[string]any
}

func (pool *Pool) Init() *Pool {
	if pool.engine == nil {
		log.Fatalln("engine is nil")
	}

	pool.data = make(map[string]any)

	POOL = sync.NewMutex(pool)

	PoolRegister(
		"BINANCE_TICKER_VALUE",
		structure.NewCached[float64](
			0.0,
			func(_ float64) (value float64, err error) {
				var (
					data any
					flag bool
				)

				ENGINE.Apply(func(engine *Engine) {
					data, err = engine.Logic().RpcDataPull("BINANCE_PRICE")
					if err != nil {
						return
					}
				})
				if err != nil {
					return
				}

				value, flag = data.(float64)
				if !flag {
					err = fmt.Errorf(
						"failed at converting %v -> %v",
						reflect.TypeOf(data),
						reflect.TypeOf(value))
					return
				}

				return
			},
			1000))

	PoolRegister(
		"BINANCE_TICKER_HISTORY",
		structure.NewCached[[]float64](
			make([]float64, 0),
			func(previous []float64) (result []float64, err error) {
				var (
					value float64
				)

				result = previous

				PoolGet[float64]("BINANCE_TICKER_VALUE").
					IfPresent(func(cached *structure.Cached[float64]) {
						value, err = cached.Get()
						if err != nil {
							return
						}
					})
				if err != nil {
					return
				}

				result = append(result, value)
				if len(result) > 3600 {
					result = result[1:]
				}

				return
			},
			1000))

	return pool
}

func (pool *Pool) Register(key string, value any) {
	pool.data[key] = value
}

func (pool *Pool) Get(key string) (result any) {
	return pool.data[key]
}

func PoolRegister[T any](key string, value *structure.Cached[T]) {
	POOL.Apply(func(pool *Pool) {
		pool.Register(key, value)
	})
}

func PoolGet[T any](key string) (result optionals.Optional[*structure.Cached[T]]) {
	POOL.Apply(func(pool *Pool) {
		result = optionals.FromNillable[*structure.Cached[T]](pool.Get(key))
	})

	return
}
