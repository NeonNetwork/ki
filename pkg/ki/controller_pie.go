package ki

import (
	"fmt"
	"github.com/neonnetwork/ki/pkg/structure"
	"math"
	"os"
)

type ControllerPie struct {
	value []structure.Pair[string, float64]

	ControllerBase
}

func (controller *ControllerPie) Value() []structure.Pair[string, float64] {
	return controller.value
}

func (controller *ControllerPie) SetValue(value []structure.Pair[string, float64]) {
	controller.value = value

	return
}

func (controller *ControllerPie) Init() *ControllerPie {
	controller.ControllerBase.Init()

	return controller
}

func (controller *ControllerPie) Compute() (err error) {
	if os.Getenv("KI_SETUP") == "BINANCE" {
		controller.SetValue(make([]structure.Pair[string, float64], 0))
	} else {
		PoolGet[[]structure.Pair[string, float64]]("RESOURCE_TOP").
			IfPresent(func(cached *structure.Cached[[]structure.Pair[string, float64]]) {
				controller.SetValue(cached.GetMust())
			})
	}

	return
}

func (controller *ControllerPie) Render() (err error) {
	box := controller.Window().BoxAbs()
	pos := box.Position()

	value := controller.Value()
	total := 0.0
	for _, v := range value {
		total += math.Abs(v.B())
	}

	i := 0
	a := 0.0

	pos = box.Position()
	for _, it := range controller.Value() {
		k, v := it.A(), it.B()

		i++
		r := v / total

		GRAPHICS.Apply(func(graphics *Graphics) {
			err = graphics.DrawCircleSector(
				box.Center(),
				float64(box.Size().Min())/2.0,
				structure.NewPair[float64](a, a+r*360.0),
				structure.NewVector3[uint8](16*i, 16*i, 16*i))
			if err != nil {
				return
			}

			a += r * 360.0
		})
		if err != nil {
			return
		}

		pos = pos.Add(structure.NewVector2[int32](0, 32))

		_ = k
	}

	pos = box.Position()
	for _, it := range controller.Value() {
		k, v := it.A(), it.B()

		GRAPHICS.Apply(func(graphics *Graphics) {
			err = graphics.DrawText(
				fmt.Sprintf("%v - %v", k, v),
				pos,
				32.0)
			if err != nil {
				return
			}
		})

		pos = pos.Add(structure.NewVector2[int32](0, 32))
	}

	return
}
