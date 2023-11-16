package ki

import (
	"github.com/google/uuid"
	"github.com/neonnetwork/ki/pkg/structure"
)

type WindowSplit struct {
	id          uuid.UUID
	isRoot      bool
	node        *structure.BinaryTreeNode[Window]
	side        structure.BinaryTreeDirection
	position    structure.Vector2[int32]
	positionAbs structure.Vector2[int32]
	size        structure.Vector2[int32]
	sizeAbs     structure.Vector2[int32]
	axis        WindowSplitAxis
	color       structure.Vector3[uint8]
	cursor      structure.Vector2[int32]
	selected    bool
	gaps        int32
}

func (window *WindowSplit) Id() uuid.UUID {
	return window.id
}

func (window *WindowSplit) Type() WindowType {
	return WindowTypeSplit
}

func (window *WindowSplit) Node() *structure.BinaryTreeNode[Window] {
	return window.node
}

func (window *WindowSplit) SetNode(value *structure.BinaryTreeNode[Window]) {
	window.node = value

	return
}

func (window *WindowSplit) IsRoot() bool {
	return window.isRoot
}

func (window *WindowSplit) SetIsRoot(value bool) {
	window.isRoot = value

	return
}

func (window *WindowSplit) Side() structure.BinaryTreeDirection {
	return window.side
}

func (window *WindowSplit) SetSide(value structure.BinaryTreeDirection) {
	window.side = value

	return
}

func (window *WindowSplit) Position() structure.Vector2[int32] {
	return window.position
}

func (window *WindowSplit) SetPosition(value structure.Vector2[int32]) {
	window.position = value

	return
}

func (window *WindowSplit) PositionAbsolute() structure.Vector2[int32] {
	return window.positionAbs
}

func (window *WindowSplit) SetPositionAbsolute(value structure.Vector2[int32]) {
	window.positionAbs = value

	return
}

func (window *WindowSplit) Gaps() int32 {
	return window.gaps
}

func (window *WindowSplit) SetGaps(value int32) {
	window.gaps = value

	return
}

func (window *WindowSplit) Color() structure.Vector3[uint8] {
	return window.color
}

func (window *WindowSplit) SetColor(value structure.Vector3[uint8]) {
	window.color = value

	return
}

func (window *WindowSplit) Size() structure.Vector2[int32] {
	return window.size
}

func (window *WindowSplit) SetSize(value structure.Vector2[int32]) {
	window.size = value

	return
}

func (window *WindowSplit) SizeAbsolute() structure.Vector2[int32] {
	return window.sizeAbs
}

func (window *WindowSplit) SetSizeAbsolute(value structure.Vector2[int32]) {
	window.sizeAbs = value

	return
}

func (window *WindowSplit) Selected() bool {
	return window.selected
}

func (window *WindowSplit) SetSelected(value bool) {
	window.selected = value

	return
}

func (window *WindowSplit) SplitAxis() WindowSplitAxis {
	return window.axis
}

func (window *WindowSplit) SetSplitAxis(value WindowSplitAxis) {
	window.axis = value

	return
}

func (window *WindowSplit) Init() *WindowSplit {
	window.id = uuid.New()
	window.gaps = 8

	window.SetPosition(structure.NewVector2[int32](0, 0))
	window.SetPositionAbsolute(window.Position())
	window.SetSize(structure.NewVector2[int32](EngineWindowUnit, EngineWindowUnit))
	window.SetSizeAbsolute(window.Size())

	return window
}

func (window *WindowSplit) Box() structure.Box[int32] {
	return structure.NewBox[int32](
		window.Position(),
		window.Size())
}

func (window *WindowSplit) BoxAbs() structure.Box[int32] {
	return structure.NewBox[int32](
		window.PositionAbsolute(),
		window.SizeAbsolute())
}

func (window *WindowSplit) BoxRender() structure.Box[int32] {
	return structure.NewBox[int32](
		window.PositionAbsolute().Add(structure.NewVector2[int32](8, 8)),
		window.SizeAbsolute().Sub(structure.NewVector2[int32](8, 8)))
}

func (window *WindowSplit) CursorPosition() structure.Vector2[int32] {
	return window.cursor
}

func (window *WindowSplit) SetCursorPosition(value structure.Vector2[int32]) {
	window.cursor = value

	return
}

func (window *WindowSplit) Compute() (err error) {
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

	return
}

func (window *WindowSplit) Controller() Controller {
	return nil
}

func (window *WindowSplit) SetController(_ Controller) {
	return
}

func (window *WindowSplit) Render() (err error) {
	return
}
