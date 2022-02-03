package ui

import (
	"fmt"
	"strings"
	"time"

	"npj1610/hypnos-fractal-viewer/render"
)

type TextUI struct {
	fps, width, height int
	screen             [][][]int
	render             render.Render
	dictionary         map[int]rune
}

func (ui TextUI) convert(n int) rune {
	return ui.dictionary[n]
}

func (ui TextUI) FPS() int {
	return ui.fps
}

func (ui TextUI) Width() int {
	return ui.width
}

func (ui TextUI) Height() int {
	return ui.height
}

func (ui TextUI) Render() render.Render {
	return ui.render
}

func (ui TextUI) Start() {
	for {
		//send frame request, get frame (select for waiting input/ctrl+c?)
		ui.screen = <-ui.render.ScreenChan()

		var sb strings.Builder
		for y := 0; y < ui.Height(); y++ {
			sb.WriteRune('\n')
			for x := 0; x < ui.Width(); x++ {
				sb.WriteRune(ui.convert(ui.screen[y][x][0]))
			}
		}
		fmt.Print(sb.String())
		time.Sleep(time.Duration(int(time.Second) / ui.FPS()))
	}
}
