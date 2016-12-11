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


#include <stdlib.h>
#include <stdio.h>
#include <string.h>


#define FILE_NAMES_NUMBER 9
#define FILE_NAME (*FILE_NAMES)

#define DEFAULT_ITERATION_NUMBER 5
#define DEFAULT_ARRAY_LENGTH (32768-2)

const char *const FILE_NAMES[FILE_NAMES_NUMBER] = {
	"benchmark.parameters.txt",
	"c/benchmark.parameters.txt",
	"c/Release/benchmark.parameters.txt",
	"benchmark/benchmark.parameters.txt",
	"benchmark/c/benchmark.parameters.txt",
	"benchmark/c/Release/benchmark.parameters.txt",
	"../benchmark.parameters.txt",
	"../../benchmark.parameters.txt",
	"../../../benchmark.parameters.txt"
};
const char *const NUMBER_OF_ITERATIONS_STR = "number of iterations";
const char *const SIZE_OF_ARRAY_STR = "length of array";


/* private functions */

static FILE *openFile() {
	FILE *file;
	int i = 0;

	do {
		file = fopen(FILE_NAMES[i], "rb");
		i += 1;
	} while ((0 == file) && (i < FILE_NAMES_NUMBER));

	return file;
}

static char *readFile() {
	FILE *file = openFile();
	char *str = 0;

	if (file) {
		fseek(file, 0, SEEK_END);

		const size_t size = (size_t) ftell(file);
		str = (char*)malloc(size + 1);
		str[size] = 0;

		fseek(file, 0, SEEK_SET);
		fread(str, size, 1, file);
		fclose(file);
	}

	return str;
}

static void printNoticeHeader() {
	printf("//                                                             //\n");
	printf("//                          Notice                             //\n");
	printf("//                                                             //\n");
}

static void printNoticeSomeUnparsedValues() {
	printf("// Some parameters could not be parsed and are set to default. //\n");
}

static void printNoticeSomeTooLowValues() {
	printf("// Some parameters are too low and are set to default.         //\n");
}

static void printNoticeFooter() {
	printf("//                                                             //\n");
	printf("/////////////////////////////////////////////////////////////////\n");
}

static void printNoticeCouldntReadFile() {
	printf("// Could not read file '%s'.             //\n", FILE_NAME);
}

static void printParametersHeader() {
	printf("//                                                             //\n");
	printf("//                        Parameters                           //\n");
	printf("//                                                             //\n");
}

static void printParametersFooter() {
	printf("//                                                             //\n");
	printf("/////////////////////////////////////////////////////////////////\n");
}

static void printNotice(const int fileRead, const int parsedIterationNumber, const int parsedArraySize) {
	if (fileRead) {
		if (parsedIterationNumber <= 0 || parsedArraySize <= 0) {
			printNoticeHeader();

			if (parsedIterationNumber < 0 || parsedArraySize < 0) {
				printNoticeSomeUnparsedValues();
			}

			if (parsedIterationNumber == 0 || parsedArraySize == 0) {
				printNoticeSomeTooLowValues();
			}
			printNoticeFooter();
		}
	}
	else {
		printNoticeHeader();
		printNoticeCouldntReadFile();
		printNoticeFooter();
	}
}

static void printParameter(const char *paramName, const int validParam, const int paramValue, const int defaultParamValue) {
	if (validParam)
		printf("// %25s:%12d                      //\n", paramName, paramValue);
	else
		printf("// %25s:%12d (default value)      //\n", paramName, defaultParamValue);
}

static void printParameters(const int fileRead, const int parsedIterationNumber, const int parsedArraySize) {
	printParametersHeader();
	printParameter(NUMBER_OF_ITERATIONS_STR, fileRead && parsedIterationNumber>0, parsedIterationNumber, DEFAULT_ITERATION_NUMBER);
	printParameter(SIZE_OF_ARRAY_STR, fileRead && parsedArraySize>0, parsedArraySize, DEFAULT_ARRAY_LENGTH);
	printParametersFooter();
}

static int parseNumber(const char *parameters) {
	int result = -1;
	char c = (*parameters);

	if ((c > 0x2F) && (c < 0x3A)) {
		result = c - 0x30;
		parameters += 1;
		c = (*parameters);

		while ((c > 0x2F) && (c < 0x3A)) {
			result = result * 10 + c - 0x30;
			parameters += 1;
			c = (*parameters);
		}
	}

	return result;
}

static int parseIterNumber(const char *parameters) {
	return parseNumber(parameters);
}

