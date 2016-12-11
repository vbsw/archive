/*
 * This is free and unencumbered software released into the public domain.
 * 
 * Anyone is free to copy, modify, publish, use, compile, sell, or
 * distribute this software, either in source code form or as a compiled
 * binary, for any purpose, commercial or non-commercial, and by any
 * means.
 * 
 * In jurisdictions that recognize copyright laws, the author or authors
 * of this software dedicate any and all copyright interest in the
 * software to the public domain. We make this dedication for the benefit
 * of the public at large and to the detriment of our heirs and
 * successors. We intend this dedication to be an overt act of
 * relinquishment in perpetuity of all present and future rights to this
 * software under copyright law.
 * 
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 * IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
 * OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
 * ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
 * OTHER DEALINGS IN THE SOFTWARE.
 * 
 * For more information, please refer to <http://unlicense.org>
 */


package main

import (
	"os"
	"strings"
	"strconv"
)




var (
	argFile string
	argW int
	argH int
	argC int
	argDelay int
	argReplayDelay int
)

func argsVersion() bool {
	return   2 == len(os.Args) &&
	       ( 0 == strings.Compare(os.Args[1], "version") ||
	         0 == strings.Compare(os.Args[1], "--version") )
}

func argsHelp() bool {
	return   2 == len(os.Args) &&
	       ( 0 == strings.Compare(os.Args[1], "help") ||
	         0 == strings.Compare(os.Args[1], "--help") )
}

func argsControls() bool {
	return   2 == len(os.Args) &&
	       ( 0 == strings.Compare(os.Args[1], "controls") ||
	         0 == strings.Compare(os.Args[1], "--controls") )
}

func argsExamples() bool {
	return   2 == len(os.Args) &&
	       ( 0 == strings.Compare(os.Args[1], "examples") ||
	         0 == strings.Compare(os.Args[1], "--examples") )
}

func argsAnimation() bool {
	argc := len(os.Args)

	if 4 == argc || 5 == argc {
		nums := strings.Split(os.Args[2], "x")
		argFile = os.Args[1]

		// WxH
		if 2 == len(nums) {
			var num int
			var err error

			num, err = strconv.Atoi(nums[0])

			if nil == err && 0 < num {
				argW = num
				num, err = strconv.Atoi(nums[1])

				if nil == err && 0 < num {
					argH = num
					argC = argW * argH
					num, err = strconv.Atoi(os.Args[3])

					if nil == err && 0 < num {
						argDelay = num

						if 5 == argc {
							num, err = strconv.Atoi(os.Args[4])

							if nil == err && 0 < num {
								argReplayDelay = num

								return true
							}

						} else {
							argReplayDelay = argDelay

							return true
						}
					}
				}
			}

		// WxHxC
		} else if 3 == len(nums) {
			var num int
			var err error

			num, err = strconv.Atoi(nums[0])

			if nil == err && 0 < num {
				argW = num
				num, err = strconv.Atoi(nums[1])

				if nil == err && 0 < num {
					argH = num
					num, err = strconv.Atoi(nums[2])

					if nil == err && 0 < num {
						argC = num
						num, err = strconv.Atoi(os.Args[3])

						if nil == err && 0 < num {
							argDelay = num

							if 5 == argc {
								num, err = strconv.Atoi(os.Args[4])

								if nil == err && 0 < num {
									argReplayDelay = num

									return true
								}

							} else {
								argReplayDelay = argDelay

								return true
							}
						}
					}
				}
			}
		}
	}
	return false
}



