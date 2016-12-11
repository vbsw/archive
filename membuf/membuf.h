#ifndef  MEMBUF_H
#define  MEMBUF_H

#include <stdlib.h>


extern void *
MemBuf_allocate
(
	const size_t in_chunkSize
);


extern void
MemBuf_free
(
	void *inout_buffer
);


extern void *
MemBuf_getMemory
(
	void *const inout_buffer,
	const size_t in_requestedMemorySize
);


extern size_t
MemBuf_getTotalSize
(
	const void *in_buffer
);


extern size_t
MemBuf_getUsedSize
(
	const void *in_buffer
);

#endif /* MEMBUF_H */
