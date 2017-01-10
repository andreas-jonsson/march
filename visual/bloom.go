// +-------------------=M=a=r=c=h=-=E=n=g=i=n=e=---------------------+
// | Copyright (C) 2016-2017 Andreas T Jonsson. All rights reserved. |
// | Contact <mail@andreasjonsson.se>                                |
// +-----------------------------------------------------------------+

package visual

import (
	"image"
	"unsafe"

	"github.com/goxjs/gl"
	"github.com/goxjs/gl/glutil"
)

type Bloom struct {
	programID      gl.Program
	textureID      gl.Texture
	positionAttrib gl.Attrib

	vertexBufferID,
	uvBufferID gl.Buffer

	size image.Point
}

func NewBloom(size image.Point) (*Bloom, error) {
	b := &Bloom{size: size}

	var err error
	b.programID, err = glutil.CreateProgram(bloomVertexShaderSrc, bloomFragmentShaderSrc)
	if err != nil {
		return nil, err
	}

	squareVerticesData := []float32{
		-1, -1,
		1, -1,
		-1, 1,
		1, 1,
	}

	b.vertexBufferID = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, b.vertexBufferID)

	ptr := unsafe.Pointer(&squareVerticesData[0])
	gl.BufferData(gl.ARRAY_BUFFER, (*[1 << 30]byte)(ptr)[:len(squareVerticesData)*4], gl.STATIC_DRAW)

	textureUVData := []float32{
		0, 0,
		1, 0,
		0, 1,
		1, 1,
	}

	b.uvBufferID = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, b.uvBufferID)

	ptr = unsafe.Pointer(&textureUVData[0])
	gl.BufferData(gl.ARRAY_BUFFER, (*[1 << 30]byte)(ptr)[:len(textureUVData)*4], gl.STATIC_DRAW)

	b.textureID = gl.CreateTexture()
	gl.BindTexture(gl.TEXTURE_2D, b.textureID)

	gl.TexImage2D(gl.TEXTURE_2D, 0, size.X, size.Y, gl.RGB, gl.UNSIGNED_BYTE, nil)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	return b, nil
}

func (b *Bloom) Destroy() {
	gl.DeleteTexture(b.textureID)
	gl.DeleteBuffer(b.vertexBufferID)
	gl.DeleteBuffer(b.uvBufferID)
	gl.DeleteProgram(b.programID)
}

func (b *Bloom) Render() {
	gl.Disable(gl.CULL_FACE)
	gl.Disable(gl.DEPTH_TEST)
	gl.DepthMask(false)

	gl.Disable(gl.BLEND)

	//gl.Enable(gl.BLEND)
	//gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, b.textureID)

	gl.CopyTexImage2D(gl.TEXTURE_2D, 0, gl.RGB, 0, 0, b.size.X, b.size.Y, 0)
	gl.GenerateMipmap(gl.TEXTURE_2D)

	gl.UseProgram(b.programID)

	pos := gl.GetAttribLocation(b.programID, "a_position")
	uv := gl.GetAttribLocation(b.programID, "a_uv")

	gl.BindBuffer(gl.ARRAY_BUFFER, b.vertexBufferID)
	gl.VertexAttribPointer(pos, 2, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(pos)

	gl.BindBuffer(gl.ARRAY_BUFFER, b.uvBufferID)
	gl.VertexAttribPointer(uv, 2, gl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(uv)

	samp := gl.GetUniformLocation(b.programID, "s_texture")
	gl.Uniform1i(samp, 0)

	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)
}

var bloomVertexShaderSrc = `
	#version 120

	attribute vec4 a_position;
	attribute vec4 a_uv;
	varying vec2 v_uv;

	void main()
	{
	    gl_Position = a_position;
	    v_uv = a_uv.xy;
	}
`

var bloomFragmentShaderSrc = `
	#version 120

	#define N 8

	uniform sampler2D s_texture;
	varying vec2 v_uv;

	void main()
	{
		vec3 col = texture2D(s_texture, v_uv, 0).xyz;
		for (int i = 1; i < N; i++)
			col += texture2D(s_texture, v_uv, i).xyz;

		gl_FragColor = vec4(col / (N - 1), 1);
	}
`
