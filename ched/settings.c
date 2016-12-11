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


#include <stdlib.h>

#include "settings.h"
#include "args.h"
#include "output.h"
#include "input.h"


static int Private_window_width  = 480;
static int Private_window_height = 360;
static const char* Private_default_file_data = "\nwindow\n\twidth 480\n\theight 360\n\n";
static const size_t Private_default_file_data_size = 32;

const char* Settings_file_path = "ched.settings.txt";


static void
Private_ReadSettings
(
	const char* data,
	const size_t size
)
{
	/* TODO interpret settings data instead of setting default */
	Settings_SetDefault();
}

void
Settings_LoadDefault ( void )
{
	size_t size = 0;
	const char* file_data = Input_ReadFile( &size, Settings_file_path );

	if ( file_data != 0 )
	{
		Private_ReadSettings( file_data, size );
	}
	else
	{
		Output_CreatingDefaultSettings( Settings_file_path );
		Output_CreateFile( Settings_file_path, Private_default_file_data, Private_default_file_data_size );
		file_data = Input_ReadFile( &size, Settings_file_path );

		if ( file_data != 0 )
		{
			Private_ReadSettings( file_data, size );
		}
		else
		{
			Output_ErrorCanNotCreateFile( Settings_file_path );
			Output_SettingDefaultSettings();
			Settings_SetDefault();
		}
	}
}

void
Settings_SetDefault ( void )
{
	Private_window_width  = 980;
	Private_window_height = 480;
}

void
Settings_Load ( void )
{
	if ( Args_IsSettings() )
	{
		if ( Args_GetSettingsFile() != 0 )
		{
			size_t size = 0;
			const char *const file_data = Input_ReadFile( &size, Args_GetSettingsFile() );

			if ( file_data != 0 )
			{
				Private_ReadSettings( file_data, size );
			}
			else
			{
				Output_WarningSettingsFileNotReadable( Args_GetSettingsFile() );
				Output_LoadingDefaultSettings();
				Settings_SetDefault();
			}
		}
		else
		{
			Output_NoSettingsFileSet();
			Output_LoadingDefaultSettings();
			Settings_SetDefault();
		}
	}
	else
	{
		/* no warning, because it is intended to laod from default path */
		Settings_LoadDefault();
	}
}

int
Settings_GetWindowWidth ( void )
{
	return Private_window_width;
}

int
Settings_GetWindowHeight ( void )
{
	return Private_window_height;
}

