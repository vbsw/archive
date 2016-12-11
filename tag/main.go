
//          Copyright Vitali Baumtrok 2014, 2015.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)

package main


import (
	"os"
	"fmt"
	"github.com/vitalibaumtrok/gobase/convert"
	"github.com/vitalibaumtrok/gobase/read"
	"github.com/vitalibaumtrok/gobase/search"
	"github.com/vitalibaumtrok/gobase/tag"
	"github.com/vitalibaumtrok/gobase/write"
)


const (
	MAJOR_VER = 2
	MINOR_VER = 1
	PATCH_VER = 1
	DEFAULT_WORD_SEPERATOR = " "
	DEFAULT_QUOTE = '\''
)


func main() {
	name := os.Args[0]
	argc := len(os.Args)

	if 1==argc {
		//printVersion()
		//printCopyright()
		//fmt.Println()
		printUsage(name)

	} else {
		cmd := os.Args[1]

		if isCommandHelp(cmd) {
			if 2==argc {
				printUsage(name)
			} else if 3==argc {
				printHelpCmd(name, os.Args[2])
			} else {
				printTooMuchArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			}

		} else if isCommandVersion(cmd) {
			if 2==argc {
				printVersion()
				printCopyright()
			} else {
				printTooMuchArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			}

		} else if isCommandW(cmd) {
			if 3<argc {
				params := convert.StringArrayToMergedStringArray(os.Args[3:], DEFAULT_QUOTE)
				if 0<len(params) {
					processW(os.Args[2], params[0], params[1:])
				} else {
					printNotEnoughArgumentsFor(cmd)
					printCmdHelpHint(name, cmd)
				}
			} else {
				printNotEnoughArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			}

		} else if isCommandS(cmd) {
			if 2<argc {
				params := convert.StringArrayToMergedStringArray(os.Args[3:], DEFAULT_QUOTE)
				if 0<len(params) {
					processS(os.Args[2], params)
				} else {
					printNotEnoughArgumentsFor(cmd)
					printCmdHelpHint(name, cmd)
				}
			} else {
				printNotEnoughArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			}

		} else if isCommandM(cmd) {
			if 2<argc {
				params := convert.StringArrayToMergedStringArray(os.Args[3:], DEFAULT_QUOTE)
				if 0<len(params) {
					processM(os.Args[2], params)
				} else {
					printNotEnoughArgumentsFor(cmd)
					printCmdHelpHint(name, cmd)
				}
			} else {
				printNotEnoughArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			}

		} else if isCommandD(cmd) {
			if 3<argc {
				params := convert.StringArrayToMergedStringArray(os.Args[3:], DEFAULT_QUOTE)
				if 2==len(params) {
					processD(os.Args[2], params[0], params[1])
				} else if 1==len(params) {
					processD(os.Args[2], params[0], "")
				} else if 0==len(params) {
					printNotEnoughArgumentsFor(cmd)
					printCmdHelpHint(name, cmd)
				} else {
					printTooMuchArgumentsFor(cmd)
					printCmdHelpHint(name, cmd)
				}
			} else {
				printNotEnoughArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			}

		} else if isCommandRm(cmd) {
			if 3<argc {
				params := convert.StringArrayToMergedStringArray(os.Args[3:], DEFAULT_QUOTE)
				if 0<len(params) {
					processRm(os.Args[2], params[0], params[1:])
				} else {
					printNotEnoughArgumentsFor(cmd)
					printCmdHelpHint(name, cmd)
				}
			} else {
				printNotEnoughArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			}

		} else if isCommandRmw(cmd) {
			if 2<argc {
				params := convert.StringArrayToMergedStringArray(os.Args[3:], DEFAULT_QUOTE)
				processRmw(os.Args[2], params)
			} else {
				printNotEnoughArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			}

		} else if isCommandExp(cmd) {
			if 4==argc {
				processExp(os.Args[2], os.Args[3])
			} else if 4<argc {
				printTooMuchArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			} else {
				printNotEnoughArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			}

		} else if isCommandImp(cmd) {
			if 4==argc {
				processImp(os.Args[2], os.Args[3])
			} else if 4<argc {
				printTooMuchArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			} else {
				printNotEnoughArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			}

		} else if isCommandPw(cmd) {
			if 2<argc {
				params := convert.StringArrayToMergedStringArray(os.Args[3:], DEFAULT_QUOTE)
				if 1==len(params) {
					processPw(os.Args[2], params[0])
				} else if 0==len(params) {
					processPw(os.Args[2], DEFAULT_WORD_SEPERATOR)
				} else {
					printTooMuchArgumentsFor(cmd)
					printCmdHelpHint(name, cmd)
				}
			} else {
				printNotEnoughArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			}

		} else if isCommandPt(cmd) {
			if 2<argc {
				params := convert.StringArrayToMergedStringArray(os.Args[3:], DEFAULT_QUOTE)
				if 1==len(params) {
					processPt(os.Args[2], params[0])
				} else if 0==len(params) {
					processPt(os.Args[2], DEFAULT_WORD_SEPERATOR)
				} else {
					printTooMuchArgumentsFor(cmd)
					printCmdHelpHint(name, cmd)
				}
			} else {
				printNotEnoughArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			}

		} else if isCommandPwd(cmd) {
			if 3<argc {
				params := convert.StringArrayToMergedStringArray(os.Args[3:], DEFAULT_QUOTE)
				if 1==len(params) {
					processPwd(os.Args[2], params[0])
				} else if 0==len(params) {
					printNotEnoughArgumentsFor(cmd)
					printCmdHelpHint(name, cmd)
				} else {
					printTooMuchArgumentsFor(cmd)
					printCmdHelpHint(name, cmd)
				}
			} else {
				printNotEnoughArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			}

		} else if isCommandPwt(cmd) {
			if 3<argc {
				params := convert.StringArrayToMergedStringArray(os.Args[3:], DEFAULT_QUOTE)
				if 2==len(params) {
					processPwt(os.Args[2], params[0], params[1])
				} else if 1==len(params) {
					processPwt(os.Args[2], params[0], DEFAULT_WORD_SEPERATOR)
				} else if 0==len(params) {
					printNotEnoughArgumentsFor(cmd)
					printCmdHelpHint(name, cmd)
				} else {
					printTooMuchArgumentsFor(cmd)
					printCmdHelpHint(name, cmd)
				}
			} else {
				printNotEnoughArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			}

		} else if isCommandPwn(cmd) {
			if 2<argc {
				params := convert.StringArrayToMergedStringArray(os.Args[3:], DEFAULT_QUOTE)
				if 0==len(params) {
					processPwn(os.Args[2])
				} else {
					printTooMuchArgumentsFor(cmd)
					printCmdHelpHint(name, cmd)
				}
			} else {
				printNotEnoughArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			}

		} else if isCommandPtn(cmd) {
			if 2<argc {
				params := convert.StringArrayToMergedStringArray(os.Args[3:], DEFAULT_QUOTE)
				if 0==len(params) {
					processPtn(os.Args[2])
				} else {
					printTooMuchArgumentsFor(cmd)
					printCmdHelpHint(name, cmd)
				}
			} else {
				printNotEnoughArgumentsFor(cmd)
				printCmdHelpHint(name, cmd)
			}
		} else {
			printUnknownCommand(cmd)
			printCmdHelpHint(name, "")
		}
	}
}

