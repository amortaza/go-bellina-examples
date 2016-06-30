package main

import (
	"g4"
	"xel"
	"runtime"
	"bellina"
	"eos"
)

func scene() {
	bl.Root()
	{
		bl.Pos(0,0)
		bl.Dim(640,480)
		bl.Color(.31,.731,.31)
		bl.Flag(bl._COLOR_SOLID | bl._BORDER_ALL)

		bl.Font("arial", 18)
		bl.FontColor(1,0,1)
		bl.Label("Hello world")

		bl.BorderThickness([]int32{2,2,2,2})
		bl.BorderColor(1,1,1)

		bl.Div()
		{
			bl.ID("one")
			bl.Pos(60, 60)
			bl.Dim(200,200)
			bl.Color(.871,0,.871)
			bl.BorderThickness([]int32{1,1,1,1})
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()
		}
		bl.End()
	}
	bl.End()
}

func onLoop() {
	scene()

	g4.Clear(.4,.6,.6,1)

	g4.PushView(xel.WinWidth, xel.WinHeight)

	//canvas := eos.RenderCanvas(bl.Root_Node)
	//canvas.Paint(false, 64, 64, nil)
	//canvas.Free()

	g4.PopView()
}

func onAfterGL() {
	bl.Init()
}

func onBeforeDelete() {
	bl.Uninit()
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


