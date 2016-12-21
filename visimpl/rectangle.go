
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package visimpl

import (
	"github.com/vbsw/vis"
)

type tRectangle struct {
	x, y, z, width, height float32
	graphics *tGraphics
}

func (this *tRectangle) X ( ) float32 {
	return this.x
}

func (this *tRectangle) Y ( ) float32 {
	return this.y
}

func (this *tRectangle) Z ( ) float32 {
	return this.z
}

func (this *tRectangle) SetXYZ ( x, y, z float32 ) {
	this.x = x
	this.y = y
	this.z = z
}

func (this *tRectangle) Width ( ) float32 {
	return this.width
}

func (this *tRectangle) Height ( ) float32 {
	return this.height
}

func (this *tRectangle) SetWidthHeight ( width, height float32 ) {
	this.width = width
	this.height = height
}

func (this *tRectangle) NewTextureFill ( texture vis.Texture ) vis.RectangleTextureFill {
	fill := new(tRectangleTextureFill)
	fill.rectangle = this
	return fill
}

func (this *tRectangle) NewColorFill ( r, g, b, a float32 ) vis.ColorFill {
	fill := new(tColorFill)
	fill.r = r
	fill.g = g
	fill.b = b
	fill.a = a
	fill.rectangle = this
	return fill
}

func (this *tRectangle) Destroy ( ) vis.Rectangle {
	this.graphics = nil
	return nil
}

