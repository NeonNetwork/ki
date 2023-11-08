package ki

import (
	"fmt"
	"log"
	"os"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/heartbytenet/bblib/containers/optionals"
	"github.com/heartbytenet/bblib/containers/sync"
	"github.com/heartbytenet/bblib/objects"
	"github.com/heartbytenet/go-lerpc/pkg/lerpc"

	"github.com/neonnetwork/ki/pkg/structure"
)

const (
	EngineWindowUnit = 1048576
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var (
	ENGINE             *sync.Mutex[*Engine] = nil
	RpcExecuteModeHttp                      = lerpc.ClientModeHttpOnly
	RpcSecure          uint32               = 0
)

type Engine struct {
	client    *Client
	rpcClient *lerpc.Client
	pool      *Pool
	graphics  *Graphics
	logic     *Logic

	windows  *structure.BinaryTreeNode[Window]
	selected *structure.BinaryTreeNode[Window]

	windowsFloating []Window
}

func (engine *Engine) Client() *Client {
	return engine.client
}

func (engine *Engine) RpcClient() *lerpc.Client {
	return engine.rpcClient
}

func (engine *Engine) Pool() *Pool {
	return engine.pool
}

func (engine *Engine) Graphics() *Graphics {
	return engine.graphics
}

func (engine *Engine) Logic() *Logic {
	return engine.logic
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
	engine.client = objects.Init[Client](&Client{})
	engine.pool = objects.Init[Pool](&Pool{engine: engine})
	engine.graphics = objects.Init[Graphics](&Graphics{engine: engine})
	engine.logic = objects.Init[Logic](&Logic{engine: engine})

	engine.rpcClient = (&lerpc.Client{}).Init("localhost:12000", "")

	engine.windows = nil
	engine.selected = engine.windows

	engine.windowsFloating = make([]Window, 0)

	ENGINE = sync.NewMutex(engine)

	return engine
}

func (engine *Engine) Start() (err error) {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.SetConfigFlags(rl.FlagWindowAlwaysRun)
	rl.SetConfigFlags(rl.FlagWindowHighdpi)

	if os.Getenv("VSYNC") != "" {
		rl.SetConfigFlags(rl.FlagVsyncHint)
	}

	rl.SetTargetFPS(360)

	rl.InitWindow(1280, 720, "Ki [raylib]")
	rl.HideCursor()

	engine.rpcClient.Mode(&RpcExecuteModeHttp)
	engine.rpcClient.Secure(&RpcSecure)
	err = engine.rpcClient.Start(0)
	if err != nil {
		return
	}

	err = engine.graphics.Start()
	if err != nil {
		return
	}

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

func (engine *Engine) Screenshot() (err error) {
	err = os.MkdirAll("./data/screenshots/", 0755)
	if err != nil {
		return
	}

	rl.TakeScreenshot(fmt.Sprintf("./data/screenshots/%v.png", time.Now().UnixMilli()))

	return
}