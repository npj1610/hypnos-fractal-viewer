package types

type TextScreen struct {
	ScreenBase
	screen [][]rune
}

func (s *TextScreen) Screen() *[][]rune {
	return &s.screen
}

func (s TextScreen) Copy() TextScreen {
	output := TextScreen{ScreenBase: s.ScreenBase}
	output.screen = make([][]rune, s.Height())
	for row := range output.screen {
		output.screen[row] = make([]rune, s.Width())
		for column := range output.screen[row] {
			output.screen[row][column] = s.screen[row][column]
		}
	}
	return output
}

func NewTextScreen(sb ScreenBase) TextScreen {
	output := TextScreen{ScreenBase: sb}
	output.screen = make([][]rune, sb.Height())
	for row := range output.screen {
		output.screen[row] = make([]rune, sb.Width())
	}
	return output
}
