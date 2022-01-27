package render

type TextRender struct {
	functions   []func([]int) []int
	screenChan  chan [][][]int
	requestChan chan bool //might carry some info later
	screen      [][][]int
}

func (tr TextRender) Functions() []func([]int) []int {
	return tr.functions
}

func (tr TextRender) ScreenChan() chan [][][]int {
	return tr.screenChan
}

func (tr TextRender) RequestChan() chan bool {
	return tr.requestChan
}

func (tr TextRender) Start() {
	for {
		for range tr.RequestChan() {
			tr.ScreenChan() <- tr.screen
		}
	}
}