func printVersion() {
	fmt.Printf(
		"tag (version %d.%d.%d) - a program to manage words with tags\n",
		MAJOR_VER,
		MINOR_VER,
		PATCH_VER)
}

func printCopyright() {
	fmt.Println("Copyright (C) 2014, 2015 Vitali Baumtrok")
	fmt.Println("This is free software; see the source for copying conditions. There is NO")
	fmt.Println("warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.")
}

func printUsage(name string) {
	fmt.Printf("Usage:     %s help [COMMAND]\n", name)
	fmt.Printf("           %s version\n", name)
	fmt.Printf("           %s COMMAND [params]\n", name)
	fmt.Println("COMMANDs:")
	fmt.Println("  w        add a word with tag(s) to a file")
	fmt.Println("  s        search for words by tag(s)")
	fmt.Println("  m        merge files to one file and print number of files merged")
	fmt.Println("  d        set the description for a word")
	fmt.Println("  rm       remove tag(s) from a word")
	fmt.Println("  rmw      remove word(s)")
	fmt.Println("  exp      export data to an xml file")
	fmt.Println("  imp      import data from an xml file")
	fmt.Println("  pw       print all words")
	fmt.Println("  pt       print all tags")
	fmt.Println("  pwd      print word description")
	fmt.Println("  pwt      print word tags")
	fmt.Println("  pwn      print the words number")
	fmt.Println("  ptn      print the tags number")
}

