package structure

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/rand"
)

type Vector2[T Number] struct {
	x T
	y T
}

func NewVector2[T Number](x T, y T) Vector2[T] {
	return Vector2[T]{
		x: x,
		y: y,
	}
}

func MapVector2[T Number, U Number](vector Vector2[T], fn func(T) U) Vector2[U] {
	return NewVector2[U](
		fn(vector.X()),
		fn(vector.Y()))
}

func NewVector2Random[T Number](n int) Vector2[T] {
	return MapVector2[byte, T](
		NewVector2[byte](0, 0),
		func(v byte) T {
			return T(rand.Intn(n))
		})
}

func (vector Vector2[T]) X() T {
	return vector.x
}

func (vector Vector2[T]) Y() T {
	return vector.y
}

func (vector Vector2[T]) ToRaylib() rl.Vector2 {
	return rl.NewVector2(
		float32(vector.x),
		float32(vector.y))
}

func (vector Vector2[T]) Copy() Vector2[T] {
	return NewVector2(
		vector.x,
		vector.y)
}

func (vector Vector2[T]) Add(value Vector2[T]) Vector2[T] {
	return NewVector2[T](
		vector.x+value.x,
		vector.y+value.y)
}

func (vector Vector2[T]) Sub(value Vector2[T]) Vector2[T] {
	return NewVector2[T](
		vector.x-value.x,
		vector.y-value.y)
}

func (vector Vector2[T]) Mul(value Vector2[T]) Vector2[T] {
	return NewVector2[T](
		vector.x*value.x,
		vector.y*value.y)
}

func (vector Vector2[T]) Div(value Vector2[T]) Vector2[T] {
	return NewVector2[T](
		Div(vector.x, value.x),
		Div(vector.y, value.y))
}

func (vector Vector2[T]) Abs() Vector2[T] {
	return NewVector2[T](
		Abs(vector.x),
		Abs(vector.y))
}

func (vector Vector2[T]) Inv() Vector2[T] {
	return NewVector2[T](
		Inv(vector.x),
		Inv(vector.y))
}
