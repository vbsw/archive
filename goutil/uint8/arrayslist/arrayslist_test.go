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
	"testing"
	"github.com/vitalibaumtrok/tag/conv"
)

func TestInit(t *testing.T) {
	larr := new(Type)
	larr.Init()

	if larr.indices==nil { t.Errorf("indices are nil") }
	if len(larr.indices)>0 { t.Errorf("indices are not empty") }
	if cap(larr.indices)<=0 { t.Errorf("capacity of indices is 0") }
	if larr.arrays==nil { t.Errorf("arrays are nil") }
	if len(larr.arrays)>0 { t.Errorf("arrays are not empty") }
	if cap(larr.arrays)<=0 { t.Errorf("capacity of arrays is 0") }
}

func TestInitWith(t *testing.T) {
	bytes_size := 8*(1+2)
	arrays_size := 0
	larr := new(Type)
	bytes := make([]byte, bytes_size)
	conv.UIntToBytes(bytes, 8*2)
	larr.InitWith(bytes)

	if larr.indices==nil { t.Errorf("indices are nil") }
	if len(larr.indices)!=2 { t.Errorf("indices don't have size of 2") }
	if cap(larr.indices)<2 { t.Errorf("capacity of indices is less then 2") }
	if larr.arrays==nil { t.Errorf("arrays are nil") }
	if len(larr.arrays)!=arrays_size {
		t.Errorf("arrays don't have size of %d, but %d", arrays_size, len(larr.arrays))
	}
	if cap(larr.arrays)<arrays_size {
		t.Errorf("capacity of arrays is less then %d", arrays_size)
	}

	if larr.indices[0]!=0 { t.Errorf("arrays index is not 0") }
}

func TestLength(t *testing.T) {
	bytes_size := 8*(1+2)
	larr := new(Type)
	bytes := make([]byte, bytes_size)
	conv.UIntToBytes(bytes, 8*2)
	larr.InitWith(bytes)

	if larr.Length()!=1 {
		t.Errorf("length of indices is not %d, but %d", 1, larr.Length())
	}
}

func TestSizeAll(t *testing.T) {
	bytes_size := 8*(1+2)
	larr := new(Type)
	bytes := make([]byte, bytes_size)
	conv.UIntToBytes(bytes, 8*2)
	larr.InitWith(bytes)

	if len(larr.indices)!=2 {
		t.Errorf("length of indices is not 2, but %d", len(larr.indices))
	}
	if larr.SizeAll()!=uint64(bytes_size)-8 {
		t.Errorf("size of all is not %d, but %d", bytes_size, larr.SizeAll())
	}
}

func TestInsert(t *testing.T) {
	arrays0_size := uint64(30)
	arrays1_size := uint64(50)
	arrays0 := make([]byte, arrays0_size)
	arrays1 := make([]byte, arrays1_size)
	larr := New()

	larr.Insert(0, arrays0)
	if len(larr.indices)!=2 {
		t.Errorf("length of indices is not 2, but %d", len(larr.indices))
	}
	if larr.indices[0]!=0 {
		t.Errorf("first index does't point to 0, but %d", larr.indices[0])
	}
	if larr.indices[1]!=arrays0_size {
		t.Errorf("size of 0 is not %d, but %d", arrays0_size, larr.indices[1])
	}
	if cap(larr.indices)<2 {
		t.Errorf("capacity is not %d, but %d", 2, cap(larr.indices))
	}

	larr.Insert(0, arrays1)
	if len(larr.indices)!=4 {
		t.Errorf("length of indices is not 4, but %d", len(larr.indices))
	}
	if larr.indices[0]!=0 {
		t.Errorf("first index does't point to 0, but %d", larr.indices[0])
	}
	if larr.indices[1]!=arrays1_size {
		t.Errorf("size of 0 is not %d, but %d", arrays1_size, larr.indices[1])
	}
	if larr.indices[2]!=arrays1_size {
		t.Errorf("second index does't point to %d, but %d",
				 arrays1_size, larr.indices[2])
	}
	if larr.indices[3]!=arrays0_size {
		t.Errorf("size of 1 is not %d, but %d", arrays0_size, larr.indices[3])
	}

	larr.Insert(0, arrays1)
	larr.Insert(0, arrays1)
	larr.Insert(0, arrays1)
	larr.Insert(0, arrays1)
	larr.Insert(0, arrays1)
	larr.Insert(0, arrays1)
	larr.Insert(0, arrays1)
	if len(larr.indices)!=9*2 {
		t.Errorf("length of indices is not 4, but %d", len(larr.indices))
	}
	
	if larr.Insert(1000, arrays1) {
		t.Errorf("insertion at impossible index: %d", 1000)
	}
}

