
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package visimpl

import (
	"github.com/vbsw/vis"
)

type tColorFill struct {
	r, g, b, a float32
	rectangle *tRectangle
}

func (this *tColorFill) R ( ) float32 {
	return this.r
}

func (this *tColorFill) G ( ) float32 {
	return this.g
}

func (this *tColorFill) B ( ) float32 {
	return this.b
}

func (this *tColorFill) A ( ) float32 {
	return this.a
}

func (this *tColorFill) SetRGBA ( r, g, b, a float32 ) {
	this.r = r
	this.g = g
	this.b = b
	this.a = a
}

func (this *tColorFill) Destroy ( ) vis.ColorFill {
	return nil
}

