
//          Copyright Vitali Baumtrok 2015.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


/**
 * @file shif.c
 * @author Vitali Baumtrok
 * @date 2015-06-27
 * @brief reference implementation for the SHiF parser
 * 
 * Reference implementation for the SHiF parser.
 */

#include <stdlib.h>
#include <stdio.h>

#define SHIF_EROR_WRONG_NODE_NAME 1
#define SHIF_EROR_WRONG_INDENT_DEPTH 2

#define SHIF_TRIM_BLANKS_LEFT 1
#define SHIF_TRIM_BLANKS_RIGHT 2
#define SHIF_TRIM_TABS_LEFT 4
#define SHIF_TRIM_TABS_RIGHT 8

static const char *const TAB = "\t";

/**
 * @brief Parses the indent token.
 * 
 * Parses the indent token. New line is considert to be CR or CRLF or LF.
 * 
 * @param inout_curr The current character to parse.
 * @param out_indentTokenBegin Pointer to where the beginning of the indent
 * token should be saved.
 * @param out_indentTokenLimit Pointer to where the end of the indent
 * token should be saved.
 */
static void
internal_SHiF_parseIndentToken
(
	const char **const inout_curr,
	const char **const out_indentTokenBegin,
	const char **const out_indentTokenLimit
)
{
	const char cFirst = **inout_curr;

	// unix new line
	if ( '\n' == cFirst )
	{
		*out_indentTokenBegin = TAB;
		*out_indentTokenLimit = TAB + 1;
		(*inout_curr) += 1;
	}
	// windows new line
	else if ( '\r' == cFirst )
	{
		*out_indentTokenBegin = TAB;
		*out_indentTokenLimit = TAB + 1;

		if ( '\n' == *((*inout_curr)+1) )
			(*inout_curr) += 2;
		else
			(*inout_curr) += 1;
	}
	// parse indent token
	else
	{
		const char *curr = *inout_curr;

		*out_indentTokenBegin = curr;
		curr += 1;

		while ( 1 )
		{
			const char cCurr = *curr;

			// unix new line
			if ( '\n' == cCurr )
			{
				*out_indentTokenLimit = curr;
				(*inout_curr) = curr + 1;
				break;
			}
			// windows new line
			else if ( '\r' == cCurr )
			{
				*out_indentTokenLimit = curr;

				if ( '\n' == *(curr+1) )
					(*inout_curr) = curr + 2;
				else
					(*inout_curr) = curr + 1;
				break;
			}
			// string end
			else if ( 0 == cCurr )
			{
				*out_indentTokenLimit = curr;
				(*inout_curr) = curr;
				break;
			}
			// continue to parse indent token
			else
			{
				curr += 1;
			}
		}
	}
}

/**
 * @brief Computes the indent depth.
 * 
 * Computes the indent depth and sets the curr text pointer at the end of
 * indent tokens.
 * 
 * @param inout_curr A c string to parse. Must not be 0 or empty.
 * @param in_indentTokenBegin Pointer to a character where the indent token
 * begins.
 * @param in_indentTokenLimit Pointer to a character where the indent token
 * ends.
 */
static int
internal_SHiF_getIndentDepth
(
	const char **const inout_curr,
	const char *const in_indentTokenBegin,
	const char *const in_indentTokenLimit,
	const size_t indentDepthSize
)
{
	const char *curr = *inout_curr;
	int indentDepth = 0;

	while ( 1 )
	{
		const char *indentTokenCurr = in_indentTokenBegin;
		
		while ( 1 )
		{
			const char cCurr = *curr;
			const char cIndentTokenCurr = *indentTokenCurr;

			// indent tokens continue
			if ( cCurr == cIndentTokenCurr )
			{
				curr += 1;
				indentTokenCurr += 1;

				// next indent token
				if ( indentTokenCurr == in_indentTokenLimit )
				{
					indentDepth += 1;
					break;
				}
			}
			// indent tokens end
			else
			{
				*inout_curr += ((size_t)indentDepth * indentDepthSize);

				return indentDepth;
			}
		}
	}
}

/**
 * @brief Skip blanks and tabs.
 * 
 * Skip blanks and tabs.
 * 
 * @param inout_cstring A c string to parse. Must not be 0.
 */
