
//     Copyright 2016 Vitali Baumtrok (vbsw@maibox.org)
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package time

type MultipleTimer struct {
	nanosStart   []int64
	nanosPaused  []int64
	nanosDelay   []int64
	nanosDelta   []int64
	nanosTotal   []int64
	microsDelta  []int64
	microsTotal  []int64
	millisDelta  []int64
	millisTotal  []int64
	secondsDelta []int64
	secondsTotal []int64
}

// Creates a new MultipleTimer and resets it.
// The start time for all timers is set to the same one value, MonotonicNanos().
func NewMultipleTimer(extent int) *MultipleTimer {
	timer := new(MultipleTimer)
	timer.nanosStart = make([]int64, extent)
	timer.nanosPaused = make([]int64, extent)
	timer.nanosDelay = make([]int64, extent)
	timer.nanosDelta = make([]int64, extent)
	timer.nanosTotal = make([]int64, extent)
	timer.microsDelta = make([]int64, extent)
	timer.microsTotal = make([]int64, extent)
	timer.millisDelta = make([]int64, extent)
	timer.millisTotal = make([]int64, extent)
	timer.secondsDelta = make([]int64, extent)
	timer.secondsTotal = make([]int64, extent)

	timer.resetNanosStartAll()

	return timer
}

// Resets nanosStart for all timers.
func (this *MultipleTimer) resetNanosStartAll() {
	nanosStartNew := MonotonicNanos()
	for i := range this.nanosStart {
		this.nanosStart[i] = nanosStartNew
	}
}

// Resets current and delta time to zero, and the start time
// to MonotonicNanos(). The argument, timerId, is the index
// representing the timer.
func (this *MultipleTimer) Reset(timerId int) {
	this.nanosPaused[timerId] = 0
	this.nanosDelay[timerId] = 0
	this.nanosDelta[timerId] = 0
	this.nanosTotal[timerId] = 0
	this.microsDelta[timerId] = 0
	this.microsTotal[timerId] = 0
	this.millisDelta[timerId] = 0
	this.millisTotal[timerId] = 0
	this.secondsDelta[timerId] = 0
	this.secondsTotal[timerId] = 0
	this.nanosStart[timerId] = MonotonicNanos()
}

// Resets current and delta time to zero for all timers.
// The start time for all timers is set to the same one value, MonotonicNanos().
func (this *MultipleTimer) ResetAll() {
	for i := range this.nanosPaused {
		this.nanosPaused[i] = 0
	}
	copy(this.nanosDelay, this.nanosPaused)
	copy(this.nanosDelta, this.nanosPaused)
	copy(this.nanosTotal, this.nanosPaused)
	copy(this.microsDelta, this.nanosPaused)
	copy(this.microsTotal, this.nanosPaused)
	copy(this.millisDelta, this.nanosPaused)
	copy(this.millisTotal, this.nanosPaused)
	copy(this.secondsDelta, this.nanosPaused)
	copy(this.secondsTotal, this.nanosPaused)

	this.resetNanosStartAll()
}

// Pauses the timer. Update() can't be called until Continue().
// The argument, timerId, is the index representing the timer.
func (this *MultipleTimer) Pause(timerId int) {
	if this.nanosPaused[timerId] == 0 {
		this.nanosPaused[timerId] = MonotonicNanos()
	}
}

// Pauses all timers. Update() can't be called until Continue().
func (this *MultipleTimer) PauseAll() {
	nanosPausedNew := MonotonicNanos()
	for i, nanosPaused := range this.nanosPaused {
		if nanosPaused == 0 {
			this.nanosPaused[i] = nanosPausedNew
		}
	}
}

// The timer is no more paused. The argument, timerId,
// is the index representing the timer.
func (this *MultipleTimer) Continue(timerId int) {
	if this.nanosPaused[timerId] > 0 {
		nanoContinue := MonotonicNanos()
		this.nanosDelay[timerId] += nanoContinue - this.nanosPaused[timerId]
		this.nanosPaused[timerId] = 0
	}
}

// All the timers are no more paused.
func (this *MultipleTimer) ContinueAll() {
	nanoContinue := MonotonicNanos()
	for i, nanosPaused := range this.nanosPaused {
		if nanosPaused > 0 {
			this.nanosDelay[i] += nanoContinue - nanosPaused
			this.nanosPaused[i] = 0
		}
	}
}

