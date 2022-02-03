package render

import (
	"npj1610/hypnos-fractal-viewer/fractal"
)

type TextRender struct {
	functions              []func(TextRender, []int) []int
	screenChan             chan [][][]int
	screen                 [][][]int
	fractal                fractal.Fractal
	counter, width, height int
	last_max               float64
}

func (tr TextRender) Width() int {
	return tr.width
}

func (tr TextRender) Height() int {
	return tr.height
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
		tr.screen = make([][][]int, tr.Height())
		for line := range tr.screen {
			tr.screen[line] = make([][]int, tr.Width())
			for space := range tr.screen[line] {
				tr.screen[line][space] = tr.functions[0](tr, grid[line][space])
				if 0 < tr.screen[line][space][0] {
					median = median + pointWeight*float64(tr.screen[line][space][0])
				}
			}
		}
		tr.last_max = median / float64(tr.Width()*tr.Height())
		median = 0
		tr.ScreenChan() <- tr.screen
		frame[0][0] = frame[0][0] + 0.01*(-0.7463-frame[0][0])
		frame[0][1] = frame[0][1] + 0.01*(0.1127-frame[0][1])
		frame[1][0] = frame[1][0] - 0.01*(-0.7463-frame[0][0])
		frame[1][1] = frame[1][1] - 0.01*(0.1127-frame[0][1])
	}
}
