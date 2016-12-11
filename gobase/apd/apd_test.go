//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package apd


import (
	"testing"
	"github.com/vitalibaumtrok/gobase/allocate"
	"github.com/vitalibaumtrok/gobase/deallocate"
	"github.com/vitalibaumtrok/gobase/write"
	"github.com/vitalibaumtrok/gobase/read"
)


func TestStringToFileOnPath(t *testing.T) {
	str1 := "hallo"
	str2 := "ciao"
	path := "test.tmp"
	err := allocate.FileOnPath(path)
	if err!=nil {
		t.Errorf("Could not create file: %s", err.Error())
	} else {
		err = write.StringToFileOnPath(str1, path)
		if err==nil {
			err = StringToFileOnPath(path, str2)
			if err==nil {
				var fileStr string
				fileStr, err = read.StringFromFileOnPath(path)
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
		err = deallocate.FileOnPath(path)
		if err!=nil {
			t.Errorf("Could not delete file: %s", err.Error())
		}
	}
}









