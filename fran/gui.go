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
	"unsafe"
	"github.com/vitalibaumtrok/glut"
	"github.com/vitalibaumtrok/gl10"
	"github.com/vitalibaumtrok/gl11"
)




const (
	gui_KEY_ESCAPE = 27
	gui_KEY_PLUS = 43
	gui_KEY_MINUS = 45
	gui_KEY_C = 67
	gui_KEY_H = 72
	gui_KEY_I = 73
	gui_KEY_O = 79
	gui_KEY_P = 80
	gui_KEY_R = 82
	gui_KEY_c = 99
	gui_KEY_h = 104
	gui_KEY_i = 105
	gui_KEY_o = 111
	gui_KEY_p = 112
	gui_KEY_r = 114
)

var (
	texture [1]uint
	m [16]float32
	windowWidth int
	windowHeight int
)

func guiInit() {
	glut.Init()
}

func guiStart() {
	m[0]=1;
	m[5]=1;
	m[10]=1;
	m[15]=1;

	frameInit()

	glut.InitDisplayMode(glut.DOUBLE | glut.RGBA)
	glut.InitWindowSize(frameWidth, frameHeight)
	glut.CreateWindow("fran 0.3.0");
	glut.ReshapeFunc(guiReshape)
	glut.DisplayFunc(guiDisplay)
	glut.KeyboardFunc(guiKeyboard)
	glut.SpecialFunc(guiSpecial)

	if 1 < argC {
		glut.TimerFunc(frameDelay, guiUpdateFrame, frameValue)
	}

	gl10.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl10.BlendFunc(gl10.SRC_ALPHA, gl10.ONE_MINUS_SRC_ALPHA)
	gl10.Enable(gl10.BLEND)
	gl10.Enable(gl10.TEXTURE_2D);

	gl11.GenTextures(1, texture[:] );
	gl11.BindTexture(gl10.TEXTURE_2D, texture[0]);
	gl10.TexImage2D(gl10.TEXTURE_2D, 0, gl10.RGBA, imgWidth, imgHeight, 0, gl10.RGBA, gl10.UNSIGNED_BYTE, unsafe.Pointer(&imgRGBA[0]));
	gl10.TexParameteri(gl10.TEXTURE_2D, gl10.TEXTURE_MAG_FILTER, gl10.NEAREST);
	gl10.TexParameteri(gl10.TEXTURE_2D, gl10.TEXTURE_MIN_FILTER, gl10.NEAREST);
	gl10.TexEnvf(gl10.TEXTURE_ENV, gl10.TEXTURE_ENV_MODE, gl10.REPLACE);

	glut.MainLoop()
}

func guiReshape(width, height int) {
	frameRatio := float32(frameWidth) / float32(frameHeight)
	ratio := float32(width) / float32(height);
	windowWidth = width
	windowHeight = height

	// broader new window
	if frameRatio < ratio {
		m[0] = frameRatio / ratio;
		m[5] = 1;

	// narrower new window
	} else if frameRatio > ratio {
		m[0] = 1;
		m[5] = ratio / frameRatio;

	} else {
		m[0] = 1;
		m[5] = 1;
	}

	gl10.Viewport(0, 0, width, height);
	gl10.LoadIdentity();
	gl10.MultMatrixf(m[:]);
}

func guiDisplay() {
	

	minX := imgFrameMinX[frameIndex]
	minY := imgFrameMinY[frameIndex]
	maxX := imgFrameMaxX[frameIndex]
	maxY := imgFrameMaxY[frameIndex]

	gl10.Clear(gl10.COLOR_BUFFER_BIT)
	gl10.Begin(gl10.QUADS)
		gl10.TexCoord2f(minX, minY); gl10.Vertex3f(-1, 1, 0)
		gl10.TexCoord2f(maxX, minY); gl10.Vertex3f(1, 1, 0)
		gl10.TexCoord2f(maxX, maxY); gl10.Vertex3f(1, -1, 0)
		gl10.TexCoord2f(minX, maxY); gl10.Vertex3f(-1, -1, 0)
	gl10.End()

	glut.SwapBuffers()
}

