
//     Copyright 2016 Vitali Baumtrok (vbsw@maibox.org)
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package time

type SecondsTimer struct {
	nanosStart   int64
	nanosPaused  int64
	nanosDelay   int64
	nanosTotal   int64
	secondsDelta float64
	secondsTotal float64
}

// Creates a new SecondsTimer and resets it.
// The start time is set to MonotonicNanos().
func NewSecondsTimer() *SecondsTimer {
	timer := new(SecondsTimer)
	timer.nanosStart = MonotonicNanos()
	return timer
}

// Resets current and delta time to zero.
// The start time is set to MonotonicNanos().
func (this *SecondsTimer) Reset() {
	this.nanosPaused = 0
	this.nanosDelay = 0
	this.nanosTotal = 0
	this.secondsDelta = 0
	this.secondsTotal = 0
	this.nanosStart = MonotonicNanos()
}

// Pauses the timer. Update() can't be called until Continue().
func (this *SecondsTimer) Pause() {
	if this.nanosPaused == 0 {
		this.nanosPaused = MonotonicNanos()
	}
}

// The timer is no more paused.
func (this *SecondsTimer) Continue() {
	if this.nanosPaused > 0 {
		nanoContinue := MonotonicNanos()
		this.nanosDelay += nanoContinue - this.nanosPaused
		this.nanosPaused = 0
	}
}

// Updates current and delta time (with MonotonicNanos() as the reference time).
func (this *SecondsTimer) Update() {
	if this.nanosPaused == 0 {
		nanosTotalNew := MonotonicNanos() - this.nanosStart - this.nanosDelay
		nanosDelta := nanosTotalNew - this.nanosTotal
		this.nanosTotal = nanosTotalNew
		this.secondsDelta = float64(nanosDelta) / (1000 * 1000 * 1000)
		this.secondsTotal = float64(this.nanosTotal) / (1000 * 1000 * 1000)
	}
}

// Returns the time in seconds between last two Update() calls
// (or between Reset() and Update() or the time set by Reset()).
func (this *SecondsTimer) Delta() float64 {
	return this.secondsDelta
}

// Returns the time in seconds between the Reset() call and Update() call (or the time set by Reset()).
func (this *SecondsTimer) Total() float64 {
	return this.secondsTotal
}
