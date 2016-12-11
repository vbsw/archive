
//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


#include <stdlib.h>
#include <string.h>

#include "strmap.h"


void
StrMap_appendParse ( StrMap *const strMap, const char * data, const size_t strSize ) {
	const char *const limit = data+strSize;

	// each iteration is a new line
	for ( ; data<limit; data+=1 ) {
		const char c = *data;

		// new line unix
		if ( '\n'==c ) {
			data += 1;

		// new line windows
		} else if ( '\r'==c && data+1<limit && '\n'==*(data+1) ) {
			data += 2;

		// white space
		} else if ( ' '==c || '\t'==c ) {
			const char *dataN = data+1;
			for ( ; dataN<limit && '\n'!=*dataN; ++dataN );
			data = dataN;

		// string pair
		} else {
			const char *keyEnd = data;

			// parse key
			while ( 1 ) {
				if ( keyEnd<limit ) {
					const char cEnd = *keyEnd;
					const char *const keyEndN = keyEnd+1;

					// new line unix
					if ( '\n'==cEnd ) {
						StrMap_append ( strMap, data, (size_t) (keyEnd-data), 0, 0 );
						data = keyEnd;
						break;

					// new line windows
					} else if ( '\r'==cEnd && keyEndN<limit && '\n'==*keyEndN ) {
						StrMap_append ( strMap, data, (size_t) (keyEnd-data), 0, 0 );
						data = keyEndN;
						break;

					// key end
					} else if ( '='==cEnd ) {
						const char *const valueStart = keyEnd+1;

						// parse value
						if ( valueStart<limit ) {
							const char *valueEnd = valueStart;

							// search end value
							while ( 1 ) {
								if ( valueEnd<limit ) {
									const char cEnd = *valueEnd;
									const char *valueEndN = valueEnd+1;

									// new line unix
									if ( '\n'==cEnd ) {
										StrMap_append ( strMap, data, (size_t) (keyEnd-data), valueStart, (size_t) (valueEnd-valueStart) );
										data = valueEnd;
										break;

									// new line windows
									} else if ( '\r'==cEnd && valueEndN<limit && '\n'==*valueEndN ) {
										StrMap_append ( strMap, data, (size_t) (keyEnd-data), valueStart, (size_t) (valueEnd-valueStart) );
										data = valueEndN;
										break;

									// continue search
									} else {
										valueEnd += 1;
									}

								// file end
								} else {
									StrMap_append ( strMap, data, (size_t) (keyEnd-data), valueStart, (size_t) (valueEnd-valueStart) );
									data = valueEnd;
									break;
								}
							}

						// file end
						} else {
							StrMap_append ( strMap, data, (size_t) (keyEnd-data), 0, 0 );
							data = keyEnd;
						}
						break;

					// continue search
					} else {
						keyEnd += 1;
					}

				// file end
				} else {
					StrMap_append ( strMap, data, (size_t) (keyEnd-data), 0, 0 );
					data = keyEnd;
					break;
				}
			} // parse key end
		} // string pair end
	} // parse data end
}


void
StrMap_writeFile ( StrMap *const strMap, FILE *const file ) {
	void **const keysLimit = (*strMap).keys.elements + (*strMap).keys.len;
	void **keyElements = (*strMap).keys.elements;
	void **valueElements = (*strMap).values.elements;
	size_t *keySizes = (*strMap).keys.sizes;
	size_t *valueSizes = (*strMap).values.sizes;

	for ( ; keyElements<keysLimit; ++keyElements, ++valueElements, ++keySizes, ++valueSizes ) {
		fwrite ( *keyElements, 1, *keySizes, file );
		fputc ( '=', file );
		fwrite ( *valueElements, 1, *valueSizes, file );
		fputc ( '\n', file );
	}
	fflush ( file );
}


