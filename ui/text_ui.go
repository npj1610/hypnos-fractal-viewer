package ui

import (
	"fmt"
	"time"
)

type TextUI int

func (ui TextUI) FPS() int {
	return int(ui)
}

func (ui TextUI) Start() {
	for {
		fmt.Println("HELLO")
		time.Sleep(time.Duration(float64(time.Second) / float64(ui.FPS())))
	}
}
