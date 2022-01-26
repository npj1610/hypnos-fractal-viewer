package ui

type UI interface {
	Start()
	FPS() int
}

func CreateTextUI(fps int) UI {
	return TextUI(fps)
}
