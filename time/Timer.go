
//     Copyright 2016 Vitali Baumtrok (vbsw@maibox.org)
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package time

type Timer struct {
	nanosStart   int64
	nanosPaused  int64
	nanosDelay   int64
	nanosDelta   int64
	nanosTotal   int64
	microsDelta  int64
	microsTotal  int64
	millisDelta  int64
	millisTotal  int64
	secondsDelta int64
	secondsTotal int64
}

// Creates a new Timer and resets it.
// The start time is set to MonotonicNanos().
func NewTimer() *Timer {
	timer := new(Timer)
	timer.nanosStart = MonotonicNanos()
	return timer
}

// Resets current and delta time to zero.
// The start time is set to MonotonicNanos().
func (this *Timer) Reset() {
	this.nanosPaused = 0
	this.nanosDelay = 0
	this.nanosDelta = 0
	this.nanosTotal = 0
	this.microsDelta = 0
	this.microsTotal = 0
	this.millisDelta = 0
	this.millisTotal = 0
	this.secondsDelta = 0
	this.secondsTotal = 0
	this.nanosStart = MonotonicNanos()
}

// Pauses the timer. Update() can't be called until Continue().
func (this *Timer) Pause() {
	if this.nanosPaused == 0 {
		this.nanosPaused = MonotonicNanos()
	}
}

// The timer is no more paused.
func (this *Timer) Continue() {
	if this.nanosPaused > 0 {
		nanoContinue := MonotonicNanos()
		this.nanosDelay += nanoContinue - this.nanosPaused
		this.nanosPaused = 0
	}
}

// Updates current and delta time (with MonotonicNanos() as the reference time).
func (this *Timer) Update() {
	if this.nanosPaused == 0 {
		nanosTotalNew := MonotonicNanos() - this.nanosStart - this.nanosDelay
		microsTotalNew := nanosTotalNew / (1000)
		millisTotalNew := nanosTotalNew / (1000 * 1000)
		secondsTotalNew := nanosTotalNew / (1000 * 1000 * 1000)
		this.nanosDelta = nanosTotalNew - this.nanosTotal
		this.nanosTotal = nanosTotalNew
		this.microsDelta = microsTotalNew - this.microsTotal
		this.microsTotal = microsTotalNew
		this.millisDelta = millisTotalNew - this.millisTotal
		this.millisTotal = millisTotalNew
		this.secondsDelta = secondsTotalNew - this.secondsTotal
		this.secondsTotal = secondsTotalNew
	}
}

// Returns the time in nanoseconds between last two Update() calls
// (or between Reset() and Update() or the time set by Reset()).
func (this *Timer) NanosDelta() int64 {
	return this.nanosDelta
}

// Returns the time in nanoseconds between the Reset() and Update() call (or the time set by Reset()).
func (this *Timer) NanosTotal() int64 {
	return this.nanosTotal
}

// Returns the time in microseconds between last two Update() calls
// (or between Reset() and Update() or the time set by Reset()).
// CAUTION: This is an accumulated value and can be greater then the delta
// returned by NanosDelta(), i.e. MicrosDelta()*1000 > NanosDelta().
func (this *Timer) MicrosDelta() int64 {
	return this.microsDelta
}

// Returns the time in milliseconds between the Reset() and Update() call (or the time set by Reset()).
func (this *Timer) MicrosTotal() int64 {
	return this.microsTotal
}

// Returns the time in milliseconds between last two Update() calls
// (or between Reset() and Update() or the time set by Reset()).
// CAUTION: This is an accumulated value and can be greater then the delta
// returned by NanosDelta(), i.e. MillisDelta()*1000000 > NanosDelta().
func (this *Timer) MillisDelta() int64 {
	return this.millisDelta
}

// Returns the time in milliseconds between the Reset() and Update() call (or the time set by Reset()).
func (this *Timer) MillisTotal() int64 {
	return this.millisTotal
}

// Returns the time in seconds between last two Update() calls
// (or between Reset() and Update() or the time set by Reset()).
// CAUTION: This is an accumulated value and can be greater then the delta
// returned by NanosDelta(), i.e. SecondsDelta()*1000000000 > NanosDelta().
func (this *Timer) SecondsDelta() int64 {
	return this.secondsDelta
}

// Returns the time in seconds between the Reset() and Update() call (or the time set by Reset()).
func (this *Timer) SecondsTotal() int64 {
	return this.secondsTotal
}
