package ki

import (
	"github.com/neonnetwork/ki/pkg/structure"
)

func (engine *Engine) Compute() (err error) {
	err = engine.ComputeWindows()
	if err != nil {
		return
	}

	return
}

func (engine *Engine) ComputeWindows() (err error) {
	engine.WindowRootNode().IfPresent(func(value *structure.BinaryTreeNode[Window]) {
		err = engine.ComputeWindowsStep(value)
	})
	if err != nil {
		return
	}

	return
}

func (engine *Engine) ComputeWindowsStep(node *structure.BinaryTreeNode[Window]) (err error) {
	var (
		window Window
	)

	window = node.Value()

	window.SetSelected(engine.selected == node)
	window.SetIsRoot(engine.windows == node)

	if window.IsRoot() {
		window.SetPositionAbsolute(structure.NewVector2[int32](0, 0))
		window.SetSizeAbsolute(engine.Screen())
	}

	err = window.Compute()
	if err != nil {
		return
	}

	node.Left().IfPresent(func(value *structure.BinaryTreeNode[Window]) {
		err = engine.ComputeWindowsStep(value)
	})
	if err != nil {
		return
	}

	node.Right().IfPresent(func(value *structure.BinaryTreeNode[Window]) {
		err = engine.ComputeWindowsStep(value)
	})
	if err != nil {
		return
	}

	return
}
