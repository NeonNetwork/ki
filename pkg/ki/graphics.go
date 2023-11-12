package ki

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/heartbytenet/bblib/containers/sync"
	"github.com/neonnetwork/ki/pkg/structure"
	"image/color"
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

	rl.SetTextureFilter(graphics.font.Texture, rl.FilterBilinear)

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

func (graphics *Graphics) DrawText(value string, pos structure.Vector2[int32], fontSize float64) (err error) {
	rl.DrawTextEx(
		graphics.Font(),
		value,
		pos.ToRaylib(),
		float32(fontSize),
		0.0,
		rl.RayWhite)

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

func (graphics *Graphics) DrawLine(a, b structure.Vector2[int32], thick float64, fill color.RGBA) (err error) {
	rl.DrawLineEx(
		a.ToRaylib(),
		b.ToRaylib(),
		float32(thick),
		fill)

	return
}

func (graphics *Graphics) DrawGraph(
	value []float64,
	valueMin float64,
	valueMax float64,
	box structure.Box[int32],
) (err error) {
	points := make([]structure.Vector2[int32], 0)

	for i, v := range value {
		pointX := float64(i) / float64(len(value)) * box.Size().ToFloat64().X()
		pointY := (1.0 - (v-valueMin)/(valueMax-valueMin)) * box.Size().ToFloat64().Y()

		points = append(points, structure.NewVector2[int32](pointX, pointY))
	}

	for _, point := range points {
		err = graphics.DrawLine(
			box.Position().Add(point.Mul(structure.NewVector2[int32](1, 0))).Add(box.Size().Mul(structure.NewVector2[int32](0, 1))),
			box.Position().Add(point.Mul(structure.NewVector2[int32](1, 1))),
			2.0,
			rl.RayWhite)
		if err != nil {
			return
		}
	}

	return
}

func (graphics *Graphics) DrawCircleSector(
	pos structure.Vector2[int32],
	radius float64,
	angles structure.Pair[float64, float64],
	color structure.Vector3[uint8],
) (err error) {
	rl.DrawCircleSector(
		pos.ToRaylib(),
		float32(radius),
		float32(angles.A()),
		float32(angles.B()),
		32,
		color.ToColor())

	return
}

func (graphics *Graphics) DrawRectangle(box structure.Box[int32], color color.RGBA) (err error) {
	rl.DrawRectangle(
		box.X(),
		box.Y(),
		box.W(),
		box.H(),
		color)

	return
}
