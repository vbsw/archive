
//     Copyright 2016 Vitali Baumtrok (vbsw@maibox.org)
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


// Provides functions to read the real and monotonic time in nanoseconds.
// Provides also timers to messure time between two events.
package time

import (
	"github.com/vbsw/semver"
	"time"
	_ "unsafe" // required to use //go:linkname
)

//go:noescape
//go:linkname nanotime runtime.nanotime
func nanotime() int64

var version semver.Version

// Current version is 2.1.0.
func Version() semver.Version {
	if version == nil {
		version = semver.New(2, 1, 0)
	}
	return version
}

// Returns system time in nanoseconds.
func RealNanos() int64 {
	return time.Now().UnixNano()
}

// Returns time in nanoseconds. Guaranteed always to be greater then the previous returned value.
func MonotonicNanos() int64 {
	return nanotime()
}
