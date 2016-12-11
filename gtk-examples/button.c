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

static void
ButtonPressCallback ( GtkWidget* widget, GdkEvent* event, gpointer user_data )
{
	static int     counter      = 1;
	GdkEventButton* button_event = (GdkEventButton*) event;
	GdkEventType    event_type   = button_event->type;
	gboolean        left_button  = (button_event->button == 1);
	gboolean        right_button = (button_event->button == 3);

	if ( event_type == GDK_BUTTON_PRESS )
	{
		if ( GTK_IS_BUTTON( widget ) )
		{
			if ( left_button  || right_button )
			{
				if ( right_button )
					counter = 1;

				if ( counter < 10 ) 
				{
					static char text[20];
					sprintf( text, "Button (pressed: %d)", counter++ );
					gtk_button_set_label( GTK_BUTTON( widget ), text );
				}
				else
				{
					gtk_button_set_label( GTK_BUTTON( widget ), "BOOOOOOM!" );
				}
			}
		}
	}
}

GtkWidget*
CreateWindow()
{
	GtkWidget* window = gtk_window_new( GTK_WINDOW_TOPLEVEL );
	char*      title  = "Example GTK+ 3 Button";
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

	GtkWidget* window = CreateWindow();
	GtkWidget* button = gtk_button_new();
	char*      text   = "Button (pressed: 0)";

	gtk_button_set_label( GTK_BUTTON( button ), text );
	g_signal_connect( G_OBJECT( button ), "button-press-event", G_CALLBACK( ButtonPressCallback ), NULL );

	gtk_container_add( GTK_CONTAINER( window ), button );
	gtk_widget_show_all( GTK_WIDGET( window ) );
	gtk_main();

	return 0;
}

