package ui

import (
	"npj1610/hypnos-fractal-viewer/render"
	"npj1610/hypnos-fractal-viewer/types"
)

type UI interface {
	Start()
}

func NewTextUI(screen types.ScreenBase, fps int, dictionary map[int]rune, render render.Render) UI {
	ui := TextUI{ScreenInt: types.ScreenInt{ScreenBase: screen}, fps: fps, render: render}
	ui.dictionary = dictionary
	return ui
}
