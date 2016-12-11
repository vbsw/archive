
//          Copyright 2015, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


#ifndef CHARANALYZER_H
#define CHARANALYZER_H


#include <stdlib.h>


#define CharAnalyzer_CHAR_LETTERS 0
#define CharAnalyzer_CHAR_DOT 1
#define CharAnalyzer_CHAR_COMMA 2
#define CharAnalyzer_CHAR_NEW_LINE 3
#define CharAnalyzer_CHAR_DIGITS 4
#define CharAnalyzer_CHAR_BLANKS 5
#define CharAnalyzer_CHAR_TABS 6
#define CharAnalyzer_CHAR_ROUND_BRAKET_OPENED 7
#define CharAnalyzer_CHAR_ROUND_BRAKET_CLOSED 8
#define CharAnalyzer_CHAR_CURLY_BRAKET_OPENED 9
#define CharAnalyzer_CHAR_CURLY_BRAKET_CLOSED 10
#define CharAnalyzer_CHAR_ANGLE_BRAKET_OPENED 11
#define CharAnalyzer_CHAR_ANGLE_BRAKET_CLOSED 12
#define CharAnalyzer_CHAR_SQUARE_BRAKET_OPENED 13
#define CharAnalyzer_CHAR_SQUARE_BRAKET_CLOSED 14
#define CharAnalyzer_CHAR_UNDERSCORE 15
#define CharAnalyzer_CHAR_BACKSLASH 16
#define CharAnalyzer_CHAR_OTHERS 17
#define CharAnalyzer_CHAR_NONE 18


typedef struct
{
	char*    chartype;
	size_t*  extent;
	size_t*  offset;
	size_t   len;
	size_t   cap;
}
CharAnalyzer_Result_T;


extern void
CharAnalyzer_Init
(
	CharAnalyzer_Result_T *const result,
	const size_t                 capacity
);
extern void
CharAnalyzer_Destroy
(
	CharAnalyzer_Result_T *const result
);
extern void
CharAnalyzer_IncrLen
(
	CharAnalyzer_Result_T *const result
);
extern void
CharAnalyzer_Parse
(
	CharAnalyzer_Result_T *const result,
	const char *const            textchars,
	const size_t                 textsize
);


#endif /* CHARANALYZER_H */
