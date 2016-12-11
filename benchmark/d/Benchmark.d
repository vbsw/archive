/*
* This is free and unencumbered software released into the public domain.
*
* Anyone is free to copy, modify, publish, use, compile, sell, or
* distribute this software, either in source code form or as a compiled
* binary, for any purpose, commercial or non-commercial, and by any
* means.
*
* In jurisdictions that recognize copyright laws, the author or authors
* of this software dedicate any and all copyright interest in the
* software to the public domain. We make this dedication for the benefit
* of the public at large and to the detriment of our heirs and
* successors. We intend this dedication to be an overt act of
* relinquishment in perpetuity of all present and future rights to this
* software under copyright law.
*
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
* EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
* MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
* IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
* OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
* ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
* OTHER DEALINGS IN THE SOFTWARE.
*
* For more information, please refer to <http://unlicense.org>
*/


import std.stdio, std.file, core.stdc.stdlib, core.stdc.string;


__gshared immutable string   INSERT_REMOVE_TEST_NAME     = "insert/remove";
__gshared immutable string   INSERT_REMOVE_MOD_TEST_NAME = "insert/remove (mod)";
__gshared immutable string   BINARY_SEARCH_TEST_NAME     = "binary search";
__gshared immutable string[] FILE_NAMES                  = [
	"benchmark.parameters.txt",
	"c/benchmark.parameters.txt",
	"c/Release/benchmark.parameters.txt",
	"benchmark/benchmark.parameters.txt",
	"benchmark/c/benchmark.parameters.txt",
	"benchmark/c/Release/benchmark.parameters.txt",
	"../benchmark.parameters.txt",
	"../../benchmark.parameters.txt",
	"../../../benchmark.parameters.txt"
];
__gshared immutable string   NUMBER_OF_ITERATIONS_STR    = "number of iterations";
__gshared immutable string   SIZE_OF_ARRAY_STR           = "length of array";
__gshared immutable string   FILE_NAME                   = FILE_NAMES[0];
__gshared immutable int      DEFAULT_ITERATION_NUMBER    = 5;
__gshared immutable int      DEFAULT_ARRAY_LENGTH        = 32768 - 2;


void printBenchmarkHeader () {
	writeln ( "\n" );
	writeln ( "/////////////////////////////////////////////////////////////////" );
	writeln ( "//                                                             //" );
	writeln ( "//                       D benchmark                           //" );
	writeln ( "//                                                             //" );
	writeln ( "/////////////////////////////////////////////////////////////////" );
}

int searchValidFile () {
	foreach ( int i, string s; FILE_NAMES ) {
		if ( std.file.exists ( s ) )
			return i;
	}
	return -1;
}

ubyte[] readFile () {
	const int fileIndex = searchValidFile ();

	if ( fileIndex < 0 )
		return null;
	else
		return cast(ubyte[]) std.file.read ( FILE_NAMES[fileIndex] );
}

int parseNumber ( ubyte[] bytes ) {
	if ( bytes.length > 0 ) {
		ubyte b = bytes[0];
		int result = -1;

		if ( (b > 0x2F) && (b < 0x3A) ) {
			result = b - 0x30;

			for ( int i = 1; i < bytes.length; i++ ) {
				b = bytes[i];
				
				if ( (b > 0x2F) && (b < 0x3A) )
					result = result * 10 + b - 0x30;
				else
					break;
			}
		}
		return result;
	}
	else {
		return -1;
	}
}

int parseIterNumber ( ubyte[] bytes ) {
	return parseNumber ( bytes );
}

int parseArraySize ( ubyte[] bytes ) {
	foreach ( int i, ubyte b; bytes )
		if ( b == 0x0A )
			return parseNumber ( bytes[i+1..bytes.length] );
	return -1;
}

void parseParameters ( ubyte[] parameters, ref int parsedIterationNumber, ref int parsedArraySize ) {
	parsedIterationNumber = parseIterNumber ( parameters );
	parsedArraySize = parseArraySize ( parameters );
}

void setDefaultSettings ( ref int iterNumber, ref int arrayLength ) {
	iterNumber = DEFAULT_ITERATION_NUMBER;
	arrayLength = DEFAULT_ARRAY_LENGTH;
}

void assignSettings ( ubyte[] parameters, ref int parsedIterationNumber, ref int parsedArraySize, ref int iterNumber, ref int arrayLength ) {
	if ( parameters != null )
		parseParameters ( parameters, parsedIterationNumber, parsedArraySize );
	else
		setDefaultSettings ( iterNumber, arrayLength );
}

void correctSettings ( const int parsedIterationNumber, const int parsedArraySize, ref int iterNumber, ref int arrayLength ) {
	if ( parsedIterationNumber > 0 )
		iterNumber = parsedIterationNumber;
	else
		iterNumber = DEFAULT_ITERATION_NUMBER;

	if ( parsedArraySize > 0 )
		arrayLength = parsedArraySize;
	else
		arrayLength = DEFAULT_ARRAY_LENGTH;
}

