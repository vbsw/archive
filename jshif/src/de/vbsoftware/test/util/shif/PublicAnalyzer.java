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


package de.vbsoftware.test.util.shif;


import java.lang.reflect.Field;
import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Method;

import de.vbsoftware.util.shif.Analyzer;
import de.vbsoftware.util.shif.Result;


/**
 * @author Vitali Baumtrok
 */
public class PublicAnalyzer
{

	private final Analyzer analyzer;

	private final Field field_result;
	private final Field field_indentToken;
	private final Field field_invalidIndentDepth;
	private final Field field_invalidNodeName;
	private final Field field_indentDepthCurr;
	private final Field field_indentDepthPrev;
	private final Field field_lineNumber;
	private final Field field_iCurr;

	private final Method method_skipLine;
	private final Method method_skipEmptyLines;
	private final Method method_isIndentToken;
	private final Method method_skipBlanksOrTabs;
	private final Method method_processIndentDepth;
	private final Method method_processIndentToken;
	private final Method method_hasNext;
	private final Method method_processNode;

	private static Field createField ( final String name ) throws NoSuchFieldException, SecurityException
	{
		final Field field = Analyzer.class.getDeclaredField ( name );

		field.setAccessible ( true );

		return field;
	}

	private static Method createMethod ( final String name, final Class<?>... parameterTypes ) throws NoSuchMethodException, SecurityException
	{
		final Method method = Analyzer.class.getDeclaredMethod ( name, parameterTypes );

		method.setAccessible ( true );

		return method;
	}

	public PublicAnalyzer ( final String text ) throws NoSuchFieldException, SecurityException, NoSuchMethodException
	{
		analyzer = new Analyzer ( text );
		field_result = createField ( "result" );
		field_indentToken = createField ( "indentToken" );
		field_invalidIndentDepth = createField ( "invalidIndentDepth" );
		field_invalidNodeName = createField ( "invalidNodeName" );
		field_indentDepthCurr = createField ( "indentDepthCurr" );
		field_indentDepthPrev = createField ( "indentDepthPrev" );
		field_lineNumber = createField ( "lineNumber" );
		field_iCurr = createField ( "iCurr" );
		method_skipLine = createMethod ( "skipLine" );
		method_skipEmptyLines = createMethod ( "skipEmptyLines" );
		method_isIndentToken = createMethod ( "isIndentToken" );
		method_skipBlanksOrTabs = createMethod ( "skipBlanksOrTabs", boolean.class, boolean.class );
		method_processIndentDepth = createMethod ( "processIndentDepth" );
		method_processIndentToken = createMethod ( "processIndentToken" );
		method_hasNext = createMethod ( "hasNext" );
		method_processNode = createMethod ( "processNode" );
	}

	public PublicAnalyzer ( final String text, final int skip ) throws NoSuchFieldException, SecurityException, NoSuchMethodException
	{
		analyzer = new Analyzer ( text, skip );
		field_result = createField ( "result" );
		field_indentToken = createField ( "indentToken" );
		field_invalidIndentDepth = createField ( "invalidIndentDepth" );
		field_invalidNodeName = createField ( "invalidNodeName" );
		field_indentDepthCurr = createField ( "indentDepthCurr" );
		field_indentDepthPrev = createField ( "indentDepthPrev" );
		field_lineNumber = createField ( "lineNumber" );
		field_iCurr = createField ( "iCurr" );
		method_skipLine = createMethod ( "skipLine" );
		method_skipEmptyLines = createMethod ( "skipEmptyLines" );
		method_isIndentToken = createMethod ( "isIndentToken" );
		method_skipBlanksOrTabs = createMethod ( "skipBlanksOrTabs", boolean.class, boolean.class );
		method_processIndentDepth = createMethod ( "processIndentDepth" );
		method_processIndentToken = createMethod ( "processIndentToken" );
		method_hasNext = createMethod ( "hasNext" );
		method_processNode = createMethod ( "processNode" );
	}

	public String getIndentToken ( ) throws IllegalArgumentException, IllegalAccessException
	{
		return (String) field_indentToken.get ( analyzer );
	}

	public boolean getInvalidIndentDepth ( ) throws IllegalArgumentException, IllegalAccessException
	{
		return field_invalidIndentDepth.getBoolean ( analyzer );
	}

	public boolean getInvalidNodeName ( ) throws IllegalArgumentException, IllegalAccessException
	{
		return field_invalidNodeName.getBoolean ( analyzer );
	}

	public int getIndentDepthCurr ( ) throws IllegalArgumentException, IllegalAccessException
	{
		return field_indentDepthCurr.getInt ( analyzer );
	}

	public void setIndentDepthCurr ( final int indentDepthCurr ) throws IllegalArgumentException, IllegalAccessException
	{
		field_indentDepthCurr.setInt ( analyzer, indentDepthCurr );
	}

	public int getIndentDepthPrev ( ) throws IllegalArgumentException, IllegalAccessException
	{
		return field_indentDepthPrev.getInt ( analyzer );
	}

	public int getLineNumber ( ) throws IllegalArgumentException, IllegalAccessException
	{
		return field_lineNumber.getInt ( analyzer );
	}

	public int getICurr ( ) throws IllegalArgumentException, IllegalAccessException
	{
		return field_iCurr.getInt ( analyzer );
	}

	public void setICurr ( final int value ) throws IllegalArgumentException, IllegalAccessException
	{
		field_iCurr.setInt ( analyzer, value );
	}

	public Result getResult ( ) throws IllegalArgumentException, IllegalAccessException
	{
		return (Result) field_result.get ( analyzer );
	}

	public void skipLine ( ) throws IllegalAccessException, IllegalArgumentException, InvocationTargetException
	{
		method_skipLine.invoke ( analyzer );
	}

	public void skipEmptyLines ( ) throws IllegalAccessException, IllegalArgumentException, InvocationTargetException
	{
		method_skipEmptyLines.invoke ( analyzer );
	}

	public boolean isIndentToken ( ) throws IllegalAccessException, IllegalArgumentException, InvocationTargetException
	{
		return (boolean) method_isIndentToken.invoke ( analyzer );
	}

	public void setIndentToken ( final String indentToken ) throws IllegalArgumentException, IllegalAccessException
	{
		field_indentToken.set ( analyzer, indentToken );
	}

	public void skipBlanksOrTabs ( final boolean skipBlanks, final boolean skipTabs ) throws IllegalAccessException, IllegalArgumentException, InvocationTargetException
	{
		method_skipBlanksOrTabs.invoke ( analyzer, skipBlanks, skipTabs );
	}

	public void processIndentDepth ( ) throws IllegalAccessException, IllegalArgumentException, InvocationTargetException
	{
		method_processIndentDepth.invoke ( analyzer );
	}

	public void processIndentToken ( ) throws IllegalAccessException, IllegalArgumentException, InvocationTargetException
	{
		method_processIndentToken.invoke ( analyzer );
	}

	public boolean hasNext ( ) throws IllegalAccessException, IllegalArgumentException, InvocationTargetException
	{
		return (boolean) method_hasNext.invoke ( analyzer );
	}

	public void processNode ( ) throws IllegalAccessException, IllegalArgumentException, InvocationTargetException
	{
		method_processNode.invoke ( analyzer );
	}

}
