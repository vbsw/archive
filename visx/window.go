
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package visx

import (
	"gitlab.com/vbit/vis"
)

type tWindow struct {

}

func (this *tWindow) Show ( ) {

}

func (this *tWindow) DimensionsInPixels ( ) (float32,float32) {
	return 0, 0
}

func (this *tWindow) SetDimensionsInPixels ( width, height float32 ) {

}

func (this *tWindow) Title ( ) string {
	return ""
}

func (this *tWindow) SetTitle ( title string ) {

}

func (this *tWindow) Destroy ( ) vis.IWindow {
	return nil
}

