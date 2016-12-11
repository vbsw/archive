
//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


/**
 * @file buffer.h
 * @author Vitali Baumtrok
 * @date 2014-12-08
 * @brief memory buffer
 * @details The Buffer reserves memory space and expands it as needed in
 * preferred chunk sizes. If the reserved space isn't enough then a new chunk
 * of space is reserved. If the data to store exceeds the
 * preferredChunkCapacity then a new chunk is allocated with the
 * exactly required size.
 */


#ifndef  BUFFER_H
#define  BUFFER_H


/**
 * @brief memory buffer
 * 
 * The Buffer holds memory. It acts also as an allocator. Once allocated
 * the memory can not be freed partly, but only as a whole. Use the functions
 * to interact with the Buffer.
 */
typedef struct {
	void **chunks;
	size_t *lengths;
	size_t *capacities;
	size_t len;
	size_t cap;
	size_t preferredChunkCapacity;
} Buffer;


/**
 * @brief initializes the buffer
 * 
 * Allocates the needed space to store the memory of the buffer. This serves as
 * an initialization for the buffer. 
 * <br><br>
 * If the buffer already contains valid space it is not freed.
 * 
 * @param buffer The buffer to initialize.
 * @param preferredChunkCapacity The size of buffer chunks in bytes. If the
 * data to store exceeds the preferredChunkCapacity then a new
 * chunk is allocated with the exactly required size.
 */
extern void
Buffer_allocate ( Buffer *const buffer, const size_t preferredChunkCapacity );

/**
 * @brief destroys the buffer
 * 
 * Deallocates all memory the buffer is holding. To use the buffer again you
 * have to call Buffer_allocate.
 * 
 * @param buffer The buffer to destroy.
 */
extern void
Buffer_deallocate ( Buffer *const buffer );


/**
 * @brief sets the content to zero (but keeps the buffer)
 * 
 * The content of the buffer is set to zero (i.e. buffer is empty). Memory
 * is not deallocated.
 * 
 * @param buffer The buffer.
 */
extern void
Buffer_reset ( Buffer *const buffer );


/**
 * @brief allocates memory
 * 
 * Allocates memory of the given size and returns it.
 * 
 * @param buffer The buffer.
 * @param dataSize The size of memory to return in bytes.
 * @return Pointer to the memory of size dataSize.
 */
extern void *
Buffer_get ( Buffer *const buffer, const size_t dataSize );


/**
 * @brief copies data to buffer
 * 
 * Copies data to buffer.
 * 
 * @param buffer The buffer.
 * @param data Pointer to data to copy to buffer. Must not be zero.
 * @param dataSize Size of the data to copy to buffer in bytes. Must not be zero.
 * @return Pointer to the copy.
 */
extern void *
Buffer_put ( Buffer *const buffer, const void *const data, const size_t dataSize );


/**
 * @brief copies data to buffer
 * 
 * Copies data to buffer. This is the same as Buffer_put with
 * the difference that it appends a zero to the data. This is usefull, if
 * you want to store a zero terminated string. cs means c string.
 * 
 * @param buffer The buffer.
 * @param data Pointer to data to copy to buffer. It can be zero.
 * @param dataSize Size of the data to copy to buffer in bytes. It can be zero.
 * @return Pointer to the copy.
 */
extern void *
Buffer_putc ( Buffer *const buffer, const void *const data, const size_t dataSize );


#endif /* BUFFER_H */















