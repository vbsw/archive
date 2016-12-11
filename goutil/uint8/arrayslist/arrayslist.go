/* 
 * util - a library for the programming language Go
 * Copyright 2014 Vitali Baumtrok
 * 
 * This file is part of util.
 * 
 * tag is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * 
 * tag is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * 
 * You should have received a copy of the GNU General Public License
 * along with tag.  If not, see <http://www.gnu.org/licenses/>.
 * 
 */

package arrayslist

import (
	goBytes "bytes"
	caparr8 "github.com/vitalibaumtrok/util/uint8/caparr"
	caparr64 "github.com/vitalibaumtrok/util/uint64/caparr"
	"github.com/vitalibaumtrok/util/uint64/decode"
	"github.com/vitalibaumtrok/util/uint64/encode"
)

const (
	MINIMUM_INITIAL_BYTES_SIZE = 8
)

type Type struct {
	indices []uint64
	arrays []uint8
}

func New() (t *Type) {
	t = new(Type)
	t.Init()
	return t
}

func NewWith(bytes []uint8) (t *Type) {
	t = new(Type)
	t.InitWith(bytes)
	return t
}

func (t *Type) Init() {
	t.indices = caparr64.New(0)
	t.arrays = caparr8.New(0)
}

func (t *Type) InitWith(bytes []uint8) {
	size := decode.NumberFrom(bytes)
	offset := size+8
	indices := bytes[8:offset]
	indices_size := uint64(len(indices))
	arrays := bytes[offset:]
	arrays_size := uint64(len(arrays))

	t.indices = caparr64.New(indices_size/8)
	t.arrays = caparr8.New(arrays_size)

	for i := range t.indices {
		t.indices[i] = decode.NumberFrom(indices[i*8:])
	}
	copy(t.arrays, arrays)
}

func (t *Type) Length() (length uint64) {
	length = uint64(len(t.indices)/2)
	return length
}

func (t *Type) SizeAll() (size uint64) {
	indices_size := uint64(len(t.indices))*8
	arrays_size := uint64(len(t.arrays))
	size = indices_size + arrays_size
	return size
}

func (t *Type) SizeEncoded() (size uint64) {
	return t.SizeAll()+8
}

func (t *Type) ArraySize(arr_index uint64) (size uint64) {
	curr_index := arr_index*2
	size = t.indices[curr_index+1]
	return size
}

func (t *Type) LinearSearch(arr []uint8) (arr_index uint64, match bool) {
	indices_length := len(t.indices)
	for i:=0; i<indices_length; i+=2 {
		array_index := t.indices[i]
		array_length := t.indices[i+1]
		go_match := goBytes.Compare(arr,
									t.arrays[array_index:array_index+array_length])
		match = (go_match==0)
		if match { break } else { arr_index++ }
	}
	return arr_index, match
}

func (t *Type) BinarySearch(arr []uint8) (arr_index uint64, match bool) {
	low := uint64(0)
	high := uint64(len(t.indices)) / 2
	for low!=high {
		mid := low + ((high-low) / 2)
		arr_index = mid*2
		array_index := t.indices[arr_index]
		array_length := t.indices[arr_index+1]
		go_match := goBytes.Compare(arr,
									t.arrays[array_index:array_index+array_length])
		if go_match<0 {
			high = mid
		} else if go_match>0 {
			arr_index += 2
			if low==mid {
				break
			}
			low = mid
		} else {
			match = true
			break
		}
	}
	arr_index /= 2
	return arr_index, match
}