static void
internal_SHiF_skipBlanksTabsToRight
(
	const char **const inout_cstring
)
{
	const char *curr = *inout_cstring;

	while ( 1 )
	{
		const char cCurr = *curr;

		if ( (0 != cCurr) && (' ' == cCurr || '\t' == cCurr) )
			curr += 1;
		else
			break;
	}
	*inout_cstring = curr;
}

/**
 * @brief Skip blanks.
 * 
 * Skip blanks.
 * 
 * @param inout_cstring A c string to parse. Must not be 0.
 */
static void
internal_SHiF_skipBlanksToRight
(
	const char **const inout_cstring
)
{
	const char *curr = *inout_cstring;

	while (1)
	{
		const char cCurr = *curr;

		if ( 0 != cCurr && ' ' == cCurr )
			curr += 1;
		else
			break;
	}
	*inout_cstring = curr;
}

/**
 * @brief Skip tabs.
 * 
 * Skip tabs.
 * 
 * @param inout_cstring A c string to parse. Must not be 0.
 */
static void
internal_SHiF_skipTabsToRight
(
	const char **const inout_cstring
)
{
	const char *curr = *inout_cstring;

	while (1)
	{
		const char cCurr = *curr;

		if ( 0 != cCurr && '\t' == cCurr )
			curr += 1;
		else
			break;
	}
	*inout_cstring = curr;
}

/**
 * @brief Skip blanks and/or tabs.
 * 
 * Skip blanks and/or tabs.
 * 
 * @param inout_cstring A c string to parse. Must not be 0.
 * @param in_trimBlanks Boolean value if blanks are skipped.
 * @param in_trimTabs Boolean value if tabs are skipped.
 */
static void
internal_SHiF_trimLeft
(
	const char **const inout_cstring,
	const int in_trimBlanks,
	const int in_trimTabs
)
{
	if ( in_trimBlanks )
		if ( in_trimTabs )
			internal_SHiF_skipBlanksTabsToRight ( inout_cstring );
		else
			internal_SHiF_skipBlanksToRight ( inout_cstring );
	else
		if ( in_trimTabs )
			internal_SHiF_skipTabsToRight ( inout_cstring );
}

/**
 * @brief Adjust a pointer from right to left skipping blans and tabs.
 * 
 * Adjust a pointer from right to left skipping blans and tabs.
 * 
 * @param inout_string A c string to parse. Must not be 0.
 * @param in_leftLimit The limit to which to adjust. Must not be equal to
 * inout_string
 */
static void
internal_SHiF_skipBlanksTabsToLeft
(
	const char **const inout_string,
	const char *const in_leftLimit
)
{
	const char *curr = (*inout_string) - 1;

	while ( 1 )
	{
		const char cCurr = *curr;

		if ( ' '==cCurr || '\t'==cCurr )
		{
			if ( in_leftLimit != curr )
			{
				curr -= 1;
			}
			else
			{
				*inout_string = curr;
				break;
			}
		}
		else
		{
			*inout_string = curr + 1;
			break;
		}
	}
}

/**
 * @brief Adjust a pointer from right to left skipping blans.
 * 
 * Adjust a pointer from right to left skipping blans.
 * 
 * @param inout_string A c string to parse. Must not be 0.
 * @param in_leftLimit The limit to which to adjust.
 */
static void
internal_SHiF_skipBlanksToLeft (
	const char **const inout_string,
	const char *const in_leftLimit
)
{
	const char *curr = (*inout_string) - 1;

	while ( 1 )
	{
		const char cCurr = *curr;

		if ( ' '==cCurr )
		{
			if ( in_leftLimit != curr )
			{
				curr -= 1;
			}
			else
			{
				*inout_string = curr;
				break;
			}
		}
		else
		{
			*inout_string = curr + 1;
			break;
		}
	}
}

/**
 * @brief Adjust a pointer from right to left skipping tabs.
 * 
 * Adjust a pointer from right to left skipping tabs.
 * 
 * @param inout_string A c string to parse. Must not be 0.
 * @param in_leftLimit The limit to which to adjust.
 */
static void
internal_SHiF_skipTabsToLeft
(
	const char **const inout_string,
	const char *const in_leftLimit
)
{
	const char *curr = (*inout_string) - 1;

	while ( 1 ) {
		const char cCurr = *curr;

		if ( '\t'==cCurr )
		{
			if ( in_leftLimit != curr )
			{
				curr -= 1;
			}
			else
			{
				*inout_string = curr;
				break;
			}
		}
		else
		{
			*inout_string = curr + 1;
			break;
		}
	}
}

