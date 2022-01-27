package main

import (
	"npj1610/hypnos-fractal-viewer/render"
	"npj1610/hypnos-fractal-viewer/ui"
)

func main() {
	width := 119
	height := 32

	var render render.Render = render.CreateTextRender(width, height)
	var ui ui.UI = ui.CreateTextUI(25, width, height, render) //, screen)

	go render.Start()
	ui.Start()
}