func printUnknownCommand(cmd string) {
	fmt.Printf("Unknown command: %s\n", cmd)
}

func printTooMuchArgumentsFor(cmd string) {
	fmt.Printf("Too much arguments for command '%s'.\n", cmd)
}

func printNotEnoughArgumentsFor(cmd string) {
	fmt.Printf("Not enough valid arguments for command '%s'.\n", cmd)
}

func printCmdHelpHint(name, cmd string) {
	fmt.Printf("For help type: %s help %s\n", name, cmd)
}

func printLineUsage() {
	fmt.Println("Usage:")
}

func printLineExamples() {
	fmt.Println("Examples:")
}

func printLineExplanation() {
	fmt.Println("Explanation:")
}

func printLineReturns() {
	fmt.Println("Returns:")
}

func printError(err error) {
	println("Error:", err.Error())
}

func printHelpCmd(name, cmd string) {
	if isCommandHelp(cmd) {
		fmt.Printf("You won't get more help then by typing:  %s help\n", name)

	} else if isCommandVersion(cmd) {
		fmt.Println("Prints version and a short description of this program.")

	} else if isCommandW(cmd) {
		printLineUsage()
		fmt.Printf("  %s %s FILE WORD {TAG}\n", name, cmd)
		printLineExamples()
		fmt.Printf("  %s %s a.db hammer tools\n", name, cmd)
		fmt.Printf("  %s %s a.db hammer red owned\n", name, cmd)
		fmt.Printf("  %s %s a.db georg\n", name, cmd)
		printLineExplanation()
		fmt.Println("  word 'hammer' is added with the tag 'tools' to file a.db")
		fmt.Println("  tags 'red' and 'owned' are added to word 'hammer' in file a.db")
		fmt.Println("  word 'georg' is added without tags to file a.db")

	} else if isCommandS(cmd) {
		printLineUsage()
		fmt.Printf("  %s %s FILE {TAG}\n", name, cmd)
		printLineExamples()
		fmt.Printf("  %s %s a.db tools\n", name, cmd)
		fmt.Printf("  %s %s a.db red owned\n", name, cmd)
		printLineExplanation()
		fmt.Println("  search in file a.db for entries with the tag 'tools'")
		fmt.Println("  search in file a.db for entries with the tags 'red' and 'owned'")

	} else if isCommandM(cmd) {
		printLineUsage()
		fmt.Printf("  %s %s FILE {FILE}\n", name, cmd)
		printLineExamples()
		fmt.Printf("  %s %s a.db b.db\n", name, cmd)
		fmt.Printf("  %s %s a.db b.db fileX\n", name, cmd)
		fmt.Printf("  %s %s a.db\n", name, cmd)
		printLineExplanation()
		fmt.Println("  the file 'b.db' is merged into file 'a.db'")
		fmt.Println("  the files 'b.db' and fileX are merged into file 'a.db'")
		fmt.Println("  nothing is merged into file 'a.db'; 'a.db' stays unchanged")

	} else if isCommandD(cmd) {
		printLineUsage()
		fmt.Printf("  %s %s FILE WORD 'DESCRIPTION'\n", name, cmd)
		printLineExamples()
		fmt.Printf("  %s %s a.db pencil 'a useful tool'\n", name, cmd)
		fmt.Printf("  %s %s a.db georg person\n", name, cmd)
		printLineExplanation()
		fmt.Println("  adds to the word pencil the description 'a useful tool'")
		fmt.Println("  adds to the word georg the description 'person'")
		printLineReturns()
		fmt.Println("  1 if word exists, otherwise 0")

	} else if isCommandRm(cmd) {
		printLineUsage()
		fmt.Printf("  %s %s FILE WORD {TAG}\n", name, cmd)
		printLineExamples()
		fmt.Printf("  %s %s a.db hammer tools\n", name, cmd)
		fmt.Printf("  %s %s a.db pencil red owned\n", name, cmd)
		fmt.Printf("  %s %s a.db georg\n", name, cmd)
		printLineExplanation()
		fmt.Println("  removes from 'hammer' the tag 'tools' in file a.db")
		fmt.Println("  removes from 'pencil' the tags 'red' and 'owned' in file a.db")
		fmt.Println("  nothing is removed from 'georg' in file a.db")

	} else if isCommandRmw(cmd) {
		printLineUsage()
		fmt.Printf("  %s %s FILE {WORD}\n", name, cmd)
		printLineExamples()
		fmt.Printf("  %s %s a.db house\n", name, cmd)
		fmt.Printf("  %s %s a.db hammer georg\n", name, cmd)
		printLineExplanation()
		fmt.Println("  word 'house' is removed from file a.db")
		fmt.Println("  words 'hammer' and 'georg' are removed from file a.db")

	} else if isCommandExp(cmd) {
		printLineUsage()
		fmt.Printf("  %s %s FILE FILE\n", name, cmd)
		printLineExamples()
		fmt.Printf("  %s %s a.db a.txt\n", name, cmd)
		printLineExplanation()
		fmt.Println("  all data from file a.db is exported to file a.txt")

	} else if isCommandImp(cmd) {
		printLineUsage()
		fmt.Printf("  %s %s FILE FILE\n", name, cmd)
		printLineExamples()
		fmt.Printf("  %s %s a.db a.txt\n", name, cmd)
		printLineExplanation()
		fmt.Println("  all data from file a.txt is imported to file a.db")

	} else if isCommandPw(cmd) {
		printLineUsage()
		fmt.Printf("  %s %s FILE\n", name, cmd)
		printLineExamples()
		fmt.Printf("  %s %s a.db\n", name, cmd)
		printLineExplanation()
		fmt.Println("  prints all words stored in file a.db")

	} else if isCommandPt(cmd) {
		printLineUsage()
		fmt.Printf("  %s %s FILE\n", name, cmd)
		printLineExamples()
		fmt.Printf("  %s %s a.db\n", name, cmd)
		printLineExplanation()
		fmt.Println("  prints all tags stored in file a.db")

	} else if isCommandPwd(cmd) {
		printLineUsage()
		fmt.Printf("  %s %s FILE WORD\n", name, cmd)
		printLineExamples()
		fmt.Printf("  %s %s a.db georg\n", name, cmd)
		printLineExplanation()
		fmt.Println("  prints the description of the word 'georg'")

	} else if isCommandPwt(cmd) {
		printLineUsage()
		fmt.Printf("  %s %s FILE WORD\n", name, cmd)
		printLineExamples()
		fmt.Printf("  %s %s a.db hammer\n", name, cmd)
		printLineExplanation()
		fmt.Println("  prints all tags of the word 'hammer'")

	} else if isCommandPwn(cmd) {
		printLineUsage()
		fmt.Printf("  %s %s FILE\n", name, cmd)
		printLineExamples()
		fmt.Printf("  %s %s a.db\n", name, cmd)
		fmt.Printf("  %s %s file\n", name, cmd)
		printLineExplanation()
		fmt.Println("  prints the number of all tags saved in 'a.db'")
		fmt.Println("  prints the number of all tags saved in 'file'")

	} else if isCommandPtn(cmd) {
		printLineUsage()
		fmt.Printf("  %s %s FILE\n", name, cmd)
		printLineExamples()
		fmt.Printf("  %s %s a.db\n", name, cmd)
		fmt.Printf("  %s %s file\n", name, cmd)
		printLineExplanation()
		fmt.Println("  prints the number of all words saved in 'a.db'")
		fmt.Println("  prints the number of all words saved in 'file'")

	} else {
		printUnknownCommand(cmd)
	}
}

