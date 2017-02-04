
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
	c_INFO_NONE int = iota
	c_INFO_TO_MANY_ARGUMENTS
	c_INFO_HELP
	c_INFO_VERSION
	c_INFO_SCREENS
	c_INFO_UNKNOWN_ARGUMENT
)

type tClInfo struct {
	infoCode int
	unknownArgument string
}

func newCLInfo ( ) *tClInfo {
	clinfo := new(tClInfo)

	if len(os.Args) > 2 {
		clinfo.infoCode = c_INFO_TO_MANY_ARGUMENTS

	} else if len(os.Args) == 2 {

		if isHelp(os.Args[1]) {
			clinfo.infoCode = c_INFO_HELP

		} else if isVersion(os.Args[1]) {
			clinfo.infoCode = c_INFO_VERSION

		} else if isScreens(os.Args[1]) {
			clinfo.infoCode = c_INFO_SCREENS

		} else {
			clinfo.infoCode = c_INFO_UNKNOWN_ARGUMENT
			clinfo.unknownArgument = os.Args[1]
		}

	} else {
		clinfo.infoCode = c_INFO_NONE
	}
	return clinfo
}

func (this *tClInfo) None ( ) bool {
	return this.infoCode == c_INFO_NONE || this.infoCode == c_INFO_SCREENS
}

func (this *tClInfo) Screens ( ) bool {
	return this.infoCode == c_INFO_SCREENS
}

func (this *tClInfo) Print ( ) {
	switch this.infoCode {

		case c_INFO_TO_MANY_ARGUMENTS:
		printError("too many arguments")

		case c_INFO_HELP:
		fmt.Println("Usage:")
		fmt.Println("   xdspl           Prints the display server name.")
		fmt.Println("   xdspl help      Prints this help.")
		fmt.Println("   xdspl version   Prints version and copying conditions.")
		fmt.Println("   xdspl screens   Prints the count of screens and their resolution.")

		case c_INFO_VERSION:
		fmt.Println("XDspl (version " + version.String() + ")")
		fmt.Println("This software is distributed under the terms of the Boost Software License, version 1.0.")
		fmt.Println("(See accompanying file LICENSE or copy at <http://www.boost.org/LICENSE_1_0.txt>)")

		case c_INFO_UNKNOWN_ARGUMENT:
		printError("unknown argument '" + this.unknownArgument + "'")
	}
}

func isHelp ( arg string ) bool {
	return arg == "--help" || arg == "help" || arg == "-help"
}

func isVersion ( arg string ) bool {
	return arg == "--version" || arg == "version" || arg == "-version"
}

func isScreens ( arg string ) bool {
	return arg == "--screens" || arg == "screens" || arg == "-screens"
}

func printError ( errMsg string ) {
	fmt.Println("error:", errMsg)
	fmt.Println("Use --help for description on valid aguments.")
}

