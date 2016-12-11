
//          Copyright Vitali Baumtrok 2015.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


/**
 * @file shif.test.main.c
 * @author Vitali Baumtrok
 * @date 2015-05-03
 * @brief SHiF function tests
 * 
 * Tests the functions of SHiF. To compile use for example the command
 * 'gcc shif.test.main.c shif.c -o main'.
 */

#include "shif.h"

#include <stdlib.h>
#include <stdio.h>

typedef struct
{
	int functionCalls;
	int errorCode;
	int indentDepth;
	int nodeNameOk;
	int nodeValueOk;
}
TestDataType;

static const char *const NAME_parse = "SHiF_parse";

static const char *const TEST_STRING_1 = "";
static const char *const TEST_STRING_2 = "  \nasdf asdf";
static const char *const TEST_STRING_3 = "  \nasdf\n  bla asdf2\n";
static const char *const TEST_STRING_4 = "  \n asdf";
static const char *const TEST_STRING_5 = "  \nasdf\n    hallo";
static const char *const TEST_STRING_6 = "  \nasdf \t \t asdf\t  \t ";

static int
failure
(
	const int assertTrue,
	const char *const functionName,
	const char *const message
)
{
	if ( 0 == assertTrue )
	{
		printf ( "%s : %s\n", functionName, message );
		return 1;
	}
	else
	{
		return 0;
	}
}

/**
 * @brief Compare two strings.
 * 
 * Compare two strings.
 * 
 * @param stringABegin Beginning of string A. Can be 0.
 * @param stringAEnd End of string A. Must not be 0, if string A is not 0.
 * @param stringB String B to compare. Can be 0.
 * @return 1 if equal, otherwise 0.
 */
static int
isEqual
(
	const char *stringABegin,
	const char *stringAEnd,
	const char *stringB
)
{
	if ( 0 != stringABegin && 0 != stringB )
	{
		while ( 1 )
		{
			if ( stringABegin == stringAEnd )
			{
				if ( 0 == *stringB )
					return 1;
				else
					return 0;
			}
			else
			{
				if ( *stringABegin == *stringB )
				{
					stringABegin += 1;
					stringB += 1;
				}
				else
				{
					return 0;
				}
			}
		}
	}
}

static void
printFromPtr
(
	const char *begin,
	const char *const end
)
{
	if ( 0!=begin )
	{
		printf("\"");

		while (begin != end)
		{
			printf ("%c", *begin);
			begin += 1;
		}
		printf("\"");
	}
	printf("\n");
}

static int
newNodeFct
( 
		void *data,
		int indentDepth,
		char *nodeNameBegin,
		char *nodeNameEnd,
		char *nodeValueBegin,
		char *nodeValueEnd
)
{
	TestDataType *const testData = (TestDataType *const) data;
	(*testData).functionCalls += 1;
	(*testData).indentDepth = indentDepth;

	(*testData).nodeNameOk = isEqual ( nodeNameBegin, nodeNameEnd, "asdf" ) | isEqual ( nodeNameBegin, nodeNameEnd, "bla" );
	(*testData).nodeValueOk = isEqual ( nodeValueBegin, nodeValueEnd, "asdf" );

	return 0;
}

static int
errorFct
(
	void *data,
	int errorCode,
	int errorLineNum
)
{
	return errorCode;
}

