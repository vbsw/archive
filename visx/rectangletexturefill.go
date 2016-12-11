
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package visx

import (
	"gitlab.com/vbit/vis"
)

type tRectangleTextureFill struct {

}

func (this *tRectangleTextureFill) Location ( ) (float32,float32) {
	return 0, 0
}

func (this *tRectangleTextureFill) SetLocation ( x, y float32 ) {

}

func (this *tRectangleTextureFill) Dimensions ( ) (float32,float32) {
	return 0, 0
}

func (this *tRectangleTextureFill) SetDimensions ( width, height float32 ) {

}

func (this *tRectangleTextureFill) Destroy ( ) vis.IRectangleTextureFill {
	return nil
}

