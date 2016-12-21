
//          Copyright 2016 Vitali Baumtrok
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package main

import (
	"github.com/vbsw/semver"
)

var app App

type App struct {
	version semver.Version
	clinfo *tCLInfo
	settings *tSettings
	window *tWindow
	graphics *tGraphics
	scene *tScene
	engine *tEngine
}

func main ( ) {
	app.version = semver.New(0, 1, 0)
	app.clinfo = newCLInfo()

	if app.clinfo.None() {
		app.settings = newSettings()
		app.window = newWindow()
		app.graphics = newGraphics()
		app.scene = newScene()
		app.engine = newEngine()

		app.window.Show()
		app.engine.Start()
		app.window.Destroy()

	} else {
		app.clinfo.Print()
	}
}
