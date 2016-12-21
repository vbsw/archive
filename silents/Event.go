
//          Copyright 2016 Vitali Baumtrok
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package main

import (
	"github.com/vbsw/xlib"
)

type tEvent struct {
	xevent xlib.XEvent
}

func (this *tEvent) KeyPressed ( ) bool {
	return this.xevent.Type() == xlib.KeyPress
}

func (this *tEvent) KeyReleased ( ) bool {
	return this.xevent.Type() == xlib.KeyRelease
}
