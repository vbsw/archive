
//          Copyright Vitali Baumtrok 2015.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


/**
 * @file cb.c
 * @author Vitali Baumtrok
 * @date 2015-04-04
 * @brief basic functions
 * 
 * Content: memory buffer allocator, string array binary search.
 */

#include <stdlib.h>
#include <stdio.h>

/**
 * @brief Computes rounded size.
 * 
 * The next best size is computed and returned. The retured size is
 * mod(size_of(size_t))==0.
 * 
 * @param in_size The size in bytes.
 * @return The new rounded size.
 */
static size_t
computeRoundedMemorySize ( const size_t in_size ) {
	const size_t rest = in_size % sizeof ( size_t );

	if ( rest )
		return in_size - rest + sizeof ( size_t );
	else
		return in_size;
}

/**
 * @brief Allocates a buffer of given size.
 * 
 * Allocates a buffer of given size. If not enough memory could be allocated
 * 0 is returned.
 * 
 * @param in_bufferSize The size of buffer in bytes.
 * @return The buffer.
 */
void *
cb_allocateBuffer ( const size_t in_bufferSize ) {
	const size_t bufferSize = computeRoundedMemorySize ( in_bufferSize );
	void *const buffer = (void *const) malloc ( bufferSize );

	if ( buffer ) {
		void **const bufferMeta = (void **const) malloc ( 4 * sizeof (void **) );

		if ( bufferMeta ) {
			bufferMeta[0] = buffer;
			bufferMeta[1] = buffer;
			bufferMeta[2] = buffer + bufferSize;
			bufferMeta[3] = 0;

			return (void*) bufferMeta;

		} else {
			free ( buffer );
			return 0;
		}

	} else {
		return 0;
	}
}

/**
 * @brief Frees the buffer.
 * 
 * All memory of the given buffer is released.
 * 
 * @param inout_buffer The buffer to deallocate.
 */
void
cb_freeBuffer ( void *const inout_buffer ) {
	void **bufferMeta = (void **) inout_buffer;
	void *buffer = bufferMeta[0];

	while (1) {
		void **nextBufferMeta = (void **) bufferMeta[3];

		free ( buffer );
		free ( bufferMeta );

		if ( nextBufferMeta ) {
			bufferMeta = nextBufferMeta;
			buffer = bufferMeta[0];

		} else {
			break;
		}
	}
}

/**
 * @brief Returns free memory from the buffer.
 * 
 * Searches for free memory in buffer and returns it. If no free memory is
 * found new memory is allocated. If no memory could be allocated 0 is
 * returned.
 * 
 * @param inout_buffer The buffer.
 * @param in_memorySize The memory size in bytes.
 */
void *
cb_getBufferMemory ( void *const inout_buffer, const size_t in_memorySize ) {
	void **bufferMeta = (void **) inout_buffer;
	const size_t bufferSize = (size_t) bufferMeta[2] - (size_t) bufferMeta[0];
	const size_t memorySize = computeRoundedMemorySize ( in_memorySize );

	// memory size matches in buffer size
	if ( bufferSize>=memorySize ) {
		while (1) {
			const size_t freeMemoryAvailableSize = (size_t) bufferMeta[2] - (size_t) bufferMeta[1];

			// free memory found
			if ( freeMemoryAvailableSize >= memorySize ) {
				void *result = bufferMeta[1];
				bufferMeta[1] += memorySize;

				return result;

			// continue search
			} else {
				void **nextBufferMeta = (void **) bufferMeta[3];

				// more memory available to check
				if ( nextBufferMeta ) {
					bufferMeta = nextBufferMeta;

				// new memory needed
				} else {
					nextBufferMeta = (void **) cb_allocateBuffer ( bufferSize );

					// return valid memory
					if ( nextBufferMeta ) {
						bufferMeta[3] = (void *) nextBufferMeta;
						nextBufferMeta[1] += memorySize;

						return nextBufferMeta[0];

					// return no memory
					} else {
						return 0;
					}
				}
			}
		}

	// new buffer chunk must be allocated
	} else {
		// search for last chunk
		while ( bufferMeta[3] ) {
			bufferMeta = (void **) bufferMeta[3];
		}

		// allocate buffer
		bufferMeta[3] = (void *) cb_allocateBuffer ( memorySize );
		bufferMeta = (void **) bufferMeta[3];

		// return valid memory
		if ( bufferMeta ) {
			// because the requested size is the same as allocated size
			bufferMeta[1] = bufferMeta[2];

			return bufferMeta[0];

		// return no memory
		} else {
			return 0;
		}
	}
}

/**
 * @brief Searches for a c-string in an array of ascending order.
 * 
 * Searches for a c-string in an array of ascending order.
 * Ascending order is ["a", "ax", "b", "bx"]. This is a binary search.
 * 
 * @param out_index Output parameter! A pointer to an unsigned integer which
 * is the index where the string would fit. The pointer can be zero.
 * @param in_stringArray The array where to search for the string. Must not
 * be zero.
 * @param in_length The length of in_stringArray.
 * @param in_string The string to search for.
 * @return 1 if match, otherwise 0.
 */
