package ki

import "github.com/neonnetwork/ki/pkg/structure"

func (engine *Engine) WindowListWalkNode(fn func (*structure.BinaryTreeNode[Window]) error) (err error) {
	engine.WindowRootNode().IfPresent(func(value *structure.BinaryTreeNode[Window]) {
		err = engine.WindowListWalkNodeStep(value, fn)
	})
	if err != nil {
		return
	}
	
	return
}

func (engine *Engine) WindowListWalkNodeStep(node *structure.BinaryTreeNode[Window], fn func (*structure.BinaryTreeNode[Window]) error) (err error) {
	err = fn(node)
	if err != nil {
		return
	}
	
	node.Left().IfPresent(func(value *structure.BinaryTreeNode[Window]) {
		err = engine.WindowListWalkNodeStep(value, fn)
	})
	if err != nil {
		return
	}
	
	node.Right().IfPresent(func(value *structure.BinaryTreeNode[Window]) {
		err = engine.WindowListWalkNodeStep(value, fn)
	})
	if err != nil {
		return
	}
	
	return
}