/**
 * @brief Adjust a pointer from right to left skipping blans and/or tabs.
 * 
 * Adjust a pointer from right to left skipping blans and/or tabs.
 * 
 * @param inout_string A c string to parse. Must not be 0.
 * @param in_leftLimit The limit to which to adjust.
 * @param in_trimBlanks Boolean value if blanks are skipped.
 * @param in_trimTabs Boolean value if tabs are skipped.
 */
static void
internal_SHiF_trimRight
(
	const char **const inout_string,
	const char *const in_leftLimit,
	const int in_trimBlanks,
	const int in_trimTabs
)
{
	if ( in_trimBlanks )
		if ( in_trimTabs )
			internal_SHiF_skipBlanksTabsToLeft ( inout_string, in_leftLimit );
		else
			internal_SHiF_skipBlanksToLeft ( inout_string, in_leftLimit );
	else
		if ( in_trimTabs )
			internal_SHiF_skipTabsToLeft ( inout_string, in_leftLimit );
}

/**
 * @brief Internal function to parse a SHiF c string.
 * 
 * Internal function to parse a SHiF c string. The given functions return
 * a value unequal to 0 to return from this function. All parameters except
 * in_cstring may be 0.
 * 
 * @param in_cstring A c string to parse. Must not be 0 or empty.
 * @param inout_data User data.
 * @param in_newNodeFct Pointer to a function. The function is called
 * when a node occurs.
 * @param in_errorFct Pointer to a function. The function is called
 * when an error accurs.
 * @param in_trimOption Bits to indicate whether blanks and/or tabs around
 * a node value should be skipped.
 * @return The value is equal to the value returned by one of the
 * functions. Normally it's 0.
 */
