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


#include <gtk/gtk.h>
#include <string.h>

#include "gui.h"
#include "settings.h"
#include "strings.h"
#include "clock.h"


static GtkWidget* Private_program_log;             /* GtkTextView */
static GtkWidget* Private_program_help_label;      /* GtkLabel */
static GtkWidget* Private_program_about_label;     /* GtkLabel */
static GtkWidget* Private_program_quit_yes_btn;    /* GtkButton */
static GtkWidget* Private_program_quit_cancel_btn; /* GtkButton */
static GtkWidget* Private_program_quit_label;      /* GtkLabel */
static GtkWidget* Private_program_quit;            /* GtkGrid */
static GtkWidget* Private_program_list;            /* GtkTreeView */
static GtkWidget* Private_program_list_scroll;     /* GtkScrolledWindow */
static GtkWidget* Private_program_log_scroll;      /* GtkScrolledWindow */
static GtkWidget* Private_program_pane;            /* GtkPaned */
static GtkWidget* Private_program_btn;             /* GtkToggleButton */
static GtkWidget* Private_settings_btn;            /* GtkToggleButton */
static GtkWidget* Private_editor_btn;              /* GtkToggleButton */
static GtkWidget* Private_window_grid;             /* GtkGrid */
static GtkWidget* Private_window;                  /* GtkWindow */

static int Private_nav_view = 0;
static int Private_nav_view_part_prev = -1;
static int Private_nav_view_part_curr = -1;


void
GUI_Log
(
	const char* text,
	const int   prepend_time
)
{
	GtkTextBuffer* text_buffer = gtk_text_view_get_buffer( GTK_TEXT_VIEW( Private_program_log ) );
	GtkTextIter pos;

	gtk_text_buffer_get_end_iter( text_buffer, &pos);

	if ( prepend_time )
	{
		const char* time_str = Clock_Get();

		gtk_text_buffer_insert( text_buffer, &pos, "[", 1 );
		gtk_text_buffer_get_end_iter( text_buffer, &pos);
		gtk_text_buffer_insert( text_buffer, &pos, time_str, Clock_time_size );
		gtk_text_buffer_get_end_iter( text_buffer, &pos);
		gtk_text_buffer_insert( text_buffer, &pos, "] ", 2 );
		gtk_text_buffer_get_end_iter( text_buffer, &pos);
	}
	gtk_text_buffer_insert( text_buffer, &pos, text, -1 );
}

static gboolean
Private_CloseWindow
(
	GtkWidget* widget,
	GdkEvent*  event,
	gpointer   data
)
{
	return FALSE;
}

static void
Private_DestroyWindow
(
	GtkWidget* widget,
	gpointer   data
)
{
	gtk_main_quit();
}

static void
Private_ProgramQuitBtnPressed
(
	GtkWidget* widget,
	GdkEvent*  event,
	gpointer   user_data
)
{
	gtk_main_quit();
}

static void
Private_SelectProgramLog ( void )
{
	GtkTreeView* view = GTK_TREE_VIEW( Private_program_list );
	GtkTreeModel* model = gtk_tree_view_get_model( view );
	GtkTreeIter iter;

	if ( gtk_tree_model_get_iter_first( model, &iter ) )
	{
		GtkTreePath* path = gtk_tree_model_get_path( model, &iter );
		GtkTreeSelection* selection = gtk_tree_view_get_selection( view );

		gtk_tree_selection_unselect_all( selection );
		//gtk_tree_selection_select_iter( selection, &iter );
		gtk_tree_view_set_cursor( view, path, NULL, FALSE );
		gtk_tree_path_free( path );
	}
}

static void
Private_SelectProgramHelp ( void )
{
	GtkTreeView* view = GTK_TREE_VIEW( Private_program_list );
	GtkTreeModel* model = gtk_tree_view_get_model( view );
	GtkTreeIter iter;

	if ( gtk_tree_model_get_iter_first( model, &iter ) )
		if ( gtk_tree_model_iter_next( model, &iter ) )
		{
			GtkTreePath* path = gtk_tree_model_get_path( model, &iter );
			GtkTreeSelection* selection = gtk_tree_view_get_selection( view );

			gtk_tree_selection_unselect_all( selection );
			//gtk_tree_selection_select_iter( selection, &iter );
			gtk_tree_view_set_cursor( view, path, NULL, FALSE );
			gtk_tree_path_free( path );
		}
}

