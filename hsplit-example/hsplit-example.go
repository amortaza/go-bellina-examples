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
	"github.com/amortaza/go-basic-widgets/hsplit"
)

func init_() {
	bl.Plugin( click.NewPlugin() )
	bl.Plugin( double_click.NewPlugin(1000) )
	bl.Plugin( mouse_drag.NewPlugin() )
	bl.Plugin( drag.NewPlugin() )
	bl.Plugin( resize.NewPlugin() )
	bl.Plugin( focus.NewPlugin() )
	bl.Plugin( zindex.NewPlugin() )
	bl.Plugin( animation.NewPlugin() )
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

		bl.BorderThickness([]int32{2,2,2,2})
		bl.BorderColor(1,1,1)

		bl.Div()
		{
			bl.ID("black")
			bl.Pos(60, 60)
			bl.Dim(360,360)
			bl.Color(0,0,0)
			bl.BorderThickness([]int32{1,1,1,1})
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()

			bl.Div()
			{
				bl.ID("green")
				bl.Pos(10, 10)
				bl.Dim(60,60)
				bl.Color(.0,.1,.0)
				bl.BorderThickness([]int32{1,1,1,1})
				bl.BorderColor(1,1,1)
				bl.BorderTopsCanvas()
			}
			bl.End()

			bl.Div()
			{
				bl.ID("blue")
				bl.Pos(160, 160)
				bl.Dim(60,60)
				bl.Color(0,0,.5)
				bl.BorderThickness([]int32{1,1,1,1})
				bl.BorderColor(1,1,1)
				bl.BorderTopsCanvas()
			}
			bl.End()

			hsplit.Use()
		}
		bl.End()

		bl.Div()
		{
			bl.ID("black2")
			bl.Pos(260, 260)
			bl.Dim(360,360)
			bl.Color(0,0,0)
			bl.BorderThickness([]int32{1,1,1,1})
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()

			bl.Div()
			{
				bl.ID("green2")
				bl.Pos(10, 10)
				bl.Dim(60,60)
				bl.Color(.0,.1,.0)
				bl.BorderThickness([]int32{1,1,1,1})
				bl.BorderColor(1,1,1)
				bl.BorderTopsCanvas()
			}
			bl.End()

			bl.Div()
			{
				bl.ID("blue2")
				bl.Pos(160, 160)
				bl.Dim(60,60)
				bl.Color(0,0,.5)
				bl.BorderThickness([]int32{1,1,1,1})
				bl.BorderColor(1,1,1)
				bl.BorderTopsCanvas()
			}
			bl.End()

			hsplit.Use()
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


