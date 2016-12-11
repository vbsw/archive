
//          Copyright 2015, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


#include <stdio.h>
#include <string.h>
#include "charanalyzer.c"


#define PRINT_RESULT(FUNC_NAME) if ( error_ ## FUNC_NAME == 0 ) printf( "passed\n" ); else printf( "failed(%i)\n", error_ ## FUNC_NAME );


static int error_CharAnalyzer_Init    = 0;
static int error_CharAnalyzer_Destroy = 0;
static int error_CharAnalyzer_IncrLen = 0;
static int error_CharAnalyzer_Parse   = 0;


static void
Helper_PrintResult
(
	CharAnalyzer_Result_T* result
)
{
	printf
	(
		"(%i, %i, %i, %i, %i)\n",
		(int)((*result).len),
		(int)((*result).cap),
		(int)((*result).chartype!=0),
		(int)((*result).extent!=0),
		(int)((*result).offset!=0)
	);
}
static void
Helper_InitResult
(
	CharAnalyzer_Result_T* result,
	size_t                 length,
	size_t                 capacity
)
{
	(*result).len      = length;
	(*result).cap      = capacity;
	(*result).chartype = (char*) malloc( sizeof(char) * capacity );
	(*result).extent   = (size_t*) malloc( sizeof(size_t) * capacity );
	(*result).offset   = (size_t*) malloc( sizeof(size_t) * capacity );
}
static void
Helper_DestroyResult
(
	CharAnalyzer_Result_T* result
)
{
	if ( (*result).chartype != 0 )
		free ( (*result).chartype );
	if ( (*result).extent != 0 )
		free ( (*result).extent );
	if ( (*result).offset != 0 )
		free ( (*result).offset );
	(*result).len      = 0;
	(*result).cap      = 0;
	(*result).chartype = 0;
	(*result).extent   = 0;
	(*result).offset   = 0;
}
static void
Test_CharAnalyzer_Init ( )
{
	const size_t capacity = 40;
	CharAnalyzer_Result_T result;
	result.len = 10000;
	CharAnalyzer_Init( &result, capacity );

	if ( result.len != 0 ) {
		error_CharAnalyzer_IncrLen = 1; return;
	}
	if ( result.chartype == 0 ) {
		error_CharAnalyzer_Init = 1; return;
	}
	if ( result.extent == 0 ) {
		error_CharAnalyzer_Init = 2; return;
	}
	if ( result.offset == 0 ) {
		error_CharAnalyzer_Init = 3; return;
	}
	if ( result.cap != capacity ) {
		error_CharAnalyzer_Init = 4; return;
	}
}
static void
Test_CharAnalyzer_Destroy ( )
{
	const size_t length   = 8;
	const size_t capacity = 16;
	CharAnalyzer_Result_T result;
	Helper_InitResult( &result, length, capacity );
	CharAnalyzer_Destroy( &result );

	if ( result.len != 0 ) {
		error_CharAnalyzer_Destroy = 1; return;
	}
	if ( result.cap != 0 ) {
		error_CharAnalyzer_Destroy = 2; return;
	}
	if ( result.chartype != 0 ) {
		error_CharAnalyzer_Destroy = 3; return;
	}
	if ( result.extent != 0 ) {
		error_CharAnalyzer_Destroy = 4; return;
	}
	if ( result.offset != 0 ) {
		error_CharAnalyzer_Destroy = 5; return;
	}
}
static void
Test_CharAnalyzer_IncrLen ( )
{
	const size_t length   = 8;
	const size_t capacity = 16;
	CharAnalyzer_Result_T result;
	Helper_InitResult( &result, length, capacity );

	CharAnalyzer_IncrLen( &result );

	if ( result.len != length + 1 ) {
		printf("dann: %i, %i\n", (int)result.len, (int)(length + 1));
		error_CharAnalyzer_IncrLen = 1; return;
	}
	if ( result.cap != capacity ) {
		error_CharAnalyzer_IncrLen = 2; return;
	}
	if ( result.chartype == 0 ) {
		error_CharAnalyzer_IncrLen = 3; return;
	}
	if ( result.extent == 0 ) {
		error_CharAnalyzer_IncrLen = 4; return;
	}
	if ( result.offset == 0 ) {
		error_CharAnalyzer_IncrLen = 5; return;
	}

	result.len = capacity;
	CharAnalyzer_IncrLen( &result );

	if ( result.len != capacity + 1 ) {
		error_CharAnalyzer_IncrLen = 6; return;
	}
	if ( result.cap <= capacity ) {
		error_CharAnalyzer_IncrLen = 7; return;
	}
	if ( result.chartype == 0 ) {
		error_CharAnalyzer_IncrLen = 8; return;
	}
	if ( result.extent == 0 ) {
		error_CharAnalyzer_IncrLen = 9; return;
	}
	if ( result.offset == 0 ) {
		error_CharAnalyzer_IncrLen = 10; return;
	}

	Helper_DestroyResult( &result );
}
static void
Test_CharAnalyzer_Parse ( )
{
	const size_t length   = 8;
	const size_t capacity = 16;
	const char text[28] = "a.,\r\n0 \t(){}[]<>_@aaa\\\n\r123";
	CharAnalyzer_Result_T result;
	Helper_InitResult( &result, length, capacity );

	CharAnalyzer_Parse( &result, text, sizeof(text) - 1 );

	if ( result.len != 22 ) {
		error_CharAnalyzer_Parse = 1; return;
	}
	if ( result.chartype[0] != CharAnalyzer_CHAR_LETTERS ) {
		error_CharAnalyzer_Parse = 2; return;
	}
	if ( result.extent[0] != 1 ) {
		error_CharAnalyzer_Parse = 3; return;
	}
	if ( result.chartype[1] != CharAnalyzer_CHAR_DOT ) {
		error_CharAnalyzer_Parse = 4; return;
	}
	if ( result.chartype[2] != CharAnalyzer_CHAR_COMMA ) {
		error_CharAnalyzer_Parse = 5; return;
	}
	if ( result.chartype[3] != CharAnalyzer_CHAR_NEW_LINE ) {
		error_CharAnalyzer_Parse = 6; return;
	}
	if ( result.extent[3] != 2 ) {
		error_CharAnalyzer_Parse = 7; return;
	}
	if ( result.chartype[4] != CharAnalyzer_CHAR_DIGITS ) {
		error_CharAnalyzer_Parse = 8; return;
	}
	if ( result.chartype[5] != CharAnalyzer_CHAR_BLANKS ) {
		error_CharAnalyzer_Parse = 9; return;
	}
	if ( result.chartype[6] != CharAnalyzer_CHAR_TABS ) {
		error_CharAnalyzer_Parse = 10; return;
	}
	if ( result.chartype[7] != CharAnalyzer_CHAR_ROUND_BRAKET_OPENED ) {
		error_CharAnalyzer_Parse = 11; return;
	}
	if ( result.chartype[8] != CharAnalyzer_CHAR_ROUND_BRAKET_CLOSED ) {
		error_CharAnalyzer_Parse = 12; return;
	}
	if ( result.chartype[9] != CharAnalyzer_CHAR_CURLY_BRAKET_OPENED ) {
		error_CharAnalyzer_Parse = 13; return;
	}
	if ( result.chartype[10] != CharAnalyzer_CHAR_CURLY_BRAKET_CLOSED ) {
		error_CharAnalyzer_Parse = 14; return;
	}
	if ( result.chartype[11] != CharAnalyzer_CHAR_ANGLE_BRAKET_OPENED ) {
		error_CharAnalyzer_Parse = 15; return;
	}
	if ( result.chartype[12] != CharAnalyzer_CHAR_ANGLE_BRAKET_CLOSED ) {
		error_CharAnalyzer_Parse = 16; return;
	}
	if ( result.chartype[13] != CharAnalyzer_CHAR_SQUARE_BRAKET_OPENED ) {
		error_CharAnalyzer_Parse = 17; return;
	}
	if ( result.chartype[14] != CharAnalyzer_CHAR_SQUARE_BRAKET_CLOSED ) {
		error_CharAnalyzer_Parse = 18; return;
	}
	if ( result.chartype[15] != CharAnalyzer_CHAR_UNDERSCORE ) {
		error_CharAnalyzer_Parse = 19; return;
	}
	if ( result.chartype[16] != CharAnalyzer_CHAR_OTHERS ) {
		error_CharAnalyzer_Parse = 20; return;
	}
	if ( result.chartype[17] != CharAnalyzer_CHAR_LETTERS ) {
		error_CharAnalyzer_Parse = 21; return;
	}
	if ( result.extent[17] != 3 ) {
		error_CharAnalyzer_Parse = 22; return;
	}
	if ( result.chartype[18] != CharAnalyzer_CHAR_BACKSLASH ) {
		error_CharAnalyzer_Parse = 23; return;
	}
	if ( result.chartype[19] != CharAnalyzer_CHAR_NEW_LINE ) {
		error_CharAnalyzer_Parse = 24; return;
	}
	if ( result.extent[19] != 1 ) {
		error_CharAnalyzer_Parse = 25; return;
	}
	if ( result.chartype[20] != CharAnalyzer_CHAR_OTHERS ) {
		error_CharAnalyzer_Parse = 26; return;
	}
	if ( result.chartype[21] != CharAnalyzer_CHAR_DIGITS ) {
		error_CharAnalyzer_Parse = 27; return;
	}
	if ( result.extent[21] != 3 ) {
		error_CharAnalyzer_Parse = 28; return;
	}

	Helper_DestroyResult( &result );
}
int
main ( int argc, const char** argv )
{
	printf( "Character Analyzer - Function Tests\n" );
	printf( "-----------------------------------\n" );

	printf( "CharAnalyzer_Init            " );
	Test_CharAnalyzer_Init();
	PRINT_RESULT( CharAnalyzer_Init );

	printf( "CharAnalyzer_Destroy         " );
	Test_CharAnalyzer_Destroy();
	PRINT_RESULT( CharAnalyzer_Destroy );

	printf( "CharAnalyzer_IncrLen         " );
	Test_CharAnalyzer_IncrLen();
	PRINT_RESULT( CharAnalyzer_IncrLen );

	printf( "CharAnalyzer_Parse           " );
	Test_CharAnalyzer_Parse();
	PRINT_RESULT( CharAnalyzer_Parse );

	return 0;
}

