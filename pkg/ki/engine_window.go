package ki

import (
	"github.com/neonnetwork/ki/pkg/structure"
)

func (engine *Engine) WindowChildAdd(
	window *structure.BinaryTreeNode[Window],
	child Window,
	axis WindowSplitAxis,
	direction structure.BinaryTreeDirection,
) (result *structure.BinaryTreeNode[Window]) {
	var (
		data structure.Pair[structure.Vector2[int32], structure.Vector2[int32]]
	)

	if window == nil {
		data = structure.NewPair(
			structure.NewVector2[int32](0, 0),
			structure.NewVector2[int32](EngineWindowUnit, EngineWindowUnit))
	} else {
		window.Value().SetSplitAxis(axis)
		data = window.Value().Split(direction)
	}

	child.SetSide(direction)
	child.SetPosition(data.A())
	child.SetSize(data.B())

	if window == nil {
		result = structure.NewBinaryTreeNode(child)

		engine.windows = result
	} else {
		result = window.ChildAdd(child, direction)
	}

	return
}
