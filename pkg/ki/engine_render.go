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
	return
}
