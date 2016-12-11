
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)

#ifndef WINDOW_H
#define WINDOW_H

#define WINDOW_EVENT_NONE 0
#define WINDOW_EVENT_KEY_PRESS 1
#define WINDOW_EVENT_KEY_RELEASE 2
#define WINDOW_EVENT_WINDOW_CLOSE 3
#define WINDOW_EVENT_REDRAW 4

extern void window_destroy ( void );
extern int  window_events_has_next ( void );
extern void window_events_next ( int* event_type, int* value1, int* value2 );
extern void window_init ( const char* title, int width, int height );
extern void window_reshape ( void );
extern void window_show ( void );
extern void window_toggle_fullscreen ( void );

#endif /* WINDOW_H */

