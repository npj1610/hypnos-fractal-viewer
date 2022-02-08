package main

import (
	"npj1610/hypnos-fractal-viewer/fractal"
	"npj1610/hypnos-fractal-viewer/render"
	"npj1610/hypnos-fractal-viewer/types"
	"npj1610/hypnos-fractal-viewer/ui"
)

//Vale la pena la concurrencia? (si no, cambia los canales por llamadas a funciones?) Render genera los frames, aumenta los bufers!
//tantos hilos de render como cores, va generando frames en paralelo y los encola en el canal para UI
//Calculadora es funcion? puede también calcular en paralelo
//Si calculadora es funcion, usar un lock a estado interno para permitir reentrada
//Se calculan los pixeles y se dejan huecos (dependencias al frame anterior), una vez calculados, se envian datos por canal
//concurencia? un hilo de la calculadora hace select a las dependencias mientras el otro sigue calculando los pixeles nuevos

func main() {
	screen := types.NewScreenBase(119, 32)
	limit := 10000
	dictionary := map[int]rune{0: '@', 3: '·', 5: 'º', 4: '8', 2: '&', 1: ' '}

	var colorizer fractal.ColorizerMandelbrot = fractal.NewColorizerMandelbrotConfig(len(dictionary))
	var fractal fractal.FractalComplex = fractal.NewMandelbrot(screen, limit, colorizer)
	var render render.Render = render.NewTextRender(screen, fractal)
	var ui ui.UI = ui.NewTextUI(screen, 25, dictionary, render)

	go render.Start()
	ui.Start()
}
