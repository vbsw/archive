//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package apd


import (
	"os"
	"fmt"
	"errors"
)


func BytesToFileOnPath(path string, bytes []byte) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
	if err==nil {
		_, err = file.Write(bytes)
		if err!=nil {
			err2 := file.Close()
			if err2!=nil {
				errStr := fmt.Sprintf("%s; %s", err.Error(), err2.Error())
				err = errors.New(errStr)
			}
		} else {
			err = file.Close()
		}
	}
	return err
}

func StringToFileOnPath(path string, text string) error {
	return BytesToFileOnPath(path, []byte(text))
}

func AnyBytesToByteArray(dstBytes []byte, size int) []byte {
	length := len(dstBytes)
	desiredLength := length+size
	if desiredLength>cap(dstBytes) {
		newBytes := make([]byte, desiredLength)
		copy(newBytes, dstBytes)
		dstBytes = newBytes
	} else {
		dstBytes = dstBytes[:desiredLength]
	}
	return dstBytes
}














