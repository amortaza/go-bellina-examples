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
	"github.com/amortaza/go-basic-widgets/simple/edit"
	"github.com/amortaza/go-bellina-plugins/zindex"
	"github.com/amortaza/go-basic-widgets/simple/button"
	"github.com/amortaza/go-bellina-plugins/hover"
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

			bl.On2("focus",
				func(v interface{}) {
					e := v.(focus.Event)
					fmt.Println("red: key " + e.ClickedNodeId)
				},
				func(v interface{}) {
					e := v.(focus.Event)
					fmt.Println("red: clicked on " + e.ClickedNodeId + " taken away from " + e.LoseFocusNodeId)

				}, func(v interface{}) {
					e := v.(focus.Event)
					fmt.Println("red: lose " + e.LoseFocusNodeId + " by ** " + e.ClickedNodeId)
				},
			)
		}
		bl.End()

		bl.Div()
		{
			bl.Id("green")
			bl.Pos(250, 120)
			bl.Dim(160,120)
			bl.Color(0,.1,.0)
			bl.BorderThickness([]int32{1,1,1,1})
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()

			bl.On2("focus",
				func(v interface{}) {
					e := v.(focus.Event)
					fmt.Println("green: key " + e.ClickedNodeId)
				},
				func(v interface{}) {
					e := v.(focus.Event)
					fmt.Println("green: clicked on " + e.ClickedNodeId + " taken away from " + e.LoseFocusNodeId)

				}, func(v interface{}) {
					e := v.(focus.Event)
					fmt.Println("green: lose " + e.LoseFocusNodeId + " by ** " + e.ClickedNodeId)
				},
			)
		}
		bl.End()

		bl.Div()
		{
			bl.Id("blue")
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


