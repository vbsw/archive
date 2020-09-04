# Window

[![GoDoc](https://godoc.org/github.com/vbsw/window?status.svg)](https://godoc.org/github.com/vbsw/window)

## About
Window is a Go package that provides functions to create a window with an OpenGL context. Window is published at <https://github.com/vbsw/window>.

The goal of this project is to provide a simple way to create a window with OpenGL context for testing purposes.

## Copyright
Copyright 2017, Vitali Baumtrok (vbsw@mailbox.org).

Window is distributed under the Boost Software License, version 1.0. (See accompanying file LICENSE or copy at <http://www.boost.org/LICENSE_1_0.txt>)

Window is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the Boost Software License for more details.

## Example

	package main
	import (
		"github.com/vbsw/window"
	)

	func main() {
		properties := window.NewProperties()
		properties.Width = 480
		properties.Height = 640
		properties.Title = "Hello, world!"

		window.Create(properties, mainLoop)
	}

	func mainLoop(event *window.Event) {
		if event.EventCode == window.CLOSE {
			window.Destroy()
		}
	}

## Compiling On Windows
This package uses GLFW. GLFW contains C files and needs Cgo to be compiled. Cgo needs another compiler to compile C files. On a 64 bit system a 64 bit compiler is needed. Go to <http://tdm-gcc.tdragon.net> and download 64 bit gcc binaries. Install it. Compile with:

	$ go install github.com/vbsw/window

Dependencies of this package:

	github.com/vbsw/semver
	github.com/vbsw/go-gl/glfw/v3.2/glfw

## Using Go
Get this project:

	$ go get github.com/vbsw/window

Update a local copy:

	$ go get -u github.com/vbsw/window

Compile:

	$ go install github.com/vbsw/window

Run tests:

	$ go test github.com/vbsw/window

## Using Git
Get the master branch and all refs of this project:

	$ git clone https://github.com/vbsw/window.git

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

## References

- <https://golang.org/doc/install>
- <https://git-scm.com/book/en/v2/Getting-Started-Installing-Git>

