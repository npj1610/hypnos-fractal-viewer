package fractal

import (
	"npj1610/hypnos-fractal-viewer/types"
)

type FractalComplex interface {
	CalcScreen(positions types.CoordinatesComplex) types.ScreenInt
}

func NewMandelbrot(screen types.ScreenBase, limit int, colorizer ColorizerMandelbrot) Mandelbrot {
	return Mandelbrot{ScreenInt: types.ScreenInt{ScreenBase: screen}, initial: 0, maxVal: 4, limit: limit, colorizer: colorizer}
}
