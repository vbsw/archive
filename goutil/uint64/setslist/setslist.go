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

package setslist

import (
	"github.com/vitalibaumtrok/util/uint64/caparr"
	"github.com/vitalibaumtrok/util/uint64/decode"
	"github.com/vitalibaumtrok/util/uint64/encode"
)

const (
	MINIMUM_INITIAL_BYTES_SIZE = 8
)

type Type struct {
	indices []uint64
	sets []uint64
}

func New() (t *Type) {
	t = new(Type)
	t.Init()
	return t
}

func NewWith(bytes []byte) (t *Type) {
	t = new(Type)
	t.InitWith(bytes)
	return t
}

func (t *Type) Init() {
	t.indices = caparr.New(0)
	t.sets = caparr.New(0)
}

func (t *Type) InitWith(bytes []byte) {
	size := decode.NumberFrom(bytes)
	sets_offset := size+8
	indices := bytes[8:sets_offset]
	indices_size := uint64(len(indices))
	sets := bytes[sets_offset:]
	sets_size := uint64(len(sets))

	t.indices = caparr.New(indices_size/8)
	t.sets = caparr.New(sets_size/8)

	for i := range t.indices {
		t.indices[i] = decode.NumberFrom(indices[i*8:])
	}
	for i := range t.sets {
		t.sets[i] = decode.NumberFrom(sets[i*8:])
	}
}

func (t *Type) Length() (length uint64) {
	length = uint64(len(t.indices)/2)
	return length
}

func (t *Type) SizeAll() (size uint64) {
	indices_size := uint64(len(t.indices))*8
	sets_size := uint64(len(t.sets))*8
	size = indices_size + sets_size
	return size
}

func (t *Type) SizeEncoded() (size uint64) {
	return t.SizeAll()+MINIMUM_INITIAL_BYTES_SIZE
}

func (t *Type) GetSet(s_index uint64) (set []uint64) {
	curr_index := s_index*2
	set_index := t.indices[curr_index]
	set_length := t.indices[curr_index+1]
	set = t.sets[set_index:set_index+set_length]
	return set
}

func (t *Type) IncrementValuesFrom(s_value uint64) {
	for i, value := range t.sets {
		if value>=s_value {
			t.sets[i]++
		}
	}
}

func (t *Type) DecrementValuesFrom(s_value uint64) {
	for i, value := range t.sets {
		if value>=s_value {
			t.sets[i]--
		}
	}
}

func (t *Type) AddEmptySet(s_index uint64) {
	indices_length := uint64(len(t.indices))
	indices_next_length := indices_length+2
	curr_index := s_index*2
	next_index := curr_index+2

	t.indices = caparr.Ensure(t.indices, indices_next_length)
	t.indices = t.indices[:indices_next_length]
	copy(t.indices[next_index:], t.indices[curr_index:indices_length])
	t.indices[curr_index+1] = 0

	if curr_index>0 {
		prev_index := curr_index-2
		set_index := t.indices[prev_index]
		set_length := t.indices[prev_index+1]
		t.indices[curr_index] = set_index+set_length
	} else {
		t.indices[curr_index] = 0
	}
}

func (t *Type) RemoveSet(s_index uint64) {
	curr_index := s_index*2
	set_length := t.indices[curr_index+1]
	indices_length := uint64(len(t.indices))
	copy(t.indices[curr_index:indices_length-2], t.indices[curr_index+2:])
	t.indices = t.indices[:indices_length-2]
	t.deleteSet(curr_index, set_length)
}

func (t *Type) HasSetValue(s_index uint64, s_value uint64) (has bool) {
	curr_index := s_index*2
	set_index := t.indices[curr_index]
	set_length := t.indices[curr_index+1]
	set := t.sets[set_index:set_index+set_length]

	for _, value := range set {
		if value==s_value {
			has = true
			break
		}
	}
	return has
}

func (t *Type) AddSetValue(s_index uint64, s_value uint64) {
	curr_index := s_index*2
	next_index := curr_index+2
	set_index := t.indices[curr_index]
	sets_length := uint64(len(t.sets))
	sets_new_length := sets_length+1
	indices_length := uint64(len(t.indices))

	t.sets = caparr.Ensure(t.sets, sets_new_length)
	t.sets = t.sets[:sets_new_length]
	copy(t.sets[set_index+1:sets_new_length], t.sets[set_index:sets_length])
	t.sets[set_index] = s_value
	t.indices[curr_index+1]++
	for i:=next_index; i<indices_length; i+=2 {
		t.indices[i]++
	}
}

func (t *Type) RemoveSetValue(s_index uint64, s_value uint64) (removed bool) {
	curr_index := s_index*2
	set_index := t.indices[curr_index]
	set_length := t.indices[curr_index+1]
	set := t.sets[set_index:set_index+set_length]

	for offset, value := range set {
		if value==s_value {
			next_index := curr_index+2
			sets_length := uint64(len(t.sets))
			removed = true
			t.indices[curr_index+1]--
			indices_length := uint64(len(t.indices))
			for i:=next_index; i<indices_length; i+=2 {
				t.indices[i]--
			}

			// shift to left
			if set_index+uint64(offset)+1<sets_length {
				copy(t.sets[set_index+uint64(offset):sets_length-1],
					t.sets[set_index+uint64(offset)+1:])
			}
			t.sets = t.sets[:sets_length-1]
			break
		}
	}
	return removed
}

func (t *Type) RemoveAllSetValues(s_value uint64) (removed uint64) {
	var i uint64
	length := t.Length()
	for i=0; i<length; i++ {
		if t.RemoveSetValue(i, s_value) {
			removed++
		}
	}
	return removed
}

func (t *Type) EncodeTo(bytes []byte) {
	indices_length := uint64(len(t.indices))
	indices_size := indices_length*8
	indices_bytes := bytes[8:]
	sets_bytes := indices_bytes[indices_size:]
	encode.NumberTo(bytes, indices_size)
	for i, number := range t.indices {
		encode.NumberTo(indices_bytes[i*8:], number)
	}
	for i, number := range t.sets {
		encode.NumberTo(sets_bytes[i*8:], number)
	}
}

func (t *Type) deleteSet(s_index uint64, s_length uint64) {
	sets_length := uint64(len(t.sets))
	sets_new_length := sets_length-s_length
	copy(t.sets[s_index:sets_new_length], t.sets[s_index+s_length:])
	t.sets = t.sets[:sets_new_length]
}