static int
internal_SHiF_parseNotEmptyString
(
	char *const in_cstring,
	void *const inout_data,
	int ( *const in_newNodeFct )
	(
		void *data,
		int indentDepth,
		char *nodeNameBegin,
		char *nodeNameEnd,
		char *nodeValueBegin,
		char *nodeValueEnd
	),
	int ( *const in_errorFct )
	(
		void *data,
		int errorCode,
		int errorLineNum
	),
	const int in_trimOption
)
{
	const int trimBlanksLeft = in_trimOption & SHIF_TRIM_BLANKS_LEFT;
	const int trimBlanksRight = in_trimOption & SHIF_TRIM_BLANKS_RIGHT;
	const int trimTabsLeft = in_trimOption & SHIF_TRIM_TABS_LEFT;
	const int trimTabsRight = in_trimOption & SHIF_TRIM_TABS_RIGHT;
	char *curr = in_cstring;
	const char *indentTokenBegin;
	const char *indentTokenLimit;
	int lineNum = 2;
	int prevIndentDepth = -1;

	// parses the first line (that's why lineNum=2)
	internal_SHiF_parseIndentToken ( (const char**) &curr, &indentTokenBegin, &indentTokenLimit );

	const size_t indentTokenSize = (size_t) (indentTokenLimit - indentTokenBegin);

	// iterate over each line
	while ( 1 )
	{
		char cCurr = *curr;

		// string end
		if ( 0 == cCurr )
		{
			return 0;
		}
		// valid line
		else
		{
			const int currIndentDepth = internal_SHiF_getIndentDepth ( (const char**) &curr, indentTokenBegin, indentTokenLimit, indentTokenSize );
			cCurr = *curr;

			// unix new line
			if ( '\n' == cCurr )
			{
				curr += 1;
				lineNum += 1;
			}
			// windows new line
			else if ( '\r' == cCurr )
			{
				if ( '\n' == *(curr+1) )
					curr += 2;
				else
					curr += 1;
				lineNum += 1;
			}
			// error: wrong node name
			else if ( '\t'==cCurr || ' '==cCurr )
			{
				if ( in_errorFct )
				{
					const int errorCode = in_errorFct ( inout_data, SHIF_EROR_WRONG_NODE_NAME, lineNum );

					// end parsing
					if ( errorCode ) {
						return errorCode;
					}
					// continue parsing
					else
					{
						// skip line
						while ( 1 )
						{
							curr += 1;
							cCurr = *curr;

							// unix new line
							if ( '\n' == cCurr )
							{
								curr += 1;
								lineNum += 1;
								break;
							}
							// windows new line
							else if ( '\r' == cCurr )
							{
								if ( '\n' == *(curr+1) )
									curr += 2;
								else
									curr += 1;
								lineNum += 1;
								break;
							}
							// string end
							else if ( 0 == cCurr )
							{
								return 0;
							}
						}
					}

				}
				else
				{
					return SHIF_EROR_WRONG_NODE_NAME;
				}
			}
			// error: wrong indent depth
			else if ( prevIndentDepth+1 < currIndentDepth )
			{
				if ( in_errorFct )
				{
					const int errorCode = in_errorFct ( inout_data, SHIF_EROR_WRONG_INDENT_DEPTH, lineNum );

					// end parsing
					if ( errorCode )
					{
						return errorCode;
					}
					// continue parsing
					else
					{
						// skip line
						while ( 1 )
						{
							curr += 1;
							cCurr = *curr;

							// unix new line
							if ( '\n' == cCurr )
							{
								curr += 1;
								lineNum += 1;
								break;
							}
							// windows new line
							else if ( '\r' == cCurr )
							{
								if ( '\n' == *(curr+1) )
									curr += 2;
								else
									curr += 1;
								lineNum += 1;
								break;
							}
							// string end
							else if ( 0 == cCurr )
							{
								return 0;
							}
						}
					}
				}
				else
				{
					return SHIF_EROR_WRONG_INDENT_DEPTH;
				}
			}
			// valid node
			else
			{
				//
				// parse node name
				//

				char *const nodeNameBegin = curr;
				curr += 1;

				while ( 1 )
				{
					cCurr = *curr;

					// unix new line
					if ( '\n' == cCurr )
					{
						if ( in_newNodeFct )
						{
							const int errorCode = in_newNodeFct ( inout_data, currIndentDepth, nodeNameBegin, curr, 0, 0 );
							if ( errorCode )
								return errorCode;
						}
						curr += 1;
						lineNum += 1;
						prevIndentDepth = currIndentDepth;
						break;
					}
					// windows new line
					else if ( '\r' == cCurr )
					{
						if ( in_newNodeFct )
						{
							const int errorCode = in_newNodeFct ( inout_data, currIndentDepth, nodeNameBegin, curr, 0, 0 );
							if (errorCode)
								return errorCode;
						}
						if ( '\n' == *(curr+1) )
							curr += 2;
						else
							curr += 1;
						lineNum += 1;
						prevIndentDepth = currIndentDepth;
						break;
					}
					// string end
					else if ( 0 == cCurr )
					{
						if ( in_newNodeFct )
						{
							const int errorCode = in_newNodeFct ( inout_data, currIndentDepth, nodeNameBegin, curr, 0, 0 );

							if ( errorCode )
								return errorCode;
						}
						return 0;
					}
					// node value begin
					else if ( ' ' == cCurr )
					{
						char *const nodeNameEnd = curr;

						//
						// parse node value
						//

						curr += 1;
						internal_SHiF_trimLeft ( (const char **) &curr, trimBlanksLeft, trimTabsLeft );
						cCurr = *curr;

						// unix new line (no node value)
						if ( '\n' == cCurr )
						{
							if ( in_newNodeFct )
							{
								const int errorCode = in_newNodeFct ( inout_data, currIndentDepth, nodeNameBegin, nodeNameEnd, 0, 0 );

								if ( errorCode )
									return errorCode;
							}
							curr += 1;
							lineNum += 1;
							prevIndentDepth = currIndentDepth;
							break;
						}
						// windows new line (no node value)
						else if ( '\r' == cCurr )
						{
							if ( in_newNodeFct )
							{
								const int errorCode = in_newNodeFct ( inout_data, currIndentDepth, nodeNameBegin, nodeNameEnd, 0, 0 );

								if ( errorCode )
									return errorCode;
							}
							if ( '\n' == *(curr+1) )
								curr += 2;
							else
								curr += 1;
							lineNum += 1;
							prevIndentDepth = currIndentDepth;
							break;
						}
						// string end (no node value)
						else if ( 0 == cCurr )
						{
							if ( in_newNodeFct )
								return in_newNodeFct ( inout_data, currIndentDepth, nodeNameBegin, nodeNameEnd, 0, 0 );
							else
								return 0;
						}
						// parse node value
						else
						{
							char *const nodeValueBegin = curr;
							curr += 1;

							while ( 1 )
							{
								cCurr = *curr;

								// unix new line
								if ( '\n' == cCurr )
								{
									char *nodeValueEnd = curr;
									internal_SHiF_trimRight ( (const char **) &nodeValueEnd, (const char *const) nodeValueBegin, trimBlanksRight, trimTabsRight );

									if ( in_newNodeFct )
									{
										int errorCode;

										if ( nodeValueBegin == nodeValueEnd )
											errorCode = in_newNodeFct ( inout_data, currIndentDepth, nodeNameBegin, nodeNameEnd, 0, 0 );
										else
											errorCode = in_newNodeFct ( inout_data, currIndentDepth, nodeNameBegin, nodeNameEnd, nodeValueBegin, nodeValueEnd );

										if ( errorCode )
											return errorCode;
									}
									curr += 1;
									lineNum += 1;
									prevIndentDepth = currIndentDepth;
									break;
								}
								// windows new line
								else if ( '\r' == cCurr )
								{
									char *nodeValueEnd = curr;
									internal_SHiF_trimRight ( (const char **) &nodeValueEnd, (const char *const) nodeValueBegin, trimBlanksRight, trimTabsRight );

									if ( in_newNodeFct )
									{
										int errorCode;

										if ( nodeValueBegin == nodeValueEnd )
											errorCode = in_newNodeFct ( inout_data, currIndentDepth, nodeNameBegin, nodeNameEnd, 0, 0 );
										else
											errorCode = in_newNodeFct ( inout_data, currIndentDepth, nodeNameBegin, nodeNameEnd, nodeValueBegin, nodeValueEnd );

										if ( errorCode )
											return errorCode;
									}

									if ( '\n' == *(curr+1) )
										curr += 2;
									else
										curr += 1;
									lineNum += 1;
									prevIndentDepth = currIndentDepth;
									break;
								}
								// string end
								else if ( 0 == cCurr )
								{
									char *nodeValueEnd = curr;
									internal_SHiF_trimRight ( (const char **) &nodeValueEnd, (const char *const) nodeValueBegin, trimBlanksRight, trimTabsRight );

									if ( in_newNodeFct )
										if ( nodeValueBegin == nodeValueEnd )
											return in_newNodeFct ( inout_data, currIndentDepth, nodeNameBegin, nodeNameEnd, 0, 0 );
										else
											return in_newNodeFct ( inout_data, currIndentDepth, nodeNameBegin, nodeNameEnd, nodeValueBegin, nodeValueEnd );
									else
										return 0;
								}
								// continue with next node value character
								else
								{
									curr += 1;
								}
							}
							break;
						}
					}
					// continue with next node name character
					else
					{
						curr += 1;
					}
				}
			}
		}
	}
}

