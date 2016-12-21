
//          Copyright 2016 Vitali Baumtrok
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package main

import (
	"fmt"
	"github.com/vbsw/xlib"
)

type tWindow struct {
	width uint
	height uint
	display *xlib.Display
	window xlib.Window
	event *tEvent
}

func newWindow ( ) *tWindow {
	window := new(tWindow)
	window.width = app.settings.WindowWidth()
	window.height = app.settings.WindowHeight()
	window.display = xlib.XOpenDisplay()
	if window.display != nil {
		screen := xlib.XDefaultScreenOfDisplay(window.display)
		rootWindow := xlib.XRootWindowOfScreen(screen)
		window.window = xlib.XCreateSimpleWindow(window.display,rootWindow,0,0,window.width,window.height,0,0,0)
		window.event = new(tEvent)
		xlib.XSelectInput(window.display,window.window,xlib.KeyPressMask)
	} else {
		fmt.Println("could not connect to display server")
	}
	return window
}

func (this *tWindow) Show ( ) {
	xlib.XMapWindow(this.display,this.window)
}

func (this *tWindow) NextEvent ( ) bool {
	this.event.xevent = xlib.XNextEvent(this.display)
	return true
}

func (this *tWindow) Event ( ) *tEvent {
	return this.event
}

func (this *tWindow) Destroy ( ) {
	xlib.XCloseDisplay(this.display);
	this.display = nil
}

