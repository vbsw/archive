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
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Arrays;
import java.util.Properties;


/**
 * @author Vitali Baumtrok
 */
public class Settings
{
	public static final String EXTERNAL_PROP_FILE = "default.properties"; //$NON-NLS-1$
	public static final String EMPTY_STRING = ""; //$NON-NLS-1$
	public static final Path EXTERNAL_PROP_PATH = getExternalPropertiesPath();

	protected final String[] stringValues;
	protected final int[] intValues;
	protected final float[] floatValues;
	protected final boolean[] booleanValues;
	protected final String[] propertyKeys;
	protected final boolean[] saved;

	public Settings ( final int num )
	{
		stringValues = new String[num];
		intValues = new int[num];
		floatValues = new float[num];
		booleanValues = new boolean[num];
		propertyKeys = new String[num];
		saved = new boolean[num];

		Arrays.fill(stringValues,EMPTY_STRING);
		Arrays.fill(propertyKeys,EMPTY_STRING);
	}

	public void reload ( )
	{
		final Properties properties = readExternalProperties();

		if ( properties!=null )
		{
			for ( int i = 0; i<propertyKeys.length; i += 1 )
			{
				final String propKey = propertyKeys[i];
				final String propVal = properties.getProperty(propKey);

				if ( propVal!=null )
				{
					stringValues[i] = propVal;
					saved[i] = true;
				}
				else
				{
					stringValues[i] = EMPTY_STRING;
					saved[i] = false;
				}
			}
		}
		else
		{
			for ( int i = 0; i<propertyKeys.length; i += 1 )
			{
				stringValues[i] = EMPTY_STRING;
				saved[i] = false;
			}
		}
	}

	public void reset ( )
	{
		reload();
	}

	public Path[] getExternalPropertiesPaths ( )
	{
		return new Path[]
		{
				EXTERNAL_PROP_PATH
		};
	}

	protected Properties readExternalProperties ( )
	{
		final Path[] externalPropPaths = getExternalPropertiesPaths();

		for ( Path externalPropPath: externalPropPaths )
		{
			if ( Files.exists(externalPropPath) )
			{
				try ( final InputStream stream = Files.newInputStream(externalPropPath) )
				{
					final Properties properties = new Properties();

					properties.load(stream);

					return properties;
				}
				catch ( final Exception e )
				{}
			}
		}

		return null;
	}

	private static Path getExternalPropertiesPath ( )
	{
		final Path jarDirectoryPath = JARFile.DIRECTORY;
		final Path externalPropPath = Paths.get(jarDirectoryPath.toString(),EXTERNAL_PROP_FILE);

		return externalPropPath;
	}
}
