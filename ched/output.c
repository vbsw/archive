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


#include <stdio.h>

#include "strings.h"


void Output_Version ( void ) { printf( "%s", Strings_version ); }
void Output_Copyright ( void ) { printf( "%s", Strings_copyright ); }
void Output_WarrantyWarning ( void ) { printf( "%s", Strings_warranty_warning ); }
void Output_EmptyLine ( void ) { printf( "\n" ); }
void Output_TooManyArguments ( void ) { printf( "%s", Strings_too_many_arguments ); }
void Output_HelpHint ( void ) { printf( "%s", Strings_help_hint ); }
void Output_Help ( void ) { printf( "%s", Strings_help ); }
void Output_NoSettingsFileSet ( void ) { printf( "%s", Strings_no_settings_file_set ); }
void Output_CreatingDefaultSettings ( const char *const arg ) { printf( "%s%s\n", Strings_creating_default_settings, arg ); }
void Output_SettingDefaultSettings ( void ) { printf( "%s", Strings_setting_default_settings ); }
void Output_LoadingDefaultSettings ( void ) { printf( "%s", Strings_loading_default_settings ); }
void Output_ErrorCanNotCreateFile ( const char *const arg ) { printf( "%s%s\n", Strings_error_can_not_create_file, arg ); }
void Output_ErrorUnknownArgument ( const char *const arg ) { printf( "%s%s\n", Strings_error_unknown_cmd, arg ); }
void Output_WarningSettingsFileNotReadable ( const char *const arg ) { printf( "%s%s\n", Strings_settings_file_not_readable, arg ); }

void
Output_CreateFile
(
	const char *const file_path,
	const char *const data,
	const size_t size
)
{
	FILE* file = fopen( file_path, "wb" );
	fwrite( (const void*) data, sizeof(char), size, file );
	fclose( file );
}

