# gl10

## Abstract
This is OpenGL 1.0 binding for the programming language Go. It works on linux.

## Installation

	go get github.com/vitalibaumtrok/gl10
	go install github.com/vitalibaumtrok/gl10

To update the package type

	go get -u github.com/vitalibaumtrok/gl10

_Note: You need the programming language Go and git to run these commands._

## Development
Programming language Go is used.

## Example
You need GLUT binding for this example. To install it type

	go get github.com/vitalibaumtrok/glut
	go install github.com/vitalibaumtrok/glut

The following example code draws a line

	package main

	import (
		"math"
		"github.com/vitalibaumtrok/gl10"
		"github.com/vitalibaumtrok/glut"
	)

	var rad float64 = 0

	func main() {
		glut.Init()
		glut.InitDisplayMode(glut.SINGLE | glut.RGBA)
		glut.CreateWindow("Testing OpenGL 1.0 binding for Go");
		glut.ReshapeFunc(reshape)
		glut.DisplayFunc(display)
		glut.MainLoop()
	}

	func reshape(width, height int) {
		gl10.Viewport(0, 0, width, height)
		gl10.MatrixMode(gl10.PROJECTION)
		gl10.LoadIdentity()
	}

	func display() {
		sinrad := math.Sin(rad)
		cosrad := math.Cos(rad)

		gl10.ClearColor(1.0, 1.0, 1.0, 0.0)
		gl10.Clear(gl10.COLOR_BUFFER_BIT)
		gl10.Color3f(0.0, 0.0, 0.0);
		gl10.LineWidth(3.0);
		gl10.Begin(gl10.LINES);
			gl10.Vertex2d(sinrad, cosrad)
			gl10.Vertex2d(-sinrad, -cosrad)
		gl10.End();
		gl10.Flush();

		rad += 0.0001
		glut.PostRedisplay()
	}

## Coverage

	glClearColor, glClear, glBegin, glEnd, glFlush, glLineWidth, glMatrixMode,
	glLoadIdentity, glViewport, glMultMatrixd, glMultMatrixf, glBlendFunc,
	glEnable, glTexEnvf, glTexEnvi, glTexParameterf, glTexParameteri,
	glTexImage1D, glTexImage2D, glGetTexImage, glTexCoord, glVertex, glColor

## Copyright
Copyright 2014 Vitali Baumtrok

This program is distributed under the terms of the Boost Software License,
Version 1.0, as published on <http://www.boost.org/LICENSE_1_0.txt>.
