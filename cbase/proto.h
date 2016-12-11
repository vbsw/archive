
//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


/**
 * @file proto.h
 * @author Vitali Baumtrok
 * @date 2014-11-20
 * @brief functions for prototyping
 * 
 * The functions here are for fast prototyping. It's not possible to handle errors with this functions.
 */

/**
 * @mainpage Abstract
 * 
 * <i>cbase</i> is a programing library for the language C. It contains simple data structures with functions for general purpose.<br><br>
 * Function tests are included in source code.
 */


#ifndef  FUNC_H
#define  FUNC_H


/**
 * @brief gets data from file
 * 
 * Allocates new memory, copies the content of the file to it and returns a
 * pointer to that memory.
 * 
 * @param fname Name of the file to read from.
 * @return Pointer to the data from file.
 */
extern void *
proto_fget ( const char *const fname );


/**
 * @brief appends data to file
 * 
 * Writes to file the given data. If file does not exist, then it is created.
 * 
 * @param fname Name of the file to write to.
 * @param data Pointer to memory to be written to file.
 * @param size Size in bytes of elements to be written.
 * @param count Number of elements, each one with a size of size bytes.
 */
extern void
proto_fpnd ( const char *const fname, const void *const data, const size_t size, const size_t count );


#endif /* FUNC_H */






















































