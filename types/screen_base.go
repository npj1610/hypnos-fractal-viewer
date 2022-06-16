package types

type ScreenBase struct {
	width, height int
}

func NewScreenBase(width, height int) ScreenBase {
	return ScreenBase{width: width, height: height}
}

func (s ScreenBase) Width() int {
	return s.width
}

func (s ScreenBase) Height() int {
	return s.height
}
