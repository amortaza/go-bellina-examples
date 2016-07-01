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
	"github.com/amortaza/go-basic-widgets/edit"
	"github.com/amortaza/go-basic-widgets/button"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/layout/docker"
	"github.com/amortaza/go-bellina-plugins/layout/horiz"
	"github.com/amortaza/go-bellina-plugins/layout/vert"
	"github.com/amortaza/go-bellina-plugins/hover"
)

func init_() {
}

func tick() {

	bl.Root()
	{
		bl.Pos(64,64)
		bl.Dim(800,600)
		bl.Color(.3,.5,.5)
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
			bl.Dim(164,148)
			bl.Color(.1,0,.0)
			bl.BorderThickness([]int32{1,1,1,1})
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()

			vert.SetPercent(25)
		}
		bl.End()

		bl.Div()
		{
			bl.Id("green")
			bl.Pos(60, 60)
			bl.Dim(164,148)
			bl.Color(.0,.10,.0)
			bl.BorderThickness([]int32{1,1,1,1})
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()

			vert.FillRemaining()

			resize.Use()
		}
		bl.End()

		vert.SetSpacing(20)
		vert.Use()
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


