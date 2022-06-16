package visual

import (
	"npj1610/hypnos-fractal-viewer/math"
	"npj1610/hypnos-fractal-viewer/types"
)

func NewTextRender(screen types.ScreenBase, fractal math.Mandelbrot, colorizer TextMandelbrotColorizer) TextRender {
	var screenChan = make(chan types.TextScreen, 100)
	return TextRender{types.NewTextScreen(screen), screenChan, fractal, colorizer}
}

type TextRender struct {
	types.TextScreen

	screenChan chan types.TextScreen
	fractal    math.Mandelbrot
	colorizer  TextMandelbrotColorizer
}

func (tr TextRender) ScreenChan() chan types.TextScreen {
	return tr.screenChan
}

func (tr TextRender) Start() {
	topLeft := complex(-2, 1)
	bottomRight := complex(1, -1)

	for {
		size := bottomRight - topLeft
		rightstep := complex(real(size)/float64(tr.Width()), 0)
		downstep := complex(0, imag(size)/float64(tr.Height()))

		tr.colorizer.PreCalc(tr.fractal, &tr.TextScreen)

		for row := range *tr.Screen() {
			for point := range (*tr.Screen())[row] {
				complexPoint := topLeft + complex(float64(point), 0)*rightstep + complex(float64(row), 0)*downstep
				(*tr.Screen())[row][point] = tr.colorizer.ForPoint(tr.fractal, row, point, tr.fractal.CalcPoint(complexPoint))
			}
		}

		tr.colorizer.PostCalc(tr.fractal, &tr.TextScreen)

		tr.ScreenChan() <- tr.TextScreen

		topLeft += complex(
			0.01*(-0.7463-real(topLeft)),
			0.01*(0.1127-imag(topLeft)),
		)
		bottomRight -= complex(
			0.01*(-0.7463-real(topLeft)),
			0.01*(0.1127-imag(topLeft)),
		)
	}
}
