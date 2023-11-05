package ki

import (
	"math"
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
	controller.ticks++
	
	controller.AddValue(math.Sin(float64(controller.ticks) / 60.0))
	
	if (int32(len(controller.Value())) * 2) > controller.Window().BoxAbs().Size().X() {
		controller.value = controller.value[1:]
	}

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
