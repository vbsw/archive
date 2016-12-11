
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package visx


import (
	"testing"
	"gitlab.com/vbit/vis"
)

func TestCreate ( t *testing.T ) {
	Create()

	vis.Destroy()
	if vis.Queue != nil { t.Error() }
	if vis.Graphics != nil { t.Error() }
	if vis.Window != nil { t.Error() }
}

