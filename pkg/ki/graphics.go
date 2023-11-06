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

	textureCursor rl.Texture2D
	textureWindow rl.Texture2D
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

	graphics.textureCursor = rl.LoadTexture("./data/image/texture_cursor.png")
	graphics.textureWindow = rl.LoadTexture("./data/image/texture_window.png")

	return
}

func (graphics *Graphics) DrawTexture(value rl.Texture2D, position structure.Vector2[int32]) (err error) {
	rl.DrawTexture(
		value,
		position.X(),
		position.Y(),
		rl.White)

	return
}

func (graphics *Graphics) DrawTextureVector(
	value rl.Texture2D,
	position structure.Vector2[int32],
	size structure.Vector2[int32],
) (err error) {
	rl.DrawTextureRec(
		value,
		structure.NewBox[float32](
			structure.NewVector2[float32](0, 0),
			size.ToFloat32(),
		).ToRaylibRectangle(),
		position.ToRaylib(),
		rl.White)

	return
}

func (graphics *Graphics) DrawTextCentered(value string, box structure.Box[int32]) (err error) {
	textData := value
	fontSize := float32(box.Size().Min()) / float32(len(textData)) * 1.5
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

func (graphics *Graphics) DrawGraph(value []float64, box structure.Box[int32]) (err error) {
	points := make([]structure.Vector2[int32], 0)

	valueMin, valueMax := structure.MinMax[float64](value...)

	for i, v := range value {
		pointX := float64(i) / float64(len(value)) * box.Size().ToFloat64().X()
		pointY := (v - valueMin) / (valueMax - valueMin) * box.Size().ToFloat64().Y()

		points = append(points, structure.NewVector2[int32](pointX, pointY))
	}

	for _, point := range points {
		rl.DrawLineEx(
			box.Position().Add(point.Mul(structure.NewVector2[int32](1, 1))).Add(box.Size().Mul(structure.NewVector2[int32](0, 1))).ToRaylib(),
			box.Position().Add(point.Mul(structure.NewVector2[int32](1, 1))).ToRaylib(),
			2.0,
			rl.RayWhite)

		//		rl.DrawCircle(
		//			box.Position().X()+point.X(),
		//			box.Position().Y()+point.Y(),
		//			2.0,
		//			rl.RayWhite)
	}

	return
}
