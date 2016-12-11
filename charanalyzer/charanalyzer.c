
//          Copyright 2015, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


#include "charanalyzer.h"


static inline void
Private_AddCharType
(
	CharAnalyzer_Result_T *const result,
	const size_t i,
	const size_t extent,
	const char chartype
)
{
	CharAnalyzer_IncrLen( result );
	(*result).chartype[(*result).len - 1] = chartype;
	(*result).extent[(*result).len - 1]   = extent;
	(*result).offset[(*result).len - 1]   = i;
}
static inline void
Private_AddCharTypeOrIncrExtent
(
	CharAnalyzer_Result_T *const result,
	const size_t i,
	const size_t extent,
	const char chartype
)
{
	if ( (*result).chartype[(*result).len - 1] == chartype )
	{
		(*result).extent[(*result).len - 1]++;
	}
	else
	{
		Private_AddCharType( result, i, extent, chartype);
	}
}
void
CharAnalyzer_Init
(
	CharAnalyzer_Result_T *const result,
	const size_t                 capacity
)
{
	(*result).len      = 0;
	(*result).cap      = 0;
	(*result).chartype = (char*) malloc( sizeof(char) * capacity );
	(*result).extent   = 0;
	(*result).offset   = 0;

	if ( (*result).chartype != 0 )
	{
		(*result).extent   = (size_t*) malloc( sizeof(size_t) * capacity );

		if ( (*result).extent != 0 )
		{
			(*result).offset   = (size_t*) malloc( sizeof(size_t) * capacity );

			if ( (*result).offset != 0 )
			{
				(*result).len = 0;
				(*result).cap = capacity;
			}
			else
			{
				free( (*result).chartype );
				free( (*result).extent );
				(*result).chartype = 0;
				(*result).extent   = 0;
			}
		}
		else
		{
			free( (*result).chartype );
			(*result).chartype = 0;
		}
	}
}
void
CharAnalyzer_Destroy
(
	CharAnalyzer_Result_T *const result
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
void
CharAnalyzer_IncrLen
(
	CharAnalyzer_Result_T *const result
)
{
	(*result).len++;

	if ( (*result).len > (*result).cap )
	{
		size_t new_capacity = (*result).cap * 2;
		char*  new_chartype = (char*) malloc( new_capacity * sizeof(char) );

		if ( new_chartype != 0 )
		{
			size_t* new_extent = (size_t*) malloc( new_capacity * sizeof(size_t) );

			if ( new_extent != 0 )
			{
				size_t* new_offset = (size_t*) malloc( new_capacity * sizeof(size_t) );

				if ( new_offset != 0 )
				{
					size_t old_length = (*result).len - 1;
					memcpy( new_chartype, (*result).chartype, old_length * sizeof(char) );
					memcpy( new_extent, (*result).extent, old_length * sizeof(size_t) );
					memcpy( new_offset, (*result).offset, old_length * sizeof(size_t) );
					free( (*result).chartype );
					free( (*result).extent );
					free( (*result).offset );
					(*result).cap = new_capacity;
					(*result).chartype = new_chartype;
					(*result).extent   = new_extent;
					(*result).offset   = new_offset;
				}
				else
				{
					free( new_chartype );
					free( new_extent );
				}
			}
			else
			{
				free( new_chartype );
			}
		}
	}
}
void
CharAnalyzer_Parse
(
	CharAnalyzer_Result_T *const result,
	const char *const            textchars,
	const size_t                 textsize
)
{
	size_t i      = 0;
	(*result).len = 0;

	if ( i < textsize )
	{
		char c              = textchars[0];
		(*result).extent[0] = 1;
		(*result).offset[0] = 0;


		/* parse first character */

		if ( (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') )
		{
			(*result).chartype[0] = CharAnalyzer_CHAR_LETTERS;
		}
		else if ( c == '.' )
		{
			(*result).chartype[0] = CharAnalyzer_CHAR_DOT;
		}
		else if ( c == ',' )
		{
			(*result).chartype[0] = CharAnalyzer_CHAR_COMMA;
		}
		else if ( c == '\r' )
		{
			i++;

			if ( (i < textsize) && (textchars[i] == '\n') )
			{
				(*result).chartype[0] = CharAnalyzer_CHAR_NEW_LINE;
				(*result).extent[0]   = 2;
			}
			else
			{
				i--;
				(*result).chartype[0] = CharAnalyzer_CHAR_OTHERS;
			}
		}
		else if ( c == '\n' )
		{
			(*result).chartype[0] = CharAnalyzer_CHAR_NEW_LINE;
		}
		else if ( c >= '0' && c <= '9' )
		{
			(*result).chartype[0] = CharAnalyzer_CHAR_DIGITS;
		}
		else if ( c == ' ' )
		{
			(*result).chartype[0] = CharAnalyzer_CHAR_BLANKS;
		}
		else if ( c == '\t' )
		{
			(*result).chartype[0] = CharAnalyzer_CHAR_TABS;
		}
		else if ( c == '(' )
		{
			(*result).chartype[0] = CharAnalyzer_CHAR_ROUND_BRAKET_OPENED;
		}
		else if ( c == ')' )
		{
			(*result).chartype[0] = CharAnalyzer_CHAR_ROUND_BRAKET_CLOSED;
		}
		else if ( c == '{' )
		{
			(*result).chartype[0] = CharAnalyzer_CHAR_CURLY_BRAKET_OPENED;
		}
		else if ( c == '}' )
		{
			(*result).chartype[0] = CharAnalyzer_CHAR_CURLY_BRAKET_CLOSED;
		}
		else if ( c == '[' )
		{
			(*result).chartype[0] = CharAnalyzer_CHAR_ANGLE_BRAKET_OPENED;
		}
		else if ( c == ']' )
		{
			(*result).chartype[0] = CharAnalyzer_CHAR_ANGLE_BRAKET_CLOSED;
		}
		else if ( c == '<' )
		{
			(*result).chartype[0] = CharAnalyzer_CHAR_SQUARE_BRAKET_OPENED;
		}
		else if ( c == '>' )
		{
			(*result).chartype[0] = CharAnalyzer_CHAR_SQUARE_BRAKET_CLOSED;
		}
		else if ( c == '_' )
		{
			(*result).chartype[0] = CharAnalyzer_CHAR_UNDERSCORE;
		}
		else
		{
			(*result).chartype[0] = CharAnalyzer_CHAR_OTHERS;
		}
		CharAnalyzer_IncrLen( result );


		/* parse rest */

		for ( i++; i < textsize; i++ )
		{
			char cPrev = c;
			c          = textchars[i];

			if ( (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') )
			{
				Private_AddCharTypeOrIncrExtent( result, i, 1, CharAnalyzer_CHAR_LETTERS );
			}
			else if ( c == '.' )
			{
				Private_AddCharType( result, i, 1, CharAnalyzer_CHAR_DOT );
			}
			else if ( c == ',' )
			{
				Private_AddCharType( result, i, 1, CharAnalyzer_CHAR_COMMA );
			}
			else if ( c == '\r' )
			{
				i++;

				if ( (i < textsize) && (textchars[i] == '\n') )
				{
					Private_AddCharType( result, i - 1, 2, CharAnalyzer_CHAR_NEW_LINE );
				}
				else
				{
					i--;
					Private_AddCharTypeOrIncrExtent( result, i, 1, CharAnalyzer_CHAR_OTHERS );
				}
			}
			else if ( c == '\n' )
			{
				Private_AddCharType( result, i, 1, CharAnalyzer_CHAR_NEW_LINE );
			}
			else if ( (c >= '0') && (c <= '9') )
			{
				Private_AddCharTypeOrIncrExtent( result, i, 1, CharAnalyzer_CHAR_DIGITS );
			}
			else if ( c == ' ' )
			{
				Private_AddCharTypeOrIncrExtent( result, i, 1, CharAnalyzer_CHAR_BLANKS );
			}
			else if ( c == '\t' )
			{
				Private_AddCharTypeOrIncrExtent( result, i, 1, CharAnalyzer_CHAR_TABS );
			}
			else if ( c == '(' )
			{
				Private_AddCharType( result, i, 1, CharAnalyzer_CHAR_ROUND_BRAKET_OPENED );
			}
			else if ( c == ')' )
			{
				Private_AddCharType( result, i, 1, CharAnalyzer_CHAR_ROUND_BRAKET_CLOSED );
			}
			else if ( c == '{' )
			{
				Private_AddCharType( result, i, 1, CharAnalyzer_CHAR_CURLY_BRAKET_OPENED );
			}
			else if ( c == '}' )
			{
				Private_AddCharType( result, i, 1, CharAnalyzer_CHAR_CURLY_BRAKET_CLOSED );
			}
			else if ( c == '[' )
			{
				Private_AddCharType( result, i, 1, CharAnalyzer_CHAR_ANGLE_BRAKET_OPENED );
			}
			else if ( c == ']' )
			{
				Private_AddCharType( result, i, 1, CharAnalyzer_CHAR_ANGLE_BRAKET_CLOSED );
			}
			else if ( c == '<' )
			{
				Private_AddCharType( result, i, 1, CharAnalyzer_CHAR_SQUARE_BRAKET_OPENED );
			}
			else if ( c == '>' )
			{
				Private_AddCharType( result, i, 1, CharAnalyzer_CHAR_SQUARE_BRAKET_CLOSED );
			}
			else if ( c == '_' )
			{
				Private_AddCharType( result, i, 1, CharAnalyzer_CHAR_UNDERSCORE );
			}
			else if ( c == '\\' )
			{
				Private_AddCharType( result, i, 1, CharAnalyzer_CHAR_BACKSLASH );
			}
			else
			{
				Private_AddCharTypeOrIncrExtent( result, i, 1, CharAnalyzer_CHAR_OTHERS );
			}
		}
	}
}

