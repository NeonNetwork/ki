package ki

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
	"github.com/neonnetwork/ki/pkg/structure"
)

type WindowRoot struct {
	id       uuid.UUID
	node     *structure.BinaryTreeNode[Window]
	axis     WindowSplitAxis
	selected bool
}

func (window *WindowRoot) Id() uuid.UUID {
	return window.id
}

func (window *WindowRoot) Node() *structure.BinaryTreeNode[Window] {
	return window.node
}

func (window *WindowRoot) SetNode(value *structure.BinaryTreeNode[Window]) {
	window.node = value

	return
}

func (window *WindowRoot) IsRoot() bool {
	return true
}

func (window *WindowRoot) SetIsRoot(value bool) {
	return
}

func (window *WindowRoot) Side() structure.BinaryTreeDirection {
	return structure.BinaryTreeLeft
}

func (window *WindowRoot) SetSide(value structure.BinaryTreeDirection) {
	return
}

func (window *WindowRoot) Position() structure.Vector2[int32] {
	return structure.NewVector2[int32](0, 0)
}

func (window *WindowRoot) SetPosition(value structure.Vector2[int32]) {
	return
}

func (window *WindowRoot) PositionAbsolute() structure.Vector2[int32] {
	return structure.NewVector2[int32](0, 0)
}

func (window *WindowRoot) SetPositionAbsolute(value structure.Vector2[int32]) {
	return
}

func (window *WindowRoot) Size() structure.Vector2[int32] {
	return structure.NewVector2[int32](EngineWindowUnit, EngineWindowUnit)
}

func (window *WindowRoot) SetSize(value structure.Vector2[int32]) {
	return
}

func (window *WindowRoot) SizeAbsolute() structure.Vector2[int32] {
	return structure.NewVector2[int32](0, 0)
}

func (window *WindowRoot) SetSizeAbsolute(value structure.Vector2[int32]) {
	return
}

func (window *WindowRoot) Color() structure.Vector3[uint8] {
	return structure.NewVector3[uint8](0, 0, 0)
}

func (window *WindowRoot) Selected() bool {
	return window.selected
}

func (window *WindowRoot) SetSelected(value bool) {
	window.selected = value

	return
}

func (window *WindowRoot) SplitAxis() WindowSplitAxis {
	return window.axis
}

func (window *WindowRoot) SetSplitAxis(value WindowSplitAxis) {
	window.axis = value

	return
}

func (window *WindowRoot) Init() *WindowRoot {
	window.id = uuid.New()

	return window
}

func (window *WindowRoot) Box() structure.Box[int32] {
	return structure.NewBox[int32](
		window.Position(),
		window.Size())
}

func (window *WindowRoot) BoxAbs() structure.Box[int32] {
	return structure.NewBox[int32](
		window.PositionAbsolute(),
		window.SizeAbsolute())
}

func (window *WindowRoot) Render() (err error) {
	if window.Selected() {
		rl.DrawRectangleLinesEx(
			window.BoxAbs().ToRaylibRectangle(),
			2,
			rl.RayWhite)
	}

	return
}

func (window *WindowRoot) Split(direction structure.BinaryTreeDirection) (result structure.Pair[structure.Vector2[int32], structure.Vector2[int32]]) {
	return WindowSplitCommon(window, direction)
}
