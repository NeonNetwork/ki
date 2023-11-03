package ki

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
	"github.com/neonnetwork/ki/pkg/structure"
)

type WindowNone struct {
	id       uuid.UUID
	node     *structure.BinaryTreeNode[Window]
	axis     WindowSplitAxis
	selected bool
}

func (window *WindowNone) Id() uuid.UUID {
	return window.id
}

func (window *WindowNone) Type() WindowType {
	return WindowTypeNone
}

func (window *WindowNone) Node() *structure.BinaryTreeNode[Window] {
	return window.node
}

func (window *WindowNone) SetNode(value *structure.BinaryTreeNode[Window]) {
	window.node = value

	return
}

func (window *WindowNone) IsRoot() bool {
	return false
}

func (window *WindowNone) SetIsRoot(value bool) {
	return
}

func (window *WindowNone) Side() structure.BinaryTreeDirection {
	return structure.BinaryTreeLeft
}

func (window *WindowNone) SetSide(value structure.BinaryTreeDirection) {
	return
}

func (window *WindowNone) Position() structure.Vector2[int32] {
	return structure.NewVector2[int32](0, 0)
}

func (window *WindowNone) SetPosition(value structure.Vector2[int32]) {
	return
}

func (window *WindowNone) PositionAbsolute() structure.Vector2[int32] {
	return structure.NewVector2[int32](0, 0)
}

func (window *WindowNone) SetPositionAbsolute(value structure.Vector2[int32]) {
	return
}

func (window *WindowNone) Size() structure.Vector2[int32] {
	return structure.NewVector2[int32](EngineWindowUnit, EngineWindowUnit)
}

func (window *WindowNone) SetSize(value structure.Vector2[int32]) {
	return
}

func (window *WindowNone) SizeAbsolute() structure.Vector2[int32] {
	return structure.NewVector2[int32](0, 0)
}

func (window *WindowNone) SetSizeAbsolute(value structure.Vector2[int32]) {
	return
}

func (window *WindowNone) Color() structure.Vector3[uint8] {
	return structure.NewVector3[uint8](0, 0, 0)
}

func (window *WindowNone) Selected() bool {
	return window.selected
}

func (window *WindowNone) SetSelected(value bool) {
	window.selected = value

	return
}

func (window *WindowNone) SplitAxis() WindowSplitAxis {
	return window.axis
}

func (window *WindowNone) SetSplitAxis(value WindowSplitAxis) {
	window.axis = value

	return
}

func (window *WindowNone) Init() *WindowNone {
	window.id = uuid.New()

	return window
}

func (window *WindowNone) Box() structure.Box[int32] {
	return structure.NewBox[int32](
		window.Position(),
		window.Size())
}

func (window *WindowNone) BoxAbs() structure.Box[int32] {
	return structure.NewBox[int32](
		window.PositionAbsolute(),
		window.SizeAbsolute())
}

func (window *WindowNone) Render() (err error) {
	if window.Selected() {
		rl.DrawRectangleLinesEx(
			window.BoxAbs().ToRaylibRectangle(),
			2,
			rl.RayWhite)
	}

	return
}
