// +-------------------=M=a=r=c=h=-=E=n=g=i=n=e=---------------------+
// | Copyright (C) 2016-2017 Andreas T Jonsson. All rights reserved. |
// | Contact <mail@andreasjonsson.se>                                |
// +-----------------------------------------------------------------+

package play

import "github.com/andreas-jonsson/march/game"

type playState struct {
}

func NewPlayState() *playState {
	return &playState{}
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

var anim = 0.0

func (s *playState) Update(gctl game.GameControl) error {
	return nil
}

func (s *playState) Render() error {
	return nil
}
