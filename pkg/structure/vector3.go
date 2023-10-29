package structure

import (
	"golang.org/x/exp/rand"
	"image/color"
)

type Vector3[T Number] struct {
	x T
	y T
	z T
}

func NewVector3[T Number](x T, y T, z T) Vector3[T] {
	return Vector3[T]{
		x: x,
		y: y,
		z: z,
	}
}

func MapVector3[T Number, U Number](vector Vector3[T], fn func(T) U) Vector3[U] {
	return NewVector3[U](
		fn(vector.X()),
		fn(vector.Y()),
		fn(vector.Z()))
}

func NewVector3Random[T Number](n int) Vector3[T] {
	return MapVector3[byte, T](
		NewVector3[byte](0, 0, 0),
		func(v byte) T {
			return T(rand.Intn(n))
		})
}

func (vector Vector3[T]) X() T {
	return vector.x
}

func (vector Vector3[T]) Y() T {
	return vector.y
}

func (vector Vector3[T]) Z() T {
	return vector.z
}

func (vector Vector3[T]) ToColor() color.RGBA {
	return color.RGBA{
		R: byte(vector.X()),
		G: byte(vector.Y()),
		B: byte(vector.Z()),
		A: 255,
	}
}
