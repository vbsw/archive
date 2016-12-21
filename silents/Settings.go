
//          Copyright 2016 Vitali Baumtrok
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package main

import (
)

type tSettings struct {
}

func newSettings ( ) *tSettings {
	settings := new(tSettings)
	return settings
}

func (this *tSettings) WindowWidth ( ) uint {
	return 360
}

func (this *tSettings) WindowHeight ( ) uint {
	return 640
}

