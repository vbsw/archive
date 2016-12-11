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

#include "args.h"
#include "output.h"
#include "parse.h"


static int   Private_too_many     = 0;
static int   Private_any          = 0;
static int   Private_help         = 0;
static int   Private_version      = 0;
static int   Private_settings     = 0;
static char* Private_settings_str = 0;
static int   Private_unknown      = 0;
static char* Private_unknown_str  = 0;

static void
Private_Reset ( void )
{
	Private_too_many     = 0;
	Private_any          = 0;
	Private_help         = 0;
	Private_version      = 0;
	Private_settings     = 0;
	Private_settings_str = 0;
	Private_unknown      = 0;
	Private_unknown_str  = 0;
}

void
Args_Read
(
	int    argc,
	char** argv
)
{
	Private_Reset();

	if ( argc > 1 )
	{
		Private_any = 1;

		if ( argc == 2 )
		{

			if ( Parse_Help( argv[1] ) )
			{
				Private_help = 1;
			}
			else if ( Parse_Version( argv[1] ) )
			{
				Private_version = 1;
			}
			else if ( Parse_Settings( argv[1] ) )
			{
				Private_settings = 1;
			}
			else
			{
				Private_unknown     = 1;
				Private_unknown_str = argv[1];
			}
		}
		else if ( argc == 3 )
		{
			if ( Parse_Settings( argv[1] ) )
			{
				Private_settings     = 1;
				Private_settings_str = argv[2];
			}
			else
			{
				Private_unknown     = 1;
				Private_unknown_str = argv[1];
			}
		}
		else
		{
			Private_too_many = 1;
		}
	}
}

int
Args_IsAny ( void ) { return Private_any; }
int
Args_IsTooMany ( void ) { return Private_too_many; }
int
Args_IsHelp ( void ) { return Private_help; }
int
Args_IsVersion ( void ) { return Private_version; }
int
Args_IsSettings ( void ) { return Private_settings; }
char*
Args_GetSettingsFile ( void ) { return Private_settings_str; }
int
Args_IsUnknown ( void ) { return Private_unknown; }
char*
Args_GetUnknown ( void ) { return Private_unknown_str; }

