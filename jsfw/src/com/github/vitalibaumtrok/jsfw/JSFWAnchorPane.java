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


import javafx.beans.property.ReadOnlyObjectProperty;
import javafx.scene.Node;
import javafx.scene.Scene;
import javafx.scene.layout.AnchorPane;
import javafx.scene.layout.Region;
import javafx.scene.shape.Rectangle;


/**
 * @author Vitali Baumtrok
 */
public class JSFWAnchorPane extends AnchorPane
{
	protected static final double CORRECTION_FACTOR = 1.003;

	protected double scalingWidth = 0;
	protected double scalingHeight = 0;
	protected boolean scalingEnabled = false;

	public JSFWAnchorPane ( )
	{
		final JSFWSizeChangeListener sizeChangeListener = new JSFWSizeChangeListener ( this );
		final JSFWSceneChangeListener sceneChangeListener = new JSFWSceneChangeListener ( sizeChangeListener );
		final ReadOnlyObjectProperty <Scene> sceneProperty = sceneProperty ( );

		sceneProperty.addListener ( sceneChangeListener );
	}

	public void setScalingSize ( final double width, final double height )
	{
		if ( scalingEnabled )
		{
			if ( width > 0 && height > 0 )
			{
				if ( scalingWidth != width || scalingHeight != height )
				{
					setSizeFixed ( width, height );
				}
			}
			else if ( scalingWidth > 0 && scalingHeight > 0 )
			{
				setSizeVariable ( );
			}
		}
		scalingWidth = width;
		scalingHeight = height;
	}

	public void setScaling ( final boolean enabled )
	{
		if ( scalingEnabled == false )
		{
			if ( enabled )
			{
				if ( scalingWidth > 0 && scalingHeight > 0 )
				{
					setSizeFixed ( scalingWidth, scalingHeight );
				}
			}
		}
		else
		{
			if ( enabled == false )
			{
				if ( scalingWidth > 0 && scalingHeight > 0 )
				{
					setSizeVariable ( );
				}
			}
		}
		scalingEnabled = enabled;
	}

	protected void setSizeFixed ( final double width, final double height )
	{
		setWidth ( width );
		setHeight ( height );
		setMinSize ( width, height );
		setMaxSize ( width, height );
		setClipSize ( width, height );
	}

	protected void setClipSize ( double width, double height )
	{
		final Node clip = getClip ( );

		if ( clip == null )
		{
			final Rectangle rect = new Rectangle ( width, height );

			setClip ( rect );
		}
		else if ( clip instanceof Rectangle )
		{
			final Rectangle rect = (Rectangle) clip;

			rect.setWidth ( width );
			rect.setHeight ( height );
		}
	}

	protected void setSizeVariable ( )
	{
		setClip ( null );
		setMinSize ( 0, 0 );
		setMaxSize ( Double.MAX_VALUE, Double.MAX_VALUE );
	}

	protected void scaleToFitParent ( )
	{
		Node parent = getParent ( );
		Region region = null;

		while ( parent != null && region == null )
		{
			if ( parent instanceof Region )
			{
				region = (Region) parent;
			}
			else
			{
				parent = parent.getParent ( );
			}
		}

		if ( region != null )
		{
			final double width = region.getWidth ( );
			final double height = region.getHeight ( );

			scaleToFit ( width, height );
		}
		else
		{
			scaleToFitScene ( );
		}
	}

	protected void scaleToFitScene ( )
	{
		final Scene scene = getScene ( );

		if ( scene != null )
		{
			final double width = scene.getWidth ( );
			final double height = scene.getHeight ( );

			scaleToFit ( width, height );
		}
	}

	protected void scaleToFit ( final double width, final double height )
	{
		if ( scalingEnabled && scalingWidth > 0 && scalingHeight > 0 )
		{
			final double paneWidth = getWidth ( );
			final double paneHeight = getHeight ( );

			if ( paneWidth > 0 && paneHeight > 0 )
			{
				final double widthRatio = width / paneWidth;
				final double heightRatio = height / paneHeight;
				final double scaleRatio;

				scaleRatio = ( widthRatio < heightRatio ? widthRatio : heightRatio ) * CORRECTION_FACTOR;

				setScaleX ( scaleRatio );
				setScaleY ( scaleRatio );
			}
		}
	}
}
