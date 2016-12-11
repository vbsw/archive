# CDS

## About
CDS (C Data Structures) is a library of data structures for the programming language C.

CDS is created by Vitali Baumtrok <vbit@mailbox.org> and published at <https://gitlab.com/vbit/cds>.

## Copying
CDS is distributed under the terms of the Boost Software License, version 1.0.
(See accompanying file LICENSE or copy at <http://www.boost.org/LICENSE_1_0.txt>)

CDS is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the Boost Software License for more details.

## Dependencies (GNU/Linux)
Packages

	libc6-dev

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

	$ git clone https://gitlab.com/vbit/cds.git

See local and remote branches:

	$ git branch -a

Checkout other branches than master, for example the development branch:

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

