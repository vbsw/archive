//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package tag


import(
	"bytes"
	"encoding/xml"
	"strconv"
	"strings"
)

import (
	"github.com/vitalibaumtrok/gobase/convert"
	"github.com/vitalibaumtrok/gobase/deallocate"
	"github.com/vitalibaumtrok/gobase/insert"
	"github.com/vitalibaumtrok/gobase/search"
)


const (
	MINIMUM_SIZE = (4+4)*4
	XML_HEADER = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<!DOCTYPE taggedwords SYSTEM \"taggedwords.dtd\">"
)


type Tag struct {
	DescsOffsets []int
	DescsLengths []int
	Descs []byte
	WordsOffsets []int
	WordsLengths []int
	Words []byte
	TagsOffsets []int
	TagsLengths []int
	Tags []byte
	MapsOffsets []int
	MapsLengths []int
	Maps []int
}

type XMLTaggedWords struct {
	XMLName xml.Name `xml:"taggedwords"`
	Words []XMLWord `xml:"word"`
}

type XMLWord struct {
	XMLName xml.Name `xml:"word"`
	Desc string `xml:",chardata"`
	Name string `xml:"name,attr"`
	Tags string `xml:"tags,attr"`
}


func NewStruct(wordsOffsets, wordsLengths []int, words []byte, descsOffsets, descsLengths []int, descs []byte, tagsOffsets, tagsLengths []int, tags []byte, mapsOffsets, mapsLengths, maps []int) (t Tag) {
	t.DescsOffsets = descsOffsets
	t.DescsLengths = descsLengths
	t.Descs = descs
	t.WordsOffsets = wordsOffsets
	t.WordsLengths = wordsLengths
	t.Words = words
	t.TagsOffsets = tagsOffsets
	t.TagsLengths = tagsLengths
	t.Tags = tags
	t.MapsOffsets = mapsOffsets
	t.MapsLengths = mapsLengths
	t.Maps = maps
	return t
}

func NewRef(wordsOffsets, wordsLengths []int, words []byte, descsOffsets, descsLengths []int, descs []byte, tagsOffsets, tagsLengths []int, tags []byte, mapsOffsets, mapsLengths, maps []int) (t *Tag) {
	t = new(Tag)
	t.DescsOffsets = descsOffsets
	t.DescsLengths = descsLengths
	t.Descs = descs
	t.WordsOffsets = wordsOffsets
	t.WordsLengths = wordsLengths
	t.Words = words
	t.TagsOffsets = tagsOffsets
	t.TagsLengths = tagsLengths
	t.Tags = tags
	t.MapsOffsets = mapsOffsets
	t.MapsLengths = mapsLengths
	t.Maps = maps
	return t
}

func Init(encodedData []byte) []byte {
	if MINIMUM_SIZE>len(encodedData) {
		encodedData = make([]byte, MINIMUM_SIZE)
		// 4 is the size of a part
		convert.IntToByteArray(encodedData, 4)
		convert.IntToByteArray(encodedData[4:], 4)
		convert.IntToByteArray(encodedData[8:], 4)
		convert.IntToByteArray(encodedData[12:], 4)
	}
	return encodedData
}

func (this *Tag) SetDesc(descIndex int, oDesc []byte) {
	descOffset := this.DescsOffsets[descIndex]
	descLength := this.DescsLengths[descIndex]
	oDescLength := len(oDesc)

	if descLength==oDescLength {
		copy(this.Descs[descOffset:], oDesc)
	} else {
		diff := oDescLength-descLength
		nextDescOffset := descOffset+descLength
		successorDescsOffsets := this.DescsOffsets[descIndex+1:]
		successorDescs := this.Descs[nextDescOffset:]
		// adjust length
		this.DescsLengths[descIndex] = oDescLength

		// adjust offsets
		for i := range successorDescsOffsets {
			successorDescsOffsets[i] += diff
		}

		// expand array
		if 0<diff {
			this.Descs = append(this.Descs, oDesc[:diff]...)

		// shrink array
		} else {
			this.Descs = this.Descs[:len(this.Descs)+diff]
		}
		copy(this.Descs[nextDescOffset+diff:], successorDescs)
		copy(this.Descs[descOffset:], oDesc)
	}
}

