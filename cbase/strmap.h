
//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


/**
 * @file strmap.h
 * @author Vitali Baumtrok
 * @date 2014-12-08
 * @brief string map
 * @details String map is a container to hold string pairs. Strings are not
 * copied. Use it to store for example properties.<br>
 * <br>
 * The header strmap.h includes stdio.h, buffer.h and list.h.
 */


#ifndef  STRMAP_H
#define  STRMAP_H


#include <stdio.h>

#include "buffer.h"
#include "list.h"


/**
 * @brief string map
 * 
 * String map is a container to hold string pairs. Use the functions
 * to interact with StrMap.
 */
typedef struct {
	List keys;
	List values;
} StrMap;


/**
 * @brief initializes StrMap
 * 
 * Allocates the needed space to store the memory. This serves as an
 * initialization for StrMap. 
 * <br><br>
 * If StrMap already contains valid space it is not freed.
 * 
 * @param strMap StrMap to initialize.
 */
static inline void
StrMap_allocate ( StrMap *const strMap ) {
	List_allocate ( & (*strMap).keys );
	List_allocate ( & (*strMap).values );
}


/**
 * @brief destroys StrMap
 * 
 * Deallocates all memory StrMap is holding. To use StrMap again you
 * have to call StrMap_allocate.
 * 
 * @param strMap StrMap to destroy.
 */
static inline void
StrMap_deallocate ( StrMap *const strMap ) {
	List_deallocate ( & (*strMap).keys );
	List_deallocate ( & (*strMap).values );
}


/**
 * @brief sets the content to zero (but keeps the buffer)
 * 
 * The content of StrMap is set to zero. Memory is not deallocated.<br>
 * <b>CAUTION</b>: Memory of StrMap will be overwritten.
 * 
 * @param strMap StrMap to reset.
 */
static inline void
StrMap_reset ( StrMap *const strMap ) {
	(*strMap).keys.len = 0;
	(*strMap).values.len = 0;
}


/**
 * @brief appends a key and its value to StrMap
 * 
 * Appends a key and its value to StrMap. It's not checked whether the key is
 * already in StrMap.
 * 
 * @param strMap StrMap to add data to. StrMap must be already initialized.
 * @param key The key characters (not terminated by 0). Must not be zero.
 * @param keySize The size of key in bytes. Must not be zero.
 * @param value Pointer to the value characters (not terminated by 0).
 * @param valueSize The size of value in bytes.
 */
