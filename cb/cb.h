#ifndef  CB_H
#define  CB_H

#include <stdlib.h>

extern void *
cb_allocateBuffer ( const size_t in_bufferSize );

extern void
cb_freeBuffer ( void *const inout_buffer );

extern void *
cb_getBufferMemory ( void *const inout_buffer, const size_t in_memorySize );

extern int
cb_stringBinarySearchAsc ( size_t *const out_index, const char *const *const in_stringArray, const size_t in_length, const char *const in_string );

extern int
cb_stringBinarySearchDesc ( size_t *const out_index, const char *const *const in_stringArray, const size_t in_length, const char *const in_string );

#endif /* CB_H */