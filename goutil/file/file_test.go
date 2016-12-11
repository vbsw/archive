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

package file

import (
	"testing"
)

func TestExists(t *testing.T) {
	path := "file.go"
	b, err := Exists(path)
	if b {
		if err!=nil {
			t.Errorf("File access caused an error: %s", err.Error())
		}
	} else {
		if err==nil {
			t.Errorf("File does not exist")
		} else {
			t.Errorf("File does not exist: %s", err.Error())
		}
	}
}

func TestDelete(t *testing.T) {
	path := "test.tmp"
	err := Create(path)
	if err!=nil {
		t.Errorf("Could not create file: %s", err.Error())
	} else {
		err = Delete(path)
		if err!=nil {
			t.Errorf("Could not delete file: %s", err.Error())
		}
	}
}

func TestWriteStr(t *testing.T) {
	str := "hallo"
	path := "test.tmp"
	err := Create(path)
	if err!=nil {
		t.Errorf("Could not create file: %s", err.Error())
	} else {
		err = WriteStr(path, str)
		if err!=nil {
			t.Errorf("Could not write to file: %s", err.Error())
		}
		err = Delete(path)
		if err!=nil {
			t.Errorf("Could not delete file: %s", err.Error())
		}
	}
}

func TestAppendStr(t *testing.T) {
	str1 := "hallo"
	str2 := "ciao"
	path := "test.tmp"
	err := Create(path)
	if err!=nil {
		t.Errorf("Could not create file: %s", err.Error())
	} else {
		err = WriteStr(path, str1)
		if err==nil {
			err = AppendStr(path, str2)
			if err==nil {
				var fileStr string
				fileStr, err = ReadStr(path)
				if err==nil {
					if (fileStr[:len(str1)]!=str1) ||
						(fileStr[len(str1):]!=str2) {
						t.Errorf("Append has wrong result")
					}
				} else {
					t.Errorf("Could not read file: %s", err.Error())
				}
			} else {
				t.Errorf("Could not append to file: %s", err.Error())
			}
		} else {
			t.Errorf("Could not write to file: %s", err.Error())
		}
		err = Delete(path)
		if err!=nil {
			t.Errorf("Could not delete file: %s", err.Error())
		}
	}
}

func TestClear(t *testing.T) {
	str := "hallo"
	path := "test.tmp"
	err := Create(path)
	if err==nil {
		err = WriteStr(path, str)
		if err==nil {
			err = Clear(path)
			if err==nil {
				var fileStr string
				fileStr, err = ReadStr(path)
				if err==nil {
					if (len(fileStr)>0) {
						t.Errorf("File isn't empty")
					}
				} else {
					t.Errorf("Could not read file: %s", err.Error())
				}
			} else {
				t.Errorf("Could not clear file: %s", err.Error())
			}
		} else {
			t.Errorf("Could not write to file: %s", err.Error())
		}
		err = Delete(path)
		if err!=nil {
			t.Errorf("Could not delete file: %s", err.Error())
		}
	} else {
		t.Errorf("Could not create file: %s", err.Error())
	}
}









