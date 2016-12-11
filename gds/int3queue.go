
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package gds

const (
	int3queue_INITIAL_CAPACITY = 30
)

type Int3Queue struct {
	values []int
	start int
}

func NewInt3Queue ( ) *Int3Queue {
	queue := NewInt3QueueFromCapacity(int3queue_INITIAL_CAPACITY)
	return queue
}

func NewInt3QueueFromCapacity ( capacity int ) *Int3Queue {
	queue := new(Int3Queue)
	queue.values = make([]int,0,capacity)
	queue.start = 0
	return queue
}

func (this *Int3Queue) Pop ( ) (int,int,int) {
	if this.start < len(this.values) {
		index := this.start
		val1 := this.values[index]
		val2 := this.values[index+1]
		val3 := this.values[index+2]
		this.start += 3;

		if this.start == len(this.values) {
			this.values = this.values[:0]
			this.start = 0;
		}
		return val1, val2, val3

	} else {
		return 0, 0, 0;
	}
}

func (this *Int3Queue) Push ( val1, val2, val3 int ) {
	length := len(this.values)

	// copy to front instead of reallocating
	if length == cap(this.values) && this.start >= length/2+3 {
		from := this.start
		to := len(this.values)
		copy(this.values,this.values[from:to])
		this.values = this.values[:length-this.start]
		this.start = 0
	}

	this.values = append(this.values,val1,val2,val3)
}

func (this *Int3Queue) Empty ( ) bool {
	return this.start == len(this.values)
}