static void
Private_SelectProgramAbout ( void )
{
	GtkTreeView* view = GTK_TREE_VIEW( Private_program_list );
	GtkTreeModel* model = gtk_tree_view_get_model( view );
	GtkTreeIter iter;

	if ( gtk_tree_model_get_iter_first( model, &iter ) )
		if ( gtk_tree_model_iter_next( model, &iter ) )
			if ( gtk_tree_model_iter_next( model, &iter ) )
			{
				GtkTreePath* path = gtk_tree_model_get_path( model, &iter );
				GtkTreeSelection* selection = gtk_tree_view_get_selection( view );

				gtk_tree_selection_unselect_all( selection );
				//gtk_tree_selection_select_iter( selection, &iter );
				gtk_tree_view_set_cursor( view, path, NULL, FALSE );
				gtk_tree_path_free( path );
			}
}

static void
Private_SelectProgramQuit ( void )
{
	GtkTreeView* view = GTK_TREE_VIEW( Private_program_list );
	GtkTreeModel* model = gtk_tree_view_get_model( view );
	GtkTreeIter iter;

	if ( gtk_tree_model_get_iter_first( model, &iter ) )
		if ( gtk_tree_model_iter_next( model, &iter ) )
			if ( gtk_tree_model_iter_next( model, &iter ) )
				if ( gtk_tree_model_iter_next( model, &iter ) )
				{
					GtkTreePath* path = gtk_tree_model_get_path( model, &iter );
					GtkTreeSelection* selection = gtk_tree_view_get_selection( view );

					gtk_tree_selection_unselect_all( selection );
					//gtk_tree_selection_select_iter( selection, &iter );
					gtk_tree_view_set_cursor( view, path, NULL, FALSE );
					gtk_tree_path_free( path );
				}
}

static void
Private_SelectProgram ( void )
{
	/* TODO */
	GUI_Log( "selected: Program\n", 1 );
}

static void
Private_SelectSettings ( void )
{
	GUI_Log( "selected: Settings (NOT IMPLEMENTED)\n", 1 );
}

static void
Private_SelectEditor ( void )
{
	GUI_Log( "selected: Editor (NOT IMPLEMENTED)\n", 1 );
}

static void
Private_ProgramQuitCancelBtnPressed
(
	GtkWidget* widget,
	GdkEvent*  event,
	gpointer   user_data
)
{
	if ( Private_nav_view_part_prev == 0 )
		Private_SelectProgramLog();
	else if ( Private_nav_view_part_prev == 1 )
		Private_SelectProgramHelp();
	else if ( Private_nav_view_part_prev == 2 )
		Private_SelectProgramAbout();
	else
		Private_SelectProgramLog();
}

static gboolean
Private_KeyPress
(
	GtkWidget* widget,
	GdkEvent*  event,
	gpointer   user_data
)
{
	GdkEventKey *const key_event = (GdkEventKey *const) event;
	const guint keyval = key_event->keyval;

	/* keys in gdk/gdkkeysyms.h */
	switch ( keyval )
	{
		case GDK_KEY_Return:
		if ( Private_nav_view == 0 && Private_nav_view_part_curr == 3 )
			gtk_main_quit();
		return TRUE;

		case GDK_KEY_Escape:
		Private_SelectProgramQuit();
		return TRUE;

		case GDK_KEY_F2:
		Private_SelectProgram();
		return TRUE;

		case GDK_KEY_F3:
		Private_SelectSettings();
		return TRUE;

		case GDK_KEY_F4:
		Private_SelectEditor();
		return TRUE;
	}

	return FALSE;
}

