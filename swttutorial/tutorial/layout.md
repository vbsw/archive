# SWT Tutorial

## Layout
The class Layout is a **layout manager**, i.e. manages the layout of Controls.

	public void adjustLayout ( final Composite composite ) {
		final Layout layout = new FillLayout ( SWT.HORIZONTAL );
		composite.setLayout ( layout );
	}

Controls in a Composite are positioned depending on the Layout and their **layout data**.
Yes, you set the layout data on the Control, not on the Container or the Layout.
This way the layout manager knows how to position each Control.

	public void adjustInGrid ( final Composite composite, final Button button1, final Button button2, final Button button3 ) {
		composite.setLayout ( new GridLayout ( 2, true ) );
		button1.setLayoutData ( new GridData ( GridData.FILL, GridData.FILL, true, true, 1, 1 ) );
		button2.setLayoutData ( new GridData ( GridData.FILL, GridData.FILL, true, true, 1, 1 ) );
		button3.setLayoutData ( new GridData ( GridData.FILL, GridData.FILL, true, true, 2, 1 ) );
	}

Not every Layout has a corresponding layout data class. To set the spacing between the
Controls inside a Composite you have to modify the layout and, if available, the layout data of the Control.

To **remove the border** of a Composite choose the style SWT.NONE for it.

	new Composite ( window, SWT.NONE );

[Layout classes reference](http://help.eclipse.org/luna/index.jsp?topic=/org.eclipse.platform.doc.isv/reference/api/org/eclipse/swt/layout/package-summary.html).

## Navigation
[Readme](https://github.com/vitalibaumtrok/swttutorial/blob/master/README.md)

1. [Library Download](https://github.com/vitalibaumtrok/swttutorial/blob/master/tutorial/libdownload.md)
2. [API Reference](https://github.com/vitalibaumtrok/swttutorial/blob/master/tutorial/docref.md)
3. [Display](https://github.com/vitalibaumtrok/swttutorial/blob/master/tutorial/display.md)
4. [Widgets](https://github.com/vitalibaumtrok/swttutorial/blob/master/tutorial/widgets.md)
5. Layout
