package ki

import (
	"fmt"
	"github.com/heartbytenet/bblib/containers/optionals"
	"github.com/heartbytenet/bblib/containers/sync"
	"github.com/neonnetwork/ki/pkg/structure"
	"log"
	"reflect"
	"sort"
	"time"
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
		"BINANCE_PRICE",
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
		"BINANCE_PRICE_HISTORY",
		structure.NewCached[[]float64](
			make([]float64, 0),
			func(previous []float64) (result []float64, err error) {
				var (
					value float64
				)

				result = previous

				PoolGet[float64]("BINANCE_PRICE").
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
				for len(result) > 7200 {
					result = result[1:]
				}

				return
			},
			1000))

	PoolRegister(
		"RESOURCE_CPU",
		structure.NewCached[float64](
			0.0,
			func(_ float64) (value float64, err error) {
				var (
					data any
					flag bool
				)

				ENGINE.Apply(func(engine *Engine) {
					data, err = engine.Logic().RpcDataPull("RESOURCE_CPU")
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
		"RESOURCE_TOP",
		structure.NewCached[[]structure.Pair[string, float64]](
			nil,
			func(prev []structure.Pair[string, float64]) (result []structure.Pair[string, float64], err error) {
				var (
					data any
				)

				result = make([]structure.Pair[string, float64], 0)

				ENGINE.Apply(func(engine *Engine) {
					data, err = engine.Logic().RpcDataPull("RESOURCE_TOP")
					if err != nil {
						return
					}
				})
				if err != nil {
					return
				}

				for key, val := range data.(map[string]any) {
					value, ok := val.(float64)
					if !ok {
						continue
					}

					result = append(result, structure.NewPair(key, value))
				}

				sort.Slice(result, func(i, j int) bool {
					return result[i].B() > result[j].B()
				})

				return
			},
			1000))

	PoolRegister(
		"RESOURCE_CPU_HISTORY",
		structure.NewCached[[]float64](
			make([]float64, 0),
			func(previous []float64) (result []float64, err error) {
				var (
					value float64
				)

				result = previous

				PoolGet[float64]("RESOURCE_CPU").
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
				for len(result) > 7200 {
					result = result[1:]
				}

				return
			},
			1000))

	PoolRegister(
		"TEXT_LIST_DATA",
		structure.NewCached[[]string](
			make([]string, 0),
			func(previous []string) (result []string, err error) {
				result = previous

				value := fmt.Sprintf("Timestamp=%v", time.Now().Unix())

				result = append(result, value)

				for len(result) > 1024 {
					result = result[1:]
				}

				return
			},
			500))

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
