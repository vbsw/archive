# SHiF

## Abstract
Simple Hierarchal Format (or short SHiF) is a document format to store data
as human readable text in a hierarchal structure.

The accompanying source code is a reference implementation for a SHiF
parser.

## SHiF Parser

### Copyrigth
Copyright Vitali Baumtrok 2015.

The SHiF parser is distributed under the Boost Software License, Version 1.0.
(See accompanying file LICENSE_1_0.txt or copy at
http://www.boost.org/LICENSE_1_0.txt)

### Clone Repository
[Download](http://git-scm.com/downloads) and install Git.

Clone this project to folder "shif" with:

	$ git clone https://github.com/vitalibaumtrok/shif shif

### Function Tests
Compile with:

	$ gcc shif.test.main.c shif.c -o main

Run main with:

	$ ./main

## SHiF Specification

### Line End
A text line ends with the tokens used by the operating system to end lines,
for example LF (0x0A) or CR LF (0x0D0A).

### Indent Token
The characters of the first line define the _indent token_. If the first
line is empty then the _indent token_ is TAB (0x09).

### Node
A node has a _node name_, a _node value_, an _indent depth_ and _subnodes_.
A node is defined by defining its _node name_.

### Root Node
The most top node is called _root node_ and its _node name_ is "root". It has
no _node value_ and its _indent depth_ is -1.

### Subnode
A node can have 0 or more _subnodes_, which are nodes. A _subnode_ can only
be the _subnode_ of one node, which is called _parent node_.

If a node has the _indent depth_ of n, then its _subnodes_ have the
_indent depth_ of n+1. _Subnodes_ of _indent depth_ bigger then n+1 are
not allowed and should cause an error when parsed.

### Node Name
The _node name_ is defined by characters starting by other then the
_indent token_, blank (i.e. SP (0x20)), tab (i.e. TAB (0x09)) or
the line end token. A _node name_ starting with blank or tab should
cause an error when parsed.

The _node name_ stands at the beginning of a new line and can be preceeded
by 0 or more _indent tokens_.

### Node Value
The _node value_ stands in the same line as the _node name_ and is
introduced by one blank, i.e. SP (0x20). The blank is not part of the
_node value_. A _node value_ must always be processed by a parser, even if it's
empty or does not exist.

A parser may remove all blanks or tabs at the beginning and ending of a
_node value_. So a _node value_ can be indented behind the _node name_.

### Indent Depth
The number of _indent tokens_ proceeding the _node name_ defines the
_indent depth_ of a node.

### Example A
The first line contains the sharp character. This means the _indent token_
is a sharp character.

	#
	document
	#author
	##name
	###first Tom
	###second Sawyer

### Example B
The first line is empty, which means the _indent token_ is tab.

	
	document
		id 123456789
		title Just an Example
		content
			paragraph
				line Hello,
				line world!
	History
		change 1
			date 2015-04-04
			comment basic structure
