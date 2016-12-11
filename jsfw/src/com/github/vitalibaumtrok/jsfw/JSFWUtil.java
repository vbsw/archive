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


import javafx.geometry.Bounds;
import javafx.scene.CacheHint;
import javafx.scene.Node;
import javafx.scene.paint.Color;
import javafx.scene.text.Font;
import javafx.scene.text.Text;


/**
 * @author Vitali Baumtrok
 */
public class JSFWUtil
{
	public static Text createText ( final String textString, final Font font )
	{
		final Text text = new Text ( textString );

		text.setFill ( Color.WHITE );
		text.setFont ( font );
		text.setCacheHint ( CacheHint.QUALITY );
		text.setCache ( true );

		return text;
	}

	public static void centerHorizontaly ( final Node node, final double areaWidth, final double y )
	{
		final Bounds nodeBounds = node.getLayoutBounds ( );
		final double nodeWidth = nodeBounds.getWidth ( );
		final double x = ( areaWidth - nodeWidth ) / 2;

		node.setTranslateX ( x );
		node.setTranslateY ( y );
	}
}
