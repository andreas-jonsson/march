// +build !mobile

// +-------------------=M=a=r=c=h=-=E=n=g=i=n=e=---------------------+
// | Copyright (C) 2016-2017 Andreas T Jonsson. All rights reserved. |
// | Contact <mail@andreasjonsson.se>                                |
// +-----------------------------------------------------------------+

package entry

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/andreas-jonsson/march/game"
	"github.com/andreas-jonsson/march/game/menu"
	"github.com/andreas-jonsson/march/game/play"
	"github.com/andreas-jonsson/march/platform"
)

func Entry() {
	if err := platform.Init(); err != nil {
		log.Panicln(err)
	}
	defer platform.Shutdown()

	//rnd, err := platform.NewRenderer(platform.ConfigWithFullscreen, platform.ConfigWithNoVSync)
	rnd, err := platform.NewRenderer(platform.ConfigWithDiv(2), platform.ConfigWithNoVSync) //, platform.ConfigWithDebug)
	if err != nil {
		log.Panicln(err)
	}
	defer rnd.Shutdown()
	platform.LogGLInfo()

	layerSize := image.Rect(0, 0, 160, 90)
	layers := []*image.Paletted{
		image.NewPaletted(layerSize, color.Palette{
			color.RGBA{0, 0, 0, 255},
			color.RGBA{255, 255, 255, 255},
		}),
	}

	states := map[string]game.GameState{
		"menu": menu.NewMenuState(layers),
		"play": play.NewPlayState(layers),
	}

	g, err := game.NewGame(states)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Shutdown()

	var gctl game.GameControl = g
	if err := g.SwitchState("menu", gctl); err != nil {
		log.Panicln(err)
	}

	for g.Running() {
		rnd.Clear()

		if err := g.Update(); err != nil {
			log.Panicln(err)
		}

		_, _, fps := g.Timing()
		rnd.SetWindowTitle(fmt.Sprintf("March - %d fps", fps))

		if err := g.Render(); err != nil {
			log.Panicln(err)
		}

		rnd.Present()
	}
}
