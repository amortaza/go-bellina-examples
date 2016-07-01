package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina-plugins/click"
	"github.com/amortaza/go-bellina-plugins/double-click"
	"github.com/amortaza/go-bellina-plugins/mouse-drag"
	"github.com/amortaza/go-bellina-plugins/drag"
	"github.com/amortaza/go-bellina-plugins/resize"
	"github.com/amortaza/go-bellina-plugins/focus"
	"github.com/amortaza/go-bellina-plugins/zindex"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/animation"
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

		bl.BorderThickness([]int32{2,2,2,2})
		bl.BorderColor(1,1,1)

		bl.Div()
		{
			bl.Id("red")
			bl.Pos(60, 60)
			bl.Dim(160,120)
			bl.Color(.1,0,.0)
			bl.BorderThickness([]int32{1,1,1,1})
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()

			nodeId := bl.Current_Node.Id

			click.On(func(interface{}) {
				animation.StartPath(nodeId, "one", 100, 600, 60, "linear", func(shadow *bl.ShadowNode, value float32) {
					shadow.Left = int32( value)
					shadow.Top++
				})

				animation.StartPath(nodeId, "two", 1, .05, 50, "linear", func(shadow *bl.ShadowNode, value float32) {
					shadow.NodeOpacity = []float32{value,value,value,value}
				})
			})

			shadow := bl.EnsureShadow()
			bl.Current_Node.Left = shadow.Left
			bl.Current_Node.Top = shadow.Top
			bl.Current_Node.NodeOpacity = shadow.NodeOpacity
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


