package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/resize"
)

func init_() {
}

func tick() {

	bl.Root()
	{
		bl.Pos(64,64)
		bl.Dim(800,600)
		bl.Color(.1,.1,.1)
		bl.BorderThickness([]int{2,2,2,2})
		bl.BorderColor(1,1,1)
		bl.Flag(bl.Z_COLOR_SOLID | bl.Z_BORDER_ALL)

		bl.OnMouseMove(func(e *bl.MouseMoveEvent) {
			//fmt.Println("BLACK")
		})

		bl.Div()
		{
			bl.Id("red")
			bl.Pos(160, 160)
			bl.Dim(360,360)
			bl.Color(.1,0,0)
			bl.BorderThickness([]int{1,1,1,1})
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()
			bl.Flag(bl.Z_COLOR_SOLID | bl.Z_BORDER_ALL)

			resize.On(func(interface{}) {
			     //fmt.Println("Resize")
			})

			bl.Div()
			{
				bl.Id("green")
				bl.Pos(80, 80)
				bl.Dim(180,180)
				bl.Color(.0,.1,0)
				bl.BorderThickness([]int{1,1,1,1})
				bl.BorderColor(1,1,1)
				bl.BorderTopsCanvas()
				bl.Flag(bl.Z_COLOR_SOLID | bl.Z_BORDER_ALL)
				//bl.OnMouseMove(func(e *bl.MouseMoveEvent) {
				//	fmt.Println("GREEN")
				//	e.BubbleToParent=false
				//})
			}
			bl.End()
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


