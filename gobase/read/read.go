//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package read


import (
	"io/ioutil"
)


func BytesFromFileOnPath(path string) (bytes []byte, err error) {
	return ioutil.ReadFile(path)
}

func StringFromFileOnPath(path string) (string, error) {
	bytes, err := BytesFromFileOnPath(path)
	return string(bytes), err
}














