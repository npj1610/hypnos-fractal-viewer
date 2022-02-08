package render

import (
	"npj1610/hypnos-fractal-viewer/fractal"
	"npj1610/hypnos-fractal-viewer/types"
)

type Render interface {
	Start()
	ScreenChan() chan types.ScreenInt
}

func NewTextRender(screen types.ScreenBase, f fractal.FractalComplex) Render {
	tr := TextRender{ScreenInt: types.ScreenInt{ScreenBase: screen}, fractal: f}
	tr.screenChan = make(chan types.ScreenInt, 100)
	return tr
}
