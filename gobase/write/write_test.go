//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package write


import (
	"testing"
	"github.com/vitalibaumtrok/gobase/allocate"
	"github.com/vitalibaumtrok/gobase/deallocate"
	"github.com/vitalibaumtrok/gobase/write"
	"github.com/vitalibaumtrok/gobase/read"
)


func TestStringToFileOnPath(t *testing.T) {
	str := "hallo"
	path := "test.tmp"
	err := allocate.FileOnPath(path)
	if err!=nil {
		t.Errorf("Could not create file: %s", err.Error())
	} else {
		err = StringToFileOnPath(path, str)
		if err!=nil {
			t.Errorf("Could not write to file: %s", err.Error())
		} else {
			str2, err2 := read.StringFromFileOnPath(path)
			if err2!=nil {
				err = deallocate.FileOnPath(path)
				if err!=nil {
					t.Errorf("Multiple errors: %s; %s", err2.Error(), err.Error())
				} else {
					t.Errorf("Could not read from file: %s", err2.Error())
				}
			} else {
				if str!=str2 {
					err = deallocate.FileOnPath(path)
					if err!=nil {
						t.Errorf("Wront string after write and more: %s; %s", err2.Error(), err.Error())
					} else {
						t.Errorf("Wront string after write: %s", err2.Error())
					}
				} else {
					err = deallocate.FileOnPath(path)
					if err!=nil {
						t.Errorf("Could not delete file: %s", err.Error())
					}
				}
			}
		}
	}
}