func (this *Tag) DescByIndex(descIndex int) []byte {
	descOffset := this.DescsOffsets[descIndex]
	descLength := this.DescsLengths[descIndex]
	return this.Descs[descOffset:descOffset+descLength]
}

func (this *Tag) WordsArray() (words [][]byte) {
	words = make([][]byte, len(this.WordsOffsets))
	for i, wordOffset := range this.WordsOffsets {
		wordLength := this.WordsLengths[i]
		words[i] = this.Words[wordOffset:wordOffset+wordLength]
	}
	return words
}

func (this *Tag) WordByIndex(wordIndex int) []byte {
	wordOffset := this.WordsOffsets[wordIndex]
	wordLength := this.WordsLengths[wordIndex]
	return this.Words[wordOffset:wordOffset+wordLength]
}

func (this *Tag) TagsArray() (tags [][]byte) {
	tags = make([][]byte, len(this.TagsOffsets))
	for i, tagOffset := range this.TagsOffsets {
		tagLength := this.TagsLengths[i]
		tags[i] = this.Tags[tagOffset:tagOffset+tagLength]
	}
	return tags
}

func (this *Tag) TagByIndex(tagIndex int) []byte {
	tagOffset := this.TagsOffsets[tagIndex]
	tagLength := this.TagsLengths[tagIndex]
	return this.Tags[tagOffset:tagOffset+tagLength]
}

func (this *Tag) TagsByWordIndex(wordIndex int) (tags [][]byte) {
	var tagsIndices []int

	for i, mapOffset := range this.MapsOffsets {
		mapLength := this.MapsLengths[i]
		mapp := this.Maps[mapOffset:mapOffset+mapLength]
		_, refMatch := search.IntegerBinary(mapp, wordIndex)

		if refMatch {
			tagsIndices = append(tagsIndices, i)
		}
	}

	tags = make([][]byte, len(tagsIndices))
	for i := range tags {
		tags[i] = this.TagByIndex(tagsIndices[i])
	}

	return tags
}

func (this *Tag) MapByIndex(mapIndex int) []int {
	mapOffset := this.MapsOffsets[mapIndex]
	mapLength := this.MapsLengths[mapIndex]
	return this.Maps[mapOffset:mapOffset+mapLength]
}

func (this *Tag) EncodedSize() int {
	// encoded parts lengths: 4*4
	// encoded numbers of descs/words/tags/maps: 4*4
	// number of descs/words: 2*(2*byteLength)
	// number of tags/maps: 2*(2*byteLength)
	return 4*4 + 4*4 + 2*(2*4*len(this.DescsOffsets)) + 2*(2*4*len(this.TagsOffsets)) + len(this.Descs) + len(this.Words) + len(this.Tags) + 4*len(this.Maps)
}

