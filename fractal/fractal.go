package fractal

type Fractal interface {
	GetGrid(positions [][]float64) [][][]int
}

func CreateMandelbrot(width, height, limit int) Mandelbrot {
	return Mandelbrot{initial: 0, maxVal: 4, width: width, height: height, limit:limit}
}