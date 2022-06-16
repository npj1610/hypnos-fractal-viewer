package visual

import (
	"npj1610/hypnos-fractal-viewer/math"
	"npj1610/hypnos-fractal-viewer/types"
)

type TextMBAutobalance struct {
	types.ScreenInt

	histogram, divisions []int
	total                int
	dictionary           map[int]rune
}

func (colorizer *TextMBAutobalance) PreCalc(mb math.Mandelbrot, screen *types.TextScreen) {
	colorizer.histogram = make([]int, mb.Limit())
	for v := range colorizer.histogram {
		colorizer.histogram[v] = 0
	}
	colorizer.divisions = make([]int, len(colorizer.dictionary))
	colorizer.total = 0

	colorizer.ScreenInt = types.NewScreenInt(screen.ScreenBase)
}

func (colorizer *TextMBAutobalance) ForPoint(mb math.Mandelbrot, row int, column int, value int) rune {
	//ignores values in the set when balancing colors
	if value != 0 {
		colorizer.histogram[value]++
		colorizer.total++
	}

	//stores temporary value
	(*colorizer.Screen())[row][column] = value

	return ' '
}

func (colorizer *TextMBAutobalance) PostCalc(mb math.Mandelbrot, screen *types.TextScreen) {
	//acumulates
	for v := range colorizer.histogram {
		if v != 0 {
			colorizer.histogram[v] = colorizer.histogram[v] + colorizer.histogram[v-1]
		}
	}
	//divides
	size := colorizer.total / len(colorizer.dictionary)
	nextDivision := 0
	for v := range colorizer.histogram {
		if size*nextDivision <= colorizer.histogram[v] {
			colorizer.divisions[nextDivision] = v
			nextDivision++
			//if there are spares after last division, include them
			if len(colorizer.dictionary) <= nextDivision {
				colorizer.divisions[nextDivision-1] = len(colorizer.histogram) - 1
				break
			}

		}
	}
	//colorizes
	for row := range *colorizer.Screen() {
		for column := range (*colorizer.Screen())[row] {
			value := (*colorizer.Screen())[row][column]
			for division := range colorizer.divisions {
				if value <= colorizer.divisions[division] {
					value = division
					break
				}
			}
			(*screen.Screen())[row][column] = colorizer.dictionary[value]
		}
	}

}

func NewTextMBAutobalance(dictionary map[int]rune) *TextMBAutobalance {
	return &TextMBAutobalance{dictionary: dictionary}
}
