
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package main

import (
	"os"
	"fmt"
)

const (
	INFO_NONE int = iota
	INFO_TO_MANY_ARGUMENTS
	INFO_HELP
	INFO_VERSION
	INFO_UNKNOWN_ARGUMENT
)

type tInfoRequest struct {
	infoCode int
	unknownArgument string
}

func printError ( errMsg string ) {
	fmt.Println("error:", errMsg)
	fmt.Println("Use --help for description on valid aguments.")
}

func newInfoRequest ( ) *tInfoRequest {
	infoRequest := new(tInfoRequest)

	if len(os.Args) > 2 {
		infoRequest.infoCode = INFO_TO_MANY_ARGUMENTS

	} else if len(os.Args) == 2 {

		if os.Args[1] == "--help" {
			infoRequest.infoCode = INFO_HELP

		} else if os.Args[1] == "--version" {
			infoRequest.infoCode = INFO_VERSION

		} else {
			infoRequest.infoCode = INFO_UNKNOWN_ARGUMENT
			infoRequest.unknownArgument = os.Args[1]
		}

	} else {
		infoRequest.infoCode = INFO_NONE
	}
	return infoRequest
}

func (this *tInfoRequest) None ( ) bool {
	return this.infoCode == INFO_NONE
}

func (this *tInfoRequest) Print ( ) {
	switch this.infoCode {

		case INFO_TO_MANY_ARGUMENTS:
		printError("too many arguments")

		case INFO_HELP:
		fmt.Println("Usage:")
		fmt.Println("   silents             Starts game.")
		fmt.Println("   silents --help      Prints this help.")
		fmt.Println("   silents --version   Prints version and copying conditions.")

		case INFO_VERSION:
		fmt.Print("Silent Shooter (version ", version, ")")
		fmt.Println()
		fmt.Println("This software is distributed under the terms of the Boost Software License, version 1.0.")
		fmt.Println("(See accompanying file LICENSE or copy at <http://www.boost.org/LICENSE_1_0.txt>)")

		case INFO_UNKNOWN_ARGUMENT:
		printError("unknown argument '" + this.unknownArgument + "'")
	}
}

