/**
 * JSFW (a framework to create shoot 'em up games)
 * Copyright (C) 2016 Vitali Baumtrok
 *
 * This file is part of JSFW.
 *
 * JSFW is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * JSFW is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */


package com.github.vitalibaumtrok.jsfw;


import java.util.ArrayList;

import javafx.animation.Transition;


/**
 * @author Vitali Baumtrok
 */
public class JSFWFrameAnimation extends Transition
{
	protected final ArrayList <Double> durationValues = createArrayList ( );
	protected final JSFWFrameView frameView;
	protected double totalDuration;

	public JSFWFrameAnimation ( final JSFWFrameView frameView )
	{
		this.frameView = frameView;
		this.totalDuration = 0;
	}

	public void Clear ( )
	{
		totalDuration = 0;

		durationValues.clear ( );
	}

	public double getDuration ( final int index )
	{
		return durationValues.get ( index );
	}

	public void addDuration ( final double duration )
	{
		totalDuration += duration;

		durationValues.add ( duration );
	}

	protected ArrayList <Double> createArrayList ( )
	{
		final ArrayList <Double> list = new ArrayList<> ( 7 );

		return list;
	}

	@Override
	protected void interpolate ( final double frac )
	{
		if ( durationValues.size ( ) > 0 )
		{
			final int limit = durationValues.size ( );
			int index = 0;
			int nextIndex = index + 1;
			double duration = durationValues.get ( index );
			double fraction = duration / totalDuration;

			while ( nextIndex < limit && fraction <= frac )
			{
				index = nextIndex;
				nextIndex += 1;
				duration += durationValues.get ( index );
				fraction = duration / totalDuration;
			}
			frameView.setFrameVisible ( index );
		}
	}
}
