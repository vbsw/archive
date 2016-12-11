
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


// VISX (Visual Interface System Implementation for X) provides implementation for VIS based on X and OpenGL. This is version 0.1.0.
package visx

import (
	"gitlab.com/vbit/semver"
	"gitlab.com/vbit/vis"
)

// Version is 0.1.0.
var Version semver.IVersion

func init ( ) {
    Version = semver.New(0,1,0)
}

func Create ( ) {
	if vis.Queue == nil && vis.Graphics == nil && vis.Window == nil {
		queue := new(tQueue)
		graphics := new(tGraphics)
		window := new(tWindow)
		vis.Queue = queue
		vis.Graphics = graphics
		vis.Window = window
	}
}

