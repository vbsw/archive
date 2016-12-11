
//          Copyright Vitali Baumtrok 2015.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


/**
 * @file test.c
 * @author Vitali Baumtrok
 * @date 2015-05-06
 * @brief membuf function tests
 * 
 * Tests the functions of membuf. To compile it use for example the command
 * 'gcc test.c membuf.c -o main'.
 */

#include "membuf.h"

#include <stdlib.h>
#include <stdio.h>

static const char *const NAME_allocate  = "MemBuf_allocate";
static const char *const NAME_getMemory  = "MemBuf_getMemory";

static int
failure
(
	const int result,
	const char *const functionName,
	const char *const message
)
{
	if ( 0 == result )
	{
		printf ( "%s : %s\n", functionName, message );
		return 1;
	}
	else
	{
		return 0;
	}
}

static int
test_allocate
()
{
	const size_t memorySize = 4 * sizeof ( int );
	void **buffer = (void **) MemBuf_allocate ( memorySize );

	if ( failure ( 0!=buffer, NAME_allocate, "no buffer allocated (out of memory?)" ) ) { return 1; }
	else
	{
		if ( failure ( 0!=buffer[0], NAME_allocate, "no buffer meta set (1)" ) )
			return 1;
		if ( failure ( 0!=buffer[1], NAME_allocate, "no buffer meta set (2)" ) )
			return 1;
		if ( failure ( 0!=buffer[2], NAME_allocate, "no buffer meta set (3)" ) )
			return 1;
		if ( failure ( 0==buffer[3], NAME_allocate, "no buffer meta set (4)" ) )
			return 1;
		if ( failure ( buffer[0]==buffer[1], NAME_allocate, "wrong buffer meta set (1)" ) )
			return 1;
		if ( failure ( buffer[2]==buffer[0]+memorySize, NAME_allocate, "wrong buffer meta set (2)" ) )
			return 1;
	}
	MemBuf_free ( buffer );

	return 0;
}

static int
test_getMemory () {
	const size_t memorySize = 4 * sizeof ( int );
	void **buffer = (void **) MemBuf_allocate ( memorySize );

	if ( failure ( 0!=buffer, NAME_getMemory, "no buffer allocated (out of memory?)" ) ) { return 1; }
	else
	{
		if ( failure ( 0!=buffer[0], NAME_getMemory, "no buffer meta set (1)" ) )
			return 1;
		if ( failure ( 0!=buffer[1], NAME_getMemory, "no buffer meta set (2)" ) )
			return 1;
		if ( failure ( 0!=buffer[2], NAME_getMemory, "no buffer meta set (3)" ) )
			return 1;
		if ( failure ( 0==buffer[3], NAME_getMemory, "no buffer meta set (4)" ) )
			return 1;
		if ( failure ( buffer[0]==buffer[1], NAME_getMemory, "wrong buffer meta set (1)" ) )
			return 1;
		if ( failure ( buffer[2]==buffer[0]+memorySize, NAME_getMemory, "wrong buffer meta set (2)" ) )
			return 1;
	}

	void *memory = MemBuf_getMemory ( buffer, memorySize/2 );

	if ( failure ( 0!=memory, NAME_getMemory, "no memory allocated (out of memory?) (1)" ) ) { return 1; }
	else
	{
		if ( failure ( buffer[0]!=buffer[1], NAME_getMemory, "wrong buffer meta set (3)" ) )
			return 1;
		if ( failure ( 0==buffer[3], NAME_getMemory, "wrong buffer meta set (4)" ) )
			return 1;
	}

	memory = MemBuf_getMemory ( buffer, memorySize/2 );

	if ( failure ( 0!=memory, NAME_getMemory, "no memory allocated (out of memory?) (2)" ) ) { return 0; }
	else
	{
		if ( failure ( buffer[1]==buffer[2], NAME_getMemory, "wrong buffer meta set (5)" ) )
			return 0;
		if ( failure ( 0==buffer[3], NAME_getMemory, "wrong buffer meta set (6)" ) )
			return 0;
	}

	memory = MemBuf_getMemory ( buffer, memorySize/2 );

	if ( failure ( 0!=memory, NAME_getMemory, "no memory allocated (out of memory?) (3)" ) ) { return 1; }
	else
	{
		if ( failure ( buffer[1]==buffer[2], NAME_getMemory, "wrong buffer meta set (7)" ) )
			return 1;
		if ( failure ( 0!=buffer[3], NAME_getMemory, "wrong buffer meta set (8)" ) )
			return 1;
		if ( failure ( ((void**) buffer[3])[0]!=((void**) buffer[3])[1], NAME_getMemory, "wrong buffer meta set (9)" ) )
			return 1;
		if ( failure ( ((void**) buffer[3])[1]!=((void**) buffer[3])[2], NAME_getMemory, "wrong buffer meta set (10)" ) )
			return 1;
	}

	memory = MemBuf_getMemory ( buffer, 2 * memorySize );

	if ( failure ( 0!=memory, NAME_getMemory, "no memory allocated (out of memory?) (4)" ) ) { return 1; }
	else
	{
		if ( failure ( buffer[1]==buffer[2], NAME_getMemory, "wrong buffer meta set (11)" ) )
			return 1;
		if ( failure ( 0!=buffer[3], NAME_getMemory, "wrong buffer meta set (12)" ) )
			return 1;
		if ( failure ( 0!=((void**) buffer[3])[3], NAME_getMemory, "wrong buffer meta set (13)" ) )
			return 1;
		if ( failure ( ((void**) ((void**) buffer[3])[3])[1]==((void**) ((void**) buffer[3])[3])[2], NAME_getMemory, "wrong buffer meta set (14)" ) )
			return 1;
	}

	MemBuf_free ( buffer );

	return 0;
}

int
main
(
	int argc,
	const char **argv
)
{
	if ( test_allocate () ) return 0;
	if ( test_getMemory () ) return 0;

	printf ( "MemBuf tests passed\n" );
	return 0;
}
