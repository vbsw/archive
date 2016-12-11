
//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


#include <stdlib.h>
#include <string.h>


typedef struct {
	char *values;
	size_t len;
	size_t cap;
} StackC;


typedef struct {
	size_t *values;
	size_t len;
	size_t cap;
} StackS;


typedef struct {
	void **values;
	size_t len;
	size_t cap;
} StackP;




void
StackC_allocate ( StackC *const stack ) {
	const size_t length = (256-8) / sizeof(char);
	(*stack).values = (char*) malloc ( length * sizeof(char) );
	(*stack).len = 0;
	(*stack).cap = length;
}


void
StackC_deallocate ( StackC *const stack ) {
	free ( (*stack).values );
	(*stack).len = 0;
	(*stack).cap = 0;
}

void
StackC_push ( StackC *const stack, const char el ) {
	const size_t length = (*stack).len;

	// assure capacity
	if ( length==(*stack).cap ) {
		const size_t newLength = 2*length;
		char *const newValues = (char*) malloc ( newLength * sizeof(char) );

		memcpy ( newValues, (*stack).values, length * sizeof(char) );
		free ( (*stack).values );
		(*stack).values = newValues;
		(*stack).cap = newLength;
	}

	(*stack).values[length] = el;
	(*stack).len += 1;
}




void
StackS_allocate ( StackS *const stack ) {
	const size_t length = (256-8) / sizeof(size_t);
	(*stack).values = (size_t*) malloc ( length * sizeof(size_t) );
	(*stack).len = 0;
	(*stack).cap = length;
}


void
StackS_deallocate ( StackS *const stack ) {
	free ( (*stack).values );
	(*stack).len = 0;
	(*stack).cap = 0;
}

void
StackS_push ( StackS *const stack, const size_t el ) {
	const size_t length = (*stack).len;

	// assure capacity
	if ( length==(*stack).cap ) {
		const size_t newLength = 2*length;
		size_t *const newValues = (size_t*) malloc ( newLength * sizeof(size_t) );

		memcpy ( newValues, (*stack).values, length * sizeof(size_t) );
		free ( (*stack).values );
		(*stack).values = newValues;
		(*stack).cap = newLength;
	}

	(*stack).values[length] = el;
	(*stack).len += 1;
}




void
StackP_allocate ( StackP *const stack ) {
	const size_t length = (512-8) / sizeof(void*);
	(*stack).values = (void**) malloc ( length * sizeof(void*) );
	(*stack).len = 0;
	(*stack).cap = length;
}


void
StackP_deallocate ( StackP *const stack ) {
	free ( (*stack).values );
	(*stack).len = 0;
	(*stack).cap = 0;
}

void
StackP_push ( StackP *const stack, void *const el ) {
	const size_t length = (*stack).len;

	// assure capacity
	if ( length==(*stack).cap ) {
		const size_t newLength = 2*length;
		void **const newValues = (void**) malloc ( newLength * sizeof(void*) );

		memcpy ( newValues, (*stack).values, length * sizeof(void*) );
		free ( (*stack).values );
		(*stack).values = newValues;
		(*stack).cap = newLength;
	}

	(*stack).values[length] = el;
	(*stack).len += 1;
}




























