static void
Private_ProgramListChanged
(
	GtkTreeSelection* selection,
	gpointer          data
)
{
	GtkTreeIter iter;
	GtkTreeModel* model;

	if ( gtk_tree_selection_get_selected( selection, &model, &iter ) )
	{
		gchar* label;

		gtk_tree_model_get( model, &iter, 0, &label, -1 );

		/* log */
		if ( !strcmp( label, Strings_menu_label_program_log ) )
		{
			if ( Private_nav_view_part_curr != 0 )
			{
				Private_nav_view_part_prev = Private_nav_view_part_curr;
				Private_nav_view_part_curr = 0;

				if ( Private_nav_view_part_prev == 1 )
				{
					g_object_ref( (gpointer) Private_program_help_label );
					gtk_container_remove( GTK_CONTAINER( Private_program_pane ), Private_program_help_label );
				}
				else if ( Private_nav_view_part_prev == 2 )
				{
					g_object_ref( (gpointer) Private_program_about_label );
					gtk_container_remove( GTK_CONTAINER( Private_program_pane ), Private_program_about_label );
				}
				else if ( Private_nav_view_part_prev == 3 )
				{
					g_object_ref( (gpointer) Private_program_quit );
					gtk_container_remove( GTK_CONTAINER( Private_program_pane ), Private_program_quit );
				}
				gtk_paned_add2( GTK_PANED( Private_program_pane ), Private_program_log_scroll );
				gtk_widget_show_all( GTK_WIDGET( Private_window ) );
			}
		}
		/* help */
		else if ( !strcmp( label, Strings_menu_label_program_help ) )
		{
			if ( Private_nav_view_part_curr != 1 )
			{
				Private_nav_view_part_prev = Private_nav_view_part_curr;
				Private_nav_view_part_curr = 1;

				if ( Private_nav_view_part_prev == 0 )
				{
					g_object_ref( (gpointer) Private_program_log_scroll );
					gtk_container_remove( GTK_CONTAINER( Private_program_pane ), Private_program_log_scroll );
				}
				else if ( Private_nav_view_part_prev == 2 )
				{
					g_object_ref( (gpointer) Private_program_about_label );
					gtk_container_remove( GTK_CONTAINER( Private_program_pane ), Private_program_about_label );
				}
				else if ( Private_nav_view_part_prev == 3 )
				{
					g_object_ref( (gpointer) Private_program_quit );
					gtk_container_remove( GTK_CONTAINER( Private_program_pane ), Private_program_quit );
				}
				gtk_paned_add2( GTK_PANED( Private_program_pane ), Private_program_help_label );
				gtk_widget_show_all( GTK_WIDGET( Private_window ) );
			}
		}
		/* about */
		else if ( !strcmp( label, Strings_menu_label_program_about ) )
		{
			if ( Private_nav_view_part_curr != 2 )
			{
				Private_nav_view_part_prev = Private_nav_view_part_curr;
				Private_nav_view_part_curr = 2;

				if ( Private_nav_view_part_prev == 0 )
				{
					g_object_ref( (gpointer) Private_program_log_scroll );
					gtk_container_remove( GTK_CONTAINER( Private_program_pane ), Private_program_log_scroll );
				}
				else if ( Private_nav_view_part_prev == 1 )
				{
					g_object_ref( (gpointer) Private_program_help_label );
					gtk_container_remove( GTK_CONTAINER( Private_program_pane ), Private_program_help_label );
				}
				else if ( Private_nav_view_part_prev == 3 )
				{
					g_object_ref( (gpointer) Private_program_quit );
					gtk_container_remove( GTK_CONTAINER( Private_program_pane ), Private_program_quit );
				}
				gtk_paned_add2( GTK_PANED( Private_program_pane ), Private_program_about_label );
				gtk_widget_show_all( GTK_WIDGET( Private_window ) );
			}
		}
		/* quit */
		else if ( !strcmp( label, Strings_menu_label_program_quit ) )
		{
			if ( Private_nav_view_part_curr != 3 )
			{
				Private_nav_view_part_prev = Private_nav_view_part_curr;
				Private_nav_view_part_curr = 3;

				if ( Private_nav_view_part_prev == 0 )
				{
					g_object_ref( (gpointer) Private_program_log_scroll );
					gtk_container_remove( GTK_CONTAINER( Private_program_pane ), Private_program_log_scroll );
				}
				else if ( Private_nav_view_part_prev == 1 )
				{
					g_object_ref( (gpointer) Private_program_help_label );
					gtk_container_remove( GTK_CONTAINER( Private_program_pane ), Private_program_help_label );
				}
				else if ( Private_nav_view_part_prev == 2 )
				{
					g_object_ref( (gpointer) Private_program_about_label );
					gtk_container_remove( GTK_CONTAINER( Private_program_pane ), Private_program_about_label );
				}
				gtk_paned_add2( GTK_PANED( Private_program_pane ), Private_program_quit );
				gtk_widget_show_all( GTK_WIDGET( Private_window ) );
			}
			else
			{
				if ( Private_nav_view_part_prev == 0 )
					Private_SelectProgramLog();
				else if ( Private_nav_view_part_prev == 1 )
					Private_SelectProgramHelp();
				else if ( Private_nav_view_part_prev == 2 )
					Private_SelectProgramAbout();
				else
					Private_SelectProgramLog();
			}
		}

		g_free( label );
	}
}

