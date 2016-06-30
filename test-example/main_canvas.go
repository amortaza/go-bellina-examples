package main

import (
	"g4"
	"xel"
	"runtime"
)

var troll *g4.Texture
var box *g4.Texture
var str1 *g4.StringTexture
var str2 *g4.StringTexture
var arial18 *g4.G4Font
var canvas *g4.Canvas
var x int32

func onLoop() {
	g4.Clear(.4,.6,.6,1)

	g4.PushViewport(xel.WinWidth, xel.WinHeight)
	g4.PushOrtho(xel.WinWidth, xel.WinHeight)

	g4.DrawTextureRect(box, x,175,400,400,[]float32{1,1,1,1})
	g4.DrawTextureRect(troll, 150,10,400,400,[]float32{1,0,1,0})

	g4.DrawColorRect(50, 10, 200, 200, []float32{1.0,0.0,0.0,1.0}, []float32{0.0,1.0,0.0,1.0}, []float32{0.0,0.0,1.0,1.0}, []float32{0.0,0.0,0.0,1.0})
	g4.DrawColorRect(50, 250, 200, 200, []float32{1.0,0.0,0.0,1.0}, []float32{0.0,1.0,0.0,1.0}, []float32{0.0,0.0,1.0,0.0}, []float32{0.0,0.0,0.0,0.0})

	x+=3
	canvas.Begin()
	{
		canvas.Clear(0,0,0)

		g4.DrawStringRect(str1, x/3, x/3, []float32{1, 0, 1}, []float32{0, .5, .86}, 1)
		g4.DrawStringRect(str1, 0, (arial18.Height+1)*6, []float32{1, 0, 1}, []float32{1, 1, 1}, 1)
	}
	canvas.End()

	canvas.Paint(50,300,[]float32{.6,.6,.6,.9})

	g4.PopOrtho()
	g4.PopViewport()
}

func onAfterGL() {
	g4.Init()

	arial18 = g4.NewG4Font("github.com/amortaza/go-bellina-examples/assets/fonts/arial.ttf", 18)

	troll = g4.NewTexture()
	troll.LoadImage("github.com/amortaza/go-bellina-examples/assets/images/troll2.png")

	box = g4.NewTexture()
	box.LoadImage("github.com/amortaza/go-bellina-examples/assets/images/crate.png")

	str1 = g4.NewStringTexture("Welcome    to Clown World!", arial18)
	str2 = g4.NewStringTexture("to Clown World!", arial18)

	canvas = g4.NewCanvas(640,300)
}

func onBeforeDelete() {
	troll.Free()
	box.Free()
	str1.Free()
	str2.Free()
	canvas.Free()

	g4.Uninit()
}

func onResize(width,height int32) {
}

func onMouseMove(x,y int) {
}

func onMouseButton(button xel.MouseButton, action xel.ButtonAction) {
}

func onKey(key xel.KeyboardKey, action xel.ButtonAction) {
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


