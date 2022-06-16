package visual

import (
	"npj1610/hypnos-fractal-viewer/math"
	"npj1610/hypnos-fractal-viewer/types"
)

type TextMandelbrotColorizer interface {
	PreCalc(mb math.Mandelbrot, screen *types.TextScreen)
	ForPoint(mb math.Mandelbrot, row int, column int, value int) rune
	PostCalc(mb math.Mandelbrot, screen *types.TextScreen)
}
