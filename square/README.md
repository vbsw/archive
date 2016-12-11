# Square Effect

## About
Square Effect is a shoot 'em up game with squares as ships.

Square Effect is created by Vitali Baumtrok <vbit@mailbox.org> and published at <https://gitlab.com/vbit/square>.

## Copying
Square Effect is distributed under the terms of the Boost Software License, version 1.0.
(See accompanying file LICENSE or copy at <http://www.boost.org/LICENSE_1_0.txt>)

Square Effect is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the Boost Software License for more details.

## Dependencies (GNU/Linux)
Packages

	libc6-dev
	libX11-dev
	libgl1-mesa-dev

Libraries

	cds  <https://gitlab.com/vbit/cds>
	ogfx <https://gitlab.com/vbit/ogfx>

Programs

	gcc
	make

## Compiling (GNU/Linux)
To compile you need cds and ogfx. Place them in same directory as square, then:

	$ cd square
	$ make

## Audio
There is no audio in Square Effect.

## Controls

	left      j or arrow left
	right     l or arrow right
	up        i or arrow up
	down      k or arrow down
	action    f or enter
	back      escape

## Using Git
How to install Git: <https://git-scm.com/book/en/v2/Getting-Started-Installing-Git>

Get the master branch and all refs of this project:

	$ git clone https://gitlab.com/vbit/square.git
	$ cd square

See local and remote branches:

	$ git branch -a

See tracked remote branches:

	$ git branch -vv

Update all tracked branches and all refs:

	$ git fetch

