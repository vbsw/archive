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


import java.util.ArrayList;
import java.util.Collections;


/**
 * @author Vitali Baumtrok
 */
final class FileTable
{
	static final ArrayList<FileTableDAO> urlsFileList = new ArrayList<> ( );
	static final FileTableDAO propertyDAO = new FileTableDAO ( Ressource.EXTERNAL_PROP_FILE, State.FILE_NOT_LOADED );
	static final FileTableDAO fxmlDAO = new FileTableDAO ( Ressource.EXTERNAL_FXML_FILE, State.FILE_NOT_LOADED );
	static final FileTableDAO cssDAO = new FileTableDAO ( Ressource.EXTERNAL_CSS_FILE, State.FILE_NOT_LOADED );

	static FileTableDAO selectedItem;

	static void init ( )
	{
		final ArrayList<String> urlsFileNameList = loadURLsFileNameList ( );

		for ( String fileName: urlsFileNameList )
		{
			final FileTableDAO file = new FileTableDAO ( fileName, State.FILE_NOT_LOADED );

			urlsFileList.add ( file );
		}
	}

	static FileTableDAO getFileByName ( final String fileName )
	{
		for ( FileTableDAO file: urlsFileList )
		{
			final String currFileName = file.getName ( );

			if ( currFileName.equals ( fileName ) )
			{
				return file;
			}
		}

		return null;
	}

	private static ArrayList<String> loadURLsFileNameList ( )
	{
		final ArrayList<String> urlsFileNameList = Ressource.loadURLsFilesFromJarDir ( );
		boolean missingDefaultURLsFile = true;

		for ( String fileName: urlsFileNameList )
		{

			if ( fileName.equals ( Ressource.EXTERNAL_URLS_FILE ) )
			{
				missingDefaultURLsFile = false;
				break;
			}
		}

		if ( missingDefaultURLsFile )
		{
			urlsFileNameList.add ( Ressource.EXTERNAL_URLS_FILE );
		}
		Collections.sort ( urlsFileNameList );

		return urlsFileNameList;
	}
}
