
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package visimpl

import (
	"github.com/vbsw/vis"
)

type tGraphics struct {

}

func newGraphics ( ) *tGraphics {
	graphics := new(tGraphics)
	return graphics
}

func (this *tGraphics) NewRectangle ( ) vis.Rectangle {
	rect := new(tRectangle)
	rect.graphics = this
	return rect
}

func (this *tGraphics) NewTexture ( rgba []byte ) vis.Texture {
	texture := new(tTexture)
	texture.graphics = this
	return texture
}

func (this *tGraphics) SwapBuffers ( ) {

}

