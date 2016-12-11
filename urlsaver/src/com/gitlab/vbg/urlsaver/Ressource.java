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


import java.io.IOException;
import java.io.InputStream;
import java.net.URI;
import java.net.URL;
import java.nio.file.DirectoryStream;
import java.nio.file.Files;
import java.nio.file.NoSuchFileException;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;

import javafx.fxml.FXMLLoader;
import javafx.scene.Parent;


/**
 * @author Vitali Baumtrok
 */
final class Ressource
{
	static final String EXTERNAL_PROP_FILE = "urlsaver.properties"; //$NON-NLS-1$
	static final String EXTERNAL_FXML_FILE = "urlsaver.fxml"; //$NON-NLS-1$
	static final String EXTERNAL_CSS_FILE = "urlsaver.css"; //$NON-NLS-1$
	static final String EXTERNAL_URLS_FILE = "default.urls"; //$NON-NLS-1$

	static final Path jarDirectoryPath = getJarDirectory ( );
	static final Path externalPropPath = Paths.get ( jarDirectoryPath.toString ( ), EXTERNAL_PROP_FILE );
	static final Path externalFXMLPath = Paths.get ( jarDirectoryPath.toString ( ), EXTERNAL_FXML_FILE );
	static final Path externalCSSPath = Paths.get ( jarDirectoryPath.toString ( ), EXTERNAL_CSS_FILE );
	static final Path externalUrlsFilePath = Paths.get ( jarDirectoryPath.toString ( ), EXTERNAL_URLS_FILE );

	private static final String INTERNAL_PROP_FILE = "com/gitlab/vbg/urlsaver/default.properties"; //$NON-NLS-1$
	private static final String INTERNAL_FXML_FILE = "com/gitlab/vbg/urlsaver/default.fxml"; //$NON-NLS-1$
	private static final String INTERNAL_CSS_FILE = "com/gitlab/vbg/urlsaver/default.css"; //$NON-NLS-1$
	private static final String INTERNAL_URLS_FILE = "com/gitlab/vbg/urlsaver/default.urls"; //$NON-NLS-1$

	static Properties loadProperties ( )
	{
		final Properties externalProperties = loadExternalProperties ( );

		if ( externalProperties != null )
		{
			FileTable.propertyDAO.setState ( State.FILE_AVAILABLE );
			FileTable.propertyDAO.setLoaded ( State.FILE_LOADED );

			return externalProperties;
		}
		else
		{
			final Properties internalProperties = loadInternalProperties ( );

			FileTable.propertyDAO.setState ( State.FILE_MISSING );
			FileTable.propertyDAO.setLoaded ( State.FILE_NOT_LOADED );

			return internalProperties;
		}
	}

	static Parent loadFXML ( )
	{
		final Parent externalFXML = loadExternalFXML ( );

		if ( externalFXML != null )
		{
			FileTable.fxmlDAO.setState ( State.FILE_AVAILABLE );
			FileTable.fxmlDAO.setLoaded ( State.FILE_LOADED );

			return externalFXML;
		}
		else
		{
			final Parent internalFXML = loadInternalFXML ( );

			FileTable.fxmlDAO.setState ( State.FILE_MISSING );
			FileTable.fxmlDAO.setLoaded ( State.FILE_NOT_LOADED );

			return internalFXML;
		}
	}

	static String loadCSSURL ( )
	{
		final String externalStylesheetURL = getExternalStylesheetURL ( );

		if ( externalStylesheetURL != null )
		{
			FileTable.cssDAO.setState ( State.FILE_AVAILABLE );
			FileTable.cssDAO.setLoaded ( State.FILE_LOADED );

			return externalStylesheetURL;
		}
		else
		{
			final String internalStylesheetURL = getInternalStylesheetURL ( );

			FileTable.cssDAO.setState ( State.FILE_MISSING );
			FileTable.cssDAO.setLoaded ( State.FILE_NOT_LOADED );

			return internalStylesheetURL;
		}
	}

	static void createFile ( final FileTableDAO file )
	{
		final String fileName = file.getName ( );

		if ( fileName.equals ( EXTERNAL_PROP_FILE ) )
		{
			copyFromJar ( INTERNAL_PROP_FILE, externalPropPath );
		}
		else if ( fileName.equals ( EXTERNAL_FXML_FILE ) )
		{
			copyFromJar ( INTERNAL_FXML_FILE, externalFXMLPath );
		}
		else if ( fileName.equals ( EXTERNAL_CSS_FILE ) )
		{
			copyFromJar ( INTERNAL_CSS_FILE, externalCSSPath );
		}
		else if ( file.isUrlsExtension ( ) )
		{
			final Path path = Paths.get ( jarDirectoryPath.toString ( ), fileName );

			copyFromJar ( INTERNAL_URLS_FILE, path );
		}
	}

	static void deleteFile ( final FileTableDAO file )
	{
		try
		{
			Files.delete ( file.path );
			file.setLoaded ( State.FILE_NOT_LOADED );
		}
		catch ( final NoSuchFileException e )
		{
			file.setLoaded ( State.FILE_NOT_LOADED );
		}
		catch ( IOException e )
		{
			e.printStackTrace ( );
		}
	}

