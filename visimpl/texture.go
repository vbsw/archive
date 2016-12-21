
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package visimpl

import (
	"github.com/vbsw/vis"
)

type tTexture struct {
	widthInPixels, heightInPixels int
	graphics *tGraphics
}

func (this *tTexture) WidthInPixels ( ) int {
	return this.widthInPixels
}

func (this *tTexture) HeightInPixels ( ) int {
	return this.heightInPixels
}

func (this *tTexture) Destroy ( ) vis.Texture {
	this.graphics = nil
	return nil
}