static int
test_parse
()
{
	TestDataType testData;
	testData.functionCalls = 0;
	testData.errorCode = 0;
	testData.indentDepth = 0;
	testData.nodeNameOk = 0;
	testData.nodeValueOk = 0;

	SHiF_parse ( (char*)TEST_STRING_1, (void*) &testData, newNodeFct, 0, 0 );

	if ( failure ( 0==testData.functionCalls, NAME_parse, "wrong function calls number (1)" ) ) { return 1; }
	else
	{
		testData.functionCalls = 0;
		testData.errorCode = 0;
		testData.indentDepth = 0;
		testData.nodeNameOk = 0;
		testData.nodeValueOk = 0;
	}

	SHiF_parse ( (char*)TEST_STRING_2, (void*) &testData, newNodeFct, 0, 0 );

	if ( failure ( 1==testData.functionCalls, NAME_parse, "wrong function calls number (2)" ) ) { return 1; }
	else
	{
		if ( failure ( 0==testData.indentDepth, NAME_parse, "wrong indent depth (1)" ) )
			return 1;

		if ( failure ( 1==testData.nodeNameOk, NAME_parse, "wrong node name (1)" ) )
			return 1;

		if ( failure ( 1==testData.nodeValueOk, NAME_parse, "wrong node value (1)" ) )
			return 1;

		testData.functionCalls = 0;
		testData.errorCode = 0;
		testData.indentDepth = 0;
		testData.nodeNameOk = 0;
		testData.nodeValueOk = 0;
	}

	SHiF_parse ( (char*)TEST_STRING_3, (void*) &testData, newNodeFct, 0, 0 );

	if ( failure ( 2==testData.functionCalls, NAME_parse, "wrong function calls number (3)" ) ) { return 1; }
	else
	{
		if ( failure ( 1==testData.indentDepth, NAME_parse, "wrong indent depth (2)" ) )
			return 1;

		if ( failure ( 1==testData.nodeNameOk, NAME_parse, "wrong node name (2)" ) )
			return 1;

		if ( failure ( 0==testData.nodeValueOk, NAME_parse, "wrong node value (2)" ) )
			return 1;

		testData.functionCalls = 0;
		testData.errorCode = 0;
		testData.indentDepth = 0;
		testData.nodeNameOk = 0;
		testData.nodeValueOk = 0;
	}

	int errorCode = SHiF_parse ( (char*)TEST_STRING_4, (void*) &testData, newNodeFct, errorFct, 0 );

	if ( failure ( SHIF_EROR_WRONG_NODE_NAME==errorCode, NAME_parse, "error unrecognized (1)" ) ) { return 1; }
	else
	{
		testData.functionCalls = 0;
		testData.errorCode = 0;
		testData.indentDepth = 0;
		testData.nodeNameOk = 0;
		testData.nodeValueOk = 0;
	}

	errorCode = SHiF_parse ( (char*)TEST_STRING_5, (void*) &testData, newNodeFct, errorFct, 0 );

	if ( failure ( SHIF_EROR_WRONG_INDENT_DEPTH==errorCode, NAME_parse, "error unrecognized (2)" ) ) { return 1; }
	else
	{
		testData.functionCalls = 0;
		testData.errorCode = 0;
		testData.indentDepth = 0;
		testData.nodeNameOk = 0;
		testData.nodeValueOk = 0;
	}

	errorCode = SHiF_parse ( (char*)TEST_STRING_6, (void*) &testData, newNodeFct, errorFct, SHIF_TRIM_ALL );

	if ( failure ( 0==errorCode, NAME_parse, "error occurred (1)" ) ) { return 1; }
	else
	{
		if ( failure ( 1==testData.functionCalls, NAME_parse, "wrong function calls number (4)" ) )
			return 1;

		if ( failure ( 0==testData.indentDepth, NAME_parse, "wrong indent depth (3)" ) )
			return 1;

		if ( failure ( 1==testData.nodeNameOk, NAME_parse, "wrong node name (3)" ) )
			return 1;

		if ( failure ( 1==testData.nodeValueOk, NAME_parse, "wrong node value (3)" ) )
			return 1;

		testData.functionCalls = 0;
		testData.errorCode = 0;
		testData.indentDepth = 0;
		testData.nodeNameOk = 0;
		testData.nodeValueOk = 0;
	}

	return 0;
}

int
main
(
	int argc,
	const char **argv
)
{
	if ( test_parse () )
		return 0;

	printf ( "SHiF tests passed\n" );
	return 0;
}
