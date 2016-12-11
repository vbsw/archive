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


import javafx.collections.ObservableList;
import javafx.geometry.Insets;
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
public class JSFWScreenMenu
{
	protected final JSFWScreens screens;
	protected final JSFWLayeredPane layeredPane;
	protected final double fixedWidth;
	protected final double fixedHeight;
	protected final Background background;
	protected final Text titleText;
	protected final Text subtitleText;
	protected final Text versionText;

	public JSFWScreenMenu ( final JSFWScreens screens, final JSFWLayeredPane layeredPane )
	{
		this.screens = screens;
		this.layeredPane = layeredPane;
		this.fixedWidth = layeredPane.getWidth ( );
		this.fixedHeight = layeredPane.getHeight ( );
		this.background = createBackground ( );
		this.titleText = createTitleText ( "Shoot 'em Up Game" );
		this.subtitleText = createSubtitleText ( "Created with JSFW" );
		this.versionText = createVerstionText ( "Version 0.1.0" );
	}

	public void fill ( )
	{
		filllayer0 ( );
		fillLayer1 ( );
	}

	public void empty ( )
	{
		emptylayer0 ( );
		emptyLayer1 ( );
	}

	public void moveSelectorUp ( )
	{
		// TODO Auto-generated method stub
	}

	public void moveSelectorDown ( )
	{
		// TODO Auto-generated method stub
	}

	public void select ( )
	{
		// TODO Auto-generated method stub
	}

	protected void filllayer0 ( )
	{
		final JSFWAnchorPane layer = layeredPane.getLayer ( 0 );

		layer.setBackground ( background );
		layer.setScaling ( false );
	}

	protected void fillLayer1 ( )
	{
		final JSFWAnchorPane layer = layeredPane.getLayer ( 1 );
		final ObservableList <Node> children = layer.getChildren ( );

		children.add ( titleText );
		children.add ( subtitleText );
		children.add ( versionText );
		layer.setScalingSize ( fixedWidth, fixedHeight );
		layer.setScaling ( true );
	}

	protected void emptylayer0 ( )
	{
		final JSFWAnchorPane layer = layeredPane.getLayer ( 0 );

		layer.setBackground ( null );
	}

	protected void emptyLayer1 ( )
	{
		final JSFWAnchorPane layer = layeredPane.getLayer ( 1 );
		final ObservableList <Node> children = layer.getChildren ( );

		children.remove ( titleText );
		children.remove ( subtitleText );
		children.remove ( versionText );
	}

	protected Background createBackground ( )
	{
		final Color color = Color.BLACK;
		final BackgroundFill bgFill = new BackgroundFill ( color, CornerRadii.EMPTY, Insets.EMPTY );
		final Background bg = new Background ( bgFill );

		return bg;
	}

	protected Text createTitleText ( final String textString )
	{
		final Font font = new Font ( "Arial Bold", 32 );
		final Text text = JSFWUtil.createText ( textString, font );
		final double y = fixedHeight * 0.25;

		JSFWUtil.centerHorizontaly ( text, fixedWidth, y );

		return text;
	}

	protected Text createSubtitleText ( final String textString )
	{
		final Font font = new Font ( "Arial", 16 );
		final Text text = JSFWUtil.createText ( textString, font );
		final double y = fixedHeight * 0.33;

		JSFWUtil.centerHorizontaly ( text, fixedWidth, y );

		return text;
	}

	protected Text createVerstionText ( final String textString )
	{
		final Font font = new Font ( "Arial Bold", 16 );
		final Text text = JSFWUtil.createText ( textString, font );
		final double y = fixedHeight * 0.7;

		JSFWUtil.centerHorizontaly ( text, fixedWidth, y );

		return text;
	}
}
