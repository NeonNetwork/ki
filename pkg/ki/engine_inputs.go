package ki

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/heartbytenet/bblib/objects"
	"github.com/neonnetwork/ki/pkg/structure"
)

func (engine *Engine) HandleInputs() (err error) {
	defer engine.WindowTreePrint()

	if rl.IsKeyPressed(rl.KeyF) {
		rl.ToggleFullscreen()
	}

	if rl.IsKeyDown(rl.KeyLeftShift) {
		if rl.IsKeyPressed(rl.KeyJ) {
			engine.selected.
				Prev().
				IfPresent(func(value *structure.BinaryTreeNode[Window]) {
					engine.selected = value
				})
		}

		if rl.IsKeyPressed(rl.KeyH) {
			engine.selected.
				Left().
				IfPresent(func(value *structure.BinaryTreeNode[Window]) {
					engine.selected = value
				})
		}

		if rl.IsKeyPressed(rl.KeyL) {
			engine.selected.
				Right().
				IfPresent(func(value *structure.BinaryTreeNode[Window]) {
					engine.selected = value
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
