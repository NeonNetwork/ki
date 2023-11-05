package ki

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/neonnetwork/ki/pkg/structure"
)

func (engine *Engine) Render() (err error) {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Black)

	err = engine.RenderWindows()
	if err != nil {
		return
	}

	err = engine.RenderWindowsSelected()
	if err != nil {
		return
	}

	err = engine.RenderCursor()
	if err != nil {
		return
	}

	rl.DrawFPS(8, 8)

	return
}

func (engine *Engine) RenderCursor() (err error) {
	rl.DrawRectangleV(
		engine.Cursor().Sub(structure.NewVector2[int32](8, 8)).ToRaylib(),
		structure.NewVector2[int32](16, 16).ToRaylib(),
		rl.RayWhite)

	return
}

func (engine *Engine) RenderWindows() (err error) {
	engine.WindowRootNode().IfPresent(func(value *structure.BinaryTreeNode[Window]) {
		err = engine.RenderWindowStep(value)
	})
	if err != nil {
		return
	}

	return
}

func (engine *Engine) RenderWindowStep(node *structure.BinaryTreeNode[Window]) (err error) {
	var (
		window Window
	)

	window = node.Value()

	err = window.Render()
	if err != nil {
		return
	}

	node.Left().IfPresent(func(value *structure.BinaryTreeNode[Window]) {
		err = engine.RenderWindowStep(value)
	})
	if err != nil {
		return
	}

	node.Right().IfPresent(func(value *structure.BinaryTreeNode[Window]) {
		err = engine.RenderWindowStep(value)
	})
	if err != nil {
		return
	}

	return
}

func (engine *Engine) RenderWindowsSelected() (err error) {
	err = engine.WindowListWalkNode(func(value *structure.BinaryTreeNode[Window]) error {
		window := value.Value()

		if window.Selected() {
			rl.DrawRectangleLinesEx(
				window.BoxAbs().ToRaylibRectangle(),
				2,
				rl.RayWhite)
		}

		return nil
	})
	if err != nil {
		return
	}

	return
}