void
StrMap_clone ( StrMap *const srcStrMap, StrMap *const destStrMap ) {
	const size_t keysLen = (*srcStrMap).keys.len;
	const size_t keysCap = (*srcStrMap).keys.cap;
	const size_t valuesLen = (*srcStrMap).values.len;
	const size_t valuesCap = (*srcStrMap).values.cap;

	void **const newKeysElements = (void **const) malloc ( keysCap * sizeof(void**) );
	size_t *const newKeysSizes = (size_t *const) malloc ( keysCap * sizeof(size_t*) );
	void **const newValuesElements = (void **const) malloc ( valuesCap * sizeof(void**) );
	size_t *const newValuesSizes = (size_t *const) malloc ( valuesCap * sizeof(size_t*) );

	memcpy ( newKeysElements, (*srcStrMap).keys.elements, keysLen * sizeof(void*) );
	memcpy ( newKeysSizes, (*srcStrMap).keys.sizes, keysLen * sizeof(size_t) );
	memcpy ( newValuesElements, (*srcStrMap).values.elements, valuesLen * sizeof(void*) );
	memcpy ( newValuesSizes, (*srcStrMap).values.sizes, valuesLen * sizeof(size_t) );

	(*destStrMap).keys.elements = newKeysElements;
	(*destStrMap).keys.sizes = newKeysSizes;
	(*destStrMap).keys.len = keysLen;
	(*destStrMap).keys.cap = keysCap;
	(*destStrMap).values.elements = newValuesElements;
	(*destStrMap).values.sizes = newValuesSizes;
	(*destStrMap).values.len = valuesLen;
	(*destStrMap).values.cap = valuesCap;
}


size_t
StrMap_search ( StrMap *const strMap, const char *const key, const size_t keySize ) {
	void **keys = (*strMap).keys.elements;
	size_t *sizes = (size_t *) (*strMap).keys.sizes;
	void **const limit = keys + (*strMap).keys.len;

	for ( ; keys<limit; keys+=1, sizes+=1 ) {
		const size_t size = *sizes;

		if ( keySize==size ) {
			const char *key1 = (const char *) *keys;
			const char *key2 = key;
			const char *const kLimit = key1+keySize;

			for ( ; key1<kLimit && *key1==*key2; key1+=1, key2+=1 );

			if ( key1==kLimit )
				return (size_t) (keys-(*strMap).keys.elements);
		}
	}
	return (*strMap).keys.len;
}


size_t
StrMap_searchAsc ( StrMap *const strMap, size_t *const index, const char *const key, const size_t keySize ) {
	const size_t lastKeyIndex = (*strMap).keys.len-1;
	void **midKey;
	void **leftKey = (*strMap).keys.elements;
	void **rightKey = (*strMap).keys.elements + lastKeyIndex;
	size_t *midSize;
	size_t *leftSize = (*strMap).keys.sizes;
	size_t *rightSize = (*strMap).keys.sizes + lastKeyIndex;

	while ( leftKey<=rightKey ) {
		midKey = leftKey + ( (rightKey-leftKey) / 2 );
		midSize = leftSize + ( (rightSize-leftSize) / 2 );
		const size_t midSizeVal = (*midSize);
		const char *key1 = (const char*) *midKey;
		const char *key2 = (const char*) key;

		// compare elements

		if ( midSizeVal==keySize ) {
			const char *const limit = key1+keySize;

			while ( key1<limit ) {

				// continue in right side
				if ( *key1<*key2 ) {
					leftKey = midKey+1;
					leftSize = midSize+1;
					break;

				// continue in left side
				} else if ( *key1>*key2 ) {
					rightKey = midKey-1;
					rightSize = midSize-1;
					break;
				}
				++key1;
				++key2;
			}

			// match
			if ( key1==limit ) {
				if ( 0!=index )
					*index = (size_t) (midKey-(*strMap).keys.elements);
				return 1;
			}

		} else if ( midSizeVal<keySize ) {
			const char *const limit = key1+midSizeVal;

			while ( key1<limit ) {

				// continue in right side
				if ( *key1<*key2 ) {
					leftKey = midKey+1;
					leftSize = midSize+1;
					break;

				// continue in left side
				} else if ( *key1>*key2 ) {
					rightKey = midKey-1;
					rightSize = midSize-1;
					break;
				}
				++key1;
				++key2;
			}

			// match, continue in right side
			if ( key1==limit ) {
				leftKey = midKey+1;
				leftSize = midSize+1;
			}

		} else {
			const char *const limit = key1+keySize;

			while ( key1<limit ) {

				// continue in right side
				if ( *key1<*key2 ) {
					leftKey = midKey+1;
					leftSize = midSize+1;
					break;

				// continue in left side
				} else if ( *key1>*key2 ) {
					rightKey = midKey-1;
					rightSize = midSize-1;
					break;
				}
				++key1;
				++key2;
			}

			// match, continue in left side
			if ( key1==limit ) {
				rightKey = midKey-1;
				rightSize = midSize-1;
			}
		}
	}

	if ( 0!=index )
		*index = (size_t) ( ( leftKey==midKey ? midKey : leftKey ) - (*strMap).keys.elements );
	return 0;
}


