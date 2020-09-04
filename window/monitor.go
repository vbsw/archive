
//          Copyright 2017, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//     (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package window

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

func currentMonitor(window *glfw.Window) *glfw.Monitor {
	var monitorMatching *glfw.Monitor
	var overlapAreaHighest int

	monitors := glfw.GetMonitors()

	for _, monitor := range monitors {
		overlapArea := overlapAreaOfWindowAndMonitor(window, monitor)

		if overlapAreaHighest < overlapArea {
			overlapAreaHighest = overlapArea
			monitorMatching = monitor
		}
	}
	return monitorMatching
}

func max(val1, val2 int) int {
	if val1 < val2 {
		return val2

	} else {
		return val1
	}
}

func min(val1, val2 int) int {
	if val1 < val2 {
		return val1

	} else {
		return val2
	}
}

func overlapAreaOfWindowAndMonitor(window *glfw.Window, monitor *glfw.Monitor) int {
	windowWidth, windowHeight := window.GetSize()
	windowX, windowY := window.GetPos()
	videoMode := monitor.GetVideoMode()
	monitorWidth, monitorHeight := videoMode.Width, videoMode.Height
	monitorX, monitorY := monitor.GetPos()

	minRightX := min(windowX+windowWidth, monitorX+monitorWidth)
	minLowY := min(windowY+windowHeight, monitorY+monitorHeight)
	sideA := max(0, minRightX-max(windowX, monitorX))
	sideB := max(0, minLowY-max(windowY, monitorY))
	area := sideA * sideB

	return area
}
