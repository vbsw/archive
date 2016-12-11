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

#ifndef OGLWIN_H
#define OGLWIN_H

#define ERROR_NONE 0
#define ERROR_NO_DISPLAY 1
#define ERROR_NO_VISUAL_INFO 2

#define EVENT_NONE 0
#define EVENT_KEY_PRESSED 1
#define EVENT_KEY_RELEASED 2
#define EVENT_WINDOW_REDRAW 3
#define EVENT_WINDOW_CLOSING 4

extern void
open_window ( void );
extern void
close_window ( void );
extern void
window_size ( int *const width, int *const height );
extern void
window_set_size ( const int width, const int height );
extern void
window_set_title ( const char* str );
extern int
connection_exists ( void );
extern void
next_event ( int* event, int* value1, int* value2 );
extern void
window_request_close ( void );
extern int
error_id ( void );

#endif /* OGLWIN_H */
