package ki

import (
	"github.com/heartbytenet/bblib/containers/optionals"
	"github.com/neonnetwork/ki/pkg/structure"
	"os"
)

type ControllerList struct {
	value []string

	textSize  int32
	textColor optionals.Optional[structure.Vector3[uint8]]

	ControllerBase
}

func (controller *ControllerList) Value() []string {
	return controller.value
}

func (controller *ControllerList) SetValue(value []string) {
	controller.value = value

	return
}

func (controller *ControllerList) AddValue(value string) {
	controller.value = append(controller.Value(), value)

	return
}

func (controller *ControllerList) Init() *ControllerList {
	controller.ControllerBase.Init()

	controller.textSize = 48
	controller.textColor = optionals.None[structure.Vector3[uint8]]()

	return controller
}

func (controller *ControllerList) Compute() (err error) {
	if os.Getenv("KI_SETUP") == "BINANCE" {
		PoolGet[[]string]("BINANCE_LIST").IfPresent(func(cache *structure.Cached[[]string]) {
			value := cache.GetMust()

			controller.SetValue(value)
		})
	} else {
		PoolGet[[]string]("RESOURCE_LIST").IfPresent(func(cache *structure.Cached[[]string]) {
			value := cache.GetMust()

			controller.SetValue(value)
		})
	}

	if controller.textColor.IsEmpty() {
		controller.textColor = optionals.Some[structure.Vector3[uint8]](
			controller.Window().Color().
				ToFloat64().
				Mul(structure.NewVector3[float64](1.0, 1.0, 1.0).Sub(structure.NewVector3Random[float64](400).Div(structure.NewVector3[float64](100.0, 100.0, 100.0)))).
				ToUint8())
	}

	return
}

func (controller *ControllerList) Render() (err error) {
	box := controller.Window().BoxRender()

	colorX := controller.Window().Color().ToColor()
	colorX.A = 32
	colorY := controller.textColor.GetDefault(structure.NewVector3[uint8](0, 0, 0)).ToColor()
	colorY.A = 32

	textSize := controller.textSize
	textLen := box.H() / textSize

	textData := controller.Value()
	for len(controller.Value()) > int(textLen) {
		if len(controller.Value()) > 0 {
			controller.SetValue(textData[1:])
		}
		textData = controller.Value()
	}

	if len(textData) > int(textLen) {
		textData = textData[:textLen]
	}

	GRAPHICS.Apply(func(graphics *Graphics) {
		err = graphics.DrawRectangle(
			box,
			colorX)
	})
	if err != nil {
		return
	}

	GRAPHICS.Apply(func(graphics *Graphics) {
		vEnd := box.Position().Add(box.Size().Mul(structure.NewVector2[int32](0, 1)))
		delta := structure.NewVector2[int32](0, 0)

		for index, value := range textData {
			vPos := vEnd.Sub(delta)
			vSize := structure.NewVector2[int32](box.W(), textSize)

			textBoxColor := colorX
			if (index % 2) == 0 {
				textBoxColor = colorY
			}

			err = graphics.DrawRectangle(
				structure.NewBox[int32](vPos, vSize),
				textBoxColor)
			if err != nil {
				return
			}

			err = graphics.DrawText(
				textData[len(textData)-1-index],
				vPos,
				float64(textSize))
			if err != nil {
				return
			}

			delta = delta.Add(structure.NewVector2[int32](0, textSize))

			_, _ = index, value
		}
	})
	if err != nil {
		return
	}

	return
}
