package main

import (
	"runtime"
	"xel"
	"g4"
	"g4/ace"
)

var troll *ace.Texture
var box *ace.Texture

func onLoop() {
	g4.Ortho(xel.WinWidth, xel.WinHeight)

	g4.Clear()

	g4.TextureRect(box, 90,175,400,400,[]float32{1,1,1,1})
	g4.TextureRect(troll, 150,10,400,400,[]float32{1,0,1,0})

	g4.RgbRect(50, 10, 200, 200, []float32{1.0,0.0,0.0,1.0}, []float32{0.0,1.0,0.0,1.0}, []float32{0.0,0.0,1.0,1.0}, []float32{0.0,0.0,0.0,1.0})
	g4.RgbRect(50, 250, 200, 200, []float32{1.0,0.0,0.0,1.0}, []float32{0.0,1.0,0.0,1.0}, []float32{0.0,0.0,1.0,0.0}, []float32{0.0,0.0,0.0,0.0})
}

func onAfterGL() {
	g4.Init()

	troll = ace.NewTexture()
	troll.LoadImage("images/troll2.png")

	box = ace.NewTexture()
	box.LoadImage("images/crate.png")
}

func onBeforeDelete() {
	troll.Free()
	box.Free()

	g4.Uninit()
}

func onResize(width,height int32) {
	g4.Viewport(width,height)
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


