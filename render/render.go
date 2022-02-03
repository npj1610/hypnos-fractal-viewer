package render

import (
	"npj1610/hypnos-fractal-viewer/fractal"
)

//una sola clase render ajustable mediante metodo factoria
//une unas dimensiones de entrada con unas de salida
//Coloración y Movimiento son submódulos (compartidos? declarativos?)

type Render interface {
	//SetOutputRange([]types.IntRange) //[2]int in types/types.go!
	Start()
	ScreenChan() chan [][][]int
}

func CreateTextRender(width int, height int, fractal fractal.Fractal) Render {
	tr := TextRender{width: width, height: height, fractal: fractal}
	tr.screenChan = make(chan [][][]int, 100)
	function := func(tr TextRender, point []int) []int {
		out := make([]int, 1)
		if point[0] < 1 {
			out[0] = 0
		} else if float64(point[0]) < tr.last_max/2 {
			out[0] = 1
		} else {
			out[0] = 2
		}
		return out
	}
	tr.functions = []func(TextRender, []int) []int{function}
	return tr
}
