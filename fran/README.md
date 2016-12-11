# fran

## Abstract
fran is a program to play frame animations stored in PNGs. It works on Linux.

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

## Installation
Install the programming language [Go](https://golang.org/doc/install) and
the source code management system [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git).

Get this project with

	$ go get github.com/vitalibaumtrok/fran

or update your local copy with

	$ go get -u github.com/vitalibaumtrok/fran

or compile it with

	$ go install github.com/vitalibaumtrok/fran

## Development
Programming language Go is used.

Dependencies:

	github.com/vitalibaumtrok/glut
	github.com/vitalibaumtrok/gl10
	github.com/vitalibaumtrok/gl11

## Usage
To play an animation stored in file "move.png", which contains 8x4 images, with a delay of 30 milliseconds between each image, type:

	fran move.png 8x4 30

To play an animation stored in file "anim.png", which contains 8x4 images, with a delay of 30 milliseconds between each image and a delay of 50 milliseconds between replays, type:

	fran anim.png 8x4 30 50

You can control the animation with the following input:

	[key]            [effect]
	o                set window to fit original size of image
	p                pause or play animation
	r                reload image
	+                zoom in
	-                zoom out
	left arrow       pause animation and show previous frame
	right arrow      pause animation and show next frame
	up arrow         increase the delay between images by 1 millisecond
	down arrow       decrease the delay between images by 1 millisecond
	ctrl+arrow       show previous/next image or increase/decrease the delay by 10

