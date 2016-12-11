
//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <errno.h>
#include <sys/time.h>

#include "proto.h"
#include "strmapbuf.h"


static const char *const STR_OUT_FILE = "output.txt";
static const char *const STR_ANY = "any";
static const char *const STR_ANYANY = "anyany";
static const char *const STR_BIG_ANY = "anyfqwerjqwfsadfooiuqzouizousdzfouiqzwoeufzoduzasofuzoqiuwezouifoiuaszfouizqowuiezfiouzasdiufzoqiuwzefoiuzaiusdfzoiuqzweoiufzoiaduszfoiuqzweoiufzoiauszdfuzwef";
static const char *const STR_UNEXPECTED_OUT = "Unexpected output";
static const char *const STR_WRONG_OUTPUT = "wrong output:";
static const char *const STR_COULD_NOT_WRITE_F = "Could not write to file:";
static const char *const STR_PROPERTIES = "key1=value1\n\nkey2=value2\n key3=value3\nkey4=value4";

static const char *const STR_KEY1 = "key1";
static const char *const STR_KEY4 = "key4";
static const char *const STR_VALUE1 = "value1";

static const int SIZE_STR_ANY = 3;
static const int SIZE_STR_BIG_ANY = 158;
static const int SIZE_STR_PROPERTIES = 49;


void printerr () {
	printf ( "  errno %i - %s\n", errno, strerror ( errno ) );
}


static const char *const STR_TEST_FGET_ERROR = "proto_fget - error";
static const char *const STR_TEST_FGET_FAILED = "proto_fget - failed";

void test_proto_fget () {
	FILE *file = fopen ( STR_OUT_FILE, "wb" );
	
	if ( 0!=file ) {

		if ( fwrite ( STR_ANY, 1, 3, file ) ) {

			if ( 0==fflush ( file ) ) {
				char *data = (char*) proto_fget ( STR_OUT_FILE );

				if ( 0==strcmp ( data, STR_ANY ) ) {
					printf ( "proto_fget         - ok\n" );

				} else {
					printf ( "%s\n", STR_TEST_FGET_FAILED );
					printf ( "  %s %s\n", STR_WRONG_OUTPUT, data );
				}
				free ( data );

			} else {
				printf ( "%s\n", STR_TEST_FGET_ERROR );
				printf ( "  %s %s\n", STR_COULD_NOT_WRITE_F, STR_OUT_FILE );
			}

		} else {
			printf ( "%s\n", STR_TEST_FGET_ERROR );
			printf ( "  %s %s\n", STR_COULD_NOT_WRITE_F, STR_OUT_FILE );
		}
		fclose ( file );
		remove ( STR_OUT_FILE );

	} else {
		printf ( "%s\n", STR_TEST_FGET_ERROR );
		printerr ();
	}
}


static const char *const STR_TEST_FPND_ERROR = "proto_fpnd  - error";
static const char *const STR_TEST_FPND_FAILED = "proto_fpnd  - failed";

void test_proto_fpnd () {
	remove ( STR_OUT_FILE );
	proto_fpnd ( STR_OUT_FILE, STR_ANY, 1, 3 );
	proto_fpnd ( STR_OUT_FILE, STR_ANY, 1, 3 );
	char *data = (char*) proto_fget ( STR_OUT_FILE );

	if ( 0!=data ) {

		if ( 0==strcmp ( data, STR_ANYANY ) ) {
			printf ( "proto_fpnd         - ok\n" );

		} else {
			printf ( "%s\n", STR_TEST_FPND_FAILED );
			printf ( "  %s %s\n", STR_WRONG_OUTPUT, data );
		}
		free ( data );

	} else {
		printf ( "%s\n", STR_TEST_FPND_ERROR );
		printerr ();
	}
	remove ( STR_OUT_FILE );
}


void test_List_search () {
	List list;
	List_allocate ( &list );
	List_append ( &list, (void*)STR_ANY, SIZE_STR_ANY );
	List_append ( &list, (void*)STR_BIG_ANY, SIZE_STR_BIG_ANY );

	if ( 1==List_search ( &list, (void*)STR_BIG_ANY, SIZE_STR_BIG_ANY ) )
		printf ( "List_search        - ok\n" );
	else
		printf ( "List_search        - failed\n" );
}


static const char *const STR_TEST_LIST_SEARCH_ASC_FAILED = "List_searchAsc   - failed";

