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


import com.github.vitalibaumtrok.jsfw.JSFWLayeredPane;
import com.github.vitalibaumtrok.jsfw.JSFWScreenMenu;
import com.github.vitalibaumtrok.jsfw.JSFWScreens;


/**
 * @author Vitali Baumtrok
 */
public class Screens extends JSFWScreens
{
	public Screens ( final JSFWLayeredPane layeredPane )
	{
		super ( layeredPane );
	}

	@Override
	protected JSFWScreenMenu createScreenMenu ( final JSFWLayeredPane layeredPane )
	{
		return new ScreenMenu ( this, layeredPane );
	}
}
