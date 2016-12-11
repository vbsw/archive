
//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


/**
 * @file strmapbuf.h
 * @author Vitali Baumtrok
 * @date 2014-12-08
 * @brief string map buffer
 * @details String map buffer is a container with a buffer to hold string
 * pairs. Strings are copied when put in map. Use it to store for example
 * properties.<br>
 * <br>
 * The header strmapbuf.h includes strmap.h (stdio.h, buffer.h and list.h).
 */


#ifndef  STRMAPBUF_H
#define  STRMAPBUF_H


#include <stdio.h>

#include "strmap.h"


/**
 * @brief string map buffer
 * 
 * String map buffer is a container to hold string pairs. Use the functions
 * of StrMapBuf.
 */
typedef struct {
	Buffer buffer;
	StrMap strMap;
} StrMapBuf;


/**
 * @brief initializes StrMapBuf
 * 
 * Allocates the needed space to store the memory. This serves as an
 * initialization for StrMapBuf. 
 * <br><br>
 * If StrMapBuf already contains valid space it is not freed.
 * 
 * @param strMapBuf StrMapBuf to initialize.
 */
static inline void
StrMapBuf_allocate ( StrMapBuf *const strMapBuf ) {
	Buffer_allocate ( & (*strMapBuf).buffer, 2048-8 );
	StrMap_allocate ( & (*strMapBuf).strMap );
}


/**
 * @brief destroys StrMapBuf
 * 
 * Deallocates all memory StrMapBuf is holding. To use StrMapBuf again you
 * have to call StrMap_allocate.
 * 
 * @param strMapBuf StrMapBuf to destroy.
 */
static inline void
StrMapBuf_deallocate ( StrMapBuf *const strMapBuf ) {
	Buffer_deallocate ( & (*strMapBuf).buffer );
	StrMap_deallocate ( & (*strMapBuf).strMap );
}


/**
 * @brief sets the content to zero (buffer is not deallocated)
 * 
 * The content of StrMapBuf is set to zero. Memory is not deallocated.<br>
 * <b>CAUTION</b>: Memory of StrMapBuf will be overwritten.
 * 
 * @param strMap StrMapBuf to reset.
 */
static inline void
StrMapBuf_reset ( StrMapBuf *const strMapBuf ) {
	Buffer_reset ( & (*strMapBuf).buffer );
	StrMap_reset ( & (*strMapBuf).strMap );
}


/**
 * @brief appends a key and its value to StrMapBuf
 * 
 * Appends a key and its value to StrMapBuf. It's not checked whether the key
 * is already in StrMapBuf. Values are copied to StrMapBuf.
 * 
 * @param strMapBuf StrMapBuf to add data to. StrMapBuf must be already
 * initialized.
 * @param key The key characters (not terminated by 0). Pointer must not be
 * zero.
 * @param keySize The size of key in bytes. Must not be zero.
 * @param value The value characters (not terminated by 0). Pointer may be
 * zero.
 * @param valueSize The size of value in bytes. It may be zero.
 */
static inline void
StrMapBuf_append ( StrMapBuf *const strMapBuf, const char *const key, const size_t keySize, const char *const value, const size_t valueSize ) {
	StrMap_append ( & (*strMapBuf).strMap,
					(char*) Buffer_putc ( & (*strMapBuf).buffer, (const void*) key, keySize ),
					keySize,
					(char*) Buffer_putc ( & (*strMapBuf).buffer, (const void*) value, valueSize ),
					valueSize );
}


/**
 * @brief parses data and appends its content
 * 
 * data should not be zero terminated. The format of data must be like:<br>
 * DATA = [LINE{NEWLINE LINE}]<br>
 * LINE = [WHITESPACE {NOTNEWLINE}] NEWLINE | KEY[=[VALUE]]<br>
 * WHITESPACE = BLANK | TAB<br>
 * <br>
 * Example:<br>
 * key1=value1<br>
 * key2=value2<br>
 * <br>
 * Values are copied to StrMapBuf.
 * 
 * @param strMap StrMapBuf to add data to. StrMapBuf must be already
 * initialized.
 * @param data Data to parse and to append.
 * @param strSize The size of data in bytes.
 */
