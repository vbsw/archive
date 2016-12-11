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

	/* keys in gdk/gdkkeysyms.h */
	if ( key_value == GDK_KEY_Escape )
	{
		if ( GTK_IS_WINDOW( widget ) )
		{
			gtk_window_close( GTK_WINDOW( widget ) );
			event_is_consumed = TRUE;
		}
	}
	else if ( key_value == GDK_KEY_1 )
	{
		if ( GTK_IS_NOTEBOOK( widget ) )
		{
			gtk_notebook_set_current_page( GTK_NOTEBOOK( widget ), 0);
			event_is_consumed = TRUE;
		}
	}
	else if ( key_value == GDK_KEY_2 )
	{
		if ( GTK_IS_NOTEBOOK( widget ) )
		{
			gtk_notebook_set_current_page( GTK_NOTEBOOK( widget ), 1);
			event_is_consumed = TRUE;
		}
	}
	else if ( key_value == GDK_KEY_3 )
	{
		if ( GTK_IS_NOTEBOOK( widget ) )
		{
			gtk_notebook_set_current_page( GTK_NOTEBOOK( widget ), 2);
			event_is_consumed = TRUE;
		}
	}

	return event_is_consumed;
}

GtkWidget*
CreateWindow()
{
	GtkWidget* window = gtk_window_new( GTK_WINDOW_TOPLEVEL );
	char*      title  = "Example GTK+ 3 Notebook";
	guint      width  = 480;
	guint      height = 360;

	gtk_window_set_title( GTK_WINDOW( window ), title );
	gtk_window_set_default_size( GTK_WINDOW( window ), width, height );

	g_signal_connect( G_OBJECT( window ), "delete-event", G_CALLBACK( CloseWindowCallback ), NULL );
	g_signal_connect( G_OBJECT( window ), "destroy", G_CALLBACK( DestroyWindowCallback ), NULL );
	g_signal_connect( G_OBJECT( window ), "key-press-event", G_CALLBACK( KeyPressCallback ), NULL );

	return window;
}

GtkWidget*
CreateNotebook()
{
	GtkWidget* notebook = gtk_notebook_new();
	GtkWidget* label1   = gtk_label_new( "monkey (key: 1)" );
	GtkWidget* child1   = gtk_label_new( "content monkey" );
	GtkWidget* label2   = gtk_label_new( "reindeer (key: 2)" );
	GtkWidget* child2   = gtk_label_new( "content reindeer" );
	GtkWidget* label3   = gtk_label_new( "racoon (key: 3)" );
	GtkWidget* child3   = gtk_label_new( "content racoon" );

	gtk_notebook_set_show_border( GTK_NOTEBOOK( notebook ), FALSE );
	gtk_notebook_set_tab_pos( GTK_NOTEBOOK( notebook ), GTK_POS_LEFT );
	gtk_notebook_append_page( GTK_NOTEBOOK( notebook ), child1, label1 );
	gtk_notebook_append_page( GTK_NOTEBOOK( notebook ), child2, label2 );
	gtk_notebook_append_page( GTK_NOTEBOOK( notebook ), child3, label3 );

	g_signal_connect( G_OBJECT( notebook ), "key-press-event", G_CALLBACK( KeyPressCallback ), NULL );

	return notebook;
}

int
main ( int argc, char** argv )
{
	gtk_init( &argc, &argv );

	GtkWidget* window   = CreateWindow();
	GtkWidget* notebook = CreateNotebook();

	gtk_container_add( GTK_CONTAINER( window ), notebook );
	gtk_widget_show_all( GTK_WIDGET( window ) );
	gtk_main();

	return 0;
}

