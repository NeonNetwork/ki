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
	axis     WindowSplitAxis
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

func (window *WindowImage) ScaledPosition(screen structure.Vector2[int32]) structure.Vector2[int32] {
	return structure.MapVector2[float64, int32](
		structure.MapVector2[int32, float64](window.Position(), structure.ConvertNumberInt32toFloat64).
			Div(structure.NewVector2[float64](float64(EngineWindowUnit), float64(EngineWindowUnit))).
			Mul(structure.MapVector2[int32, float64](screen, structure.ConvertNumberInt32toFloat64)),
		structure.ConvertNumberFloat64toInt32)
}

func (window *WindowImage) ScaledSize(screen structure.Vector2[int32]) structure.Vector2[int32] {
	return structure.MapVector2[float64, int32](
		structure.MapVector2[int32, float64](window.Size(), structure.ConvertNumberInt32toFloat64).
			Div(structure.NewVector2[float64](float64(EngineWindowUnit), float64(EngineWindowUnit))).
			Mul(structure.MapVector2[int32, float64](screen, structure.ConvertNumberInt32toFloat64)),
		structure.ConvertNumberFloat64toInt32)
}

func (window *WindowImage) ScaledBox(screen structure.Vector2[int32]) structure.Box[int32] {
	return structure.NewBox(window.ScaledPosition(screen), window.ScaledSize(screen))
}

func (window *WindowImage) Render(screen structure.Vector2[int32], cursor structure.Vector2[int32]) (err error) {
	position := window.ScaledPosition(screen)
	size := window.ScaledSize(screen)

	c := window.Color().ToColor()
	if window.ScaledBox(screen).CollisionPoint(cursor) {
		c = rl.Red
	}

	rl.DrawTextureRec(
		window.texture,
		rl.NewRectangle(0, 0, float32(size.X()), float32(size.Y())),
		position.ToRaylib(),
		c)

	if window.Selected() {
		rl.DrawRectangleLinesEx(
			window.ScaledBox(screen).ToRaylibRectangle(),
			2,
			rl.RayWhite)
	}

	return
}

func (window *WindowImage) Split(axis WindowSplitAxis, direction structure.BinaryTreeDirection) (result structure.Pair[structure.Vector2[int32], structure.Vector2[int32]]) {
	window.SetSplitAxis(axis)

	return WindowSplitCommon(window, direction)
}

func (window *WindowImage) SplitAxis() WindowSplitAxis {
	return window.axis
}

func (window *WindowImage) SetSplitAxis(value WindowSplitAxis) {
	window.axis = value
}
