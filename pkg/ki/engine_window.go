package ki

import (
	"fmt"
	"github.com/heartbytenet/bblib/objects"
	"strings"

	"github.com/neonnetwork/ki/pkg/structure"
)

// WindowSplitCommon handles the splitting of windows
// this returns the box values for the current and newly created window
func (engine *Engine) WindowSplitCommon(window Window, direction structure.BinaryTreeDirection) (result structure.Pair[structure.Box[int32], structure.Box[int32]]) {
	var (
		axis WindowSplitAxis
		vert bool
		zero structure.Vector2[int32]
		size structure.Vector2[int32]
		diff structure.Vector2[int32]
	)

	axis = window.SplitAxis()                // W          || H
	vert = axis == WindowSplitVertical       // true       || false
	zero = structure.NewVector2[int32](0, 0) // [0 , 0] || [0, 0]
	size = structure.NewVector2[int32](EngineWindowUnit, EngineWindowUnit)
	size = size.Div(structure.NewVector2[int32](2, 1).Order(vert)) // [W / 2, H] || [W, H / 2]
	diff = size.Mul(structure.NewVector2[int32](1, 0).Order(vert)) // [W / 2, 0] || [0, H / 2]

	if direction == structure.BinaryTreeLeft {
		result = structure.NewPair(
			structure.NewBox(zero.Add(diff), size),
			structure.NewBox(zero.Add(zero), size))
	} else {
		result = structure.NewPair(
			structure.NewBox(zero.Add(zero), size),
			structure.NewBox(zero.Add(diff), size))
	}

	return
}

func (engine *Engine) WindowChildAdd(
	node *structure.BinaryTreeNode[Window],
	value Window,
	axis WindowSplitAxis,
	direction structure.BinaryTreeDirection,
) (result *structure.BinaryTreeNode[Window]) {
	var (
		split  Window
		window Window
		data   structure.Pair[structure.Box[int32], structure.Box[int32]]
	)

	if node == nil {
		data = structure.NewPair(
			structure.NewBox(structure.NewVector2[int32](0, 0), structure.NewVector2[int32](0, 0)),
			structure.NewBox(structure.NewVector2[int32](0, 0), structure.NewVector2[int32](EngineWindowUnit, EngineWindowUnit)))
	} else {
		window = node.Value()

		split = objects.Init[WindowSplit](&WindowSplit{})
		split.SetSide(window.Side())
		split.SetSplitAxis(axis)
		split.SetPosition(window.Position())
		split.SetSize(window.Size())

		data = engine.WindowSplitCommon(split, direction)

		window.SetSide(direction.Opposite())
		window.SetPosition(data.A().Position())
		window.SetSize(data.A().Size())
	}

	value.SetSide(direction)
	value.SetPosition(data.B().Position())
	value.SetSize(data.B().Size())

	if node == nil {
		result = structure.NewBinaryTreeNode(value)

		engine.windows = result
	} else {
		node.SetValue(split)
		split.SetNode(node)

		window.SetNode(node.ChildAdd(window, window.Side()))

		result = node.ChildAdd(value, direction)
	}

	value.SetNode(result)

	return
}

func (engine *Engine) WindowTreePrint() {
	fmt.Println()
	engine.WindowRootNode().IfPresent(func(value *structure.BinaryTreeNode[Window]) {
		engine.WindowTreePrintStep(value, 0, 8)
	})
}

func (engine *Engine) WindowTreePrintStep(node *structure.BinaryTreeNode[Window], level int, step int) {
	window := node.Value()

	fmt.Print(strings.Repeat(" ", level))
	if window.Selected() {
		fmt.Print("+")
	} else {
		fmt.Print("-")
	}
	fmt.Printf(
		"%s->T=%v;R=%v;B=%v;A[%v]\n",
		window.Id().String()[:4],
		window.Type(),
		window.IsRoot(),
		window.Box(),
		window.BoxAbs())

	node.Left().IfPresent(func(value *structure.BinaryTreeNode[Window]) {
		engine.WindowTreePrintStep(value, level+step, step)
	})

	node.Right().IfPresent(func(value *structure.BinaryTreeNode[Window]) {
		engine.WindowTreePrintStep(value, level+step, step)
	})
}
