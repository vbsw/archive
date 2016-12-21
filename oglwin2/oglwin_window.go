
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package oglwin

type Window struct {
	width, height int
	eventPolling bool
}

// Creates and returns a new window. Only 2, 4, 8 and 16 are accepted as MSAA samples, all other values are ignored.
func New ( doubleBuffer bool, msaaSamples int ) (*Window, error) {
	var config tConfiguration
	config.init(doubleBuffer, msaaSamples)
	window, err := newFromConfiguration(&config)
	window.eventPolling = (err == nil)
	return window, err
}

func newFromConfiguration ( config *tConfiguration ) (*Window, error) {
	window := new(Window)
	return window, nil
}

// Sets the width and height of the window.
func (this *Window) SetWH ( width, height int ) {
	this.width = width
	this.height = height
}

// Returns the width of the window.
func (this *Window) W ( ) int {
	return 0
}

// Returns the height of the window.
func (this *Window) H ( ) int {
	return 0
}

// Returns the x coordinate of the mouse cursor.
func (this *Window) MouseX ( ) int {
	return 0
}

// Returns the y coordinate of the mouse cursor.
func (this *Window) MouseY ( ) int {
	return 0
}

// Sets the window visible.
func (this *Window) Show ( ) {
}

// Polls the next event and returns true, if next event is available.
// Returns false if window has been destroyed.
func (this *Window) NextEvent ( ) bool {
	return this.eventPolling
}

// Returns the event code of the current event.
func (this *Window) EventCode ( ) int {
	return 0
}

// Returns the key code of the last key event.
func (this *Window) KeyCode ( ) int {
	return 0
}

// Returns spin range of the wheel .
func (this *Window) WheelValue ( ) int {
	return 1
}

// Swaps buffers, if possible.
func (this *Window) SwapBuffers ( ) {
}

// Destroys this window.
// The function NextEvent returns always false, now.
func (this *Window) Destroy ( ) {
	this.eventPolling = false
}