func isCommandVersion(s string) bool {
	return (len(s)==1 && s[0:1]=="v") || (len(s)==7 && s[0:7]=="version") ||
	(len(s)==2 && s[0:2]=="-v") || (len(s)==8 && s[0:8]=="-version") ||
	(len(s)==3 && s[0:3]=="--v") || (len(s)==9 && s[0:9]=="--version")
}

func isCommandHelp(s string) bool {
	return (len(s)==1 && s[0:1]=="h") || (len(s)==4 && s[0:4]=="help") ||
	(len(s)==2 && s[0:2]=="-h") || (len(s)==5 && s[0:5]=="-help") ||
	(len(s)==3 && s[0:3]=="--h") || (len(s)==6 && s[0:6]=="--help")
}

func isCommandW(s string) bool {
	return len(s)==1 && s[0:1]=="w"
}

func isCommandS(s string) bool {
	return len(s)==1 && s[0:1]=="s"
}

func isCommandM(s string) bool {
	return len(s)==1 && s[0:1]=="m"
}

func isCommandD(s string) bool {
	return len(s)==1 && s[0:1]=="d"
}

func isCommandRm(s string) bool {
	return len(s)==2 && s[0:2]=="rm"
}

func isCommandRmw(s string) bool {
	return len(s)==3 && s[0:3]=="rmw"
}

