package ki

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/neonnetwork/ki/pkg/structure"
	"os"
)

type ControllerGraph struct {
	ticks int64
	value []float64

	valueYmin  float64
	valueYmax  float64
	valueYline float64

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

	if os.Getenv("KI_SETUP") == "BINANCE" {
		PoolGet[[]float64]("BINANCE_PRICE_HISTORY").
			IfPresent(func(cached *structure.Cached[[]float64]) {
				value, err = cached.Get()
				if err != nil {
					return
				}
			})

		if err != nil {
			return
		}
	} else {
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
	}

	for int32(len(value)) > (controller.Window().BoxAbs().W() / 8) {
		value = value[1:]
	}

	controller.SetValue(value)

	controller.valueYmin, controller.valueYmax = structure.MinMax[float64](value...)

	// Compute mouse
	cursor := controller.Window().CursorPosition()

	cursorY := controller.Window().BoxAbs().H() - cursor.Y()

	controller.valueYline = controller.valueYmin + (controller.valueYmax-controller.valueYmin)*(float64(cursorY)/float64(controller.Window().BoxAbs().H()))

	return
}

func (controller *ControllerGraph) Render() (err error) {
	box := controller.Window().BoxAbs()
	cursor := controller.Window().CursorPosition()

	GRAPHICS.Apply(func(graphics *Graphics) {
		err = graphics.DrawGraph(
			controller.Value(),
			controller.valueYmin,
			controller.valueYmax,
			controller.Window().BoxAbs())
		if err != nil {
			return
		}

		err = graphics.DrawLine(
			box.Position().Add(cursor.Mul(structure.NewVector2[int32](0, 1))),
			box.Position().Add(cursor.Mul(structure.NewVector2[int32](0, 1))).Add(box.Size().Mul(structure.NewVector2[int32](1, 0))),
			1.0,
			rl.Red)
		if err != nil {
			return
		}

		err = graphics.DrawLine(
			box.Position().Add(cursor.Mul(structure.NewVector2[int32](1, 0))),
			box.Position().Add(cursor.Mul(structure.NewVector2[int32](1, 0))).Add(box.Size().Mul(structure.NewVector2[int32](0, 1))),
			1.0,
			rl.Red)
		if err != nil {
			return
		}

		err = graphics.DrawText(
			fmt.Sprintf("%0.4f", controller.valueYline),
			box.Position().Add(cursor).Add(structure.NewVector2[int32](8, 8)),
			32.0)
		if err != nil {
			return
		}
	})
	if err != nil {
		return
	}

	return
}
