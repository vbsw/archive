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




type Rect3Int struct {
	Pos Tuple3Int
	Size Tuple2Int
}

func (rect *Rect3Int) Set(x, y, z, width, height int32) {
	rect.Pos.Att1 = x
	rect.Pos.Att2 = y
	rect.Pos.Att3 = z
	rect.Size.Att1 = width
	rect.Size.Att2 = height
}

func (rect *Rect3Int) X() int32 {
	return rect.Pos.Att1
}

func (rect *Rect3Int) Y() int32 {
	return rect.Pos.Att2
}

func (rect *Rect3Int) Z() int32 {
	return rect.Pos.Att3
}

func (rect *Rect3Int) Width() int32 {
	return rect.Size.Att1
}

func (rect *Rect3Int) Height() int32 {
	return rect.Size.Att2
}

func (rect *Rect3Int) Area() int32 {
	return rect.Size.Att1 * rect.Size.Att2
}
 



