//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package convert


import (
	"testing"
)


func TestStringArrayToMergedStringArray(t *testing.T) {
	var words []string
	mark := rune('\'')

	words = append(words, "'asdf", "bla'", "'", "hallo'", "''", "ag'", "'ps", "pp's")
	mergedWords := StringArrayToMergedStringArray(words, mark)

	if 5!=len(mergedWords) {
		t.Errorf("Words number isn't five: %d", len(mergedWords))
	}

	if "asdf bla"!=mergedWords[0] {
		t.Errorf("Wrong merged word: %s", mergedWords[0])
	}

	if " hallo"!=mergedWords[1] {
		t.Errorf("Wrong merged word: %s", mergedWords[0])
	}

	if "pp's"!=mergedWords[4] {
		t.Errorf("Wrong merged word: %s", mergedWords[0])
	}
}













