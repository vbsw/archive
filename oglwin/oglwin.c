/**
 * oglwin (graphical window with OpenGL context)
 * Copyright (C) 2016 Vitali Baumtrok
 *
 * This file is part of oglwin.
 *
 * oglwin is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * oglwin is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with oglwin.  If not, see <http://www.gnu.org/licenses/>.
 */

#include <stdlib.h>
#include <stdio.h>
#include <X11/Xlib.h>
#include <GL/glx.h>

#define ERROR_NONE 0
#define ERROR_NO_DISPLAY 1
#define ERROR_NO_VISUAL_INFO 2

#define EVENT_NONE 0
#define EVENT_KEY_PRESSED 1
#define EVENT_KEY_RELEASED 2
#define EVENT_WINDOW_REDRAW 3
#define EVENT_WINDOW_CLOSING 4

static int error = ERROR_NONE;
static Display* display = NULL;
static int screen_id;
static const char* title = 0;
static Window window;
static Atom wm_protocols;
static Atom wm_delete_window_atom;
static int window_width = 720;
static int window_height = 480;
static int window_event = EVENT_NONE;
static int window_event_value1;
static int window_event_value2;
static Time window_event_time;

static GLint visual_attributes[5] = { GLX_RGBA, GLX_DEPTH_SIZE, 24, GLX_DOUBLEBUFFER, None };

static void
create_display ( void ) {
	if (display==NULL && error==0) {
		display = XOpenDisplay(NULL);

		if (display==NULL && error==0) {
			error = ERROR_NO_DISPLAY;
		}
	}
}

static void
create_window ( void ) {
	if (display!=NULL && error==ERROR_NONE) {
		screen_id = DefaultScreen(display);
		Screen *const screen = ScreenOfDisplay(display,screen_id);
		Window root = RootWindowOfScreen(screen);
		XVisualInfo *const visual_info = glXChooseVisual(display,screen_id,visual_attributes);

		if (visual_info!=NULL) {
			Colormap colormap = XCreateColormap(display,root,visual_info->visual,AllocNone);
			XSetWindowAttributes set_window_attributes;
			set_window_attributes.colormap = colormap;
			set_window_attributes.event_mask = ExposureMask|KeyPressMask|KeyReleaseMask|SubstructureNotifyMask;
			window = XCreateWindow(display,root,0,0,window_width,window_height,0,visual_info->depth,InputOutput,visual_info->visual,CWColormap|CWEventMask,&set_window_attributes);
			wm_protocols = XInternAtom(display,"WM_PROTOCOLS",False);
			wm_delete_window_atom = XInternAtom(display,"WM_DELETE_WINDOW",False);

			XSetWMProtocols(display,window,&wm_delete_window_atom,1);
			XMapWindow(display,window);

			if (title!=0) {
				XStoreName(display,window,title);
			}

			XFree(visual_info);

		} else {
			XCloseDisplay(display);
			display = NULL;
			error = ERROR_NO_VISUAL_INFO;
		}
	}
}

static void 
get_window_size ( void ) {
	const int screen = XDefaultScreen(display);
	Window root;
	int x, y, border_width, depth;

	XGetGeometry(display,window,&root,&x,&y,&window_width,&window_height,&border_width,&depth);
}

void
open_window ( void ) {
	create_display();
	create_window();
}

void
close_window ( void ) {
	if (display!=NULL) {
		XCloseDisplay(display);
		display = NULL;
	}
}

void
window_size ( int *const width, int *const height ) {
	width[0] = window_width;
	height[0] = window_height;
}

void
window_set_size ( const int width, const int height ) {
	window_width = width;
	window_height = height;

	if ( (display!=NULL) && (error==ERROR_NONE) ) {
		XResizeWindow(display,window,width,height);
	}
}

void
window_set_title ( const char* str )
{
	if (title!=0) {
		free((void*)title);
	}
	title = str;

	if ( (display!=NULL) && (error==ERROR_NONE) ) {
		XStoreName(display,window,title);
	}
}

