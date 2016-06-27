package main

import (
	"runtime"
	"xel"
	"g4"
)

var arial18 *g4.G4Font
var str1 *g4.StringTexture
var canvas *g4.Canvas

func onLoop() {
	g4.Clear(.4,.6,.6,1)

	g4.PushViewport(xel.WinWidth, xel.WinHeight)
	g4.PushOrtho(xel.WinWidth, xel.WinHeight)

	//g4.DrawColorRect(50, 10, 200, 200, []float32{1.0,0.0,0.0,1.0}, []float32{0.0,1.0,0.0,1.0}, []float32{0.0,0.0,1.0,1.0}, []float32{0.0,0.0,0.0,1.0})
	//g4.DrawColorRect(50, 250, 200, 200, []float32{1.0,0.0,0.0,1.0}, []float32{0.0,1.0,0.0,1.0}, []float32{0.0,0.0,1.0,0.0}, []float32{0.0,0.0,0.0,0.0})

	g4.DrawStringRect(str1, 10, 10, []float32{1, 1, 1}, []float32{0.5, .5, .5}, 1)

	canvas.Begin()
	{
		g4.Clear(.4,.6,.6,1)

		g4.DrawStringRect(str1, 10, 0, []float32{1, 1, 1}, []float32{0.5, .5, .5}, 1)
	}
	canvas.End()

	canvas.Paint(true, 0, 25,[]float32{1,1,1,1})

	g4.PopOrtho()
	g4.PopViewport()
}

func onAfterGL() {
	g4.Init()

	f := g4.LoadTrueTypeFromFile("github.com/amortaza/go-bellina-examples/assets/fonts/arial.ttf")
	arial18 = g4.NewG4Font(f, 8)
	str1 = g4.NewStringTexture("Welcome to Clown World!", arial18)
	canvas = g4.NewCanvas(320,100)
}

func onBeforeDelete() {
	canvas.Free()
	arial18.Free()
	str1.Free()
	g4.Uninit()
}

func onResize(width,height int32) {
}

func onMouseMove(x,y int32) {
}

func onMouseButton(button xel.MouseButton, action xel.ButtonAction) {
}

func onKey(key xel.KeyboardKey, action xel.ButtonAction, a,b,c bool) {
}

func init() {
	runtime.LockOSThread()
}

func main() {

	xel.Init(800, 600)

	xel.SetCallbacks(onAfterGL, onLoop, onBeforeDelete, onResize, onMouseMove, onMouseButton, onKey)

	xel.Loop()

	xel.Uninit()
}


