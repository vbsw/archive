
//          Copyright Vitali Baumtrok 2015.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


/**
 * @file cb.test.main.c
 * @author Vitali Baumtrok
 * @date 2015-04-04
 * @brief cb function tests
 * 
 * Tests the functions of cb. To compile it use for example the command
 * 'gcc cb.test.main.c cb.c -o main'.
 */

#include "cb.h"

#include <stdlib.h>
#include <stdio.h>

static const char *const NAME_allocateBuffer  = "cb_allocateBuffer";
static const char *const NAME_getBufferMemory  = "cb_getBufferMemory";
static const char *const NAME_stringBinarySearchAsc  = "cb_stringBinarySearchAsc";
static const char *const NAME_stringBinarySearchDesc = "cb_stringBinarySearchDesc";

static int
failure ( const int result, const char *const functionName, const char *const message ) {
	if ( 0 == result ) {
		printf ( "%s : %s\n", functionName, message );
		return 1;
	}

	else {
		return 0;
	}
}

static int
test_allocateBuffer () {
	const size_t memorySize = 4 * sizeof ( int );
	void **buffer = (void **) cb_allocateBuffer ( memorySize );

	if ( failure ( 0!=buffer, NAME_stringBinarySearchAsc, "no buffer allocated (out of memory?)" ) ) { return 1; }
	else {
		if ( failure ( 0!=buffer[0], NAME_stringBinarySearchAsc, "no buffer meta set (1)" ) )
			return 1;
		if ( failure ( 0!=buffer[1], NAME_stringBinarySearchAsc, "no buffer meta set (2)" ) )
			return 1;
		if ( failure ( 0!=buffer[2], NAME_stringBinarySearchAsc, "no buffer meta set (3)" ) )
			return 1;
		if ( failure ( 0==buffer[3], NAME_stringBinarySearchAsc, "no buffer meta set (4)" ) )
			return 1;
		if ( failure ( buffer[0]==buffer[1], NAME_stringBinarySearchAsc, "wrong buffer meta set (1)" ) )
			return 1;
		if ( failure ( buffer[2]==buffer[0]+memorySize, NAME_stringBinarySearchAsc, "wrong buffer meta set (2)" ) )
			return 1;
	}
	cb_freeBuffer ( buffer );

	return 0;
}

