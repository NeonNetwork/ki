package ki

import "github.com/neonnetwork/ki/pkg/structure"

func (engine *Engine) HandleCursor() (err error) {
	cursor := engine.Cursor().Sub(structure.NewVector2[int32](EngineWindowGaps, EngineWindowGaps))

	err = engine.WindowListWalkNode(func(node *structure.BinaryTreeNode[Window]) error {
		window := node.Value()

		box := window.BoxAbs()

		window.SetCursorPosition(cursor.Sub(box.Position()))

		return nil
	})
	if err != nil {
		return
	}

	return
}
