package visual

import (
	"npj1610/hypnos-fractal-viewer/math"
	"npj1610/hypnos-fractal-viewer/types"
)

func NewTextRender(screen types.ScreenBase, fractal math.Mandelbrot, colorizer TextMandelbrotColorizer, zoom TextMBZoom) TextRender {
	var screenChan = make(chan types.TextScreen, 100)
	return TextRender{types.NewTextScreen(screen), screenChan, fractal, colorizer, zoom}
}

type TextRender struct {
	types.TextScreen

	screenChan chan types.TextScreen
	fractal    math.Mandelbrot
	colorizer  TextMandelbrotColorizer
	zoom       TextMBZoom
}

func (tr TextRender) ScreenChan() chan types.TextScreen {
	return tr.screenChan
}

func (tr TextRender) Start() {
	win := tr.zoom.StartingWindow()

	for {
		topLeft := win.start
		rightstep := complex(real(win.top)/float64(tr.Width()), imag(win.top)/float64(tr.Width()))
		downstep := complex(real(win.side)/float64(tr.Height()), imag(win.side)/float64(tr.Height()))

		tr.colorizer.PreCalc(tr.fractal, &tr.TextScreen)

		for row := range *tr.Screen() {
			for point := range (*tr.Screen())[row] {
				complexPoint := topLeft + complex(float64(point), 0)*rightstep + complex(float64(row), 0)*downstep
				(*tr.Screen())[row][point] = tr.colorizer.ForPoint(tr.fractal, row, point, tr.fractal.CalcPoint(complexPoint))
			}
		}

		tr.colorizer.PostCalc(tr.fractal, &tr.TextScreen)

		tr.ScreenChan() <- tr.TextScreen.Copy()

		win = tr.zoom.UpdateWindow(win)
	}
}

/*
BENCHMARKING
import (
    "log"
    "os"
    "runtime"
    "runtime/pprof"
)

cpu, err := os.Create("prof\\cpu" + strconv.FormatInt(int64(i), 10) + ".prof")
if err != nil {
	log.Fatal(err)
}
pprof.StartCPUProfile(cpu)


pprof.StopCPUProfile()
*/
