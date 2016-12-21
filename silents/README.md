# Silent Shooter

## About
Silent Shooter is a shoot 'em up game with squares as ships.
It is programmed in Go and published at <https://github.com/vbsw/silents>.

## Copyright
Copyright 2016 Vitali Baumtrok (vbsw@mailbox.org)

Silent Shooter is distributed under the terms of the Boost Software License, version 1.0.
(See accompanying file LICENSE or copy at <http://www.boost.org/LICENSE_1_0.txt>)

Silent Shooter is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the Boost Software License for more details.

## Controls

	left      j or arrow left
	right     l or arrow right
	up        i or arrow up
	down      k or arrow down
	action    f or enter
	back      escape

## Audio
There is no audio in Silent Shooter.

## Dependencies (GNU/Linux)
Packages

	libc6-dev
	libX11-dev
	libgl1-mesa-dev

Libraries

	vis     https://github.com/vbsw/vis
	visimpl https://github.com/vbsw/visimpl

Programs

	gcc
	go

## Compiling (GNU/Linux)
To compile you need cds and ogfx. Place them in same directory as silents, then:

	$ cd silents
	$ make

## Using Git
Get the master branch and all refs of this project:

	$ git clone https://github.com/vbsw/square.git
	$ cd square

See local and remote branches:

	$ git branch -a

See tracked remote branches:

	$ git branch -vv

Update all tracked branches and all refs:

	$ git fetch

## References

- https://golang.org/doc/install
- https://git-scm.com/book/en/v2/Getting-Started-Installing-Git

