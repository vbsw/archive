
//          Copyright Vitali Baumtrok 2015.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)

package de.vbsoftware.game.kit.special;

/**
 * 
 * @author Vitali Baumtrok
 */
public abstract class GamePath3DSegment
{
	public final double	maxEntitySeconds;

	public GamePath3DSegment ( final double maxEntitySeconds )
	{
		this.maxEntitySeconds = maxEntitySeconds;
	}

	public abstract void calculateLocation ( GameLocation2D location, double entitySeconds );
}
