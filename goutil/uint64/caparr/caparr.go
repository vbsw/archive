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

package caparr

import (
	caparr8 "github.com/vitalibaumtrok/util/uint8/caparr"
)

const (
	INITIAL_SIZE = caparr8.INITIAL_SIZE / 8
)

func New(length uint64) ([]uint64) {
	initial_size := Next(length)
	return make([]uint64, length, initial_size)
}

func Next(min_capacity uint64) (new_capacity uint64) {
	new_capacity = INITIAL_SIZE
	for new_capacity<=min_capacity {
		new_capacity = new_capacity*2 + (8/8)
	}
	return new_capacity
}

func Ensure(arr []uint64, new_length uint64) ([]uint64) {
	arr_capacity := uint64(cap(arr))
	if arr_capacity<new_length {
		arr_length := uint64(len(arr))
		arr_new_capacity := Next(new_length)
		new_arr := make([]uint64, arr_length, arr_new_capacity)
		copy(new_arr, arr)
		return new_arr
	}
	return arr
}









