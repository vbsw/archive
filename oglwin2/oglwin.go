
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


// OGLWin (OpenGL Window) provides functions to create windows with OpenGL context.
package oglwin

import (
	"github.com/vbsw/semver"
)

const (
	BUTTON1_PRESSED = iota
	BUTTON1_RELEASED
	BUTTON2_PRESSED
	BUTTON2_RELEASED
	BUTTON3_PRESSED
	BUTTON3_RELEASED
	CLOSE
	KEY_PRESSED
	KEY_RELEASED
	MOUSE
	REDRAW
	WHEEL_DOWN
	WHEEL_UP
)

var version semver.Version

// Current version is 0.1.0.
func Version ( ) semver.Version {
	if version == nil {
		semver.New(0, 1, 0)
	}
	return version
}

