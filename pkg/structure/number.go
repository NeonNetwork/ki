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

func ConvertNumberInt32toFloat64(v int32) float64 {
	return float64(v)
}

func ConvertNumberFloat64toInt32(v float64) int32 {
	return int32(v)
}

func MinMax[T Number](values ...T) (T, T) {
	var (
		vmin T
		vmax T
	)

	if len(values) < 1 {
		return T(0.0), T(0.0)
	}

	for _, v := range values {
		if v > vmax {
			vmax = v
		}

		if v < vmin {
			vmin = v
		}
	}

	return vmin, vmax
}