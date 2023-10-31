package ki

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (engine *Engine) Render() (err error) {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Black)

	err = engine.RenderWindows()
	if err != nil {
		return
	}

	rl.DrawRectangle(rl.GetMouseX()-8, rl.GetMouseY()-8, 16, 16, rl.White)
	rl.DrawFPS(8, 8)

	return
}

func (engine *Engine) RenderWindows() (err error) {
	err = engine.IterateWindows(func(window Window, data any) (any, error) {
		err := window.Render()
		if err != nil {
			return data, err
		}

		return data, nil
	}, nil)
	if err != nil {
		return
	}

	return
}
