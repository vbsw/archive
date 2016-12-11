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

package decode

import (
	"errors"
)

const ERR_MSG = "can't decode number"

func NumberFrom(arr []uint8) (number uint64) {
	return uint64(arr[0]) | uint64(arr[1])<<8 | uint64(arr[2])<<16 |
		uint64(arr[3])<<24 | uint64(arr[4])<<32 | uint64(arr[5])<<40 |
		uint64(arr[6])<<48 | uint64(arr[7])<<56
}

func AsciiNumberFrom(arr []uint8) (number uint64, err error) {
	b := arr[0]
	if b>47 && b<58 {
		number = uint64(b-48)
	} else {
		err = errors.New(ERR_MSG)
	}
	return number, err
}