void printNoticeHeader () {
	writeln ( "//                                                             //" );
	writeln ( "//                          Notice                             //" );
	writeln ( "//                                                             //" );
}

void printNoticeSomeUnparsedValues () {
	writeln ( "// Some parameters could not be parsed and are set to default. //" );
}

void printNoticeSomeTooLowValues () {
	writeln ( "// Some parameters are too low and are set to default.         //" );
}

void printNoticeFooter () {
	writeln ( "//                                                             //" );
	writeln ( "/////////////////////////////////////////////////////////////////" );
}

void printNoticeCouldntReadFile () {
	writeln ( "// Could not read file '", FILE_NAME, "'.             //" );
}

void printNotice ( const bool fileRead, const int parsedIterationNumber, const int parsedArraySize ) {
	if ( fileRead ) {
		if ( parsedIterationNumber <= 0 || parsedArraySize <= 0 ) {
			printNoticeHeader ();

			if ( parsedIterationNumber < 0 || parsedArraySize < 0 ) {
				printNoticeSomeUnparsedValues ();
			}

			if ( parsedIterationNumber == 0 || parsedArraySize == 0 ) {
				printNoticeSomeTooLowValues ();
			}
			printNoticeFooter ();
		}
	}
	else {
		printNoticeHeader ();
		printNoticeCouldntReadFile ();
		printNoticeFooter ();
	}
}

private static void printParametersHeader () {
	writeln ( "//                                                             //" );
	writeln ( "//                        Parameters                           //" );
	writeln ( "//                                                             //" );
}

private static void printParameter ( string paramName, const bool validParam, const int paramValue, const int defaultParamValue ) {
	if ( validParam )
		writefln ( "// %25s:%12d                      //", paramName, paramValue );
	else
		writefln ( "// %25s:%12d (default value)      //", paramName, defaultParamValue );
}

private static void printParametersFooter () {
	writeln ( "//                                                             //" );
	writeln ( "/////////////////////////////////////////////////////////////////" );
}

private static void printParameters ( const bool fileRead, const int parsedIterationNumber, const int parsedArraySize ) {
	printParametersHeader ();
	printParameter ( NUMBER_OF_ITERATIONS_STR, fileRead && (parsedIterationNumber>0), parsedIterationNumber, DEFAULT_ITERATION_NUMBER );
	printParameter ( SIZE_OF_ARRAY_STR, fileRead && (parsedArraySize>0), parsedArraySize, DEFAULT_ARRAY_LENGTH );
	printParametersFooter ();
}

void readParameters ( ref int iterNumber, ref int arrayLength ) {
	ubyte[] fileBytes = readFile ();
	int parsedIterationNumber = 0;
	int parsedArraySize = 0;

	assignSettings ( fileBytes, parsedIterationNumber, parsedArraySize, iterNumber, arrayLength );
	correctSettings ( parsedIterationNumber, parsedArraySize, iterNumber, arrayLength );
	printNotice ( fileBytes!=null, parsedIterationNumber, parsedArraySize );
	printParameters ( fileBytes!=null, parsedIterationNumber, parsedArraySize );
}

void printResultsHeader () {
	writeln ( "//                                                             //" );
	writeln ( "//                         Results                             //" );
	writeln ( "//                                                             //" );
}

double getTimeInSeconds () {
	return cast(double) std.datetime.Clock.currStdTime () / 10000000.0;
}

void printTestTime ( string testName, const double passedSeconds ) {
	writefln ( "// %25s:%12.3f seconds              //", testName, passedSeconds );
}

void printOutOfMemory ( string testName ) {
	writefln ( "// %25s:       error (out of memory)      //", testName);
}

void benchmarkInsertRemove ( const int iterNumber, const int arrayLength ) {
	int[] arr = new int[arrayLength];

	arr.length = 0;
	arr.assumeSafeAppend ();

	const double startTime = getTimeInSeconds ();
	const int arrLimit = arrayLength - 1;

	for ( int iterCurr = 0; iterCurr < iterNumber; iterCurr += 1 ) {

		// inserting
		for ( int currLength = 0, flag = 0; currLength < arrLimit; currLength += 1 ) {

			if ( flag != 0 ) {
				std.array.insertInPlace ( arr, 1, currLength );
				flag = 0;
			} else {
				std.array.insertInPlace ( arr, 0, currLength );
				flag = 1;
			}
		}

		// removing
		for ( int currLength = arrLimit, flag = arrayLength % 2; currLength > 0; currLength -= 1 ) {

			if ( flag != 0 ) {
				arr = std.algorithm.mutation.remove ( arr, 1 );
				flag = 0;
			} else {
				arr = std.algorithm.mutation.remove ( arr, 0 );
				flag = 1;
			}
		}
	}
	printTestTime ( INSERT_REMOVE_TEST_NAME, ( getTimeInSeconds () - startTime ) / cast(double) iterNumber );
}

