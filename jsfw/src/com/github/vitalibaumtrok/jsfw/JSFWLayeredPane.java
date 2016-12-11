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
import javafx.scene.Node;
import javafx.scene.layout.StackPane;


/**
 * @author Vitali Baumtrok
 */
public class JSFWLayeredPane extends StackPane
{
	public JSFWLayeredPane ( )
	{
		this ( 2 );
	}

	public JSFWLayeredPane ( final int layersNumber )
	{
		ensureLayers ( layersNumber );
	}

	public void ensureLayers ( final int layersNumber )
	{
		final ObservableList <Node> nodes = getChildren ( );

		while ( nodes.size ( ) < layersNumber )
		{
			final JSFWAnchorPane pane = new JSFWAnchorPane ( );

			nodes.add ( pane );
		}
	}

	public JSFWAnchorPane getLayer ( final int layerIndex )
	{
		final ObservableList <Node> nodes = getChildren ( );
		final Node node = nodes.get ( layerIndex );

		return (JSFWAnchorPane) node;
	}
}