func (this *Tag) Unmarshal(encodedData []byte) {
	// find out the sizes
	encodedDescsSize := convert.ByteArrayToInt(encodedData)
	encodedWordsSize := convert.ByteArrayToInt(encodedData[4:])
	encodedTagsSize := convert.ByteArrayToInt(encodedData[8:])
	encodedMapsSize := convert.ByteArrayToInt(encodedData[12:])
	encodedDescs := encodedData[16:]
	encodedWords := encodedData[16+encodedDescsSize:]
	encodedTags := encodedData[16+encodedDescsSize+encodedWordsSize:]
	encodedMaps := encodedData[16+encodedDescsSize+encodedWordsSize+encodedTagsSize:]
	descsNumber := convert.ByteArrayToInt(encodedDescs)
	wordsNumber := convert.ByteArrayToInt(encodedWords)
	tagsNumber := convert.ByteArrayToInt(encodedTags)
	mapsNumber := convert.ByteArrayToInt(encodedMaps)
	// 2*4, because indices and lengths are first
	descsOffset := 4 + 2*4*descsNumber
	wordsOffset := 4 + 2*4*wordsNumber
	tagsOffset := 4 + 2*4*tagsNumber
	mapsOffset := 4 + 2*4*mapsNumber

	// allocate memory
	this.DescsOffsets = make([]int, descsNumber)
	this.DescsLengths = make([]int, descsNumber)
	this.Descs = make([]byte, encodedDescsSize-descsOffset)
	this.WordsOffsets = make([]int, wordsNumber)
	this.WordsLengths = make([]int, wordsNumber)
	this.Words = make([]byte, encodedWordsSize-wordsOffset)
	this.TagsOffsets = make([]int, tagsNumber)
	this.TagsLengths = make([]int, tagsNumber)
	this.Tags = make([]byte, encodedTagsSize-tagsOffset)
	this.MapsOffsets = make([]int, mapsNumber)
	this.MapsLengths = make([]int, mapsNumber)
	this.Maps = make([]int, (encodedMapsSize-mapsOffset) / 4)

	// convert data
	convert.ByteArrayToIntArray(this.DescsOffsets, encodedDescs[4:])
	convert.ByteArrayToIntArray(this.DescsLengths, encodedDescs[4+4*descsNumber:])
	copy(this.Descs, encodedDescs[descsOffset:])
	convert.ByteArrayToIntArray(this.WordsOffsets, encodedWords[4:])
	convert.ByteArrayToIntArray(this.WordsLengths, encodedWords[4+4*wordsNumber:])
	copy(this.Words, encodedWords[wordsOffset:])
	convert.ByteArrayToIntArray(this.TagsOffsets, encodedTags[4:])
	convert.ByteArrayToIntArray(this.TagsLengths, encodedTags[4+4*tagsNumber:])
	copy(this.Tags, encodedTags[tagsOffset:])
	convert.ByteArrayToIntArray(this.MapsOffsets, encodedMaps[4:])
	convert.ByteArrayToIntArray(this.MapsLengths, encodedMaps[4+4*mapsNumber:])
	convert.ByteArrayToIntArray(this.Maps, encodedMaps[mapsOffset:])
}

func (this *Tag) Marshal(encodedData []byte) {
	// find out the sizes
	descsNumber := len(this.DescsOffsets)
	wordsNumber := len(this.WordsOffsets)
	tagsNumber := len(this.TagsOffsets)
	mapsNumber := len(this.MapsOffsets)
	// 2*4, because indices and lengths are first
	descsOffset := 4 + 2*4*descsNumber
	wordsOffset := 4 + 2*4*wordsNumber
	tagsOffset := 4 + 2*4*tagsNumber
	mapsOffset := 4 + 2*4*mapsNumber
	encodedDescsSize := descsOffset+len(this.Descs)
	encodedWordsSize := wordsOffset+len(this.Words)
	encodedTagsSize := tagsOffset+len(this.Tags)
	encodedMapsSize := mapsOffset + 4*len(this.Maps)
	encodedDescs := encodedData[16:]
	encodedWords := encodedData[16+encodedDescsSize:]
	encodedTags := encodedData[16+encodedDescsSize+encodedWordsSize:]
	encodedMaps := encodedData[16+encodedDescsSize+encodedWordsSize+encodedTagsSize:]

	// convert part sizes
	convert.IntToByteArray(encodedData, encodedDescsSize)
	convert.IntToByteArray(encodedData[4:], encodedWordsSize)
	convert.IntToByteArray(encodedData[8:], encodedTagsSize)
	convert.IntToByteArray(encodedData[12:], encodedMapsSize)
	// convert data
	convert.IntToByteArray(encodedDescs, descsNumber)
	convert.IntArrayToByteArray(encodedDescs[4:], this.DescsOffsets)
	convert.IntArrayToByteArray(encodedDescs[4+4*descsNumber:], this.DescsLengths)
	copy(encodedDescs[descsOffset:], this.Descs)
	convert.IntToByteArray(encodedWords, wordsNumber)
	convert.IntArrayToByteArray(encodedWords[4:], this.WordsOffsets)
	convert.IntArrayToByteArray(encodedWords[4+4*wordsNumber:], this.WordsLengths)
	copy(encodedWords[wordsOffset:], this.Words)
	convert.IntToByteArray(encodedTags, tagsNumber)
	convert.IntArrayToByteArray(encodedTags[4:], this.TagsOffsets)
	convert.IntArrayToByteArray(encodedTags[4+4*tagsNumber:], this.TagsLengths)
	copy(encodedTags[tagsOffset:], this.Tags)
	convert.IntToByteArray(encodedMaps, mapsNumber)
	convert.IntArrayToByteArray(encodedMaps[4:], this.MapsOffsets)
	convert.IntArrayToByteArray(encodedMaps[4+4*mapsNumber:], this.MapsLengths)
	convert.IntArrayToByteArray(encodedMaps[mapsOffset:], this.Maps)
}

