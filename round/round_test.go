
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//     (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package round

import "testing"


func TestHalfToEven(t *testing.T) {
	if HalfToEven(0.0) != 0 {
		t.Error()
	}
	if HalfToEven(23.5) != 24 {
		t.Error()
	}
	if HalfToEven(23.9999999999999999) != 24 {
		t.Error()
	}
	if HalfToEven(24.0) != 24 {
		t.Error()
	}
	if HalfToEven(24.5) != 24 {
		t.Error()
	}
	if HalfToEven(24.9999999999999999) != 25 {
		t.Error()
	}
	if HalfToEven(-23.5) != -24 {
		t.Error()
	}
	if HalfToEven(-23.9999999999999999) != -24 {
		t.Error()
	}
	if HalfToEven(-24.0) != -24 {
		t.Error()
	}
	if HalfToEven(-24.5) != -24 {
		t.Error()
	}
	if HalfToEven(-24.9999999999999999) != -25 {
		t.Error()
	}
	if HalfToEven(0.4999999999999999) != 0 {
		t.Error()
	}
	if HalfToEven(1.4999999999999998) != 1 {
		t.Error()
	}
	if HalfToEven(1.4999999999999999) != 2 {
		t.Error()
	}
}

func TestHalfToOdd(t *testing.T) {
	if HalfToOdd(0.0) != 0 {
		t.Error()
	}
	if HalfToOdd(23.5) != 23 {
		t.Error()
	}
	if HalfToOdd(23.9999999999999999) != 24 {
		t.Error()
	}
	if HalfToOdd(24.0) != 24 {
		t.Error()
	}
	if HalfToOdd(24.5) != 25 {
		t.Error()
	}
	if HalfToOdd(24.9999999999999999) != 25 {
		t.Error()
	}
	if HalfToOdd(-23.5) != -23 {
		t.Error()
	}
	if HalfToOdd(-23.9999999999999999) != -24 {
		t.Error()
	}
	if HalfToOdd(-24.0) != -24 {
		t.Error()
	}
	if HalfToOdd(-24.5) != -25 {
		t.Error()
	}
	if HalfToOdd(-24.9999999999999999) != -25 {
		t.Error()
	}
	if HalfToOdd(0.4999999999999999) != 0 {
		t.Error()
	}
	if HalfToOdd(2.499999999999999) != 2 {
		t.Error()
	}
	if HalfToOdd(2.4999999999999999) != 3 {
		t.Error()
	}
}

