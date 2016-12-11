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


package geometry




type Rect3FloatArr struct {
	Pos Tuple3FloatArr
	Size Tuple2FloatArr
}

func (rect *Rect3FloatArr) Set(index int32, x, y, z, width, height float32) {
	rect.Pos.Att1Arr[index] = x
	rect.Pos.Att2Arr[index] = y
	rect.Pos.Att3Arr[index] = z
	rect.Size.Att1Arr[index] = width
	rect.Size.Att2Arr[index] = height
}

func (rect *Rect3FloatArr) X(index int32) float32 {
	return rect.Pos.Att1Arr[index]
}

func (rect *Rect3FloatArr) Y(index int32) float32 {
	return rect.Pos.Att2Arr[index]
}

func (rect *Rect3FloatArr) Z(index int32) float32 {
	return rect.Pos.Att3Arr[index]
}

func (rect *Rect3FloatArr) Width(index int32) float32 {
	return rect.Size.Att1Arr[index]
}

func (rect *Rect3FloatArr) Height(index int32) float32 {
	return rect.Size.Att2Arr[index]
}

func (rect *Rect3FloatArr) Area(index int32) float32 {
	return rect.Size.Att1Arr[index] * rect.Size.Att2Arr[index]
}

func (rect *Rect3FloatArr) Length() int32 {
	return int32(len(rect.Pos.Att1Arr))
}

func (rect *Rect3FloatArr) Capacity() int32 {
	return int32(cap(rect.Pos.Att1Arr))
}




