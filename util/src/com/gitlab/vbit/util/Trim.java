/**
 * com.gitlab.vbit.util (helper classes)
 * CopyRight 2016 Vitali Baumtrok
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


/**
 * @author Vitali Baumtrok
 */
public class Trim
{
	public int getLeftIndex ( final String string )
	{
		for ( int start = 0; start<string.length(); start += 1 )
		{
			if ( string.charAt(start)>0x20 )
			{
				return start;
			}
		}

		return 0;
	}

	public int getRightIndex ( final String string )
	{
		for ( int end = string.length(); end>0; end -= 1 )
		{
			if ( string.charAt(end-1)>0x20 )
			{
				return end;
			}
		}

		return 0;
	}

	public int getLeftIndex ( final char[] array )
	{
		for ( int start = 0; start<array.length; start += 1 )
		{
			if ( array[start]>0x20 )
			{
				return start;
			}
		}

		return 0;
	}

	public int getRightIndex ( final char[] array )
	{
		for ( int end = array.length; end>0; end -= 1 )
		{
			if ( array[end-1]>0x20 )
			{
				return end;
			}
		}

		return 0;
	}

	public int getLeftIndex ( final StringBuffer buffer )
	{
		for ( int start = 0; start<buffer.length(); start += 1 )
		{
			if ( buffer.charAt(start)>0x20 )
			{
				return start;
			}
		}

		return 0;
	}

	public int getRightIndex ( final StringBuffer buffer )
	{
		for ( int end = buffer.length(); end>0; end -= 1 )
		{
			if ( buffer.charAt(end-1)>0x20 )
			{
				return end;
			}
		}

		return 0;
	}

	public int getLeftIndex ( final StringBuilder builder )
	{
		for ( int start = 0; start<builder.length(); start += 1 )
		{
			if ( builder.charAt(start)>0x20 )
			{
				return start;
			}
		}

		return 0;
	}

	public int getRightIndex ( final StringBuilder builder )
	{
		for ( int end = builder.length(); end>0; end -= 1 )
		{
			if ( builder.charAt(end-1)>0x20 )
			{
				return end;
			}
		}

		return 0;
	}

	public String left ( final String string )
	{
		final int start = getLeftIndex(string);

		if ( start>0 )
		{
			return string.substring(start);
		}
		else
		{
			return string;
		}
	}

	public String right ( final String string )
	{
		final int end = getLeftIndex(string);

		if ( end<string.length() )
		{
			return string.substring(0,end);
		}
		else
		{
			return string;
		}
	}

	public String both ( final String string )
	{
		final int start = getLeftIndex(string);
		final int end = getRightIndex(string);

		if ( start>0||end<string.length() )
		{
			return string.substring(start,end);
		}
		else
		{
			return string;
		}
	}

	public char[] left ( final char[] array, final int start )
	{
		if ( start>0 )
		{
			System.arraycopy(array,start,array,0,array.length-start);
		}

		return array;
	}

	public StringBuffer left ( final StringBuffer buffer )
	{
		synchronized ( buffer )
		{
			final int start = getLeftIndex(buffer);

			if ( start>0 )
			{
				buffer.delete(start,buffer.length());
			}

			return buffer;
		}
	}

	public StringBuffer right ( final StringBuffer buffer )
	{
		synchronized ( buffer )
		{
			final int end = getLeftIndex(buffer);

			if ( end<buffer.length() )
			{
				buffer.delete(0,end);
			}

			return buffer;
		}
	}

	public StringBuffer both ( final StringBuffer buffer )
	{
		synchronized ( buffer )
		{
			final int start = getLeftIndex(buffer);
			final int end = getRightIndex(buffer);

			if ( end<buffer.length() )
			{
				buffer.setLength(end);
			}

			if ( start>0 )
			{
				buffer.delete(0,start);
			}

			return buffer;
		}
	}

	public StringBuilder left ( final StringBuilder builder )
	{
		final int start = getLeftIndex(builder);

		if ( start>0 )
		{
			builder.delete(start,builder.length());
		}

		return builder;
	}

	public StringBuilder right ( final StringBuilder builder )
	{
		final int end = getLeftIndex(builder);

		if ( end<builder.length() )
		{
			builder.delete(0,end);
		}

		return builder;
	}

	public StringBuilder both ( final StringBuilder builder )
	{
		final int start = getLeftIndex(builder);
		final int end = getRightIndex(builder);

		if ( end<builder.length() )
		{
			builder.setLength(end);
		}

		if ( start>0 )
		{
			builder.delete(0,start);
		}

		return builder;
	}

	public String leftToString ( final char[] array )
	{
		final int start = getLeftIndex(array);

		if ( start>0 )
		{
			return new String(array,start,array.length-start);
		}
		else
		{
			return new String(array);
		}
	}

	public String rightToString ( final char[] array )
	{
		final int end = getRightIndex(array);

		if ( end<array.length )
		{
			return new String(array,0,end);
		}
		else
		{
			return new String(array);
		}
	}

	public String leftToString ( final StringBuffer buffer )
	{
		synchronized ( buffer )
		{
			final int start = getLeftIndex(buffer);

			if ( start>0 )
			{
				return buffer.substring(start,buffer.length());
			}
			else
			{
				return buffer.toString();
			}
		}
	}

	public String rightToString ( final StringBuffer buffer )
	{
		synchronized ( buffer )
		{
			final int end = getLeftIndex(buffer);

			if ( end<buffer.length() )
			{
				return buffer.substring(0,end);
			}
			else
			{
				return buffer.toString();
			}
		}
	}

	public String bothToString ( final StringBuffer buffer )
	{
		synchronized ( buffer )
		{
			final int start = getLeftIndex(buffer);
			final int end = getRightIndex(buffer);

			if ( start>0||end<buffer.length() )
			{
				return buffer.substring(start,end);
			}
			else
			{
				return buffer.toString();
			}
		}
	}

	public String leftToString ( final StringBuilder builder )
	{
		final int start = getLeftIndex(builder);

		if ( start>0 )
		{
			return builder.substring(start,builder.length());
		}
		else
		{
			return builder.toString();
		}
	}

	public String rightToString ( final StringBuilder builder )
	{
		final int end = getLeftIndex(builder);

		if ( end<builder.length() )
		{
			return builder.substring(0,end);
		}
		else
		{
			return builder.toString();
		}
	}

	public String bothToString ( final StringBuilder builder )
	{
		final int start = getLeftIndex(builder);
		final int end = getRightIndex(builder);

		if ( start>0||end<builder.length() )
		{
			return builder.substring(start,end);
		}
		else
		{
			return builder.toString();
		}
	}
}
