
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


// Implementation of Visual Interface System ([VIS](https://github.com/vbsw/vis)). It's based on X and OpenGL.
package visimpl

import (
	"github.com/vbsw/vis"
)

type tVisualSystem struct {
	eventQueue *tEventQueue
	graphics *tGraphics
	window *tWindow
}

func New ( ) vis.VisualSystem {
	vsys := new(tVisualSystem)
	vsys.eventQueue = newEventQueue()
	vsys.graphics = newGraphics()
	vsys.window = newWindow()
	return vsys
}

func (this *tVisualSystem) EventQueue ( ) vis.EventQueue {
	return this.eventQueue
}

func (this *tVisualSystem) Graphics ( ) vis.Graphics {
	return this.graphics
}

func (this *tVisualSystem) Window ( ) vis.Window {
	return this.window
}

func (this *tVisualSystem) Destroy ( ) vis.VisualSystem {
	this.eventQueue.polling = false
	this.eventQueue = nil
	this.graphics = nil
	this.window = nil
	return nil
}