static void
Private_CreateProgramLog ( void )
{
	Private_program_log        = gtk_text_view_new();
	GdkRGBA bgcolor            = { 0, 0, 0, 1 };
	GdkRGBA fgcolor            = { 0.1, 0.9, 0.1, 1 };
	PangoFontDescription* font = pango_font_description_from_string ("monospace 10");

	gtk_widget_set_hexpand( Private_program_log, TRUE );
	gtk_widget_set_vexpand( Private_program_log, TRUE );
	gtk_text_view_set_editable( GTK_TEXT_VIEW( Private_program_log ), FALSE );
	gtk_widget_override_background_color( GTK_WIDGET( Private_program_log ), GTK_STATE_FLAG_NORMAL, &bgcolor );
	gtk_widget_override_color( GTK_WIDGET( Private_program_log ), GTK_STATE_FLAG_NORMAL, &fgcolor );
	gtk_text_view_set_wrap_mode( GTK_TEXT_VIEW( Private_program_log ), GTK_WRAP_WORD );
	gtk_text_view_set_cursor_visible( GTK_TEXT_VIEW( Private_program_log ), FALSE );
	gtk_widget_override_font( GTK_WIDGET( Private_program_log ), font);
}

static void
Private_CreateProgramHelp ( void )
{
	Private_program_help_label = gtk_label_new( NULL );

	//gtk_label_set_justify( GTK_LABEL( Private_program_help_label ), GTK_JUSTIFY_CENTER );
	gtk_widget_set_hexpand( Private_program_help_label, TRUE );
	gtk_widget_set_vexpand( Private_program_help_label, TRUE );
}

static void
Private_CreateProgramAbout ( void )
{
	Private_program_about_label = gtk_label_new( NULL );

	//gtk_label_set_justify( GTK_LABEL( Private_program_about_label ), GTK_JUSTIFY_CENTER );
	gtk_widget_set_hexpand( Private_program_about_label, TRUE );
	gtk_widget_set_vexpand( Private_program_about_label, TRUE );
}

static void
Private_CreateProgramQuit ( void )
{
	Private_program_quit_label      = gtk_label_new( NULL );
	Private_program_quit_yes_btn    = gtk_button_new();
	Private_program_quit_cancel_btn = gtk_button_new();
	Private_program_quit            = gtk_grid_new();

	gtk_widget_set_halign( Private_program_quit_label, GTK_ALIGN_START );
	/* gtk_widget_set_hexpand( Private_program_quit_label, TRUE ); */
	g_signal_connect( G_OBJECT( Private_program_quit_yes_btn ), "button-press-event", G_CALLBACK( Private_ProgramQuitBtnPressed ), NULL );
	g_signal_connect( G_OBJECT( Private_program_quit_cancel_btn ), "button-press-event", G_CALLBACK( Private_ProgramQuitCancelBtnPressed ), NULL );
	gtk_widget_set_hexpand( Private_program_quit, TRUE );
	gtk_widget_set_vexpand( Private_program_quit, TRUE );
	gtk_grid_attach( GTK_GRID( Private_program_quit ), Private_program_quit_label, 0, 0, 3, 1 );
	gtk_grid_attach( GTK_GRID( Private_program_quit ), Private_program_quit_yes_btn, 0, 1, 1, 1 );
	gtk_grid_attach( GTK_GRID( Private_program_quit ), Private_program_quit_cancel_btn, 1, 1, 1, 1 );
	gtk_grid_set_row_spacing( GTK_GRID( Private_program_quit ), 10 );
	gtk_grid_set_column_spacing( GTK_GRID( Private_program_quit ), 10 );
	gtk_container_set_border_width( GTK_CONTAINER( Private_program_quit ), 10 );
}

