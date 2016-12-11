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

import (
)




var (
	imgWidth int
	imgHeight int
	imgRGBA []byte
	imgFrameMinX []float32
	imgFrameMinY []float32
	imgFrameMaxX []float32
	imgFrameMaxY []float32
)

func imgConvert() {
	imgWidth = twoPotency(loadWidth)
	imgHeight = twoPotency(loadHeight)
	imgSize := int(imgWidth * imgHeight) * 4

	if cap(imgRGBA) < imgSize  {
		imgRGBA = make([]byte, imgSize, int(float32(imgSize) * 1.2))

	} else {
		imgRGBA = imgRGBA[:imgSize]
	}

	if cap(imgFrameMinX) < argC {
		imgFrameMinX = make([]float32, argC)
		imgFrameMinY = make([]float32, argC)
		imgFrameMaxX = make([]float32, argC)
		imgFrameMaxY = make([]float32, argC)
	}

	imgSetFrames()
	imgCopyRGBA()
}

func imgSetFrames() {
	frameWidth := loadWidth / argW
	frameHeight := loadHeight / argH
	frameWidthFloat := float32(frameWidth) / float32(imgWidth)
	frameHeightFloat := float32(frameHeight) / float32(imgHeight)
	frameHeightFloatMin := float32(0)

	for h, i := 0, 0; h < argH && i < argC; h++ {
		frameWidthFloatMin := float32(0)

		for w := 0; w < argW && i < argC; w++ {
			imgFrameMinX[i] = frameWidthFloatMin
			imgFrameMaxX[i] = frameWidthFloatMin + frameWidthFloat
			imgFrameMinY[i] = frameHeightFloatMin
			imgFrameMaxY[i] = frameHeightFloatMin + frameHeightFloat
			frameWidthFloatMin += frameWidthFloat
			i += 1
		}
		frameHeightFloatMin += frameHeightFloat
	}
}

func imgCopyRGBA() {
	dstLineLength := imgWidth * 4
	srcLineLength := loadWidth * 4
	for line := 0; line < loadHeight; line++ {
		loadOffset := line * srcLineLength
		copy(imgRGBA[line*dstLineLength:], loadRGBA[loadOffset:loadOffset+srcLineLength])
	}
}

func twoPotency(val int) int {
	var nextVal int = 1

	for nextVal < val {
		nextVal = nextVal * 2
	}
	return nextVal
}




