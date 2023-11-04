package ki

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/heartbytenet/bblib/containers/optionals"

	"github.com/neonnetwork/ki/pkg/structure"
)

const (
	EngineWindowUnit = 65536
)

type Engine struct {
	windows  *structure.BinaryTreeNode[Window]
	selected *structure.BinaryTreeNode[Window]

	windowsFloating []Window
}

func (engine *Engine) WindowRoot() optionals.Optional[Window] {
	return optionals.FlatMap[*structure.BinaryTreeNode[Window], Window](
		engine.WindowRootNode(),
		func(value *structure.BinaryTreeNode[Window]) optionals.Optional[Window] {
			return optionals.Some(value.Value())
		})
}

func (engine *Engine) WindowRootNode() optionals.Optional[*structure.BinaryTreeNode[Window]] {
	if engine.windows == nil {
		return optionals.None[*structure.BinaryTreeNode[Window]]()
	}

	return optionals.Some(engine.windows)
}

func (engine *Engine) WindowSelected() optionals.Optional[Window] {
	return optionals.FlatMap[*structure.BinaryTreeNode[Window], Window](
		engine.WindowSelectedNode(),
		func(value *structure.BinaryTreeNode[Window]) optionals.Optional[Window] {
			return optionals.Some(value.Value())
		})
}

func (engine *Engine) WindowSelectedNode() optionals.Optional[*structure.BinaryTreeNode[Window]] {
	if engine.selected == nil {
		return optionals.None[*structure.BinaryTreeNode[Window]]()
	}

	return optionals.Some(engine.selected)
}

func (engine *Engine) Init() *Engine {
	engine.windows = nil
	engine.selected = engine.windows

	engine.windowsFloating = make([]Window, 0)

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

		err = engine.Compute()
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
