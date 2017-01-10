// +-------------------=M=a=r=c=h=-=E=n=g=i=n=e=---------------------+
// | Copyright (C) 2016-2017 Andreas T Jonsson. All rights reserved. |
// | Contact <mail@andreasjonsson.se>                                |
// +-----------------------------------------------------------------+

package play

import (
	"image"
	"log"

	"image/png"

	"github.com/andreas-jonsson/march/data"
	"github.com/andreas-jonsson/march/game"
	"github.com/andreas-jonsson/march/platform"
	"github.com/andreas-jonsson/march/visual"
)

type playState struct {
	marcher *visual.Marcher
	bloom   *visual.Bloom

	testImage *image.Paletted
	layers    []*image.Paletted
}

func NewPlayState(layers []*image.Paletted) *playState {
	s := &playState{layers: layers}

	var err error
	s.marcher, err = visual.NewMarcher()
	if err != nil {
		log.Panicln(err)
	}

	s.bloom, err = visual.NewBloom(image.Pt(720, 450))
	if err != nil {
		log.Panicln(err)
	}

	r, err := data.FS.Open("test.png")
	if err != nil {
		log.Panicln(err)
	}
	defer r.Close()

	img, err := png.Decode(r)
	if err != nil {
		log.Panicln(err)
	}

	s.testImage = img.(*image.Paletted)

	return s
}

func (s *playState) Name() string {
	return "play"
}

func (s *playState) Enter(from game.GameState, args ...interface{}) error {
	return nil
}

func (s *playState) Exit(to game.GameState) error {
	return nil
}

func (s *playState) Update(gctl game.GameControl) error {
	for event := gctl.PollEvent(); event != nil; event = gctl.PollEvent() {
		switch event.(type) {
		case *platform.MouseButtonEvent, *platform.QuitEvent:
			gctl.Terminate()
		}
	}

	s.marcher.BuildGeometry(s.testImage)
	return nil
}

func (s *playState) Render() error {
	s.marcher.Render()
	s.bloom.Render()
	return nil
}
