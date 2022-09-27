package inout

import (
	"fmt"
	"strings"
	"time"

	"npj1610/hypnos-fractal-viewer/types"
	"npj1610/hypnos-fractal-viewer/visual"
)

func NewTextUI(screen *types.ScreenBase, fps int, render visual.TextRender) TextUI {
	return TextUI{&types.TextScreen{ScreenBase: screen}, fps, render}
}

type TextUI struct {
	*types.TextScreen

	fps    int
	render visual.TextRender
}

func (ui TextUI) FPS() int {
	return ui.fps
}

func (ui TextUI) Start() {
	for {
		//get frame (select for waiting input/ctrl+c?)
		ui.TextScreen = <-ui.render.ScreenChan()

		var sb strings.Builder
		for y := 0; y < ui.Height(); y++ {
			sb.WriteRune('\n')
			for x := 0; x < ui.Width(); x++ {
				sb.WriteRune((*ui.Screen())[y][x])
			}
		}
		fmt.Print(sb.String())
		time.Sleep(time.Duration(int(time.Second) / ui.FPS()))
	}
}
