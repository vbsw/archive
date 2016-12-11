# strl

## Abstract
strl is a library of functions, written in C, to manage C strings in a list.
The list is implemented as a buffered array.

## Copyright
Copyright 2016, Vitali Baumtrok.

strl is distributed under the terms of the Boost Software License, Version 1.0.
(See accompanying file LICENSE or copy at <http://www.boost.org/LICENSE_1_0.txt>)

## Download
You can download strl from <https://github.com/vitalibaumtrok/strl/archive/master.zip>.

## Tests
Compile and run tests:

	$ make
	$ ./test

## Using Git
How to install Git: <https://git-scm.com/book/en/v2/Getting-Started-Installing-Git>

Get the master branch and all refs of this project:

	$ git clone https://github.com/vitalibaumtrok/strl.git
	$ cd strl

See local and remote branches:

	$ git branch -a

Checkout other branches than master, for example the development branch:

	$ git branch development origin/development
	$ git checkout development

See tracked remote branches:

	$ git branch -vv

Update all tracked branches and all refs:

	$ git fetch

## Example

	void* strl = strlnew ( );
	strlapnd ( strl, "Hello" );
	strlapnd ( strl, "World" );
	printf ( "%s, %s!\n", strlarr(strl)[0], strlarr(strl)[1] );

	strlrm ( strl, 1 );
	strlapnd ( strl, "Bob" );
	printf ( "%s, %s!\n", strlarr(strl)[0], strlarr(strl)[1] );

	strldest ( strl );

Output:

	Hello, World!
	Hello, Bob!

## Documentation

__strlnew__ : Allocates new list with default capacity.

	return           The new allocated list or 0, if malloc fails.

__strlnewcap__ : Allocates new list with given capacity.

	size_t capacity  The capacity of new list. Must not be 0 or the default capacity function will fail.
	return           The new allocated list or 0, if malloc fails.

__strldest__ : Deallocates the list and its content.

	void* strl      The list to deallocate.

__strllen__ : Returns the length, i.e. number of elements, of the list.

	const void* strl The list to retrieve the length from.
	return           Number of elements the list contains.

__strlcap__ : Returns the capacity of the list, i.e. the maximum number of elements the list can hold before reallocating.

	const void* strl The list to retrieve the capacity from.
	return           Maximum number of elements the list can hold.

__strlerr__ : Returns the error number from the list. The number is > 0, if malloc fails somewhere.

	const void* strl The list to retrieve the error number from.
	return           0, if no error.

__strlerrclear__ : Resets the error number.

	void* strl       The list to reset the error number.

__strlcapfunc__ : Sets the capacity computing function for the list.

	void* strl       The list to set the capacity function.
	size_t (*capfunc) ( size_t n, size_t capacity ) Capacity function.

__strlapnd__ : Makes a copy of str and appends it to the list. Resets the error number.

	void* strl       The list to append the C string to.
	const char* str  The C string to copy and to append.

__strlapndref__ : Appends str to the list (does not copy str). Resets the error number.

	void* strl       The list to append the C string to.
	char* str        The C string to append.

__strlrm__ : Removes an element from the list at index and releases it from memory.

	void* strl       The list to remove the element from.
	size_t index     The position of the element to release from memory and to remove from list.

__strlrmref__ : Removes an element from the list at index.

	void* strl       The list to remove the element from.
	size_t index     The position of the element to remove from list.
	return           Pointer to the C string, that has been removed.

__strlclear__ : Releases all the C strings from memory and sets the length of the list to 0.

	void* strl       The list to remove all elements from.

__strlclearref__ : Sets the length of the list to 0. (Does not release the C strings from memory.)

	void* strl       The list to remove all elements from.

__strlins__ : Makes a copy of str and inserts it at index into list. Resets the error number.

	void* strl       The list to insert the C string into.
	size_t index     The position where to insert the element.
	const char* str  The C string to copy and to insert.

__strlinsref__ : Inserts str at index into list (does not copy str). Resets the error number.

	void* strl       The list to insert the C string into.
	size_t index     The position where to insert the element.
	char* str        The C string to insert.

