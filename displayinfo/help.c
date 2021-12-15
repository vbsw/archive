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

extern int args_is_help();
extern int args_is_none();
extern int args_is_copyright();
extern int args_is_version();
extern int args_is_resolution();
extern int args_is_server();
extern int args_is_resolution();
extern int args_is_server();
extern const char* xserver_default_display_server();
extern void xserver_resolution(int *const width, int *const height);


static void help_print_help() {
	printf("Usage:        displayinfo <command>\n");
	printf("COMMAND:\n");
	printf("  help        prints this help\n");
	printf("  version     prints version\n");
	printf("  copyright   prints copyright and copying conditions\n");
	printf("  resolution  prints resolution of current screen\n");
	printf("  server      prints default X server connection string\n");
}

static void help_print_version() {
	printf("1.0.1\n");
}

static void help_print_copyright() {
	printf("Copyright (C) 2017 Vitali Baumtrok (vbsw@mailbox.org)\n");
	printf("This program is free software: you can redistribute it and/or modify\n");
	printf("it under the terms of the GNU General Public License as published by\n");
	printf("the Free Software Foundation, either version 3 of the License, or\n");
	printf("(at your option) any later version.\n");
}

static void help_print_resolution() {
	int width, height;
	xserver_resolution(&width,&height);
	if (width>=0 && height>=0) {
		printf("%ix%i\n",width,height);
	} else {
		printf("n.a.\n");
	}
}

static void help_print_server() {
	const char *const server = xserver_default_display_server();
	if (server!=0) {
		printf("%s\n",server);
	} else {
		printf("n.a.");
	}
}


/* ************************* */
/*    exported functions     */
/* ************************* */

void help_print_hint() {
	printf("Use \"displayinfo help\" for valid arguments.\n");
}

void help_print() {
	if (args_is_help() || args_is_none()) {
		help_print_help();

	} else if (args_is_version()) {
		help_print_version();

	} else if (args_is_copyright()) {
		help_print_copyright();

	} else if (args_is_resolution()) {
		help_print_resolution();

	} else if (args_is_server()) {
		help_print_server();
	}
}