// adds directly description data (maps are not modified!)
func (this *Tag) AddDesc(descIndex int, desc []byte) {
	var descOffset int
	descLength := len(desc)
	if 0<descIndex {
		if len(this.DescsOffsets)==descIndex {
			prevIndex := descIndex-1
			descOffset = this.DescsOffsets[prevIndex]+this.DescsLengths[prevIndex]
		} else {
			descOffset = this.DescsOffsets[descIndex]
		}
	} else {
		descOffset = 0
	}
	// insert description
	this.Descs = insert.ByteWord(this.Descs, descOffset, desc)
	this.DescsOffsets = insert.IntegerIncrementSuccessors(this.DescsOffsets, descIndex, descOffset, descLength)
	this.DescsLengths = insert.Integer(this.DescsLengths, descIndex, descLength)
}

// removes directly description data (maps are not modified!)
func (this *Tag) RemoveDesc(descIndex int) {
	descOffset := this.DescsOffsets[descIndex]
	descLength := this.DescsLengths[descIndex]
	this.Descs = deallocate.BytesFromArray(this.Descs, descOffset, descLength)
	this.DescsOffsets = deallocate.IntegerFromArray(this.DescsOffsets, descIndex)
	this.DescsLengths = deallocate.IntegerFromArray(this.DescsLengths, descIndex)
	successorDescsOffsets := this.DescsOffsets[descIndex:]

	// adjust offsets
	for i := range successorDescsOffsets {
		successorDescsOffsets[i] -= descLength
	}
}

// adds directly word data (maps are not modified!)
func (this *Tag) AddWord(wordIndex int, word []byte) {
	var wordOffset int
	wordLength := len(word)
	if 0<wordIndex {
		if len(this.WordsOffsets)==wordIndex {
			prevIndex := wordIndex-1
			wordOffset = this.WordsOffsets[prevIndex]+this.WordsLengths[prevIndex]
		} else {
			wordOffset = this.WordsOffsets[wordIndex]
		}
	} else {
		wordOffset = 0
	}
	// insert word
	this.Words = insert.ByteWord(this.Words, wordOffset, word)
	this.WordsOffsets = insert.IntegerIncrementSuccessors(this.WordsOffsets, wordIndex, wordOffset, wordLength)
	this.WordsLengths = insert.Integer(this.WordsLengths, wordIndex, wordLength)
}

// removes directly word data (maps are not modified!)
func (this *Tag) RemoveWord(wordIndex int) {
	wordOffset := this.WordsOffsets[wordIndex]
	wordLength := this.WordsLengths[wordIndex]
	this.Words = deallocate.BytesFromArray(this.Words, wordOffset, wordLength)
	this.WordsOffsets = deallocate.IntegerFromArray(this.WordsOffsets, wordIndex)
	this.WordsLengths = deallocate.IntegerFromArray(this.WordsLengths, wordIndex)
	successorWordsOffsets := this.WordsOffsets[wordIndex:]

	// adjust offsets
	for i := range successorWordsOffsets {
		successorWordsOffsets[i] -= wordLength
	}
}

// adds directly tag data (maps are not modified!)
func (this *Tag) AddTag(tagIndex int, tagg []byte) {
	var tagOffset int
	tagLength := len(tagg)
	if 0<tagIndex {
		if len(this.TagsOffsets)==tagIndex {
			prevIndex := tagIndex-1
			tagOffset = this.TagsOffsets[prevIndex]+this.TagsLengths[prevIndex]
		} else {
			tagOffset = this.TagsOffsets[tagIndex]
		}
	} else {
		tagOffset = 0
	}
	// insert tag
	this.Tags = insert.ByteWord(this.Tags, tagOffset, tagg)
	this.TagsOffsets = insert.IntegerIncrementSuccessors(this.TagsOffsets, tagIndex, tagOffset, tagLength)
	this.TagsLengths = insert.Integer(this.TagsLengths, tagIndex, tagLength)
}

