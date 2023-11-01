package ki

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/neonnetwork/ki/pkg/structure"
	"log"
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

type IterateDataRender struct {
	Position structure.Vector2[int32]
	Last     Window
}

func (data IterateDataRender) SetLast(value Window) IterateDataRender {
	data.Last = value

	return data
}

func (engine *Engine) RenderWindows() (err error) {
	err = engine.IterateWindows(
		func(window Window, data any) (any, error) {
			value := data.(IterateDataRender)

			if window.IsRoot() {
				err = window.Render()
				if err != nil {
					return nil, err
				}
			}

			return value.SetLast(window), nil
		},
		IterateDataRender{
			Position: structure.NewVector2[int32](0, 0),
		})
	if err != nil {
		return
	}

	return
}