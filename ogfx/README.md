# OGfx

## About
OGfx (OpenGL Graphics Utility) is a library that provides graphical output using OpenGL on Linux. It is written in C.

OGfx is created by Vitali Baumtrok <vbit@mailbox.org> and published at <https://gitlab.com/vbit/ogfx>.

## Copying
OGfx is distributed under the terms of the Boost Software License, version 1.0.
(See accompanying file LICENSE or copy at <http://www.boost.org/LICENSE_1_0.txt>)

OGfx is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the Boost Software License for more details.

## Dependencies (GNU/Linux)
Packages

	libc6-dev
	libX11-dev
	libgl1-mesa-dev

Libraries

	cds <https://gitlab.com/vbit/cds>

Programs

	gcc
	make
	ar

ar is in package binutils.

## Compiling (GNU/Linux)
To create static link library on Linux run:

	$ make

## Using Git
Get the master branch and all refs of this project:

	$ git clone https://gitlab.com/vbit/ogfx.git

See local and remote branches:

	$ git branch -a

Check out branches other than master, for example the "development" branch:

	$ git branch development origin/development
	$ git checkout development

See tracked remote branches:

	$ git branch -vv

Update all tracked branches and all refs:

	$ git fetch

## References

- <https://git-scm.com/book/en/v2/Getting-Started-Installing-Git>
- <http://www.adp-gmbh.ch/cpp/gcc/create_lib.html>
- <http://www.rapidtables.com/code/linux/gcc/gcc-wall.htm>
- <https://www.gnu.org/software/binutils/>
- <https://www.opengl.org/wiki/Programming_OpenGL_in_Linux:_GLX_and_Xlib>
- <http://www.x.org/releases/X11R7.6/doc/libX11/specs/libX11/libX11.html>
- <https://www.opengl.org/sdk/docs/man2/xhtml/#glX>
- <https://tronche.com/gui/x/xlib/>
- <https://en.wikibooks.org/wiki/X_Window_Programming/Xlib>

