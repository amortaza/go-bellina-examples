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
	"github.com/amortaza/go-bellina-plugins/hover"
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

		bl.BorderThickness([]int32{2,2,2,2})
		bl.BorderColor(1,1,1)

		bl.Div()
		{
			bl.Id("black2")
			bl.Pos(160, 160)
			bl.Dim(360,360)
			bl.Color(0,0,0)
			bl.BorderThickness([]int32{1,1,1,1})
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()

			bl.Font("arial", 12)
			bl.FontColor(1,1,1)
			bl.FontNudge(3,3)
			bl.Label("Hello world")

			hover.On(func(v interface{}) {
				e := v.(*hover.Event)

				if e.IsInEvent {
					bl.EnsureShadowById(e.InNodeId).Color1(.51, .51, 0).FontColor(0,0,0)

				} else {
					bl.EnsureShadowById(e.OutNodeId).Color1(.51, 0, .51).FontColor(.50,0,0)
				}
			} )

			bl.EnsureShadow().Color1_to_Node().FontColor_to_Node()
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


