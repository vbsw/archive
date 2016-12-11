/**
 * com.gitlab.vbit.util (helper classes)
 * Copyright 2016 Vitali Baumtrok
 * 
 * This file is part of com.gitlab.vbit.util.
 * 
 * com.gitlab.vbit.util is free software: you can redistribute it and/or
 * modify it under the terms of the GNU General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 * 
 * com.gitlab.vbit.util is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 * 
 * You should have received a copy of the GNU General Public License
 * along with com.gitlab.vbit.util.  If not,
 * see <http://www.gnu.org/licenses/>.
 */


package com.gitlab.vbit.util;


import java.io.InputStream;
import java.net.MalformedURLException;
import java.net.URI;
import java.net.URISyntaxException;
import java.net.URL;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;


/**
 * @author Vitali Baumtrok
 */
public class JARFile
{
	public static final Path DIRECTORY = getDirectory();

	private static Path getDirectory ( )
	{
		try
		{
			final String classFileExtension = ".class"; //$NON-NLS-1$
			final String jarPrefix = "jar:"; //$NON-NLS-1$
			final String className = JARFile.class.getSimpleName()+classFileExtension;
			final URL classUrl = JARFile.class.getResource(className);
			final String classUrlStr = classUrl.toString();

			if ( classUrlStr.startsWith(jarPrefix) )
			{
				return getJarDirectory(classUrlStr);
			}
			else
			{
				return getDevelopmentBinDirecotry(classUrlStr);
			}
		}
		catch ( final Exception e )
		{
			e.printStackTrace();

			return null;
		}
	}

	private static Path getJarDirectory ( final String classUrlStr ) throws MalformedURLException, URISyntaxException
	{
		final String canonicalClassName = JARFile.class.getCanonicalName();
		final String classFileExtension = ".class"; //$NON-NLS-1$
		final int classPathLength = canonicalClassName.length()+classFileExtension.length();
		final String jarDirStr = classUrlStr.substring(4,classUrlStr.length()-classPathLength-2);
		final URL jarDirUrl = new URL(jarDirStr);
		final URI jarDirUri = jarDirUrl.toURI();
		final Path path = Paths.get(jarDirUri);

		return path.getParent();
	}

	private static Path getDevelopmentBinDirecotry ( final String classUrlStr ) throws MalformedURLException, URISyntaxException
	{
		final String canonicalClassName = JARFile.class.getCanonicalName();
		final String classFileExtension = ".class"; //$NON-NLS-1$
		final int classPathLength = canonicalClassName.length()+classFileExtension.length();
		final String binDir = "/bin"; //$NON-NLS-1$
		final String jarDirStr = classUrlStr.substring(0,classUrlStr.length()-classPathLength-binDir.length());
		final URL jarDirUrl = new URL(jarDirStr);
		final URI jarDirUri = jarDirUrl.toURI();
		final Path path = Paths.get(jarDirUri);

		return path;
	}

	public static InputStream getInputStreamFromInsideJar ( final String pathStr )
	{
		final ClassLoader classLoader = JARFile.class.getClassLoader();
		final InputStream stream = classLoader.getResourceAsStream(pathStr);

		return stream;
	}

	public static void copyFromJar ( final String from, final Path to )
	{
		try ( final InputStream inStream = getInputStreamFromInsideJar(from) )
		{
			Files.copy(inStream,to);
		}
		catch ( final Exception e )
		{
			e.printStackTrace();
		}
	}
}
