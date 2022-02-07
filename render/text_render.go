package render

import (
	"npj1610/hypnos-fractal-viewer/fractal"
	"npj1610/hypnos-fractal-viewer/types"
)

type TextRender struct {
	types.ScreenInt

	functions  []func(TextRender, []int) []int
	screenChan chan [][][]int
	fractal    fractal.Fractal
	last_max   float64
}

func (tr TextRender) Functions() []func(TextRender, []int) []int {
	return tr.functions
}

func (tr TextRender) ScreenChan() chan [][][]int {
	return tr.screenChan
}

func (tr TextRender) Start() {
	var pointWeight float64 = 1 / float64(tr.Width()*tr.Height())
	frame := make([][]float64, 2)
	frame[0] = []float64{-2, 1}
	frame[1] = []float64{1, -1}
	var median float64 = 0
	tr.last_max = 40
	for {
		grid := tr.fractal.GetGrid(frame)
		*tr.Screen() = make([][][]int, tr.Height())
		for line := range *tr.Screen() {
			(*tr.Screen())[line] = make([][]int, tr.Width())
			for space := range (*tr.Screen())[line] {
				(*tr.Screen())[line][space] = tr.functions[0](tr, grid[line][space])
				if 0 < (*tr.Screen())[line][space][0] {
					median = median + pointWeight*float64((*tr.Screen())[line][space][0])
				}
			}
		}
		tr.last_max = median / float64(tr.Width()*tr.Height())
		median = 0
		tr.ScreenChan() <- *tr.Screen()
		frame[0][0] = frame[0][0] + 0.01*(-0.7463-frame[0][0])
		frame[0][1] = frame[0][1] + 0.01*(0.1127-frame[0][1])
		frame[1][0] = frame[1][0] - 0.01*(-0.7463-frame[0][0])
		frame[1][1] = frame[1][1] - 0.01*(0.1127-frame[0][1])
	}
}
