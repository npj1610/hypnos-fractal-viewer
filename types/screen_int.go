package types

type ScreenInt struct {
	*ScreenBase
	screen [][]int
}

func (s *ScreenInt) Screen() *[][]int {
	return &s.screen
}

//Pointer to Screen Int allows efficient concurrency
func NewScreenInt(sb *ScreenBase) *ScreenInt {
	output := ScreenInt{ScreenBase: sb}
	output.screen = make([][]int, sb.Height())
	for row := range output.screen {
		output.screen[row] = make([]int, sb.Width())
	}
	return &output
}