int
cb_stringBinarySearchAsc ( size_t *const out_index, const char *const *const in_stringArray, const size_t in_length, const char *const in_string ) {
	if ( 0 < in_length ) {
		const size_t lastElementIndex = in_length - 1;
		const char *const *leftEl = in_stringArray;
		const char *const *rightEl = in_stringArray + lastElementIndex;
		const char *const *midEl;

		// iterate over all strings
		while ( 1 ) {
			midEl = leftEl + ( (rightEl-leftEl) / 2 );
			const char *el1 = *midEl;
			const char *el2 = in_string;

			// iterate over all characters
			while ( 1 ) {
				const char c1 = *el1;
				const char c2 = *el2;

				// string 1 ends
				if ( 0 == c1 ) {

					// string 2 ends: return match
					if ( 0 == c2 ) {
						if ( 0 != out_index )
							*out_index = (size_t) (midEl - in_stringArray);
						return 1;
					}

					// string 2 longer
					else {
						leftEl = midEl + 1;

						// limit exceeded: return
						if ( leftEl > rightEl ) {
							if ( 0 != out_index )
								*out_index = (size_t) (leftEl - in_stringArray);
							return 0;
						}

						// more elements: continue on right side
						else {
							break;
						}
					}
				}

				// string 2 ends: continue on left side
				else if ( 0 == c2 ) {
					rightEl = midEl - 1;

					// limit exceeded: return
					if ( leftEl > rightEl ) {
						if ( 0 != out_index )
							*out_index = (size_t) (leftEl - in_stringArray);
						return 0;
					}

					// more elements: continue on left side
					else {
						break;
					}
				}

				// character smaller: continue on right side
				else if ( c1 < c2 ) {
					leftEl = midEl + 1;

					// limit exceeded: return
					if ( leftEl > rightEl ) {
						if ( 0 != out_index )
							*out_index = (size_t) (leftEl - in_stringArray);
						return 0;
					}

					// more elements: continue on right side
					else {
						break;
					}
				}

				// character bigger: continue on left side
				else if ( c1 > c2 ) {
					rightEl = midEl - 1;

					// limit exceeded: return
					if ( leftEl > rightEl ) {
						if ( 0 != out_index )
							*out_index = (size_t) (leftEl - in_stringArray);
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
		if ( 0 != out_index )
			*out_index = 0;
		return 0;
	}
}

/**
 * @brief Searches for a c-string in an array of descending order.
 * 
 * Searches for a c-string in an array of descending order.
 * Descending order is ["bx", "b", "ax", "a"]. This is a binary search.
 * 
 * @param out_index Output parameter! A pointer to an unsigned integer which
 * is the index where the string would fit. The pointer can be zero.
 * @param in_stringArray The array where to search for the string. Must not
 * be zero.
 * @param in_length The length of in_stringArray.
 * @param in_string The string to search for.
 * @return 1 if match, otherwise 0.
 */
int
cb_stringBinarySearchDesc ( size_t *const out_index, const char *const *const in_stringArray, const size_t in_length, const char *const in_string ) {
	if ( 0 < in_length ) {
		const size_t lastElementIndex = in_length - 1;
		const char *const *leftEl = in_stringArray;
		const char *const *rightEl = in_stringArray + lastElementIndex;
		const char *const *midEl;

		// iterate over all strings
		while ( 1 ) {
			midEl = leftEl + ( (rightEl-leftEl) / 2 );
			const char *el1 = *midEl;
			const char *el2 = in_string;

			// iterate over all characters
			while ( 1 ) {
				const char c1 = *el1;
				const char c2 = *el2;

				// string 1 ends
				if ( 0 == c1 ) {

					// string 2 ends: return match
					if ( 0 == c2 ) {
						if ( 0 != out_index )
							*out_index = (size_t) (midEl - in_stringArray);
						return 1;
					}

					// string 2 longer
					else {
						rightEl = midEl - 1;

						// limit exceeded: return
						if ( leftEl > rightEl ) {
							if ( 0 != out_index )
								*out_index = (size_t) (leftEl - in_stringArray);
							return 0;
						}

						// more elements: continue on left side
						else {
							break;
						}
					}
				}

				// string 2 ends: continue on right side
				else if ( 0 == c2 ) {
					leftEl = midEl + 1;

					// limit exceeded: return
					if ( leftEl > rightEl ) {
						if ( 0 != out_index )
							*out_index = (size_t) (leftEl - in_stringArray);
						return 0;
					}

					// more elements: continue on right side
					else {
						break;
					}
				}

				// character smaller: continue on left side
				else if ( c1 < c2 ) {
					rightEl = midEl - 1;

					// limit exceeded: return
					if ( leftEl > rightEl ) {
						if ( 0 != out_index )
							*out_index = (size_t) (leftEl - in_stringArray);
						return 0;
					}

					// more elements: continue on left side
					else {
						break;
					}
				}

				// character bigger: continue on right side
				else if ( c1 > c2 ) {
					leftEl = midEl + 1;

					// limit exceeded: return
					if ( leftEl > rightEl ) {
						if ( 0 != out_index )
							*out_index = (size_t) (leftEl - in_stringArray);
						return 0;
					}

					// more elements: continue on right side
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
		if ( 0 != out_index )
			*out_index = 0;
		return 0;
	}
}
