
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


#include <stdlib.h>
#include <string.h>

#define IQUEUE_INITIAL_CAPACITY 15

typedef struct {
	int* values;
	size_t start;
	size_t end;
	size_t capacity;
} iqueue_t;

iqueue_t* iqueue_destroy ( iqueue_t *const queue ) {
	free ( queue->values );
	return queue;
}

void iqueue_ensure_capacity ( iqueue_t *const queue ) {
	const size_t end = queue->end;
	const size_t capacity = queue->capacity;

	if ( end == capacity ) {
		const size_t start = queue->start;
		const size_t count = end - start;

		if ( start < capacity/2 ) {
			const size_t new_capacity = capacity*2;
			int *const new_values = (int*) malloc ( sizeof(int)*new_capacity );
			memcpy ( new_values + start, queue->values + start, sizeof(int)*count );
			free ( queue->values );
			queue->values = new_values;
			queue->capacity = new_capacity;
		} else {
			memmove ( queue->values, queue->values + start, sizeof(int)*count );
			queue->start = 0;
			queue->end = count;
		}
	}
}

iqueue_t* iqueue_free ( iqueue_t *const queue ) {
	iqueue_destroy ( queue );
	free ( queue );
	return 0;
}

void iqueue_init_from_capacity ( iqueue_t* queue, size_t initial_capacity ) {
	queue->values = (int*) malloc ( sizeof(int)*initial_capacity );
	queue->start = 0;
	queue->end = 0;
	queue->capacity = initial_capacity;
}

void iqueue_init ( iqueue_t *const queue ) {
	iqueue_init_from_capacity ( queue, IQUEUE_INITIAL_CAPACITY );
}

iqueue_t* iqueue_new_from_capacity ( size_t initial_capacity ) {
	iqueue_t *const queue = (iqueue_t *const) malloc ( sizeof(iqueue_t) );
	iqueue_init_from_capacity ( queue, initial_capacity );
	return queue;
}

iqueue_t* iqueue_new ( void ) {
	iqueue_t *const queue = iqueue_new_from_capacity ( IQUEUE_INITIAL_CAPACITY );
	return queue;
}

int iqueue_has_next ( iqueue_t *const queue ) {
	return queue->start < queue->end;
}

void iqueue_pop ( iqueue_t *const queue, int *const value ) {
	if ( queue->start < queue->end ) {
		const size_t index = queue->start;
		value[0] = queue->values[index];
		queue->start += 1;

		if ( queue->start == queue->end ) {
			queue->start = 0;
			queue->end = 0;
		}

	} else {
		value[0] = 0;
	}
}

void iqueue_push ( iqueue_t* queue, int value ) {
	iqueue_ensure_capacity ( queue );

	const size_t index = queue->end;
	queue->values[index] = value;
	queue->end += 1;
}

