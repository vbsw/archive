/**
 * DisplayInfo (info on display)
 * Copyright (C) 2017 Vitali Baumtrok
 *
 * This file is part of DisplayInfo.
 *
 * DisplayInfo is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * DisplayInfo is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with DisplayInfo.  If not, see <http://www.gnu.org/licenses/>.
 */


#include <string.h>


/* ************************* */
/*    exported functions     */
/* ************************* */

int parse_help(char* string) {
	return   !strcmp(string,"help")
	      || !strcmp(string,"--help")
	      || !strcmp(string,"-h")
	      || !strcmp(string,"-help");
}

int parse_version(char* string) {
	return   !strcmp(string,"version")
	      || !strcmp(string,"--version")
	      || !strcmp(string,"-v")
	      || !strcmp(string,"-version");
}

int parse_copyright(char* string) {
	return   !strcmp(string,"copyright")
	      || !strcmp(string,"--copyright")
	      || !strcmp(string,"-c")
	      || !strcmp(string,"-copyright");
}

int parse_resolution(char* string) {
	return   !strcmp(string,"resolution")
	      || !strcmp(string,"--resolution")
	      || !strcmp(string,"-r")
	      || !strcmp(string,"-resolution")
	      || !strcmp(string,"-res")
	      || !strcmp(string,"res")
	      || !strcmp(string,"--res");
}

int parse_server(char* string) {
	return   !strcmp(string,"server")
	      || !strcmp(string,"--server")
	      || !strcmp(string,"-s")
	      || !strcmp(string,"-server");
}

