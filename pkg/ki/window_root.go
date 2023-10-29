package ki

import (
	"github.com/google/uuid"
	"github.com/neonnetwork/ki/pkg/structure"
)

type RootWindow struct {
	id       uuid.UUID
	position structure.Vector2[int32]
	size     structure.Vector2[int32]
}

func (window *RootWindow) Id() uuid.UUID {
	return window.id
}

func (window *RootWindow) Position() structure.Vector2[int32] {
	return window.position
}

func (window *RootWindow) Size() structure.Vector2[int32] {
	return window.size
}

func (window *RootWindow) Box() structure.Box[int32] {
	return structure.NewBox[int32](
		window.Position(),
		window.Size())
}

func (window *RootWindow) SetPosition(value structure.Vector2[int32]) {
	window.position = value
}

func (window *RootWindow) SetSize(value structure.Vector2[int32]) {
	window.size = value
}

func (window *RootWindow) Init() *RootWindow {
	window.id = uuid.New()

	return window
}

func (window *RootWindow) Render(cursor structure.Vector2[int32]) (err error) {
	return
}

func (window *RootWindow) Split() (result structure.Pair[structure.Vector2[int32], structure.Vector2[int32]]) {
	size := window.size.Copy()
	cut := size.Div(structure.NewVector2[int32](2, 1))

	result = structure.NewPair(
		window.Position().Add(cut.Mul(structure.NewVector2[int32](1, 0))),
		cut)

	return
}