func isCommandExp(s string) bool {
	return len(s)==3 && s[0:3]=="exp"
}

func isCommandImp(s string) bool {
	return len(s)==3 && s[0:3]=="imp"
}

func isCommandPw(s string) bool {
	return len(s)==2 && s[0:2]=="pw"
}

func isCommandPt(s string) bool {
	return len(s)==2 && s[0:2]=="pt"
}

func isCommandPwd(s string) bool {
	return len(s)==3 && s[0:3]=="pwd"
}

func isCommandPwt(s string) bool {
	return len(s)==3 && s[0:3]=="pwt"
}

func isCommandPwn(s string) bool {
	return len(s)==3 && s[0:3]=="pwn"
}

func isCommandPtn(s string) bool {
	return len(s)==3 && s[0:3]=="ptn"
}

// insert word with tags in file on given path
func processW(path, word string, tags []string) {
	encodedData, err := read.BytesFromFileOnPath(path)

	if err==nil {
		var tagg tag.Tag
		wordBytes := []byte(word)
		tagsBytes := convert.StringArrayToByteArray(tags)
		encodedData = tag.Init(encodedData)
		tagg.Unmarshal(encodedData)

		if tagg.AddWordAndTags(wordBytes, tagsBytes) {
			encodedData = make([]byte, tagg.EncodedSize())
			tagg.Marshal(encodedData)
			err = write.BytesToFileOnPath(path, encodedData)

			if err!=nil {
				printError(err)
			}
		}

	} else {
		printError(err)
	}
}

func processS(path string, tags []string) {
	encodedData, err := read.BytesFromFileOnPath(path)

	if err==nil {
		var tagg tag.Tag
		tagsBytes := convert.StringArrayToByteArray(tags)
		encodedData = tag.Init(encodedData)
		tagg.Unmarshal(encodedData)
		words := tagg.WordsByTags(tagsBytes)

		for _, word := range words {
			println(string(word))
		}

		if 0==len(words) {
			println()
		}

	} else {
		printError(err)
	}
}

