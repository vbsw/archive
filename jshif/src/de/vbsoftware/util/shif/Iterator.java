/**
 * JSHiF - a program to process documents in Simple Hierarchal Format
 * Copyright 2015 Vitali Baumtrok
 * 
 * This file is part of JSHiF.
 *
 * JSHiF is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * JSHiF is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */


package de.vbsoftware.util.shif;


/**
 * @author Vitali Baumtrok
 */
class Iterator implements java.util.Iterator<Result>
{

	private final Analyzer analyzer;

	public Iterator ( final Analyzer analyzer )
	{
		this.analyzer = analyzer;
	}

	@Override
	public boolean hasNext ( )
	{
		return analyzer.hasNext ( );
	}

	@Override
	public Result next ( )
	{
		analyzer.processNext ( );

		return analyzer.result;
	}

	@Override
	public void remove ( )
	{
		throw new UnsupportedOperationException ( );
	}

}
