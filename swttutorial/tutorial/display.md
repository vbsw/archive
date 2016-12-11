# SWT Tutorial

## Display
A new Display allocates **system resources**.

The Display manages the creation of **graphical elements**.

	final Display display = new Display ();
	final Shell shell = new Shell ( display );

The Display manages **event handling**. The method _readAndDispatch_ causes the processing of input.
All GUI elements, that are created on this Display, process the input.

	display.readAndDispatch ();

The first Display created becomes the **default Display**. It's right to assume that only one Display is created,
which you can access with the static method _getDefault_. (If the default Display is null, then a new one is created.)

	final Display displayA = new Display ();
	final Display displayB = Display.getDefault ();
	assert ( displayA==displayB );

To **release system resources** call the method _dispose_. If this Display was the default one, then the default one is set to null.

	display.dispose ();

The **main loop** should be something like this:

	public static void main ( String args[] ) {
		final Display display = new Display ();
		final Shell mainWindow = new Shell ( display );
		while ( !mainWindow.isDisposed () )
			if ( !display.readAndDispatch () )
				display.sleep ();
		display.dispose ();
	}

[Display class reference](http://help.eclipse.org/luna/topic/org.eclipse.platform.doc.isv/reference/api/org/eclipse/swt/widgets/Display.html).

## Navigation
[Readme](https://github.com/vitalibaumtrok/swttutorial/blob/master/README.md)

1. [Library Download](https://github.com/vitalibaumtrok/swttutorial/blob/master/tutorial/libdownload.md)
2. [API Reference](https://github.com/vitalibaumtrok/swttutorial/blob/master/tutorial/docref.md)
3. Display
4. [Widgets](https://github.com/vitalibaumtrok/swttutorial/blob/master/tutorial/widgets.md)
5. [Layout](https://github.com/vitalibaumtrok/swttutorial/blob/master/tutorial/layout.md)
