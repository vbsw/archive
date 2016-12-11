
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        <http://www.boost.org/LICENSE_1_0.txt>)


#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "strl.h"

#define CAPACITY_INITIAL 14
#define INDEX_LEN 0
#define INDEX_CAP 1
#define INDEX_ERR 2


size_t
strlnextcap ( size_t n, size_t capacity ) {
	size_t capacity_additional = 0;

	while ( capacity_additional < n ) {
		capacity_additional = capacity;
		capacity += capacity_additional;
	}
	return capacity;
}

void*
strlnewcap ( size_t capacity ) {
	void *const char_arr = malloc ( sizeof(char *) * capacity );

	if ( char_arr ) {
		void**const strl_ptrs = (void**) malloc ( sizeof(void*) * 3 );

		if ( strl_ptrs ) {
			size_t*const strl_numbers = (size_t*) malloc ( sizeof(size_t) * 3 );

			if ( strl_numbers ) {
				strl_numbers[INDEX_LEN] = 0;
				strl_numbers[INDEX_CAP] = capacity;
				strl_numbers[INDEX_ERR] = 0;
				strl_ptrs[0] = (void*) strl_numbers;
				strl_ptrs[1] = (void*) &strlnextcap;
				strl_ptrs[2] = char_arr;

				return (void*) strl_ptrs;

			} else {
				free ( char_arr );
				free ( strl_ptrs );

				return 0;
			}

		} else {
			free ( char_arr );

			return 0;
		}

	} else {
		return 0;
	}
}

void*
strlnew ( void ) {
	return strlnewcap ( CAPACITY_INITIAL );
}

void
strldest ( void* strl ) {
	void**const strl_ptrs = (void**) strl;
	size_t*const strl_numbers = (size_t*) strl_ptrs[0];
	char**const char_arr = (char**) strl_ptrs[2];
	const size_t len = strl_numbers[INDEX_LEN];
	size_t i;

	for ( i = 0; i < len; i += 1 ) {
		free ( char_arr[i] );
	}

	free ( strl_ptrs[0] );
	free ( strl_ptrs[2] );
	free ( strl );
}

size_t
strllen ( const void* strl ) {
	const size_t*const*const strl_numbers = (const size_t*const*const) strl;

	return strl_numbers[0][INDEX_LEN];
}

size_t
strlcap ( const void* strl ) {
	const size_t*const*const strl_numbers = (const size_t*const*const) strl;

	return strl_numbers[0][INDEX_CAP];
}

unsigned short int
strlerr ( const void* strl ) {
	const size_t*const*const strl_numbers = (const size_t*const*const) strl;

	return (unsigned short int) strl_numbers[0][INDEX_ERR];
}

extern void
strlerrclear ( void* strl ) {
	const size_t*const*const strl_numbers = (const size_t*const*const) strl;

	strl_numbers[0][INDEX_ERR] = 0;
}

void
strlcapfunc ( void* strl, size_t (*capfunc) ( size_t n, size_t capacity ) ) {
	void**const strl_ptrs = (void**) strl;
	strl_ptrs[1] = (void*) capfunc;
}

void
strlapnd ( void* strl, const char* str ) {
	void**const strl_ptrs = (void**) strl;
	size_t*const strl_numbers = (size_t*) strl_ptrs[0];
	const size_t len = strl_numbers[INDEX_LEN];
	const size_t cap = strl_numbers[INDEX_CAP];
	const size_t str_size = strlen ( str ) + sizeof(char);
	char*const str_copy = (char*) malloc ( str_size );

	if ( str_copy ) {

		if ( len >= cap ) {
			size_t (*const capfunc) ( size_t n, size_t capacity ) = (size_t (*) ( size_t n, size_t capacity )) (strl_ptrs[1]);
			const size_t cap_new = capfunc ( 1, cap );
			char**const char_arr_new = (char**) malloc ( sizeof(char**) * cap_new );

			if ( char_arr_new ) {
				char**const char_arr = (char**) strl_ptrs[2];
				strl_ptrs[2] = (void*) char_arr_new;
				char_arr_new[len] = str_copy;
				strl_numbers[INDEX_LEN] += 1;
				strl_numbers[INDEX_ERR] = 0;

				memcpy ( str_copy, str, str_size );
				memcpy ( char_arr_new, char_arr, sizeof(char*) * len );
				free ( char_arr );

			} else {
				strl_numbers[INDEX_ERR] = 1;

				free ( str_copy );
			}

		} else {
			char**const char_arr = (char**) strl_ptrs[2];
			char_arr[len] = str_copy;
			strl_numbers[INDEX_LEN] += 1;
			strl_numbers[INDEX_ERR] = 0;

			memcpy ( str_copy, str, str_size );
		}

	} else {
		strl_numbers[INDEX_ERR] = 1;
	}
}

void
strlapndref ( void* strl, char* str ) {
	void**const strl_ptrs = (void**) strl;
	size_t*const strl_numbers = (size_t*) strl_ptrs[0];
	const size_t len = strl_numbers[INDEX_LEN];
	const size_t cap = strl_numbers[INDEX_CAP];
	const size_t str_size = strlen ( str ) + sizeof(char);

	if ( len >= cap ) {
		size_t (*const capfunc) ( size_t n, size_t capacity ) = (size_t (*) ( size_t n, size_t capacity )) (strl_ptrs[1]);
		const size_t cap_new = capfunc ( 1, cap );
		char**const char_arr_new = (char**) malloc ( sizeof(char**) * cap_new );

		if ( char_arr_new ) {
			char**const char_arr = (char**) strl_ptrs[2];
			strl_ptrs[2] = (void *) char_arr_new;
			char_arr_new[len] = str;
			strl_numbers[INDEX_LEN] += 1;
			strl_numbers[INDEX_ERR] = 0;

			memcpy ( char_arr_new, char_arr, sizeof(char*) * len );

		} else {
			strl_numbers[INDEX_ERR] = 1;
		}

	} else {
		char**const char_arr = (char**) strl_ptrs[2];
		char_arr[len] = str;
		strl_numbers[INDEX_LEN] += 1;
		strl_numbers[INDEX_ERR] = 0;
	}
}

