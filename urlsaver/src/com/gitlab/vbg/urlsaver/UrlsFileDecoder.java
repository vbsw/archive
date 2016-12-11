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


import java.nio.charset.Charset;
import java.util.ArrayList;

import javafx.concurrent.Task;


/**
 * The SBiF specification can be found at <a href="https://github.com/vitalibaumtrok/sbif">https://github.com/vitalibaumtrok/sbif</a>.
 * <br><br>
 * The encoding of urls file is:
 * <br><code>
 * - FILE ::= SIZE THREE_FRAGMENTS_CONTAINER<br>
 * - THREE_FRAGMENTS_CONTAINER ::= SOFS FS THREE_FRAGMENTS<br>
 * - THREE_FRAGMENTS ::= FRAGMENT_ONE FRAGMENT_TWO FRAGMENT_THREE<br>
 * - FRAGMENT_ONE ::= SOFS FS KEY_WORDS<br>
 * - FRAGMENT_TWO ::= SOFS FS URL_WORDS<br>
 * - FRAGMENT_THREE ::= SOFS FS KEY_MAPS<br>
 * </code>
 * <br><br>
 * Default empty file:
 * <br><code>
 * new byte[]{56,0,0,0,0,0,0,0, 24,0,0,0,0,0,0,0, 8,0,0,0,0,0,0,0, 8,0,0,0,0,0,0,0, 8,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0, 0,0,0,0,0,0,0,0};
 * </code>
 * 
 * @author Vitali Baumtrok
 */
class UrlsFileDecoder extends Task<UrlsData>
{
	private static final Charset UTF8_CHARSET = Charset.forName ( "UTF-8" );
	private static final byte NUL_BYTE = 0;
	private static final byte TAB_BYTE = 9;
	private static final byte LF_BYTE = 10;
	private static final byte CR_BYTE = 13;
	private static final byte SP_BYTE = 32;
	private static final byte ZERO_BYTE = 48;

	private final UrlsData urlsData = new UrlsData ( );
	private final String fileName;

	private byte[] bytes;
	private int bytesRead;
	private int totalSize;
	private int fragmentSizeOne;
	private int fragmentSizeTwo;
	private int fragmentOneSOFS;
	private int fragmentTwoSOFS;
	private int fragmentThreeSOFS;
	private int fragmentOneSOFSRead;
	private int fragmentTwoSOFSRead;
	private int fragmentThreeSOFSRead;
	private int fragmentOneDataRead;
	private int fragmentTwoDataRead;
	private int fragmentThreeDataRead;

	UrlsFileDecoder ( final String fileName )
	{
		this.fileName = fileName;
	}

	static boolean isOnlyWhiteSpace ( final byte[] bytes )
	{
		for ( byte c: bytes )
		{
			if ( c != NUL_BYTE || c != TAB_BYTE || c != LF_BYTE || c != CR_BYTE || c != SP_BYTE || c != ZERO_BYTE )
			{
				return false;
			}
		}
		return true;
	}

	void reset ( final byte[] bytes )
	{
		this.bytes = bytes;
		this.bytesRead = 0;
		this.fragmentOneSOFSRead = 0;
		this.fragmentTwoSOFSRead = 0;
		this.fragmentThreeSOFSRead = 0;
		this.fragmentOneDataRead = 0;
		this.fragmentTwoDataRead = 0;
		this.fragmentThreeDataRead = 0;

		this.keyWords.clear ( );
		this.urlWords.clear ( );
		this.mapsPool.addAll ( this.maps );
		this.maps.clear ( );
	}

	int decodeSizes ( )
	{
		if ( bytes.length > 64 )
		{
			totalSize = bytesToInt ( bytes, 0 ) + 8;
			totalSize = bytes.length >= totalSize ? totalSize : bytes.length;
			fragmentSizeOne = bytesToInt ( bytes, 16 );
			fragmentSizeTwo = bytesToInt ( bytes, 24 );
			// fragmentSizeTree = bytesToInt ( bytes, 32 );
			fragmentOneSOFS = bytesToInt ( bytes, 40 );
			fragmentTwoSOFS = bytesToInt ( bytes, 40 + fragmentSizeOne );
			fragmentThreeSOFS = bytesToInt ( bytes, 40 + fragmentSizeOne + fragmentSizeTwo );

			bytesRead = totalSize > 64 ? 64 : totalSize;
		}
		else
		{
			totalSize = bytes.length;
			bytesRead = bytes.length;
		}

		return bytesRead;
	}

