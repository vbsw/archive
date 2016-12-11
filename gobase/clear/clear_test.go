//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package clear


import (
	"testing"
	"github.com/vitalibaumtrok/gobase/allocate"
	"github.com/vitalibaumtrok/gobase/deallocate"
	"github.com/vitalibaumtrok/gobase/write"
	"github.com/vitalibaumtrok/gobase/read"
)


func TestFileOnPath(t *testing.T) {
	str := "hallo"
	path := "test.tmp"
	err := allocate.FileOnPath(path)
	if err==nil {
		err = write.StringToFileOnPath(str, path)
		if err==nil {
			err = FileOnPath(path)
			if err==nil {
				var fileStr string
				fileStr, err = read.StringFromFileOnPath(path)
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
		err = deallocate.FileOnPath(path)
		if err!=nil {
			t.Errorf("Could not delete file: %s", err.Error())
		}
	} else {
		t.Errorf("Could not create file: %s", err.Error())
	}
}














