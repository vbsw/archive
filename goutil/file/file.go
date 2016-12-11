/* 
 * util - a library for the programming language Go
 * Copyright 2014 Vitali Baumtrok
 * 
 * This file is part of util.
 * 
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * 
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * 
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 * 
 */

package file

import (
	"os"
	"fmt"
	"errors"
	"io/ioutil"
)

func Exists(path string) (exists bool, err error) {
	_, errTmp := os.Stat(path)
	if errTmp==nil {
		exists = true
	} else if os.IsNotExist(errTmp) {
		exists = false
	} else {
		exists = false
		err = errTmp
	}
	return exists, err
}

func Readable(path string) (readable bool, err error) {
	file, errTmp := os.Open(path)
	if errTmp==nil {
		readable = true
		err = file.Close()
	} else {
		err = errTmp
	}
	return readable, err
}

func Writable(path string) (writable bool, err error) {
	// TODO implement file.Writable() bool
	writable = true
	return writable, err
}

func Executable(path string) (executable bool, err error) {
	// TODO implement file.Executable() bool
	executable = true
	return executable, err
}

func Read(path string) (bytes []uint8, err error) {
	return ioutil.ReadFile(path)
}

func ReadStr(path string) (text string, err error) {
	bytes, errTmp := Read(path)
	return string(bytes), errTmp
}

func Write(path string, bytes []uint8) (err error) {
	return ioutil.WriteFile(path, bytes, 0666)
}

func WriteStr(path string, text string) (err error) {
	return Write(path, []uint8(text))
}

func Create(path string) (err error) {
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

func Delete(path string) (err error) {
	return os.Remove(path)
}

func Clear(path string) (err error) {
	return WriteStr(path, "")
}

func Append(path string, bytes []uint8) (err error) {
	var file *os.File
	file, err = os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
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

func AppendStr(path string, text string) (err error) {
	return Append(path, []uint8(text))
}









