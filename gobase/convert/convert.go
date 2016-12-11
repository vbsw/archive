//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package convert


import (
	"bytes"
	"unicode/utf8"
)


const (
	MASK_4th_BYTE = 0xff000000
	MASK_3rd_BYTE = 0x00ff0000
	MASK_2nd_BYTE = 0x0000ff00
	MASK_1st_BYTE = 0x000000ff
)


// Converts 4 bytes to an integer.
func ByteArrayToIntArray(dstIntegers []int, srcBytes []byte) {
	var bytesOffset uint

	for i := range dstIntegers {
		integer := int(srcBytes[bytesOffset])
		bytesOffset++
		integer = (integer<<8) + int(srcBytes[bytesOffset])
		bytesOffset++
		integer = (integer<<8) + int(srcBytes[bytesOffset])
		bytesOffset++
		integer = (integer<<8) + int(srcBytes[bytesOffset])
		bytesOffset++
		dstIntegers[i] = integer
	}
}

// Converts an integer to 4 bytes.
func IntArrayToByteArray(dstBytes []byte, srcIntegers []int) {
	var bytesOffset uint

	for _, integer := range srcIntegers {
		dstBytes[bytesOffset] = byte((integer&MASK_4th_BYTE) >> 24)
		bytesOffset++
		dstBytes[bytesOffset] = byte((integer&MASK_3rd_BYTE) >> 16)
		bytesOffset++
		dstBytes[bytesOffset] = byte((integer&MASK_2nd_BYTE) >> 8)
		bytesOffset++
		dstBytes[bytesOffset] = byte(integer&MASK_1st_BYTE)
		bytesOffset++
	}
}

// Converts an integer to 4 bytes.
func IntToByteArray(dstBytes []byte, srcInteger int) {
	dstBytes[0] = byte((srcInteger&MASK_4th_BYTE) >> 24)
	dstBytes[1] = byte((srcInteger&MASK_3rd_BYTE) >> 16)
	dstBytes[2] = byte((srcInteger&MASK_2nd_BYTE) >> 8)
	dstBytes[3] = byte(srcInteger&MASK_1st_BYTE)
}

// Reads an integer as 4 bytes.
func ByteArrayToInt(bytes []byte) (integer int) {
	integer = int(bytes[0])
	integer = (integer<<8) + int(bytes[1])
	integer = (integer<<8) + int(bytes[2])
	integer = (integer<<8) + int(bytes[3])
	return integer
}

func StringArrayToByteArray(strings []string) (bytes [][]byte) {
	bytes = make([][]byte, len(strings))
	for i, str := range strings {
		bytes[i] = []byte(str)
	}
	return bytes
}

func ByteArrayToStringArray(bytes [][]byte) (strings []string) {
	strings = make([]string, len(bytes))
	for i, byteWord := range bytes {
		strings[i] = string(byteWord)
	}
	return strings
}

// merges strings enclosed by mark
func StringArrayToMergedStringArray(words []string, mark rune) (mergedWords []string) {
	var mergedWord bytes.Buffer
	wordsLength := len(words)
	startMergeWordIndex := -1
	startMergeRuneLength := -1
	mergedWords = make([]string, 0, wordsLength)

	for i, word := range words {

		// not merging currently
		if 0>startMergeWordIndex {
			runeAtWordBegin, runeAtWordBeginLength := utf8.DecodeRuneInString(word)

			// start merging
			if mark==runeAtWordBegin {
				var runeAtWordEnd rune
				var runeAtWordEndIndex int
				restWord := word[runeAtWordBeginLength:]

				for j, r := range restWord {
					runeAtWordEnd = r
					runeAtWordEndIndex = j
				}

				// end merging
				if mark==runeAtWordEnd {
					restWord = restWord[:runeAtWordEndIndex]

					if 0<len(restWord) {
						mergedWords = append(mergedWords, restWord)
					}

				} else {
					startMergeWordIndex = i
					startMergeRuneLength = runeAtWordBeginLength
				}

			// append word
			} else {
				mergedWords = append(mergedWords, word)
			}

		// merging currently
		} else {
			var runeAtWordEnd rune
			var runeAtWordEndLength int

			for j, r := range word {
				runeAtWordEnd = r
				runeAtWordEndLength = j
			}

			// end merging
			if mark==runeAtWordEnd {
				wordsToMerge := words[startMergeWordIndex:i+1]
				mergedWordLength := -1
				runeAtWordEndLength = len(word)-runeAtWordEndLength

				for _, w := range wordsToMerge {
					mergedWordLength += 1+len(w)
				}
				mergedWord.WriteString(wordsToMerge[0])
				wordsToMerge = wordsToMerge[1:]

				for _, w := range wordsToMerge {
					mergedWord.WriteString(" ")
					mergedWord.WriteString(w)
				}
				mergedWordStr := mergedWord.String()
				mergedWords = append(mergedWords, mergedWordStr[startMergeRuneLength:len(mergedWordStr)-runeAtWordEndLength])
				startMergeWordIndex = -1
				startMergeRuneLength = -1
				mergedWord.Reset()
			}
		}
	}
	
	// add not merged words
	if 0<=startMergeWordIndex {
		wordsToMerge := words[startMergeWordIndex:]
		for _, w := range wordsToMerge {
			mergedWords = append(mergedWords, w)
		}
	}
	return mergedWords
}

func StringToEscapedString(text string) string {
	const (
		CTX_ESCAPE = 0
		CTX_OTHER = 1
		CHAR_BACK_SLASH = '\\'
		CHAR_N = 'n'
		CHAR_T = 't'
		CHAR_R = 'r'
		CHAR_NEW_LINE = '\n'
		CHAR_TAB = '\t'
		CHAR_CARRIAGE_RETURN = '\r'
	)

	if 1<len(text) {
		var buffer bytes.Buffer
		ctx := CTX_OTHER
		textBytes := []byte(text)

		for _, b := range textBytes {

			if CTX_ESCAPE==ctx {

				if CHAR_BACK_SLASH==b {
					ctx = CTX_OTHER
					buffer.WriteByte(CHAR_BACK_SLASH)
				} else {
					if CHAR_N==b {
						ctx = CTX_OTHER
						buffer.WriteByte(CHAR_NEW_LINE)
					} else if CHAR_T==b {
						ctx = CTX_OTHER
						buffer.WriteByte(CHAR_TAB)
					} else if CHAR_R==b {
						ctx = CTX_OTHER
						buffer.WriteByte(CHAR_CARRIAGE_RETURN)
					}
				}

			} else {

				if CHAR_BACK_SLASH==b {
					ctx = CTX_ESCAPE
				} else {
					buffer.WriteByte(b)
				}
			}
		}
		return buffer.String()

	} else {
		return text
	}
}














