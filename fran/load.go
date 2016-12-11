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
	"os"
	"bytes"
	"image"
	"image/png"
)




var (
	loadSize int
	loadBuffer []byte
	loadRGBA []byte
	loadWidth int
	loadHeight int
)

func loadPNG() bool {
	if loadFile() {
		img, imgErr := png.Decode(bytes.NewReader(loadBuffer))

		if imgErr == nil {
			switch imgGen := img.(type) {
				case *image.NRGBA:
				loadRGBA = imgGen.Pix
				loadWidth = imgGen.Rect.Max.X - imgGen.Rect.Min.X
				loadHeight = imgGen.Rect.Max.Y - imgGen.Rect.Min.Y
				imgConvert()
				return true

				case *image.RGBA:
				loadRGBA = imgGen.Pix
				loadWidth = imgGen.Rect.Max.X - imgGen.Rect.Min.X
				loadHeight = imgGen.Rect.Max.Y - imgGen.Rect.Min.Y
				imgConvert()
				return true

				case *image.NRGBA64:
				loadRGBA = imgGen.Pix
				loadWidth = imgGen.Rect.Max.X - imgGen.Rect.Min.X
				loadHeight = imgGen.Rect.Max.Y - imgGen.Rect.Min.Y
				imgConvert()
				return true

				case *image.RGBA64:
				loadRGBA = imgGen.Pix
				loadWidth = imgGen.Rect.Max.X - imgGen.Rect.Min.X
				loadHeight = imgGen.Rect.Max.Y - imgGen.Rect.Min.Y
				imgConvert()
				return true

				default:
				printCanNotInterpretImage()
			}
		}
	}
	return false
}

func loadFile() bool {
	file, errFile := os.Open(argFile)

	if errFile == nil {
		defer file.Close()

		stat, errStat := file.Stat()

		if errStat == nil {
			loadSize = int(stat.Size())

			if cap(loadBuffer)  < int(loadSize) {
				loadBuffer = make([]byte, int(float32(loadSize) * 1.2))

			} else {
				loadBuffer = loadBuffer[:cap(loadBuffer)]
			}
			n, errFileRead := file.Read(loadBuffer)
			loadBuffer = loadBuffer[:n]

			return errFileRead == nil
		}
	}
	return false
}




