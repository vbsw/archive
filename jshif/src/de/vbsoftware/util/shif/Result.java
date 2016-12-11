/**
 * JSHiF - a program to process documents in Simple Hierarchal Format
 * Copyright 2015 Vitali Baumtrok
 * 
 * This file is part of JSHiF.
 *
 * JSHiF is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * JSHiF is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */


package de.vbsoftware.util.shif;


import java.util.ArrayList;


/**
 * @author Vitali Baumtrok
 */
public class Result
{

	protected final Analyzer analyzer;

	protected int nameBegin = 0;
	protected int nameEnd = 0;
	protected int valueBegin = 0;
	protected int valueEnd = 0;

	public Result ( final Analyzer analyzer )
	{
		this.analyzer = analyzer;
	}

	public boolean isValid ( )
	{
		return !( analyzer.invalidIndentDepth || analyzer.invalidNodeName );
	}

	public boolean isInvalidIndentDepth ( )
	{
		return analyzer.invalidIndentDepth;
	}

	public boolean isInvalidNodeName ( )
	{
		return analyzer.invalidNodeName;
	}

	public int getLineNumber ( )
	{
		return analyzer.lineNumber;
	}

	public int getIndentDelta ( )
	{
		return analyzer.indentDepthCurr - analyzer.indentDepthPrev;
	}

	public String getNodeNameAsString ( )
	{
		return analyzer.text.substring ( nameBegin, nameEnd );
	}

	public String getNodeValueAsString ( )
	{
		return analyzer.text.substring ( valueBegin, valueEnd );
	}

	public int getNameAsInteger ( ) throws NumberFormatException
	{
		return Integer.parseInt ( analyzer.text.substring ( nameBegin, nameEnd ) );
	}

	public int getValueAsInteger ( ) throws NumberFormatException
	{
		return Integer.parseInt ( analyzer.text.substring ( valueBegin, valueEnd ) );
	}

	public float getNameAsFloat ( ) throws NumberFormatException
	{
		return Float.parseFloat ( analyzer.text.substring ( nameBegin, nameEnd ) );
	}

	public float getValueAsFloat ( ) throws NumberFormatException
	{
		return Float.parseFloat ( analyzer.text.substring ( valueBegin, valueEnd ) );
	}

	public ArrayList<String> getNameAsStringList ( final char separator )
	{
		return getStringList ( nameBegin, nameEnd, separator );
	}

	public ArrayList<String> getValueAsStringList ( final char separator )
	{
		return getStringList ( valueBegin, valueEnd, separator );
	}

	public ArrayList<Integer> getNameAsIntegerList ( final char separator ) throws NumberFormatException
	{
		return getIntegerList ( nameBegin, nameEnd, separator );
	}

	public ArrayList<Integer> getValueAsIntegerList ( final char separator ) throws NumberFormatException
	{
		return getIntegerList ( valueBegin, valueEnd, separator );
	}

	public ArrayList<Float> getNameAsFloatList ( final char separator ) throws NumberFormatException
	{
		return getFloatList ( nameBegin, nameEnd, separator );
	}

	public ArrayList<Float> getValueAsFloatList ( final char separator ) throws NumberFormatException
	{
		return getFloatList ( valueBegin, valueEnd, separator );
	}

	public boolean isNewSubnode ( )
	{
		return analyzer.indentDepthCurr == analyzer.indentDepthPrev + 1;
	}

	public boolean isNewNode ( )
	{
		return analyzer.indentDepthCurr == analyzer.indentDepthPrev;
	}

	public boolean isNewTopnode ( )
	{
		return analyzer.indentDepthCurr < analyzer.indentDepthPrev;
	}

	private ArrayList<String> getStringList ( int iBegin, int iEnd, final char separator )
	{
		final ArrayList<String> list = new ArrayList<> ( );
		final String text = analyzer.text;

		while ( true )
		{
			if ( iEnd >= text.length ( ) )
			{
				if ( iBegin < iEnd )
				{
					list.add ( text.substring ( iBegin, iEnd ) );
				}
				break;
			}
			else
			{
				final char c = text.charAt ( iEnd );

				if ( c == separator )
				{
					if ( iBegin < iEnd )
					{
						list.add ( text.substring ( iBegin, iEnd ) );
					}
					iEnd++;
					iBegin = iEnd;
				}
				else
				{
					iEnd++;
				}
			}
		}

		return list;
	}

	private ArrayList<Integer> getIntegerList ( int iBegin, int iEnd, final char separator ) throws NumberFormatException
	{
		final ArrayList<Integer> list = new ArrayList<> ( );
		final String text = analyzer.text;

		while ( true )
		{
			if ( iEnd >= text.length ( ) )
			{
				if ( iBegin < iEnd )
				{
					final Integer integer = Integer.parseInt ( text.substring ( iBegin, iEnd ) );

					list.add ( integer );
				}
				break;
			}
			else
			{
				final char c = text.charAt ( iEnd );

				if ( c == separator )
				{
					if ( iBegin < iEnd )
					{
						final Integer integer = Integer.parseInt ( text.substring ( iBegin, iEnd ) );

						list.add ( integer );
					}
					iEnd++;
					iBegin = iEnd;
				}
				else
				{
					iEnd++;
				}
			}
		}

		return list;
	}

	private ArrayList<Float> getFloatList ( int iBegin, int iEnd, final char separator ) throws NumberFormatException
	{
		final ArrayList<Float> list = new ArrayList<> ( );
		final String text = analyzer.text;

		while ( true )
		{
			if ( iEnd >= text.length ( ) )
			{
				if ( iBegin < iEnd )
				{
					final Float f = Float.parseFloat ( text.substring ( iBegin, iEnd ) );

					list.add ( f );
				}
				break;
			}
			else
			{
				final char c = text.charAt ( iEnd );

				if ( c == separator )
				{
					if ( iBegin < iEnd )
					{
						final Float f = Float.parseFloat ( text.substring ( iBegin, iEnd ) );

						list.add ( f );
					}
					iEnd++;
					iBegin = iEnd;
				}
				else
				{
					iEnd++;
				}
			}
		}

		return list;
	}

}
