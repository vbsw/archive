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

import org.eclipse.swt.widgets.Display;
import org.eclipse.swt.widgets.Shell;

/**
 * A class to create a simple window of size 480x360 pixel.
 * 
 * @author Vitali Baumtrok
 */
public class XmplSimpleWindow {

	/**
	 * The name of the window.
	 */
	private static final String	NAME	= "SWT Tutorial: Simple Window";

	/**
	 * The width of the window.
	 */
	private static final int	WIDTH	= 480;

	/**
	 * The height of the window.
	 */
	private static final int	HEIGHT	= 360;

	/**
	 * Creates a window of size 480x360 pixel placed in the middle of the
	 * given display.
	 * 
	 * @param display
	 *        The display to create the window on.
	 * @return A window.
	 */
	public static Shell createAndOpenWindow ( final Display display ) {
		final int x = ( display.getBounds ().width - WIDTH ) / 2;
		final int y = ( display.getBounds ().height - HEIGHT ) / 2;
		final Shell window = new Shell ( display );

		window.setText ( NAME );
		window.setSize ( WIDTH, HEIGHT );
		window.setLocation ( x, y );
		window.open ();

		return window;
	}

}
