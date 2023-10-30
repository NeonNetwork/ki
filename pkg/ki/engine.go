package ki

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/heartbytenet/bblib/objects"
	"github.com/neonnetwork/ki/pkg/structure"
)

const (
	EngineWindowUnit = 65536
)

type Engine struct {
	windows  *structure.BinaryTreeNode[Window]
	selected *structure.BinaryTreeNode[Window]
}

func (engine *Engine) Init() *Engine {
	engine.windows = structure.NewBinaryTreeNode[Window](objects.Init[RootWindow](&RootWindow{
		position: structure.NewVector2[int32](0, 0),
		size:     structure.NewVector2[int32](EngineWindowUnit, EngineWindowUnit),
	}))

	engine.selected = engine.windows

	return engine
}

func (engine *Engine) Start() (err error) {
	rl.SetConfigFlags(rl.FlagWindowResizable)

	rl.SetTargetFPS(360)

	rl.InitWindow(1280, 720, "Ki [raylib]")
	rl.HideCursor()

	for engine.Running() {
		err = engine.HandleInputs()
		if err != nil {
			return
		}

		err = engine.Render()
		if err != nil {
			return
		}
	}

	return
}

func (engine *Engine) Close() (err error) {
	rl.CloseWindow()

	return
}

func (engine *Engine) Wait() {
	return
}

func (engine *Engine) Running() bool {
	return !rl.WindowShouldClose()
}

func (engine *Engine) Screen() structure.Vector2[int32] {
	return structure.NewVector2[int32](
		int32(rl.GetRenderWidth()),
		int32(rl.GetRenderHeight()))
}

func (engine *Engine) Cursor() structure.Vector2[int32] {
	return structure.NewVector2[int32](
		rl.GetMouseX(),
		rl.GetMouseY())
}

func (engine *Engine) HandleInputs() (err error) {
	if rl.IsKeyPressed(rl.KeyF) {
		rl.ToggleFullscreen()
	}

	if rl.IsKeyPressed(rl.KeyH) {
		engine.selected = engine.WindowChildAdd(engine.selected, objects.Init[WindowImage](&WindowImage{}), structure.BinaryTreeLeft)
	}

	if rl.IsKeyPressed(rl.KeyL) {
		engine.selected = engine.WindowChildAdd(engine.selected, objects.Init[WindowImage](&WindowImage{}), structure.BinaryTreeRight)
	}

	return
}

func (engine *Engine) WindowChildAdd(window *structure.BinaryTreeNode[Window], child Window, direction structure.BinaryTreeDirection) (result *structure.BinaryTreeNode[Window]) {
	data := window.Value().Split(direction)

	child.SetPosition(data.A())
	child.SetSize(data.B())

	result = window.ChildAdd(child, direction)
	return
}

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
	err = engine.RenderWindow([]int{}, engine.windows)
	if err != nil {
		return
	}

	return
}

func (engine *Engine) RenderWindow(path []int, window *structure.BinaryTreeNode[Window]) (err error) {
	window.Value().SetSelected(window.Value().Id() == engine.selected.Value().Id())

	err = window.Value().Render(engine.Screen(), engine.Cursor())
	if err != nil {
		return
	}

	window.
		Left().
		IfPresent(func(value *structure.BinaryTreeNode[Window]) {
			err = engine.RenderWindow(append(path, -1), value)
			if err != nil {
				return
			}
		})
	if err != nil {
		return
	}

	window.
		Right().
		IfPresent(func(value *structure.BinaryTreeNode[Window]) {
			err = engine.RenderWindow(append(path, +1), value)
			if err != nil {
				return
			}
		})
	if err != nil {
		return
	}

	return
}
