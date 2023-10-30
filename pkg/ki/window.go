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
	Position() structure.Vector2[int32]
	Size() structure.Vector2[int32]
	Box() structure.Box[int32]
	Render(screen structure.Vector2[int32], cursor structure.Vector2[int32]) error
	Split(axis WindowSplitAxis, direction structure.BinaryTreeDirection) structure.Pair[structure.Vector2[int32], structure.Vector2[int32]]
	SplitAxis() WindowSplitAxis
	SetSplitAxis(axis WindowSplitAxis)
	SetPosition(vector structure.Vector2[int32])
	SetSize(vector structure.Vector2[int32])

	Selected() bool
	SetSelected(bool)
}

func WindowSplitCommon(window Window, direction structure.BinaryTreeDirection) (result structure.Pair[structure.Vector2[int32], structure.Vector2[int32]]) {
	var (
		axis WindowSplitAxis
		vert bool
		zero structure.Vector2[int32]
		size structure.Vector2[int32]
		diff structure.Vector2[int32]
	)

	axis = window.SplitAxis()                                      // W          || H
	vert = axis == WindowSplitVertical                             // true       || false
	zero = structure.NewVector2[int32](0, 0)                       // [0    , 0] || [0, 0]
	size = window.Size().Copy()                                    // [W    , H] || [W, H]
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
