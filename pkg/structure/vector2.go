package structure

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/rand"
)

type Vector2[T Number] struct {
	x T
	y T
}

func NewVector2[T Number, V Number](x V, y V) Vector2[T] {
	return Vector2[T]{
		x: T(x),
		y: T(y),
	}
}

func NewVector2FromRaylib[T Number](value rl.Vector2) Vector2[T] {
	return NewVector2[T](
		value.X,
		value.Y)
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
	return NewVector2[T, T](
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

func (vector Vector2[T]) Rev() Vector2[T] {
	return NewVector2[T](
		vector.Y(),
		vector.X())
}

func (vector Vector2[T]) Order(reversed bool) Vector2[T] {
	if reversed {
		return vector.Rev()
	} else {
		return vector.Copy()
	}
}

func (vector Vector2[T]) Max() T {
	if vector.X() > vector.Y() {
		return vector.X()
	} else {
		return vector.Y()
	}
}

func (vector Vector2[T]) Min() T {
	if vector.X() < vector.Y() {
		return vector.X()
	} else {
		return vector.Y()
	}
}

func (vector Vector2[T]) ToInt32() Vector2[int32] {
	return NewVector2[int32](
		vector.X(),
		vector.Y())
}

func (vector Vector2[T]) ToFloat32() Vector2[float32] {
	return NewVector2[float32](
		vector.X(),
		vector.Y())
}

func (vector Vector2[T]) ToFloat64() Vector2[float64] {
	return NewVector2[float64](
		vector.X(),
		vector.Y())
}
