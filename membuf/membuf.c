
//          Copyright Vitali Baumtrok 2015.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


/**
 * @file membuf.c
 * @author Vitali Baumtrok
 * @date 2015-05-06
 * @brief memory buffer
 * 
 * Functions for a memory buffer (allocator), that allocates only.
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
internal_MemBuf_computeRoundedMemorySize
(
	const size_t in_size
)
{
	const size_t rest = in_size % sizeof ( size_t );

	if ( rest )
		return in_size - rest + sizeof ( size_t );
	else
		return in_size;
}

/**
 * @brief Allocate buffer.
 * 
 * Allocate a buffer with a given chunk size. If not enough memory
 * could be allocated 0 is returned, otherwise 1.
 * 
 * @param in_chunkSize The minimum chunk size in bytes.
 * @return The allocated buffer or 0, if couldn't allocate memory.
 */
void *
MemBuf_allocate
(
	const size_t in_chunkSize
)
{
	void **const buffer = (void **const) malloc ( sizeof(void**) * 4 );

	if ( buffer )
	{
		const size_t actualChunkSize = internal_MemBuf_computeRoundedMemorySize ( in_chunkSize );
		void *const chunk = malloc ( actualChunkSize );

		if ( chunk )
		{
			buffer[0] = chunk;
			buffer[1] = chunk;
			buffer[2] = (void*) ( (size_t) chunk + actualChunkSize );
			buffer[3] = 0;

			return (void*) buffer;
		}
		else
		{
			free ( buffer );
			return 0;
		}
	}
	else
	{
		return 0;
	}
}

/**
 * @brief Deallocate the buffer.
 * 
 * Deallocate the buffer.
 * 
 * @param buffer Buffer to deallocate.
 */
void
MemBuf_free
(
	void *inout_buffer
)
{
	while ( 1 ) {
		void **const bufferArray = (void **const) inout_buffer;
		void *const next = (void *const) (bufferArray[3]);

		free ( bufferArray[0] );
		free ( inout_buffer );

		if ( next )
			inout_buffer = next;
		else
			break;
	}
}

/**
 * @brief Get available memory from the buffer.
 * 
 * Searches for available memory in buffer and returns it. If no available
 * memory is found new memory is allocated. If no memory could be allocated
 * 0 is returned.
 * 
 * @param inout_buffer The buffer.
 * @param in_requestedMemorySize The size of requested memory in bytes.
 */
void *
MemBuf_getMemory
(
	void *inout_buffer,
	const size_t in_requestedMemorySize
)
{
	void **bufferArray = (void **) inout_buffer;
	void *arrayEnd = bufferArray[2];
	void *arrayUsed = bufferArray[1];
	void *next = bufferArray[3];

	const size_t bufferSize = (size_t) arrayEnd - (size_t) bufferArray[0];
	const size_t actualRequestedMemorySize = internal_MemBuf_computeRoundedMemorySize ( in_requestedMemorySize );

	// requested memory size matches the buffer size
	if ( in_requestedMemorySize <= bufferSize )
	{
		while (1)
		{
			const size_t availableChunkMemorySize = (size_t) arrayEnd - (size_t) arrayUsed;

			// available memory found
			if ( in_requestedMemorySize <= availableChunkMemorySize )
			{
				bufferArray[1] = (void *) ( (size_t) arrayUsed + actualRequestedMemorySize );

				return arrayUsed;
			}
			// more memory to search
			else if ( next )
			{
				bufferArray = (void **) next;
				arrayUsed = bufferArray[1];
				arrayEnd = bufferArray[2];
				next = bufferArray[3];
			}
			// new memory needed
			else
			{
				next = MemBuf_allocate ( bufferSize );

				if ( next )
				{
					bufferArray[3] = next;
					bufferArray = (void **) next;
					bufferArray[1] = (void *) ( (size_t) bufferArray[1] + actualRequestedMemorySize );

					return bufferArray[0];
				}
				// return; no memory
				else
				{
					return 0;
				}
			}
		}
	}
	// new buffer chunk must be allocated
	else
	{
		while ( 1 )
		{
			// search for last chunk
			if ( next )
			{
				bufferArray = (void **) next;
				next = bufferArray[3];
			}
			// allocate new chunk
			else
			{
				next = MemBuf_allocate ( actualRequestedMemorySize );

				if ( next )
				{
					bufferArray[3] = next;
					bufferArray = (void **) next;
					bufferArray[1] = (void *) ( (size_t) bufferArray[1] + actualRequestedMemorySize );

					return bufferArray[0];
				}
				// return; no memory
				else
				{
					return 0;
				}
			}
		}
	}
}

size_t
MemBuf_getTotalSize
(
	const void *in_buffer
)
{
	size_t totalSize = 0;

	do
	{
		const void *const *const bufferArray = (const void *const *const) in_buffer;
		const void *const arrayEnd = bufferArray[2];

		totalSize += (size_t) arrayEnd - (size_t) bufferArray[0];
		in_buffer = bufferArray[3];
	}
	while ( in_buffer );

	return totalSize;
}

size_t
MemBuf_getUsedSize
(
	const void *in_buffer
)
{
	size_t usedSize = 0;

	do
	{
		const void *const *const bufferArray = (const void *const *const) in_buffer;
		const void *const arrayUsed = bufferArray[1];
		const void *const arrayEnd = bufferArray[2];

		if ( arrayEnd < arrayUsed )
			usedSize += (size_t) arrayEnd - (size_t) bufferArray[0];
		else
			usedSize += (size_t) arrayUsed - (size_t) bufferArray[0];
		in_buffer = bufferArray[3];
	}
	while ( in_buffer );

	return usedSize;
}
