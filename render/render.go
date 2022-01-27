package render

type Render interface {
	//SetOutputRange([]types.IntRange) //[2]int in types/types.go!
	Start()
	RequestChan() chan bool
	ScreenChan() chan [][][]int
}

func CreateTextRender(width int, height int) Render {
	tr := TextRender{}
	tr.requestChan = make(chan bool, 1)
	tr.screenChan = make(chan [][][]int, 1)
	tr.screen = make([][][]int, height)
	for line := range tr.screen {
		tr.screen[line] = make([][]int, width)
		for space := range tr.screen[line] {
			tr.screen[line][space] = make([]int, 1)
			tr.screen[line][space][0] = 1 + (line+space)%2
		}
	}
	return tr
}
