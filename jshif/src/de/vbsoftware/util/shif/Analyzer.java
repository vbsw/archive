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


/**
 * @author Vitali Baumtrok
 */
public class Analyzer implements Iterable<Result>
{

	public static final String DEFAULT_INDENT_TOKEN = "\t";
	public static final int SKIP_NONE = 0;
	public static final int SKIP_BLANKS_LEFT = 1;
	public static final int SKIP_BLANKS_RIGHT = 2;
	public static final int SKIP_TABS_LEFT = 4;
	public static final int SKIP_TABS_RIGHT = 8;
	public static final int SKIP_ALL = 15;

	protected final Result result;
	protected final String text;
	protected final boolean skipBlanksLeft;
	protected final boolean skipTabsLeft;
	protected final boolean skipBlanksRight;
	protected final boolean skipTabsRight;

	protected String indentToken = null;
	protected boolean invalidIndentDepth = false;
	protected boolean invalidNodeName = false;
	protected int indentDepthCurr = -1;
	protected int indentDepthPrev = -1;
	protected int lineNumber = -1;
	protected int iCurr = 0;

	public Analyzer ( final String text )
	{
		this ( text, Analyzer.SKIP_ALL );
	}

	public Analyzer ( final String text, final int skip )
	{
		this.text = text;
		this.result = new Result ( this );
		this.skipBlanksLeft = ( skip & SKIP_BLANKS_LEFT ) > 0;
		this.skipTabsLeft = ( skip & SKIP_TABS_LEFT ) > 0;
		this.skipBlanksRight = ( skip & SKIP_BLANKS_RIGHT ) > 0;
		this.skipTabsRight = ( skip & SKIP_TABS_RIGHT ) > 0;
	}

	@Override
	public java.util.Iterator<Result> iterator ( )
	{
		return new Iterator ( this );
	}

	public String getIndentToken ( )
	{
		if ( indentToken == null )
			processIndentToken ( );
		return indentToken;
	}

	protected void processNext ( )
	{
		invalidIndentDepth = false;
		invalidNodeName = false;

		if ( indentToken == null )
		{
			processIndentToken ( );
		}
		skipEmptyLines ( );
		processIndentDepth ( );
		processNode ( );
	}

	protected boolean hasNext ( )
	{
		return iCurr < text.length ( );
	}

	/**
	 * Parses the indent token. Increments the line number.
	 */
	protected void processIndentToken ( )
	{
		final String text = this.text;
		int iCurr = this.iCurr;

		// does not increment the line number, if nothing to parse 
		if ( iCurr < text.length ( ) )
		{
			final int iBegin = iCurr;
			int iEnd = iBegin;
			indentDepthPrev = indentDepthCurr;
			indentDepthCurr = 0;
			lineNumber++;

			while ( true )
			{
				if ( iCurr < text.length ( ) )
				{
					final char c = text.charAt ( iCurr );

					// Unix new line
					if ( c == '\n' )
					{
						iEnd = iCurr;
						iCurr++;
						break;
					}
					// Windows new line
					else if ( c == '\r' )
					{
						iEnd = iCurr;
						iCurr++;

						if ( iCurr < text.length ( ) && text.charAt ( iCurr ) == '\n' )
						{
							iCurr++;
						}
						break;
					}
					// valid indent token character
					else
					{
						iCurr++;
					}
				}
				else
				{
					iEnd = iCurr;
					break;
				}
			}

			if ( iBegin < iEnd )
			{
				indentToken = text.substring ( iBegin, iEnd );
			}
			else
			{
				indentToken = Analyzer.DEFAULT_INDENT_TOKEN;
			}
			this.iCurr = iCurr;
		}
	}

	/**
	 * Skips lines with spaces or tabs. Increments the line number.
	 */
	protected void skipEmptyLines ( )
	{
		final String text = this.text;
		int iCurr = this.iCurr;

		// does not increment the line number, if nothing to parse 
		if ( iCurr < text.length ( ) )
		{
			int iBegin = iCurr;

			while ( true )
			{
				if ( iCurr < text.length ( ) )
				{
					final char c = text.charAt ( iCurr );

					// Unix new line
					if ( c == '\n' )
					{
						lineNumber++;
						iCurr++;
						iBegin = iCurr;
					}
					// Windows new line
					else if ( c == '\r' )
					{
						lineNumber++;
						iCurr++;

						if ( iCurr < text.length ( ) && text.charAt ( iCurr ) == '\n' )
						{
							iCurr++;
						}
						iBegin = iCurr;
					}
					// white space
					else if ( c == ' ' || c == '\t' )
					{
						iCurr++;
					}
					// no white space; do not increment the line number
					else
					{
						iCurr = iBegin;
						break;
					}
				}
				else
				{
					lineNumber++;
					break;
				}
			}
			this.iCurr = iCurr;
		}
	}

