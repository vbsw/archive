# input

## Abstract
Global unsynchronized event buffer, written in programming language Go.

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

## Development
Programming language Go is used.

## Installation
Install the programming language [Go](https://golang.org/doc/install) and
the source code management system [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).

Get this project with

	$ go get github.com/vitalibaumtrok/input

or update your local copy with

	$ go get -u github.com/vitalibaumtrok/input

or compile it with

	$ go install github.com/vitalibaumtrok/input

## Example

### Writing to Buffer

	input.AddKeyEvent(input.KEY_PRESSED, 27)
	input.AddPosition2FloatEvent(input.POSITION_2FLOAT_NORMAL, 0, 10.0, 20.0)

### Reading from Buffer

	input.SwapReadWriteBuffers()

	for input.HasNextKeyEvent() {

		switch input.KeyEvent() {

			case input.KEY_PRESSED:
			switch input.KeyCode() {
				case 37:
				goLeft()

				case 38:
				goUp()

				case 39:
				goRight()

				case 40:
				goDown()
			}

			case input.KEY_RELEASED:
			if input.KeyCode() == 27 {
				quitApplication()
			}
		}
	}
