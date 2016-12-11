//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package write


import (
	"io/ioutil"
)


func BytesToFileOnPath(path string, bytes []byte) error {
	return ioutil.WriteFile(path, bytes, 0666)
}

func StringToFileOnPath(path string, text string) error {
	return BytesToFileOnPath(path, []byte(text))
}














