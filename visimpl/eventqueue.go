
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package visimpl

import (
	"github.com/vbsw/events"
)

type tEventQueue struct {
	polling, autoRedraw bool
	event events.Event
	keyCode int
}

func newEventQueue ( ) *tEventQueue {
	eventQueue := new(tEventQueue)
	eventQueue.polling = true
	eventQueue.event = events.NONE
	eventQueue.keyCode = 0
	return eventQueue
}

func (this *tEventQueue) PollEvents ( ) bool {
	return this.polling
}

func (this *tEventQueue) EventCode ( ) events.Event {
	return this.event
}

func (this *tEventQueue) KeyCode ( ) int {
	return this.keyCode
}

func (this *tEventQueue) EnableAutoRedraw ( ) {
	this.autoRedraw = true
}

func (this *tEventQueue) DisableAutoRedraw ( ) {
	this.autoRedraw = false
}

