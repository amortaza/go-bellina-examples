package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-basic-widgets/vscroll"
	"github.com/amortaza/go-basic-widgets/hscroll"
	"github.com/amortaza/go-bellina-plugins/layout/docker"
	"strconv"
)

func init_() {
}
var pcts = ""

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

		bl.BorderThickness(bl.FourTwosInt)
		bl.BorderColor(1,1,1)

		bl.Div()
		{
			bl.Id("black")
			bl.Pos(60, 60)
			bl.Dim(640,480)
			bl.Color(0,0,0)
			bl.BorderThickness(bl.FourOnesInt)
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()

			bl.Font("arial", 78)
			bl.FontColor(.51,.51,0)
			bl.FontNudge(50,50)
			bl.Label(pcts)

			hscroll.Id("one").Thickness(40).On(func(v interface{}) {
				e := v.(*hscroll.Event)

				//fmt.Println(e.PercentStart, e.PercentEnd)
				pcts = strconv.Itoa(int(e.PercentStart * 100))
			})


			hscroll.Div()
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

			vscroll.Id("two").Thickness(40).On(func(v interface{}) {
				e := v.(*vscroll.Event)

				//fmt.Println(e.PercentStart, e.PercentEnd)
				pcts = strconv.Itoa(int(e.PercentStart * 100))
			})

			vscroll.Div()
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


