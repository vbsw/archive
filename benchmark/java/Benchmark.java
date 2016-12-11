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


import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;


public class Benchmark {

	private static final String INSERT_REMOVE_TEST_NAME = "insert/remove";
	private static final String INSERT_REMOVE_MOD_TEST_NAME = "insert/remove (mod)";
	private static final String BINARY_SEARCH_TEST_NAME = "binary search";

	private static final String[] FILE_NAMES = new String[] {
			"benchmark.parameters.txt", "c/benchmark.parameters.txt",
			"c/Release/benchmark.parameters.txt",
			"benchmark/benchmark.parameters.txt",
			"benchmark/c/benchmark.parameters.txt",
			"benchmark/c/Release/benchmark.parameters.txt",
			"../benchmark.parameters.txt",
			"../../benchmark.parameters.txt",
			"../../../benchmark.parameters.txt" };
	private static final String NUMBER_OF_ITERATIONS_STR = "number of iterations";
	private static final String SIZE_OF_ARRAY_STR = "length of array";

	private static final String FILE_NAME = FILE_NAMES[0];

	private static final int DEFAULT_ITERATION_NUMBER = 5;
	private static final int DEFAULT_ARRAY_LENGTH = 32768 - 2;
	
	private static class Int {
		public int value = 0;
	}

	private static class Word implements Comparable<Word> {

		public final byte[] arr;

		public Word ( final int wordLength, final byte initialValue ) {
			arr = new byte[wordLength];
			Arrays.fill ( arr, initialValue );
		}

		@Override
		public int compareTo ( final Word w ) {
			int i = 0;

			while ( true ) {
				if ( i == arr.length ) {
					if ( i == w.arr.length ) {
						return 0;
					} else {
						return -1;
					}
				} else if ( i == w.arr.length ) {
					return 1;
				} else {
					final byte c1 = arr[i];
					final byte c2 = w.arr[i];

					if ( c1 < c2 )
						return -1;
					else if ( c2 < c1 )
						return 1;
					else
						i++;
				}
			}
		}

	}

	private static void printBenchmarkHeader () {
		System.out.println ( "\n" );
		System.out.println ( "/////////////////////////////////////////////////////////////////" );
		System.out.println ( "//                                                             //" );
		System.out.println ( "//                    Java benchmark                           //" );
		System.out.println ( "//                                                             //" );
		System.out.println ( "/////////////////////////////////////////////////////////////////" );
	}

	private static File openFile () {
		File file = null;
		
		for ( String fileName : FILE_NAMES ) {
			file = new File ( fileName );
			try {
				if ( file.isFile () && file.canRead () )
					break;
				else
					file = null;
			}
			catch ( SecurityException e ) {
				file = null;
			}
		}
		return file;
	}

	private static String readFile () {
		final File file = openFile ();
		
		if ( file != null ) {
			final byte[] bytes = new byte[ (int) file.length () ];
			FileInputStream inStream;
			try {
				inStream = new FileInputStream ( file );
				inStream.read ( bytes );
				inStream.close ();
				
				return new String ( bytes, "UTF-8" );
			} catch ( IOException e ) {
				return null;
			}
		} else {
			return null;
		}
	}

