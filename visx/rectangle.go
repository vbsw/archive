
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package visx

import (
	"gitlab.com/vbit/vis"
)

type tRectangle struct {

}

func (this *tRectangle) Location ( ) (float32,float32,float32) {
	return 0, 0, 0
}

func (this *tRectangle) SetLocation ( x, y, z float32 ) {

}

func (this *tRectangle) Dimensions ( ) (float32,float32) {
	return 0, 0
}

func (this *tRectangle) SetDimensions ( width, height float32 ) {

}

func (this *tRectangle) NewTextureFill ( texture vis.ITexture ) vis.IRectangleTextureFill {
	fill := new(tRectangleTextureFill)
	return fill
}

func (this *tRectangle) NewColorFill ( r, g, b, a float32 ) vis.IColorFill {
	fill := new(tColorFill)
	fill.r = r
	fill.g = g
	fill.b = b
	fill.a = a
	return fill
}

func (this *tRectangle) Destroy ( ) vis.IRectangle {
	return nil
}

