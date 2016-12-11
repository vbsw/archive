/*
 * This is free and unencumbered software released into the public domain.
 * 
 * Anyone is free to copy, modify, publish, use, compile, sell, or
 * distribute this software, either in source code form or as a compiled
 * binary, for any purpose, commercial or non-commercial, and by any
 * means.
 * 
 * In jurisdictions that recognize copyright laws, the author or authors
 * of this software dedicate any and all copyright interest in the
 * software to the public domain. We make this dedication for the benefit
 * of the public at large and to the detriment of our heirs and
 * successors. We intend this dedication to be an overt act of
 * relinquishment in perpetuity of all present and future rights to this
 * software under copyright law.
 * 
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 * IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
 * OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
 * ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
 * OTHER DEALINGS IN THE SOFTWARE.
 * 
 * For more information, please refer to <http://unlicense.org>
 */
package de.vbsoftware.sample.swt;

import org.eclipse.swt.SWT;
import org.eclipse.swt.layout.GridData;
import org.eclipse.swt.layout.GridLayout;
import org.eclipse.swt.widgets.Button;
import org.eclipse.swt.widgets.Composite;
import org.eclipse.swt.widgets.Display;
import org.eclipse.swt.widgets.Shell;

/**
 * A class to create a window with three buttons inside. A GridLayout is used to
 * arrange the buttons.
 * 
 * @author Vitali Baumtrok
 */
public class XmplGridLayout {

	/**
	 * The name of the window.
	 */
	private static final String	NAME	= "SWT Tutorial: Grid Layout";

	/**
	 * The width of the window.
	 */
	private static final int	WIDTH	= 480;

	/**
	 * The height of the window.
	 */
	private static final int	HEIGHT	= 360;

	/**
	 * Creates a window. The window contains three buttons with no function.
	 * Two buttons in a row and one under them.
	 * 
	 * @param display
	 *        The display to create the window on.
	 * @return A window.
	 */
	public static Shell createAndOpenWindow ( final Display display ) {
		final int x = ( display.getBounds ().width - WIDTH ) / 2;
		final int y = ( display.getBounds ().height - HEIGHT ) / 2;
		final Shell window = new Shell ( display );
		final Button buttonAlice = createButton ( window, "Alice" );
		final Button buttonBob = createButton ( window, "Bob" );
		final Button buttonCharlie = createButton ( window, "Charlie" );

		adjustInGrid ( window, buttonAlice, buttonBob, buttonCharlie );
		window.setText ( NAME );
		window.setSize ( WIDTH, HEIGHT );
		window.setLocation ( x, y );
		window.open ();

		return window;
	}

	/**
	 * Arranges the buttons in a grid. Two buttons in the first row, one button
	 * in the second row.
	 * 
	 * @param composite
	 *        The container for the buttons.
	 * @param b1
	 *        The button at top left.
	 * @param b2
	 *        The button at top right.
	 * @param b3
	 *        The button at bottom.
	 */
	public static void adjustInGrid ( final Composite composite, final Button b1, final Button b2, final Button b3 ) {
		final GridLayout gridLayout = new GridLayout ( 2, true );

		gridLayout.marginTop = 0;
		gridLayout.marginRight = 0;
		gridLayout.marginBottom = 0;
		gridLayout.marginLeft = 0;
		gridLayout.marginWidth = 0;
		gridLayout.marginHeight = 0;
		gridLayout.horizontalSpacing = 0;
		gridLayout.verticalSpacing = 0;

		b1.setLayoutData ( new GridData ( GridData.FILL, GridData.FILL, true, true, 1, 1 ) );
		b2.setLayoutData ( new GridData ( GridData.FILL, GridData.FILL, true, true, 1, 1 ) );
		b3.setLayoutData ( new GridData ( GridData.FILL, GridData.FILL, true, true, 2, 1 ) );
		composite.setLayout ( gridLayout );
	}

	/**
	 * Creates a push button.
	 * 
	 * @param parent
	 *        The parent to add the buttons to.
	 * @param name
	 *        The name of the button.
	 */
	public static Button createButton ( final Composite parent, final String name ) {
		final Button button = new Button ( parent, SWT.PUSH );

		button.setText ( name );

		return button;
	}

}
