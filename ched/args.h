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


#ifndef ARGS_H
#define ARGS_H


extern void
Args_Read ( int argc, char **argv );
extern int
Args_IsAny ( void );
extern int
Args_IsTooMany ( void );
extern int
Args_IsHelp ( void );
extern int
Args_IsVersion ( void );
extern int
Args_IsSettings ( void );
extern char*
Args_GetSettingsFile ( void );
extern int
Args_IsUnknown ( void );
extern char*
Args_GetUnknown ( void );


#endif /* ARGS_H */
