package ki

import (
	"github.com/neonnetwork/ki/pkg/structure"
)

type ControllerGraph struct {
	ticks int64
	value []float64

	ControllerBase
}

func (controller *ControllerGraph) Value() []float64 {
	return controller.value
}

func (controller *ControllerGraph) SetValue(value []float64) {
	controller.value = value

	return
}

func (controller *ControllerGraph) AddValue(value float64) {
	controller.value = append(controller.value, value)

	return
}

func (controller *ControllerGraph) Init() *ControllerGraph {
	controller.ControllerBase.Init()

	controller.value = make([]float64, 0)
	controller.ticks = 0

	return controller
}

func (controller *ControllerGraph) Compute() (err error) {
	var (
		value []float64
	)

	controller.ticks++

	PoolGet[[]float64]("RESOURCE_CPU_HISTORY").
		IfPresent(func(cached *structure.Cached[[]float64]) {
			value, err = cached.Get()
			if err != nil {
				return
			}
		})
	if err != nil {
		return
	}

	for int32(len(value)) > (controller.Window().BoxAbs().W() / 8) {
		value = value[1:]
	}

	controller.SetValue(value)

	return
}

func (controller *ControllerGraph) Render() (err error) {
	GRAPHICS.Apply(func(graphics *Graphics) {
		err = graphics.DrawGraph(
			controller.Value(),
			controller.Window().BoxAbs())
		if err != nil {
			return
		}
	})
	if err != nil {
		return
	}

	return
}
