
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package visimpl


import (
	"testing"
)

func TestCreate ( t *testing.T ) {
	vsys := New()
	eventQueue := vsys.EventQueue()
	vsys = vsys.Destroy()

	if vsys != nil { t.Error() }
	if eventQueue.PollEvents() != false { t.Error() }
}

