//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)

package is

import (
	"testing"
)

func TestExistentFile(t *testing.T) {
	path := "file.go"
	b, err := ExistentFile(path)
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









