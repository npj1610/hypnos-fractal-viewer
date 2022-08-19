package visual

import (
	m "math"
	"math/cmplx"
)

type TextMBBasicZoom struct {
	rate        float64    //per frame multiplicative zoom rate
	center      complex128 //center point
	centerShift float64    //percentage of centershift per frame
	rotation    float64    //rotation per frame in radians
}

func NewTextMBBasicZoom(r float64, c complex128, cs float64, phi float64) TextMBBasicZoom {
	return TextMBBasicZoom{rate: r, center: c, centerShift: cs, rotation: phi}
}

func (zoom TextMBBasicZoom) UpdateWindow(win TextMBWindow) TextMBWindow {
	winCenter := win.start + win.side/2 + win.top/2

	shift := zoom.center - winCenter
	newCenter := winCenter + complex(zoom.centerShift*real(shift), zoom.centerShift*imag(shift))
	newTop := complex(real(win.top)/zoom.rate, imag(win.top)/zoom.rate)
	newSide := complex(real(win.side)/zoom.rate, imag(win.side)/zoom.rate)

	newStart := newCenter - newTop/2 - newSide/2

	rot := cmplx.Pow(complex(m.E, 0), complex(0, zoom.rotation))
	newTop = newTop * rot
	newSide = newSide * rot

	newStart = newCenter + (newStart-newCenter)*rot

	return TextMBWindow{start: newStart, top: newTop, side: newSide}
}

func (zoom TextMBBasicZoom) StartingWindow() TextMBWindow {
	return TextMBWindow{start: complex(-2, 1), top: complex(3, 0), side: complex(0, -2)}
}
