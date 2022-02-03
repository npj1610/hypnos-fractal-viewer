package fractal

import (
	"math"
	"math/cmplx"
)

type Mandelbrot struct {
	initial			complex128
	maxVal			float64
	limit			int
	grid			[][][]int
	width, height	int
}

func (mb Mandelbrot) Limit() int {
	return mb.limit
}

func (mb Mandelbrot) Height() int {
	return mb.height
}

func (mb Mandelbrot) Width() int {
	return mb.width
}

func (mb Mandelbrot) GetGrid(positions [][]float64) [][][]int {
	upleft := complex(positions[0][0], positions[0][1])
	//downright := complex(positions[1][0], positions[1][1])
	rightstep := complex((positions[1][0] - positions[0][0])/float64(mb.Width()), 0)
	downstep := complex(0, (positions[1][1] - positions[0][1])/float64(mb.Height()))

	mb.grid = make([][][]int, mb.Height())
	for row := range mb.grid {
		mb.grid[row] = make([][]int, mb.Width())
		for point := range mb.grid[row] {

			mb.grid[row][point] = make([]int, 1)
			mb.grid[row][point][0] = mb.calcPoint(upleft + complex(float64(point), 0)*rightstep + complex(float64(row), 0)*downstep)
		}
	}
	return mb.grid
}

func (mb Mandelbrot) calcPoint(c complex128) int {
	p := mb.initial
	var counter int = 0
	for math.Pow(real(p), 2) + math.Pow(imag(p), 2) <= 4 {
		counter = counter + 1
		if mb.Limit() < counter {
			counter = 0
			break
		}
		p = cmplx.Pow(p, complex128(2)) + c
		//fmt.Print(p)
	}
	return counter
}