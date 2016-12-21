
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package visimpl

type tWindow struct {
	widthInPixels, heightInPixels int
	title string
}

func newWindow ( ) *tWindow {
	window := new(tWindow)
	window.title = "Prototype Application"
	window.widthInPixels = 720
	window.heightInPixels = 480
	return window
}

func (this *tWindow) Show ( ) {

}

func (this *tWindow) WidthInPixels ( ) int {
	return this.widthInPixels
}

func (this *tWindow) HeightInPixels ( ) int {
	return this.heightInPixels
}

func (this *tWindow) SetWidthHeightInPixels ( width, height int ) {
	this.widthInPixels = width
	this.heightInPixels = height
}

func (this *tWindow) Title ( ) string {
	return this.title
}

func (this *tWindow) SetTitle ( title string ) {
	this.title = title
}