/**
 * @brief Parse a SHiF c string.
 * 
 * Parse a SHiF c string. The given functions return a value unequal to 0
 * to return from this function. All parameters may be 0.
 * 
 * @param in_cstring A c string to parse.
 * @param inout_data User data.
 * @param in_newNodeFct Pointer to a function. The function is called
 * when a node occurs.
 * @param in_errorFct Pointer to a function. The function is called
 * when an error accurs.
 * @param in_trimOption Bits to indicate whether blanks and/or tabs around
 * a node value should be skipped.
 * @return The value is equal to the value returned by one of the
 * functions. Normally it's 0.
 */
int
SHiF_parse
(
	char *const in_cstring,
	void *const inout_data,
	int ( *const in_newNodeFct )
	(
		void *data,
		int indentDepth,
		char *nodeNameBegin,
		char *nodeNameEnd,
		char *nodeValueBegin,
		char *nodeValueEnd
	),
	int ( *const in_errorFct )
	(
		void *data,
		int errorCode,
		int errorLineNum
	),
	const int in_trimOption
)
{
	if ( in_cstring && (0 != *in_cstring) )
		return
		internal_SHiF_parseNotEmptyString
		(
			in_cstring,
			inout_data,
			in_newNodeFct,
			in_errorFct,
			in_trimOption
		);

	else
		return 0;
}
