package fractal

type ColorizerMandelbrot interface {
	PreCalc(mb *Mandelbrot)
	ForPoint(mb *Mandelbrot, row int, column int, coord complex128)
	ForRow(mb *Mandelbrot, row int, line float64)
	PostCalc(mb *Mandelbrot)
}

type ColorizerMandelbrotConfig struct {
	histogram, divisions []int
	total, colors        int
}

func (cmc *ColorizerMandelbrotConfig) PreCalc(mb *Mandelbrot) {
	cmc.histogram = make([]int, mb.Limit())
	for v := range cmc.histogram {
		cmc.histogram[v] = 0
	}
	cmc.divisions = make([]int, cmc.colors)
	cmc.total = 0
}

func (cmc *ColorizerMandelbrotConfig) ForPoint(mb *Mandelbrot, row int, column int, coord complex128) {
	value := (*mb.Screen())[row][column][0]

	if value != 0 {
		cmc.histogram[value]++
		cmc.total++
	}
}

func (cmc *ColorizerMandelbrotConfig) ForRow(mb *Mandelbrot, row int, line float64) {}

func (cmc *ColorizerMandelbrotConfig) PostCalc(mb *Mandelbrot) {
	//acumulates
	for v := range cmc.histogram {
		if v != 0 {
			cmc.histogram[v] = cmc.histogram[v] + cmc.histogram[v-1]
		}
	}
	size := cmc.total / cmc.colors
	nextDivision := 0
	for v := range cmc.histogram {
		if size*nextDivision <= cmc.histogram[v] {
			cmc.divisions[nextDivision] = v
			nextDivision++
			if cmc.colors <= nextDivision {
				cmc.divisions[nextDivision-1] = len(cmc.histogram) - 1
				break
			}

		}
	}
	//colorizes
	for row := range *mb.Screen() {
		for column := range (*mb.Screen())[row] {
			value := (*mb.Screen())[row][column][0]
			for division := range cmc.divisions {
				if value <= cmc.divisions[division] {
					value = division
					break
				}
			}
			(*mb.Screen())[row][column][0] = value
		}
	}

}

func NewColorizerMandelbrotConfig(colors int) *ColorizerMandelbrotConfig {
	return &ColorizerMandelbrotConfig{colors: colors}
}