	static ArrayList<String> loadURLsFilesFromJarDir ( )
	{
		final ArrayList<String> list = new ArrayList<> ( );

		try ( DirectoryStream<Path> directoryStream = Files.newDirectoryStream ( jarDirectoryPath ) )
		{
			for ( Path file: directoryStream )
			{
				final Path fileNamePath = file.getFileName ( );
				final String fileName = fileNamePath.toString ( );

				if ( UrlsFile.isValidExtension ( fileName ) )
				{
					list.add ( fileName );
				}
			}
		}
		catch ( final Exception e )
		{
			e.printStackTrace ( );
		}

		return list;
	}

	static byte[] readFile ( final String fileName )
	{
		final Path path = Paths.get ( jarDirectoryPath.toString ( ), fileName );

		try
		{
			final byte[] bytes = Files.readAllBytes ( path );

			return bytes;
		}
		catch ( IOException e )
		{
			return null;
		}
	}

	static boolean writeFile ( final String fileName, final byte[] bytes )
	{
		final Path path = Paths.get ( jarDirectoryPath.toString ( ), fileName );

		try
		{
			Files.write ( path, bytes );

			return true;
		}
		catch ( IOException e )
		{
			e.printStackTrace ( );

			return false;
		}
	}

	private static Properties loadExternalProperties ( )
	{
		if ( Files.exists ( externalPropPath ) )
		{
			try ( final InputStream stream = Files.newInputStream ( externalPropPath ) )
			{
				final Properties properties = new Properties ( );

				properties.load ( stream );

				return properties;
			}
			catch ( final Exception e )
			{}
		}

		return null;
	}

	private static Properties loadInternalProperties ( )
	{
		try ( InputStream stream = getInputStreamFromInsideJar ( INTERNAL_PROP_FILE ) )
		{
			final Properties properties = new Properties ( );

			properties.load ( stream );

			return properties;
		}
		catch ( final Exception e )
		{
			e.printStackTrace ( );

			return null;
		}
	}

	private static Parent loadExternalFXML ( )
	{
		if ( Files.exists ( externalFXMLPath ) )
		{
			try ( final InputStream stream = Files.newInputStream ( externalFXMLPath ) )
			{
				final FXMLLoader fxmlLoader = new FXMLLoader ( );
				final Parent fxml = fxmlLoader.load ( stream );

				return fxml;
			}
			catch ( final Exception e )
			{}
		}

		return null;
	}

	private static Parent loadInternalFXML ( )
	{
		try ( InputStream stream = getInputStreamFromInsideJar ( INTERNAL_FXML_FILE ) )
		{
			final FXMLLoader fxmlLoader = new FXMLLoader ( );
			final Parent fxml = fxmlLoader.load ( stream );

			return fxml;
		}
		catch ( final Exception e )
		{
			e.printStackTrace ( );

			return null;
		}
	}

	private static String getExternalStylesheetURL ( )
	{
		if ( Files.exists ( externalCSSPath ) )
		{
			try
			{
				final ClassLoader classLoader = Ressource.class.getClassLoader ( );
				final URL url = classLoader.getResource ( EXTERNAL_CSS_FILE );
				final String urlStr = url.toExternalForm ( );

				return urlStr;
			}
			catch ( final Exception e )
			{}
		}

		return null;
	}

	private static String getInternalStylesheetURL ( )
	{
		try
		{
			final ClassLoader classLoader = Ressource.class.getClassLoader ( );
			final URL url = classLoader.getResource ( INTERNAL_CSS_FILE );
			final String urlStr = url.toExternalForm ( );

			return urlStr;
		}
		catch ( final Exception e )
		{
			e.printStackTrace ( );

			return null;
		}
	}

	private static InputStream getInputStreamFromInsideJar ( final String pathStr )
	{
		final ClassLoader classLoader = Ressource.class.getClassLoader ( );
		final InputStream stream = classLoader.getResourceAsStream ( pathStr );

		return stream;
	}

	private static Path getJarDirectory ( )
	{
		try
		{
			final String classFileExtension = ".class"; //$NON-NLS-1$
			final String jarPrefix = "jar:"; //$NON-NLS-1$
			final String className = Ressource.class.getSimpleName ( ) + classFileExtension;
			final URL classUrl = Ressource.class.getResource ( className );
			final String classUrlStr = classUrl.toString ( );

			if ( classUrlStr.startsWith ( jarPrefix ) )
			{
				final String jarDirStr = classUrlStr.substring ( 4, classUrlStr.length ( ) - 41 );
				final URL jarDirUrl = new URL ( jarDirStr );
				final URI jarDirUri = jarDirUrl.toURI ( );
				final Path path = Paths.get ( jarDirUri );

				return path.getParent ( );
			}
			else
			{
				final String jarDirStr = classUrlStr.substring ( 0, classUrlStr.length ( ) - 40 );
				final URL jarDirUrl = new URL ( jarDirStr );
				final URI jarDirUri = jarDirUrl.toURI ( );
				final Path path = Paths.get ( jarDirUri );

				return path;
			}
		}
		catch ( final Exception e )
		{
			e.printStackTrace ( );

			return null;
		}
	}

	private static void copyFromJar ( final String from, final Path to )
	{
		try ( final InputStream inStream = getInputStreamFromInsideJar ( from ) )
		{
			Files.copy ( inStream, to );
		}
		catch ( final Exception e )
		{
			e.printStackTrace ( );
		}
	}
}
