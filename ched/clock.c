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


#include <time.h>

#include "buffer.h"


const size_t Clock_time_size = 19;

char*
Clock_Get ( void )
{
	/* yyyy-mm-dd HH:MM:SS */
	char* data = Buffer_GetMemory( Clock_time_size, 1 );
	const time_t t = time( NULL );
	struct tm local_time = *localtime( &t );
	const int year = local_time.tm_year + 1900;
	const int year_thousands = year / 1000;
	const int year_thousands_m = year_thousands * 1000;
	const int year_hundreds = ( year - year_thousands_m ) / 100;
	const int year_hundreds_m = year_hundreds * 100;
	const int year_tens = ( year - ( year_thousands_m + year_hundreds_m ) ) / 10;
	const int year_tens_m = year_tens * 10;
	const int year_ones = year - ( year_thousands_m + year_hundreds_m + year_tens_m );
	const int month = local_time.tm_mon + 1;
	const int month_tens = month / 10;
	const int month_ones = month - month_tens * 10;
	const int day = local_time.tm_mday;
	const int day_tens = day / 10;
	const int day_ones = day - day_tens * 10;
	const int hour = local_time.tm_hour;
	const int hour_tens = hour / 10;
	const int hour_ones = hour - hour_tens * 10;
	const int minutes = local_time.tm_min;
	const int minutes_tens = minutes / 10;
	const int minutes_ones = minutes - minutes_tens * 10;
	const int seconds = local_time.tm_sec;
	const int seconds_tens = seconds / 10;
	const int seconds_ones = seconds - seconds_tens * 10;
	data[0] = year_thousands + 48;
	data[1] = year_hundreds + 48;
	data[2] = year_tens + 48;
	data[3] = year_ones + 48;
	data[4] = '-';
	data[5] = month_tens + 48;
	data[6] = month_ones + 48;
	data[7] = '-';
	data[8] = day_tens + 48;
	data[9] = day_ones + 48;
	data[10] = ' ';
	data[11] = hour_tens + 48;
	data[12] = hour_ones + 48;
	data[13] = ':';
	data[14] = minutes_tens + 48;
	data[15] = minutes_ones + 48;
	data[16] = ':';
	data[17] = seconds_tens + 48;
	data[18] = seconds_ones + 48;

	return data;
}

