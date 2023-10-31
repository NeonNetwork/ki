package ki

import (
	"github.com/google/uuid"
	"github.com/neonnetwork/ki/pkg/structure"
)

const (
	WindowSplitHorizontal WindowSplitAxis = iota
	WindowSplitVertical
)

type WindowSplitAxis int8

type Window interface {
	Id() uuid.UUID
	IsRoot() bool
	SetIsRoot(bool)
	Side() structure.BinaryTreeDirection
	SetSide(structure.BinaryTreeDirection)
	Position() structure.Vector2[int32]
	SetPosition(vector structure.Vector2[int32])
	PositionAbsolute() structure.Vector2[int32]
	SetPositionAbsolute(structure.Vector2[int32])
	Size() structure.Vector2[int32]
	SetSize(vector structure.Vector2[int32])
	SizeAbsolute() structure.Vector2[int32]
	SetSizeAbsolute(structure.Vector2[int32])
	SplitAxis() WindowSplitAxis
	SetSplitAxis(axis WindowSplitAxis)
	Selected() bool
	SetSelected(bool)

	Box() structure.Box[int32]
	BoxAbs() structure.Box[int32]

	Render() error
	Split(direction structure.BinaryTreeDirection) structure.Pair[structure.Vector2[int32], structure.Vector2[int32]]
}

func WindowSplitCommon(window Window, direction structure.BinaryTreeDirection) (result structure.Pair[structure.Vector2[int32], structure.Vector2[int32]]) {
	var (
		axis WindowSplitAxis
		vert bool
		zero structure.Vector2[int32]
		size structure.Vector2[int32]
		diff structure.Vector2[int32]
	)

	axis = window.SplitAxis()                // W          || H
	vert = axis == WindowSplitVertical       // true       || false
	zero = structure.NewVector2[int32](0, 0) // [0    , 0] || [0, 0]
	size = structure.NewVector2[int32](EngineWindowUnit, EngineWindowUnit)
	size = size.Div(structure.NewVector2[int32](2, 1).Order(vert)) // [W / 2, H] || [W, H / 2]
	diff = size.Mul(structure.NewVector2[int32](1, 0).Order(vert)) // [W / 2, 0] || [0, H / 2]

	if direction == structure.BinaryTreeLeft {
		result = structure.NewPair(
			window.Position().Add(zero),
			size)

		window.SetSize(size)
		window.SetPosition(window.Position().Add(diff))
	} else {
		result = structure.NewPair(
			window.Position().Add(diff),
			size)

		window.SetSize(size)
		window.SetPosition(window.Position().Add(zero))
	}

	return
}
