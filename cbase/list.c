
//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include <stdio.h>


typedef struct {
	void **elements;
	size_t *sizes;
	size_t len;
	size_t cap;
} List;


// capacity in number of elements
#define LIST_INITIAL_CAP 7


void
List_allocate ( List *const list ) {
	(*list).elements = (void**) malloc ( LIST_INITIAL_CAP * sizeof(void*) );
	(*list).sizes = (size_t*) malloc ( LIST_INITIAL_CAP * sizeof(size_t) );
	(*list).len = 0;
	(*list).cap = LIST_INITIAL_CAP;
}


void
List_deallocate ( List *const list ) {
	free ( (*list).elements );
	free ( (*list).sizes );
}


void
List_append ( List *const list, void *const element, const size_t elementSize ) {
	const size_t len = (*list).len;
	const size_t cap = (*list).cap;

	// ensure capacity
	if ( len==cap ) {
		const size_t oldElementsSize = cap * sizeof(void*);
		const size_t oldLengthsSize = cap * sizeof(size_t);
		const size_t newElementsSize = oldElementsSize*2;
		const size_t newLengthsSize = oldLengthsSize*2;

		void **const newElements = (void**) malloc ( newElementsSize );
		size_t *const newSizes = (size_t*) malloc ( newLengthsSize );

		memcpy ( newElements, (*list).elements, oldElementsSize );
		memcpy ( newSizes, (*list).sizes, oldLengthsSize );

		free ( (*list).elements );
		free ( (*list).sizes );

		(*list).elements = newElements;
		(*list).sizes = newSizes;
		(*list).cap = cap*2;
	}

	(*list).elements[len] = element;
	(*list).sizes[len] = elementSize;
	(*list).len += 1;
}


void
List_insert ( List *const list, const size_t index, void *const element, const size_t elementSize ) {
	const size_t len = (*list).len;
	const size_t cap = (*list).cap;

	// ensure capacity
	if ( len==cap ) {
		const size_t oldElementsSize = cap * sizeof(void*);
		const size_t oldLengthsSize = cap * sizeof(size_t);
		const size_t newElementsSize = oldElementsSize*2;
		const size_t newLengthsSize = oldLengthsSize*2;

		void **const newElements = (void**const) malloc ( newElementsSize );
		size_t *const newSizes = (size_t*const) malloc ( newLengthsSize );

		memcpy ( newElements, (*list).elements, oldElementsSize );
		memcpy ( newSizes, (*list).sizes, oldLengthsSize );

		free ( (*list).elements );
		free ( (*list).sizes );

		(*list).elements = newElements;
		(*list).sizes = newSizes;
		(*list).cap = cap*2;
	}

	// append
	if ( len==index ) {
		(*list).elements[len] = element;
		(*list).sizes[len] = elementSize;
		(*list).len += 1;

	// insert
	} else {
		void **const elements = (*list).elements;
		void **const elementsInsert = elements+index;
		size_t *const sizes = (*list).sizes;
		size_t *const sizesInsert = sizes+index;
		const size_t moveLength = len-index;

		memmove ( elementsInsert+1, elementsInsert, moveLength * sizeof(void*) );
		memmove ( sizesInsert+1, sizesInsert, moveLength * sizeof(size_t) );

		(*elementsInsert) = element;
		(*sizesInsert) = elementSize;
	}
}


void
List_remove ( List *const list, const size_t index ) {
	const size_t len = (*list).len;
	const size_t nextIndex = index+1;

	if ( len>nextIndex ) {
		void **const elements = (*list).elements;
		void **const elementsRemove = elements+index;
		size_t *const sizes = (*list).sizes;
		size_t *const sizesRemove = sizes+index;
		const size_t moveLength = len-nextIndex;

		memmove ( elementsRemove, elementsRemove+1, moveLength * sizeof(void*) );
		memmove ( sizesRemove, sizesRemove+1, moveLength * sizeof(size_t) );
	}

	--(*list).len;
}


void *
List_elementClone ( List *const list, const size_t index, const size_t elementSize ) {
	void *const element = malloc ( elementSize );
	memcpy ( element, (*list).elements[index], elementSize );
	return element;
}


