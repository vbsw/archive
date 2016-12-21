
//          Copyright 2016 Vitali Baumtrok
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package main

import (
	"fmt"
)

type tEngine struct {
}

func newEngine ( ) *tEngine {
	engine := new(tEngine)
	return engine
}

func (this *tEngine) Start ( ) {
	for app.window.NextEvent() {
		event := app.window.Event()
		if event.KeyPressed() {
			fmt.Println("key pressed")
		}
	}
}
