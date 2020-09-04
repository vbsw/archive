
//          Copyright 2017, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//     (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package window

// The event INITIALIZE is the first REDRAW event.
const (
	CLOSE int = iota
	RESIZE
	REDRAW
	KEY_PRESS
	KEY_AUTO_PRESS
	KEY_RELEASE
	BUTTON_PRESS
	BUTTON_RELEASE
	CURSOR_MOVE
	WHEEL_SPIN
	INITIALIZE
)

// The field TimeDelta contains the time in seconds since
// the last redraw event. If this is the first redraw event,
// then the time is zero.
type Event struct {
	EventCode    int
	KeyCode      int
	ButtonCode   int
	CursorX      int
	CursorY      int
	Wheel        int
	WindowWidth  int
	WindowHeight int
	TimeDelta    float32
	Time         float64
}

func NewEvent(eventCode int) *Event {
	event := new(Event)
	event.EventCode = eventCode

	return event
}
