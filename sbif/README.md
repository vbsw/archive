# SBiF

## Abstract
Simple Binary Format (or short SBiF) is a document format to store data
in sequently laid out chunks.

The accompanying source code is a reference implementation for a SBiF
parser.

## SBiF Parser

### Copyrigth
Copyright Vitali Baumtrok 2015.

The SBiF parser is distributed under the Boost Software License, Version 1.0.
(See accompanying file LICENSE_1_0.txt or copy at
http://www.boost.org/LICENSE_1_0.txt)

### Clone Repository
[Download](http://git-scm.com/downloads) and install Git.

Clone this project to folder "sbif" with:

	$ git clone https://github.com/vitalibaumtrok/sbif sbif

### Function Tests
Compile with:

	$ gcc sbif.test.main.c sbif.c -o main

Run main with:

	$ ./main

## SBiF Specification

Data in SBiF has the following layout

	    FILE ::= SIZE FRAGMENT
	FRAGMENT ::= SOFS FS FF | CD
	    SOFS ::= SIZE
	      FS ::= SIZE*
	      FF ::= FRAGMENT*

**FILE** is a sequence of two elements (SIZE and FRAGMENT). The first element SIZE represents the size (in bytes) of the element FRAGMENT.

**SIZE** is a (signed or unsigned) number encoded in 8 bytes. The first byte is the most significant byte.

**FRAGMENT** is either a sequence of three elements (SOFS, FS and FF), or just one element (CD).

**SOFS** stands for "Size Of Fragement Sizes". It represents the size (in bytes) of the element FS, not the number of elements in FS.

**FS** stands for "Fragment Sizes". It is a sequence of N elements which are SIZEs. The n-th element of FS represents the size (in bytes) of the n-th element in FF. It must apply N = M, where M is the number of elements in FF, i.e. |FS| = |FF|. 

**FF** stands for "Fragments". It is a sequence of M elements which are FRAGMENTs. It must apply N = M, where N is the number of elements in FS, i.e. |FS| = |FF|.

**CD** stands for "Custom Data". It is data that is undefined by SBiF and must be parsed by the user.
