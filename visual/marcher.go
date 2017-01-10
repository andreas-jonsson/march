// +-------------------=M=a=r=c=h=-=E=n=g=i=n=e=---------------------+
// | Copyright (C) 2016-2017 Andreas T Jonsson. All rights reserved. |
// | Contact <mail@andreasjonsson.se>                                |
// +-----------------------------------------------------------------+

package visual

import (
	"image"
	"log"
	"reflect"
	"unsafe"

	"github.com/goxjs/gl"
	"github.com/goxjs/gl/glutil"
)

type Marcher struct {
	programID      gl.Program
	vertexBufferID gl.Buffer
	positionAttrib gl.Attrib

	vertexBuffer []float32
}

func NewMarcher() (*Marcher, error) {
	m := new(Marcher)

	var err error
	m.programID, err = glutil.CreateProgram(vertexShaderSrc, fragmentShaderSrc)
	if err != nil {
		return nil, err
	}

	m.vertexBufferID = gl.CreateBuffer()
	m.positionAttrib = gl.GetAttribLocation(m.programID, "a_position")

	return m, nil
}

func (m *Marcher) Destroy() {
	gl.DeleteProgram(m.programID)
	gl.DeleteBuffer(m.vertexBufferID)
}

func (m *Marcher) BuildGeometry(img *image.Paletted) {
	if len(img.Palette) > 2 {
		log.Panic("Expected monocrome image to process!")
	}

	m.vertexBuffer = m.vertexBuffer[:0]

	size := img.Bounds().Size()
	size.X--
	size.Y--

	const space = 1.25

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			p0 := img.ColorIndexAt(x, y)
			p1 := img.ColorIndexAt(x+1, y)
			p2 := img.ColorIndexAt(x+1, y+1)
			p3 := img.ColorIndexAt(x, y+1)

			index := p3<<3 | p2<<2 | p1<<1 | p0
			vert := configurations[index]

			for i := 0; i < len(vert); i += 2 {
				m.vertexBuffer = append(m.vertexBuffer, vert[i]+float32(x)*space)
				m.vertexBuffer = append(m.vertexBuffer, vert[i+1]+float32(y)*space)
			}
		}
	}
}

func (m *Marcher) DebugConfigurations() {
	m.vertexBuffer = m.vertexBuffer[:0]
	for offset, vert := range configurations {
		for i := 0; i < len(vert); i += 2 {
			m.vertexBuffer = append(m.vertexBuffer, vert[i]+float32(offset)*1.2)
			m.vertexBuffer = append(m.vertexBuffer, vert[i+1])
		}
	}
}

func (m *Marcher) Render() {
	ln := len(m.vertexBuffer)
	if ln == 0 {
		return
	}

	gl.Disable(gl.CULL_FACE)
	gl.Disable(gl.BLEND)
	gl.Disable(gl.DEPTH_TEST)
	gl.DepthMask(false)

	gl.UseProgram(m.programID)

	header := *(*reflect.SliceHeader)(unsafe.Pointer(&m.vertexBuffer))
	header.Len *= 4
	header.Cap *= 4
	data := *(*[]byte)(unsafe.Pointer(&header))

	gl.BindBuffer(gl.ARRAY_BUFFER, m.vertexBufferID)
	gl.BufferData(gl.ARRAY_BUFFER, data, gl.STREAM_DRAW)

	gl.VertexAttribPointer(m.positionAttrib, 2, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(m.positionAttrib)

	gl.DrawArrays(gl.TRIANGLES, 0, ln/2)
}

var vertexShaderSrc = `
	#version 120

	attribute vec4 a_position;

	void main()
	{
		//const vec2 res = vec2(720, 450);
		const vec2 res = vec2(160, 90);
		vec2 halfRes = res * 0.5;

		vec2 p = a_position.xy * 1;

		vec2 pos = (vec2(p.x, res.y - p.y) / halfRes) - 1;
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

var configurations = [...][]float32{
	//case 0
	{},

	//case 1
	{
		0.0, 0.0,
		0.5, 0.0,
		0.0, 0.5,
	},

	//case 2
	{
		0.5, 0.0,
		1.0, 0.0,
		1.0, 0.5,
	},

	//case 3
	{
		0.0, 0.0,
		1.0, 0.0,
		0.0, 0.5,

		0.0, 0.5,
		1.0, 0.0,
		1.0, 0.5,
	},

	//case 4
	{
		1.0, 0.5,
		1.0, 1.0,
		0.5, 1.0,
	},

	//case 5
	{
		0.0, 0.0,
		0.5, 0.0,
		0.0, 0.5,

		0.0, 0.5,
		0.5, 0.0,
		0.5, 1.0,

		0.5, 1.0,
		0.5, 0.0,
		1.0, 0.5,

		1.0, 0.5,
		1.0, 1.0,
		0.5, 1.0,
	},

	//case 6
	{
		0.5, 0.0,
		1.0, 0.0,
		0.5, 1.0,

		0.5, 1.0,
		1.0, 0.0,
		1.0, 1.0,
	},

	//case 7
	{
		0.0, 0.0,
		1.0, 0.0,
		0.0, 0.5,

		0.0, 0.5,
		1.0, 0.0,
		1.0, 1.0,

		1.0, 1.0,
		0.5, 1.0,
		0.0, 0.5,
	},

	//case 8
	{
		0.0, 0.5,
		0.5, 1.0,
		0.0, 1.0,
	},

	//case 9
	{
		0.0, 0.0,
		0.5, 0.0,
		0.0, 1.0,

		0.0, 1.0,
		0.5, 0.0,
		0.5, 1.0,
	},

	//case 10
	{
		0.5, 0.0,
		1.0, 0.0,
		1.0, 0.5,

		1.0, 0.5,
		0.5, 1.0,
		0.5, 0.0,

		0.5, 0.0,
		0.5, 1.0,
		0.0, 0.5,

		0.0, 0.5,
		0.5, 1.0,
		0.0, 1.0,
	},

	//case 11
	{
		0.0, 0.0,
		1.0, 0.0,
		1.0, 0.5,

		1.0, 0.5,
		0.5, 1.0,
		0.0, 0.0,

		0.0, 0.0,
		0.5, 1.0,
		0.0, 1.0,
	},

	//case 12
	{
		0.0, 0.5,
		1.0, 0.5,
		1.0, 1.0,

		1.0, 1.0,
		0.0, 1.0,
		0.0, 0.5,
	},

	//case 13
	{
		0.0, 0.0,
		0.5, 0.0,
		1.0, 0.5,

		1.0, 0.5,
		1.0, 1.0,
		0.0, 1.0,

		0.0, 1.0,
		0.0, 0.0,
		1.0, 0.5,
	},

	//case 14
	{
		0.5, 0.0,
		1.0, 0.0,
		1.0, 1.0,

		1.0, 1.0,
		0.0, 1.0,
		0.5, 0.0,

		0.5, 0.0,
		0.0, 1.0,
		0.0, 0.5,
	},

	//case 15
	{
		0.0, 0.0,
		1.0, 0.0,
		1.0, 1.0,

		1.0, 1.0,
		0.0, 1.0,
		0.0, 0.0,
	},
}
