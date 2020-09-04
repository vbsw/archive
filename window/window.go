
//          Copyright 2017, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//     (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


// Functions to create a simple window with OpenGL context.
package window

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/vbsw/semver"
)

var windowDestroy bool
var eventQueue *tEventQueue
var windowWidth int
var windowHeight int
var initialWindowProperties *Properties

// The version is 0.3.0.
func Version() semver.Version {
	return semver.New(0, 3, 0)
}

// Returns the name of the event as string.
func EventName(eventCode int) string {
	switch eventCode {
	case CLOSE:
		return "CLOSE"
	case RESIZE:
		return "RESIZE"
	case REDRAW:
		return "REDRAW"
	case KEY_PRESS:
		return "KEY_PRESS"
	case KEY_AUTO_PRESS:
		return "KEY_AUTO_PRESS"
	case KEY_RELEASE:
		return "KEY_RELEASE"
	case BUTTON_PRESS:
		return "BUTTON_PRESS"
	case CURSOR_MOVE:
		return "CURSOR_MOVE"
	case WHEEL_SPIN:
		return "WHEEL_SPIN"
	case INITIALIZE:
		return "INITIALIZE"
	default:
		return "unknown"
	}
}

// Creates a resizable window with minimum width and height. If antiAlising
// is true, then the framebuffer uses 2x multisampling (MSAA), otherwise
// none.
func Create(properties *Properties, callback func(event *Event)) {
	err := glfw.Init()

	if err == nil {
		defer glfw.Terminate()

		eventQueue = newEventQueue()
		initialWindowProperties = properties.Clone()
		window, err := createWindow()

		if err == nil {
			windowWidth = initialWindowProperties.Width
			windowHeight = initialWindowProperties.Height
			windowDestroy = false

			window.MakeContextCurrent()
			glfw.SwapInterval(1)
			mainLoop(window, callback)
			window.Destroy()

		} else {
			panic(err)
		}

	} else {
		panic(err)
	}
}

// Closes the created window.
func Destroy() {
	windowDestroy = true
}

func createWindow() (*glfw.Window, error) {
	glfw.DefaultWindowHints()

	if initialWindowProperties.AntiAliasing {
		glfw.WindowHint(glfw.Samples, 2)
	}
	if initialWindowProperties.Title == "" {
		initialWindowProperties.Title = "OpenGL Application"
	}
	if initialWindowProperties.Fullscreen {
		glfw.WindowHint(glfw.Maximized, 1)
		glfw.WindowHint(glfw.Decorated, 0)

		windowNew, err := glfw.CreateWindow(initialWindowProperties.Width, initialWindowProperties.Height, initialWindowProperties.Title, nil, nil)

		if err == nil {
			windowNew.SetKeyCallback(keyCallback)
			windowNew.SetCursorPosCallback(cursorCallback)
			windowNew.SetMouseButtonCallback(buttonCallback)
			windowNew.SetScrollCallback(scrollCallback)
			windowNew.MakeContextCurrent()
			windowWidth, windowHeight = windowNew.GetSize()
		}
		return windowNew, err

	} else {
		windowNew, err := glfw.CreateWindow(initialWindowProperties.Width, initialWindowProperties.Height, initialWindowProperties.Title, nil, nil)

		if err == nil {
			windowWidth = initialWindowProperties.Width
			windowHeight = initialWindowProperties.Height
			windowNew.SetSizeLimits(initialWindowProperties.Width, initialWindowProperties.Height, 10000, 10000)
			windowNew.SetKeyCallback(keyCallback)
			windowNew.SetCursorPosCallback(cursorCallback)
			windowNew.SetMouseButtonCallback(buttonCallback)
			windowNew.SetScrollCallback(scrollCallback)
			windowNew.SetFramebufferSizeCallback(resizeCallback)
			eventQueue.insertEventFront(NewEvent(RESIZE))
		}
		return windowNew, err
	}
}

func mainLoop(window *glfw.Window, callback func(event *Event)) {
	for !windowDestroy {
		pollNextEvent(window)
		callback(eventQueue.events[0])

		if eventQueue.events[0].EventCode == REDRAW {
			window.SwapBuffers()
		}
		eventQueue.removeEventFront()
	}
}

func pollNextEvent(window *glfw.Window) {
	if eventQueue.isEmpty() {
		glfw.PollEvents()
		eventQueue.appendRedrawEventIfMissing()
	}
	if window.ShouldClose() {
		eventQueue.insertEventFront(NewEvent(CLOSE))
		window.SetShouldClose(false)
	}
	cursorfX, cursorfY := window.GetCursorPos()
	eventQueue.events[0].CursorX = int(cursorfX)
	eventQueue.events[0].CursorY = int(cursorfY)
	eventQueue.events[0].WindowWidth = windowWidth
	eventQueue.events[0].WindowHeight = windowHeight

	time := glfw.GetTime()
	eventQueue.updateEventTime(time)

	if !eventQueue.isTimeInitialized() {
		eventQueue.removeEventFront()
		pollNextEvent(window)
	}
}

func keyCallback(window *glfw.Window, key glfw.Key, scanCode int, action glfw.Action, mods glfw.ModifierKey) {
	event := new(Event)
	event.KeyCode = scanCode

	switch action {
	case glfw.Press:
		event.EventCode = KEY_PRESS

	case glfw.Repeat:
		event.EventCode = KEY_AUTO_PRESS

	case glfw.Release:
		event.EventCode = KEY_RELEASE
	}
	eventQueue.appendEvent(event)
}

func cursorCallback(window *glfw.Window, xpos float64, ypos float64) {
	event := NewEvent(CURSOR_MOVE)

	eventQueue.appendEvent(event)
}

func buttonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	event := new(Event)
	event.ButtonCode = int(button)

	switch action {
	case glfw.Press:
		event.EventCode = BUTTON_PRESS

	case glfw.Release:
		event.EventCode = BUTTON_RELEASE
	}
	eventQueue.appendEvent(event)
}

func scrollCallback(w *glfw.Window, xoff float64, yoff float64) {
	event := new(Event)
	event.EventCode = WHEEL_SPIN
	event.Wheel = int(yoff)

	eventQueue.appendEvent(event)
}

func resizeCallback(w *glfw.Window, width int, height int) {
	windowWidth = width
	windowHeight = height

	eventQueue.replaceEvent(RESIZE)
}
