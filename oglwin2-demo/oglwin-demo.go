
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


// This is a demo for the Go package oglwin (<https://github.com/vbsw/oglwin>).
package main

import (
	"fmt"
	"github.com/vbsw/oglwin"
)

func main ( ) {
	window, _ := oglwin.New(true, 0)
	window.SetWH(640, 360)
	window.Show()

	for window.NextEvent() {

		switch window.EventCode() {

		case oglwin.KEY_PRESSED:
			fmt.Println("Key", window.KeyCode())
			if ( window.KeyCode() == 9 ) {
				window.Destroy()
			}

		case oglwin.REDRAW:
			fmt.Println("resize ", window.W(), "x", window.H())
			window.SwapBuffers()
		}
	}
}

