
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package visimpl

import (
	"github.com/vbsw/vis"
)

type tRectangleTextureFill struct {
	x, y, width, height float32
	rectangle *tRectangle
}

func (this *tRectangleTextureFill) X ( ) float32 {
	return this.x
}

func (this *tRectangleTextureFill) Y ( ) float32 {
	return this.y
}

func (this *tRectangleTextureFill) SetXY ( x, y float32 ) {
	this.x = x
	this.y = y
}

func (this *tRectangleTextureFill) Width ( ) float32 {
	return this.width
}

func (this *tRectangleTextureFill) Height ( ) float32 {
	return this.height
}

func (this *tRectangleTextureFill) SetWidthHeight ( width, height float32 ) {
	this.width = width
	this.height = height
}

func (this *tRectangleTextureFill) Destroy ( ) vis.RectangleTextureFill {
	this.rectangle = nil
	return nil
}

