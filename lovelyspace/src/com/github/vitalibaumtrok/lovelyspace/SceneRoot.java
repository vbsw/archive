/**
 * Lovely Space (a shoot 'em up game)
 * Copyright (C) 2016 Vitali Baumtrok
 *
 * This file is part of Lovely Space.
 *
 * Lovely Space is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Lovely Space is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */


package com.github.vitalibaumtrok.lovelyspace;


import com.github.vitalibaumtrok.jsfw.JSFWAnchorPane;
import com.github.vitalibaumtrok.jsfw.JSFWLayeredPane;

import javafx.collections.ObservableList;
import javafx.geometry.Bounds;
import javafx.geometry.Insets;
import javafx.scene.CacheHint;
import javafx.scene.Node;
import javafx.scene.layout.Background;
import javafx.scene.layout.BackgroundFill;
import javafx.scene.layout.CornerRadii;
import javafx.scene.paint.Color;
import javafx.scene.text.Font;
import javafx.scene.text.Text;


/**
 * @author Vitali Baumtrok
 */
public class SceneRoot extends JSFWLayeredPane
{
	public SceneRoot ( )
	{
		super ( 2 );

		initBackground ( );
		initInfo ( );
	}

	private void initBackground ( )
	{
		final JSFWAnchorPane layer = getLayer ( 0 );
		final Color color = Color.BLACK;
		final BackgroundFill bgFill = new BackgroundFill ( color, CornerRadii.EMPTY, Insets.EMPTY );
		final Background bg = new Background ( bgFill );

		layer.setBackground ( bg );
	}

	private void initInfo ( )
	{
		final JSFWAnchorPane layer = getLayer ( 1 );
		final ObservableList <Node> children = layer.getChildren ( );
		final Text titleText = createText ( Window.TITLE, Fonts.ARIAL_BOLD_32, 100 );
		final Text infoText = createText ( "(nothing implmented yet)", Fonts.ARIAL_12, 140 );
		final Text versionText = createText ( "Version " + Version.STRING, Fonts.ARIAL_BOLD_16, 250 );

		children.add ( titleText );
		children.add ( infoText );
		children.add ( versionText );
		layer.setScalingSize ( Window.WIDTH, Window.HEIGHT );
		layer.setScaling ( true );
	}

	private Text createText ( final String textString, final Font font, final double y )
	{
		final Text text = new Text ( textString );

		text.setFill ( Color.WHITE );
		text.setFont ( font );
		text.setCacheHint ( CacheHint.QUALITY );
		text.setCache ( true );

		final Bounds textBounds = text.getLayoutBounds ( );
		final double textWidth = textBounds.getWidth ( );
		final double x = ( Window.WIDTH - textWidth ) / 2;

		text.setTranslateX ( x );
		text.setTranslateY ( y );

		return text;
	}
}
