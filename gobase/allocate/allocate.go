//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package allocate


import (
	"os"
	"fmt"
	"errors"
)


func FileOnPath(path string) (err error) {
	var file *os.File
	file, err = os.Create(path)
	if err==nil {
		err = file.Close()
	} else {
		err2 := file.Close()
		if err2!=nil {
			errStr := fmt.Sprintf("%s; %s", err.Error(), err2.Error())
			err = errors.New(errStr)
		}
	}
	return err
}














