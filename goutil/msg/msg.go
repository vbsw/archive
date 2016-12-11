/* 
 * util - a library for the programming language Go
 * Copyright 2014 Vitali Baumtrok
 * 
 * This file is part of util.
 * 
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * 
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * 
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 * 
 */

package msg

import (
	"fmt"
	"time"
)

const (
	TIME_FORMAT = "2006-01-02 15:04:05"
	TIME_SPEAKER_TEMPLATE = "%s  %s: %s\n"
	SPEAKER_TEMPLATE = "%s: %s\n"
	ERROR_STR = "Error"
	WARNING_STR = "Warning"
	INFO_STR = "Info"
	TRACE_STR = "Trace"
	DEBUG_STR = "Debug"
)

func TimeAndErrorLn(msg string) (string) {
	return fmt.Sprintf(TIME_SPEAKER_TEMPLATE, Time(), ERROR_STR, msg)
}

func TimeAndWarningLn(msg string) (string) {
	return fmt.Sprintf(TIME_SPEAKER_TEMPLATE, Time(), WARNING_STR, msg)
}

func TimeAndInfoLn(msg string) (string) {
	return fmt.Sprintf(TIME_SPEAKER_TEMPLATE, Time(), INFO_STR, msg)
}

func TimeAndTraceLn(msg string) (string) {
	return fmt.Sprintf(TIME_SPEAKER_TEMPLATE, Time(), TRACE_STR, msg)
}

func TimeAndDebugLn(msg string) (string) {
	return fmt.Sprintf(TIME_SPEAKER_TEMPLATE, Time(), DEBUG_STR, msg)
}

func ErrorLn(msg string) (string) {
	return fmt.Sprintf(SPEAKER_TEMPLATE, ERROR_STR, msg)
}

func WarningLn(msg string) (string) {
	return fmt.Sprintf(SPEAKER_TEMPLATE, WARNING_STR, msg)
}

func InfoLn(msg string) (string) {
	return fmt.Sprintf(SPEAKER_TEMPLATE, INFO_STR, msg)
}

func TraceLn(msg string) (string) {
	return fmt.Sprintf(SPEAKER_TEMPLATE, TRACE_STR, msg)
}

func DebugLn(msg string) (string) {
	return fmt.Sprintf(SPEAKER_TEMPLATE, DEBUG_STR, msg)
}

func SpeakerLn(speaker, msg string) (string) {
	return fmt.Sprintf(SPEAKER_TEMPLATE, speaker, msg)
}

func Time() (string) {
	currTime := time.Now()
	return currTime.Format(TIME_FORMAT)
}