// Updates current and delta time (with MonotonicNanos() as the reference time).
// The argument, timerId, is the index representing the timer.
func (this *MultipleTimer) Update(timerId int) {
	if this.nanosPaused[timerId] == 0 {
		nanosTotalNew := MonotonicNanos() - this.nanosStart[timerId] - this.nanosDelay[timerId]
		microsTotalNew := nanosTotalNew / (1000)
		millisTotalNew := nanosTotalNew / (1000 * 1000)
		secondsTotalNew := nanosTotalNew / (1000 * 1000 * 1000)
		this.nanosDelta[timerId] = nanosTotalNew - this.nanosTotal[timerId]
		this.nanosTotal[timerId] = nanosTotalNew
		this.microsDelta[timerId] = microsTotalNew - this.microsTotal[timerId]
		this.microsTotal[timerId] = microsTotalNew
		this.millisDelta[timerId] = millisTotalNew - this.millisTotal[timerId]
		this.millisTotal[timerId] = millisTotalNew
		this.secondsDelta[timerId] = secondsTotalNew - this.secondsTotal[timerId]
		this.secondsTotal[timerId] = secondsTotalNew
	}
}

// Updates current and delta time for all timers
// (with MonotonicNanos() as the one single reference time).
func (this *MultipleTimer) UpdateAll() {
	for i, nanosPaused := range this.nanosPaused {
		if nanosPaused == 0 {
			nanosTotalNew := MonotonicNanos() - this.nanosStart[i] - this.nanosDelay[i]
			microsTotalNew := nanosTotalNew / (1000)
			millisTotalNew := nanosTotalNew / (1000 * 1000)
			secondsTotalNew := nanosTotalNew / (1000 * 1000 * 1000)
			this.nanosDelta[i] = nanosTotalNew - this.nanosTotal[i]
			this.nanosTotal[i] = nanosTotalNew
			this.microsDelta[i] = microsTotalNew - this.microsTotal[i]
			this.microsTotal[i] = microsTotalNew
			this.millisDelta[i] = millisTotalNew - this.millisTotal[i]
			this.millisTotal[i] = millisTotalNew
			this.secondsDelta[i] = secondsTotalNew - this.secondsTotal[i]
			this.secondsTotal[i] = secondsTotalNew
		}
	}
}

// Returns the time in nanoseconds between last two Update() calls
// (or between Reset() and Update() or the time set by Reset()).
// The argument, timerId, is the index representing the timer.
func (this *MultipleTimer) NanosDelta(timerId int) int64 {
	return this.nanosDelta[timerId]
}

// Returns the time in nanoseconds between the Reset() and Update() call (or the time set by Reset()).
// The argument, timerId, is the index representing the timer.
func (this *MultipleTimer) NanosTotal(timerId int) int64 {
	return this.nanosTotal[timerId]
}

// Returns the time in microseconds between last two Update() calls
// (or between Reset() and Update() or the time set by Reset()).
// The argument, timerId, is the index representing the timer.
// CAUTION: This is an accumulated value and can be greater then the delta
// returned by NanosDelta(), i.e. MicrosDelta()*1000 > NanosDelta().
func (this *MultipleTimer) MicrosDelta(timerId int) int64 {
	return this.microsDelta[timerId]
}

// Returns the time in milliseconds between the Reset() and Update() call (or the time set by Reset()).
// The argument, timerId, is the index representing the timer.
func (this *MultipleTimer) MicrosTotal(timerId int) int64 {
	return this.microsTotal[timerId]
}

// Returns the time in milliseconds between last two Update() calls
// (or between Reset() and Update() or the time set by Reset()).
// The argument, timerId, is the index representing the timer.
// CAUTION: This is an accumulated value and can be greater then the delta
// returned by NanosDelta(), i.e. MillisDelta()*1000000 > NanosDelta().
func (this *MultipleTimer) MillisDelta(timerId int) int64 {
	return this.millisDelta[timerId]
}

// Returns the time in milliseconds between the Reset() and Update() call (or the time set by Reset()).
// The argument, timerId, is the index representing the timer.
func (this *MultipleTimer) MillisTotal(timerId int) int64 {
	return this.millisTotal[timerId]
}

// Returns the time in seconds between last two Update() calls
// (or between Reset() and Update() or the time set by Reset()).
// The argument, timerId, is the index representing the timer.
// CAUTION: This is an accumulated value and can be greater then the delta
// returned by NanosDelta(), i.e. SecondsDelta()*1000000000 > NanosDelta().
func (this *MultipleTimer) SecondsDelta(timerId int) int64 {
	return this.secondsDelta[timerId]
}

// Returns the time in seconds between the Reset() and Update() call (or the time set by Reset()).
// The argument, timerId, is the index representing the timer.
func (this *MultipleTimer) SecondsTotal(timerId int) int64 {
	return this.secondsTotal[timerId]
}