func TestSizeOf(t *testing.T) {
	arrays0_size := uint64(5)
	arrays1_size := uint64(10)
	arrays0 := make([]byte, arrays0_size)
	arrays1 := make([]byte, arrays1_size)
	larr := New()

	larr.Insert(0, arrays0)
	if larr.SizeOf(0)!=arrays0_size {
		t.Errorf("size of 0 is not %d, but %d", arrays0_size, larr.SizeOf(0))
	}
	larr.Insert(1, arrays1)
	if larr.SizeOf(0)!=arrays0_size {
		t.Errorf("size of 0 is not %d, but %d", arrays0_size, larr.SizeOf(0))
	}
	if larr.SizeOf(1)!=arrays1_size {
		t.Errorf("size of 1 is not %d, but %d", arrays1_size, larr.SizeOf(0))
	}
}

func TestCopyTo(t *testing.T) {
	arrays0_size := uint64(30)
	arrays1_size := uint64(50)
	arrays0 := make([]byte, arrays0_size)
	arrays1 := make([]byte, arrays1_size)
	bytes := make([]byte, 8+8*6+arrays0_size+arrays1_size+100)
	larr := New()

	conv.UIntToBytes(arrays0, 30)
	larr.Insert(0, arrays0)
	larr.Insert(1, arrays1)
	larr.EncodeTo(bytes)
	if conv.BytesToUInt(bytes)!=8*4 {
		t.Errorf("offset not %d, but %d", 8*4, conv.BytesToUInt(bytes))
	}
	if conv.BytesToUInt(bytes[8+8*4:])!=30 {
		t.Errorf("number not %d, but %d", 30, conv.BytesToUInt(bytes[8+8*4:]))
	}
	larr.Insert(1, arrays1)
	larr.EncodeTo(bytes)
	if conv.BytesToUInt(bytes)!=8*6 {
		t.Errorf("offset not %d, but %d", 8*6, conv.BytesToUInt(bytes))
	}
}

func TestLinearSearch(t *testing.T) {
	datum0_size := uint64(5)
	datum1_size := uint64(10)
	datum2_size := uint64(10)
	datum3_size := uint64(10)
	datum4_size := uint64(10)
	datum0 := make([]byte, datum0_size)
	datum1 := make([]byte, datum1_size)
	datum2 := make([]byte, datum2_size)
	datum3 := make([]byte, datum3_size)
	datum4 := make([]byte, datum4_size)
	larr := New()

	conv.UIntToBytes(datum1, 20)
	conv.UIntToBytes(datum2, 30)
	conv.UIntToBytes(datum3, 40)
	conv.UIntToBytes(datum4, 50)
	larr.Insert(0, datum0)
	larr.Insert(0, datum1)
	larr.Insert(0, datum2)
	larr.Insert(0, datum3)

	if goBytes.Compare(larr.Get(3), datum0)!=0 { t.Errorf("Insert failed") }
	i, match := larr.LinearSearch(datum0)
	if !match { t.Errorf("datum not found") }
	if i!=3 { t.Errorf("wrong index: %d", i) }

	i, match = larr.LinearSearch(datum4)
	if match { t.Errorf("datum found") }
	if i!=4 { t.Errorf("wrong index: %d", i) }
}

