package math

func NewMandelbrot() Mandelbrot {
	return Mandelbrot{initial: 0, maxVal: 4, limit: 2000}
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
		if mb.MaxVal() < real(p)*real(p)+imag(p)*imag(p) {
			return counter
		}
		p = p*p + c
	}
	return 0
}
