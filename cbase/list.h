
//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


/**
 * @file list.h
 * @author Vitali Baumtrok
 * @date 2014-12-08
 * @brief an array list to hold data
 * @details The struct List holds data and its corresponding sizes. The struct
 * is simple enough to access its elements directly.
 * You can use this list in combination with the Buffer. Store
 * data in Buffer and put its pointers in List.
 * After using the List, free the memory of Buffer and
 * List.<br><br>
 * This documentation is just a reminder what the functions do. Please look
 * in to the source code to understand more. Be mindful of the missing bounds
 * checks.
 */


#ifndef  LIST_H
#define  LIST_H

/**
 * @brief array list
 * 
 * This list stores pointers to void, i.e. pointers to unknown data. It also
 * stores the corresponding size of the data in bytes. The size don't need to
 * be the real size of the stored data, because the size is never used in
 * functions to store data, only in search functions. You can access the
 * elements directly, for example void *e = list.elements[0]; to get the
 * first element in the list, or list.len = 0; to reset the List.
 */
typedef struct {
	void **elements;
	size_t *sizes;
	size_t len;
	size_t cap;
} List;


/**
 * @brief initializes the list
 * 
 * Allocates the needed space to store the elements of the list. This serves as
 * an initialization for the list.
 * <br><br>
 * If the list already contains valid space it is not freed.
 * 
 * @param list The list to initialize.
 */
extern void
List_allocate ( List *const list );


/**
 * @brief destroys the list
 * 
 * Deallocates the memory the list is holding to store its elements, i.e.
 * the struct members elements and sizes are lost.
 * The elements themselfs are not deallocated and the members len
 * and cap of the list struct stay untouched.
 * 
 * @param list The list to destroy.
 */
extern void
List_deallocate ( List *const list );


/**
 * @brief appends an element to the list
 * 
 * Appends an element to the list.
 * 
 * @param list The list to which the element is to append.
 * @param element The element to append.
 * @param elementSize The size of the element in bytes.
 */
extern void
List_append ( List *const list, void *const element, const size_t elementSize );


/**
 * @brief inserts an element into the list
 * 
 * Inserts an element into the list.<br><br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param list The list in which the element is to insert.
 * @param index The index where to insert the element.
 * @param element The element to insert.
 * @param elementSize The size of the element in bytes.
 */
extern void
List_insert ( List *const list, const size_t index, void *const element, const size_t elementSize );


/**
 * @brief removes an element from the list
 * 
 * Removes an element from the list and its corresponding size. If the index
 * is bigger then the length of the list the last element is removed.<br><br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param list The list from which the element is to remove.
 * @param index The index of the element to remove.
 */
extern void
List_remove ( List *const list, const size_t index );


/**
 * @brief copy an element from the list
 * 
 * Makes a copy of an element. The size of the copy is given as argument.
 * 
 * @param list The list from which the element is to copy.
 * @param index The index of the element to copy.
 * @param elementSize The size of the copy, i.e. the element, in bytes.
 * @return Pointer to the copy of the element.
 */
extern void *
List_elementClone ( List *const list, const size_t index, const size_t elementSize );


/**
 * @brief searches for an element
 * 
 * Searches for an element in the list. Returns the index of the element if
 * found, otherwise the index is equal to the length of the list. This is a
 * linear search.
 * 
 * @param list The list where to search for the element.
 * @param element The element to search for.
 * @param elementSize The size of the element in bytes.
 * @return The index of the matching element, otherwise the length of the list.
 */
extern size_t
List_search ( List *const list, const void *const element, const size_t elementSize );


/**
 * @brief searches for an element in a list having its elements in ascending order
 * 
 * Searches for an element in the list having its element in ascanding order.
 * Ascending order is ["a", "ax", "b", "bx"]. This is a binary search.
 * Comparison happens byte wise.
 * 
 * @param list The list where to search for the element.
 * @param index Output parameter! The index where the element would fit. If
 * the parameter is zero it's not used.
 * @param element The element to search for.
 * @param elementSize The size of the element in bytes.
 * @return 1 if match, otherwise 0.
 */
extern int
List_searchAsc ( List *const list, size_t *const index, const void *const element, const size_t elementSize );


/**
 * @brief searches for an element in a list having its elements in descending order
 * 
 * Searches for an element in the list having its element in descending order.
 * Descending order is ["bx", "b", "ax", "a"]. This is a binary search.
 * Comparison happens byte wise.
 * 
 * @param list The list where to search for the element.
 * @param index Output parameter! The index where the element would fit. If
 * the parameter is zero it's not used.
 * @param element The element to search for.
 * @param elementSize The size of the element in bytes.
 * @return 1 if match, otherwise 0.
 */
extern int
List_searchDesc ( List *const list, size_t *const index, const void *const element, const size_t elementSize );


#endif /* LIST_H */




















