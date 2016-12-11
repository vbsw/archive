# round

## Abstract
round is a package that provides round functions written in the programming language Go. This package does NOT claim to satisfy the requirement of the IEEE 754-1985 standard!

round is created by Vitali Baumtrok <vbsoftware@mailbox.org> and published at <https://github.com/vitalibaumtrok/round>.

## Copying
round is distributed under the Boost Software License, Version 1.0. (See accompanying file LICENSE or copy at <http://www.boost.org/LICENSE_1_0.txt>)

## Using Go
How to install Go: <https://golang.org/doc/install>

Get this project:

	$ go get github.com/vitalibaumtrok/round

Update a local copy:

	$ go get -u github.com/vitalibaumtrok/round

Compile:

	$ go install github.com/vitalibaumtrok/round

Run tests:

	$ go test github.com/vitalibaumtrok/round

## Using Git
How to install Git: <https://git-scm.com/book/en/v2/Getting-Started-Installing-Git>

Get the master branch and all refs of this project:

	$ git clone https://github.com/vitalibaumtrok/round.git
	$ cd round

See all tags:

	$ git tag -l

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

- <https://github.com/golang/go/issues/4594>
- <https://en.wikipedia.org/wiki/IEEE_floating_point#Rounding_rules>
- <http://www.jhauser.us/arithmetic/SoftFloat.html>
- <https://en.wikipedia.org/wiki/Double-precision_floating-point_format>


