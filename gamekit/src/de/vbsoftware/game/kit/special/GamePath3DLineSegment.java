
//          Copyright Vitali Baumtrok 2015.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)

package de.vbsoftware.game.kit.special;

import de.vbsoftware.game.kit.GamePathSegment;

/**
 * 
 * @author Vitali Baumtrok
 */
public class GamePath3DLineSegment extends GamePathSegment
{
	public final double	maxEntitySeconds;
	public final double	endX;
	public final double	endY;
	public final double	endZ;

	public GamePath3DLineSegment ( final double endX, final double endY, final double endZ, final double maxEntitySeconds )
	{
		this.endX = endX;
		this.endY = endY;
		this.endZ = endZ;
		this.maxEntitySeconds = maxEntitySeconds;
	}

	public void calculateLocation ( final GameLocation3D location, final double entitySeconds )
	{
		if ( entitySeconds < maxEntitySeconds )
		{
			final double percent = entitySeconds / maxEntitySeconds;
			location.x = location.startX + endX * percent;
			location.y = location.startY + endY * percent;
			location.z = location.startZ + endZ * percent;
		}
		else
		{
			location.x = location.startX + endX;
			location.y = location.startY + endY;
			location.z = location.startZ + endZ;
			location.isEnd = true;
		}
	}
}
