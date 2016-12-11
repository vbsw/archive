/**
 * ChEd (a text editor)
 * Copyright (C) 2016 Vitali Baumtrok
 *
 * This file is part of ChEd.
 *
 * ChEd is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * ChEd is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */


#include "parse.h"


extern int
Parse_Help
(
	const char *const cstr
)
{
	return    cstr    != 0
	       && cstr[0] != 0
	       && cstr[0] == '-'
	       && cstr[1] != 0
	       && cstr[1] == '-'
	       && cstr[2] != 0
	       && cstr[2] == 'h'
	       && cstr[3] != 0
	       && cstr[3] == 'e'
	       && cstr[4] != 0
	       && cstr[4] == 'l'
	       && cstr[5] != 0
	       && cstr[5] == 'p'
	       && cstr[6] == 0;
}

extern int
Parse_Version
(
	const char *const cstr
)
{
	return    cstr    != 0
	       && cstr[0] != 0
	       && cstr[0] == '-'
	       && cstr[1] != 0
	       && cstr[1] == '-'
	       && cstr[2] != 0
	       && cstr[2] == 'v'
	       && cstr[3] != 0
	       && cstr[3] == 'e'
	       && cstr[4] != 0
	       && cstr[4] == 'r'
	       && cstr[5] != 0
	       && cstr[5] == 's'
	       && cstr[6] != 0
	       && cstr[6] == 'i'
	       && cstr[7] != 0
	       && cstr[7] == 'o'
	       && cstr[8] != 0
	       && cstr[8] == 'n'
	       && cstr[9] == 0;
}

extern int
Parse_Settings
(
	const char *const cstr
)
{
	return    cstr    != 0
	       && cstr[0] != 0
	       && cstr[0] == '-'
	       && cstr[1] != 0
	       && cstr[1] == '-'
	       && cstr[2] != 0
	       && cstr[2] == 's'
	       && cstr[3] != 0
	       && cstr[3] == 'e'
	       && cstr[4] != 0
	       && cstr[4] == 't'
	       && cstr[5] != 0
	       && cstr[5] == 't'
	       && cstr[6] != 0
	       && cstr[6] == 'i'
	       && cstr[7] != 0
	       && cstr[7] == 'n'
	       && cstr[8] != 0
	       && cstr[8] == 'g'
	       && cstr[9] != 0
	       && cstr[9] == 's'
	       && cstr[10] == 0;
}

