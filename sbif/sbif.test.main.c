
//          Copyright Vitali Baumtrok 2015.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


/**
 * @file sbif.test.main.c
 * @author Vitali Baumtrok
 * @date 2015-06-27
 * @brief SBiF function tests
 * 
 * Tests the functions of SBiF. To compile use for example the command
 * 'gcc sbif.test.main.c sbif.c -o main'.
 */

#include "sbif.h"

#include <stdlib.h>
#include <stdio.h>


static const char *const FN_encodeSIZE = "SBiF_encodeSIZE";
static const char *const FN_decodeSIZE = "SBiF_decodeSIZE";
static const char *const FN_encodeSIZEArray = "SBiF_encodeSIZEArray";
static const char *const FN_decodeSIZEArray = "SBiF_decodeSIZEArray";


static int
failure
(
    const int assertTrue,
    const char *const functionName,
    const int failureId
)
{
    if ( 0 == assertTrue )
    {
        printf ( "%s : failure (%i)\n", functionName, failureId );
        return 1;
    }
    else
    {
        return 0;
    }
}


static int
test_encodeSIZE ()
{
    unsigned char arr[8] = { 0, 0, 0, 0, 0, 0, 0, 0 };
    size_t SIZENum = 257;
	int failc = 1;

    SBiF_encodeSIZE ( arr, SIZENum );
    if ( failure ( arr[0]==0 && arr[1]==0 && arr[2]==0 && arr[3]==0 && arr[4]==0 && arr[5]==0 && arr[6]==1 && arr[7]==1, FN_encodeSIZE, failc++ ) ) { return 1; }
    else
    {
        arr[0] = 0;
        arr[1] = 0;
        arr[2] = 0;
        arr[3] = 0;
        arr[4] = 0;
        arr[5] = 0;
        arr[6] = 0;
        arr[7] = 0;
		SIZENum = 65792;
    }

    SBiF_encodeSIZE ( arr, SIZENum );
    if ( failure ( arr[0]==0 && arr[1]==0 && arr[4]==0 && arr[5]==1 && arr[6]==1 && arr[7]==0, FN_encodeSIZE, failc++ ) ) { return 1; }
    else
    {
        arr[0] = 0;
        arr[1] = 0;
        arr[2] = 0;
        arr[3] = 0;
        arr[4] = 0;
        arr[5] = 0;
        arr[6] = 0;
        arr[7] = 0;
		SIZENum = 0;
    }

    SBiF_encodeSIZE ( arr, SIZENum );
    if ( failure ( arr[0]==0 && arr[1]==0 && arr[4]==0 && arr[5]==0 && arr[6]==0 && arr[7]==0, FN_encodeSIZE, failc++ ) ) { return 1; }

    return 0;
}


static int
test_decodeSIZE ()
{
    unsigned char arr[8] = { 0, 0, 0, 0, 0, 1, 0, 1 };
    size_t SIZENum = 0;
	int failc = 1;

    SBiF_decodeSIZE ( &SIZENum, arr );
    if ( failure ( SIZENum==65537, FN_decodeSIZE, failc++ ) ) { return 1; }
    else
    {
        arr[0] = 1;
        arr[1] = 0;
        arr[2] = 0;
        arr[3] = 0;
        arr[4] = 0;
        arr[5] = 0;
        arr[6] = 0;
        arr[7] = 0;
		SIZENum = 0;
    }

    SBiF_decodeSIZE ( &SIZENum, arr );
    if ( failure ( SIZENum>>56==1, FN_decodeSIZE, failc++ ) ) { return 1; }
    else
    {
        arr[0] = 0;
        arr[1] = 0;
        arr[2] = 0;
        arr[3] = 0;
        arr[4] = 0;
        arr[5] = 0;
        arr[6] = 0;
        arr[7] = 0;
		SIZENum = 0;
    }

    SBiF_decodeSIZE ( &SIZENum, arr );
    if ( failure ( SIZENum==0, FN_decodeSIZE, failc++ ) ) { return 1; }

    return 0;
}


