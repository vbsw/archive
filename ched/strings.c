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


#include "strings.h"


const char* Strings_version;
const char* Strings_copyright;
const char* Strings_warranty_warning;
const char* Strings_too_many_arguments;
const char* Strings_help_hint;
const char* Strings_help;
const char* Strings_error_can_not_create_file;
const char* Strings_error_unknown_cmd;
const char* Strings_window_title;
const char* Strings_no_settings_file_set;
const char* Strings_settings_file_not_readable;
const char* Strings_creating_default_settings;
const char* Strings_setting_default_settings;
const char* Strings_loading_default_settings;
const char* Strings_menu_label_help;
const char* Strings_menu_label_status;
const char* Strings_menu_label_program;
const char* Strings_menu_label_settings;
const char* Strings_menu_label_editor;
const char* Strings_menu_label_program_tooltip;
const char* Strings_menu_label_settings_tooltip;
const char* Strings_menu_label_editor_tooltip;
const char* Strings_menu_label_program_log;
const char* Strings_menu_label_program_help;
const char* Strings_menu_label_program_about;
const char* Strings_menu_label_program_quit;
const char* Strings_info_program_started;
const char* Strings_btn_label_program_quit_headline;
const char* Strings_btn_label_program_quit_yes;
const char* Strings_btn_label_program_quit_cancel;
const char* Strings_program_about;
const char* Strings_program_help;

void
Strings_SetDefault ( void )
{
	Strings_version                         = "ChEd 0.1.1\n";
	Strings_copyright                       = "Copyright (C) 2016 Vitali Baumtrok\n";
	Strings_warranty_warning                = "This is free software; see the source for copying conditions. There is NO\nwarranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.\n";
	Strings_too_many_arguments              = "Too many arguments.\n";
	Strings_help_hint                       = "Use --help for more information on how to use ChEd.\n";
	Strings_help                            = "Usage:       ched [--settings FILE]\n             ched COMMAND\nCOMMANDs:\n --help      Prints this help\n --version   Prints version and copyright\n";
	Strings_error_can_not_create_file       = "Error: Can not create file: ";
	Strings_error_unknown_cmd               = "Error: Unknown command: ";
	Strings_window_title                    = "ChEd 0.1.1";
	Strings_no_settings_file_set            = "Warning: No settings file set.\n";
	Strings_settings_file_not_readable      = "Warning: Could not read settings file: ";
	Strings_creating_default_settings       = "Warning: Creating default settings file: ";
	Strings_setting_default_settings        = "Warning: Settings set to default.\n";
	Strings_loading_default_settings        = "Warning: Loading default settings.\n";
	Strings_menu_label_help                 = "Help";
	Strings_menu_label_status               = "Status";
	Strings_menu_label_program              = "Program";
	Strings_menu_label_settings             = "Settings";
	Strings_menu_label_editor               = "Editor";
	Strings_menu_label_program_tooltip      = "Program View (F2)";
	Strings_menu_label_settings_tooltip     = "Settings View (F3)";
	Strings_menu_label_editor_tooltip       = "Editor View (F4)";
	Strings_menu_label_program_log          = "Log";
	Strings_menu_label_program_help         = "Help";
	Strings_menu_label_program_about        = "About";
	Strings_menu_label_program_quit         = "Quit";
	Strings_info_program_started            = "Program started!";
	Strings_btn_label_program_quit_headline = "Quit program (Enter/Escape)?";
	Strings_btn_label_program_quit_yes      = "yes";
	Strings_btn_label_program_quit_cancel   = "cancel";
	Strings_program_help                    = "Shortcuts\n\nesc    quit program\nf2       program view\nf3       settings view\nf4       editor view";
	Strings_program_about                   = "ChEd 0.1.1 (a text editor)\nCopyright (C) 2016 Vitali Baumtrok\n\nChEd is free software: you can redistribute it and/or modify\nit under the terms of the GNU General Public License as published by\nthe Free Software Foundation, either version 3 of the License, or\n(at your option) any later version.\n\nChEd is distributed in the hope that it will be useful,\nbut WITHOUT ANY WARRANTY; without even the implied warranty of\nMERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the\nGNU General Public License for more details.\n\nYou should have received a copy of the GNU General Public License\nalong with this program.  If not, see <http://www.gnu.org/licenses/>.";
}

void
Strings_Load ( void )
{
	/* TODO */
	Strings_SetDefault();
}

