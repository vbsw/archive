/**
 * JSHiF - a program to process documents in Simple Hierarchal Format
 * Copyright 2015, Vitali Baumtrok
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


import java.io.IOException;
import java.io.Writer;
import java.util.ArrayList;


/**
 * @author Vitali Baumtrok
 */
public class Node
{

	public final ArrayList<Node> subnodes = new ArrayList<> ( );

	public String name;
	public String value;
	public Node parent;

	public Node ( )
	{
		this ( "", "", null );
	}

	public Node ( final String name, final String value )
	{
		this ( name, value, null );
	}

	public Node ( final String name, final String value, final Node parent )
	{
		this.name = name;
		this.value = value;
		this.parent = parent;
	}

	/**
	 * Searches for the first subnode with the given name and returns it.
	 * 
	 * @param nodeName The name which should the subnode have.
	 * 
	 * @return Subnode with the given name.
	 */
	public Node getNodeByName ( final String nodeName )
	{
		for ( final Node node: subnodes )
		{
			if ( node != null && node.name != null && node.name.equals ( nodeName ) )
			{
				return node;
			}
		}
		return null;
	}

	/**
	 * Searches for the first node in all children nodes with the given name
	 * and returns it.
	 * 
	 * @param nodeName The name which should the subnode have.
	 * 
	 * @return Subnode with the given name.
	 */
	public Node getNodeByNameRecursive ( final String nodeName )
	{
		for ( final Node node: subnodes )
		{
			if ( node != null && node.name != null && node.name.equals ( nodeName ) )
			{
				return node;
			}
		}
		for ( final Node node: subnodes )
		{
			if ( node != null )
			{
				final Node n = node.getNodeByNameRecursive ( nodeName );

				if ( n != null )
				{
					return node;
				}
			}
		}
		return null;
	}

	public void addNode ( final Node node )
	{
		node.parent = this;
		subnodes.add ( node );
	}

	public void addNode ( final String name, final String value )
	{
		addNode ( new Node ( name, value, this ) );
	}

	/**
	 * The same as {@link append}, but writes the indent token in first line.
	 * 
	 * @param writer
	 * @throws IOException
	 */
	public void write ( final Writer writer ) throws IOException
	{
		writePrivate ( writer, Analyzer.DEFAULT_INDENT_TOKEN, 0, System.getProperty ( "line.separator" ) );
	}

	/**
	 * The same as {@link append}, but writes the indent token in first line.
	 * 
	 * @param writer
	 * @param indentToken
	 * @throws IOException
	 */
	public void write ( final Writer writer, final String indentToken ) throws IOException
	{
		writePrivate ( writer, indentToken, 0, System.getProperty ( "line.separator" ) );
	}

	/**
	 * The same as {@link append}, but writes the indent token in first line.
	 * 
	 * @param writer
	 * @param indentToken
	 * @param indentDepth
	 * @throws IOException
	 */
	public void write ( final Writer writer, final String indentToken, final int indentDepth ) throws IOException
	{
		writePrivate ( writer, indentToken, indentDepth, System.getProperty ( "line.separator" ) );
	}

	/**
	 * Writes the content.
	 * 
	 * @param writer
	 * @throws IOException
	 */
	public void append ( final Writer writer ) throws IOException
	{
		appendPrivate ( writer, Analyzer.DEFAULT_INDENT_TOKEN, 0, System.getProperty ( "line.separator" ) );
	}

	/**
	 * Writes the content.
	 * 
	 * @param writer
	 * @param indentToken
	 * @throws IOException
	 */
	public void append ( final Writer writer, final String indentToken ) throws IOException
	{
		appendPrivate ( writer, indentToken, 0, System.getProperty ( "line.separator" ) );
	}

	/**
	 * Writes the content.
	 * 
	 * @param writer
	 * @param indentToken
	 * @param indentDepth
	 * @throws IOException
	 */
	public void append ( final Writer writer, final String indentToken, final int indentDepth ) throws IOException
	{
		appendPrivate ( writer, indentToken, indentDepth, System.getProperty ( "line.separator" ) );
	}

	/**
	 * The same as {@link append}, but writes the indent token in first line.
	 * 
	 * @param writer
	 * @param indentToken
	 * @param indentDepth
	 * @param lineSeparator
	 * @throws IOException
	 */
	private void writePrivate ( final Writer writer, final String indentToken, final int indentDepth, final String lineSeparator ) throws IOException
	{
		if ( indentToken != Analyzer.DEFAULT_INDENT_TOKEN )
		{
			writer.write ( indentToken );
			writer.write ( lineSeparator );
		}

		appendPrivate ( writer, indentToken, indentDepth, lineSeparator );
	}

	/**
	 * Writes the content.
	 * 
	 * @param writer
	 * @param indentToken
	 * @param indentDepth
	 * @param lineSeparator
	 * @throws IOException
	 */
	private void appendPrivate ( final Writer writer, final String indentToken, final int indentDepth, final String lineSeparator ) throws IOException
	{
		if ( name != null && !name.isEmpty ( ) )
		{
			// indent
			for ( int i = 0; i < indentDepth; i++ )
			{
				writer.write ( indentToken );
			}
			// node name
			writer.write ( name );

			// node value
			if ( value != null && !value.isEmpty ( ) )
			{
				writer.write ( ' ' );
				writer.write ( value );
			}
			// new line
			writer.write ( lineSeparator );
		}

		for ( final Node node: subnodes )
		{
			node.writePrivate ( writer, indentToken, indentDepth + 1, lineSeparator );
		}
	}

}
