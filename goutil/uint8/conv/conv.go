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

package conv

func AsciiToLowerCase(bytes []uint8) {
	for i, b := range bytes {
		if (b>64) && (b<91) {
			bytes[i] = (b|32)
		}
	}
}

func AsciiToHigherCase(bytes []uint8) {
	for i, b := range bytes {
		if (b>96) && (b<123) {
			bytes[i] = (b&95)
		}
	}
}










