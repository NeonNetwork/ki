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

const (
	WindowTypeNone WindowType = iota
	WindowTypeSplit
	WindowTypeImage
)

type WindowType int32

type Window interface {
	Id() uuid.UUID
	Type() WindowType
	Node() *structure.BinaryTreeNode[Window]
	SetNode(value *structure.BinaryTreeNode[Window])
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
	Color() structure.Vector3[uint8]
	SetColor(structure.Vector3[uint8])

	Box() structure.Box[int32]
	BoxAbs() structure.Box[int32]

	Controller() Controller
	SetController(Controller)

	Compute() error
	Render() error
}
