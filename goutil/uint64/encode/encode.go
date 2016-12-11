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

package encode

import (
	"errors"
)

const ERR_MSG = "can't encode number"

func NumberTo(arr []uint8, number uint64) {
	arr[0] = uint8(number)
	arr[1] = uint8(number>>8)
	arr[2] = uint8(number>>16)
	arr[3] = uint8(number>>24)
	arr[4] = uint8(number>>32)
	arr[5] = uint8(number>>40)
	arr[6] = uint8(number>>48)
	arr[7] = uint8(number>>56)
}

func AsciiNumberTo(arr []uint8, number uint64) (err error) {
	if number<10 {
		arr[0] = uint8(number+48)
	} else {
		err = errors.New(ERR_MSG)
	}
	return err
}










