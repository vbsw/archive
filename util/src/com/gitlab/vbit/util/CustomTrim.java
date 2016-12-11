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


/**
 * @author Vitali Baumtrok
 */
public class CustomTrim extends Trim
{
	private final char[] chars;

	public CustomTrim ( final char... chars )
	{
		this.chars = chars;
	}

	@Override
	public int getLeftIndex ( final String string )
	{
		Outer:
		for ( int start = 0; start<string.length(); start += 1 )
		{
			for ( int i = 0; i<chars.length; i += 1 )
			{
				if ( string.charAt(start)==chars[i] )
				{
					continue Outer;
				}
			}

			return start;
		}

		return 0;
	}

	@Override
	public int getRightIndex ( final String string )
	{
		Outer:
		for ( int end = string.length(); end>0; end -= 1 )
		{
			for ( int i = 0; i<chars.length; i += 1 )
			{
				if ( string.charAt(end-1)==chars[i] )
				{
					continue Outer;
				}
			}

			return end;
		}

		return 0;
	}

	@Override
	public int getLeftIndex ( final char[] array )
	{
		Outer:
		for ( int start = 0; start<array.length; start += 1 )
		{
			for ( int i = 0; i<chars.length; i += 1 )
			{
				if ( array[start]==chars[i] )
				{
					continue Outer;
				}
			}

			return start;
		}

		return 0;
	}

	@Override
	public int getRightIndex ( final char[] array )
	{
		Outer:
		for ( int end = array.length; end>0; end -= 1 )
		{
			for ( int i = 0; i<chars.length; i += 1 )
			{
				if ( array[end-1]==chars[i] )
				{
					continue Outer;
				}
			}

			return end;
		}

		return 0;
	}

	@Override
	public int getLeftIndex ( final StringBuffer buffer )
	{
		Outer:
		for ( int start = 0; start<buffer.length(); start += 1 )
		{
			for ( int i = 0; i<chars.length; i += 1 )
			{
				if ( buffer.charAt(start)==chars[i] )
				{
					continue Outer;
				}
			}

			return start;
		}

		return 0;
	}

	@Override
	public int getRightIndex ( final StringBuffer buffer )
	{
		Outer:
		for ( int end = buffer.length(); end>0; end -= 1 )
		{
			for ( int i = 0; i<chars.length; i += 1 )
			{
				if ( buffer.charAt(end-1)==chars[i] )
				{
					continue Outer;
				}
			}

			return end;
		}

		return 0;
	}

	@Override
	public int getLeftIndex ( final StringBuilder builder )
	{
		Outer:
		for ( int start = 0; start<builder.length(); start += 1 )
		{
			for ( int i = 0; i<chars.length; i += 1 )
			{
				if ( builder.charAt(start)==chars[i] )
				{
					continue Outer;
				}
			}

			return start;
		}

		return 0;
	}

	@Override
	public int getRightIndex ( final StringBuilder builder )
	{
		Outer:
		for ( int end = builder.length(); end>0; end -= 1 )
		{
			for ( int i = 0; i<chars.length; i += 1 )
			{
				if ( builder.charAt(end-1)==chars[i] )
				{
					continue Outer;
				}
			}

			return end;
		}

		return 0;
	}
}