static void
Private_CreateProgramList ( void )
{
	Private_program_list      = gtk_tree_view_new();
	GtkTreeViewColumn* column = gtk_tree_view_column_new();
	GtkCellRenderer* renderer = gtk_cell_renderer_text_new();
	GtkTreeModel* model       = GTK_TREE_MODEL( gtk_list_store_new( 1, G_TYPE_STRING ) );

	gtk_tree_view_column_pack_start( column, renderer, FALSE );
	gtk_tree_view_column_add_attribute( column, renderer, "text", 0 );
	gtk_tree_view_set_headers_visible( GTK_TREE_VIEW( Private_program_list ), FALSE);
	gtk_tree_view_set_enable_search( GTK_TREE_VIEW( Private_program_list ), FALSE);
	gtk_tree_view_append_column( GTK_TREE_VIEW( Private_program_list ), column );
	gtk_tree_view_set_model( GTK_TREE_VIEW( Private_program_list ), model );
	gtk_tree_selection_set_mode( gtk_tree_view_get_selection( GTK_TREE_VIEW( Private_program_list ) ), GTK_SELECTION_SINGLE );
	g_signal_connect( G_OBJECT( gtk_tree_view_get_selection( GTK_TREE_VIEW( Private_program_list ) ) ), "changed", G_CALLBACK( Private_ProgramListChanged ), NULL );
}

static void
Private_CreateProgramPane ( void )
{
	Private_program_pane         = gtk_paned_new( GTK_ORIENTATION_HORIZONTAL );
	Private_program_list_scroll  = gtk_scrolled_window_new( NULL, NULL );
	Private_program_log_scroll   = gtk_scrolled_window_new( NULL, NULL );

	/* g_object_set( G_OBJECT( Private_program_list_scroll ), "margin-left", 10, NULL ); */
	/* gtk_container_set_border_width( GTK_CONTAINER( Private_program_list_scroll ), 10 ); */
	gtk_container_add( GTK_CONTAINER( Private_program_list_scroll ), Private_program_list );
	gtk_container_add( GTK_CONTAINER( Private_program_log_scroll ), Private_program_log );
	gtk_paned_set_position( GTK_PANED( Private_program_pane ), 200 );
	gtk_paned_add1( GTK_PANED( Private_program_pane ), Private_program_list_scroll );
	gtk_paned_add2( GTK_PANED( Private_program_pane ), Private_program_log_scroll );
}

static void
Private_CreateViewButtons ( void )
{
	Private_program_btn  = gtk_toggle_button_new();
	Private_settings_btn = gtk_toggle_button_new();
	Private_editor_btn   = gtk_toggle_button_new();

	gtk_toggle_button_set_active( GTK_TOGGLE_BUTTON( Private_program_btn ), TRUE );
	gtk_widget_set_sensitive( Private_program_btn, FALSE );
	gtk_widget_set_sensitive( Private_settings_btn, FALSE );
	gtk_widget_set_sensitive( Private_editor_btn, FALSE );
}

