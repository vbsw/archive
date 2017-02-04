
//          Copyright 2016 Vitali Baumtrok
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


// Prints the name of the default display server, if it is running, otherwise "n.a." (not available).
package main


import (
	"fmt"
	"strconv"
	"github.com/vbsw/semver"
	"github.com/vbsw/xlib"
)


var version semver.Version

func main() {
	version = semver.New(0,1,0)
	clinfo := newCLInfo()

	if clinfo.None() {
		display := xlib.XOpenDisplay()

		if ( display != nil ) {

			if clinfo.Screens() {
				displayScreenCount := xlib.XScreenCount(display)
				if displayScreenCount > 0 {
					resolutions := screensResolutions(display,displayScreenCount)
					fmt.Println(displayScreenCount,resolutions)
				} else {
					fmt.Println(displayScreenCount)
				}

			} else {
				displayName := xlib.XDisplayString(display)
				fmt.Println(displayName)
			}
			xlib.XCloseDisplay(display)

		} else {
			fmt.Println("n.a.")
		}

	} else {
		clinfo.Print()
	}
}

func screensResolutions ( display *xlib.Display, screenCount int ) string {
	resolutions := ""
	for i:=0; i<screenCount; i+=1 {
		screen := xlib.XScreenOfDisplay(display,i)
		screenWidth := xlib.XWidthOfScreen(screen)
		screenHeight := xlib.XHeightOfScreen(screen)
		if len(resolutions) == 0 {
			resolutions = strconv.Itoa(screenWidth) + "x" + strconv.Itoa(screenHeight)
		} else {
			resolutions = resolutions + " " + strconv.Itoa(screenWidth) + "x" + strconv.Itoa(screenHeight)
		}
	}
	return resolutions
}

