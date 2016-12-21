
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package main

import (
	"github.com/vbsw/semver"
)

var version semver.Version

func main ( ) {
	version = semver.New(0,1,0)
	infoRequest := newInfoRequest()

	if infoRequest.None() {
		game := newGame()
		game.Init()
		game.Loop()

	} else {
		infoRequest.Print()
	}
}

func mainGameTitle ( ) string {
	return "Silent Shooter (" + version.String() + ")"
}

