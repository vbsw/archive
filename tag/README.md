# tag

## Abstract
tag, version 2.1.0, is a program to manage words with tags.

## Functionality
- saves a word with tags in a file
- searches words by tags in a file
- exports/imports words and tags to/from text file

## Development
Programming language Go is used.

The encoding of tags is described on <https://github.com/vitalibaumtrok/gobase/blob/master/tag/spec.md>

## Installation
Install the programming language Go as described here <https://golang.org/doc/install>. Install Git.

Downloadd the sources with

	$ go get github.com/vitalibaumtrok/gobase
	$ go get github.com/vitalibaumtrok/tag

If you already have the sources, update them with

	$ go get -u github.com/vitalibaumtrok/gobase
	$ go get -u github.com/vitalibaumtrok/tag

Install _tag_ with

	$ go install github.com/vitalibaumtrok/tag

## Copyright
Copyright 2014 Vitali Baumtrok

This program is distributed under the terms of the Boost Software License, Version 1.0, as published on <http://www.boost.org/LICENSE_1_0.txt>.