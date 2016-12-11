
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package visx

import (
	"gitlab.com/vbit/vis"
)

type tGraphics struct {

}

func (this *tGraphics) NewRectangle ( ) vis.IRectangle {
	rect := new(tRectangle)
	return rect
}

func (this *tGraphics) NewTexture ( rgba []byte ) vis.ITexture {
	texture := new(tTexture)
	return texture
}

func (this *tGraphics) SwapBuffers ( ) {

}

func (this *tGraphics) Destroy ( ) vis.IGraphics {
	return nil
}

