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
	"testing"
)

func TestNext(t *testing.T) {
	capacity1 := uint64(10)
	capacity2 := uint64(1000)
	if Next(capacity1)<capacity1 {
		t.Errorf("capacity too low")
	}
	if Next(capacity2)<capacity2 {
		t.Errorf("capacity too low")
	}
}









