/**
 * Key Words - a program to manage data by key words
 * Copyright 2015 Vitali Baumtrok
 * 
 * This file is part of Key Words.
 *
 * Key Words is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Key Words is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */


#include <stdlib.h>


#include <gtk/gtk.h>


static void quitProgram ( void *const data )
{
	GtkWidget *const window = ( (GtkWidget**) data )[0];

	gtk_widget_destroy ( window );
}


static void buttonFunc_quitProgram ( GtkWidget *widget, gpointer data )
{
	quitProgram ( data );
}


static gboolean windowFunc_closeGTKMainWindow ( GtkWidget *widget, GdkEvent *event, gpointer data )
{
	quitProgram ( data );

	return TRUE;
}


/**
 * This is called when quitProgram is called.
 */
static void windowFunc_destroyGTKMainWindow ( GtkWidget *widget, gpointer data )
{
	gtk_main_quit ( );
}


static void createGTKMainWindow ( void *const data )
{
	GtkWidget *label = gtk_label_new ( "Quit This Amazing Program" );
	GtkWidget *button = gtk_button_new ( );
	GtkWidget *window = gtk_window_new ( GTK_WINDOW_TOPLEVEL );
	GtkWidget *vBox = gtk_vbox_new ( FALSE, 0 );
	GtkWidget *hBox = gtk_hbox_new ( FALSE, 0 );
	( (void**) data )[0] = (void*) window;

	gtk_misc_set_padding (GTK_MISC (label), 20, 3);
	gtk_window_set_title ( GTK_WINDOW (window), "Key Words" );
	gtk_window_set_default_size ( GTK_WINDOW (window), 480, 360 );
	gtk_box_pack_start ( GTK_BOX (vBox), hBox, TRUE, FALSE, 0 );
	gtk_container_add ( GTK_CONTAINER (button), label );
	gtk_container_add ( GTK_CONTAINER (window), vBox );
	gtk_box_pack_start ( GTK_BOX (hBox), button, TRUE, FALSE, 0 );
	gtk_widget_show ( label );
	gtk_widget_show ( hBox );
	gtk_widget_show ( vBox );
	gtk_widget_show ( button );
	gtk_widget_show ( window );
	g_signal_connect ( button, "clicked", G_CALLBACK (buttonFunc_quitProgram), data );
	g_signal_connect ( window, "delete-event", G_CALLBACK (windowFunc_closeGTKMainWindow), data );
	g_signal_connect ( window, "destroy", G_CALLBACK (windowFunc_destroyGTKMainWindow), data );
	g_object_set ( window, "border-width", 10, NULL );
}


static void createGUI ( int argc, char **argv, void *data )
{
	gtk_init ( &argc, &argv );
	createGTKMainWindow ( data );
}


static void startProgram ( void *const data )
{
	gtk_main ( );
}


static void setupProgram ( int argc, char **argv, void *const data )
{
	// TODO: read parameters
}


int main ( int argc, char **argv )
{
	void *data = malloc ( sizeof (void*) * 1 );

	setupProgram ( argc, argv, data );
	createGUI ( argc, argv, data );
	startProgram ( data );

	return 0;
}

