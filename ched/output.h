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


#ifndef OUTPUT_H
#define OUTPUT_H


extern void
Output_Version ( void );
extern void
Output_Copyright ( void );
extern void
Output_WarrantyWarning ( void );
extern void
Output_EmptyLine ( void );
extern void
Output_TooManyArguments ( void );
extern void
Output_HelpHint ( void );
extern void
Output_Help ( void );
extern void
Output_NoSettingsFileSet ( void );
extern void
Output_SettingDefaultSettings ( void );
extern void
Output_CreatingDefaultSettings ( const char *const arg );
extern void
Output_LoadingDefaultSettings ( void );
extern void
Output_ErrorCanNotCreateFile ( const char *const arg );
extern void
Output_ErrorUnknownArgument ( const char *const arg );
extern void
Output_WarningSettingsFileNotReadable ( const char *const arg );
extern void
Output_CreateFile ( const char *const file_path, const char *const data, const size_t size );


#endif /* OUTPUT_H */
