# SWT Tutorial

## Control is a Widget!
The **class Control** is a subclass of class Widget. All visual widgets are controls.

	+-----------------------------------+
	|         java.lang.Object          |
	+-----------------------------------+
	                  ^
	                  |
	+-----------------------------------+
	|  org.eclipse.swt.widgets.Widget   |
	+-----------------------------------+
	                  ^
	                  |
	+-----------------------------------+
	|  org.eclipse.swt.widgets.Control  |
	+-----------------------------------+

For example, the class Button is a direct subclass of Control.

[Control class reference](http://help.eclipse.org/luna/topic/org.eclipse.platform.doc.isv/reference/api/org/eclipse/swt/widgets/Control.html)

## Composite is a container!
The **class Composite** is a container to place widgets into.

	+--------------------------------------+
	|   org.eclipse.swt.widgets.Control    |
	+--------------------------------------+
	                   ^
	                   |
	+--------------------------------------+
	|  org.eclipse.swt.widgets.Scrollable  |
	+--------------------------------------+
	                   ^
	                   |
	+--------------------------------------+
	|  org.eclipse.swt.widgets.Composite   |
	+--------------------------------------+

The class Group is a direct subclass of Composite and is the same as Composite, but has a border and a title.

On a Composite you can set a **layout**. (Shell is a Composite)

	final Shell shell = new Shell ( Display.getDefault () );
	shell.setLayout ( new FillLayout () );

[Composite class reference](http://help.eclipse.org/luna/topic/org.eclipse.platform.doc.isv/reference/api/org/eclipse/swt/widgets/Composite.html)

## Navigation
[Readme](https://github.com/vitalibaumtrok/swttutorial/blob/master/README.md)

1. [Library Download](https://github.com/vitalibaumtrok/swttutorial/blob/master/tutorial/libdownload.md)
2. [API Reference](https://github.com/vitalibaumtrok/swttutorial/blob/master/tutorial/docref.md)
3. [Display](https://github.com/vitalibaumtrok/swttutorial/blob/master/tutorial/display.md)
4. Widgets
5. [Layout](https://github.com/vitalibaumtrok/swttutorial/blob/master/tutorial/layout.md)
