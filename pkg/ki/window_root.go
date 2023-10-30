package ki

import (
	"github.com/google/uuid"
	"github.com/neonnetwork/ki/pkg/structure"
)

type RootWindow struct {
	id       uuid.UUID
	position structure.Vector2[int32]
	size     structure.Vector2[int32]
	axis     WindowSplitAxis
	selected bool
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

func (window *RootWindow) Selected() bool {
	return window.selected
}

func (window *RootWindow) SetSelected(value bool) {
	window.selected = value
}

func (window *RootWindow) Init() *RootWindow {
	window.id = uuid.New()

	return window
}

func (window *RootWindow) Render(screen structure.Vector2[int32], cursor structure.Vector2[int32]) (err error) {
	return
}

func (window *RootWindow) Split(axis WindowSplitAxis, direction structure.BinaryTreeDirection) (result structure.Pair[structure.Vector2[int32], structure.Vector2[int32]]) {
	window.SetSplitAxis(axis)

	return WindowSplitCommon(window, direction)
}

func (window *RootWindow) SplitAxis() WindowSplitAxis {
	return window.axis
}

func (window *RootWindow) SetSplitAxis(value WindowSplitAxis) {
	window.axis = value
}
