package ui

import (
	"npj1610/hypnos-fractal-viewer/render"
)

type UI interface {
	Start()
}

func CreateTextUI(fps int, width int, height int, render render.Render) UI { //}, screen [][]rune) UI {
	//prepare screen
	ui := TextUI{fps: fps, width: width, height: height, render: render}
	ui.dictionary = map[int]rune{0: ' ', 1: 'º', 2: '@'}
	return ui
}
