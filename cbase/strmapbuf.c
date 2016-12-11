
//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


#include <stdlib.h>
#include <string.h>

#include "strmapbuf.h"


void
StrMapBuf_appendParse ( StrMapBuf *const strMapBuf, const char * data, const size_t dataSize ) {
	const char *const limit = data+dataSize;

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
						StrMapBuf_append ( strMapBuf, data, (size_t) (keyEnd-data), 0, 0 );
						data = keyEnd;
						break;

					// new line windows
					} else if ( '\r'==cEnd && keyEndN<limit && '\n'==*keyEndN ) {
						StrMapBuf_append ( strMapBuf, data, (size_t) (keyEnd-data), 0, 0 );
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
										StrMapBuf_append ( strMapBuf, data, (size_t) (keyEnd-data), valueStart, (size_t) (valueEnd-valueStart) );
										data = valueEnd;
										break;

									// new line windows
									} else if ( '\r'==cEnd && valueEndN<limit && '\n'==*valueEndN ) {
										StrMapBuf_append ( strMapBuf, data, (size_t) (keyEnd-data), valueStart, (size_t) (valueEnd-valueStart) );
										data = valueEndN;
										break;

									// continue search
									} else {
										valueEnd += 1;
									}

								// file end
								} else {
									StrMapBuf_append ( strMapBuf, data, (size_t) (keyEnd-data), valueStart, (size_t) (valueEnd-valueStart) );
									data = valueEnd;
									break;
								}
							}

						// file end
						} else {
							StrMapBuf_append ( strMapBuf, data, (size_t) (keyEnd-data), 0, 0 );
							data = keyEnd;
						}
						break;

					// continue search
					} else {
						keyEnd += 1;
					}

				// file end
				} else {
					StrMapBuf_append ( strMapBuf, data, (size_t) (keyEnd-data), 0, 0 );
					data = keyEnd;
					break;
				}
			} // parse key end
		} // string pair end
	} // parse data end
}


void
StrMapBuf_readFile ( StrMapBuf *const strMapBuf, FILE *const file ) {
	fseek ( file, 0L, SEEK_END );
	const long int fsize = ftell ( file );
	fseek ( file, 0L, SEEK_SET );
	char *data = (char*) malloc ( fsize );
	fread ( data, 1, fsize, file );
	fseek ( file, 0L, SEEK_SET );
	StrMapBuf_appendParse ( strMapBuf, data, fsize );
	free ( (void*) data );
}


void
StrMapBuf_clone ( StrMapBuf *const srcStrMapBuf, StrMapBuf *const destStrMapBuf ) {
	size_t i;
	Buffer *const destBuffer = &(*destStrMapBuf).buffer;
	void **const keysElements = (*srcStrMapBuf).strMap.keys.elements;
	const size_t keysLen = (*srcStrMapBuf).strMap.keys.len;
	const size_t keysCap = (*srcStrMapBuf).strMap.keys.cap;
	void **const valuesElements = (*srcStrMapBuf).strMap.values.elements;
	const size_t valuesCap = (*srcStrMapBuf).strMap.values.cap;
	const size_t valuesLen = (*srcStrMapBuf).strMap.values.len;
	(*destStrMapBuf).strMap.keys.len = keysLen;
	(*destStrMapBuf).strMap.keys.cap = keysCap;
	(*destStrMapBuf).strMap.values.len = valuesLen;
	(*destStrMapBuf).strMap.values.cap = valuesCap;
	void **const newKeysElements = (void **const) malloc ( keysCap * sizeof(void**) );
	size_t *const newKeysSizes = (size_t *const) malloc ( keysCap * sizeof(size_t*) );
	void **const newValuesElements = (void **const) malloc ( valuesCap * sizeof(void**) );
	size_t *const newValuesSizes = (size_t *const) malloc ( valuesCap * sizeof(size_t*) );

	Buffer_allocate ( destBuffer, (*srcStrMapBuf).buffer.preferredChunkCapacity );
	memcpy ( newKeysSizes, (*srcStrMapBuf).strMap.keys.sizes, keysLen * sizeof(size_t) );
	memcpy ( newValuesSizes, (*srcStrMapBuf).strMap.values.sizes, valuesLen * sizeof(size_t) );

	for ( i=0; i<keysLen; i+=1 )
		newKeysElements[i] = Buffer_putc ( destBuffer, keysElements[i], newKeysSizes[i] );

	for ( i=0; i<valuesLen; i+=1 )
		newValuesElements[i] = Buffer_putc ( destBuffer, valuesElements[i], newValuesSizes[i] );

	// dest has no memory to free
	(*destStrMapBuf).strMap.keys.elements = newKeysElements;
	(*destStrMapBuf).strMap.keys.sizes = newKeysSizes;
	(*destStrMapBuf).strMap.values.elements = newValuesElements;
	(*destStrMapBuf).strMap.values.sizes = newValuesSizes;
}

























































