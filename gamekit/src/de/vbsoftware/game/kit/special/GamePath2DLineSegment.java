
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
public class GamePath2DLineSegment extends GamePathSegment
{
	public final double	maxEntitySeconds;
	public final double	endX;
	public final double	endY;

	public GamePath2DLineSegment ( final double endX, final double endY, final double maxEntitySeconds )
	{
		this.endX = endX;
		this.endY = endY;
		this.maxEntitySeconds = maxEntitySeconds;
	}

	public void calculateLocation ( final GameLocation2D location, final double entitySeconds )
	{
		if ( entitySeconds < maxEntitySeconds )
		{
			final double percent = entitySeconds / maxEntitySeconds;
			location.x = location.startX + endX * percent;
			location.y = location.startY + endY * percent;
		}
		else
		{
			location.x = location.startX + endX;
			location.y = location.startY + endY;
			location.isEnd = true;
		}
	}
}
