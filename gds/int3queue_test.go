
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package gds

import (
	"testing"
)

func TestInt3QueuePush ( t *testing.T ) {
	queue := NewInt3QueueFromCapacity(12)

	queue.Push(101,111,121)
	if queue.start != 0 { t.Error() }
	if len(queue.values) != 3 { t.Error() }
	if cap(queue.values) != 12 { t.Error() }

	queue.Push(102,112,122)
	if queue.start != 0 { t.Error() }
	if len(queue.values) != 6 { t.Error() }
	if cap(queue.values) != 12 { t.Error() }

	queue.Push(103,113,123)
	queue.Push(104,114,124)
	queue.start = 9
	queue.Push(105,115,125)
	if queue.start != 0 { t.Error() }
	if len(queue.values) != 6 { t.Error() }
	if cap(queue.values) != 12 { t.Error() }
}

func TestInt3QueuePop ( t *testing.T ) {
	queue := NewInt3Queue()

	queue.Push(101,111,121)
	queue.Push(102,112,122)
	queue.Push(103,113,123)
	if queue.start != 0 { t.Error() }
	if len(queue.values) != 9 { t.Error() }
	if cap(queue.values) != int3queue_INITIAL_CAPACITY { t.Error() }

	val1, val2, val3 := queue.Pop()
	if queue.Empty() { t.Error() }
	if val1 != 101 { t.Error() }
	if val2 != 111 { t.Error() }
	if val3 != 121 { t.Error() }
	if queue.start != 3 { t.Error() }
	if len(queue.values) != 9 { t.Error() }

	val1, val2, val3 = queue.Pop()
	queue.Push(104,114,124)
	if queue.Empty() { t.Error() }
	if val1 != 102 { t.Error() }
	if val2 != 112 { t.Error() }
	if val3 != 122 { t.Error() }
	if queue.start != 6 { t.Error() }
	if len(queue.values) != 12 { t.Error() }

	val1, val2, val3 = queue.Pop()
	val1, val2, val3 = queue.Pop()
	if queue.Empty() == false { t.Error() }
	if val1 != 104 { t.Error() }
	if val2 != 114 { t.Error() }
	if val3 != 124 { t.Error() }
	if queue.start != 0 { t.Error() }
	if len(queue.values) != 0 { t.Error() }
	if cap(queue.values) != int3queue_INITIAL_CAPACITY { t.Error() }
}

