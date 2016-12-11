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

#include "stdafx.h"


const char *const INSERT_REMOVE_TEST_NAME = "insert/remove";
const char *const BINARY_SEARCH_TEST_NAME = "binary search";


/* external functions */

void helper_printBenchmarkHeader();
void helper_printResultsHeader();
void helper_printResultsFooter();
void helper_readParameters(int *iterNumber, int *arraySize);
void helper_printTestTime(const char *testName, const double passedSeconds);
void helper_printOutOfMemory(const char *testName);
void helper_benchmarkInsertRemove(int *const arr, const int iterNumber, const size_t arrLength);
void helper_initPointerArray(char **arr, char *const val, const int arrayLength);
void helper_allocateStrings(char **arr, const int arrayLength);
int helper_validPointers(char **arr, const int arrayLength);
void helper_allocateWords(char **arr, const int arrayLength);
int help_BinarySearch(size_t *const out_index, const char *const *const in_stringArray, const size_t in_length, const char *const in_string);


/* benchmark functions */

static double getTimeInSeconds() {
	return omp_get_wtime();
}

static void benchmarkInsertRemove(const int iterNumber, const int arrayLength) {
	int *arr = (int*)malloc(((size_t) arrayLength) * sizeof(int));

	if (arr) {
		const double startTime = getTimeInSeconds();

		helper_benchmarkInsertRemove(arr, iterNumber, arrayLength);
		helper_printTestTime(INSERT_REMOVE_TEST_NAME, (getTimeInSeconds() - startTime) / iterNumber);
		free(arr);
	}
	else {
		helper_printOutOfMemory(INSERT_REMOVE_TEST_NAME);
	}
}

static void benchmarkBinarySearch(const int iterNumber, const int arrayLength) {
	const int wordsToSearchLength = 32766;
	char **arr = (char**)malloc(((size_t)arrayLength) * sizeof(char*));

	if (arr) {
		char **currStr = arr;
		char **const strLimit = arr + arrayLength;

		helper_initPointerArray(arr, 0, arrayLength);
		helper_allocateStrings(arr, arrayLength);

		if (helper_validPointers(arr, arrayLength)) {
			char **wordsToSearch = (char**)malloc((size_t)wordsToSearchLength * sizeof(char*));

			if (wordsToSearch) {
				char **currWordStr = wordsToSearch;
				char **const wordStrLimit = wordsToSearch + wordsToSearchLength;

				helper_initPointerArray(wordsToSearch, 0, wordsToSearchLength);
				helper_allocateWords(wordsToSearch, wordsToSearchLength);
				//printf("\n%d\n", wordsToSearch[1036][0]);

				if (helper_validPointers(wordsToSearch, wordsToSearchLength)) {
					int i, j;

					/* testing */
					const double startTime = getTimeInSeconds();
					const char aValue = wordsToSearch[0][0];
					const char bValue = aValue * 2;

					for (i = 0; i < iterNumber; i++) {
						for (j = 0; j < wordsToSearchLength; j++) {
							if (help_BinarySearch(0, arr, arrayLength, wordsToSearch[j])) {
								wordsToSearch[0][0] = aValue;
							}
							else {
								wordsToSearch[0][0] = bValue;
							}
						}
					}

					helper_printTestTime(BINARY_SEARCH_TEST_NAME, (getTimeInSeconds() - startTime) / iterNumber);
					/* end testing */

					for (currStr = arr; currStr < strLimit; currStr += 1)
						free(*currStr);
					free(arr);
					for (currWordStr = wordsToSearch; currWordStr < wordStrLimit; currWordStr += 1)
						free(*currWordStr);
					free(wordsToSearch);
				}
				else {
					for (currStr = arr; currStr < strLimit; currStr += 1)
						free(*currStr);
					free(arr);
					for (currWordStr = wordsToSearch; currWordStr < wordStrLimit; currWordStr += 1)
						if (*currWordStr != 0)
							free(*currWordStr);
					free(wordsToSearch);
					helper_printOutOfMemory(BINARY_SEARCH_TEST_NAME);
				}
			}
			else {
				for (currStr = arr; currStr < strLimit; currStr += 1)
					free(*currStr);
				free(arr);
				free(wordsToSearch);
				helper_printOutOfMemory(BINARY_SEARCH_TEST_NAME);
			}
		}
		else {
			for (currStr = arr; currStr < strLimit; currStr += 1)
				if (*currStr != 0)
					free(*currStr);
			free(arr);
			helper_printOutOfMemory(BINARY_SEARCH_TEST_NAME);
		}
	}
	else {
		helper_printOutOfMemory(BINARY_SEARCH_TEST_NAME);
	}
}

static void benchmark(const int iterNumber, const int arrayLength) {
	benchmarkInsertRemove(iterNumber, arrayLength);
	benchmarkBinarySearch(iterNumber, arrayLength);
}

int _tmain(int argc, _TCHAR* argv[]) {
	int iterNumber = 0;
	int arrayLength = 0;

	helper_printBenchmarkHeader();
	helper_readParameters(&iterNumber, &arrayLength);
	helper_printResultsHeader();
	benchmark(iterNumber, arrayLength);
	helper_printResultsFooter();
	printf("\n\npress any key...");
	_getch();

	return 0;
}

