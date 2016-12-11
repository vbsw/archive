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


import java.nio.ByteBuffer;
import java.nio.ByteOrder;


/**
 * @author Vitali Baumtrok
 */
public class SBiFAnlyzer
{
	private ByteBuffer buffer = null;

	public void parse ( )
	{}

	public ByteBuffer getBuffer ( )
	{
		return buffer;
	}

	public void resetBuffer ( final int minCapacity )
	{
		if ( buffer==null||buffer.capacity()<minCapacity )
		{
			buffer = ByteBuffer.allocate(minCapacity);

			buffer.order(ByteOrder.LITTLE_ENDIAN);
		}
		else
		{
			buffer.rewind();
			buffer.limit(buffer.capacity());
		}
	}

	public boolean isBufferWhitespaceOnly ( )
	{
		final int position = buffer.position();

		for ( int i = 0; i<buffer.limit(); i += 1 )
		{
			final byte b = buffer.get();

			if ( b>0x20||b<0 )
			{
				buffer.position(position);

				return true;
			}
		}
		buffer.position(position);

		return false;
	}
}
