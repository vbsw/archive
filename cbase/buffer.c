
//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


#include <stdlib.h>
#include <string.h>


typedef struct {
	void **chunks;
	size_t *lengths;
	size_t *capacities;
	size_t len;
	size_t cap;
	size_t preferredChunkCapacity;
} Buffer;


// alignment in bytes
#define BUFFER_ALIGNMENT 8

// capacity in number of elements
#define BUFFER_INITIAL_CAP 7


void
Buffer_allocate ( Buffer *const buffer, const size_t preferredChunkCapacity ) {
	(*buffer).chunks = (void**) malloc ( BUFFER_INITIAL_CAP * sizeof(void*) );
	(*buffer).lengths = (size_t*) malloc ( BUFFER_INITIAL_CAP * sizeof(size_t) );
	(*buffer).capacities = (size_t*) malloc ( BUFFER_INITIAL_CAP * sizeof(size_t) );
	(*buffer).len = 0;
	(*buffer).cap = BUFFER_INITIAL_CAP;
	(*buffer).preferredChunkCapacity = preferredChunkCapacity;
}


void
Buffer_deallocate ( Buffer *const buffer ) {
	void **chunks = (*buffer).chunks;
	void **const limit = chunks + (*buffer).len;

	while ( chunks<limit ) {
		free ( (*chunks) );
		++chunks;
	}

	free ( (*buffer).chunks );
	free ( (*buffer).lengths );
	free ( (*buffer).capacities );
}


void
Buffer_reset ( Buffer *const buffer ) {
	size_t * lengths = (*buffer).lengths;
	const size_t *const limit = lengths + (*buffer).len;

	for ( ; lengths<limit; lengths+=1 )
		*lengths = 0;
	(*buffer).len = 0;
}


void *
Buffer_get ( Buffer *const buffer, const size_t dataSize ) {
	const size_t modulo = dataSize%BUFFER_ALIGNMENT;
	const size_t alignedDataSize = dataSize + ( modulo==0 ? 0 : BUFFER_ALIGNMENT - modulo );
	void *result = 0;
	void **chunks = (*buffer).chunks;
	size_t *lengths = (*buffer).lengths;
	size_t *capacities = (*buffer).capacities;
	void **const limit = chunks + (*buffer).len;

	// search for free memory in chunks
	while ( chunks<limit ) {
		const size_t length = *lengths;
		const size_t capacity = *capacities;
		
		if ( capacity-length >= dataSize ) {
			result = (*chunks)+length;
			(*lengths) = length + alignedDataSize;
			break;
		}
		
		++chunks;
		++lengths;
		++capacities;
	}

	// allocate new chunk
	if ( 0==result ) {
		const size_t len = (*buffer).len;
		const size_t cap = (*buffer).cap;

		// ensure capacity
		if ( len==cap ) {
			const size_t oldChunksCapacitySize = cap * sizeof(void*);
			const size_t oldLenCapCapacitySize = cap * sizeof(size_t);
			const size_t newChunksCapacitySize = oldChunksCapacitySize*2;
			const size_t newLenCapCapacitySize = oldLenCapCapacitySize*2;

			void **newChunks = (void**) malloc ( newChunksCapacitySize );
			size_t *newLengths = (size_t*) malloc ( newLenCapCapacitySize );
			size_t *newCapacities = (size_t*) malloc ( newLenCapCapacitySize );

			memcpy ( newChunks, (*buffer).chunks, oldChunksCapacitySize );
			memcpy ( newLengths, (*buffer).lengths, oldLenCapCapacitySize );
			memcpy ( newCapacities, (*buffer).capacities, oldLenCapCapacitySize );

			free ( (*buffer).chunks );
			free ( (*buffer).lengths );
			free ( (*buffer).capacities );

			(*buffer).chunks = newChunks;
			(*buffer).lengths = newLengths;
			(*buffer).capacities = newCapacities;
			(*buffer).cap = cap*2;
		}

		// allocate preferred chunk size
		if ( (*buffer).preferredChunkCapacity>=dataSize ) {
			result = malloc ( (*buffer).preferredChunkCapacity );
			(*buffer).capacities[len] = (*buffer).preferredChunkCapacity;
		// allocate given data size
		} else {
			result = malloc ( dataSize );
			(*buffer).capacities[len] = dataSize;
		}
		(*buffer).chunks[len] = result;
		(*buffer).lengths[len] = alignedDataSize;
		++(*buffer).len;
	}

	return result;
}


void *
Buffer_put ( Buffer *const buffer, const void *const data, const size_t dataSize ) {
	void *result = Buffer_get ( buffer, dataSize );
	memcpy ( result, data, dataSize );
	return result;
}


void *
Buffer_putc ( Buffer *const buffer, const void *const data, const size_t dataSize ) {
	void *result;

	if ( 0==data || 0==dataSize ) {
		result = Buffer_get ( buffer, 1 );
		((char*)result)[0] = 0;

	} else {
		result = Buffer_get ( buffer, dataSize+1 );
		memcpy ( result, data, dataSize );
		((char*)result)[dataSize] = 0;
	}
	return result;
}




























































