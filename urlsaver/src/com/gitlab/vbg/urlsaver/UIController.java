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


import java.awt.Desktop;
import java.net.MalformedURLException;
import java.net.URI;
import java.net.URL;

import javafx.stage.WindowEvent;


/**
 * @author Vitali Baumtrok
 */
public class UIController
{
	static void loadInitialUrlsFiles ( )
	{
		if ( Settings.data.isAutoload ( ) )
		{
			for ( FileTableDAO file: FileTable.urlsFileList )
			{
				if ( file.isAvailable ( ) )
				{
					file.reader.start ( );
				}
			}
		}
	}

	static void handleSelectedFileItem ( final FileTableDAO selectedFile )
	{
		FileTable.selectedItem = selectedFile;

		if ( selectedFile.isUrlsExtension ( ) )
		{
			URLTable.file = selectedFile;
		}
		UI.refreshTabs ( );
		UI.refreshFileDetails ( );
	}

	static void handleSelectedURLItem ( final URLTableDAO selectedURL )
	{
		// TODO handleSelectedURLItem
	}

	static void handleCreateSelectedFile ( )
	{
		if ( isWrittenFileEqualToSelected ( ) )
		{
			final FileTableDAO fileTableDAO = FileTable.selectedItem;

			if ( fileTableDAO.isAvailable ( ) == false )
			{
				Ressource.createFile ( fileTableDAO );

				if ( fileTableDAO.isAvailable ( ) )
				{
					fileTableDAO.setState ( State.FILE_AVAILABLE );
					UI.refreshFileItem ( fileTableDAO, true );
				}
			}
		}
		else
		{
			// TODO: create written URLs file (update of FileTable)
		}
	}

	static void handeReloadFile ( )
	{
		if ( isWrittenFileEqualToSelected ( ) )
		{
			FileTable.selectedItem.refreshNeeded = true;

			reloadSelectedFile ( );
		}
		else
		{
			// TODO: reload written URLs file (update of FileTable)
		}
	}

	static void handleDeleteSelectedFile ( )
	{
		if ( isWrittenFileEqualToSelected ( ) )
		{
			final FileTableDAO selectedFile = FileTable.selectedItem;

			if ( selectedFile.isAvailable ( ) )
			{
				Ressource.deleteFile ( selectedFile );

				if ( selectedFile.isAvailable ( ) == false )
				{
					selectedFile.setState ( State.FILE_MISSING );
					UI.refreshFileItem ( selectedFile, true );
				}
			}
		}
		else
		{
			// TODO: deletion of URLs file (update of FileTable)
		}
	}

	static void handleOpenInBroser ( )
	{
		if ( UI.openInBrowserButton.isDisabled ( ) == false )
		{
			openURLInBrowser ( );
		}
	}

	static void handleURLTextChange ( )
	{
		final URI uri = getURIFromTextField ( );

		UI.openInBrowserButton.setDisable ( uri == null );
	}

	static void handleQuitApplication ( )
	{
		final WindowEvent closeEvent = new WindowEvent ( UI.stage, WindowEvent.WINDOW_CLOSE_REQUEST );

		UI.stage.fireEvent ( closeEvent );
	}

	static void handeCloseWindow ( final WindowEvent event )
	{
		for ( FileTableDAO file: FileTable.urlsFileList )
		{
			file.reader.terminate ( );
		}
	}

	static void handleSelectedURLTab ( )
	{
		if (URLTable.file.refreshNeeded)
		{
			URLTable.file.refresh();
		}
		UI.refreshURLs ( );
	}

	static synchronized void updateLoadingProgress ( final String fileName, final int loadedPercent )
	{
		final FileTableDAO file = FileTable.getFileByName ( fileName );

		if ( file != null )
		{
			if ( loadedPercent >= 0 )
			{
				final String loaded = String.valueOf ( loadedPercent );

				file.setLoaded ( loaded );
			}
			else
			{
				file.setLoaded ( "0" );
			}
			UI.refreshTabs ( );
			UI.refreshFileItem ( file, false );
		}
	}

	private static void reloadSelectedFile ( )
	{
		final FileTableDAO selectedFile = FileTable.selectedItem;
		final String fileName = selectedFile.getName ( );

		if ( fileName.equals ( Ressource.EXTERNAL_PROP_FILE ) )
		{
			Settings.reload ( );
			UI.applySettings ( );
			selectedFile.setLoaded ( State.FILE_LOADED );
			UI.refreshFileItem ( selectedFile, true );
		}
		else if ( fileName.equals ( Ressource.EXTERNAL_FXML_FILE ) )
		{
			UI.reloadFXML ( );
			selectedFile.setLoaded ( State.FILE_LOADED );
			UI.refreshFileItem ( selectedFile, true );
		}
		else if ( fileName.equals ( Ressource.EXTERNAL_CSS_FILE ) )
		{
			UI.reloadCSS ( );
			selectedFile.setLoaded ( State.FILE_LOADED );
			UI.refreshFileItem ( selectedFile, true );
		}
		else if ( selectedFile.isUrlsExtension ( ) )
		{
			selectedFile.reader.start ( );
		}
	}

	private static boolean isWrittenFileEqualToSelected ( )
	{
		final FileTableDAO fileTableDAO = FileTable.selectedItem;
		final String selectedFile = fileTableDAO.getName ( );
		final String writtenFile = UI.fileTextField.getText ( );

		return writtenFile.equals ( selectedFile );
	}

	private static void openURLInBrowser ( )
	{
		if ( Desktop.isDesktopSupported ( ) )
		{
			final Desktop desktop = Desktop.getDesktop ( );

			if ( desktop.isSupported ( Desktop.Action.BROWSE ) )
			{
				final URI uri = getURIFromTextField ( );

				if ( uri != null )
				{
					try
					{
						desktop.browse ( uri );
					}
					catch ( final Exception e )
					{}
				}
			}
		}
	}

	private static URI getURIFromTextField ( )
	{
		final URL url = getURLFromTextField ( );

		if ( url != null )
		{
			try
			{
				return url.toURI ( );
			}
			catch ( final Exception e )
			{}
		}

		return null;
	}

	private static URL getURLFromTextField ( )
	{
		final String urlStr = UI.urlTextField.getText ( );

		if ( urlStr != null && urlStr.length ( ) > 0 )
		{
			try
			{
				return new URL ( urlStr );
			}
			catch ( final MalformedURLException e )
			{}

			try
			{
				final String httpPrefix = "http://"; //$NON-NLS-1$
				final String httpURLStr = httpPrefix + urlStr;

				return new URL ( httpURLStr );
			}
			catch ( final MalformedURLException e )
			{}
		}

		return null;
	}
}
