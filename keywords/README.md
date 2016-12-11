# Key Words

## Abstract
Key Words is a program to manage data by key words.

## Copyright
Copyright 2015, Vitali Baumtrok.

Key Words is distributed under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

## Clone repository

[Download](http://git-scm.com/downloads) and install Git.

To clone this repository to folder "KeyWords" run:

	$ git clone https://github.com/vitalibaumtrok/keywords KeyWords

## Development
Key Words is programmed in C. There is only one branch: master.

## Dependencies
libgtk2.0-dev

## Compile
To compile do

	$ gcc -Wall -g main.c -o keywords `pkg-config --cflags --libs gtk+-2.0`

or run Makefile

	$ make

## Releases
None.
