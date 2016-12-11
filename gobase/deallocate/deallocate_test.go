//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package deallocate


import (
	"testing"
	"github.com/vitalibaumtrok/gobase/allocate"
)


func TestFileOnPath(t *testing.T) {
	path := "test.tmp"
	err := allocate.FileOnPath(path)
	if err!=nil {
		t.Errorf("Could not create file: %s", err.Error())
	} else {
		err = FileOnPath(path)
		if err!=nil {
			t.Errorf("Could not delete file: %s", err.Error())
		}
	}
}













