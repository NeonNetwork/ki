package ki

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
	"github.com/neonnetwork/ki/pkg/structure"
	"image/color"
)

type WindowImage struct {
	id       uuid.UUID
	position structure.Vector2[int32]
	size     structure.Vector2[int32]
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

func (window *WindowImage) Color() color.RGBA {
	v := byte(window.position.X())

	return color.RGBA{v, v, v, 255}
}

func (window *WindowImage) SetPosition(value structure.Vector2[int32]) {
	window.position = value
}

func (window *WindowImage) SetSize(value structure.Vector2[int32]) {
	window.size = value
}

func (window *WindowImage) Init() *WindowImage {
	window.id = uuid.New()

	return window
}

func (window *WindowImage) Render(cursor structure.Vector2[int32]) (err error) {
	c := window.Color()

	if window.Box().CollisionPoint(cursor) {
		c = rl.Red
	}

	rl.DrawRectangle(
		window.Position().X(),
		window.Position().Y(),
		window.Size().X(),
		window.Size().Y(),
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
