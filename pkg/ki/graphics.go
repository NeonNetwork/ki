package ki

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/heartbytenet/bblib/containers/sync"
	"github.com/neonnetwork/ki/pkg/structure"
	"log"
)

var (
	GRAPHICS *sync.Mutex[*Graphics] = nil
)

type Graphics struct {
	engine *Engine

	font rl.Font
}

func (graphics *Graphics) Font() rl.Font {
	return graphics.font
}

func (graphics *Graphics) Init() *Graphics {
	if graphics.engine == nil {
		log.Fatalln("engine is nil")
	}

	GRAPHICS = sync.NewMutex(graphics)

	return graphics
}

func (graphics *Graphics) Start() (err error) {
	graphics.font = rl.LoadFontEx(
		"./data/font/iosevka-regular.ttf",
		1024,
		nil)

	return
}

func (graphics *Graphics) DrawTextCentered(text string, box structure.Box[int32]) (err error) {
	textData := text
	fontSize := float32(box.Size().Min()) / float32(len(textData)/2.0)
	textSize := structure.NewVector2FromRaylib[int32](
		rl.MeasureTextEx(
			graphics.Font(),
			textData,
			fontSize,
			0.0))

	rl.DrawTextEx(
		graphics.Font(),
		textData,
		box.Position().
			Add(box.Size().
				Div(structure.NewVector2[int32](2, 2)).
				Sub(textSize.Div(structure.NewVector2[int32](2, 2)))).
			ToRaylib(),
		fontSize,
		0.0,
		rl.RayWhite)

	return
}