static int parseArraySize(const char *parameters) {
	while (1) {
		const char c = (*parameters);

		if (c == 0)
			return -1;
		else if (c == '\n')
			return parseNumber(parameters + 1);
		else
			parameters += 1;
	}
}

static void parseParameters(const char *parameters, int *parsedIterationNumber, int *parsedArraySize) {
	(*parsedIterationNumber) = parseIterNumber(parameters);
	(*parsedArraySize) = parseArraySize(parameters);
}

static void setDefaultSettings(int *iterNumber, int *arrayLength) {
	(*iterNumber) = DEFAULT_ITERATION_NUMBER;
	(*arrayLength) = DEFAULT_ARRAY_LENGTH;
}

static void assignSettings(const char *parameters, int *parsedIterationNumber, int *parsedArraySize, int *iterNumber, int *arrayLength) {
	if (parameters)
		parseParameters(parameters, parsedIterationNumber, parsedArraySize);
	else
		setDefaultSettings(iterNumber, arrayLength);
}

static void correctSettings(const int parsedIterationNumber, const int parsedArraySize, int *iterNumber, int *arrayLength) {
	if (parsedIterationNumber > 0)
		(*iterNumber) = parsedIterationNumber;
	else
		(*iterNumber) = DEFAULT_ITERATION_NUMBER;

	if (parsedArraySize > 0)
		(*arrayLength) = parsedArraySize;
	else
		(*arrayLength) = DEFAULT_ARRAY_LENGTH;
}

static void initCharArray(char *arr, const char val, const int arrayLength) {
	char *currArr = arr;
	char *const arrLimit = arr + arrayLength;

	if (arrayLength > 16) {
		char *const arrLimit16 = arr + 16;

		for (; currArr < arrLimit16; currArr += 1)
			*currArr = val;

		for (currArr += 16; currArr < arrLimit; currArr += 16)
			memcpy(currArr - 16, currArr - 32, sizeof(char) * 16);
		currArr -= 16;
	}
	for (; currArr < arrLimit; currArr += 1)
		*currArr = val;
}


/* public functions */

void helper_printBenchmarkHeader() {
	printf("\n\n");
	printf("/////////////////////////////////////////////////////////////////\n");
	printf("//                                                             //\n");
	printf("//                       C benchmark                           //\n");
	printf("//                                                             //\n");
	printf("/////////////////////////////////////////////////////////////////\n");
}

void helper_printResultsHeader() {
	printf("//                                                             //\n");
	printf("//                         Results                             //\n");
	printf("//                                                             //\n");
}

void helper_printResultsFooter() {
	printf("//                                                             //\n");
	printf("/////////////////////////////////////////////////////////////////\n");
}

void helper_readParameters(int *iterNumber, int *arrayLength) {
	const char *fileText = readFile();
	int parsedIterationNumber = 0;
	int parsedArraySize = 0;

	assignSettings(fileText, &parsedIterationNumber, &parsedArraySize, iterNumber, arrayLength);
	correctSettings(parsedIterationNumber, parsedArraySize, iterNumber, arrayLength);
	printNotice(0 != fileText, parsedIterationNumber, parsedArraySize);
	printParameters(0 != fileText, parsedIterationNumber, parsedArraySize);
}

void helper_printTestTime(const char *testName, const double passedSeconds) {
	printf("// %25s:%12.3f seconds              //\n", testName, passedSeconds);
}

void helper_printOutOfMemory(const char *testName) {
	printf("// %25s:       error (out of memory)      //\n", testName);
}

void helper_benchmarkInsertRemove(int *const arr, const int iterNumber, const int arrLength) {
	const int arrLimit = arrLength - 1;
	int iterCurr;
	int *const arrNext = arr + 1;
	int *const arrNextNext = arr + 2;

	for (iterCurr = 0; iterCurr < iterNumber; iterCurr += 1) {
		int currLength, flag;

		// inserting
		for (currLength = 0, flag = 0; currLength < arrLimit; currLength += 1) {

			if (flag) {
				memmove(arrNextNext, arrNext, (size_t)(currLength - flag));
				(*arrNext) = currLength;
				flag = 0;
			}
			else {
				memmove(arrNext, arr, (size_t)(currLength - flag));
				(*arr) = currLength;
				flag = 1;
			}
		}

		// removing
		for (currLength = arrLimit, flag = arrLength % 2; currLength > 0; currLength -= 1) {

			if (flag) {
				memmove(arrNext, arrNextNext, (size_t)(currLength - flag));
				(*arrNext) = currLength;
				flag = 0;
			}
			else {
				memmove(arr, arrNext, (size_t)(currLength - flag));
				(*arr) = currLength;
				flag = 1;
			}
		}
	}
}

