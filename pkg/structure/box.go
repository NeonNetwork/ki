package structure

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Box[T Number] struct {
	position Vector2[T]
	size     Vector2[T]
}

func NewBox[T Number](position Vector2[T], size Vector2[T]) Box[T] {
	return Box[T]{
		position: position,
		size:     size,
	}
}

func (box Box[T]) ToRaylibRectangle() rl.Rectangle {
	return rl.NewRectangle(
		float32(box.X()),
		float32(box.Y()),
		float32(box.W()),
		float32(box.H()))
}

func (box Box[T]) X() T {
	return box.position.X()
}

func (box Box[T]) Y() T {
	return box.position.Y()
}

func (box Box[T]) W() T {
	return box.size.X()
}

func (box Box[T]) H() T {
	return box.size.Y()
}

func (box Box[T]) Position() Vector2[T] {
	return box.position
}

func (box Box[T]) Size() Vector2[T] {
	return box.size
}

func (box Box[T]) Center() Vector2[T] {
	return box.Position().
		Add(box.Size().Div(NewVector2[T](2, 2)))
}

func (box Box[T]) CollisionPoint(point Vector2[T]) (result bool) {
	if point.X() < box.X() {
		result = false
		return
	}

	if point.X() >= (box.X() + box.W()) {
		result = false
		return
	}

	if point.Y() < box.Y() {
		result = false
		return
	}

	if point.Y() >= (box.Y() + box.H()) {
		result = false
		return
	}

	result = true
	return
}
