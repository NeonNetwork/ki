package structure

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"

	"golang.org/x/exp/rand"
)

type Vector3[T Number] struct {
	x T
	y T
	z T
}

func NewVector3[T Number, V Number](x V, y V, z V) Vector3[T] {
	return Vector3[T]{
		x: T(x),
		y: T(y),
		z: T(z),
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

func (vector Vector3[T]) ToRaylib() rl.Vector3 {
	return rl.NewVector3(
		float32(vector.X()),
		float32(vector.Y()),
		float32(vector.Z()))
}

func (vector Vector3[T]) ToColor() color.RGBA {
	return color.RGBA{
		R: byte(vector.X()),
		G: byte(vector.Y()),
		B: byte(vector.Z()),
		A: 255,
	}
}

func (vector Vector3[T]) Copy() Vector3[T] {
	return NewVector3[T, T](
		vector.X(),
		vector.Y(),
		vector.Z())
}

func (vector Vector3[T]) Add(value Vector3[T]) Vector3[T] {
	return NewVector3[T](
		vector.X() + value.X(),
		vector.Y() + value.Y(),
		vector.Z() + value.Z())
}

func (vector Vector3[T]) Sub(value Vector3[T]) Vector3[T] {
	return NewVector3[T](
		vector.X() - value.X(),
		vector.Y() - value.Y(),
		vector.Z() - value.Z())
}

func (vector Vector3[T]) Mul(value Vector3[T]) Vector3[T] {
	return NewVector3[T](
		vector.X() * value.X(),
		vector.Y() * value.Y(),
		vector.Z() * value.Z())
}

func (vector Vector3[T]) Div(value Vector3[T]) Vector3[T] {
	return NewVector3[T](
		Div(vector.X(), value.X()),
		Div(vector.Y(), value.Y()),
		Div(vector.Z(), value.Z()))
}

func (vector Vector3[T]) ToUint8() Vector3[uint8] {
	return NewVector3[uint8](
		vector.X(),
		vector.Y(),
		vector.Z())
}

func (vector Vector3[T]) ToInt32() Vector3[int32] {
	return NewVector3[int32](
		vector.X(),
		vector.Y(),
		vector.Z())
}

func (vector Vector3[T]) ToFloat32() Vector3[float32] {
	return NewVector3[float32](
		vector.X(),
		vector.Y(),
		vector.Z())
}

func (vector Vector3[T]) ToFloat64() Vector3[float64] {
	return NewVector3[float64](
		vector.X(),
		vector.Y(),
		vector.Z())
}

