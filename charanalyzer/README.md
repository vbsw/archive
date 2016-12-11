# Character Analyzer

## Abstract
Character Analyzer is an incremental text parser.

## Copyright
Copyright 2015, Vitali Baumtrok.

Character Analyzer is distributed under the Boost Software License, Version 1.0.
(See accompanying file LICENSE or copy at http://www.boost.org/LICENSE_1_0.txt)

## Clone Repository
[Download](http://git-scm.com/downloads) and install Git.

Clone this project to folder "CharAnalyzer":

	$ git clone https://github.com/vitalibaumtrok/charanalyzer CharAnalyzer

## Development
Programming language C is used.

## Compiling
Create an object file "charnalyzer.o":

	$ gcc -S charnalayzer.c -o charnalyzer.o

Compile and run function tests:

	$ gcc functiontests.c -o test
	$ ./test

## Example

	const char sample[14] = "Hello, World!";
	CharAnalyzer_Result_T result;

	CharAnalyzer_Init( &result, 2048 );
	CharAnalyzer_Parse( &result, sample, sizeof(sample)-1 );
	FooBar( &result );

	CharAnalyzer_Destroy( &result );

