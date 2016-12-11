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


import javafx.beans.property.ReadOnlyDoubleProperty;
import javafx.beans.value.ChangeListener;
import javafx.beans.value.ObservableValue;
import javafx.scene.Scene;


/**
 * @author Vitali Baumtrok
 */
public class JSFWSceneChangeListener implements ChangeListener <Scene>
{
	protected final ChangeListener <? super Number> sizeChangeListener;

	public JSFWSceneChangeListener ( final ChangeListener <? super Number> sizeChangeListener )
	{
		this.sizeChangeListener = sizeChangeListener;
	}

	@Override
	public void changed ( final ObservableValue <? extends Scene> observable, final Scene oldValue, final Scene newValue )
	{
		if ( oldValue != null )
		{
			final ReadOnlyDoubleProperty widthProp = oldValue.widthProperty ( );
			final ReadOnlyDoubleProperty heightProp = oldValue.heightProperty ( );

			widthProp.removeListener ( sizeChangeListener );
			heightProp.removeListener ( sizeChangeListener );
		}

		if ( newValue != null )
		{
			final ReadOnlyDoubleProperty widthProp = newValue.widthProperty ( );
			final ReadOnlyDoubleProperty heightProp = newValue.heightProperty ( );

			widthProp.addListener ( sizeChangeListener );
			heightProp.addListener ( sizeChangeListener );
		}
	}
}
