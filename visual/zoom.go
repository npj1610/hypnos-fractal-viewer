package visual

type TextMBWindow struct {
	start     complex128
	top, side complex128
}

type TextMBZoom interface {
	StartingWindow() TextMBWindow
	UpdateWindow(win TextMBWindow) TextMBWindow
}
