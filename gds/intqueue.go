
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package gds

const (
	intqueue_INITIAL_CAPACITY = 15
)

type IntQueue struct {
	values []int
	start int
}

func NewIntQueue ( ) *IntQueue {
	queue := NewIntQueueFromCapacity(intqueue_INITIAL_CAPACITY)
	return queue
}

func NewIntQueueFromCapacity ( capacity int ) *IntQueue {
	queue := new(IntQueue)
	queue.values = make([]int,0,capacity)
	queue.start = 0
	return queue
}

func (this *IntQueue) Pop ( ) int {
	if this.start < len(this.values) {
		index := this.start
		value := this.values[index]
		this.start += 1;

		if this.start == len(this.values) {
			this.values = this.values[:0]
			this.start = 0;
		}
		return value

	} else {
		return 0;
	}
}

func (this *IntQueue) Push ( value int ) {
	length := len(this.values)

	// copy to front instead of reallocating
	if length == cap(this.values) && this.start >= length/2 {
		from := this.start
		to := len(this.values)
		copy(this.values,this.values[from:to])
		this.values = this.values[:length-this.start]
		this.start = 0
	}

	this.values = append(this.values,value)
}

func (this *IntQueue) Empty ( ) bool {
	return this.start == len(this.values)
}

