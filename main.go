package main

import (
	"npj1610/hypnos-fractal-viewer/inout"
	"npj1610/hypnos-fractal-viewer/math"
	"npj1610/hypnos-fractal-viewer/types"
	"npj1610/hypnos-fractal-viewer/visual"
)

//Vale la pena la concurrencia? (si no, cambia los canales por llamadas a funciones?) Render genera los frames, aumenta los bufers!
//tantos hilos de render como cores, va generando frames en paralelo y los encola en el canal para UI
//Calculadora es funcion? puede también calcular en paralelo
//Si calculadora es funcion, usar un lock a estado interno para permitir reentrada
//Se calculan los pixeles y se dejan huecos (dependencias al frame anterior), una vez calculados, se envian datos por canal
//concurencia? un hilo de la calculadora hace select a las dependencias mientras el otro sigue calculando los pixeles nuevos

func main() { //200, 60; 119,32
	screen := types.NewScreenBase(200, 60)
	//dictionary := map[int]rune{0: '@', 1: '·', 2: 'º', 3: '8', 4: '&', 5: ' '}
	//dictionary := map[int]rune{1: 'o', 2: '.', 3: ':', 4: '-', 5: '=', 6: '+', 7: '*', 8: '#', 9: '%', 0: '@'}
	chars := [70]rune{' ', '$', '@', 'B', '%', '8', '&', 'W', 'M', '#', '*', 'o', 'a', 'h', 'k', 'b', 'd', 'p', 'q', 'w',
		'm', 'Z', 'O', '0', 'Q', 'L', 'C', 'J', 'U', 'Y', 'X', 'z', 'c', 'v', 'u', 'n', 'x', 'r', 'j', 'f', 't', '/',
		'\\', '|', '(', ')', '1', '{', '}', '[', ']', '?', '-', '_', '+', '~', '<', '>', 'i', '!', 'l', 'I', ';', ':',
		',', '"', '^', '`', '\'', '.'}
	dictionary := map[int]rune{}
	for i := range chars {
		dictionary[i] = chars[i]
	}

	var colorizer visual.TextMandelbrotColorizer = visual.NewTextMBAutobalance(dictionary)
	var fractal math.Mandelbrot = math.NewMandelbrot()
	var render visual.TextRender = visual.NewTextRender(screen, fractal, colorizer)
	var ui inout.TextUI = inout.NewTextUI(screen, 15, render)

	go render.Start()
	ui.Start()
}
