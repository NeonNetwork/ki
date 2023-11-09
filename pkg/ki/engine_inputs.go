package ki

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/heartbytenet/bblib/objects"
	"github.com/neonnetwork/ki/pkg/structure"
)

func (engine *Engine) HandleInputs() (err error) {
	if rl.IsKeyPressed(rl.KeyF) {
		rl.ToggleFullscreen()
	}

	if rl.IsKeyPressed(rl.KeyMinus) {
		defer engine.WindowTreePrint()
	}

	if rl.IsKeyPressed(rl.KeyZ) {
		err = engine.Screenshot()
		if err != nil {
			return
		}
	}

	if rl.IsKeyPressed(rl.KeyQ) {
		engine.fpsDisplay = !engine.fpsDisplay
	}

	if rl.IsKeyDown(rl.KeyLeftShift) {
		if rl.IsKeyPressed(rl.KeyJ) {
			engine.WindowSelectedNode().
				IfPresent(func(node *structure.BinaryTreeNode[Window]) {
					node.
						Prev().
						IfPresent(func(value *structure.BinaryTreeNode[Window]) {
							engine.selected = value
						})
				})
		}

		if rl.IsKeyPressed(rl.KeyH) {
			engine.WindowSelectedNode().
				IfPresent(func(node *structure.BinaryTreeNode[Window]) {
					node.
						Left().
						IfPresent(func(value *structure.BinaryTreeNode[Window]) {
							engine.selected = value
						})
				})
		}

		if rl.IsKeyPressed(rl.KeyL) {
			engine.WindowSelectedNode().
				IfPresent(func(node *structure.BinaryTreeNode[Window]) {
					node.
						Right().
						IfPresent(func(value *structure.BinaryTreeNode[Window]) {
							engine.selected = value
						})
				})
		}
	} else {
		if rl.IsKeyPressed(rl.KeyH) {
			engine.selected = engine.WindowChildAdd(
				engine.selected,
				objects.Init[WindowImage](&WindowImage{}),
				WindowSplitHorizontal,
				structure.BinaryTreeLeft)
		}

		if rl.IsKeyPressed(rl.KeyJ) {
			engine.selected = engine.WindowChildAdd(
				engine.selected,
				objects.Init[WindowImage](&WindowImage{}),
				WindowSplitVertical,
				structure.BinaryTreeRight)
		}

		if rl.IsKeyPressed(rl.KeyK) {
			engine.selected = engine.WindowChildAdd(
				engine.selected,
				objects.Init[WindowImage](&WindowImage{}),
				WindowSplitVertical,
				structure.BinaryTreeLeft)
		}

		if rl.IsKeyPressed(rl.KeyL) {
			engine.selected = engine.WindowChildAdd(
				engine.selected,
				objects.Init[WindowImage](&WindowImage{}),
				WindowSplitHorizontal,
				structure.BinaryTreeRight)
		}
	}

	return
}