func (t *Type) Set(arr_index uint64, arr []byte) (set bool) {
	indices_length := uint64(len(t.indices))
	curr_index := arr_index*2

	// can set
	if curr_index<indices_length {
		set = true
		arr_size := uint64(len(arr))
		array_index := t.indices[curr_index]
		array_size := t.indices[curr_index+1]
		next_index := curr_index+2

		// there is data to shift
		if (next_index<indices_length) && (arr_size!=array_size) {
			arrays_size := uint64(len(t.arrays))

			if arr_size<array_size {
				diff := array_size-arr_size
				// shift data
				copy(t.arrays[array_index+arr_size:],
					t.arrays[array_index+array_size:])
				t.arrays = t.arrays[:arrays_size-diff]
				// update indices
				indices_tmp_slice := t.indices[next_index:]
				indices_tmp_slice_length := len(indices_tmp_slice)
				for i:=0; i<indices_tmp_slice_length; i+=2 {
					indices_tmp_slice[i] -= diff
				}
			} else {
				diff := arr_size-array_size
				// ensure capacity
				min_arrays_capacity := arrays_size+diff
				t.arrays = caparr8.Ensure(t.arrays, min_arrays_capacity)
				// shift data
				t.arrays = t.arrays[:min_arrays_capacity]
				copy(t.arrays[array_index+arr_size:],
					t.arrays[array_index+array_size:])
				// update indices
				indices_tmp_slice := t.indices[next_index:]
				indices_tmp_slice_length := len(indices_tmp_slice)
				for i:=0; i<indices_tmp_slice_length; i+=2 {
					indices_tmp_slice[i] += diff
				}
			}
		}
		// overwrite data
		copy(t.arrays[array_index:], arr)
		// update array index
		t.indices[curr_index+1] = arr_size
	}

	return set
}

func (t *Type) Insert(arr_index uint64, arr []uint8) (inserted bool) {
	indices_length := uint64(len(t.indices))
	indices_next_length := indices_length+2
	curr_index := arr_index*2

	// can insert
	if curr_index<indices_next_length {
		inserted = true
		next_index := curr_index+2
		arrays_size := uint64(len(t.arrays))
		arr_size := uint64(len(arr))
		min_arrays_capacity := arrays_size+arr_size
		length := t.Length()
		new_length := length+1

		// ensure capacity
		t.indices = caparr64.Ensure(t.indices, new_length*2)
		t.indices = t.indices[:indices_next_length]
		t.arrays = caparr8.Ensure(t.arrays, min_arrays_capacity)
		t.arrays = t.arrays[:min_arrays_capacity]

		// insert
		if curr_index<indices_length {
			array_index := t.indices[curr_index]
			// shift data
			copy(t.arrays[array_index+arr_size:min_arrays_capacity],
				t.arrays[array_index:arrays_size])
			// shift indices
			copy(t.indices[next_index:indices_next_length],
				t.indices[curr_index:indices_length])
			// update indices
			indices_tmp_slice := t.indices[next_index:]
			indices_tmp_slice_length := len(indices_tmp_slice)
			for i:=0; i<indices_tmp_slice_length; i+=2 {
				indices_tmp_slice[i] += arr_size
			}
			// overwrite array
			copy(t.arrays[array_index:], arr)
		// append
		} else {
			t.indices[curr_index] = arrays_size
			// overwrite array
			copy(t.arrays[arrays_size:], arr)
		}
		// update array index
		t.indices[curr_index+1] = arr_size
	}

	return inserted
}

func (t *Type) Remove(arr_index uint64) {
	curr_index := arr_index*2
	indices_length := uint64(len(t.indices))
	indices_new_length := indices_length-2
	array_index := t.indices[curr_index]
	array_length := t.indices[curr_index+1]
	arrays_length := uint64(len(t.arrays))
	arrays_new_length := arrays_length-array_length

	copy(t.indices[curr_index:indices_new_length],
		 t.indices[curr_index+2:])
	t.indices = t.indices[:indices_new_length]
	for i:=curr_index; i<indices_new_length; i+=2 {
		t.indices[i] -= array_length
	}
	copy(t.arrays[array_index:arrays_new_length],
		 t.arrays[array_index+array_length:])
	t.arrays = t.arrays[:arrays_new_length]
}

func (t *Type) Get(arr_index uint64) (arr []uint8) {
	curr_index := arr_index*2
	array_index := t.indices[curr_index]
	array_length := t.indices[curr_index+1]
	arr = t.arrays[array_index:array_index+array_length]
	return arr
}

func (t *Type) GetAll() (arr_arr [][]uint8) {
	arr_arr = make([][]byte, t.Length())
	for i := range arr_arr {
		curr_index := i*2
		array_index := t.indices[curr_index]
		array_length := t.indices[curr_index+1]
		arr_arr[i] = t.arrays[array_index:array_index+array_length]
	}
	return arr_arr
}

func (t *Type) EncodeTo(bytes []uint8) {
	indices_length := uint64(len(t.indices))
	indices_size := indices_length*8
	data := bytes[8:]
	encode.NumberTo(bytes, indices_size)
	for i, number := range t.indices {
		encode.NumberTo(data[i*8:], number)
	}
	copy(data[indices_size:], t.arrays)
}















