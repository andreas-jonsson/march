// +-------------------=M=a=r=c=h=-=E=n=g=i=n=e=---------------------+
// | Copyright (C) 2016-2017 Andreas T Jonsson. All rights reserved. |
// | Contact <mail@andreasjonsson.se>                                |
// +-----------------------------------------------------------------+

package menu

import (
	"image"

	"github.com/andreas-jonsson/march/game"
)

type menuState struct {
	layers []*image.Paletted
}

func NewMenuState(layers []*image.Paletted) *menuState {
	return &menuState{layers}
}

func (s *menuState) Name() string {
	return "menu"
}

func (s *menuState) Enter(from game.GameState, args ...interface{}) error {
	return args[0].(game.GameControl).SwitchState("play", args[0])
}

func (s *menuState) Exit(to game.GameState) error {
	return nil
}

func (s *menuState) Update(gctl game.GameControl) error {
	gctl.PollAll()
	return nil
}

func (s *menuState) Render() error {
	return nil
}
