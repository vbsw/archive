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

import javafx.geometry.Rectangle2D;
import javafx.scene.image.Image;
import javafx.scene.image.ImageView;


/**
 * @author Vitali Baumtrok
 */
public class JSFWFrameView extends ImageView
{
	protected final ArrayList <Rectangle2D> viewPorts = createArrayList ( );

	public JSFWFrameView ( final Image image )
	{
		super ( image );
	}

	public JSFWFrameView ( )
	{}

	public void clear ( )
	{
		viewPorts.clear ( );
	}

	public Rectangle2D getFrame ( final int index )
	{
		return viewPorts.get ( index );
	}

	public void addFrame ( final double x, final double y, final double width, final double height )
	{
		final Rectangle2D rect = new Rectangle2D ( x, y, width, height );

		viewPorts.add ( rect );
	}

	public void setFrameVisible ( final int index )
	{
		final Rectangle2D rect = viewPorts.get ( index );

		setViewport ( rect );
	}

	protected ArrayList <Rectangle2D> createArrayList ( )
	{
		final ArrayList <Rectangle2D> list = new ArrayList<> ( 7 );

		return list;
	}
}