void helper_initPointerArray(char **arr, char *const val, const int arrayLength) {
	char **currArr = arr;
	char **const arrLimit = arr + arrayLength;

	if (arrayLength > 64) {
		char **const arrLimit64 = arr + 64;

		for (; currArr < arrLimit64; currArr += 1)
			*currArr = val;

		for (currArr += 64; currArr < arrLimit; currArr += 64)
			memcpy(currArr - 64, currArr - 128, sizeof(char*) * 64);
		currArr -= 64;
	}
	for (; currArr < arrLimit; currArr += 1)
		*currArr = val;
}

void helper_allocateStrings(char **arr, const int arrayLength) {
	const int maxValue = 127;
	const int alphabetArrayLength = arrayLength + maxValue - arrayLength % maxValue;
	const int stepLength = alphabetArrayLength / maxValue;

	int i, j;

	for (i = 0; i < arrayLength; i = j) {
		const int jLimit = i + stepLength;
		const char number = (char)(i / stepLength + 1);
		//printf("x %d %d \n", i, number);

		for (j = i; j < jLimit && j < arrayLength; j += 1) {
			const int wordLength = j % maxValue + 1;
			char *word = malloc(((size_t)(wordLength + 1))*sizeof(char));
			//printf("    %d %d \n", j, wordLength);

			if (word) {
				initCharArray(word, (char)number, wordLength);
				//printf(" %d ", j);
				word[wordLength] = 0;
				arr[j] = word;
			}
			else {
				return;
			}
		}
	}
}

int helper_validPointers(char **arr, const int arrayLength) {
	char **const arrLimit = arr + arrayLength;

	for (; arr < arrLimit; arr += 1)
		if (*arr == 0)
			return 0;
	return 1;
}

void helper_allocateWords(char **arr, const int arrayLength) {
	int i, j;

	for (i = 0; i < arrayLength;) {
		for (j = -2; (j < 128) && (i < arrayLength); j += 3, i += 1) {
			char *word = (char*)malloc(sizeof(char) * (10+1));

			if (word) {
				initCharArray(word, (char)j, 10);
				word[10] = 0;
				arr[i] = word;
			}
			else {
				return;
			}
		}
	}
}


int help_BinarySearch(size_t *const out_index, const char *const *const in_stringArray, const size_t in_length, const char *const in_string) {
	if (0 < in_length) {
		const size_t lastElementIndex = in_length - 1;
		const char *const *leftEl = in_stringArray;
		const char *const *rightEl = in_stringArray + lastElementIndex;
		const char *const *midEl;

		// iterate over all strings
		while (1) {
			midEl = leftEl + ((rightEl - leftEl) / 2);
			const char *el1 = *midEl;
			const char *el2 = in_string;

			// iterate over all characters
			while (1) {
				const char c1 = *el1;
				const char c2 = *el2;

				// string 1 ends
				if (0 == c1) {

					// string 2 ends: return match
					if (0 == c2) {
						if (0 != out_index)
							*out_index = (size_t)(midEl - in_stringArray);
						return 1;
					}

					// string 2 longer
					else {
						leftEl = midEl + 1;

						// limit exceeded: return
						if (leftEl > rightEl) {
							if (0 != out_index)
								*out_index = (size_t)(leftEl - in_stringArray);
							return 0;
						}

						// more elements: continue on right side
						else {
							break;
						}
					}
				}

				// string 2 ends: continue on left side
				else if (0 == c2) {
					rightEl = midEl - 1;

					// limit exceeded: return
					if (leftEl > rightEl) {
						if (0 != out_index)
							*out_index = (size_t)(leftEl - in_stringArray);
						return 0;
					}

					// more elements: continue on left side
					else {
						break;
					}
				}

				// character smaller: continue on right side
				else if (c1 < c2) {
					leftEl = midEl + 1;

					// limit exceeded: return
					if (leftEl > rightEl) {
						if (0 != out_index)
							*out_index = (size_t)(leftEl - in_stringArray);
						return 0;
					}

					// more elements: continue on right side
					else {
						break;
					}
				}

				// character bigger: continue on left side
				else if (c1 > c2) {
					rightEl = midEl - 1;

					// limit exceeded: return
					if (leftEl > rightEl) {
						if (0 != out_index)
							*out_index = (size_t)(leftEl - in_stringArray);
						return 0;
					}

					// more elements: continue on left side
					else {
						break;
					}
				}

				// characters equal: continue to compare
				else {
					el1 += 1;
					el2 += 1;
				}
			}
		}
	}

	else {
		if (0 != out_index)
			*out_index = 0;
		return 0;
	}
}

