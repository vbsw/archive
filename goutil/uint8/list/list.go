/* 
 * util - a library for the programming language Go
 * Copyright 2014 Vitali Baumtrok
 * 
 * This file is part of util.
 * 
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * 
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * 
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 * 
 */

package list

import (
	"github.com/vitalibaumtrok/util/uint8/caparr"
)

func New() (list []uint8) {
	return make([]uint8, 0, caparr.INITIAL_SIZE)
}

func Append(list []uint8, unsigned_integer uint8) (new_list []uint8) {
	list_length := uint64(len(list))
	list_new_length := list_length+1
	list = caparr.Ensure(list, list_new_length)
	new_list = append(list, unsigned_integer)
	return new_list
}

func AppendList(list1 []uint8, list2 []uint8) (new_list []uint8) {
	list1_length := uint64(len(list1))
	list2_length := uint64(len(list2))

	if list1_length>0 {
		if list2_length>0 {
			new_list = caparr.Ensure(list1, list1_length+list2_length)
			copy(new_list[list1_length:], list2)
		} else {
			new_list = list1
		}
	} else {
		new_list = list2
	}

	return new_list
}

func Insert(list []uint8, index uint64, value uint8) ([]uint8) {
	list_length := uint64(len(list))
	list_new_length := list_length+1
	list = caparr.Ensure(list, list_new_length)
	list = list[:list_new_length]
	if index<list_length {
		copy(list[index+1:], list[index:list_length])
	}
	list[index] = value
	return list
}

func Remove(list []uint8, index uint64) ([]uint8) {
	list_length := uint64(len(list))
	if index+1<list_length {
		copy(list[index:], list[index+1:list_length])
	}
	return list[:list_length-1]
}

func LinearSearch(list []uint8, unsigned_integer uint8) (index uint64, match bool) {
	for i, list_val := range list {
		if list_val==unsigned_integer {
			return uint64(i), true
		}
	}
	return 0, false
}









