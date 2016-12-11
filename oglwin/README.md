# oglwin

[![GoDoc](https://godoc.org/gitlab.com/vbit/oglwin?status.svg)](https://godoc.org/gitlab.com/vbit/oglwin)

## About
oglwin is a Go package that provides functions to create a graphical window with OpenGL context.

oglwin is created by Vitali Baumtrok <vbit@mailbox.org> and published at <https://gitlab.com/vbit/oglwin>.

## Copying
oglwin is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

oglwin is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with oglwin. If not, see <http://www.gnu.org/licenses/>.

## Example

	package main

	import (
		"fmt"
		"gitlab.com/vbit/oglwin"
	)

	func main() {
		oglwin.SetInitFunc(windowInit)
		oglwin.SetKeyPressFunc(windowKeyPress)
		oglwin.SetCanvasRedrawFunc(windowCanvasRedraw)
		oglwin.SetWindowCloseFunc(windowClose)
		oglwin.Show()
	}

	func windowInit() {
		oglwin.SetCanvasSize(480,360)
	}

	func windowKeyPress(keyCode int) {
		fmt.Println("Key:",keyCode)
	}

	func windowCanvasRedraw(canvasWidth, canvasHeight int) {
		fmt.Println("Canvas Size:",canvasWidth,canvasHeight)
	}

	func windowClose() bool {
		return true
	}

## Dependencies

Packages

	libc6-dev
	libX11-dev
	libgl1-mesa-dev

## Using Go
How to install Go: <https://golang.org/doc/install>

Get this project:

	$ go get gitlab.com/vbg/oglwin

Update a local copy:

	$ go get -u gitlab.com/vbg/oglwin

Compile:

	$ go install gitlab.com/vbg/oglwin

## Using Git
How to install Git: <https://git-scm.com/book/en/v2/Getting-Started-Installing-Git>

Get the master branch and all refs of this project:

	$ git clone https://gitlab.com/vbit/oglwin.git
	$ cd window

See all tags:

	$ git tag -l

See local and remote branches:

	$ git branch -a

Checkout other branches than master, for example the development branch:

	$ git branch development origin/development
	$ git checkout development

See tracked remote branches:

	$ git branch -vv

Update all tracked branches and all refs:

	$ git fetch

