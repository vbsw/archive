/**
 * URL Saver (a tool to manage URLs by keywords)
 * Copyright 2016 Vitali Baumtrok
 * 
 * This file is part of URL Saver.
 * 
 * URL Saver is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * 
 * URL Saver is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * 
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */


package com.gitlab.vbg.urlsaver;


/**
 * @author Vitali Baumtrok
 */
@SuppressWarnings ( "serial" )
final class Properties extends java.util.Properties
{
	private static final String WINDOW_TITLE = "window.title"; //$NON-NLS-1$
	private static final String WINDOW_WIDTH = "window.width"; //$NON-NLS-1$
	private static final String WINDOW_HEIGHT = "window.height"; //$NON-NLS-1$
	private static final String WINDOW_MAXIMIZED = "window.maximized"; //$NON-NLS-1$
	private static final String DATA_AUTOLOAD = "data.autoload"; //$NON-NLS-1$

	private static final String WINDOW_TITLE_DEFAULT = "URL Saver"; //$NON-NLS-1$
	private static final String WINDOW_WIDTH_DEFAULT = "720"; //$NON-NLS-1$
	private static final String WINDOW_HEIGHT_DEFAULT = "480"; //$NON-NLS-1$
	private static final String WINDOW_MAXIMIZED_DEFAULT = "false"; //$NON-NLS-1$
	private static final String DATA_AUTOLOAD_DEFAULT = "false"; //$NON-NLS-1$

	String getWindowTitle ( )
	{
		final String windowTitle = getProperty ( WINDOW_TITLE, WINDOW_TITLE_DEFAULT );

		return windowTitle;
	}

	double getWindowWidth ( )
	{
		final String windowWidthStr = getProperty ( WINDOW_WIDTH, WINDOW_WIDTH_DEFAULT );

		try
		{
			final int windowWidth = Integer.parseInt ( windowWidthStr );

			return windowWidth;
		}
		catch ( final NumberFormatException e )
		{
			return 720;
		}
	}

	double getWindowHeight ( )
	{
		final String windowHeightStr = getProperty ( WINDOW_HEIGHT, WINDOW_HEIGHT_DEFAULT );

		try
		{
			final int windowHeight = Integer.parseInt ( windowHeightStr );

			return windowHeight;
		}
		catch ( final NumberFormatException e )
		{
			return 480;
		}
	}

	boolean getWindowIsMaximized ( )
	{
		final String windowMaximizedStr = getProperty ( WINDOW_MAXIMIZED, WINDOW_MAXIMIZED_DEFAULT );
		final boolean windowMaximized = parseBoolean ( windowMaximizedStr );

		return windowMaximized;
	}

	boolean getDataIsAutoload ( )
	{
		final String dataAutoloadStr = getProperty ( DATA_AUTOLOAD, DATA_AUTOLOAD_DEFAULT );
		final boolean dataAutoload = parseBoolean ( dataAutoloadStr );

		return dataAutoload;
	}

	private boolean parseBoolean ( final String bStr )
	{
		final String trueStr = "true"; //$NON-NLS-1$
		final String yesStr = "yes"; //$NON-NLS-1$
		final String yStr = "y"; //$NON-NLS-1$
		final String oneStr = "1"; //$NON-NLS-1$

		return bStr.equals ( trueStr ) || bStr.equals ( yesStr ) || bStr.equals ( yStr ) || bStr.equals ( oneStr );
	}
}
