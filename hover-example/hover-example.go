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
	"github.com/amortaza/go-bellina-plugins/hover"
	"github.com/amortaza/go-bellina"
)

func init_() {
	bl.Plugin( click.NewPlugin() )
	bl.Plugin( double_click.NewPlugin(1000) )
	bl.Plugin( mouse_drag.NewPlugin() )
	bl.Plugin( drag.NewPlugin() )
	bl.Plugin( resize.NewPlugin() )
	bl.Plugin( focus.NewPlugin() )
	bl.Plugin( zindex.NewPlugin() )
	bl.Plugin( hover.NewPlugin() )
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
			bl.ID("red")
			bl.Pos(60, 60)
			bl.Dim(160,120)
			bl.Color(.1,0,.0)
			bl.BorderThickness([]int32{1,1,1,1})
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()

			hover.On(func(v interface{}) {
				// e := v.(.Event)
				fmt.Println("hi")
			} )
		}
		bl.End()

		bl.Div()
		{
			bl.ID("green")
			bl.Pos(250, 120)
			bl.Dim(160,120)
			bl.Color(0,.1,.0)
			bl.BorderThickness([]int32{1,1,1,1})
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()
		}
		bl.End()

		bl.Div()
		{
			bl.ID("blue")
			bl.Pos(450, 120)
			bl.Dim(160,120)
			bl.Color(0,0,.1)
			bl.BorderThickness([]int32{1,1,1,1})
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()
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

