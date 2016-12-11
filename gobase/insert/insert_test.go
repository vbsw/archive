//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package insert


import (
	"testing"
)


const (
	STRING_A = "hallo"
	STRING_B = "tschüss"
	STRING_C = "kilo"
	STRING_AB = "hallotschüss"
	STRING_ACB = "hallokilotschüss"
)


func TestByteWord(t *testing.T) {
	var emptyWords []byte

	words := ByteWord(emptyWords, 0, []byte(STRING_A))
	if len(STRING_A)!=len(words) {
		t.Errorf("Wrong word size: %d", len(words))
	}

	words = ByteWord(words, len(STRING_A), []byte(STRING_B))
	if len(STRING_A)+len(STRING_B)!=len(words) {
		t.Errorf("Wrong word size: %d", len(words))
	}
	if STRING_AB!=string(words) {
		t.Errorf("Wrong words: %s", words)
	}

	words = ByteWord(words, len(STRING_A), []byte(STRING_C))
	if len(STRING_A)+len(STRING_B)+len(STRING_C)!=len(words) {
		t.Errorf("Wrong word size: %d", len(words))
	}
	if STRING_ACB!=string(words) {
		t.Errorf("Wrong words: %s", words)
	}
}

func TestInteger(t *testing.T) {
	var integers []int
	
	integers = Integer(integers, 0, 1)
	if 0<len(integers) {
		if 1!=len(integers) {
			t.Errorf("Wrong array length: %d", len(integers))
		}
		if 1!=integers[0] {
			t.Errorf("Wrong array length: %d", len(integers))
		}
	} else {
		t.Errorf("Array is empty")
	}
}

func TestIntegerIncrementSuccessors(t *testing.T) {
	//
}














