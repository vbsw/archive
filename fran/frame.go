/*
 * This is free and unencumbered software released into the public domain.
 * 
 * Anyone is free to copy, modify, publish, use, compile, sell, or
 * distribute this software, either in source code form or as a compiled
 * binary, for any purpose, commercial or non-commercial, and by any
 * means.
 * 
 * In jurisdictions that recognize copyright laws, the author or authors
 * of this software dedicate any and all copyright interest in the
 * software to the public domain. We make this dedication for the benefit
 * of the public at large and to the detriment of our heirs and
 * successors. We intend this dedication to be an overt act of
 * relinquishment in perpetuity of all present and future rights to this
 * software under copyright law.
 * 
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 * IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
 * OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
 * ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
 * OTHER DEALINGS IN THE SOFTWARE.
 * 
 * For more information, please refer to <http://unlicense.org>
 */


package main




var (
	frameWidth int
	frameHeight int
	frameValue int
	frameIndex int
	frameCanceled []int
	frameDelay int
	frameReplayDelay int
)

func frameInit() {
	frameWidth = loadWidth / argW
	frameHeight = loadHeight / argH
	frameValue = 0
	frameIndex = 0
	frameCanceled = make([]int, 0, 7)
	frameDelay = argDelay
	frameReplayDelay = argReplayDelay

	for 120 > frameWidth || 120 > frameHeight {
		frameWidth *= 2
		frameHeight *= 2
	}
}

func frameReload() {
	frameWidth = loadWidth / argW
	frameHeight = loadHeight / argH

	for 120 > frameWidth || 120 > frameHeight {
		frameWidth *= 2
		frameHeight *= 2
	}
}

func frameIndexShift(count int) {
	frameIndex += count
	frameIndex %= argC

	if 0 > frameIndex {
		frameIndex += argC
	}
}

func frameDelayShift(count int) {
	equal := frameDelay == frameReplayDelay
	frameDelay += count

	if 0 > frameDelay {
		frameDelay = 0
	}

	if equal {
		frameReplayDelay = frameDelay
	}
}

func frameCancelCurrent() bool {
	if 0 <= frameValue {
		frameCanceled = append(frameCanceled, frameValue)
		frameValue = -1
		return true

	} else {
		return false
	}
}

func frameValueNext() int {
	counter := 0

	for i := 0; i < len(frameCanceled); i++ {
		if counter == frameCanceled[i] {
			counter += 1
			i = -1
		}
	}
	return counter
}