void benchmarkInsertRemoveMod ( const int iterNumber, const int arrayLength ) {
	int *arr = cast(int*) core.stdc.stdlib.malloc ( (cast(size_t) arrayLength) * int.sizeof );

	if (arr) {
		const double startTime = getTimeInSeconds ();
		const int arrLimit = arrayLength - 1;
		int iterCurr;
		int* arrNext = arr + 1;
		int* arrNextNext = arr + 2;

		for ( iterCurr = 0; iterCurr < iterNumber; iterCurr += 1 ) {
			int currLength, flag;

			// inserting
			for (currLength = 0, flag = 0; currLength < arrLimit; currLength += 1) {

				if (flag) {
					core.stdc.string.memmove ( arrNextNext, arrNext, cast(size_t) (currLength - flag) );
					(*arrNext) = currLength;
					flag = 0;
				}
				else {
					core.stdc.string.memmove ( arrNext, arr, cast(size_t) (currLength - flag) );
					(*arr) = currLength;
					flag = 1;
				}
			}

			// removing
			for ( currLength = arrLimit, flag = arrayLength % 2; currLength > 0; currLength -= 1 ) {

				if (flag) {
					core.stdc.string.memmove ( arrNext, arrNextNext, cast(size_t) (currLength - flag) );
					(*arrNext) = currLength;
					flag = 0;
				}
				else {
					core.stdc.string.memmove ( arr, arrNext, cast(size_t) (currLength - flag) );
					(*arr) = currLength;
					flag = 1;
				}
			}
		}
		printTestTime(INSERT_REMOVE_MOD_TEST_NAME, (getTimeInSeconds() - startTime) / cast(double) iterNumber);
		core.stdc.stdlib.free(arr);
	}
	else {
		printOutOfMemory(INSERT_REMOVE_MOD_TEST_NAME);
	}
}

ubyte[] newWord ( const int length, const ubyte value ) {
	ubyte[] chars = new ubyte[length];
	foreach ( int i; chars )
		chars[i] = value;
	return chars;
}

void allocateWords ( ubyte[][] words ) {
	const int maxValue = 127;
	const int alphabetArrayLength = words.length + maxValue - words.length % maxValue;
	const int stepLength = alphabetArrayLength / maxValue;

	for ( int i = 0, j; i < words.length; i = j ) {
		const int jLimit = i + stepLength;
		const ubyte number = cast(ubyte) ( i / stepLength + 1 );

		for ( j = i; (j < jLimit) && (j < words.length); j += 1 ) {
			const int wordLength = (j % maxValue) + 1;

			words[j] = newWord ( wordLength, number );
		}
	}
}

void allocateWordsToSearch ( ubyte[][] words ) {
	for ( int i=0; i<words.length; )
		for ( int j=-2; (j < 128) && (i < words.length); j += 3, i += 1 )
			words[i] = newWord ( 10, cast(ubyte) j );
}

private static void benchmarkBinarySearch ( const int iterNumber, const int arrayLength ) {
	const int wordsToSearchLength = 32766;
	ubyte[][] words = new ubyte[][arrayLength];
	ubyte[][] wordsToSearch = new ubyte[][wordsToSearchLength];

	allocateWords ( words );
	allocateWordsToSearch ( wordsToSearch );

	std.range.SortedRange!(ubyte[][], "a < b") sortedWords = std.range.assumeSorted ( words );

	const double startTime = getTimeInSeconds ();
	const ubyte aValue = wordsToSearch[0][0];
	const ubyte bValue = cast(ubyte) ( aValue * 2 );

	for ( int i = 0; i < iterNumber; i++ ) {
		for ( int j = 0; j < wordsToSearchLength; j++ ) {
			if ( std.algorithm.searching.canFind ( sortedWords, wordsToSearch[j] ) )
				wordsToSearch[0][0] = aValue;
			else
				wordsToSearch[0][0] = bValue;
		}
	}
	printTestTime ( BINARY_SEARCH_TEST_NAME, ( getTimeInSeconds () - startTime ) / cast(double) iterNumber );
}

void benchmark ( const int iterNumber, const int arrayLength ) {
	benchmarkInsertRemove ( iterNumber, arrayLength );
	benchmarkInsertRemoveMod ( iterNumber, arrayLength );
	benchmarkBinarySearch ( iterNumber, arrayLength );
}

void printResultsFooter () {
	writeln ( "//                                                             //" );
	writeln ( "/////////////////////////////////////////////////////////////////" );
}

void main () {
	int iterNumber = 0;
	int arrayLength = 0;

	printBenchmarkHeader ();
	readParameters ( iterNumber, arrayLength );
	printResultsHeader ();
	benchmark ( iterNumber, arrayLength );
	printResultsFooter ();
	writeln ( "\n" );
}