static int
test_getBufferMemory () {
	const size_t memorySize = 4 * sizeof ( int );
	void **buffer = (void **) cb_allocateBuffer ( memorySize );

	if ( failure ( 0!=buffer, NAME_getBufferMemory, "no buffer allocated (out of memory?)" ) ) { return 1; }
	else {
		if ( failure ( 0!=buffer[0], NAME_getBufferMemory, "no buffer meta set (1)" ) )
			return 1;
		if ( failure ( 0!=buffer[1], NAME_getBufferMemory, "no buffer meta set (2)" ) )
			return 1;
		if ( failure ( 0!=buffer[2], NAME_getBufferMemory, "no buffer meta set (3)" ) )
			return 1;
		if ( failure ( 0==buffer[3], NAME_getBufferMemory, "no buffer meta set (4)" ) )
			return 1;
		if ( failure ( buffer[0]==buffer[1], NAME_getBufferMemory, "wrong buffer meta set (1)" ) )
			return 1;
		if ( failure ( buffer[2]==buffer[0]+memorySize, NAME_getBufferMemory, "wrong buffer meta set (2)" ) )
			return 1;
	}

	void *memory = cb_getBufferMemory ( buffer, memorySize/2 );

	if ( failure ( 0!=memory, NAME_getBufferMemory, "no memory allocated (out of memory?) (1)" ) ) { return 1; }
	else {
		if ( failure ( buffer[0]!=buffer[1], NAME_getBufferMemory, "wrong buffer meta set (3)" ) )
			return 1;
		if ( failure ( 0==buffer[3], NAME_getBufferMemory, "wrong buffer meta set (4)" ) )
			return 1;
	}

	memory = cb_getBufferMemory ( buffer, memorySize/2 );

	if ( failure ( 0!=memory, NAME_getBufferMemory, "no memory allocated (out of memory?) (2)" ) ) { return 0; }
	else {
		if ( failure ( buffer[1]==buffer[2], NAME_getBufferMemory, "wrong buffer meta set (5)" ) )
			return 0;
		if ( failure ( 0==buffer[3], NAME_getBufferMemory, "wrong buffer meta set (6)" ) )
			return 0;
	}

	memory = cb_getBufferMemory ( buffer, memorySize/2 );

	if ( failure ( 0!=memory, NAME_getBufferMemory, "no memory allocated (out of memory?) (3)" ) ) { return 1; }
	else {
		if ( failure ( buffer[1]==buffer[2], NAME_getBufferMemory, "wrong buffer meta set (7)" ) )
			return 1;
		if ( failure ( 0!=buffer[3], NAME_getBufferMemory, "wrong buffer meta set (8)" ) )
			return 1;
		if ( failure ( ((void**) buffer[3])[0]!=((void**) buffer[3])[1], NAME_getBufferMemory, "wrong buffer meta set (9)" ) )
			return 1;
		if ( failure ( ((void**) buffer[3])[1]!=((void**) buffer[3])[2], NAME_getBufferMemory, "wrong buffer meta set (10)" ) )
			return 1;
	}

	memory = cb_getBufferMemory ( buffer, 2 * memorySize );

	if ( failure ( 0!=memory, NAME_getBufferMemory, "no memory allocated (out of memory?) (4)" ) ) { return 1; }
	else {
		if ( failure ( buffer[1]==buffer[2], NAME_getBufferMemory, "wrong buffer meta set (11)" ) )
			return 1;
		if ( failure ( 0!=buffer[3], NAME_getBufferMemory, "wrong buffer meta set (12)" ) )
			return 1;
		if ( failure ( 0!=((void**) buffer[3])[3], NAME_getBufferMemory, "wrong buffer meta set (13)" ) )
			return 1;
		if ( failure ( ((void**) ((void**) buffer[3])[3])[1]==((void**) ((void**) buffer[3])[3])[2], NAME_getBufferMemory, "wrong buffer meta set (14)" ) )
			return 1;
	}

	cb_freeBuffer ( buffer );

	return 0;
}

