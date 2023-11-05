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
	selected    bool
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
