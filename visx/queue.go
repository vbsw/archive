
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package visx

import (
	"gitlab.com/vbit/vis"
)

type tQueue struct {

}

func (this *tQueue) Empty ( ) bool {
	return false
}

func (this *tQueue) PushEvent ( eventTyp int, value1, value2 float32 ) {

}

func (this *tQueue) PopEvent ( ) (int,float32,float32) {
	return vis.EVENT_NONE, 0, 0
}

func (this *tQueue) Destroy ( ) vis.IQueue {
	return nil
}

