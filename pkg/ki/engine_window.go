package ki

import (
	"fmt"
	"strings"

	"github.com/neonnetwork/ki/pkg/structure"
)

func (engine *Engine) WindowSplitCommon(window Window, direction structure.BinaryTreeDirection) (result structure.Box[int32]) {
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
		result = structure.NewBox(
			window.Position().Add(zero),
			size)

		window.SetSize(size)
		window.SetPosition(window.Position().Add(diff))
	} else {
		result = structure.NewBox(
			window.Position().Add(diff),
			size)

		window.SetSize(size)
		window.SetPosition(window.Position().Add(zero))
	}

	return
}

func (engine *Engine) WindowChildAdd(
	node      *structure.BinaryTreeNode[Window],
	value     Window,
	axis      WindowSplitAxis,
	direction structure.BinaryTreeDirection,
) (result *structure.BinaryTreeNode[Window]) {
	var (
		window Window
		size   structure.Box[int32]
	)

	if node == nil {
		size = structure.NewBox(
			structure.NewVector2[int32](0, 0),
			structure.NewVector2[int32](EngineWindowUnit, EngineWindowUnit))
	} else {
		window = node.Value()

		size = engine.WindowSplitCommon(window, direction)
	}

	value.SetSide(direction)
	value.SetPosition(size.Position())
	value.SetSize(size.Size())

	if node == nil {
		result = structure.NewBinaryTreeNode(value)

		engine.windows = result
	} else {
		result = node.ChildAdd(value, direction)
	}

	value.SetNode(result)

	return
}

func (engine *Engine) WindowTreePrint() {
	engine.WindowRootNode().IfPresent(func(value *structure.BinaryTreeNode[Window]) {
		engine.WindowTreePrintStep(value, 0, 32)
	})
}

func (engine *Engine) WindowTreePrintStep(node *structure.BinaryTreeNode[Window], level int, step int) {
	window := node.Value()


	fmt.Println("")
	fmt.Print(strings.Repeat(" ", level))
	if window.Selected() {
		fmt.Print(">")
	} else {
		fmt.Print(" ")
	}
	fmt.Printf(
		"%s_%v[%v-%v]",
		window.Id().String()[:4],
		window.IsRoot(),
		window.PositionAbsolute(),
		window.SizeAbsolute())


	node.Left().IfPresent(func(value *structure.BinaryTreeNode[Window]) {
		engine.WindowTreePrintStep(value, level + step, step)
	})

	node.Right().IfPresent(func(value *structure.BinaryTreeNode[Window]) {
		engine.WindowTreePrintStep(value, level + step, step)
	})
}