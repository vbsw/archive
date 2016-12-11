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




type Tuple4FloatArr struct {
	Att1Arr []float32
	Att2Arr []float32
	Att3Arr []float32
	Att4Arr []float32
}

func (tuple *Tuple4FloatArr) Init(length, capacity int32) {
	tuple.Att1Arr = make([]float32, length, capacity)
	tuple.Att2Arr = make([]float32, length, capacity)
	tuple.Att3Arr = make([]float32, length, capacity)
	tuple.Att4Arr = make([]float32, length, capacity)
}

func (tuple *Tuple4FloatArr) Set(index int32, att1, att2, att3, att4 float32) {
	tuple.Att1Arr[index] = att1
	tuple.Att2Arr[index] = att2
	tuple.Att3Arr[index] = att3
	tuple.Att4Arr[index] = att4
}

func (tuple *Tuple4FloatArr) Append(att1, att2, att3, att4 float32) {
	tuple.Att1Arr = append(tuple.Att1Arr, att1)
	tuple.Att2Arr = append(tuple.Att2Arr, att2)
	tuple.Att3Arr = append(tuple.Att3Arr, att3)
	tuple.Att4Arr = append(tuple.Att4Arr, att4)
}

func (tuple *Tuple4FloatArr) ExpandCapacity(capacity int32) {
	att1Arr := make([]float32, len(tuple.Att1Arr), capacity)
	copy(att1Arr, tuple.Att1Arr)
	tuple.Att1Arr = att1Arr

	att2Arr := make([]float32, len(tuple.Att2Arr), capacity)
	copy(att2Arr, tuple.Att2Arr)
	tuple.Att2Arr = att2Arr

	att3Arr := make([]float32, len(tuple.Att3Arr), capacity)
	copy(att3Arr, tuple.Att3Arr)
	tuple.Att3Arr = att3Arr

	att4Arr := make([]float32, len(tuple.Att4Arr), capacity)
	copy(att4Arr, tuple.Att4Arr)
	tuple.Att4Arr = att4Arr
}

func (tuple *Tuple4FloatArr) Length() int32 {
	return int32(len(tuple.Att1Arr))
}

func (tuple *Tuple4FloatArr) Capacity() int32 {
	return int32(cap(tuple.Att1Arr))
}




