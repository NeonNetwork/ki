package ki

import (
	"fmt"

	"github.com/neonnetwork/ki/pkg/structure"
)

type ControllerNumber[T structure.Number] struct {
	value T

	ControllerBase
}

func (controller *ControllerNumber[T]) Value() T {
	return controller.value
}

func (controller *ControllerNumber[T]) SetValue(value T) {
	controller.value = value

	return
}

func (controller *ControllerNumber[T]) AddValue(value T) {
	controller.value = controller.Value() + value

	return
}

func (controller *ControllerNumber[T]) Init() *ControllerNumber[T] {
	controller.ControllerBase.Init()

	return controller
}

func (controller *ControllerNumber[T]) Compute() (err error) {
	controller.AddValue(127)

	return
}

func (controller *ControllerNumber[T]) Render() (err error) {
	GRAPHICS.Apply(func(graphics *Graphics) {
		err = graphics.DrawTextCentered(
			fmt.Sprintf("%v", controller.Value()),
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
