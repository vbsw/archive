/**
 * DisplayInfo (info on display)
 * Copyright (C) 2017 Vitali Baumtrok
 *
 * This file is part of DisplayInfo.
 *
 * DisplayInfo is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * DisplayInfo is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with DisplayInfo.  If not, see <http://www.gnu.org/licenses/>.
 */


#include <X11/Xlib.h>

extern char* util_copy_string(const char *const string);


/* ************************* */
/*    exported functions     */
/* ************************* */

const char* xserver_default_display_server() {
	Display* display = XOpenDisplay(0);

	if (display!=0) {
		const char *const server = DisplayString(display);
		const char *const server_copy = util_copy_string(server);
		XCloseDisplay(display);
		return server_copy;

	} else {
		return 0;
	}
}

void xserver_resolution(int *const width, int *const height) {
	Display* display = XOpenDisplay(0);

	if (display!=0) {
		Window window;
		int revert_to_return;
		XWindowAttributes window_attributes;
		XGetInputFocus(display,&window,&revert_to_return);
		XGetWindowAttributes(display,window,&window_attributes);

		width[0] = window_attributes.screen[0].width;
		height[0] = window_attributes.screen[0].height;

		XCloseDisplay(display);

	} else {
		width[0] = -1;
		height[0] = -1;
	}
}

