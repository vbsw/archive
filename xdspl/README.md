# XDspl

[![GoDoc](https://godoc.org/github.com/vbsw/xdspl?status.svg)](https://godoc.org/github.com/vbsw/xdspl)

## About
XDspl prints the name of the display server, if it is running, otherwise "n.a." (not available).
XDspl is programmed in Go and published at <https://github.com/vbsw/xdspl>.

## Copyright
Copyright 2016 Vitali Baumtrok (vbsw@mailbox.org)

XDspl is distributed under the terms of the Boost Software License, version 1.0.
(See accompanying file LICENSE or copy at <http://www.boost.org/LICENSE_1_0.txt>)

XDspl is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the Boost Software License for more details.

## Example A
Input

	$ xdspl

Output

	$ :0

If the display server is not running the output is

	$ n.a.

## Example B
Input

	$ xdspl screens

Output

	$ 2 1920x1080 1920x1080

If the display server is not running the output is

	$ n.a.

## References

- <https://golang.org/doc/install>
- <https://git-scm.com/book/en/v2/Getting-Started-Installing-Git>