// removes directly tag data (maps are not modified!)
func (this *Tag) RemoveTag(tagIndex int) {
	tagOffset := this.TagsOffsets[tagIndex]
	tagLength := this.TagsLengths[tagIndex]
	this.Tags = deallocate.BytesFromArray(this.Tags, tagOffset, tagLength)
	this.TagsOffsets = deallocate.IntegerFromArray(this.TagsOffsets, tagIndex)
	this.TagsLengths = deallocate.IntegerFromArray(this.TagsLengths, tagIndex)
	successorTagsOffsets := this.TagsOffsets[tagIndex:]

	// adjust offsets
	for i := range successorTagsOffsets {
		successorTagsOffsets[i] -= tagLength
	}
}

func (this *Tag) AddEmptyMap(mapIndex int) {
	var mapOffset int
	if 0<mapIndex {
		if len(this.MapsOffsets)==mapIndex {
			prevIndex := mapIndex-1
			mapOffset = this.MapsOffsets[prevIndex]+this.MapsLengths[prevIndex]
		} else {
			mapOffset = this.MapsOffsets[mapIndex]
		}
	} else {
		mapOffset = 0
	}
	this.MapsOffsets = insert.Integer(this.MapsOffsets, mapIndex, mapOffset)
	this.MapsLengths = insert.Integer(this.MapsLengths, mapIndex, 0)
}

func (this *Tag) AddRef(refAbsoluteOffset, mapIndex, ref int) {
	successorMapOffsets := this.MapsOffsets[mapIndex+1:]
	this.Maps = insert.Integer(this.Maps, refAbsoluteOffset, ref)
	this.MapsLengths[mapIndex]++

	for i := range successorMapOffsets {
		successorMapOffsets[i]++
	}
}

// increment referenced word indices equal to or greater then boundary
func (this *Tag) IncrementRefValues(boundary int) {
	for i, refWordIndex := range this.Maps {
		if refWordIndex>=boundary {
			this.Maps[i]++
		}
	}
}

// decrement referenced word indices greater then boundary
func (this *Tag) DecrementRefValues(boundary int) {
	for i, refWordIndex := range this.Maps {
		if refWordIndex>boundary {
			this.Maps[i]--
		}
	}
}

func (this *Tag) AddWordAndTags(word []byte, tags [][]byte) (added bool) {
	wordIndex, wordMatch := search.ByteWordBinary(this.WordsOffsets, this.WordsLengths, this.Words, word)

	if !wordMatch {
		added = true
		this.AddDesc(wordIndex, []byte(""))
		this.AddWord(wordIndex, word)
		this.IncrementRefValues(wordIndex)
	}

	for _, tagg := range tags {
		tagIndex, tagMatch := search.ByteWordBinary(this.TagsOffsets, this.TagsLengths, this.Tags, tagg)

		if !tagMatch {
			this.AddTag(tagIndex, tagg)
			this.AddEmptyMap(tagIndex)
		}

		// tagIndex is mapIndex
		refAbsoluteOffset, refMatch := search.SetsIntegerBinary(this.MapsOffsets, this.MapsLengths, this.Maps, tagIndex, wordIndex)

		if !refMatch {
			added = true
			this.AddRef(refAbsoluteOffset, tagIndex, wordIndex)
		}
	}

	return added
}

func (this *Tag) AddWordAndTagsSetDesc(word []byte, tags [][]byte, desc []byte) {
	wordIndex, wordMatch := search.ByteWordBinary(this.WordsOffsets, this.WordsLengths, this.Words, word)

	if wordMatch {
		this.SetDesc(wordIndex, desc)
	} else {
		this.AddDesc(wordIndex, desc)
		this.AddWord(wordIndex, word)
		this.IncrementRefValues(wordIndex)
	}

	for _, tagg := range tags {
		tagIndex, tagMatch := search.ByteWordBinary(this.TagsOffsets, this.TagsLengths, this.Tags, tagg)

		if !tagMatch {
			this.AddTag(tagIndex, tagg)
			this.AddEmptyMap(tagIndex)
		}

		// tagIndex is mapIndex
		refAbsoluteOffset, refMatch := search.SetsIntegerBinary(this.MapsOffsets, this.MapsLengths, this.Maps, tagIndex, wordIndex)

		if !refMatch {
			this.AddRef(refAbsoluteOffset, tagIndex, wordIndex)
		}
	}
}

