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


extern int parse_help(const char *const string);
extern int parse_version(const char *const string);
extern int parse_copyright(const char *const string);
extern int parse_server(const char *const string);
extern int parse_resolution(const char *const string);
extern int parse_server(const char *const string);
extern void error_set(const char *const str1);
extern void error_set_ccc(const char *const str1, const char *const str2, const char *const str3);


#define STATE_NOT_INITIALIZED 0
#define STATE_NO_ARGUMENTS 1
#define STATE_HELP 2
#define STATE_VERSION 3
#define STATE_COPYRIGHT 4
#define STATE_RESOLUTION 5
#define STATE_SERVER 6
#define STATE_UNKNOWN_ARGUMENT 7
#define STATE_TOO_MANY_ARGUMENTS 8

static int args_state = STATE_NOT_INITIALIZED;
static const char *args_unknown_argument = 0;

static void args_initiliaze_none() {
	args_state = STATE_NO_ARGUMENTS;
}

static void args_initialize_one(const int argc, const char *const *const argv) {
	const char *const argument = argv[1];

	if (parse_help(argument)) {
		args_state = STATE_HELP;

	} else if (parse_version(argument)) {
		args_state = STATE_VERSION;

	} else if (parse_copyright(argument)) {
		args_state = STATE_COPYRIGHT;

	} else if (parse_resolution(argument)) {
		args_state = STATE_RESOLUTION;

	} else if (parse_server(argument)) {
		args_state = STATE_SERVER;

	} else {
		args_state = STATE_UNKNOWN_ARGUMENT;
		error_set_ccc("error: unknown argument \"",argument,"\"\n");
	}
}

static void args_initialize_many(const int argc, const char *const *const argv) {
	args_state = STATE_TOO_MANY_ARGUMENTS;
	error_set("error: unknwon arguments\n");
}


/* ************************* */
/*    exported functions     */
/* ************************* */

void args_initialize(const int argc, const char *const *const argv) {
	if (argc==1) {
		args_initiliaze_none();

	} else if (argc==2) {
		args_initialize_one(argc,argv);

	} else {
		args_initialize_many(argc,argv);
	}
}

int args_is_valid() {
	return    args_state != STATE_NOT_INITIALIZED
	       && args_state != STATE_UNKNOWN_ARGUMENT
	       && args_state != STATE_TOO_MANY_ARGUMENTS;
}

int args_is_none() {
	return    args_state == STATE_NOT_INITIALIZED
	       || args_state == STATE_NO_ARGUMENTS;
}

int args_is_help() {
	return args_state == STATE_HELP;
}

int args_is_version() {
	return args_state == STATE_VERSION;
}

int args_is_copyright() {
	return args_state == STATE_COPYRIGHT;
}

int args_is_server() {
	return args_state == STATE_SERVER;
}

int args_is_resolution() {
	return args_state == STATE_RESOLUTION;
}

