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

import javafx.concurrent.Service;
import javafx.concurrent.Task;


/**
 * @author Vitali Baumtrok
 */
final class UrlsFileReader extends Service<UrlsData>
{
	private final String fileName;

	private boolean started = false;
	private volatile boolean running = true;
	private volatile boolean restart = false;

	UrlsFileReader ( final String fileName )
	{
		this.fileName = fileName;
	}

	synchronized void terminate ( )
	{
		if ( started )
		{
			running = false;
			restart = true;
			notify ( );
		}
	}

	synchronized ArrayList<String> getKeyWords ( )
	{
		return decoder.keyWords;
	}

	synchronized ArrayList<String> getURLWords ( )
	{
		return decoder.urlWords;
	}

	synchronized ArrayList<ArrayList<String>> getMaps ( )
	{
		return decoder.maps;
	}

	@Override
	protected Task<UrlsData> createTask ( )
	{
		final UrlsFileDecoder decoder = new UrlsFileDecoder ( fileName );

		return decoder;
	}

	@Override
	public void start ( )
	{
		restart ( );
		if ( started )
		{
			restart = true;

			notify ( );
		}
		else
		{
			started = true;

			super.start ( );
		}
	}

	private void processData ( final byte[] bytes )
	{
		if ( bytes.length > 0 )
		{
			UIController.updateLoadingProgress ( fileName, 0 );

			if ( UrlsFileDecoder.isOnlyWhiteSpace ( bytes ) == false )
			{
				decoder.reset ( bytes );
				decoder.decodeSizes ( );

				int percentageRead = decoder.getPercentageRead ( );

				if ( percentageRead < 100 )
				{
					int percentageReadPrev = 0;

					if ( percentageRead - percentageReadPrev > 0 )
					{
						percentageReadPrev = percentageRead;

						UIController.updateLoadingProgress ( fileName, percentageRead );
					}

					while ( decoder.hasKeyWordsToRead ( ) )
					{
						decoder.decodeNextKeyWord ( );

						percentageRead = decoder.getPercentageRead ( );

						if ( percentageRead - percentageReadPrev > 0 )
						{
							percentageReadPrev = percentageRead;

							UIController.updateLoadingProgress ( fileName, percentageRead );
						}
					}

					while ( decoder.hasURLWordsToRead ( ) )
					{
						decoder.decodeNextURLWord ( );

						percentageRead = decoder.getPercentageRead ( );

						if ( percentageRead - percentageReadPrev > 0 )
						{
							percentageReadPrev = percentageRead;

							UIController.updateLoadingProgress ( fileName, percentageRead );
						}
					}

					while ( decoder.hasMapsToRead ( ) )
					{
						decoder.decodeNextMap ( );

						percentageRead = decoder.getPercentageRead ( );

						if ( percentageRead - percentageReadPrev > 0 )
						{
							percentageReadPrev = percentageRead;

							UIController.updateLoadingProgress ( fileName, percentageRead );
						}
					}
				}
				else
				{
					UIController.updateLoadingProgress ( fileName, percentageRead );
				}
			}
			else
			{
				UIController.updateLoadingProgress ( fileName, 100 );
			}
		}
		else
		{
			UIController.updateLoadingProgress ( fileName, 100 );
		}
	}

	private synchronized void pause ( )
	{
		try
		{
			this.wait ( );
		}
		catch ( final InterruptedException e )
		{}
	}
}
