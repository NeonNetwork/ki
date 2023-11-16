package structure

import (
	"github.com/heartbytenet/bblib/containers/sync"
	"github.com/heartbytenet/bblib/objects"
	"log"
	"reflect"
	"time"
)

// Cached is a generic struct used to keep track of data that evolves with time
// Cached.supply is called under a goroutine each time Cached.value has expired
type Cached[T any] struct {
	value  *sync.Mutex[T]
	supply func(T) (T, error)
	error  error
	last   *sync.Mutex[int64]
	life   int64
}

func NewCached[T any](value T, supply func(T) (T, error), life int64) *Cached[T] {
	return objects.Init[Cached[T]](&Cached[T]{
		value:  sync.NewMutex(value),
		supply: supply,
		life:   life,
	})
}

func (cached *Cached[T]) Init() *Cached[T] {
	if cached.supply == nil {
		log.Fatalln("supply function is nil")
	}

	cached.last = sync.NewMutex[int64](0)

	return cached
}

func (cached *Cached[T]) Now() int64 {
	return time.Now().UnixMilli()
}

func (cached *Cached[T]) Alive() bool {
	if (cached.Now() - cached.last.Get()) >= cached.life {
		return false
	}

	return true
}

func (cached *Cached[T]) Update() {
	var (
		result T
		value  T
		err    error
	)

	value = cached.value.Get()

	result, err = cached.supply(value)
	if err != nil {
		log.Println("failed at updating cached value using supplier for", reflect.TypeOf(cached), "error:", err)
		return
	}

	cached.value.Map(func(_ T) T {
		return result
	})
}

func (cached *Cached[T]) Tick() {
	cached.last.Set(cached.Now())
}

func (cached *Cached[T]) Get() (result T, err error) {
	if !cached.Alive() {
		if cached.last.Get() == 0 {
			cached.Tick()
			cached.Update()
		} else {
			cached.Tick()
			go cached.Update()
		}
	}

	cached.value.Apply(func(value T) {
		result = value
	})

	return
}

func (cached *Cached[T]) GetMust() (result T) {
	var (
		err error
	)

	result, err = cached.Get()
	if err != nil {
		panic(err)
	}

	return
}
