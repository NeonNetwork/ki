package structure

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Inv[T Number](value T) T {
	return -value
}

func Abs[T Number](value T) T {
	if value < 0 {
		return Inv(value)
	}

	return value
}

func Div[T Number](x T, y T) T {
	if y == 0 {
		return T(0)
	}

	return x / y
}
