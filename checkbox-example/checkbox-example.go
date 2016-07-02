package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-basic-widgets/checkbox"
)

func init_() {
}

func tick() {

	bl.Root()
	{
		bl.Pos(64,64)
		bl.Dim(800,600)
		bl.Color(.1,.1,.1)
		bl.Flag(bl.Z_COLOR_SOLID | bl.Z_BORDER_ALL)

		bl.Font("arial", 6)
		bl.FontColor(1,1,1)
		bl.FontNudge(3,3)
		bl.Label("Hello world")

		bl.BorderThickness(bl.FourTwosInt)
		bl.BorderColor(1,1,1)

		bl.Div()
		{
			bl.Id("black2")
			bl.Pos(160, 160)
			bl.Dim(360,360)
			bl.Color(0,0,0)
			bl.BorderThickness(bl.FourOnesInt)
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()

			bl.Font("arial", 12)
			bl.FontColor(1,1,1)
			bl.FontNudge(3,3)
			bl.Label("Hello world")


			checkbox.Id("one").SetLabel("Sweet").On(nil)
		}
		bl.End()
	}
	bl.End()
}

func uninit() {
}

func init() {
	runtime.LockOSThread()
}

func main() {
	bl.Start( 1024, 768, "Bellina v0.2", init_, tick, uninit )

	fmt.Println("bye!")
}