static int
test_stringBinarySearchAsc () {
	const char **stringArray = (const char**) malloc ( sizeof(char*) * 7 );
	size_t index = 1000;
	stringArray[0] = "drei";
	stringArray[1] = "eins";
	stringArray[2] = "vier";
	stringArray[3] = "zwei";

	const char *string = "fuenf";

	if ( failure (	0 == cb_stringBinarySearchAsc ( &index, stringArray, 0, string ), NAME_stringBinarySearchAsc, "string found (1)" ) ) { return 1; }
	else {
		if ( failure ( index == 0, NAME_stringBinarySearchAsc, "string at wrong index (1)" ) )
			return 1;
	}

	if ( failure (	0 == cb_stringBinarySearchAsc ( &index, stringArray, 1, string ), NAME_stringBinarySearchAsc, "string found (2)" ) ) { return 1; }
	else {
		if ( failure ( index == 1, NAME_stringBinarySearchAsc, "string at wrong index (2)" ) )
			return 1;
	}

	if ( failure (	0 == cb_stringBinarySearchAsc ( &index, stringArray, 3, string ), NAME_stringBinarySearchAsc, "string found (3)" ) ) { return 1; }
	else {
		if ( failure ( index == 2, NAME_stringBinarySearchAsc, "string at wrong index (3)" ) )
			return 1;
	}

	if ( failure (	0 == cb_stringBinarySearchAsc ( &index, stringArray, 4, string ), NAME_stringBinarySearchAsc, "string found (4)" ) ) { return 1; }
	else {
		if ( failure ( index == 2, NAME_stringBinarySearchAsc, "string at wrong index (4)" ) )
			return 1;
	}

	string = "eins";

	if ( failure (	1 == cb_stringBinarySearchAsc ( &index, stringArray, 3, string ), NAME_stringBinarySearchAsc, "string not found (1)" ) ) { return 1; }
	else {
		if ( failure ( index == 1, NAME_stringBinarySearchAsc, "string at wrong index (5)" ) )
			return 1;
	}

	if ( failure (	1 == cb_stringBinarySearchAsc ( &index, stringArray, 4, string ), NAME_stringBinarySearchAsc, "string not found (2)" ) ) { return 1; }
	else {
		if ( failure ( index == 1, NAME_stringBinarySearchAsc, "string at wrong index (6)" ) )
			return 1;
	}

	string = "alpha";

	if ( failure (	0 == cb_stringBinarySearchAsc ( &index, stringArray, 1, string ), NAME_stringBinarySearchAsc, "string found (5)" ) ) { return 1; }
	else {
		if ( failure ( index == 0, NAME_stringBinarySearchAsc, "string at wrong index (7)" ) )
			return 1;
	}

	if ( failure (	0 == cb_stringBinarySearchAsc ( &index, stringArray, 3, string ), NAME_stringBinarySearchAsc, "string found (6)" ) ) { return 1; }
	else {
		if ( failure ( index == 0, NAME_stringBinarySearchAsc, "string at wrong index (8)" ) )
			return 1;
	}

	if ( failure (	0 == cb_stringBinarySearchAsc ( &index, stringArray, 4, string ), NAME_stringBinarySearchAsc, "string found (7)" ) ) { return 1; }
	else {
		if ( failure ( index == 0, NAME_stringBinarySearchAsc, "string at wrong index (9)" ) )
			return 1;
	}

	string = "zzz";

	if ( failure (	0 == cb_stringBinarySearchAsc ( &index, stringArray, 1, string ), NAME_stringBinarySearchAsc, "string found (8)" ) ) { return 1; }
	else {
		if ( failure ( index == 1, NAME_stringBinarySearchAsc, "string at wrong index (10)" ) )
			return 1;
	}

	if ( failure (	0 == cb_stringBinarySearchAsc ( &index, stringArray, 3, string ), NAME_stringBinarySearchAsc, "string found (9)" ) ) { return 1; }
	else {
		if ( failure ( index == 3, NAME_stringBinarySearchAsc, "string at wrong index (11)" ) )
			return 1;
	}

	if ( failure (	0 == cb_stringBinarySearchAsc ( &index, stringArray, 4, string ), NAME_stringBinarySearchAsc, "string found (10)" ) ) { return 1; }
	else {
		if ( failure ( index == 4, NAME_stringBinarySearchAsc, "string at wrong index (12)" ) )
			return 1;
	}

	free ( stringArray );

	return 0;
}

