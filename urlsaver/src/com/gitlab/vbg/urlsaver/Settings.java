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
final class Settings
{
	static final Window window = new Window ( );
	static final Data data = new Data ( );

	static void reload ( )
	{
		final Properties properties = Ressource.loadProperties ( );

		window.title = properties.getWindowTitle ( );
		window.width = properties.getWindowWidth ( );
		window.height = properties.getWindowHeight ( );
		window.maximized = properties.getWindowIsMaximized ( );
		data.autoload = properties.getDataIsAutoload ( );
	}

	static final class Window
	{
		private String title = "";
		private double width = 0;
		private double height = 0;
		private boolean maximized = false;

		String getTitle ( )
		{
			return title;
		}

		double getWidth ( )
		{
			return width;
		}

		double getHeight ( )
		{
			return height;
		}

		boolean isMaximized ( )
		{
			return maximized;
		}
	}

	static final class Data
	{
		private boolean autoload = false;

		boolean isAutoload ( )
		{
			return autoload;
		}
	}
}
