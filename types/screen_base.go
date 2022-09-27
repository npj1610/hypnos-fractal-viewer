package types

type ScreenBase struct {
	width, height int
}

//pointer to Screen Base allows dynamic screen size
//needs a list of callbacks to update screen size when modified
func NewScreenBase(width, height int) *ScreenBase {
	return &ScreenBase{width: width, height: height}
}

func (s ScreenBase) Width() int {
	return s.width
}

func (s ScreenBase) Height() int {
	return s.height
}
