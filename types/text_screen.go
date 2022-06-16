package types

type TextScreen struct {
	ScreenBase
	screen [][]rune
}

func (s *TextScreen) Screen() *[][]rune {
	return &s.screen
}

func NewTextScreen(sb ScreenBase) TextScreen {
	output := TextScreen{ScreenBase: sb}
	output.screen = make([][]rune, sb.Height())
	for row := range output.screen {
		output.screen[row] = make([]rune, sb.Width())
	}
	return output
}
