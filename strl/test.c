
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        <http://www.boost.org/LICENSE_1_0.txt>)


#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "strl.h"


static const char *const STR_strlnew     = "strlnew       ";
static const char *const STR_strlnewcap  = "strlnewcap    ";
static const char *const STR_strlapnd    = "strlapnd      ";
static const char *const STR_strlapndref = "strlapndref   ";
static const char *const STR_strlrm      = "strlrm        ";
static const char *const STR_strlclear   = "strlclear     ";
static const char *const STR_strlins     = "strlins       ";
static const char *const STR_strlinsref  = "strlinsref    ";


static void
print_header ( void ) {
	printf ( "\n" );
}

static void
print_footer ( void ) {
	printf ( "\n" );
}

static void
print_result ( const char *const func_name, const int error ) {
	if ( error == 0 ) {
		printf ( "%s  [passed]", func_name );
	} else {
		printf ( "%s  [failed:%d]", func_name, error );
	}
	printf ( "\n" );
}

static int
test_strlnew ( void ) {
	void * strl = strlnew ( );

	if ( strl != 0 ) {
		strldest ( strl );
		return 0;

	} else {
		return 1;
	}
}

static int
test_strlnewcap ( void ) {
	void * strl = strlnewcap ( 20 );

	if ( strl != 0 && strllen ( strl ) == 0 && strlcap ( strl ) == 20 ) {
		strldest ( strl );
		return 0;

	} else {
		return 1;
	}
}

static int
test_strlapndref ( void ) {
	void * strl = strlnewcap ( 1 );
	char * aaa = (char *) malloc ( sizeof(char) * 4 );
	char * bbb = (char *) malloc ( sizeof(char) * 4 );

	memcpy ( aaa, "aaa", 4 );
	memcpy ( bbb, "bbb", 4 );

	if ( strl == 0 ) return 1;

	strlapndref ( strl, aaa );

	if ( strlerr ( strl ) ) return 2;
	if ( strllen ( strl ) != 1 && strlcap ( strl ) != 1 ) return 3;

	strlapndref ( strl, bbb );

	if ( strlerr ( strl ) ) return 4;
	if ( strllen ( strl ) != 2 && strlcap ( strl ) < 2 ) return 5;
	if ( strcmp ( strlarr ( strl ) [0], aaa ) ) return 6;
	if ( strcmp ( strlarr ( strl ) [1], bbb ) ) return 7;

	strldest ( strl );

	return 0;
}

static int
test_strlapnd ( void ) {
	void * strl = strlnewcap ( 1 );

	if ( strl == 0 ) return 1;

	strlapnd ( strl, "aaa" );

	if ( strlerr ( strl ) ) return 2;
	if ( strllen ( strl ) != 1 && strlcap ( strl ) != 1 ) return 3;
	if ( strcmp ( strlarr ( strl ) [0], "aaa" ) ) return 4;

	strlapnd ( strl, "bbb" );

	if ( strlerr ( strl ) ) return 5;
	if ( strllen ( strl ) != 2 && strlcap ( strl ) < 2 ) return 6;
	if ( strlarr ( strl ) == 0 ) return 7;
	if ( strcmp ( strlarr ( strl ) [0], "aaa" ) ) return 8;
	if ( strcmp ( strlarr ( strl ) [1], "bbb" ) ) return 9;

	strldest ( strl );

	return 0;
}

static int
test_strlrm ( void ) {
	void * strl = strlnewcap ( 1 );

	if ( strl == 0 ) return 1;

	strlapnd ( strl, "aaa" );

	if ( strlerr ( strl ) ) return 2;
	if ( strllen ( strl ) != 1 && strlcap ( strl ) != 1 ) return 3;
	if ( strcmp ( strlarr ( strl ) [0], "aaa" ) ) return 4;

	strlapnd ( strl, "bbb" );

	if ( strlerr ( strl ) ) return 5;
	if ( strllen ( strl ) != 2 && strlcap ( strl ) < 2 ) return 6;
	if ( strlarr ( strl ) == 0 ) return 7;
	if ( strcmp ( strlarr ( strl ) [0], "aaa" ) ) return 8;
	if ( strcmp ( strlarr ( strl ) [1], "bbb" ) ) return 9;

	strlrm ( strl, 0 );

	if ( strllen ( strl ) != 1 && strlcap ( strl ) < 2 ) return 10;
	if ( strcmp ( strlarr ( strl ) [0], "bbb" ) ) return 11;

	strldest ( strl );

	return 0;
}

