/**
 * gtk-examples (a set of programs demonstrating GTK+ 3 functionality)
 * Created by Vitali Baumtrok <vbsoftware@mailbox.org>
 *
 * This file is part of gtk-examples.
 *
 * You can redistribute gtk-examples under the terms of the
 * Creative Commons Zero v1.0 Universal license
 * (see accompanying file LICENSE or copy at
 * <https://creativecommons.org/publicdomain/zero/1.0/legalcode>).
 */


#include <gtk/gtk.h>
#include <stdio.h>


static gboolean
CloseWindowCallback ( GtkWidget* widget, GdkEvent* event, gpointer data )
{
	gboolean event_is_consumed = FALSE;

	printf( "window close event\n" );

	return event_is_consumed;
}

static void
DestroyWindowCallback ( GtkWidget* widget, gpointer data )
{
	printf( "window destroy event\n" );

	gtk_main_quit();
}

static gboolean
KeyPressCallback ( GtkWidget* widget, GdkEvent* event, gpointer user_data )
{
	GdkEventKey* key_event         = (GdkEventKey*) event;
	guint        key_value         = key_event->keyval;
	gboolean     event_is_consumed = FALSE;

	printf( "key press event (key: %d)\n", key_value );

	/* keys in gdk/gdkkeysyms.h */
	if ( key_value == GDK_KEY_Escape )
	{
		if ( GTK_IS_WINDOW( widget ) )
		{
			gtk_window_close( GTK_WINDOW( widget ) );
			event_is_consumed = TRUE;
		}
	}

	return event_is_consumed;
}

GtkWidget*
CreateWindow()
{
	GtkWidget* window = gtk_window_new( GTK_WINDOW_TOPLEVEL );
	char*      title  = "Example GTK+ 3 Layout";
	guint      width  = 480;
	guint      height = 360;

	gtk_window_set_title( GTK_WINDOW( window ), title );
	gtk_window_set_default_size( GTK_WINDOW( window ), width, height );

	g_signal_connect( G_OBJECT( window ), "delete-event", G_CALLBACK( CloseWindowCallback ), NULL );
	g_signal_connect( G_OBJECT( window ), "destroy", G_CALLBACK( DestroyWindowCallback ), NULL );
	g_signal_connect( G_OBJECT( window ), "key-press-event", G_CALLBACK( KeyPressCallback ), NULL );

	return window;
}

int
main ( int argc, char** argv )
{
	gtk_init( &argc, &argv );

	GtkWidget* window           = CreateWindow();
	GtkWidget* button_A         = gtk_button_new_with_label( "centered   button   A" );
	GtkWidget* button_B         = gtk_button_new_with_label( "centered   button   B" );
	GtkWidget* button_expanding = gtk_button_new_with_label( "expanding   button" );
	GtkWidget* grid_layout      = gtk_grid_new();

	gtk_widget_set_hexpand( button_A, TRUE );
	gtk_widget_set_vexpand( button_A, TRUE );
	gtk_widget_set_hexpand( button_B, TRUE );
	gtk_widget_set_vexpand( button_B, TRUE );

	gtk_widget_set_halign( button_A, GTK_ALIGN_CENTER );
	gtk_widget_set_valign( button_A, GTK_ALIGN_END );
	gtk_widget_set_halign( button_B, GTK_ALIGN_CENTER );
	gtk_widget_set_valign( button_B, GTK_ALIGN_START );

	gtk_grid_set_row_spacing( GTK_GRID( grid_layout ), 20 );
	gtk_container_set_border_width( GTK_CONTAINER( grid_layout ), 10 );
	gtk_grid_attach( GTK_GRID( grid_layout ), button_A,         1, 0, 2, 1 );
	gtk_grid_attach( GTK_GRID( grid_layout ), button_expanding, 0, 1, 4, 1 );
	gtk_grid_attach( GTK_GRID( grid_layout ), button_B,         1, 2, 2, 1 );

	gtk_container_add( GTK_CONTAINER( window ), grid_layout );
	gtk_widget_show_all( GTK_WIDGET( window ) );
	gtk_main();

	return 0;
}