	/**
	 * Parses the indent depth. Increments the line number.
	 */
	protected void processIndentDepth ( )
	{
		if ( this.iCurr < this.text.length ( ) )
		{
			final int indentTokenLength = indentToken.length ( );
			int indentDepth = 0;
			indentDepthPrev = indentDepthCurr;
			lineNumber++;

			while ( isIndentToken ( ) )
			{
				iCurr += indentTokenLength;
				indentDepth++;
			}
			indentDepthCurr = indentDepth;
		}

		invalidIndentDepth = indentDepthPrev + 1 < indentDepthCurr;
	}

	protected boolean isIndentToken ( )
	{
		final String text = this.text;
		final String indentToken = this.indentToken;
		int iCurr = this.iCurr;

		// ensure range
		if ( iCurr < text.length ( ) && indentToken.length ( ) <= text.length ( ) - iCurr )
		{
			for ( int i = 0; i < indentToken.length ( ); i++,iCurr++ )
			{
				if ( indentToken.charAt ( i ) != text.charAt ( iCurr ) )
				{
					return false;
				}
			}
			return true;
		}
		else
		{
			return false;
		}
	}

	/**
	 * Parses node name and node value. Indent token and indent depth is not parsed. Line
	 * number is not incremented.
	 */
	protected void processNode ( )
	{
		if ( this.iCurr < this.text.length ( ) )
		{
			char cCurr = this.text.charAt ( this.iCurr );
			invalidNodeName = cCurr == ' ' || cCurr == '\t';

			if ( invalidNodeName )
			{
				skipBlanksOrTabs ( true, true );
			}
			final String text = this.text;
			int iCurr = this.iCurr;
			int iBegin = iCurr;
			int iEnd = iBegin;

			// parse node name
			for ( ; iCurr < text.length ( ); iCurr++ )
			{
				cCurr = text.charAt ( iCurr );

				if ( cCurr == ' ' || cCurr == '\t' )
				{
					iEnd = iCurr;
					iCurr++; // skip one character before node value
					break;
				}
				else if ( cCurr == '\n' || cCurr == '\r' )
				{
					iEnd = iCurr;
					break;
				}
			}

			this.iCurr = iCurr;
			result.nameBegin = iBegin;
			result.nameEnd = iEnd;

			skipBlanksOrTabs ( skipBlanksLeft, skipTabsLeft );

			// parse node value

			iCurr = this.iCurr;
			iBegin = iCurr;
			iEnd = iBegin;

			// search line end
			while ( iCurr < text.length ( ) )
			{
				final char c = text.charAt ( iCurr );

				if ( c == '\n' || c == '\r' )
				{
					break;
				}
				iCurr++;
			}

			// search node value end
			if ( iEnd < iCurr )
			{
				iEnd = iCurr - 1;

				if ( skipBlanksRight )
				{
					if ( skipTabsRight )
					{
						while ( iBegin <= iEnd )
						{
							char c = text.charAt ( iEnd );

							if ( c != ' ' && c != '\t' )
							{
								break;
							}
							iEnd--;
						}
					}
					else
					{
						while ( iBegin <= iEnd )
						{
							if ( text.charAt ( iEnd ) != ' ' )
							{
								break;
							}
							iEnd--;
						}
					}
				}
				else if ( skipTabsRight )
				{
					while ( iBegin <= iEnd )
					{
						if ( text.charAt ( iEnd ) != '\t' )
						{
							break;
						}
						iEnd--;
					}
				}
				iEnd++;
			}
			result.valueBegin = iBegin;
			result.valueEnd = iEnd;

			skipLine ( );
		}
	}

	protected void skipBlanksOrTabs ( final boolean skipBlanks, final boolean skipTabs )
	{
		final String text = this.text;
		int iCurr = this.iCurr;

		if ( skipBlanks )
		{
			if ( skipTabs )
			{
				while ( iCurr < text.length ( ) )
				{
					final char c = text.charAt ( iCurr );

					if ( c != ' ' && c != '\t' )
					{
						break;
					}
					iCurr++;
				}
			}
			else
			{
				while ( iCurr < text.length ( ) )
				{
					if ( text.charAt ( iCurr ) != ' ' )
					{
						break;
					}
					iCurr++;
				}
			}
		}
		else if ( skipTabs )
		{
			while ( iCurr < text.length ( ) )
			{
				if ( text.charAt ( iCurr ) != '\t' )
				{
					break;
				}
				iCurr++;
			}
		}
		this.iCurr = iCurr;
	}

	/**
	 * Skips this line to beginning of a new line. Does not increment
	 * line number.
	 */
	protected void skipLine ( )
	{
		final String text = this.text;
		int iCurr = this.iCurr;

		while ( iCurr < text.length ( ) )
		{
			final char c = text.charAt ( iCurr );

			// Unix new line
			if ( c == '\n' )
			{
				iCurr++;
				break;
			}
			// Windows new line
			else if ( c == '\r' )
			{
				iCurr++;

				if ( iCurr < text.length ( ) && text.charAt ( iCurr ) == '\n' )
				{
					iCurr++;
				}
				break;
			}
			// other characters
			else
			{
				iCurr++;
			}
		}
		this.iCurr = iCurr;
	}

}