func (this *Tag) RemoveTags(word []byte, tags [][]byte) (removed int) {
	wordIndex, wordMatch := search.ByteWordBinary(this.WordsOffsets, this.WordsLengths, this.Words, word)

	if wordMatch {

		for _, tagg := range tags {
			tagIndex, tagMatch := search.ByteWordBinary(this.TagsOffsets, this.TagsLengths, this.Tags, tagg)

			if tagMatch {
				mapp := this.MapByIndex(tagIndex)
				refIndex, refMatch := search.IntegerBinary(mapp, wordIndex)

				if refMatch {
					removed++
					absoluteRefOffset := this.MapsOffsets[tagIndex]+refIndex
					successorMapsOffsets := this.MapsOffsets[tagIndex+1:]
					this.Maps = deallocate.IntegerFromArray(this.Maps, absoluteRefOffset)
					// adjust lengths
					this.MapsLengths[tagIndex]--

					// adjust offsets
					for i := range successorMapsOffsets {
						successorMapsOffsets[i]--
					}

					if 1==len(mapp) {
						// remove map
						this.MapsOffsets = deallocate.IntegerFromArray(this.MapsOffsets, tagIndex)
						this.MapsLengths = deallocate.IntegerFromArray(this.MapsLengths, tagIndex)
						// remove tag
						tagOffset := this.TagsOffsets[tagIndex]
						tagLength := this.TagsLengths[tagIndex]
						this.Tags = deallocate.BytesFromArray(this.Tags, tagOffset, tagLength)
						this.TagsOffsets = deallocate.IntegerFromArray(this.TagsOffsets, tagIndex)
						this.TagsLengths = deallocate.IntegerFromArray(this.TagsLengths, tagIndex)
						successorTagsOffsets := this.TagsOffsets[tagIndex:]

						// adjust offsets
						for i := range successorTagsOffsets {
							successorTagsOffsets[i] -= tagLength
						}
					}
				}
			}
		}
	}
	return removed
}

func (this *Tag) RemoveWords(words [][]byte) (removed int) {
	for _, word := range words {
		wordIndex, wordMatch := search.ByteWordBinary(this.WordsOffsets, this.WordsLengths, this.Words, word)

		if wordMatch {
			removed++
			this.RemoveDesc(wordIndex)
			this.RemoveWord(wordIndex)

			for i:=0; i<len(this.MapsOffsets); i++ {
				mapOffset := this.MapsOffsets[i]
				mapLength := this.MapsLengths[i]
				mapp := this.Maps[mapOffset:mapOffset+mapLength]
				refIndex, refMatch := search.IntegerBinary(mapp, wordIndex)

				if refMatch {
					absoluteRefOffset := mapOffset+refIndex
					successorMapsOffsets := this.MapsOffsets[i+1:]
					this.Maps = deallocate.IntegerFromArray(this.Maps, absoluteRefOffset)
					// adjust lengths
					this.MapsLengths[i]--

					// adjust offsets
					for i := range successorMapsOffsets {
						successorMapsOffsets[i]--
					}

					if 1==len(mapp) {
						// remove map
						this.MapsOffsets = deallocate.IntegerFromArray(this.MapsOffsets, i)
						this.MapsLengths = deallocate.IntegerFromArray(this.MapsLengths, i)
						this.RemoveTag(i)
						i--
					}
				}
			}
		}
	}
	return removed
}

func (this *Tag) ExportTXT() []byte {
	var buffer bytes.Buffer
	lastIndex := len(this.WordsOffsets)-1

	for i, wordOffset := range this.WordsOffsets {
		descOffset := this.DescsOffsets[i]
		descLength := this.DescsLengths[i]
		wordLength := this.WordsLengths[i]
		desc := string(this.Descs[descOffset:descOffset+descLength])
		word := string(this.Words[wordOffset:wordOffset+wordLength])
		tags := convert.ByteArrayToStringArray(this.TagsByWordIndex(i))

		// write line number of word
		if 0<len(word) {
			var wordLines int = 1

			for _, r := range word {
				if '\n'==r {
					wordLines++
				}
			}
			buffer.WriteString(strconv.Itoa(wordLines))
		} else {
			buffer.WriteString("0")
		}

		// write line number of tags
		if 0<len(tags) {
			buffer.WriteString(" 1")
		} else {
			buffer.WriteString(" 0")
		}

		// write line number of description
		if 0<len(desc) {
			var descLines int = 1

			for _, r := range desc {
				if '\n'==r {
					descLines++
				}
			}
			buffer.WriteString(" ")
			buffer.WriteString(strconv.Itoa(descLines))
		} else {
			buffer.WriteString(" 0")
		}
		buffer.WriteString("\n")

		// write word
		buffer.WriteString(word)
		buffer.WriteString("\n")

		// write tags
		if 0<len(tags) {
			successorTags := tags[1:]
			buffer.WriteString(tags[0])

			if 0<len(successorTags) {
				for _, tag := range successorTags {
					buffer.WriteString(" ")
					buffer.WriteString(tag)
				}
			}
			buffer.WriteString("\n")
		}

		// write description
		if 0<len(desc) {
			buffer.WriteString(desc)
			buffer.WriteString("\n")
		}

		if i<lastIndex {
			buffer.WriteString("\n")
		}
	}

	return buffer.Bytes()
}

