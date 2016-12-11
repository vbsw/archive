# testogl

## Abstract
testogl is written in C and contains various test for OpenGL and GLUT for
Linux.

## Copying
This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <http://unlicense.org>

## Requirements
You need following programs and/or libraries to compile testogl:

* gcc
* Make
* GLUT (for example the packages freeglut3 and freeglut3-dev)

## Compiling
To compile a test run on command line:

	$ make TEST

replacing TEST with one of the following

* primdraw (draws primitives)
* primscale (keep aspect ratio when window size changes)
* winmod (key 'o' sets window original size; klick and hold to drag window)
* texture (shows texture)
* wireframes (shows primitives as wire frames)
* gluttext (shows text created by GLUT)

To compile for example the test 'primdraw' run:

	$ make primdraw

To start the program run:

	$ ./testogl
