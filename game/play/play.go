// +-------------------=M=a=r=c=h=-=E=n=g=i=n=e=---------------------+
// | Copyright (C) 2016-2017 Andreas T Jonsson. All rights reserved. |
// | Contact <mail@andreasjonsson.se>                                |
// +-----------------------------------------------------------------+

package play

import (
	"image"
	"log"
	"reflect"
	"unsafe"

	"image/png"

	"github.com/andreas-jonsson/march/game"
	"github.com/andreas-jonsson/march/visual"
	"github.com/andreas-jonsson/openwar/data"
	"github.com/goxjs/gl"
	"github.com/goxjs/gl/glutil"
)

type playState struct {
	programID      gl.Program
	vertexBufferID gl.Buffer
	positionAttrib gl.Attrib

	testImage *image.Paletted
}

func NewPlayState() *playState {
	s := &playState{}

	var err error
	s.programID, err = glutil.CreateProgram(vertexShaderSrc, fragmentShaderSrc)
	if err != nil {
		//return nil, err
		panic(err)
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

	s.vertexBufferID = gl.CreateBuffer()
	s.positionAttrib = gl.GetAttribLocation(s.programID, "a_position")

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
	return nil
}

func (s *playState) Render() error {
	//buf := visual.DebugGeometry(*new([]float32))

	buf := visual.BuildGeometry(*new([]float32), s.testImage)

	gl.Disable(gl.CULL_FACE)
	gl.Disable(gl.BLEND)
	gl.Disable(gl.DEPTH_TEST)
	gl.DepthMask(false)

	gl.UseProgram(s.programID)

	gl.ClearColor(0.5, 0.5, 0.5, 1)
	//gl.Viewport(100, 100, 160, 90)

	if len(buf) > 0 {
		header := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
		header.Len *= 4
		header.Cap *= 4
		data := *(*[]byte)(unsafe.Pointer(&header))

		gl.BindBuffer(gl.ARRAY_BUFFER, s.vertexBufferID)
		gl.BufferData(gl.ARRAY_BUFFER, data, gl.STREAM_DRAW)

		gl.VertexAttribPointer(s.positionAttrib, 2, gl.FLOAT, false, 0, 0)
		gl.EnableVertexAttribArray(s.positionAttrib)

		gl.DrawArrays(gl.TRIANGLES, 0, len(buf)/2)
	}

	return nil
}

var vertexShaderSrc = `
	#version 120

	attribute vec4 a_position;

	void main()
	{
		//const vec2 res = vec2(720, 450);
		const vec2 res = vec2(160, 90);
		vec2 halfRes = res * 0.5;

		vec2 pos = (vec2(a_position.x, res.y - a_position.y) / halfRes) - 1;
		gl_Position = vec4(pos, 0, 1);
	}
`

var fragmentShaderSrc = `
	#version 120

	//uniform vec3 u_color;

	void main()
	{
		gl_FragColor = vec4(1,0,0,1);
	}
`
