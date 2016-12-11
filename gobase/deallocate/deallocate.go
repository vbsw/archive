//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package deallocate


import (
	"os"
)


func FileOnPath(path string) error {
	return os.Remove(path)
}

func IntegerFromArray(integers []int, index int) []int {
	integersLength := len(integers)
	copy(integers[index:], integers[index+1:])
	return integers[:integersLength-1]
}

func BytesFromArray(bytes []byte, offset, length int) []byte {
	copy(bytes[offset:], bytes[offset+length:])
	return bytes[:len(bytes)-length]
}