	int getPercentageRead ( )
	{
		final double bytesReadDouble = (double) bytesRead;
		final double bytesLengthDouble = (double) totalSize;
		final double readDouble = bytesLengthDouble / bytesReadDouble;
		final int percentageRead = (int) ( readDouble * 100 );

		return percentageRead;
	}

	boolean hasKeyWordsToRead ( )
	{
		return fragmentOneSOFSRead < fragmentOneSOFS;
	}

	boolean hasURLWordsToRead ( )
	{
		return fragmentTwoSOFSRead < fragmentTwoSOFS;
	}

	boolean hasMapsToRead ( )
	{
		return fragmentThreeSOFSRead < fragmentThreeSOFS;
	}

	void decodeNextKeyWord ( )
	{
		final int fsOffset = 48 + fragmentOneSOFSRead;
		final int fragmentSize = bytesToInt ( bytes, fsOffset );
		final int fragmentOffset = 48 + fragmentOneSOFS + fragmentOneDataRead;
		final String keyWord = new String ( bytes, fragmentOffset, fragmentSize, UTF8_CHARSET );

		fragmentOneSOFSRead += 8;
		fragmentOneDataRead += fragmentSize;
		bytesRead += ( 8 + fragmentSize );

		keyWords.add ( keyWord );
	}

	void decodeNextURLWord ( )
	{
		final int fsOffset = 40 + fragmentSizeOne + 8 + fragmentTwoSOFSRead;
		final int fragmentSize = bytesToInt ( bytes, fsOffset );
		final int fragmentOffset = 40 + fragmentSizeOne + 8 + fragmentTwoSOFS + fragmentTwoDataRead;
		final String urlWord = new String ( bytes, fragmentOffset, fragmentSize, UTF8_CHARSET );

		fragmentTwoSOFSRead += 8;
		fragmentTwoDataRead += fragmentSize;
		bytesRead += ( 8 + fragmentSize );

		urlWords.add ( urlWord );
	}

	void decodeNextMap ( )
	{
		final int fsOffset = 40 + fragmentSizeOne + fragmentSizeTwo + 8 + fragmentThreeSOFSRead;
		final int fragmentSize = bytesToInt ( bytes, fsOffset );
		final int fragmentOffset = 40 + fragmentSizeOne + fragmentSizeTwo + 8 + fragmentThreeSOFS + fragmentThreeDataRead;
		final int fragmentEnd = fragmentOffset + fragmentSize;
		final ArrayList<String> map = getNewMap ( );

		for ( int i = fragmentOffset; i < fragmentEnd; i += 8 )
		{
			final int urlWordIndex = bytesToInt ( bytes, i );
			final String urlWord = urlWords.get ( urlWordIndex );

			map.add ( urlWord );
		}

		fragmentThreeSOFSRead += 8;
		fragmentThreeDataRead += fragmentSize;
		bytesRead += ( 8 + fragmentSize );

		maps.add ( map );
	}

	private int bytesToInt ( final byte[] bytes, final int offset )
	{
		try
		{
			final byte byte0 = bytes[offset];
			final byte byte1 = bytes[offset + 1];
			final byte byte2 = bytes[offset + 2];
			final byte byte3 = bytes[offset + 3];

			return ( byte0 & 0xFF ) | ( byte1 & 0xFF ) << 8 | ( byte2 & 0xFF ) << 16 | ( byte3 & 0xFF ) << 24;
		}
		catch ( final Exception e )
		{
			e.printStackTrace ( );

			return 0;
		}
	}

	private ArrayList<String> getNewMap ( )
	{
		final int size = mapsPool.size ( );

		if ( size == 0 )
		{
			return new ArrayList<String> ( );
		}
		else
		{
			final ArrayList<String> lastMap = mapsPool.remove ( size - 1 );

			lastMap.clear ( );

			return lastMap;
		}
	}

	@Override
	protected UrlsData call ( ) throws Exception
	{
		// TODO Auto-generated method stub
		return null;
	}
}