func (this *Tag) ImportTXT(txtBytes []byte) {
	offset := 0

	// iterate over all bytes
	for i:=0; i<len(txtBytes); i++ {
		b := txtBytes[i]

		// no white space
		if '\r'!=b && '\n'!=b && ' '!=b && '\t'!=b {
			var word []byte
			var tags [][]byte
			var desc []byte
			var linesStr []string

			// read line numbers
			for offset=i; '\n'!=txtBytes[i]; i++ {}
			if '\r'==txtBytes[i-1] {
				linesStr = strings.Split(string(txtBytes[offset:i-1]), " ")
			} else {
				linesStr = strings.Split(string(txtBytes[offset:i]), " ")
			}
			wordLines, _ := strconv.Atoi(linesStr[0])
			tagsLines, _ := strconv.Atoi(linesStr[1])
			descLines, _ := strconv.Atoi(linesStr[2])

			// new line to read word
			i++

			// read word
			for offset=i; 0<wordLines; i++ {
				if '\n'==txtBytes[i] {
					wordLines--
					break
				}
			}
			if '\r'==txtBytes[i-1] {
				word = txtBytes[offset:i-1]
			} else {
				word = txtBytes[offset:i]
			}

			// new line to read tags, description or empty line
			i++

			// read tags
			if 0<tagsLines {
				var tagsStrings []string
				for offset=i; '\n'!=txtBytes[i]; i++ {}
				if '\r'==txtBytes[i-1] {
					tagsStrings = strings.Split(string(txtBytes[offset:i-1]), " ")
				} else {
					tagsStrings = strings.Split(string(txtBytes[offset:i]), " ")
				}
				tags = make([][]byte, len(tagsStrings))
				for i, s := range tagsStrings {
					tags[i] = []byte(s)
				}

				// new line to read description
				i++
			}

			// read description
			if 0<descLines {
				for offset=i; 0<descLines; i++ {
					if '\n'==txtBytes[i] {
						descLines--
						break
					}
				}
				if '\r'==txtBytes[i-1] {
					desc = txtBytes[offset:i-1]
				} else {
					desc = txtBytes[offset:i]
				}
			}

			this.AddWordAndTagsSetDesc(word, tags, desc)
		}
	}
}

