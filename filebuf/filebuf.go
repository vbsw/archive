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


package filebuf

import (
	"os"
	"bytes"
)




var (
	Data []byte = nil
	Factor float32 = 1.2
	Name string
	InitialSize = 4000
	errorStr string
)

func Error() string {
	return errorStr
}

func Read() bool {
	file, errFile := os.Open(Name)
	errorStr = ""

	if errFile == nil {
		defer file.Close()

		stat, errStat := file.Stat()

		if errStat == nil {
			loadSize := int(stat.Size())

			if Data == nil {

				if InitialSize > loadSize {
					Data = make([]byte, InitialSize)

				} else {
					Data = make([]byte, int(float32(loadSize) * Factor))
				}

			} else {

				if cap(Data)  > loadSize {
					Data = Data[:cap(Data)]

				} else {
					Data = make([]byte, int(float32(loadSize) * Factor))
				}
			}
			n, errFileRead := file.Read(Data)
			Data = Data[:n]

			if errFileRead == nil {
				return true

			} else {
				errorStr = errFileRead.Error()

				return false
			}

		} else {
			errorStr = errStat.Error()

			return false
		}

	} else {
		errorStr = errFile.Error()

		return false
	}
}

func ReadFromFile(name string) bool {
	Name = name

	return Read()
}

func DataReader() *bytes.Reader {
	return bytes.NewReader(Data)
}



