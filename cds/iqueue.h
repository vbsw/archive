
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)

#ifndef IQUEUE_H
#define IQUEUE_H

#include <stdlib.h>

#define IQUEUE_INITIAL_CAPACITY 15

typedef struct {
	int* values;
	size_t start;
	size_t end;
	size_t capacity;
} iqueue_t;

extern iqueue_t* iqueue_destroy ( iqueue_t* queue );
extern void iqueue_ensure_capacity ( iqueue_t* queue );
extern iqueue_t* iqueue_free ( iqueue_t* queue );
extern void iqueue_init_from_capacity ( iqueue_t* queue, size_t initial_capacity );
extern void iqueue_init ( iqueue_t* queue );
extern iqueue_t* iqueue_new_from_capacity ( size_t initial_capacity );
extern iqueue_t* iqueue_new ( void );
extern int iqueue_has_next ( iqueue_t* queue );
extern void iqueue_pop ( iqueue_t* queue, int* value );
extern void iqueue_push ( iqueue_t* queue, int value );

#endif /* IQUEUE_H */

