# JSHiF

## Abstract
JSHiF is a program to process documents in Simple Hierarchal Format (SHiF).

## Copyright
Copyright 2015, Vitali Baumtrok.

JSHiF is distributed under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

## Clone repository

[Download](http://git-scm.com/downloads) and install Git.

To clone this repository to folder "JSHiF" run:

	$ git clone https://github.com/vitalibaumtrok/jshif JSHiF

## Using Eclipse (IDE)
If you want to edit this project in Eclipse, then download the source to workspace of Eclipse, creating a new folder, for example folder "JSHiF". Open Eclipse and create new Java Project with the same name as the new folder, in this case "JSHiF".

## Development
JSHiF is programmed in Java. There is only one branch: master.

## SHiF Specification

### Line End
A text line ends with bytes used by the operating system to end lines,
for example LF (0x0A) or CR LF (0x0D0A).

### Indent Token
The bytes of the first line define the _indent token_. If the first
line is empty then the _indent token_ is TAB (0x09).

### Node
A node has one _node name_, one _node value_, one _indent depth_ and multiple
_subnodes_. A node is defined by defining its _node name_.

### Root Node
The most top node is called _root node_ and its _node name_ is "root". It has
no _node value_ and its _indent depth_ is -1.

### Subnode
A node can have 0 or more _subnodes_, which are nodes. A _subnode_ can only
be the _subnode_ of one node, which is called _parent node_.

If a node has the _indent depth_ of n, then its _subnodes_ have the
_indent depth_ of n+1. _Subnodes_ of _indent depth_ bigger then n+1 are
not allowed and are skipped or cause an error when parsed.

### Node Name
The _node name_ is defined by characters starting by other then the
_indent token_. A node which _node name_ starts with blank (0x20) or
TAB (0x09) is skipped or causes an error when parsed.

The _node name_ stands at the beginning of a new line and can be preceeded
by 0 or more _indent tokens_.

### Node Value
The _node value_ stands in the same line as the _node name_ and is
introduced by one blank (0x20) or TAB (0x09). The blank or TAB is not part of
the _node value_. A _node value_ always exists. If it is not present in the
document then its empty.

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