extern void
StrMapBuf_appendParse ( StrMapBuf *const strMapBuf, const char * data, const size_t dataSize );


/**
 * @brief reads string pairs from file
 * 
 * Reads strings from file. The file format must be like:<br>
 * FILE = [LINE{NEWLINE LINE}]<br>
 * LINE = [WHITESPACE {NOTNEWLINE}] NEWLINE | KEY[=[VALUE]]<br>
 * WHITESPACE = BLANK | TAB<br>
 * <br>
 * Example:<br>
 * key1=value1<br>
 * key2=value2<br>
 * <br>
 * Values are copied to StrMapBuf.
 * 
 * @param strMapBuf StrMapBuf to add data to. StrMapBuf must be already
 * initialized.
 * @param file An open file to read from.
 */
extern void
StrMapBuf_readFile ( StrMapBuf *const strMapBuf, FILE *const file );


/**
 * @brief writes string pairs to file
 * 
 * Writes strings to file. The file will be like:<br>
 * key1=value1<br>
 * key2=value2
 * 
 * @param strMapBuf StrMapBuf to write to file. StrMapBuf must contain valid
 * data.
 * @param file An open file to write to.
 */
static inline void
StrMapBuf_writeFile ( StrMapBuf *const strMapBuf, FILE *const file ) {
	StrMap_writeFile ( & (*strMapBuf).strMap, file );
}


/**
 * @brief copies the content of StrMapBuf to another one
 * 
 * Copies the content of StrMapBuf to another one. dest must be uninitialized,
 * because new memory is allocated and previous is not released.
 * 
 * @param srcStrMapBuf StrMapBuf to make copy of.
 * @param destStrMapBuf StrMapBuf to write to.
 */
extern void
StrMapBuf_clone ( StrMapBuf *const srcStrMapBuf, StrMapBuf *const destStrMapBuf );


/**
 * @brief returns a pointer to the key
 * 
 * Returns a pointer to the key at the given index.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param strMapBuf StrMapBuf where to return the key from.
 * @param index The index of the key.
 * @return Pointer to the key.
 */
static inline char *
StrMapBuf_getKey ( StrMapBuf *const strMapBuf, const size_t index ) {
	return StrMap_getKey ( & (*strMapBuf).strMap, index );
}


/**
 * @brief returns a pointer to the value
 * 
 * Returns a pointer to the value at the given index.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param strMapBuf StrMapBuf where to return the value from.
 * @param index The index of the value.
 * @return Pointer to the value.
 */
static inline char *
StrMapBuf_getValue ( StrMapBuf *const strMapBuf, const size_t index ) {
	return StrMap_getValue ( & (*strMapBuf).strMap, index );
}


/**
 * @brief replaces a key
 * 
 * Replaces the key at index. The key is copied to StrMapBuf.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param strMapBuf StrMapBuf where to replace the key.
 * @param index The index where to replace the key.
 * @param key Pointer to a non zero terminated string to put in StrMapBuf.
 * @param keySize The size of key in bytes.
 */
static inline void
StrMapBuf_setKey ( StrMapBuf *const strMapBuf, const size_t index, const char *const key, const size_t keySize ) {
	StrMap_setKey ( & (*strMapBuf).strMap,
					index,
					(char*) Buffer_putc ( & (*strMapBuf).buffer, key, keySize ),
					keySize );
}


/**
 * @brief replaces a value
 * 
 * Replaces the value at index. The value is copied to StrMapBuf.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param strMapBuf StrMapBuf where to replace the value.
 * @param index The index where to replace the value.
 * @param value Pointer to a non zero terminated string to put in StrMapBuf.
 * @param valueSize The size of value in bytes.
 */
