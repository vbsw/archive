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


import javafx.beans.property.SimpleStringProperty;


/**
 * @author Vitali Baumtrok
 */
public class URLTableDAO
{
	private final SimpleStringProperty url = new SimpleStringProperty ( );
	private final SimpleStringProperty edited = new SimpleStringProperty ( );

	public SimpleStringProperty urlProperty ( )
	{
		return url;
	}

	public SimpleStringProperty editedProperty ( )
	{
		return edited;
	}

	public String getUrl ( )
	{
		return url.get ( );
	}

	public void setUrl ( final String urlStr )
	{
		url.set ( urlStr );
	}

	public String getEdited ( )
	{
		return edited.get ( );
	}

	public void setEdited ( final String editedStr )
	{
		edited.set ( editedStr );
	}
}
