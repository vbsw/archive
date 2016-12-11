//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package search


import (
	"testing"
)


const (
	STRING_A = "hallo"
	STRING_B = "tschüss"
	STRING_C = "kilo"
	STRING_AB = "hallotschüss"
)


func TestIntegerBinary(t *testing.T) {
	//
}

func TestSetsIntegerBinary(t *testing.T) {
	//
}

func TestByteWordBinary(t *testing.T) {
	wordsOffsets := []int { 0, len(STRING_A) }
	wordsLengths := []int { len(STRING_A), len(STRING_B) }
	words := []byte(STRING_AB)

	wordIndex, wordMatch := ByteWordBinary(wordsOffsets, wordsLengths, words, []byte(STRING_A))
	if 0!=wordIndex {
		t.Errorf("Wrong word index: %d", wordIndex)
	}
	if !wordMatch {
		t.Errorf("Word not found.")
	}

	wordIndex, wordMatch = ByteWordBinary(wordsOffsets, wordsLengths, words, []byte(STRING_B))
	if 1!=wordIndex {
		t.Errorf("Wrong word index: %d", wordIndex)
	}
	if !wordMatch {
		t.Errorf("Word not found.")
	}

	wordIndex, wordMatch = ByteWordBinary(wordsOffsets, wordsLengths, words, []byte(STRING_C))
	if 1!=wordIndex {
		t.Errorf("Wrong word index: %d", wordIndex)
	}
	if wordMatch {
		t.Errorf("Word was found.")
	}
}

func TestByteWordLinear(t *testing.T) {
	wordsLengths := []int { len(STRING_A), len(STRING_B) }
	words := []byte(STRING_AB)

	wordIndex, wordMatch := ByteWordLinear(wordsLengths, words, []byte(STRING_A))
	if 0!=wordIndex {
		t.Errorf("Wrong word index: %d", wordIndex)
	}
	if !wordMatch {
		t.Errorf("Word not found.")
	}

	wordIndex, wordMatch = ByteWordLinear(wordsLengths, words, []byte(STRING_B))
	if 1!=wordIndex {
		t.Errorf("Wrong word index: %d", wordIndex)
	}
	if !wordMatch {
		t.Errorf("Word not found.")
	}

	wordIndex, wordMatch = ByteWordLinear(wordsLengths, words, []byte(STRING_C))
	if 1!=wordIndex {
		t.Errorf("Wrong word index: %d", wordIndex)
	}
	if wordMatch {
		t.Errorf("Word was found.")
	}
}