static int
test_strlclear ( void ) {
	void * strl = strlnewcap ( 1 );

	if ( strl == 0 ) return 1;

	strlapnd ( strl, "aaa" );

	if ( strlerr ( strl ) ) return 2;
	if ( strllen ( strl ) != 1 && strlcap ( strl ) != 1 ) return 3;
	if ( strcmp ( strlarr ( strl ) [0], "aaa" ) ) return 4;

	strlapnd ( strl, "bbb" );

	if ( strlerr ( strl ) ) return 5;
	if ( strllen ( strl ) != 2 && strlcap ( strl ) < 2 ) return 6;
	if ( strlarr ( strl ) == 0 ) return 7;
	if ( strcmp ( strlarr ( strl ) [0], "aaa" ) ) return 8;
	if ( strcmp ( strlarr ( strl ) [1], "bbb" ) ) return 9;

	strlclear ( strl );

	if ( strllen ( strl ) != 0 && strlcap ( strl ) < 2 ) return 10;

	strldest ( strl );

	return 0;
}

static int
test_strlins ( void ) {
	void * strl = strlnewcap ( 1 );

	if ( strl == 0 ) return 1;

	strlapnd ( strl, "aaa" );
	strlapnd ( strl, "ccc" );

	if ( strllen ( strl ) != 2 && strlcap ( strl ) < 2 ) return 2;

	strlins ( strl, 1, "bbb" );

	if ( strllen ( strl ) != 3 && strlcap ( strl ) < 3 ) return 3;
	if ( strcmp ( strlarr ( strl ) [0], "aaa" ) ) return 4;
	if ( strcmp ( strlarr ( strl ) [1], "bbb" ) ) return 5;
	if ( strcmp ( strlarr ( strl ) [2], "ccc" ) ) return 6;

	strldest ( strl );

	return 0;
}

static int
test_strlinsref ( void ) {
	void * strl = strlnewcap ( 1 );
	char * aaa = (char *) malloc ( sizeof(char) * 4 );
	char * bbb = (char *) malloc ( sizeof(char) * 4 );
	char * ccc = (char *) malloc ( sizeof(char) * 4 );

	memcpy ( aaa, "aaa", 4 );
	memcpy ( bbb, "bbb", 4 );
	memcpy ( ccc, "ccc", 4 );

	if ( strl == 0 ) return 1;

	strlapndref ( strl, aaa );
	strlapndref ( strl, ccc );

	if ( strllen ( strl ) != 2 && strlcap ( strl ) < 2 ) return 2;

	strlinsref ( strl, 1, bbb );

	if ( strllen ( strl ) != 3 && strlcap ( strl ) < 3 ) return 3;
	if ( strcmp ( strlarr ( strl ) [0], aaa ) ) return 4;
	if ( strcmp ( strlarr ( strl ) [1], bbb ) ) return 5;
	if ( strcmp ( strlarr ( strl ) [2], ccc ) ) return 6;

	strldest ( strl );

	return 0;
}

int
main ( int argc, char * * argv ) {
	print_header ( );

	print_result ( STR_strlnew, test_strlnew ( ) );
	print_result ( STR_strlnewcap, test_strlnewcap ( ) );
	print_result ( STR_strlapndref, test_strlapndref ( ) );
	print_result ( STR_strlapnd, test_strlapnd ( ) );
	print_result ( STR_strlrm, test_strlrm ( ) );
	print_result ( STR_strlclear, test_strlclear ( ) );
	print_result ( STR_strlins, test_strlins ( ) );
	print_result ( STR_strlinsref, test_strlinsref ( ) );

	print_footer ( );

	return 0;
}

























