static int
test_encodeSIZEArray ()
{
    unsigned char arr[8] = { 0, 0, 0, 0, 0, 0, 0, 0 };
    size_t SIZENum = 257;
	int failc = 1;

    SBiF_encodeSIZEArray ( arr, &SIZENum, 8 );
    if ( failure ( arr[0]==0 && arr[1]==0 && arr[5]==0 && arr[6]==1 && arr[7]==1, FN_encodeSIZE, failc++ ) ) { return 1; }
    else
    {
        arr[0] = 0;
        arr[1] = 0;
        arr[2] = 0;
        arr[3] = 0;
        arr[4] = 0;
        arr[5] = 0;
        arr[6] = 0;
        arr[7] = 0;
		SIZENum = 65792;
    }

    SBiF_encodeSIZEArray ( arr, &SIZENum, 8 );
    if ( failure ( arr[0]==0 && arr[1]==0 && arr[4]==0 && arr[5]==1 && arr[6]==1 && arr[7]==0, FN_encodeSIZE, failc++ ) ) { return 1; }
    else
    {
        arr[0] = 0;
        arr[1] = 0;
        arr[2] = 0;
        arr[3] = 0;
        arr[4] = 0;
        arr[5] = 0;
        arr[6] = 0;
        arr[7] = 0;
		SIZENum = 0;
    }

    SBiF_encodeSIZEArray ( arr, &SIZENum, 8 );
    if ( failure ( arr[0]==0 && arr[1]==0 && arr[4]==0 && arr[5]==0 && arr[6]==0 && arr[7]==0, FN_encodeSIZE, failc++ ) ) { return 1; }
    else
    {
        arr[0] = 0;
        arr[1] = 0;
        arr[2] = 0;
        arr[3] = 0;
        arr[4] = 0;
        arr[5] = 0;
        arr[6] = 0;
        arr[7] = 0;
		SIZENum = 65792;
    }

    SBiF_encodeSIZEArray ( arr, &SIZENum, 6 );
    if ( failure ( arr[0]==0 && arr[1]==0 && arr[4]==0 && arr[5]==1 && arr[6]==0 && arr[7]==0, FN_encodeSIZE, failc++ ) ) { return 1; }

    return 0;
}


static int
test_decodeSIZEArray ()
{
    unsigned char arr[8] = { 0, 0, 0, 0, 0, 1, 0, 1 };
    size_t SIZENum = 0;
	int failc = 1;

    SBiF_decodeSIZEArray ( &SIZENum, arr, 8 );
    if ( failure ( SIZENum==65537, FN_decodeSIZE, failc++ ) ) { return 1; }
    else
    {
        arr[0] = 1;
        arr[1] = 0;
        arr[2] = 0;
        arr[3] = 0;
        arr[4] = 0;
        arr[5] = 0;
        arr[6] = 0;
        arr[7] = 0;
		SIZENum = 0;
    }

    SBiF_decodeSIZEArray ( &SIZENum, arr, 8 );
    if ( failure ( SIZENum>>56==1, FN_decodeSIZE, failc++ ) ) { return 1; }
    else
    {
        arr[0] = 0;
        arr[1] = 0;
        arr[2] = 0;
        arr[3] = 0;
        arr[4] = 0;
        arr[5] = 0;
        arr[6] = 0;
        arr[7] = 0;
		SIZENum = 0;
    }

    SBiF_decodeSIZEArray ( &SIZENum, arr, 8 );
    if ( failure ( SIZENum==0, FN_decodeSIZE, failc++ ) ) { return 1; }
    else
    {
        arr[0] = 0;
        arr[1] = 0;
        arr[2] = 0;
        arr[3] = 0;
        arr[4] = 0;
        arr[5] = 1;
        arr[6] = 0;
        arr[7] = 1;
		SIZENum = 0;
    }

    SBiF_decodeSIZEArray ( &SIZENum, arr, 6 );
    if ( failure ( SIZENum==65536, FN_decodeSIZE, failc++ ) ) { return 1; }

    return 0;
}


int
main
(
    int argc,
    const char **argv
)
{
    if ( test_encodeSIZE () || test_decodeSIZE () || test_encodeSIZEArray () || test_decodeSIZEArray () )
    {
        return 0;
    }
    else
    {
        printf ( "SBiF tests passed\n" );
        return 0;
    }
}