func processM(path string, oPaths []string) {
	encodedData, err := read.BytesFromFileOnPath(path)

	if err==nil {
		var tagg tag.Tag
		var filesMerged int
		encodedData = tag.Init(encodedData)
		tagg.Unmarshal(encodedData)

		for _, oPath := range oPaths {
			oEncodedData, oErr := read.BytesFromFileOnPath(oPath)

			if oErr!=nil {
				var oTagg tag.Tag
				filesMerged++
				oEncodedData = tag.Init(oEncodedData)
				oTagg.Unmarshal(oEncodedData)
				tagg.Merge(&oTagg)
			}
		}

		encodedData = make([]byte, tagg.EncodedSize())
		tagg.Marshal(encodedData)
		err = write.BytesToFileOnPath(path, encodedData)

		if err==nil {
			println(filesMerged)
		} else {
			printError(err)
		}

	} else {
		printError(err)
	}
}

func processD(path, word, desc string) {
	encodedData, err := read.BytesFromFileOnPath(path)

	if err==nil {
		var tagg tag.Tag
		encodedData = tag.Init(encodedData)
		tagg.Unmarshal(encodedData)
		wordIndex, wordMatch := search.ByteWordBinary(tagg.WordsOffsets, tagg.WordsLengths, tagg.Words, []byte(word))

		if wordMatch {
			tagg.SetDesc(wordIndex, []byte(desc))
			encodedData = make([]byte, tagg.EncodedSize())
			tagg.Marshal(encodedData)
			err = write.BytesToFileOnPath(path, encodedData)

			if err==nil {
				println(1)
			} else {
				printError(err)
			}

		} else {
			println(0)
		}

	} else {
		printError(err)
	}
}

// remove tags from a word
func processRm(path, word string, tags []string) {
	encodedData, err := read.BytesFromFileOnPath(path)

	if err==nil {
		var removed int

		if tag.MINIMUM_SIZE<=len(encodedData) {
			var tagg tag.Tag
			tagsBytes := convert.StringArrayToByteArray(tags)
			tagg.Unmarshal(encodedData)
			removed = tagg.RemoveTags([]byte(word), tagsBytes)

			if 0<removed {
				encodedData = make([]byte, tagg.EncodedSize())
				tagg.Marshal(encodedData)
				err = write.BytesToFileOnPath(path, encodedData)

				if err!=nil {
					printError(err)
				}
			}
		}
		println(removed)

	} else {
		printError(err)
	}
}

// remove words
func processRmw(path string, words []string) {
	encodedData, err := read.BytesFromFileOnPath(path)

	if err==nil {
		var removed int

		if tag.MINIMUM_SIZE<=len(encodedData) {
			var tagg tag.Tag
			wordsBytes := convert.StringArrayToByteArray(words)
			tagg.Unmarshal(encodedData)
			removed = tagg.RemoveWords(wordsBytes)

			if 0<removed {
				encodedData = make([]byte, tagg.EncodedSize())
				tagg.Marshal(encodedData)
				err = write.BytesToFileOnPath(path, encodedData)

				if err!=nil {
					printError(err)
				}
			}
		}
		println(removed)

	} else {
		printError(err)
	}
}

// export tags to an xml file
func processExp(path, xmlExportPath string) {
	encodedData, err := read.BytesFromFileOnPath(path)

	if err==nil {
		var tagg tag.Tag
		encodedData = tag.Init(encodedData)
		tagg.Unmarshal(encodedData)
		xmlBytes := tagg.ExportTXT()
		
		if err==nil {
			err = write.BytesToFileOnPath(xmlExportPath, xmlBytes)
			if err!=nil {
				printError(err)
			}
		} else {
			printError(err)
		}

	} else {
		printError(err)
	}
}

