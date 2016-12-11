
//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


/**
 * @file stack.h
 * @author Vitali Baumtrok
 * @date 2014-12-27
 * @brief stacks for various types
 * @details These are simple stack structures for basic data types: char,
 * size_t and pointer.
 */


#ifndef  STACK_H
#define  STACK_H


/**
 * @brief stack for char
 * 
 * Stack for char. You can access the struct members directly or use the
 * functions. For example to check whether stack is empty use (0==stack.len)
 * or StackC_isEmpty(&stack);.
 */
typedef struct {
	char *values;
	size_t len;
	size_t cap;
} StackC;


/**
 * @brief stack for size_t
 * 
 * Stack for size_t. You can access the struct members directly or use the
 * functions. For example to check whether stack is empty use (0==stack.len)
 * or StackC_isEmpty(&stack);.
 */
typedef struct {
	size_t *values;
	size_t len;
	size_t cap;
} StackS;


/**
 * @brief stack for pointers
 * 
 * Stack for pointers. You can access the struct members directly or use the
 * functions. For example to check whether stack is empty use (0==stack.len)
 * or StackC_isEmpty(&stack);.
 */
typedef struct {
	void **values;
	size_t len;
	size_t cap;
} StackP;


/**
 * @brief initializes the char stack
 * 
 * Allocates the needed space to store the memory of the stack. This serves as
 * an initialization for the stack. The initial capacity of the stack is
 * (256-8) / sizeof(char).
 * <br><br>
 * If the stack already contains valid space it is not freed.
 * 
 * @param stack The stack to initialize.
 */
extern void
StackC_allocate ( StackC *const stack );

/**
 * @brief destroys the char stack
 * 
 * Deallocates all memory the stack is holding. To use the stack again you
 * have to call StackC_allocate.
 * 
 * @param stack The stack to destroy.
 */
extern void
StackC_deallocate ( StackC *const stack );


/**
 * @brief sets the content to zero (but keeps the stack memory)
 * 
 * The content of the stack is set to zero (i.e. stack is empty). Memory
 * is not deallocated.
 * 
 * @param stack The stack to reset.
 */
static inline void
StackC_reset ( StackC *const stack ) {
	(*stack).len = 0;
}


/**
 * @brief returns the top element of the char stack
 * 
 * Returns the top element of the char stack.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param stack The stack.
 * @return Top element of the char stack.
 */
static inline char
StackC_top ( StackC *const stack ) {
	return (*stack).values[(*stack).len-1];
}


/**
 * @brief puts an element on the char stack
 * 
 * Puts an element on the char stack.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param stack The stack to put an element on.
 * @param el The element to put on the stack.
 */
extern void
StackC_push ( StackC *const stack, const char el );


/**
 * @brief removes the top element of the char stack
 * 
 * Removes the top element of the char stack and returns it.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param stack The stack.
 * @return The top element of the stack that has been removed.
 */
static inline char
StackC_pop ( StackC *const stack ) {
	(*stack).len -= 1;
	return (*stack).values[(*stack).len];
}




/**
 * @brief initializes the size_t stack
 * 
 * Allocates the needed space to store the memory of the stack. This serves as
 * an initialization for the stack. The initial capacity of the stack is
 * (256-8) / sizeof(size_t).
 * <br><br>
 * If the stack already contains valid space it is not freed.
 * 
 * @param stack The stack to initialize.
 */
extern void
StackS_allocate ( StackS *const stack );

/**
 * @brief destroys the size_t stack
 * 
 * Deallocates all memory the stack is holding. To use the stack again you
 * have to call StackC_allocate.
 * 
 * @param stack The stack to destroy.
 */
extern void
StackS_deallocate ( StackS *const stack );


/**
 * @brief sets the content to zero (but keeps the stack memory)
 * 
 * The content of the stack is set to zero (i.e. stack is empty). Memory
 * is not deallocated.
 * 
 * @param stack The stack to reset.
 */
static inline void
StackS_reset ( StackS *const stack ) {
	(*stack).len = 0;
}


/**
 * @brief returns the top element of the size_t stack
 * 
 * Returns the top element of the size_t stack.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param stack The stack.
 * @return Top element of the size_t stack.
 */
static inline size_t
StackS_top ( StackS *const stack ) {
	return (*stack).values[(*stack).len-1];
}


/**
 * @brief puts an element on the size_t stack
 * 
 * Puts an element on the size_t stack.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param stack The stack to put an element on.
 * @param el The element to put on the stack.
 */
extern void
StackS_push ( StackS *const stack, const size_t el );


/**
 * @brief removes the top element of the size_t stack
 * 
 * Removes the top element of the size_t stack and returns it.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param stack The stack.
 * @return The top element of the stack that has been removed.
 */
static inline size_t
StackS_pop ( StackS *const stack ) {
	(*stack).len -= 1;
	return (*stack).values[(*stack).len];
}




/**
 * @brief initializes the pointer stack
 * 
 * Allocates the needed space to store the memory of the stack. This serves as
 * an initialization for the stack. The initial capacity of the stack is
 * (512-8) / sizeof(void*).
 * <br><br>
 * If the stack already contains valid space it is not freed.
 * 
 * @param stack The stack to initialize.
 */
extern void
StackP_allocate ( StackP *const stack );

/**
 * @brief destroys the pointer stack
 * 
 * Deallocates all memory the stack is holding. To use the stack again you
 * have to call StackC_allocate.
 * 
 * @param stack The stack to destroy.
 */
extern void
StackP_deallocate ( StackP *const stack );


/**
 * @brief sets the content to zero (but keeps the stack memory)
 * 
 * The content of the stack is set to zero (i.e. stack is empty). Memory
 * is not deallocated.
 * 
 * @param stack The stack to reset.
 */
static inline void
StackP_reset ( StackP *const stack ) {
	(*stack).len = 0;
}


/**
 * @brief returns the top element of the pointer stack
 * 
 * Returns the top element of the pointer stack.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param stack The stack.
 * @return Top element of the pointer stack.
 */
static inline void *
StackP_top ( StackP *const stack ) {
	return (*stack).values[(*stack).len-1];
}


/**
 * @brief puts an element on the pointer stack
 * 
 * Puts an element on the pointer stack.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param stack The stack to put an element on.
 * @param el The element to put on the stack.
 */
extern void
StackP_push ( StackP *const stack, void *const el );


/**
 * @brief removes the top element of the pointer stack
 * 
 * Removes the top element of the pointer stack and returns it.<br>
 * <b>CAUTION</b>: No bounds check is done here.
 * 
 * @param stack The stack.
 * @return The top element of the stack that has been removed.
 */
static inline void *
StackP_pop ( StackP *const stack ) {
	(*stack).len -= 1;
	return (*stack).values[(*stack).len];
}


#endif /* STACK_H */















