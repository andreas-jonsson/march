// +-------------------=M=a=r=c=h=-=E=n=g=i=n=e=---------------------+
// | Copyright (C) 2016-2017 Andreas T Jonsson. All rights reserved. |
// | Contact <mail@andreasjonsson.se>                                |
// +-----------------------------------------------------------------+

package visual

import (
	"image"
	"log"
)

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

func BuildGeometry(buf []float32, img *image.Paletted) []float32 {
	if len(img.Palette) > 2 {
		log.Panic("Expected monocrome image to process!")
	}

	size := img.Bounds().Size()
	size.X--
	size.Y--

	const space = 1.2

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			p0 := img.ColorIndexAt(x, y)
			p1 := img.ColorIndexAt(x+1, y)
			p2 := img.ColorIndexAt(x+1, y+1)
			p3 := img.ColorIndexAt(x, y+1)

			index := p3<<3 | p2<<2 | p1<<1 | p0
			vert := configurations[index]

			for i := 0; i < len(vert); i += 2 {
				buf = append(buf, vert[i]+float32(x)*space)
				buf = append(buf, vert[i+1]+float32(y)*space)
			}
		}
	}

	return buf
}

func DebugGeometry(buf []float32) []float32 {
	for offset, vert := range configurations {
		for i := 0; i < len(vert); i += 2 {
			buf = append(buf, vert[i]+float32(offset)*1.2)
			buf = append(buf, vert[i+1])
		}
	}
	return buf
}
