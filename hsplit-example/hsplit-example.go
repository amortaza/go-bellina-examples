package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina-plugins/resize"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-basic-widgets/hsplit"
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
			bl.Id("black")
			bl.Pos(60, 60)
			bl.Dim(360,360)
			bl.Color(0,0,0)
			bl.BorderThickness(bl.FourOnesInt)
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()

			bl.Div()
			{
				bl.Id("green")
				bl.Pos(10, 10)
				bl.Dim(60,60)
				bl.Color(.0,.1,.0)
				bl.BorderThickness(bl.FourOnesInt)
				bl.BorderColor(1,1,1)
				bl.BorderTopsCanvas()
			}
			bl.End()

			bl.Div()
			{
				bl.Id("blue")
				bl.Pos(160, 160)
				bl.Dim(60,60)
				bl.Color(0,0,.5)
				bl.BorderThickness(bl.FourOnesInt)
				bl.BorderColor(1,1,1)
				bl.BorderTopsCanvas()
			}
			bl.End()

			hsplit.Use("cool2")
		}
		bl.End()

		bl.Div()
		{
			bl.Id("black2")
			bl.Pos(260, 260)
			bl.Dim(360,360)
			bl.Color(0,0,0)
			bl.BorderThickness(bl.FourOnesInt)
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()

			bl.Div()
			{
				bl.Id("green2")
				bl.Pos(10, 10)
				bl.Dim(60,60)
				bl.Color(.0,.1,.0)
				bl.BorderThickness(bl.FourOnesInt)
				bl.BorderColor(1,1,1)
				bl.BorderTopsCanvas()
			}
			bl.End()

			bl.Div()
			{
				bl.Id("blue2")
				bl.Pos(160, 160)
				bl.Dim(60,60)
				bl.Color(0,0,.5)
				bl.BorderThickness(bl.FourOnesInt)
				bl.BorderColor(1,1,1)
				bl.BorderTopsCanvas()
			}
			bl.End()

			hsplit.Use("cool")
		}
		bl.End()

		resize.Use()
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


