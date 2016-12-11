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


#ifndef STRINGS_H
#define STRINGS_H


extern void
Strings_SetDefault ( void );
extern void
Strings_Load ( void );

extern const char* Strings_version;
extern const char* Strings_copyright;
extern const char* Strings_warranty_warning;
extern const char* Strings_too_many_arguments;
extern const char* Strings_help_hint;
extern const char* Strings_help;
extern const char* Strings_error_can_not_create_file;
extern const char* Strings_error_unknown_cmd;
extern const char* Strings_window_title;
extern const char* Strings_no_settings_file_set;
extern const char* Strings_settings_file_not_readable;
extern const char* Strings_creating_default_settings;
extern const char* Strings_setting_default_settings;
extern const char* Strings_loading_default_settings;
extern const char* Strings_menu_label_help;
extern const char* Strings_menu_label_status;
extern const char* Strings_menu_label_program;
extern const char* Strings_menu_label_settings;
extern const char* Strings_menu_label_editor;
extern const char* Strings_menu_label_program_tooltip;
extern const char* Strings_menu_label_settings_tooltip;
extern const char* Strings_menu_label_editor_tooltip;
extern const char* Strings_menu_label_program_log;
extern const char* Strings_menu_label_program_help;
extern const char* Strings_menu_label_program_about;
extern const char* Strings_menu_label_program_quit;
extern const char* Strings_info_program_started;
extern const char* Strings_btn_label_program_quit_headline;
extern const char* Strings_btn_label_program_quit_yes;
extern const char* Strings_btn_label_program_quit_cancel;
extern const char* Strings_program_help;
extern const char* Strings_program_about;

#endif /* STRINGS_H */