void test_List_searchAsc () {
	List list;
	List_allocate ( &list );
	List_append ( &list, (void*)"a", 1 );
	size_t index;

	if ( List_searchAsc ( &list, &index, (void*)"b", 1 ) ) {
		printf ( "%s\n", STR_TEST_LIST_SEARCH_ASC_FAILED );
		printf ( "  found missing element at index: %i\n", (int)index );

	} else if ( 1!=index ) {
		printf ( "%s\n", STR_TEST_LIST_SEARCH_ASC_FAILED );
		printf ( "  1. wrong index: %i\n", (int)index );

	} else {
		List_append ( &list, (void*)"c", 1 );

		if ( List_searchAsc ( &list, &index, (void*)"b", 1 ) ) {
			printf ( "%s\n", STR_TEST_LIST_SEARCH_ASC_FAILED );
			printf ( "  found missing element at index: %i\n", (int)index );

		} else if ( 1!=index ) {
			printf ( "%s\n", STR_TEST_LIST_SEARCH_ASC_FAILED );
			printf ( "  2. wrong index: %i\n", (int)index );

		} else {
			List_append ( &list, (void*)"e", 1 );

			if ( !List_searchAsc ( &list, &index, (void*)"c", 1 ) ) {
				printf ( "%s\n", STR_TEST_LIST_SEARCH_ASC_FAILED );
				printf ( "  element not found, instead index: %i\n", (int)index );

			} else if ( 1!=index ) {
				printf ( "%s\n", STR_TEST_LIST_SEARCH_ASC_FAILED );
				printf ( "  3. wrong index: %i\n", (int)index );

			} else {
				printf ( "List_searchAsc     - ok\n" );
			}
		}
	}
}


static const char *const STR_TEST_LIST_SEARCH_DESC_FAILED = "List_searchDesc  - failed";

void test_List_searchDesc () {
	List list;
	List_allocate ( &list );
	List_append ( &list, (void*)"e", 1 );
	size_t index;

	if ( List_searchDesc ( &list, &index, (void*)"b", 1 ) ) {
		printf ( "%s\n", STR_TEST_LIST_SEARCH_DESC_FAILED );
		printf ( "  found missing element at index: %i\n", (int)index );

	} else if ( 1!=index ) {
		printf ( "%s\n", STR_TEST_LIST_SEARCH_DESC_FAILED );
		printf ( "  1: wrong index: %i\n", (int)index );

	} else {
		List_append ( &list, (void*)"c", 1 );

		if ( List_searchDesc ( &list, &index, (void*)"b", 1 ) ) {
			printf ( "%s\n", STR_TEST_LIST_SEARCH_DESC_FAILED );
			printf ( "  found missing element at index: %i\n", (int)index );

		} else if ( 2!=index ) {
			printf ( "%s\n", STR_TEST_LIST_SEARCH_DESC_FAILED );
			printf ( "  2: wrong index: %i\n", (int)index );

		} else {
			List_append ( &list, (void*)"a", 1 );

			if ( !List_searchDesc ( &list, &index, (void*)"c", 1 ) ) {
				printf ( "%s\n", STR_TEST_LIST_SEARCH_DESC_FAILED );
				printf ( "  element not found, instead index: %i\n", (int)index );

			} else if ( 1!=index ) {
				printf ( "%s\n", STR_TEST_LIST_SEARCH_DESC_FAILED );
				printf ( "  3: wrong index: %i\n", (int)index );

			} else {
				printf ( "List_searchDesc    - ok\n" );
			}
		}
	}
}


static const char *const STR_TEST_STRMAP_READ_ERROR = "StrMapBuf_readFile  - error";
static const char *const STR_TEST_STRMAP_READ_FAILED = "StrMapBuf_readFile  - failed";

void test_StrMapBuf_read () {
	StrMapBuf map;
	StrMapBuf_allocate ( &map );
	FILE *file = fopen ( STR_OUT_FILE, "w" );
	
	if ( 0!=file ) { 

		if ( fwrite ( STR_PROPERTIES, 1, SIZE_STR_PROPERTIES, file ) ) {

			if ( 0==fflush ( file ) ) {
				FILE *fileR = fopen ( STR_OUT_FILE, "r" );
				StrMapBuf_readFile ( &map, fileR );
				
				if ( 3==map.strMap.keys.len &&
						0==strcmp ( (char*) map.strMap.keys.elements[0], STR_KEY1 ) &&
						0==strcmp ( (char*) map.strMap.values.elements[0], STR_VALUE1 ) &&
						0==strcmp ( (char*) map.strMap.keys.elements[2], STR_KEY4 ) ) {
					printf ( "StrMapBuf_readFile - ok\n" );
				} else {
					printf ( "%s\n", STR_TEST_STRMAP_READ_FAILED );
					printf ( "  values don't match\n" );
				}
				fclose ( fileR );

			} else {
				printf ( "%s\n", STR_TEST_STRMAP_READ_ERROR );
				printf ( "  1: %s %s\n", STR_COULD_NOT_WRITE_F, STR_OUT_FILE );
			}

		} else {
			printf ( "%s\n", STR_TEST_STRMAP_READ_ERROR );
			printf ( "  2: %s %s\n", STR_COULD_NOT_WRITE_F, STR_OUT_FILE );
		}
		fclose ( file );
		remove ( STR_OUT_FILE );

	} else {
		printf ( "%s\n", STR_TEST_STRMAP_READ_ERROR );
		printerr ();
	}
	StrMapBuf_deallocate ( &map );
}


int main ( int argc, char **argv ) {
	test_proto_fget ();
	test_proto_fpnd ();
	test_List_search ();
	test_List_searchAsc ();
	test_List_searchDesc ();
	test_StrMapBuf_read ();

	return 0;
}




























