size_t
List_search ( List *const list, const void *const element, const size_t elementSize ) {
	void **elements = (*list).elements;
	void **const limit = elements+(*list).len;
	size_t *sizes = (*list).sizes;

	if ( 8<=sizeof(size_t) ) {

		while ( elements<limit ) {

			if ( (*sizes)==elementSize ) {
				const size_t cmpLength = elementSize/8;
				const uint64_t *el1 = (const uint64_t*) (*elements);
				const uint64_t *el2 = (const uint64_t*) element;
				const uint64_t *const limit = (const uint64_t*) (el1+cmpLength);

				// compare both elements over cmpLength
				while ( el1<limit ) {

					if ( (*el1)!=(*el2) )
						break;
					++el1;
					++el2;
				}

				// compare the rest bytes if previous bytes matches
				if ( el1==limit ) {
					const size_t bytesLength = elementSize - (cmpLength*8);
					
					if ( bytesLength ) {
						const char *el1Bytes = (const char *)(el1-1);
						const char *el2Bytes = (const char *)(el2-1);
						const char *const limitBytes = el1Bytes+bytesLength;

						while ( el1Bytes<limitBytes ) {

							if ( (*el1Bytes)!=(*el2Bytes) )
								break;
							++el1Bytes;
							++el2Bytes;
						}

						// element was found
						if ( el1Bytes==limitBytes )
							return (size_t) (elements-(*list).elements);

					} else {
						return (size_t) (elements-(*list).elements);
					}
				}
			}
			++elements;
			++sizes;
		}

	} else if ( 4==sizeof(size_t) ) {

		while ( elements<limit ) {

			if ( (*sizes)==elementSize ) {
				const size_t cmpLength = elementSize/4;
				const uint32_t *el1 = (const uint32_t*) (*elements);
				const uint32_t *el2 = (const uint32_t*) element;
				const uint32_t *const limit = (const uint32_t*) (el1+cmpLength);

				// compare both elements over cmpLength
				while ( el1<limit ) {

					if ( (*el1)!=(*el2) )
						break;
					++el1;
					++el2;
				}

				// until now both elements match; compare bytes
				if ( el1==limit ) {
					const size_t bytesLength = elementSize - (cmpLength*4);
					
					if ( bytesLength ) {
						const char *el1Bytes = (const char *)(el1-1);
						const char *el2Bytes = (const char *)(el2-1);
						const char *const limitBytes = el1Bytes+bytesLength;

						while ( el1Bytes<limitBytes ) {

							if ( (*el1Bytes)!=(*el2Bytes) )
								break;
							++el1Bytes;
							++el2Bytes;
						}

						// element was found
						if ( el1Bytes==limitBytes )
							return (size_t) (elements-(*list).elements);

					} else {
						return (size_t) (elements-(*list).elements);
					}
				}
			}
			++elements;
			++sizes;
		}

	} else if ( 2==sizeof(size_t) ) {

		while ( elements<limit ) {

			if ( (*sizes)==elementSize ) {
				const size_t cmpLength = elementSize/2;
				const uint16_t *el1 = (const uint16_t*) (*elements);
				const uint16_t *el2 = (const uint16_t*) element;
				const uint16_t *const limit = (const uint16_t*) (el1+cmpLength);

				// compare both elements over cmpLength
				while ( el1<limit ) {

					if ( (*el1)!=(*el2) )
						break;
					++el1;
					++el2;
				}

				// until now both elements match; compare bytes
				if ( el1==limit ) {
					const size_t bytesLength = elementSize - (cmpLength*2);
					
					if ( bytesLength ) {
						const char *el1Bytes = (const char *)(el1-1);
						const char *el2Bytes = (const char *)(el2-1);
						const char *const limitBytes = el1Bytes+bytesLength;

						while ( el1Bytes<limitBytes ) {

							if ( (*el1Bytes)!=(*el2Bytes) )
								break;
							++el1Bytes;
							++el2Bytes;
						}

						// element was found
						if ( el1Bytes==limitBytes )
							return (size_t) (elements-(*list).elements);

					} else {
						return (size_t) (elements-(*list).elements);
					}
				}
			}
			++elements;
			++sizes;
		}

	} else {

		while ( elements<limit ) {

			if ( (*sizes)==elementSize ) {
				const char *el1 = (const char*) (*elements);
				const char *el2 = (const char*) element;
				const char *const limit = (const char*) (el1+elementSize);

				// compare both elements
				while ( el1<limit ) {

					if ( (*el1)!=(*el2) )
						break;
					++el1;
					++el2;
				}

				// element was found
				if ( el1==limit )
					return (size_t) (elements-(*list).elements);
			}
			++elements;
			++sizes;
		}
	}

	return (*list).len;

}