func TestBinarySearch(t *testing.T) {
	datum0_size := uint64(5)
	datum1_size := uint64(10)
	datum2_size := uint64(10)
	datum3_size := uint64(10)
	datum4_size := uint64(10)
	datum5_size := uint64(10)
	datum0 := make([]byte, datum0_size)
	datum1 := make([]byte, datum1_size)
	datum2 := make([]byte, datum2_size)
	datum3 := make([]byte, datum3_size)
	datum4 := make([]byte, datum4_size)
	datum5 := make([]byte, datum5_size)
	larr := New()

	conv.UIntToBytes(datum1, 20)
	conv.UIntToBytes(datum2, 30)
	conv.UIntToBytes(datum3, 40)
	conv.UIntToBytes(datum4, 50)
	conv.UIntToBytes(datum5, 35)
	larr.Insert(0, datum0)
	larr.Insert(1, datum1)
	larr.Insert(2, datum2)
	larr.Insert(3, datum3)

	if goBytes.Compare(larr.Get(1), datum1)!=0 { t.Errorf("Insert failed") }
	i, match := larr.BinarySearch(datum1)
	if !match { t.Errorf("datum not found") }
	if i!=1 { t.Errorf("wrong index: %d", i) }

	i, match = larr.BinarySearch(datum4)
	if match { t.Errorf("datum found") }
	if i!=4 { t.Errorf("wrong index: %d", i) }

	i, match = larr.BinarySearch(datum5)
	if match { t.Errorf("datum found") }
	if i!=3 { t.Errorf("wrong index: %d", i) }
}

func TestSet(t *testing.T) {
	arrays0_size := uint64(5)
	arrays1_size := uint64(10)
	arrays0 := make([]byte, arrays0_size)
	arrays1 := make([]byte, arrays1_size)
	arrays_final_size := arrays0_size*3 + arrays1_size
	larr := New()

	larr.Insert(0, arrays0)
	larr.Insert(1, arrays1)
	larr.Insert(2, arrays0)
	larr.Insert(3, arrays1)
	if larr.Length()!=4 { t.Errorf("wrong length: %d", larr.Length()) }
	if !larr.Set(1, arrays0) { t.Errorf("datum isn't set") }
	if larr.Length()!=4 { t.Errorf("wrong length: %d", larr.Length()) }
	if uint64(len(larr.arrays))!=arrays_final_size {
		t.Errorf("wrong length: %d", larr.Length())
	}
}

func TestGet(t *testing.T) {
	arrays0_size := uint64(5)
	arrays1_size := uint64(10)
	arrays0 := make([]byte, arrays0_size)
	arrays1 := make([]byte, arrays1_size)
	larr := New()

	larr.Insert(0, arrays0)
	larr.Insert(1, arrays1)
	if uint64(len(larr.Get(0)))!=arrays0_size {
		t.Errorf("wrong size: %d", len(larr.Get(0)))
	}
	if uint64(len(larr.Get(1)))!=arrays1_size {
		t.Errorf("wrong size: %d", len(larr.Get(1)))
	}
	larr.Insert(1, arrays0)
	if uint64(len(larr.Get(1)))!=arrays0_size {
		t.Errorf("wrong size: %d", len(larr.Get(1)))
	}
	if uint64(len(larr.Get(2)))!=arrays1_size {
		t.Errorf("wrong size: %d", len(larr.Get(2)))
	}
}

func TestGetAll(t *testing.T) {
	arrays0_size := uint64(5)
	arrays1_size := uint64(10)
	arrays0 := make([]byte, arrays0_size)
	arrays1 := make([]byte, arrays1_size)
	larr := New()

	larr.Insert(0, arrays0)
	larr.Insert(1, arrays1)
	larr.Insert(1, arrays0)

	all := larr.GetAll()
	if len(all)!=3 { t.Errorf("wrong number of arrays: %d", len(all)) }

	larr.Insert(1, arrays1)
	all = larr.GetAll()
	if len(all)!=4 { t.Errorf("wrong number of arrays: %d", len(all)) }
}















