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
 * aint with com.gitlab.vbit.util.  If not,
 * see <http://www.gnu.org/licenses/>.
 */


package com.gitlab.vbit.util;


/**
 * @author Vitali Baumtrok
 */
public class IntStack
{
	public static final int INITIAL_CAPACITY = 10;
	public static final int DEFAULT_VALUE = Integer.MAX_VALUE;

	private int[] values;
	private int start;
	private int length;

	public IntStack ( final int capacity )
	{
		values = new int[capacity];

		reset();
	}

	public IntStack ( )
	{
		this(INITIAL_CAPACITY);
	}

	public void reset ( )
	{
		start = 0;
		length = 0;
	}

	public int push ( final int value )
	{
		if ( length==values.length )
		{
			final int[] valuesNew = new int[values.length*2];

			System.arraycopy(values,0,valuesNew,0,length);

			values = valuesNew;
		}
		values[length] = value;
		length += 1;

		return value;
	}

	public int popTop ( )
	{
		if ( length>0&&length>start )
		{
			return values[--length];
		}

		return DEFAULT_VALUE;
	}

	public int popBot ( )
	{
		if ( start<values.length&&start<length )
		{
			return values[start--];
		}

		return DEFAULT_VALUE;
	}

	public int top ( )
	{
		if ( length>0&&length>start )
		{
			return values[length-1];
		}

		return DEFAULT_VALUE;
	}

	public int bot ( )
	{
		if ( start<values.length&&start<length )
		{
			return values[start];
		}

		return DEFAULT_VALUE;
	}

	public boolean isEmpty ( )
	{
		return start==length;
	}
}
