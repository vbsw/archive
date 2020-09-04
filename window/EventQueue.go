
//          Copyright 2017, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//     (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package window

type tEventQueue struct {
	events       []*Event
	timePrevious float64
	timeDeltaSet bool
}

func newEventQueue() *tEventQueue {
	eventQueue = new(tEventQueue)
	eventQueue.events = make([]*Event, 0, 10)

	return eventQueue
}

func (this *tEventQueue) isEmpty() bool {
	return len(this.events) == 0
}

func (this *tEventQueue) appendEvent(event *Event) {
	this.events = append(this.events, event)
}

func (this *tEventQueue) appendRedrawEventIfMissing() {
	for _, event := range this.events {
		if event.EventCode == REDRAW || event.EventCode == INITIALIZE {
			return
		}
	}
	if this.timeDeltaSet {
		this.appendEvent(NewEvent(REDRAW))

	} else {
		this.appendEvent(NewEvent(INITIALIZE))
	}
}

func (this *tEventQueue) removeEventFront() {
	this.events = this.events[1:]
}

func (this *tEventQueue) insertEventFront(event *Event) {
	if len(this.events) < cap(this.events) {
		this.events = append(this.events, nil)
		copy(this.events[1:], this.events[:])
		this.events[0] = event

	} else {
		newEvents := make([]*Event, 1, cap(this.events)+2)
		newEvents[0] = event
		this.events = append(newEvents, this.events...)
	}
}

func (this *tEventQueue) indexOfEvent(eventCode int) int {
	for i, event := range this.events {
		if event.EventCode == eventCode {
			return i
		}
	}
	return -1
}

func (this *tEventQueue) replaceEvent(eventCode int) {
	resizeEventIndex := this.indexOfEvent(RESIZE)

	if resizeEventIndex < 0 {
		event := NewEvent(RESIZE)
		this.events = append(this.events, event)

	} else if resizeEventIndex+1 < len(this.events) {
		event := this.events[resizeEventIndex]
		copy(this.events[resizeEventIndex:], this.events[resizeEventIndex+1:])
		this.events[len(this.events)-1] = event
	}
}

func (this *tEventQueue) updateEventTime(time float64) {
	if this.timeDeltaSet {
		this.events[0].Time = time
		newTimeDelta := (this.events[0].Time - this.timePrevious) * 1000 * 1000
		this.events[0].TimeDelta = float32(int(newTimeDelta)) / (1000 * 1000)

		if this.events[0].EventCode == REDRAW {
			this.timePrevious = this.events[0].Time
		}

	} else {
		this.events[0].Time = time
		this.events[0].TimeDelta = 0

		if this.events[0].EventCode == INITIALIZE {
			this.timePrevious = this.events[0].Time
			this.timeDeltaSet = true
		}
	}
}

func (this *tEventQueue) isTimeInitialized() bool {
	initialized := this.events[0].EventCode == INITIALIZE || this.timeDeltaSet

	return initialized
}