static inline void
StrMap_append ( StrMap *const strMap, const char *const key, const size_t keySize, const char *const value, const size_t valueSize ) {
	List_append ( &(*strMap).keys, (void*) key, keySize );
	List_append ( &(*strMap).values, (void*) value, valueSize );
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
 * <b>CAUTION</b>: Data is not copied, but pointed to.
 * 
 * @param strMap StrMap to add data to. StrMap must be already initialized.
 * @param data Data to parse and to append.
 * @param strSize The size of data in bytes.
 */
extern void
StrMap_appendParse ( StrMap *const strMap, const char * data, const size_t strSize );


/**
 * @brief writes string pairs to file
 * 
 * Writes strings to file. The file will be like:<br>
 * key1=value1<br>
 * key2=value2
 * 
 * @param strMap StrMap to write to file. StrMap must contain valid data.
 * @param file An open file to write to.
 */
extern void
StrMap_writeFile ( StrMap *const strMap, FILE *const file );


/**
 * @brief copies the content of StrMap to another one
 * 
 * Copies the content of StrMap to another one. dest must be uninitialized,
 * because new memory is allocated and previous is not released.
 * 
 * @param srcStrMap StrMap to make copy of.
 * @param destStrMap StrMap to write to.
 */
extern void
StrMap_clone ( StrMap *const srcStrMap, StrMap *const destStrMap );


/**
 * @brief returns a pointer to the key
 * 
 * Returns a pointer to the key at the given index.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param strMap StrMap where to return the key from.
 * @param index The index of the key.
 * @return Pointer to the key.
 */
static inline char *
StrMap_getKey ( StrMap *const strMap, const size_t index ) {
	return (char*) (*strMap).keys.elements[index];
}


/**
 * @brief returns a pointer to the value
 * 
 * Returns a pointer to the value at the given index.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param strMap StrMap where to return the value from.
 * @param index The index of the value.
 * @return Pointer to the value.
 */
static inline char *
StrMap_getValue ( StrMap *const strMap, const size_t index ) {
	return (char*) (*strMap).values.elements[index];
}


/**
 * @brief replaces a key
 * 
 * Replaces the key at index. The key is copied to the map.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param strMap StrMap where to replace the key.
 * @param index The index where to replace the key.
 * @param key Pointer to a non zero terminated string to put in StrMap.
 * @param keySize The size of key in bytes.
 */
static inline void
StrMap_setKey ( StrMap *const strMap, const size_t index, char *const key, const size_t keySize ) {
	(*strMap).keys.elements[index] = (void*) key;
	(*strMap).keys.sizes[index] = keySize;
}


/**
 * @brief replaces a value
 * 
 * Replaces the value at index. The value is copied to the map.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param strMap StrMap where to replace the value.
 * @param index The index where to replace the value.
 * @param value Pointer to a non zero terminated string to put in StrMap.
 * @param valueSize The size of value in bytes.
 */
static inline void
StrMap_setValue ( StrMap *const strMap, const size_t index, char *const value, const size_t valueSize ) {
	(*strMap).values.elements[index] = (void*) value;
	(*strMap).values.sizes[index] = valueSize;
}


/**
 * @brief inserts a string pair
 * 
 * Inserts a string pair.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param strMap StrMap where to insert the string pair.
 * @param index The index where to insert the string pair.
 * @param key Pointer to a non zero terminated string that is inserted as key.
 * @param keySize The size of key in bytes.
 * @param value Pointer to a non zero terminated string that is inserted as
 * value.
 * @param valueSize The size of value in bytes.
 */
static inline void
StrMap_insert ( StrMap *const strMap, const size_t index, char *const key, const size_t keySize, char *const value, const size_t valueSize ) {
	List_insert ( & (*strMap).keys, index, (void*) key, keySize );
	List_insert ( & (*strMap).values, index, (void*) value, valueSize );
}


/**
 * @brief searches for the index of a key
 * 
 * Searches for the index of a key. This is a linear search.
 * 
 * @param strMap StrMap where to search for key.
 * @param key Pointer to a non zero terminated string to search for as the key.
 * The pointer must not be zero, if the keySize is not zero.
 * @param keySize The size of key in bytes. It may be zero.
 * @return The index of the key if found
 * ( StrMap_search(..) < strMap.keys.len ), otherwise the length of keys.
 */
extern size_t
StrMap_search ( StrMap *const strMap, const char *const key, const size_t keySize );


/**
 * @brief searches for a key
 * 
 * Searches for a key in the map having its keys in ascending order.
 * Ascending order is ["a", "ax", "b", "bx"]. This is a binary search.
 * 
 * @param strMap StrMap where to search for key.
 * @param index Output parameter! The index where the pair would fit. If
 * the parameter is zero it's not used.
 * @param key Pointer to a non zero terminated string to search for as the key.
 * The pointer must not be zero, if the keySize is not zero.
 * @param keySize The size of key in bytes. It may be zero.
 * @return 1 if match, otherwise 0.
 */
extern size_t
StrMap_searchAsc ( StrMap *const strMap, size_t *const index, const char *const key, const size_t keySize );


/**
 * @brief searches for a key
 * 
 * Searches for a key in the map having its keys in descending order.
 * Descending order is ["bx", "b", "ax", "a"]. This is a binary search.
 * 
 * @param strMap StrMap where to search for key.
 * @param index Output parameter! The index where the pair would fit. If
 * the parameter is zero it's not used.
 * @param key Pointer to a non zero terminated string to search for as the key.
 * The pointer must not be zero, if the keySize is not zero.
 * @param keySize The size of key in bytes. It may be zero.
 * @return 1 if match, otherwise 0.
 */
extern size_t
StrMap_searchDesc ( StrMap *const strMap, size_t *const index, const char *const key, const size_t keySize );


#endif /* STRMAP_H */















