
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package visx

import (
	"gitlab.com/vbit/vis"
)

type tTexture struct {

}

func (this *tTexture) DimensionsInPixels ( ) (float32,float32) {
	return 0, 0
}

func (this *tTexture) Destroy ( ) vis.ITexture {
	return nil
}

