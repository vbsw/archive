# OGLWin

[![GoDoc](https://godoc.org/github.com/vbsw/oglwin?status.svg)](https://godoc.org/github.com/vbsw/oglwin)

## About
OGLWin (OpenGL Window) is a Go package that provides functions to create swindows with OpenGL context.

OGLWin is created by Vitali Baumtrok <vbsw@mailbox.org> and published at <https://github.com/vbsw/oglwin>.

## Copying
You can redistribute OGLWin and/or modify it under the terms of the Boost Software License, version 1.0.
(See accompanying file LICENSE or copy at <http://www.boost.org/LICENSE_1_0.txt>)

OGLWin is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the Boost Software License for more details.

## Dependencies
To compile OGLWin under Linux you need following packages (or similar)

	libc6-dev
	libX11-dev
	libgl1-mesa-dev

## Example
See <https://github.com/vbsw/oglwin-demo>.

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

## References

- <https://golang.org/doc/install>
- <https://git-scm.com/book/en/v2/Getting-Started-Installing-Git>