	private static int parseNumber ( final String parameters ) {
		if ( parameters.length() > 0 ) {
			char c = parameters.charAt ( 0 );
			int result = -1;
			
			if ( (c > 0x2F) && (c < 0x3A) ) {
				result = c - 0x30;
				
				for ( int i = 1; i < parameters.length (); i++ ) {
					c = parameters.charAt ( i );
					
					if ( (c > 0x2F) && (c < 0x3A) )
						result = result * 10 + c - 0x30;
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

	private static int parseIterNumber ( final String parameters ) {
		return parseNumber ( parameters );
	}

	private static int parseArraySize ( String parameters ) {
		for ( int i = 0; i < parameters.length (); i++ ) {
			final char c = parameters.charAt ( 0 );
			
			if ( c == '\n' )
				return parseNumber ( parameters.substring ( i+1 ) );
		}
		return -1;
	}

	private static void parseParameters ( final String parameters, final Int parsedIterationNumber, final Int parsedArraySize ) {
		parsedIterationNumber.value = parseIterNumber ( parameters );
		parsedArraySize.value = parseArraySize ( parameters );
	}

	private static void setDefaultSettings ( final Int iterNumber, final Int arrayLength ) {
		iterNumber.value = DEFAULT_ITERATION_NUMBER;
		arrayLength.value = DEFAULT_ARRAY_LENGTH;
	}

	private static void assignSettings ( final String parameters, final Int parsedIterationNumber, final Int parsedArraySize, final Int iterNumber, final Int arrayLength ) {
		if ( parameters != null )
			parseParameters ( parameters, parsedIterationNumber, parsedArraySize );
		else
			setDefaultSettings ( iterNumber, arrayLength );
	}

	private static void correctSettings ( final int parsedIterationNumber, final int parsedArraySize, Int iterNumber, Int arrayLength ) {
		if ( parsedIterationNumber > 0 )
			iterNumber.value = parsedIterationNumber;
		else
			iterNumber.value = DEFAULT_ITERATION_NUMBER;

		if ( parsedArraySize > 0 )
			arrayLength.value = parsedArraySize;
		else
			arrayLength.value = DEFAULT_ARRAY_LENGTH;
	}

	private static void printNoticeHeader () {
		System.out.println ( "//                                                             //" );
		System.out.println ( "//                          Notice                             //" );
		System.out.println ( "//                                                             //" );
	}

	private static void printNoticeSomeUnparsedValues () {
		System.out.println ( "// Some parameters could not be parsed and are set to default. //" );
	}

	private static void printNoticeSomeTooLowValues () {
		System.out.println ( "// Some parameters are too low and are set to default.         //" );
	}

	private static void printNoticeFooter () {
		System.out.println ( "//                                                             //" );
		System.out.println ( "/////////////////////////////////////////////////////////////////" );
	}

	private static void printNoticeCouldntReadFile () {
		System.out.println ( "// Could not read file '" + FILE_NAME + "'.             //" );
	}

	private static void printNotice ( final boolean fileRead, final int parsedIterationNumber, final int parsedArraySize ) {
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
		System.out.println ( "//                                                             //" );
		System.out.println ( "//                        Parameters                           //" );
		System.out.println ( "//                                                             //" );
	}

	private static void printParameter ( final String paramName, final boolean validParam, final int paramValue, final int defaultParamValue ) {
		if ( validParam )
			System.out.format ( "// %25s:%12d                      //\n", paramName, paramValue );
		else
			System.out.format ( "// %25s:%12d (default value)      //\n", paramName, defaultParamValue );
	}

	private static void printParametersFooter () {
		System.out.println ( "//                                                             //" );
		System.out.println ( "/////////////////////////////////////////////////////////////////" );
	}

	private static void printParameters ( final boolean fileRead, final int parsedIterationNumber, final int parsedArraySize ) {
		printParametersHeader ();
		printParameter ( NUMBER_OF_ITERATIONS_STR, fileRead && parsedIterationNumber>0, parsedIterationNumber, DEFAULT_ITERATION_NUMBER );
		printParameter ( SIZE_OF_ARRAY_STR, fileRead && parsedArraySize>0, parsedArraySize, DEFAULT_ARRAY_LENGTH );
		printParametersFooter ();
	}

	private static void readParameters ( final Int iterNumber, final Int arrayLength ) {
		final String fileText = readFile ();
		final Int parsedIterationNumber = new Int ();
		final Int parsedArraySize = new Int ();

		assignSettings ( fileText, parsedIterationNumber, parsedArraySize, iterNumber, arrayLength );
		correctSettings ( parsedIterationNumber.value, parsedArraySize.value, iterNumber, arrayLength );
		printNotice ( fileText!=null, parsedIterationNumber.value, parsedArraySize.value );
		printParameters ( fileText!=null, parsedIterationNumber.value, parsedArraySize.value );
	}

	private static void printResultsHeader () {
		System.out.println ( "//                                                             //" );
		System.out.println ( "//                         Results                             //" );
		System.out.println ( "//                                                             //" );
	}

	private static double getTimeInSeconds () {
		return (double) System.nanoTime () / 1000000000.0d;
	}

	private static void printTestTime ( final String testName, final double passedSeconds ) {
		System.out.format ( "// %25s:%12.3f seconds              //\n", testName, passedSeconds );
	}

	private static void benchmarkInsertRemove ( final int iterNumber, final int arrayLength ) {
		final ArrayList<Integer> arr = new ArrayList<Integer> ( arrayLength );

		final double startTime = getTimeInSeconds ();
		final int arrLimit = arrayLength - 1;

		for ( int iterCurr = 0; iterCurr < iterNumber; iterCurr += 1 ) {

			// inserting
			for ( int currLength = 0, flag = 0; currLength < arrLimit; currLength += 1 ) {

				if ( flag != 0 ) {
					arr.add ( 1, currLength );
					flag = 0;
				} else {
					arr.add ( 0, currLength );
					flag = 1;
				}
			}

			// removing
			for ( int currLength = arrLimit, flag = arrayLength % 2; currLength > 0; currLength -= 1 ) {

				if ( flag != 0 ) {
					arr.remove ( 1 );
					flag = 0;
				} else {
					arr.remove ( 0 );
					flag = 1;
				}
			}
		}
		printTestTime ( INSERT_REMOVE_TEST_NAME, ( getTimeInSeconds () - startTime ) / iterNumber );
	}

	private static void benchmarkInsertRemoveMod ( final int iterNumber, final int arrayLength ) {
		final int[] arr = new int[arrayLength];

		final double startTime = getTimeInSeconds ();
		final int arrLimit = arrayLength - 1;

		for ( int iterCurr = 0; iterCurr < iterNumber; iterCurr += 1 ) {

			// inserting
			for ( int currLength = 0, flag = 0; currLength < arrLimit; currLength += 1 ) {

				if ( flag != 0 ) {
					System.arraycopy ( arr, 1, arr, 2, currLength-flag );
					arr[1] = currLength;
					flag = 0;
				} else {
					System.arraycopy ( arr, 0, arr, 1, currLength-flag );
					arr[0] = currLength;
					flag = 1;
				}
			}

			// removing
			for ( int currLength = arrLimit, flag = arrayLength % 2; currLength > 0; currLength -= 1 ) {

				if ( flag != 0 ) {
					System.arraycopy ( arr, 2, arr, 1, currLength-flag );
					arr[1] = currLength;
					flag = 0;
				} else {
					System.arraycopy ( arr, 1, arr, 0, currLength-flag );
					arr[0] = currLength;
					flag = 1;
				}
			}
		}
		printTestTime ( INSERT_REMOVE_MOD_TEST_NAME, ( getTimeInSeconds () - startTime ) / iterNumber );
	}

	private static void allocateStrings ( final Word[] arr ) {
		final int maxValue = 127;
		final int alphabetArrayLength = arr.length + maxValue - arr.length % maxValue;
		final int stepLength = alphabetArrayLength / maxValue;

		for ( int i = 0, j; i < arr.length; i = j ) {
			final int jLimit = i + stepLength;
			final byte number = (byte) ( i / stepLength + 1 );

			for ( j = i; j < jLimit && j < arr.length; j += 1 ) {
				final int wordLength = j % maxValue + 1;
				
				arr[j] = new Word ( wordLength, number );
			}
		}
	}

	private static void allocateWords ( final Word[] arr ) {
		for ( int i=0; i<arr.length; )
			for ( int j=-2; (j < 128) && (i < arr.length); j += 3, i += 1 )
				arr[i] = new Word ( 10, (byte) j );
	}

	private static void benchmarkBinarySearch ( final int iterNumber, final int arrayLength ) {
		final int wordsToSearchLength = 32766;
		final Word[] arr = new Word[arrayLength];
		final Word[] wordsToSearch = new Word[wordsToSearchLength];

		allocateStrings ( arr );
		allocateWords ( wordsToSearch );

		final double startTime = getTimeInSeconds ();
		final byte aValue = wordsToSearch[0].arr[0];
		final byte bValue = (byte) ( aValue * 2 );

		for ( int i = 0; i < iterNumber; i++ ) {
			for ( int j = 0; j < wordsToSearchLength; j++ ) {
				if ( Arrays.binarySearch ( arr, wordsToSearch[j] ) <= 0 )
					wordsToSearch[0].arr[0] = aValue;
				else
					wordsToSearch[0].arr[0] = bValue;
			}
		}
		printTestTime ( BINARY_SEARCH_TEST_NAME, ( getTimeInSeconds () - startTime ) / iterNumber );
	}

	private static void benchmark ( final int iterNumber, final int arrayLength ) {
		benchmarkInsertRemove ( iterNumber, arrayLength );
		benchmarkInsertRemoveMod ( iterNumber, arrayLength );
		benchmarkBinarySearch ( iterNumber, arrayLength );
	}

	private static void printResultsFooter () {
		System.out.println ( "//                                                             //" );
		System.out.println ( "/////////////////////////////////////////////////////////////////" );
	}

	public static void main ( final String[] args ) {
		final Int iterNumber = new Int ();
		final Int arrayLength = new Int ();

		printBenchmarkHeader ();
		readParameters ( iterNumber, arrayLength );
		printResultsHeader ();
		benchmark ( iterNumber.value, arrayLength.value );
		printResultsFooter ();
		System.out.println ( "\n" );
	}

}
