
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package main

import (
	"fmt"
	"github.com/vbsw/vis"
	"github.com/vbsw/visimpl"
	"github.com/vbsw/dclock"
)

type tGame struct {
	vsys vis.VisualSystem
	clock dclock.Clock
}

func newGame ( ) *tGame {
	game := new(tGame)
	game.clock = dclock.New()
	return game
}

func (this *tGame) Init ( ) {
	vsys, err := visimpl.New()

	if err == nil {
		this.vsys = vsys

		vsys.Graphics().BackgroundColorFill().SetRGB(0, 0, 0)
		vsys.EventQueue().SetMouseEvents(true)
		vsys.Window().SetWidthHeightInPixel(360, 540)
		vsys.Window().SetTitle(mainGameTitle())
		vsys.Window().Show()

	} else {
		this.vsys = nil
		fmt.Println("Error:", err)
	}

	
}

func (this *tGame) Loop ( ) {
	vsys := this.vsys
	clock := this.clock

	for vsys.EventQueue().PollEvents() {
		clock.Update()

		switch vsys.EventQueue().EventCode() {

			case vis.KEY_PRESSED:
			if vsys.EventQueue().KeyCode() == 9 {
				vsys.Destroy();
			} else {
				fmt.Println("key pressed:", vsys.EventQueue().KeyCode())
			}

			case vis.CURSOR:
			fmt.Println("mouse ", vsys.EventQueue().X(), vsys.EventQueue().Y())

			case vis.REDRAW:
			fmt.Println("resize ", vsys.Window().WidthInPixel(), "x", vsys.Window().HeightInPixel())
			vsys.Graphics().Redraw()

			case vis.CLOSE:
			vsys.Destroy();
		}
	}
}
