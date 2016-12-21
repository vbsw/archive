
//          Copyright 2016 Vitali Baumtrok
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package main

import (
	"os"
	"fmt"
)

type ParseResult int

const (
	NO_INFO ParseResult = iota
	TOO_MANY_ARGS
	UNKNOWN_ARG
	PRINT_HELP
	PRINT_VERSION
)

type tCLInfo struct {
	parseResult ParseResult
	unknownArg string
}

func newCLInfo ( ) *tCLInfo {
	clinfo := new(tCLInfo)
	clinfo.parse(os.Args)
	return clinfo
}

func (this *tCLInfo) None ( ) bool {
	return this.parseResult == NO_INFO
}

func (this *tCLInfo) Print ( ) {
	switch this.parseResult {
		case TOO_MANY_ARGS:
			this.printTooManyArgs()
			this.printHelpHint()

		case UNKNOWN_ARG:
			this.printUnknownArg()
			this.printHelpHint()

		case PRINT_HELP:
			this.printHelp()

		case PRINT_VERSION:
			this.printVersion()

		default:
			this.printHelp()
	}
}

func (this *tCLInfo) parse ( args []string ) {
	if len(args) > 2 {
		this.parseResult = TOO_MANY_ARGS

	} else if len(args) == 2 {

		if args[1] == "--help" {
			this.parseResult = PRINT_HELP

		} else if args[1] == "--version" {
			this.parseResult = PRINT_VERSION

		} else {
			this.parseResult = UNKNOWN_ARG
			this.unknownArg = args[1]
		}

	} else {
		this.parseResult = NO_INFO
	}
}

func (this *tCLInfo) printTooManyArgs ( ) {
	fmt.Println("error: too many arguments")
}

func (this *tCLInfo) printUnknownArg ( ) {
	fmt.Println("error: unknown argument '" + this.unknownArg + "'")
}

func (this *tCLInfo) printHelpHint ( ) {
	fmt.Println("Use --help for description on valid aguments.")
}

func (this *tCLInfo) printVersion ( ) {
	fmt.Println("Silent Shooter (version " + app.version.String() + ")")
	fmt.Println("This software is distributed under the terms of the")
	fmt.Println("Boost Software License, version 1.0. (See accompanying file LICENSE ")
	fmt.Println("or copy at <http://www.boost.org/LICENSE_1_0.txt>)")
}

func (this *tCLInfo) printHelp ( ) {
	fmt.Println("Usage:")
	fmt.Println("   silents             Starts game.")
	fmt.Println("   silents --help      Prints this help.")
	fmt.Println("   silents --version   Prints version and copying conditions.")
}