// import tags from an xml file
func processImp(path, xmlImportPath string) {
	encodedData, err := read.BytesFromFileOnPath(path)

	if err==nil {
		var txtBytes []byte
		txtBytes, err = read.BytesFromFileOnPath(xmlImportPath)
		
		if err==nil {
			var tagg tag.Tag
			encodedData = tag.Init(encodedData)
			tagg.Unmarshal(encodedData)
			tagg.ImportTXT(txtBytes)
			encodedData = make([]byte, tagg.EncodedSize())
			tagg.Marshal(encodedData)
			err = write.BytesToFileOnPath(path, encodedData)

			if err!=nil {
				printError(err)
			}

		} else {
			printError(err)
		}

	} else {
		printError(err)
	}
}

func processPw(path, seperator string) {
	encodedData, err := read.BytesFromFileOnPath(path)

	if err==nil {

		if tag.MINIMUM_SIZE<=len(encodedData) {
			var tagg tag.Tag
			tagg.Unmarshal(encodedData)
			words := tagg.WordsArray()

			if 0<len(words) {
				print(string(words[0]))
				words = words[1:]
				seperator = convert.StringToEscapedString(seperator)
				for _, word := range words {
					print(seperator, string(word))
				}
			}
			println()

		} else {
			println()
		}

	} else {
		printError(err)
	}
}

func processPt(path, seperator string) {
	encodedData, err := read.BytesFromFileOnPath(path)

	if err==nil {

		if tag.MINIMUM_SIZE<=len(encodedData) {
			var tagg tag.Tag
			tagg.Unmarshal(encodedData)
			tags := tagg.TagsArray()

			if 0<len(tags) {
				print(string(tags[0]))
				tags = tags[1:]
				seperator = convert.StringToEscapedString(seperator)
				for _, t := range tags {
					print(seperator, string(t))
				}
			}
			println()

		} else {
			println()
		}

	} else {
		printError(err)
	}
}

// tags of a word
func processPwd(path, word string) {
	encodedData, err := read.BytesFromFileOnPath(path)

	if err==nil {

		if tag.MINIMUM_SIZE<=len(encodedData) {
			var tagg tag.Tag
			tagg.Unmarshal(encodedData)
			wordIndex, wordMatch := search.ByteWordBinary(tagg.WordsOffsets, tagg.WordsLengths, tagg.Words, []byte(word))

			if wordMatch {
				desc := tagg.DescByIndex(wordIndex)
				println(string(desc))
			} else {
				println()
			}

		} else {
			println()
		}

	} else {
		printError(err)
	}
}

// tags of a word
func processPwt(path, word, seperator string) {
	encodedData, err := read.BytesFromFileOnPath(path)

	if err==nil {

		if tag.MINIMUM_SIZE<=len(encodedData) {
			var tagg tag.Tag
			tagg.Unmarshal(encodedData)
			wordIndex, wordMatch := search.ByteWordBinary(tagg.WordsOffsets, tagg.WordsLengths, tagg.Words, []byte(word))

			if wordMatch {
				tags := tagg.TagsByWordIndex(wordIndex)

				if 0<len(tags) {
					print(string(tags[0]))
					tags = tags[1:]
					seperator = convert.StringToEscapedString(seperator)
					for _, t := range tags {
						print(seperator, string(t))
					}
				}
				println()

			} else {
				println()
			}

		} else {
			println()
		}

	} else {
		printError(err)
	}
}

func processPwn(path string) {
	encodedData, err := read.BytesFromFileOnPath(path)

	if err==nil {

		if tag.MINIMUM_SIZE<=len(encodedData) {
			var tagg tag.Tag
			tagg.Unmarshal(encodedData)
			println(len(tagg.WordsOffsets))

		} else {
			println(0)
		}

	} else {
		printError(err)
	}
}

func processPtn(path string) {
	encodedData, err := read.BytesFromFileOnPath(path)

	if err==nil {

		if tag.MINIMUM_SIZE<=len(encodedData) {
			var tagg tag.Tag
			tagg.Unmarshal(encodedData)
			println(len(tagg.TagsOffsets))

		} else {
			println(0)
		}

	} else {
		printError(err)
	}
}





