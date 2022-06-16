package math

import (
	"math"
	"math/cmplx"
)

func NewMandelbrot() Mandelbrot {
	return Mandelbrot{initial: 0, maxVal: 4, limit: 5000}
}

type Mandelbrot struct {
	initial complex128
	maxVal  float64
	limit   int
}

func (mb Mandelbrot) Limit() int {
	return mb.limit
}

func (mb Mandelbrot) MaxVal() float64 {
	return mb.maxVal
}

func (mb Mandelbrot) Initial() complex128 {
	return mb.initial
}

func (mb Mandelbrot) CalcPoint(c complex128) int {
	p := mb.Initial()
	for counter := 0; counter < mb.Limit(); counter++ {
		if mb.MaxVal() < math.Pow(real(p), 2)+math.Pow(imag(p), 2) {
			return counter
		}
		p = cmplx.Pow(p, complex128(2)) + c
	}
	return 0
}
