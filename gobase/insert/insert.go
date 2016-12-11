//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package insert


func ByteWord(words []byte, offset int, word []byte) (newWords []byte) {
	wordLength := len(word)
	newWords = make([]byte, len(words)+wordLength)
	copy(newWords, words[:offset])
	copy(newWords[offset:], word)
	copy(newWords[offset+wordLength:], words[offset:])
	return newWords
}

func Integer(integers []int, index, value int) (newIntegers []int) {
	newIntegers = make([]int, 1+len(integers))
	copy(newIntegers, integers[:index])
	newIntegers[index] = value
	copy(newIntegers[1+index:], integers[index:])
	return newIntegers
}

func IntegerIncrementSuccessors(integers []int, index, value, increment int) []int {
	successors := integers[index:]
	for i := range successors {
		successors[i] += increment
	}
	return Integer(integers, index, value)
}














