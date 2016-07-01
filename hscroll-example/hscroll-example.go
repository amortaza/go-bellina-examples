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
	"github.com/amortaza/go-basic-widgets/vscroll"
	"github.com/amortaza/go-basic-widgets/hscroll"
	"github.com/amortaza/go-bellina-plugins/layout/docker"
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

		bl.Font("arial", 8)
		bl.FontColor(1,1,1)
		bl.FontNudge(3,3)
		bl.Label("Hello world")

		bl.BorderThickness([]int32{2,2,2,2})
		bl.BorderColor(1,1,1)

		bl.Div()
		{
			bl.Id("black")
			bl.Pos(60, 60)
			bl.Dim(640,480)
			bl.Color(0,0,0)
			bl.BorderThickness([]int32{1,1,1,1})
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()

			hscroll.SetThickness(40)
			hscroll.On(func(v interface{}) {
				//e := v.(*hscroll.Event)

				//fmt.Println(e.PercentStart, e.PercentEnd)
			})


			hscroll.Div("one")
			{
				if bl.Current_Node == nil {
					fmt.Println("we are nil")
				}
				docker.AnchorBottom()
				docker.AnchorLeft()
				docker.AnchorRight()
				docker.Use()

				bl.Color(.1,0,0)
			}
			hscroll.End()

			vscroll.SetThickness(40)
			vscroll.On(func(v interface{}) {
				e := v.(*vscroll.Event)

				fmt.Println(e.PercentStart, e.PercentEnd)
			})

			vscroll.Div("two")
			{
				docker.AnchorRight()
				docker.AnchorTop()
				docker.AnchorBottom()
				docker.Use()

				bl.Color(.1,0,0.1)
			}
			vscroll.End()
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


