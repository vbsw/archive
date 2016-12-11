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

#define ROUND(size) (size && 7? ( size & ( (((size_t)0)-(size_t)1) - 7 ) ) + 8: size)


size_t Buffer_file_size = 0;
char* Buffer_file_data = 0;
size_t Buffer_tresh_count = 0;


static size_t
Private_GetMinumumFileSize
(
	const size_t file_size
)
{
	return file_size * 2;
}

void
Buffer_AllocateMinimumFileSize
(
	size_t file_size
)
{
	if ( Buffer_file_data != 0 )
		free( Buffer_file_data );

	Buffer_file_size = Private_GetMinumumFileSize( file_size );
	Buffer_file_data = malloc( Buffer_file_size );

	if ( Buffer_file_data == 0 )
		Buffer_file_size = 0;
}

void*
Buffer_GetMemory
(
	const size_t size,
	const int increment_tresh
)
{
	const size_t rounded_size = ROUND( size );

	if ( increment_tresh )
		Buffer_tresh_count += rounded_size;

	/* TODO memory management */
	return malloc( rounded_size );
}

void
Buffer_AddTresh ( size_t size )
{
	Buffer_tresh_count += ROUND( size );
}

