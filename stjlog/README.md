# stjlog

## Abstract
stjlog is a logger written in Java. It uses static methods and
static public variables.

The reason this logger provides only static methods and static fields
is that their use can be easily traced back, because there are no
abstraction layers. This is useful for fast prototyping.

## Copying
This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <http://unlicense.org>

## Clone Repository
[Download](http://git-scm.com/downloads) and install Git.

To clone this repository to folder "StJLog" run:

	$ git clone https://github.com/vitalibaumtrok/stjlog StJLog

## Using Eclipse IDE
If you want to edit this project in Eclipse, then download it
to the workspace of Eclipse, for example to folder "StJLog".
Open Eclipse and crete new Java Project with the same name as
the folder, in this case "StJLog".

## How to use
To log a message use one of the following methods

	Log.error ( message );
	Log.warning ( message );
	Log.info ( message );
	Log.trace ( message );

To set the logger output to a file use

	Log.output = LogOutput.FILE;
	Log.file = Paths.get ( System.getProperty ( "user.home" ), "app.log" );

To change the output to terminal use

	Log.output = LogOutput.STD_OUT;

To disable a loggin level use some of the following

	Log.levelError = false;
	Log.levelWarning = false;
	Log.levelInfo = false;
	Log.levelTrace = false;

To print message with stack trace use

	try { ... }
	catch ( Throwable t )
	{
		Log.error ( "resource not available", t );
	}