static void
Private_CreateWindow ( void )
{
	Private_window      = gtk_window_new( GTK_WINDOW_TOPLEVEL );
	Private_window_grid = gtk_grid_new();
	GtkWidget* hgrid    = gtk_grid_new();

	gtk_container_set_border_width( GTK_CONTAINER( hgrid ), 10 );
	gtk_grid_attach( GTK_GRID( hgrid ), Private_program_btn, 0, 0, 1, 1 );
	gtk_grid_attach( GTK_GRID( hgrid ), Private_settings_btn, 1, 0, 1, 1 );
	gtk_grid_attach( GTK_GRID( hgrid ), Private_editor_btn, 2, 0, 1, 1 );
	gtk_grid_attach( GTK_GRID( Private_window_grid ), hgrid, 0, 0, 1, 1 );
	gtk_grid_attach( GTK_GRID( Private_window_grid ), Private_program_pane, 0, 1, 1, 1 );
	gtk_window_set_default_size( GTK_WINDOW( Private_window ), (guint) Settings_GetWindowWidth(), (guint) Settings_GetWindowHeight() );
	g_signal_connect( Private_window, "delete-event", G_CALLBACK( Private_CloseWindow ), NULL );
	g_signal_connect( Private_window, "destroy", G_CALLBACK( Private_DestroyWindow ), NULL );
	g_signal_connect( Private_window, "key-press-event", G_CALLBACK( Private_KeyPress ), NULL );
	gtk_container_add( GTK_CONTAINER( Private_window ), Private_window_grid );
}

static void
Private_UpdateLables ( void )
{
	GtkListStore* store;
	GtkTreeIter iter;

	gtk_window_set_title( GTK_WINDOW( Private_window ), Strings_window_title );
	gtk_button_set_label( GTK_BUTTON( Private_program_btn ), Strings_menu_label_program );
	gtk_widget_set_tooltip_text( Private_editor_btn, Strings_menu_label_editor_tooltip );
	gtk_button_set_label( GTK_BUTTON( Private_settings_btn ), Strings_menu_label_settings );
	gtk_widget_set_tooltip_text( Private_settings_btn, Strings_menu_label_settings_tooltip );
	gtk_button_set_label( GTK_BUTTON( Private_editor_btn ), Strings_menu_label_editor );
	gtk_widget_set_tooltip_text( Private_program_btn, Strings_menu_label_program_tooltip );
	gtk_button_set_label( GTK_BUTTON( Private_program_quit_yes_btn ), Strings_btn_label_program_quit_yes );
	gtk_button_set_label( GTK_BUTTON( Private_program_quit_cancel_btn ), Strings_btn_label_program_quit_cancel );
	gtk_label_set_text( GTK_LABEL( Private_program_quit_label ), Strings_btn_label_program_quit_headline );
	gtk_label_set_text( GTK_LABEL( Private_program_help_label ), Strings_program_help );
	gtk_label_set_text( GTK_LABEL( Private_program_about_label ), Strings_program_about );

	/* update program list labels */
	store = GTK_LIST_STORE( gtk_tree_view_get_model( GTK_TREE_VIEW( Private_program_list ) ) );

	gtk_list_store_clear( store );
	gtk_list_store_append( store, &iter );
	gtk_list_store_set( store, &iter, 0, Strings_menu_label_program_log, -1 );
	gtk_list_store_append( store, &iter );
	gtk_list_store_set( store, &iter, 0, Strings_menu_label_program_help, -1 );
	gtk_list_store_append( store, &iter );
	gtk_list_store_set( store, &iter, 0, Strings_menu_label_program_about, -1 );
	gtk_list_store_append( store, &iter );
	gtk_list_store_set( store, &iter, 0, Strings_menu_label_program_quit, -1 );
}

void
GUI_Start
(
	int    argc,
	char** argv
)
{
	gtk_init( &argc, &argv );

	Private_CreateProgramLog();
	Private_CreateProgramHelp();
	Private_CreateProgramAbout();
	Private_CreateProgramQuit();
	Private_CreateProgramList();
	Private_CreateProgramPane();
	Private_CreateViewButtons();
	Private_CreateWindow();
	Private_UpdateLables();
	Private_SelectProgramLog();

	GUI_Log( Strings_info_program_started, 1 );
	GUI_Log( "\n", 0 );

	gtk_widget_show_all( GTK_WIDGET( Private_window ) );
	gtk_main();
}

