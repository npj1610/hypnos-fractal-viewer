package types

type ScreenBasic struct {
	width, height int
}

func NewScreenBasic(width, height int) ScreenBasic {
	return ScreenBasic{width: width, height: height}
}

func (s ScreenBasic) Width() int {
	return s.width
}

func (s ScreenBasic) Height() int {
	return s.height
}

type ScreenInt struct {
	ScreenBasic
	screen [][][]int
}

func (s *ScreenInt) Screen() *[][][]int {
	return &s.screen
}