static int
test_stringBinarySearchDesc () {
	const char **stringArray = (const char**) malloc ( sizeof(char*) * 7 );
	size_t index = 1000;
	stringArray[0] = "zwei";
	stringArray[1] = "vier";
	stringArray[2] = "eins";
	stringArray[3] = "drei";

	const char *string = "fuenf";

	if ( failure (	0 == cb_stringBinarySearchAsc ( &index, stringArray, 0, string ), NAME_stringBinarySearchAsc, "string found (1)" ) ) { return 1; }
	else {
		if ( failure ( index == 0, NAME_stringBinarySearchAsc, "string at wrong index (1)" ) )
			return 1;
	}

	if ( failure (	0 == cb_stringBinarySearchDesc ( &index, stringArray, 1, string ), NAME_stringBinarySearchDesc, "string found (2)" ) ) { return 1; }
	else {
		if ( failure ( index == 1, NAME_stringBinarySearchDesc, "string at wrong index (2)" ) )
			return 1;
	}

	if ( failure (	0 == cb_stringBinarySearchDesc ( &index, stringArray, 3, string ), NAME_stringBinarySearchDesc, "string found (3)" ) ) { return 1; }
	else {
		if ( failure ( index == 2, NAME_stringBinarySearchDesc, "string at wrong index (3)" ) )
			return 1;
	}

	if ( failure (	0 == cb_stringBinarySearchDesc ( &index, stringArray, 4, string ), NAME_stringBinarySearchDesc, "string found (4)" ) ) { return 1; }
	else {
		if ( failure ( index == 2, NAME_stringBinarySearchDesc, "string at wrong index (4)" ) )
			return 1;
	}

	string = "eins";

	if ( failure (	1 == cb_stringBinarySearchDesc ( &index, stringArray, 3, string ), NAME_stringBinarySearchDesc, "string not found (1)" ) ) { return 1; }
	else {
		if ( failure ( index == 2, NAME_stringBinarySearchDesc, "string at wrong index (5)" ) )
			return 1;
	}

	if ( failure (	1 == cb_stringBinarySearchDesc ( &index, stringArray, 4, string ), NAME_stringBinarySearchDesc, "string not found (2)" ) ) { return 1; }
	else {
		if ( failure ( index == 2, NAME_stringBinarySearchDesc, "string at wrong index (6)" ) )
			return 1;
	}

	string = "alpha";

	if ( failure (	0 == cb_stringBinarySearchDesc ( &index, stringArray, 1, string ), NAME_stringBinarySearchDesc, "string found (5)" ) ) { return 1; }
	else {
		if ( failure ( index == 1, NAME_stringBinarySearchDesc, "string at wrong index (7)" ) )
			return 1;
	}

	if ( failure (	0 == cb_stringBinarySearchDesc ( &index, stringArray, 3, string ), NAME_stringBinarySearchDesc, "string found (6)" ) ) { return 1; }
	else {
		if ( failure ( index == 3, NAME_stringBinarySearchDesc, "string at wrong index (8)" ) )
			return 1;
	}

	if ( failure (	0 == cb_stringBinarySearchDesc ( &index, stringArray, 4, string ), NAME_stringBinarySearchDesc, "string found (7)" ) ) { return 1; }
	else {
		if ( failure ( index == 4, NAME_stringBinarySearchDesc, "string at wrong index (9)" ) )
			return 1;
	}

	string = "zzz";

	if ( failure (	0 == cb_stringBinarySearchDesc ( &index, stringArray, 1, string ), NAME_stringBinarySearchDesc, "string found (8)" ) ) { return 1; }
	else {
		if ( failure ( index == 0, NAME_stringBinarySearchDesc, "string at wrong index (10)" ) )
			return 1;
	}

	if ( failure (	0 == cb_stringBinarySearchDesc ( &index, stringArray, 3, string ), NAME_stringBinarySearchDesc, "string found (9)" ) ) { return 1; }
	else {
		if ( failure ( index == 0, NAME_stringBinarySearchDesc, "string at wrong index (11)" ) )
			return 1;
	}

	if ( failure (	0 == cb_stringBinarySearchDesc ( &index, stringArray, 4, string ), NAME_stringBinarySearchDesc, "string found (10)" ) ) { return 1; }
	else {
		if ( failure ( index == 0, NAME_stringBinarySearchDesc, "string at wrong index (12)" ) )
			return 1;
	}

	free ( stringArray );

	return 0;
}

int
main ( int argc, const char **argv ) {
	if ( test_allocateBuffer () ) return 0;
	if ( test_getBufferMemory () ) return 0;
	if ( test_stringBinarySearchAsc () ) return 0;
	if ( test_stringBinarySearchDesc () ) return 0;

	printf ( "cb tests passed\n" );
	return 0;
}
