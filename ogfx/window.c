
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


#include <stdlib.h>
#include <string.h>
#include "../cds/iqueue.h"

#define WINDOW_EVENT_NONE 0
#define WINDOW_EVENT_KEY_PRESS 1
#define WINDOW_EVENT_KEY_RELEASE 2
#define WINDOW_EVENT_WINDOW_CLOSE 3
#define WINDOW_EVENT_REDRAW 4

void window_destroy ( void ) {

}

int window_events_has_next ( void ) {
	return 1;
}

void window_events_next ( int *const event_type, int *const value1, int *const value2 ) {
	event_type[0] = 0;

	if ( value1 != 0 ) {
		value1[0] = 0;
	}
	if ( value2 != 0 ) {
		value2[0] = 0;
	}
}

void window_init ( const char *const title, int width, int height ) {

}


void window_reshape ( void ) {

}

void window_show ( void ) {

}

void window_toggle_fullscreen ( void ) {

}