func (this *Tag) ExportXML() []byte {
	var buffer bytes.Buffer
	buffer.WriteString(XML_HEADER)

	// custom implementation of xml export
	buffer.WriteString("\n<taggedwords>")
	for i, wordOffset := range this.WordsOffsets {
		descOffset := this.DescsOffsets[i]
		descLength := this.DescsLengths[i]
		desc := string(this.Descs[descOffset:descOffset+descLength])
		wordLength := this.WordsLengths[i]
		word := string(this.Words[wordOffset:wordOffset+wordLength])
		tags := convert.ByteArrayToStringArray(this.TagsByWordIndex(i))
		buffer.WriteString("\n\t<word name=\"")
		buffer.WriteString(word)
		buffer.WriteString("\"")

		if 0<len(tags) {
			successorTags := tags[1:]
			buffer.WriteString(" tags=\"")
			buffer.WriteString(tags[0])

			if 0<len(successorTags) {
				for _, tag := range successorTags {
					buffer.WriteString(" ")
					buffer.WriteString(tag)
				}
			}
			buffer.WriteString("\"")
		}

		if 0<len(desc) {
			buffer.WriteString(">")
			buffer.WriteString(desc)
			buffer.WriteString("</word>")
		} else {
			buffer.WriteString("/>")
		}
	}
	buffer.Write([]byte("\n</taggedwords>"))
	
	// go implementation of xml export
	/*
	var taggedWords XMLTaggedWords
	taggedWords.XMLName.Local = "taggedwords"

	for i, wordOffset := range this.WordsOffsets {
		var xmlWord XMLWord
		descOffset := this.DescsOffsets[i]
		descLength := this.DescsLengths[i]
		xmlWord.Desc = string(this.Descs[descOffset:descOffset+descLength])
		wordLength := this.WordsLengths[i]
		xmlWord.Name = string(this.Words[wordOffset:wordOffset+wordLength])
		tags := convert.ByteArrayToStringArray(this.TagsByWordIndex(i))
		xmlWord.Tags = strings.Join(tags, " ")
		taggedWords.Words = append(taggedWords.Words, xmlWord)
	}

	xmlEncoder := xml.NewEncoder(&buffer)
	xmlEncoder.Indent("", "\t")
	xmlEncoder.Encode(&taggedWords)
	*/
	return buffer.Bytes()
}

func (this *Tag) ImportXML(xmlBytes []byte) {
	var taggedWords XMLTaggedWords
	xml.Unmarshal(xmlBytes, &taggedWords)

	// is this necessary?
	//println(taggedWords.XMLName.Space)

	if 0<len(taggedWords.Words) {
		for _, xmlWord := range taggedWords.Words {
			desc := xmlWord.Desc
			word := xmlWord.Name
			tags := strings.Split(xmlWord.Tags, " ")
			tagsBytes := convert.StringArrayToByteArray(tags)
			this.AddWordAndTagsSetDesc([]byte(word), tagsBytes, []byte(desc))
		}
	}
}

func (this *Tag) WordsByTags(tags [][]byte) (words [][]byte) {
	var wordsIndices, wordsIndicesIntersection []int
	tagsLength := len(tags)

	// first tag
	if 0<tagsLength {
		tagIndex, tagMatch := search.ByteWordBinary(this.TagsOffsets, this.TagsLengths, this.Tags, tags[0])

		if tagMatch {
			mapp := this.MapByIndex(tagIndex)
			wordsIndices = append(wordsIndices, mapp...)
		}

		if 1<tagsLength {
			tags = tags[1:]
		}
	}

	// rest of tags
	for _, tag := range tags {
		tagIndex, tagMatch := search.ByteWordBinary(this.TagsOffsets, this.TagsLengths, this.Tags, tag)
		wordsIndicesIntersection = wordsIndicesIntersection[:0]

		// search for mapping
		if tagMatch {
			wordsLength := len(wordsIndices)
			mapp := this.MapByIndex(tagIndex)
			mapLength := len(mapp)

			// intersect two sets
			for i,j:=0,0; (i<wordsLength)&&(j<mapLength); {
				wordIndex := wordsIndices[i]
				ref := mapp[j]
				if wordIndex<ref {
					i++
				} else if wordIndex>ref {
					j++
				} else {
					wordsIndicesIntersection = append(wordsIndicesIntersection, wordIndex)
					i++
					j++
				}
			}
		}

		// save matches
		if 0!=len(wordsIndicesIntersection) {
			copy(wordsIndices, wordsIndicesIntersection)
			wordsIndices = wordsIndices[:len(wordsIndicesIntersection)]
		// no matches: stop search
		} else {
			wordsIndices = wordsIndices[:0]
			break
		}
	}
	words = make([][]byte, len(wordsIndices))

	for i := range words {
		words[i] = this.WordByIndex(wordsIndices[i])
	}

	return words
}

func (this *Tag) Merge(oTag *Tag) {
	for oWordIndex, oWordOffset := range oTag.WordsOffsets {
		oWordLength := oTag.WordsLengths[oWordIndex]
		oDesc := oTag.DescByIndex(oWordIndex)
		oWord := oTag.Words[oWordOffset:oWordOffset+oWordLength]
		oTags := oTag.TagsByWordIndex(oWordIndex)
		this.AddWordAndTagsSetDesc(oWord, oTags, oDesc)
	}
}