func guiKeyboard(key uint8, x, y int) {

	switch key {

	case gui_KEY_ESCAPE:
		glut.DestroyWindow(glut.GetWindow())

	case gui_KEY_C, gui_KEY_c, gui_KEY_H, gui_KEY_h, gui_KEY_I, gui_KEY_i:
		printControls()

	case gui_KEY_O, gui_KEY_o:
		delayToPrint := frameDelay != argDelay
		frameDelay = argDelay
		frameReplayDelay = argReplayDelay

		if delayToPrint {
			printFrameDelay()
		}
		glut.ReshapeWindow(frameWidth, frameHeight)
		glut.PostRedisplay()

	case gui_KEY_PLUS:
		glut.ReshapeWindow(windowWidth * 2, windowHeight * 2)
		glut.PostRedisplay()

	case gui_KEY_MINUS:
		glut.ReshapeWindow(windowWidth / 2, windowHeight / 2)
		glut.PostRedisplay()

	case gui_KEY_P, gui_KEY_p:
		guiTogglePauseAnimation()

	case gui_KEY_R, gui_KEY_r:
		if loadPNG() {
			frameReload()

			gl11.DeleteTextures(1, texture[:] );
			gl11.GenTextures(1, texture[:] );
			gl11.BindTexture(gl10.TEXTURE_2D, texture[0]);
			gl10.TexImage2D(gl10.TEXTURE_2D, 0, gl10.RGBA, imgWidth, imgHeight, 0, gl10.RGBA, gl10.UNSIGNED_BYTE, unsafe.Pointer(&imgRGBA[0]));
			gl10.TexParameteri(gl10.TEXTURE_2D, gl10.TEXTURE_MAG_FILTER, gl10.NEAREST);
			gl10.TexParameteri(gl10.TEXTURE_2D, gl10.TEXTURE_MIN_FILTER, gl10.NEAREST);
			gl10.TexEnvf(gl10.TEXTURE_ENV, gl10.TEXTURE_ENV_MODE, gl10.REPLACE);

			guiReshape(windowWidth, windowHeight)

		} else {
			printErrorCouldNotLoadFile()
		}
	}
}

func guiSpecial(key, x, y int) {
	switch key {

	case glut.KEY_LEFT:

		if !frameCancelCurrent() {

			if 0 < (glut.GetModifiers() & glut.ACTIVE_CTRL) {
				frameIndexShift(-10)

			} else {
				frameIndexShift(-1)
			}
			glut.PostRedisplay()
		}

	case glut.KEY_RIGHT:

		if !frameCancelCurrent() {

			if 0 < (glut.GetModifiers() & glut.ACTIVE_CTRL) {
				frameIndexShift(10)

			} else {
				frameIndexShift(1)
			}
			glut.PostRedisplay()
		}

	case glut.KEY_UP:

		if 0 < (glut.GetModifiers() & glut.ACTIVE_CTRL) {
			frameDelayShift(10)
			printFrameDelay()

		} else {
			frameDelayShift(1)
			printFrameDelay()
		}
		glut.PostRedisplay()

	case glut.KEY_DOWN:

		if 0 < (glut.GetModifiers() & glut.ACTIVE_CTRL) {
			frameDelayShift(-10)
			printFrameDelay()

		} else {
			frameDelayShift(-1)
			printFrameDelay()
		}
		glut.PostRedisplay()
	}
}

func guiUpdateFrame(value int) {
	if 10000 > value {
		running := true

		for i, c := range frameCanceled {
			if value == c {
				lengthNew := len(frameCanceled) - 1
				frameCanceled[i] = frameCanceled[lengthNew]
				frameCanceled = frameCanceled[:lengthNew]
				running = false
				break
			}
		}

		if running {
			frameIndex += 1

			if argC - 1 == frameIndex {
				glut.TimerFunc(frameReplayDelay, guiUpdateFrame, value)
				glut.PostRedisplay()

			} else if argC == frameIndex {
				frameIndex = 0
				glut.TimerFunc(frameDelay, guiUpdateFrame, value)
				glut.PostRedisplay()

			} else {
				glut.TimerFunc(frameDelay, guiUpdateFrame, value)
				glut.PostRedisplay()
			}
		}

	} else {
		value -= 10000
		running := true

		for i, c := range frameCanceled {
			if value == c {
				lengthNew := len(frameCanceled) - 1
				frameCanceled[i] = frameCanceled[lengthNew]
				frameCanceled = frameCanceled[:lengthNew]
				running = false
				break
			}
		}

		if running {
			if argC - 1 == frameIndex {
				glut.TimerFunc(frameReplayDelay, guiUpdateFrame, value)

			} else {
				glut.TimerFunc(frameDelay, guiUpdateFrame, value)
			}
		}
	}
}

func guiTogglePauseAnimation() {
	// continue
	if 0 > frameValue {
		frameValue = frameValueNext()
		guiUpdateFrame(frameValue + 10000)

	// pause
	} else {
		frameCanceled = append(frameCanceled, frameValue)
		frameValue = -1
	}
}


