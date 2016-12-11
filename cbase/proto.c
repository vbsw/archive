
//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


#include <stdlib.h>
#include <stdio.h>


void *
proto_fget ( const char *const fname ) {
	FILE *file = fopen ( fname, "rb" );

	if ( 0!=file ) {
		fseek ( file, 0L, SEEK_END );
		const long int fsize = ftell ( file );
		fseek ( file, 0L, SEEK_SET );
		void *const data = malloc ( fsize );
		fread ( data, 1, fsize, file );
		fclose ( file );

		return data;

	} else {
		return 0;
	}
}


void
proto_fpnd ( const char *const fname, const void *const data, const size_t size, const size_t count ) {
	FILE *file = fopen ( fname, "ab" );

	if ( 0!=file ) {
		fwrite ( data, size, count, file );
		fclose ( file );
	}
}





















