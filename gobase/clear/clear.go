//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package clear


import (
	"github.com/vitalibaumtrok/gobase/write"
)


func FileOnPath(path string) error {
	return write.StringToFileOnPath("", path)
}














