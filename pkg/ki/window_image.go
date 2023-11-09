package ki

import (
	"github.com/heartbytenet/bblib/objects"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
	"github.com/neonnetwork/ki/pkg/structure"
)

type WindowImage struct {
	id          uuid.UUID
	node        *structure.BinaryTreeNode[Window]
	isRoot      bool
	side        structure.BinaryTreeDirection
	position    structure.Vector2[int32]
	positionAbs structure.Vector2[int32]
	size        structure.Vector2[int32]
	sizeAbs     structure.Vector2[int32]
	axis        WindowSplitAxis
	selected    bool
	color       structure.Vector3[uint8]

	controller Controller
}

func (window *WindowImage) Id() uuid.UUID {
	return window.id
}

func (window *WindowImage) Type() WindowType {
	return WindowTypeImage
}

func (window *WindowImage) Node() *structure.BinaryTreeNode[Window] {
	return window.node
}

func (window *WindowImage) SetNode(value *structure.BinaryTreeNode[Window]) {
	window.node = value

	return
}

func (window *WindowImage) IsRoot() bool {
	return window.isRoot
}

func (window *WindowImage) SetIsRoot(value bool) {
	window.isRoot = value

	return
}

func (window *WindowImage) Side() structure.BinaryTreeDirection {
	return window.side
}

func (window *WindowImage) SetSide(value structure.BinaryTreeDirection) {
	window.side = value

	return
}

func (window *WindowImage) Position() structure.Vector2[int32] {
	return window.position
}

func (window *WindowImage) SetPosition(value structure.Vector2[int32]) {
	window.position = value

	return
}

func (window *WindowImage) PositionAbsolute() structure.Vector2[int32] {
	return window.positionAbs
}

func (window *WindowImage) SetPositionAbsolute(value structure.Vector2[int32]) {
	window.positionAbs = value

	return
}

func (window *WindowImage) Size() structure.Vector2[int32] {
	return window.size
}

func (window *WindowImage) SetSize(value structure.Vector2[int32]) {
	window.size = value

	return
}

func (window *WindowImage) SizeAbsolute() structure.Vector2[int32] {
	return window.sizeAbs
}

func (window *WindowImage) SetSizeAbsolute(value structure.Vector2[int32]) {
	window.sizeAbs = value

	return
}

func (window *WindowImage) Color() structure.Vector3[uint8] {
	return window.color
}

func (window *WindowImage) Selected() bool {
	return window.selected
}

func (window *WindowImage) SetSelected(value bool) {
	window.selected = value

	return
}

func (window *WindowImage) SplitAxis() WindowSplitAxis {
	return window.axis
}

func (window *WindowImage) SetSplitAxis(value WindowSplitAxis) {
	window.axis = value

	return
}

var (
	I = 0
	W = []string{"N", "G", "P"}
)

func (window *WindowImage) Init() *WindowImage {
	window.id = uuid.New()
	window.color = structure.NewVector3Random[byte](256)

	I++
	switch W[I%len(W)] {
	case "N":
		{
			window.SetController(objects.Init[ControllerNumber[float64]](&ControllerNumber[float64]{}))
			break
		}
	case "G":
		{
			window.SetController(objects.Init[ControllerGraph](&ControllerGraph{}))
			break
		}
	case "P":
		{
			window.SetController(objects.Init[ControllerPie](&ControllerPie{}))
			break
		}
	}

	return window
}

func (window *WindowImage) Box() structure.Box[int32] {
	return structure.NewBox[int32](
		window.Position(),
		window.Size())
}

func (window *WindowImage) BoxAbs() structure.Box[int32] {
	return structure.NewBox[int32](
		window.PositionAbsolute(),
		window.SizeAbsolute())
}

func (window *WindowImage) Compute() (err error) {
	var (
		parent   Window
		full     structure.Vector2[float64]
		position structure.Vector2[int32]
		size     structure.Vector2[int32]
	)

	full = structure.NewVector2[float64](EngineWindowUnit, EngineWindowUnit)

	window.
		Node().
		Prev().
		IfPresent(func(value *structure.BinaryTreeNode[Window]) {
			parent = value.Value()

			position, size = parent.PositionAbsolute(), parent.SizeAbsolute()

			position = window.
				Position().ToFloat64().
				Div(full).
				Mul(size.ToFloat64()).
				ToInt32().
				Add(position)

			size = window.
				Size().ToFloat64().
				Div(full).
				Mul(size.ToFloat64()).
				ToInt32()

			window.SetPositionAbsolute(position)
			window.SetSizeAbsolute(size)
		})

	err = window.Controller().Compute()
	if err != nil {
		return
	}

	return
}

func (window *WindowImage) Controller() Controller {
	return window.controller
}

func (window *WindowImage) SetController(value Controller) {
	window.controller = value

	value.SetWindow(window)

	return
}

func (window *WindowImage) Render() (err error) {
	position := window.PositionAbsolute()
	size := window.SizeAbsolute()

	GRAPHICS.Apply(func(graphics *Graphics) {
		rl.DrawTextureRec(
			graphics.textureWindow,
			rl.NewRectangle(0, 0, float32(size.X()), float32(size.Y())),
			position.ToRaylib(),
			window.Color().ToColor())
	})

	err = window.Controller().Render()
	if err != nil {
		return
	}

	return
}