static inline void
StrMapBuf_setValue ( StrMapBuf *const strMapBuf, const size_t index, const char *const value, const size_t valueSize ) {
	StrMap_setValue ( & (*strMapBuf).strMap,
					  index,
					  (char*) Buffer_putc ( & (*strMapBuf).buffer, value, valueSize ),
					  valueSize );
}


/**
 * @brief inserts a string pair
 * 
 * Inserts a string pair. The key and value are copied to StrMapBuf.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param strMapBuf StrMapBuf where to insert the string pair.
 * @param index The index where to insert the string pair.
 * @param key Pointer to a non zero terminated string that is inserted as key.
 * Pointer must not be zero.
 * @param keySize The size of key in bytes. Must not be zero.
 * @param value Pointer to a non zero terminated string that is inserted as
 * value. Pointer may be zero.
 * @param valueSize The size of value in bytes. It may be zero.
 */
static inline void
StrMapBuf_insert ( StrMapBuf *const strMapBuf, const size_t index, char *const key, const size_t keySize, char *const value, const size_t valueSize ) {
	StrMap_insert ( & (*strMapBuf).strMap,
					index,
					(char*) Buffer_putc ( & (*strMapBuf).buffer, (const void*) key, keySize ),
					keySize,
					(char*) Buffer_putc ( & (*strMapBuf).buffer, (const void*) value, valueSize ),
					valueSize );
}


/**
 * @brief searches for the index of a key
 * 
 * Searches for the index of a key. This is a linear search.
 * 
 * @param strMapBuf StrMapBuf where to search for key.
 * @param key Pointer to a non zero terminated string to search for as the key.
 * The pointer must not be zero, if the keySize is not zero.
 * @param keySize The size of key in bytes.
 * @return The index of the key if found
 * ( StrMapBuf_search(..) < strMapBuf.strMap.keys.len ), otherwise the length
 * of keys in strMap.
 */
static inline size_t
StrMapBuf_search ( StrMapBuf *const strMapBuf, const char *const key, const size_t keySize ) {
	return StrMap_search ( & (*strMapBuf).strMap, key, keySize );
}


/**
 * @brief searches for a key
 * 
 * Searches for a key in the map having its keys in ascending order.
 * Ascending order is ["a", "ax", "b", "bx"]. This is a binary search.
 * 
 * @param strMapBuf StrMapBuf where to search for key.
 * @param index Output parameter! The index where the pair would fit. If
 * the parameter, i.e. the pointer, is zero it's not used.
 * @param key Pointer to a non zero terminated string to search for as the key.
 * The pointer must not be zero, if the keySize is not zero.
 * @param keySize The size of key in bytes.
 * @return 1 if match, otherwise 0.
 */
static inline size_t
StrMapBuf_searchAsc ( StrMapBuf *const strMapBuf, size_t *const index, const char *const key, const size_t keySize ) {
	return StrMap_searchAsc ( & (*strMapBuf).strMap, index, key, keySize );
}


/**
 * @brief searches for a key
 * 
 * Searches for a key in the map having its keys in descending order.
 * Descending order is ["bx", "b", "ax", "a"]. This is a binary search.
 * 
 * @param strMapBuf StrMapBuf where to search for key.
 * @param index Output parameter! The index where the pair would fit. If
 * the parameter, i.e. the pointer, is zero it's not used.
 * @param key Pointer to a non zero terminated string to search for as the key.
 * The pointer must not be zero, if the keySize is not zero.
 * @param keySize The size of key in bytes.
 * @return 1 if match, otherwise 0.
 */
static inline size_t
StrMapBuf_searchDesc ( StrMapBuf *const strMapBuf, size_t *const index, const char *const key, const size_t keySize ) {
	return StrMap_searchDesc ( & (*strMapBuf).strMap, index, key, keySize );
}


#endif /* STRMAPBUF_H */















