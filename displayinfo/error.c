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


#include <stdio.h>


static const char * c1 = 0;
static const char * c2 = 0;
static const char * c3 = 0;


/* ************************* */
/*    exported functions     */
/* ************************* */

void error_set(const char *const str1) {
	c1 = str1;
	c2 = 0;
	c3 = 0;
}

void error_set_cc(const char *const str1, const char *const str2) {
	c1 = str1;
	c2 = str2;
	c3 = 0;
}

void error_set_ccc(const char *const str1, const char *const str2, const char *const str3) {
	c1 = str1;
	c2 = str2;
	c3 = str3;
}

void error_print() {
	if (c1!=0) {

		if (c2!=0) {

			if (c3!=0) {
				printf("%s%s%s",c1,c2,c3);

			} else {
			printf("%s%s",c1,c2);
			}

		} else {
			printf("%s",c1);
		}
	}
}

