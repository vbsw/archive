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


import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;

import javafx.beans.property.SimpleStringProperty;


/**
 * @author Vitali Baumtrok
 */
public class FileTableDAO
{
	private final SimpleStringProperty name;
	private final SimpleStringProperty state;
	private final SimpleStringProperty loaded;

	final UrlsFileReader reader;
	final Path path;
	final String nameWidthouExtension;

	final ArrayList<String> keyWords = new ArrayList<> ( );
	final ArrayList<String> urlWords = new ArrayList<> ( );
	final ArrayList<ArrayList<String>> maps = new ArrayList<> ( );

	boolean refreshNeeded;

	public FileTableDAO ( final String nameStr, final String readStr )
	{
		path = Paths.get ( Ressource.jarDirectoryPath.toString ( ), nameStr );
		name = new SimpleStringProperty ( nameStr );
		loaded = new SimpleStringProperty ( readStr );
		reader = new UrlsFileReader ( nameStr );
		nameWidthouExtension = removeExtension ( nameStr );
		refreshNeeded = true;

		if ( isAvailable ( ) )
		{
			state = new SimpleStringProperty ( State.FILE_AVAILABLE );
		}
		else
		{
			state = new SimpleStringProperty ( State.FILE_MISSING );
		}
	}

	public SimpleStringProperty nameProperty ( )
	{
		return name;
	}

	public SimpleStringProperty stateProperty ( )
	{
		return state;
	}

	public SimpleStringProperty loadedProperty ( )
	{
		return loaded;
	}

	public FileTableDAO ( )
	{
		this ( "", "" );
	}

	public String getName ( )
	{
		return name.get ( );
	}

	public void setName ( final String nameStr )
	{
		name.set ( nameStr );
	}

	public String getState ( )
	{
		return state.get ( );
	}

	public void setState ( final String stateStr )
	{
		state.set ( stateStr );
	}

	public String getLoaded ( )
	{
		return loaded.get ( );
	}

	public void setLoaded ( final String loadedStr )
	{
		loaded.set ( loadedStr );
	}

	public boolean isAvailable ( )
	{
		return Files.exists ( path );
	}

	public boolean isUrlsExtension ( )
	{
		final String fileName = name.get ( );

		return UrlsFile.isValidExtension ( fileName );
	}

	public String getNameWithoutExtension ( )
	{
		return nameWidthouExtension;
	}

	public void refresh ( )
	{
		final ArrayList<String> newKeyWords = reader.getKeyWords ( );
		final ArrayList<String> newURLWords = reader.getURLWords ( );
		final ArrayList<ArrayList<String>> newMaps = reader.getMaps ( );

		keyWords.clear ( );
		urlWords.clear ( );
		maps.clear ( );

		keyWords.addAll ( newKeyWords );
		urlWords.addAll ( newURLWords );
		maps.addAll ( newMaps );
	}

	private String removeExtension ( final String nameStr )
	{
		int i = nameStr.length ( ) - 1;

		while ( i >= 0 )
		{
			if ( nameStr.charAt ( i ) == '.' )
			{
				return nameStr.substring ( 0, i );
			}
			i -= 1;
		}

		return nameStr;
	}
}