char**
strlarr ( void* strl ) {
	void**const strl_ptrs = (void**) strl;
	char**const char_arr = (char**) strl_ptrs[2];

	return char_arr;
}

char*
strlrmref ( void* strl, size_t index ) {
	void**const strl_ptrs = (void**) strl;
	size_t*const strl_numbers = (size_t*) strl_ptrs[0];
	char**const char_arr = (char**) strl_ptrs[2];
	char*const str = char_arr[index];
	const size_t rest = ( strl_numbers[INDEX_LEN] - index - 1 ) * sizeof(char*);
	strl_numbers[INDEX_LEN] -= 1;

	memmove ( char_arr + index, char_arr + index + 1, rest );

	return str;
}

void
strlrm ( void* strl, size_t index ) {
	char* str = strlrmref ( strl, index );

	free ( str );
}

void
strlclear ( void* strl ) {
	void**const strl_ptrs = (void**) strl;
	size_t*const strl_numbers = (size_t*) strl_ptrs[0];
	char**const char_arr = (char**) strl_ptrs[2];
	const size_t len = strl_numbers[INDEX_LEN];
	size_t i;

	for ( i = 0; i < len; i += 1 ) {
		free ( char_arr[i] );
	}
	strl_numbers[INDEX_LEN] = 0;
}

void
strlclearref ( void* strl ) {
	void**const strl_ptrs = (void**) strl;
	size_t*const strl_numbers = (size_t*) strl_ptrs[0];
	strl_numbers[INDEX_LEN] = 0;
}

void
strlins ( void* strl, size_t index, const char* str ) {
	void**const strl_ptrs = (void**) strl;
	size_t*const strl_numbers = (size_t*) strl_ptrs[0];
	const size_t len = strl_numbers[INDEX_LEN];
	const size_t cap = strl_numbers[INDEX_CAP];
	const size_t str_size = strlen ( str ) + sizeof(char);
	char*const str_copy = (char*) malloc ( str_size );

	if ( str_copy ) {

		if ( len >= cap ) {
			size_t (*const capfunc) ( size_t n, size_t capacity ) = (size_t (*) ( size_t n, size_t capacity )) (strl_ptrs[1]);
			const size_t cap_new = capfunc ( 1, cap );
			char**const char_arr_new = (char**) malloc ( sizeof(char**) * cap_new );

			if ( char_arr_new ) {
				char**const char_arr = (char**) strl_ptrs[2];
				strl_ptrs[2] = (void*) char_arr_new;

				memcpy ( str_copy, str, str_size );
				memcpy ( char_arr_new, char_arr, sizeof(char*) * index );
				memcpy ( char_arr_new + index + 1, char_arr + index, sizeof(char*) * ( len - index ) );

				char_arr_new[index] = str_copy;
				strl_numbers[INDEX_LEN] += 1;
				strl_numbers[INDEX_ERR] = 0;
				free ( char_arr );

			} else {
				strl_numbers[INDEX_ERR] = 1;

				free ( str_copy );
			}

		} else {
			char**const char_arr = (char**) strl_ptrs[2];

			memmove ( char_arr + index + 1, char_arr + index, sizeof(char*) * ( len - index ) );

			char_arr[index] = str_copy;
			strl_numbers[INDEX_LEN] += 1;
			strl_numbers[INDEX_ERR] = 0;

			memcpy ( str_copy, str, str_size );
		}

	} else {
		strl_numbers[INDEX_ERR] = 1;
	}
}

void
strlinsref ( void* strl, size_t index, char* str ) {
	void**const strl_ptrs = (void**) strl;
	size_t*const strl_numbers = (size_t*) strl_ptrs[0];
	const size_t len = strl_numbers[INDEX_LEN];
	const size_t cap = strl_numbers[INDEX_CAP];
	const size_t str_size = strlen ( str ) + sizeof(char);

	if ( len >= cap ) {
		size_t (*const capfunc) ( size_t n, size_t capacity ) = (size_t (*) ( size_t n, size_t capacity )) (strl_ptrs[1]);
		const size_t cap_new = capfunc ( 1, cap );
		char**const char_arr_new = (char**) malloc ( sizeof(char**) * cap_new );

		if ( char_arr_new ) {
			char**const char_arr = (char**) strl_ptrs[2];
			strl_ptrs[2] = (void*) char_arr_new;

			memcpy ( char_arr_new, char_arr, sizeof(char*) * index );
			memcpy ( char_arr_new + index + 1, char_arr + index, sizeof(char*) * ( len - index ) );

			char_arr_new[index] = str;
			strl_numbers[INDEX_LEN] += 1;
			strl_numbers[INDEX_ERR] = 0;
			free ( char_arr );

		} else {
			strl_numbers[INDEX_ERR] = 1;
		}

	} else {
		char**const char_arr = (char**) strl_ptrs[2];

		memmove ( char_arr + index + 1, char_arr + index, sizeof(char*) * ( len - index ) );

		char_arr[index] = str;
		strl_numbers[INDEX_LEN] += 1;
		strl_numbers[INDEX_ERR] = 0;
	}
}


























































