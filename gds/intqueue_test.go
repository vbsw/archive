
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package gds

import (
	"testing"
)

func TestIntQueuePush ( t *testing.T ) {
	queue := NewIntQueueFromCapacity(2)

	queue.Push(101)
	if queue.start != 0 { t.Error() }
	if len(queue.values) != 1 { t.Error() }
	if cap(queue.values) != 2 { t.Error() }

	queue.Push(102)
	if queue.start != 0 { t.Error() }
	if len(queue.values) != 2 { t.Error() }
	if cap(queue.values) != 2 { t.Error() }

	queue.start = 1
	queue.Push(103)
	if queue.start != 0 { t.Error() }
	if len(queue.values) != 2 { t.Error() }
	if cap(queue.values) != 2 { t.Error() }
}

func TestIntQueuePop ( t *testing.T ) {
	queue := NewIntQueue()

	queue.Push(101)
	queue.Push(102)
	queue.Push(103)
	if queue.start != 0 { t.Error() }
	if len(queue.values) != 3 { t.Error() }
	if cap(queue.values) != intqueue_INITIAL_CAPACITY { t.Error() }

	value := queue.Pop()
	if queue.Empty() { t.Error() }
	if value != 101 { t.Error() }
	if queue.start != 1 { t.Error() }
	if len(queue.values) != 3 { t.Error() }

	value = queue.Pop()
	queue.Push(104)
	if queue.Empty() { t.Error() }
	if value != 102 { t.Error() }
	if queue.start != 2 { t.Error() }
	if len(queue.values) != 4 { t.Error() }

	value = queue.Pop()
	value = queue.Pop()
	if queue.Empty() == false { t.Error() }
	if value != 104 { t.Error() }
	if queue.start != 0 { t.Error() }
	if len(queue.values) != 0 { t.Error() }
	if cap(queue.values) != intqueue_INITIAL_CAPACITY { t.Error() }
}

