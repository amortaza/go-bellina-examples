package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina-plugins/drag"
	"github.com/amortaza/go-basic-widgets/edit"
	"github.com/amortaza/go-bellina"
)

func init_() {
}

func tick() {

	bl.Root()
	{
		bl.Pos(64,64)
		bl.Dim(800,600)
		bl.Color(.1,.1,.5)
		bl.Flag(bl.Z_COLOR_SOLID | bl.Z_BORDER_ALL)

		bl.Font("arial", 6)
		bl.FontColor(1,1,1)
		bl.FontNudge(3,3)
		bl.Label("Hello world")

		bl.BorderThickness(bl.FourTwosInt)
		bl.BorderColor(1,1,1)

		bl.Div()
		{
			bl.Id("red")
			bl.Pos(60, 60)
			bl.Dim(320,240)
			bl.Color(.1,0,.0)
			bl.BorderThickness(bl.FourOnesInt)
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()

			edit.Id("zero").On()

			//edit.Id("one")
			//edit.Div()
			//{
				//edit.Pos(10, 10)
				//edit.Size(8)
				//edit.Padding(10, 10, 20)
				//edit.Extend(250)
			//}
			//edit.End()

			//edit.Id("two")
			//edit.Div()
			//{
				//edit.Pos(10, 170)
				//edit.Size(14)
				//edit.Width(200)
			//}
			//edit.End()

			drag.On(nil)
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


