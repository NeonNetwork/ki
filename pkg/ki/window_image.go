package ki

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
	"github.com/neonnetwork/ki/pkg/structure"
	"os"
)

type WindowImage struct {
	id       uuid.UUID
	position structure.Vector2[int32]
	size     structure.Vector2[int32]
	selected bool
	color    structure.Vector3[uint8]

	texture rl.Texture2D
}

func (window *WindowImage) Id() uuid.UUID {
	return window.id
}

func (window *WindowImage) Position() structure.Vector2[int32] {
	return window.position
}

func (window *WindowImage) Size() structure.Vector2[int32] {
	return window.size
}

func (window *WindowImage) Box() structure.Box[int32] {
	return structure.NewBox[int32](
		window.Position(),
		window.Size())
}

func (window *WindowImage) Color() structure.Vector3[uint8] {
	return window.color
}

func (window *WindowImage) SetPosition(value structure.Vector2[int32]) {
	window.position = value
}

func (window *WindowImage) SetSize(value structure.Vector2[int32]) {
	window.size = value
}

func (window *WindowImage) Selected() bool {
	return window.selected
}

func (window *WindowImage) SetSelected(value bool) {
	window.selected = value
}

func (window *WindowImage) Init() *WindowImage {
	window.id = uuid.New()

	window.color = structure.NewVector3Random[byte](256)

	window.texture = rl.LoadTexture(os.Getenv("TEXTURE"))

	return window
}

func (window *WindowImage) Render(screen structure.Vector2[int32], cursor structure.Vector2[int32]) (err error) {
	position := structure.MapVector2[float64, int32](
		structure.MapVector2[int32, float64](window.Position(), structure.ConvertNumberInt32toFloat64).
			Div(structure.NewVector2[float64](float64(EngineWindowUnit), float64(EngineWindowUnit))).
			Mul(structure.MapVector2[int32, float64](screen, structure.ConvertNumberInt32toFloat64)),
			structure.ConvertNumberFloat64toInt32)

	size := structure.MapVector2[float64, int32](
		structure.MapVector2[int32, float64](window.Size(), structure.ConvertNumberInt32toFloat64).
			Div(structure.NewVector2[float64](float64(EngineWindowUnit), float64(EngineWindowUnit))).
			Mul(structure.MapVector2[int32, float64](screen, structure.ConvertNumberInt32toFloat64)),
			structure.ConvertNumberFloat64toInt32)

	c := window.Color().ToColor()
	if structure.NewBox(position, size).CollisionPoint(cursor) {
		c = rl.Red
	}

//	rl.DrawRectangle(
//		position.X(),
//		position.Y(),
//		size.X(),
//		size.Y(),
//		c)

	rl.DrawTextureRec(
		window.texture,
		rl.NewRectangle(0, 0, float32(size.X()), float32(size.Y())),
		position.ToRaylib(),
		c)

	return
}

func (window *WindowImage) Split() (result structure.Pair[structure.Vector2[int32], structure.Vector2[int32]]) {
	size := window.size.Copy()
	cut := size.Div(structure.NewVector2[int32](2, 1))

	window.size = cut

	result = structure.NewPair(
		window.Position().Add(cut.Mul(structure.NewVector2[int32](1, 0))),
		cut)

	return
}