int
List_searchAsc ( List *const list, size_t *const index, const void *const element, const size_t elementSize ) {
	const size_t lastElementIndex = (*list).len-1;
	void **midEl;
	void **leftEl = (*list).elements;
	void **rightEl = (*list).elements + lastElementIndex;
	size_t *midSize;
	size_t *leftSize = (*list).sizes;
	size_t *rightSize = (*list).sizes + lastElementIndex;

	while ( leftEl<=rightEl ) {
		midEl = leftEl + ( (rightEl-leftEl) / 2 );
		midSize = leftSize + ( (rightSize-leftSize) / 2 );
		const size_t midSizeVal = (*midSize);
		const char *el1 = (const char*) *midEl;
		const char *el2 = (const char*) element;

		// compare elements

		if ( midSizeVal==elementSize ) {
			const char *const limit = el1+elementSize;

			while ( el1<limit ) {

				// continue in right side
				if ( *el1<*el2 ) {
					leftEl = midEl+1;
					leftSize = midSize+1;
					break;

				// continue in left side
				} else if ( *el1>*el2 ) {
					rightEl = midEl-1;
					rightSize = midSize-1;
					break;
				}
				++el1;
				++el2;
			}

			// match
			if ( el1==limit ) {
				if ( 0!=index )
					*index = (size_t) (midEl-(*list).elements);
				return 1;
			}

		} else if ( midSizeVal<elementSize ) {
			const char *const limit = el1+midSizeVal;

			while ( el1<limit ) {

				// continue in right side
				if ( *el1<*el2 ) {
					leftEl = midEl+1;
					leftSize = midSize+1;
					break;

				// continue in left side
				} else if ( *el1>*el2 ) {
					rightEl = midEl-1;
					rightSize = midSize-1;
					break;
				}
				++el1;
				++el2;
			}

			// match, continue in right side
			if ( el1==limit ) {
				leftEl = midEl+1;
				leftSize = midSize+1;
			}

		} else {
			const char *const limit = el1+elementSize;

			while ( el1<limit ) {

				// continue in right side
				if ( *el1<*el2 ) {
					leftEl = midEl+1;
					leftSize = midSize+1;
					break;

				// continue in left side
				} else if ( *el1>*el2 ) {
					rightEl = midEl-1;
					rightSize = midSize-1;
					break;
				}
				++el1;
				++el2;
			}

			// match, continue in left side
			if ( el1==limit ) {
				rightEl = midEl-1;
				rightSize = midSize-1;
			}
		}
	}

	if ( 0!=index )
		*index = (size_t) ( ( leftEl==midEl ? midEl : leftEl ) - (*list).elements );
	return 0;
}


int
List_searchDesc ( List *const list, size_t *const index, const void *const element, const size_t elementSize ) {
	const size_t lastElementIndex = (*list).len-1;
	void **midEl;
	void **leftEl = (*list).elements;
	void **rightEl = (*list).elements + lastElementIndex;
	size_t *midSize;
	size_t *leftSize = (*list).sizes;
	size_t *rightSize = (*list).sizes + lastElementIndex;

	while ( leftEl<=rightEl ) {
		midEl = leftEl + ( (rightEl-leftEl) / 2 );
		midSize = leftSize + ( (rightSize-leftSize) / 2 );
		const size_t midSizeVal = (*midSize);
		const char *el1 = (const char*) *midEl;
		const char *el2 = (const char*) element;

		// compare elements

		if ( midSizeVal==elementSize ) {
			const char *const limit = el1+elementSize;

			while ( el1<limit ) {

				// continue in left side
				if ( *el1<*el2 ) {
					rightEl = midEl-1;
					rightSize = midSize-1;
					break;

				// continue in right side
				} else if ( *el1>*el2 ) {
					leftEl = midEl+1;
					leftSize = midSize+1;
					break;
				}
				++el1;
				++el2;
			}

			// match
			if ( el1==limit ) {
				if ( 0!=index )
					*index = (size_t) (midEl-(*list).elements);
				return 1;
			}

		} else if ( midSizeVal<elementSize ) {
			const char *const limit = el1+midSizeVal;

			while ( el1<limit ) {

				// continue in left side
				if ( *el1<*el2 ) {
					rightEl = midEl-1;
					rightSize = midSize-1;
					break;

				// continue in right side
				} else if ( *el1>*el2 ) {
					leftEl = midEl+1;
					leftSize = midSize+1;
					break;
				}
				++el1;
				++el2;
			}

			// match, continue in left side
			if ( el1==limit ) {
				rightEl = midEl-1;
				rightSize = midSize-1;
			}

		} else {
			const char *const limit = el1+elementSize;

			while ( el1<limit ) {

				// continue in left side
				if ( *el1<*el2 ) {
					rightEl = midEl-1;
					rightSize = midSize-1;
					break;

				// continue in right side
				} else if ( *el1>*el2 ) {
					leftEl = midEl+1;
					leftSize = midSize+1;
					break;
				}
				++el1;
				++el2;
			}

			// match, continue in right side
			if ( el1==limit ) {
				leftEl = midEl+1;
				leftSize = midSize+1;
			}
		}
	}

	if ( 0!=index )
		*index = (size_t) ( ( leftEl==midEl ? midEl : leftEl ) - (*list).elements );
	return 0;
}




























































