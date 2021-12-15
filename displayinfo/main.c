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


extern void args_initialize(int argc, char** argv);
extern int args_is_valid();
extern void help_print();
extern void help_print_hint();
extern void error_print();


int main(int argc, char** argv) {
	args_initialize(argc,argv);

	if (args_is_valid()) {
		help_print();

	} else {
		error_print();
	}

	return 0;
}

