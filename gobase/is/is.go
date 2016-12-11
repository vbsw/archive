//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)

package is

import (
	"os"
)

func ExistentFile(path string) (bool, error) {
	_, err := os.Stat(path)
	if err==nil {
		return true, err
	} else if os.IsNotExist(err) {
		return false, err
	} else {
		return false, err
	}
}

func ReadableFile(path string) (bool, error) {
	file, err := os.Open(path)
	if err==nil {
		return true, file.Close()
	} else {
		return false, err
	}
}

func DirectoryFile(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	return fileInfo.IsDir(), err
}

func WritableFile(path string) (bool, error) {
	// TODO implement WritableFile() bool
	return true, nil
}

func ExecutableFile(path string) (bool, error) {
	// TODO implement ExecutableFile() bool
	return true, nil
}









