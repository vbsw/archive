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

// Functions to create graphical window with OpenGL context.
package oglwin

// #cgo LDFLAGS: -lX11 -lGL
// #include "oglwin.h"
import "C"
import (
	"errors"
)

var (
	err error = nil
	windowTitle string = ""
	event C.int
	event_value1 C.int
	event_value2 C.int
)

func setError(error_id C.int) {
	if (err!=nil) {
		switch error_id {
			case C.ERROR_NO_DISPLAY:
			err = errors.New("failed to establish connection to X server")
			case C.ERROR_NO_VISUAL_INFO:
			err = errors.New("failed to obtain visual info for the screen")
		}
	}
}

func Open() {
	C.open_window()
	if (C.error_id()!=C.ERROR_NONE) {
		setError(C.error_id())
	}
}

func Close() {
	C.close_window()
	if (C.error_id()!=C.ERROR_NONE) {
		setError(C.error_id())
	}
}

func Size() (int,int) {
	var width, height C.int
	C.window_size(&width,&height)
	return int(width), int(height)
}

func SetSize(width, height int) {
	C.window_set_size(C.int(width),C.int(height))
}

func Title() string {
	return windowTitle
}

func SetTitle(title string) {
	windowTitleC := C.CString(title)
	windowTitle = title

	C.window_set_title(windowTitleC)
}

func NextEvent() bool {
	window_exists := C.connection_exists()

	if (window_exists!=0) {
		C.next_event(&event,&event_value1,&event_value2)
		return true

	} else {
		return false
	}
}

func RequestWindowClosing() {
	C.window_request_close()
}

func IsKeyPressedEvent() bool {
	return event==C.EVENT_KEY_PRESSED
}

func IsRedrawEvent() bool {
	return event==C.EVENT_WINDOW_REDRAW
}

func IsWindowClosingEvent() bool {
	return event==C.EVENT_WINDOW_CLOSING
}

func KeyCode() int {
	return int(event_value1)
}

func Error() error {
	return err
}

