package ki

import (
	"github.com/neonnetwork/ki/pkg/structure"
)

type IterateDataCompute struct {
	Position structure.Vector2[float64]
	Size     structure.Vector2[float64]
	Path     []structure.BinaryTreeDirection
}

func (data IterateDataCompute) Copy() IterateDataCompute {
	return IterateDataCompute{
		Size: data.Size,
		Path: data.Path,
	}
}

func (data IterateDataCompute) AddPath(value structure.BinaryTreeDirection) IterateDataCompute {
	data.Path = append(data.Path, value)

	return data
}

func (data IterateDataCompute) SetSize(value structure.Vector2[float64]) IterateDataCompute {
	data.Size = value.Copy()

	return data
}

func (data IterateDataCompute) SetPosition(value structure.Vector2[float64]) IterateDataCompute {
	data.Position = value.Copy()

	return data
}

func (data IterateDataCompute) AddPosition(value structure.Vector2[float64]) IterateDataCompute {
	data.Position = data.Position.Add(value)

	return data
}

func (engine *Engine) Compute() (err error) {
	err = engine.ComputeWindows()
	if err != nil {
		return
	}

	return
}

func (engine *Engine) ComputeWindows() (err error) {
	err = engine.IterateWindows(
		func(window Window, data any) (any, error) {
			var (
				value IterateDataCompute
			)

			value = data.(IterateDataCompute)

			// window.SetPositionAbsolute(structure.MapVector2(value.Position, structure.ConvertNumberFloat64toInt32))
			window.SetSizeAbsolute(structure.MapVector2(value.Size, structure.ConvertNumberFloat64toInt32))

			if engine.selected != nil {
				window.SetSelected(window.Id() == engine.selected.Value().Id())
			}

			value = value.Copy().
				AddPath(window.Side()).
				SetSize(
					structure.MapVector2(window.Size(), structure.ConvertNumberInt32toFloat64).
						Div(structure.NewVector2[float64](EngineWindowUnit, EngineWindowUnit)).
						Mul(value.Size),
				)

			return value, nil
		},
		IterateDataCompute{
			Position: structure.NewVector2[float64](0.0, 0.0),
			Size:     structure.MapVector2(engine.Screen(), structure.ConvertNumberInt32toFloat64),
			Path:     make([]structure.BinaryTreeDirection, 0),
		})
	if err != nil {
		return
	}

	return
}