int
connection_exists ( void ) {
	return display!=NULL;
}

void
next_event ( int* event, int* value1, int* value2 ) {
	if (window_event==EVENT_NONE) {
		XEvent xevent;
		int has_event = 0;
		int event_tmp = EVENT_NONE;
		int value1_tmp, value2_tmp;
		Time time_tmp1, time_tmp2;

		while (has_event==0) {
	 		XNextEvent(display,&xevent);

			if (xevent.type==KeyPress) {
				event_tmp  = EVENT_KEY_PRESSED;
				value1_tmp = (int) xevent.xkey.keycode;
				value2_tmp = 0;
				has_event  = 1;

			} else if (xevent.type== KeyRelease) {
				event_tmp  = EVENT_KEY_RELEASED;
				value1_tmp = (int) xevent.xkey.keycode;
				value2_tmp = 0;
				has_event  = 1;
				time_tmp1  = xevent.xkey.time;

			} else if ( (xevent.type==Expose) && (xevent.xexpose.count==0) ) {
				get_window_size();
				event_tmp  = EVENT_WINDOW_REDRAW;
				value1_tmp = window_width;
				value2_tmp = window_height;
				has_event  = 1;

			} else if (     xevent.type                 == ClientMessage
				     && xevent.xclient.message_type == wm_protocols
				     && xevent.xclient.data.l[0]    == wm_delete_window_atom ) {
				event_tmp  = EVENT_WINDOW_CLOSING;
				value1_tmp = 0;
				value2_tmp = 0;
				has_event  = 1;
			}
		}

		if (event_tmp==EVENT_KEY_RELEASED) {
			has_event = 0;

			while ( (has_event==0) && (XPending(display)) ) {
		 		XNextEvent(display,&xevent);

				if (xevent.type==KeyPress) {
					window_event        = EVENT_KEY_PRESSED;
					window_event_value1 = (int) xevent.xkey.keycode;
					window_event_value2 = 0;
					has_event           = 1;
					time_tmp2           = xevent.xkey.time;

				} else if (xevent.type== KeyRelease) {
					event_tmp           = EVENT_KEY_RELEASED;
					window_event_value1 = (int) xevent.xkey.keycode;
					window_event_value2 = 0;
					has_event           = 1;

				} else if ( (xevent.type==Expose) && (xevent.xexpose.count==0) ) {
					get_window_size();
					window_event        = EVENT_WINDOW_REDRAW;
					window_event_value1 = window_width;
					window_event_value2 = window_height;
					has_event           = 1;

				} else if (     xevent.type                 == ClientMessage
					     && xevent.xclient.message_type == wm_protocols
					     && xevent.xclient.data.l[0]    == wm_delete_window_atom ) {
					window_event        = EVENT_WINDOW_CLOSING;
					window_event_value1 = 0;
					window_event_value2 = 0;
					has_event           = 1;
				}
			}
		}

		if (           event_tmp    == EVENT_KEY_RELEASED
		            && value1_tmp   == window_event_value1
		            && window_event == EVENT_KEY_PRESSED
		            && time_tmp1    == time_tmp2 ) {
			event[0]     = window_event;
			value1[0]    = window_event_value1;
			value2[0]    = window_event_value2;
			window_event = EVENT_NONE;

		} else {
			event[0]     = event_tmp;
			value1[0]    = value1_tmp;
			value2[0]    = value2_tmp;
		}

	} else {
		event[0]     = window_event;
		value1[0]    = window_event_value1;
		value2[0]    = window_event_value2;
		window_event = EVENT_NONE;
	}
}

void
window_request_close ( void ) {
	XEvent close_event;
	close_event.xclient.type = ClientMessage;
	close_event.xclient.message_type = wm_protocols;
	close_event.xclient.format = 32;
	close_event.xclient.data.l[0] = wm_delete_window_atom;
	close_event.xclient.data.l[1] = CurrentTime;

	XSendEvent(display,window,False,NoEventMask,&close_event);
}

int
error_id ( void ) {
	return error;
}


