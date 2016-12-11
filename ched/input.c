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

#include "buffer.h"


static size_t
Private_GetFileSize
(
	FILE* file
)
{
	fseek( file, 0, SEEK_END );
	long int file_size = ftell( file );
	rewind( file );
	if ( file_size >= 0 )
		return (size_t) file_size;
	else
		return 0;
}

char*
Input_ReadFile
(
	size_t* size,
	const char* file_name
)
{
	FILE* file = fopen( file_name, "rb" );

	if ( file != 0 )
	{
		const size_t file_size = Private_GetFileSize( file );

		if ( Buffer_file_size < file_size )
			Buffer_AllocateMinimumFileSize( file_size );

		*size = fread( Buffer_file_data, sizeof( char ), file_size, file );

		fclose( file );

		return Buffer_file_data;
	}
	else
	{
		return 0;
	}
}

