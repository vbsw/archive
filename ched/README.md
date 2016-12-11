# ChEd

## Abstract
ChEd is a text editor.

## Copying
Copyright 2016, Vitali Baumtrok.

ChEd is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

ChEd is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program.  If not, see <http://www.gnu.org/licenses/>.

## Clone Repository
[Download](http://git-scm.com/downloads) and install Git.

Clone this project to folder "ChEd":

	$ git clone https://github.com/vitalibaumtrok/ched ChEd

## Development
Programming language C is used.

## Dependencies

	libgtk-3-dev

## Documentation

### Module Usage

	args
		output
		parse
		stdlib
	buffer
		stdlib
	clock
		buffer
		time
	gui
		gtk/gtk
		settings
		string
		strings
		clock
	input
		buffer
		stdio
	main
		args
		gui
		settings
		strings
		tui
	output
		stdio
		strings
	parse
		-
	settings
		args
		input
		output
		stdlib
	strings
		-
	tui
		args
		output
		stdlib

