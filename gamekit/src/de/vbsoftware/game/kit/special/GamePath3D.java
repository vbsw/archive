
//          Copyright Vitali Baumtrok 2015.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)

package de.vbsoftware.game.kit.special;

import de.vbsoftware.game.kit.GamePath;

/**
 * 
 * @author Vitali Baumtrok
 */
public class GamePath3D extends GamePath
{

	protected final GamePath3DSegment[]	pathSegments;
	protected final long				repeatNumber;

	public GamePath3D ( final long repeatNumber, final GamePath3DSegment... pathSegments )
	{
		this.pathSegments = pathSegments;
		this.repeatNumber = repeatNumber;
	}

	public void calculateLocation ( final GameLocation3D location, final double entitySeconds )
	{
		if ( !location.isEnd )
		{
			final double startX = location.startX;
			final double startY = location.startY;
			final double startZ = location.startZ;
			final int lastIndex = pathSegments.length - 1;
			long repeatNumber = this.repeatNumber;
			double entitySecondsLeft = entitySeconds;

			// loop over the path
			outer: while ( entitySeconds > 0 && repeatNumber != 0 )
			{
				for ( int i = 0; i < pathSegments.length; i++ )
				{
					final GamePath3DSegment pathSegment = pathSegments[i];
					pathSegment.calculateLocation ( location, entitySecondsLeft );
					entitySecondsLeft -= pathSegment.maxEntitySeconds;

					// more segments left (reset isEnd)
					if ( i < lastIndex )
					{
						location.isEnd = false;

						// more time left (continue segment iteration)
						if ( entitySeconds > 0 )
						{
							location.startX = location.x;
							location.startY = location.y;
							location.startZ = location.z;
						}
						// no more time
						else
						{
							location.startX = startX;
							location.startY = startY;
							location.startZ = startZ;
							break;
						}
					}
					// no more segments
					else
					{
						if ( entitySeconds > 0 )
						{
							location.startX = startX;
							location.startY = startY;
							location.startZ = startZ;
							repeatNumber--;

							// continue path iteration
							if ( repeatNumber != 0 )
							{
								location.isEnd = false;
								location.x = startX;
								location.y = startY;
								location.z = startZ;
								break;
							}
							// keep isEnd, which can be true or false
							else
							{
								break outer;
							}
						}
						else
						{
							// keep isEnd, which can be true or false
							location.startX = startX;
							location.startY = startY;
							location.startZ = startZ;
							break outer;
						}
					}
				}
			}
		}
	}

	@Override
	public double getMaxEntitySeconds ()
	{
		if ( repeatNumber <= 0 )
		{
			return -1;
		}
		else
		{
			double value = 0;

			for ( GamePath3DSegment pathSegment: pathSegments )
			{
				value += pathSegment.maxEntitySeconds;
			}

			return value * repeatNumber;
		}
	}
}
