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


/**
 * @author Vitali Baumtrok
 */
public class JSFWScreens
{
	protected final JSFWScreenMenu menu;
	protected Selection selection;

	public JSFWScreens ( final JSFWLayeredPane layeredPane )
	{
		menu = createScreenMenu ( layeredPane );
		selection = Selection.MENU;
	}

	public JSFWScreenMenu getMenu ( )
	{
		return menu;
	}

	protected JSFWScreenMenu createScreenMenu ( final JSFWLayeredPane layeredPane )
	{
		return new JSFWScreenMenu ( this, layeredPane );
	}

	public boolean isMenu ( )
	{
		return selection == Selection.MENU;
	}

	protected enum Selection
	{
		MENU;
	}
}
