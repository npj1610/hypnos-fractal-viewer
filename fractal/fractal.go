package fractal

import (
	"npj1610/hypnos-fractal-viewer/types"
)

type Fractal interface {
	GetGrid(positions [][]float64) [][][]int
}

func NewMandelbrot(screen types.ScreenBasic, limit int) Mandelbrot {
	return Mandelbrot{ScreenInt: types.ScreenInt{ScreenBasic: screen}, initial: 0, maxVal: 4, limit: limit}
}
