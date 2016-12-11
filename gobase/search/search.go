//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package search


import (
	"bytes"
)


func IntegerBinary(integers []int, integer int) (integerIndex int, match bool) {
	low := 0
	high := len(integers)-1

	for low<=high && (!match) {
		integerIndex = low + ((high-low)/2) // middle
		otherInteger := integers[integerIndex]

		if integer<otherInteger {
			high = integerIndex-1
		} else if integer>otherInteger {
			integerIndex++ // for the case low > high
			low = integerIndex
		} else {
			match = true
		}
	}

	return integerIndex, match
}

func SetsIntegerBinary(setsOffsets, setsLengths, sets []int, setIndex, integer int) (absoluteOffset int, match bool) {
	setOffset := setsOffsets[setIndex]
	setLength := setsLengths[setIndex]
	set := sets[setOffset:setOffset+setLength]
	absoluteOffset, match = IntegerBinary(set, integer)
	absoluteOffset += setOffset
	return absoluteOffset, match
}


// Searches for an index where word matches one of the entries in words.
func ByteWordBinary(wordsOffsets, wordsLengths []int, words, word []byte) (wordIndex int, match bool) {
	low := 0
	high := len(wordsOffsets)-1

	for low<=high && (!match) {
		wordIndex = low + ((high-low)/2) // middle
		wordOffset := wordsOffsets[wordIndex]
		wordLength := wordsLengths[wordIndex]
		comparison := bytes.Compare(word, words[wordOffset:wordOffset+wordLength])

		if 0>comparison {
			high = wordIndex-1
		} else if 0<comparison {
			wordIndex++ // for the case low > high
			low = wordIndex
		} else {
			match = true
		}
	}

	return wordIndex, match
}

// Searches for an index where word matches one of the entries in words.
func ByteWordLinear(wordsLengths []int, words, word []byte) (wordIndex int, match bool) {
	wordOffset := 0

	for _, wordLength := range wordsLengths {
		comparison := bytes.Compare(word, words[wordOffset:wordOffset+wordLength])

		if 0<comparison {
			wordOffset += wordLength
			wordIndex++
		} else {
			match = comparison==0
			break
		}
	}

	return wordIndex, match
}














