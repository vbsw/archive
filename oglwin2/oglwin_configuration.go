
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package oglwin

import "C"

type tConfiguration struct {
	doubleBufferC C.int
	msaaC C.int
	samplesC C.int
}

func (this *tConfiguration) init ( doubleBuffer bool, msaaSamples int ) *tConfiguration {
	if doubleBuffer {
		this.doubleBufferC = 1
	} else {
		this.doubleBufferC = 0
	}

	if msaaSamples == 2 || msaaSamples == 4 || msaaSamples == 8 || msaaSamples == 16 {
		this.msaaC = 1
		this.samplesC = C.int(msaaSamples)
	} else {
		this.msaaC = 0
		this.samplesC = 0
	}

	return this
}

