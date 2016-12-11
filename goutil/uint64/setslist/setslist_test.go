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
	"testing"
)

func TestAddEmptySet(t *testing.T) {
	var tt Type
	tt.Init()
	tt.AddEmptySet(0)

	if len(tt.indices)!=2 {
		t.Errorf("wrong number of indices: %d", len(tt.indices))
	}
	if len(tt.maps)!=0 {
		t.Errorf("wrong number of maps: %d", len(tt.maps))
	}

	tt.AddEmptySet(0)
	if len(tt.indices)!=4 {
		t.Errorf("wrong number of indices: %d", len(tt.indices))
	}
	if len(tt.maps)!=0 {
		t.Errorf("wrong number of maps: %d", len(tt.maps))
	}

	tt.AddEmptySet(1)
	if len(tt.indices)!=6 {
		t.Errorf("wrong number of indices: %d", len(tt.indices))
	}
	if len(tt.maps)!=0 {
		t.Errorf("wrong number of maps: %d", len(tt.maps))
	}
}

func TestRemoveSet(t *testing.T) {
	var tt Type
	tt.Init()
	tt.AddEmptySet(0)
	tt.AddEmptySet(0)
	tt.AddEmptySet(0)
	tt.RemoveSet(0)
	if len(tt.indices)!=4 {
		t.Errorf("wrong number of indices: %d", len(tt.indices))
	}
	if len(tt.maps)!=0 {
		t.Errorf("wrong number of maps: %d", len(tt.maps))
	}
}

func TestAddSetValue(t *testing.T) {
	var tt Type
	tt.Init()
	tt.AddEmptySet(0)
	tt.AddEmptySet(0)
	tt.AddEmptySet(0)
	tt.AddSetValue(0, 5)

	if len(tt.indices)!=6 {
		t.Errorf("wrong number of indices: %d", len(tt.indices))
	}
	if tt.maps[0]!=5 {
		t.Errorf("wrong map number: %d", tt.maps[0])
	}

	tt.AddSetValue(1, 7)
	if tt.maps[1]!=7 {
		t.Errorf("wrong map number: %d", tt.maps[1])
	}

	tt.AddSetValue(2, 1)
	tt.AddSetValue(1, 8)
	if tt.maps[2]!=7 {
		t.Errorf("wrong map number: %d", tt.maps[2])
	}
	if tt.maps[3]!=1 {
		t.Errorf("wrong map number: %d", tt.maps[3])
	}
}















