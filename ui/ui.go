package ui

import (
	"npj1610/hypnos-fractal-viewer/render"
	"npj1610/hypnos-fractal-viewer/types"
)

type UI interface {
	Start()
}

func CreateTextUI(screen types.ScreenBasic, fps int, render render.Render) UI {
	ui := TextUI{ScreenInt: types.ScreenInt{ScreenBasic: screen}, fps: fps, render: render}
	ui.dictionary = map[int]rune{0: ' ', 1: 'º', 2: '@'}
	return ui
}