size_t
StrMap_searchDesc ( StrMap *const strMap, size_t *const index, const char *const key, const size_t keySize ) {
	const size_t lastKeyIndex = (*strMap).keys.len-1;
	void **midKey;
	void **leftKey = (*strMap).keys.elements;
	void **rightKey = (*strMap).keys.elements + lastKeyIndex;
	size_t *midSize;
	size_t *leftSize = (*strMap).keys.sizes;
	size_t *rightSize = (*strMap).keys.sizes + lastKeyIndex;

	while ( leftKey<=rightKey ) {
		midKey = leftKey + ( (rightKey-leftKey) / 2 );
		midSize = leftSize + ( (rightSize-leftSize) / 2 );
		const size_t midSizeVal = (*midSize);
		const char *key1 = (const char*) *midKey;
		const char *key2 = (const char*) key;

		// compare elements

		if ( midSizeVal==keySize ) {
			const char *const limit = key1+keySize;

			while ( key1<limit ) {

				// continue in left side
				if ( *key1<*key2 ) {
					rightKey = midKey-1;
					rightSize = midSize-1;
					break;

				// continue in right side
				} else if ( *key1>*key2 ) {
					leftKey = midKey+1;
					leftSize = midSize+1;
					break;
				}
				++key1;
				++key2;
			}

			// match
			if ( key1==limit ) {
				if ( 0!=index )
					*index = (size_t) (midKey-(*strMap).keys.elements);
				return 1;
			}

		} else if ( midSizeVal<keySize ) {
			const char *const limit = key1+midSizeVal;

			while ( key1<limit ) {

				// continue in left side
				if ( *key1<*key2 ) {
					rightKey = midKey-1;
					rightSize = midSize-1;
					break;

				// continue in right side
				} else if ( *key1>*key2 ) {
					leftKey = midKey+1;
					leftSize = midSize+1;
					break;
				}
				++key1;
				++key2;
			}

			// match, continue in left side
			if ( key1==limit ) {
				rightKey = midKey-1;
				rightSize = midSize-1;
			}

		} else {
			const char *const limit = key1+keySize;

			while ( key1<limit ) {

				// continue in left side
				if ( *key1<*key2 ) {
					rightKey = midKey-1;
					rightSize = midSize-1;
					break;

				// continue in right side
				} else if ( *key1>*key2 ) {
					leftKey = midKey+1;
					leftSize = midSize+1;
					break;
				}
				++key1;
				++key2;
			}

			// match, continue in right side
			if ( key1==limit ) {
				leftKey = midKey+1;
				leftSize = midSize+1;
			}
		}
	}

	if ( 0!=index )
		*index = (size_t) ( ( leftKey==midKey ? midKey : leftKey ) - (*strMap).keys.elements );
	return 0;
}

























































