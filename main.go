package main

import (
	"npj1610/hypnos-fractal-viewer/render"
	"npj1610/hypnos-fractal-viewer/ui"
	"npj1610/hypnos-fractal-viewer/fractal"
)

//Vale la pena la concurrencia? (si no, cambia los canales por llamadas a funciones?) Render genera los frames, aumenta los bufers!
//tantos hilos de render como cores, va generando frames en paralelo y los encola en el canal para UI
//Calculadora es funcion? puede también calcular en paralelo
//Si calculadora es funcion, usar un lock a estado interno para permitir reentrada
//Se calculan los pixeles y se dejan huecos (dependencias al frame anterior), una vez calculados, se envian datos por canal
//concurencia? un hilo de la calculadora hace select a las dependencias mientras el otro sigue calculando los pixeles nuevos

func main() {
	width := 119
	height := 32

	var fractal fractal.Fractal = fractal.CreateMandelbrot(width, height, 5000)
	var render render.Render = render.CreateTextRender(width, height, fractal)
	var ui ui.UI = ui.CreateTextUI(25, width, height, render) //, screen)

	go render.Start()
	ui.Start()
}
