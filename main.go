package main

import (
	"fmt"
	"npj1610/hypnos-fractal-viewer/ui"
)

func main() {
	fmt.Println("test")
	var ui ui.UI = ui.CreateTextUI(25)
	ui.Start()
}
