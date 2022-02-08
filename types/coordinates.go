package types

type CoordinatesComplex struct {
	topLeft, bottomRight complex128
}

func (cc CoordinatesComplex) TopLeft() complex128 {
	return cc.topLeft
}

func (cc CoordinatesComplex) BottomRight() complex128 {
	return cc.bottomRight
}

func NewCoordinatesComplex(top, bottom, left, right float64) CoordinatesComplex {
	return CoordinatesComplex{topLeft: complex(left, top), bottomRight: complex(right, bottom)}
}
