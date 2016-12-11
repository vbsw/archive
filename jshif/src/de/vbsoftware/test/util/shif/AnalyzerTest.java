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


import static org.junit.Assert.assertFalse;
import static org.junit.Assert.assertTrue;

import java.lang.reflect.InvocationTargetException;
import java.util.Iterator;

import org.junit.Test;

import de.vbsoftware.util.shif.Analyzer;
import de.vbsoftware.util.shif.Result;


/**
 * @author Vitali Baumtrok
 */
public class AnalyzerTest
{

	private static final String TEST_STRING = "\n# just a comment\n\n\nfile\n\tdirectory ./\n\tname   test.txt  ";
	private static final String TEST_INDENT_TOKEN = "asdf123";
	private static final String TEST_STRING_FOR_TOKEN = TEST_INDENT_TOKEN + "\n\n\n\nbunny\n";

	@Test
	public void testSkipLine ( ) throws NoSuchFieldException, SecurityException, NoSuchMethodException, IllegalArgumentException, IllegalAccessException, InvocationTargetException
	{
		final PublicAnalyzer analyzer = new PublicAnalyzer ( AnalyzerTest.TEST_STRING );
		final int initialLineNumber = analyzer.getLineNumber ( );

		analyzer.skipLine ( );
		assertTrue ( analyzer.getICurr ( ) == 1 ); // "# just a comment"
		analyzer.skipLine ( );
		assertTrue ( analyzer.getICurr ( ) == 18 ); // empty
		analyzer.skipLine ( );
		assertTrue ( analyzer.getICurr ( ) == 19 ); // empty
		analyzer.skipLine ( );
		assertTrue ( analyzer.getICurr ( ) == 20 ); // "file"
		analyzer.setICurr ( analyzer.getICurr ( ) + 3 );
		analyzer.skipLine ( );
		assertTrue ( analyzer.getICurr ( ) == 25 ); // "\tdirectory"
		analyzer.skipLine ( );
		analyzer.skipLine ( );
		analyzer.skipLine ( );
		assertTrue ( analyzer.getICurr ( ) == AnalyzerTest.TEST_STRING.length ( ) );
		assertTrue ( analyzer.getLineNumber ( ) == initialLineNumber );
	}

	@Test
	public void testSkipEmptyLines ( ) throws IllegalAccessException, IllegalArgumentException, InvocationTargetException, NoSuchFieldException, SecurityException, NoSuchMethodException
	{
		final PublicAnalyzer analyzer = new PublicAnalyzer ( AnalyzerTest.TEST_STRING );
		final int initialLineNumber = analyzer.getLineNumber ( );

		analyzer.skipEmptyLines ( );
		assertTrue ( analyzer.getICurr ( ) == 1 ); // "# just a comment"
		analyzer.skipLine ( );
		assertTrue ( analyzer.getLineNumber ( ) == 0 );
		assertTrue ( analyzer.getICurr ( ) == 18 ); // empty
		analyzer.skipEmptyLines ( );
		assertTrue ( analyzer.getLineNumber ( ) == 2 );
		assertTrue ( analyzer.getICurr ( ) == 20 ); // "file"
		analyzer.skipEmptyLines ( );
		assertTrue ( analyzer.getICurr ( ) == 20 ); // "file"
		analyzer.skipLine ( );
		assertTrue ( analyzer.getICurr ( ) == 25 ); // "\tdirectory"
		analyzer.skipEmptyLines ( );
		assertTrue ( analyzer.getICurr ( ) == 25 ); // "\tdirectory"
		analyzer.skipLine ( );
		analyzer.skipLine ( );
		analyzer.skipLine ( );
		assertTrue ( analyzer.getICurr ( ) == AnalyzerTest.TEST_STRING.length ( ) );
		assertTrue ( analyzer.getLineNumber ( ) > initialLineNumber );
	}

	@Test
	public void testSkipBlanksOrTabs ( ) throws NoSuchFieldException, SecurityException, NoSuchMethodException, IllegalArgumentException, IllegalAccessException, InvocationTargetException
	{
		final PublicAnalyzer analyzer = new PublicAnalyzer ( AnalyzerTest.TEST_STRING );
		final int offset = AnalyzerTest.TEST_STRING.length ( ) - 12;

		analyzer.setICurr ( offset ); // "  test.txt"
		assertTrue ( AnalyzerTest.TEST_STRING.substring ( analyzer.getICurr ( ) ).equals ( "  test.txt  " ) );
		analyzer.skipBlanksOrTabs ( true, false );
		assertTrue ( analyzer.getICurr ( ) == offset + 2 );
		analyzer.setICurr ( offset );
		analyzer.skipBlanksOrTabs ( false, false );
		assertTrue ( analyzer.getICurr ( ) == offset );
		analyzer.skipBlanksOrTabs ( false, true );
		assertTrue ( analyzer.getICurr ( ) == offset );
	}

	@Test
	public void testIsIndentToken ( ) throws NoSuchFieldException, SecurityException, NoSuchMethodException, IllegalArgumentException, IllegalAccessException, InvocationTargetException
	{
		final PublicAnalyzer analyzer = new PublicAnalyzer ( AnalyzerTest.TEST_STRING );

		analyzer.setIndentToken ( "\t" );
		analyzer.setICurr ( 25 ); // "\tdirectory"
		assertTrue ( analyzer.isIndentToken ( ) );
	}

	@Test
	public void testProcessIndentDepth ( ) throws NoSuchFieldException, SecurityException, NoSuchMethodException, IllegalAccessException, IllegalArgumentException, InvocationTargetException
	{
		final PublicAnalyzer analyzer = new PublicAnalyzer ( AnalyzerTest.TEST_STRING );

		analyzer.setIndentToken ( Analyzer.DEFAULT_INDENT_TOKEN );
		analyzer.processIndentDepth ( );
		assertTrue ( analyzer.getResult ( ).isValid ( ) );
		assertTrue ( analyzer.getIndentDepthCurr ( ) == 0 );
		assertTrue ( analyzer.getIndentDepthPrev ( ) == -1 );
		assertTrue ( analyzer.getLineNumber ( ) == 0 );
		assertTrue ( analyzer.getICurr ( ) == 0 );

		analyzer.setIndentDepthCurr ( -1 );
		analyzer.setICurr ( 25 ); // "\tdirectory"
		analyzer.processIndentDepth ( );
		assertTrue ( !analyzer.getResult ( ).isValid ( ) );
		assertTrue ( analyzer.getResult ( ).isInvalidIndentDepth ( ) );
		assertFalse ( analyzer.getResult ( ).isInvalidNodeName ( ) );

		analyzer.setIndentDepthCurr ( 0 );
		analyzer.setICurr ( 25 ); // "\tdirectory"
		analyzer.processIndentDepth ( );
		assertFalse ( analyzer.getResult ( ).isInvalidIndentDepth ( ) );
		assertTrue ( analyzer.getResult ( ).isValid ( ) );
	}

	@Test
	public void testProcessIndentToken ( ) throws NoSuchFieldException, SecurityException, NoSuchMethodException, IllegalAccessException, IllegalArgumentException, InvocationTargetException
	{
		final PublicAnalyzer analyzer = new PublicAnalyzer ( AnalyzerTest.TEST_STRING );
		final PublicAnalyzer analyzerTokenString = new PublicAnalyzer ( AnalyzerTest.TEST_STRING_FOR_TOKEN );
		final PublicAnalyzer analyzerTokenOnly = new PublicAnalyzer ( AnalyzerTest.TEST_INDENT_TOKEN );

		analyzer.processIndentToken ( );
		assertTrue ( analyzer.getIndentToken ( ) == Analyzer.DEFAULT_INDENT_TOKEN );

		assertTrue ( analyzerTokenString.getIndentDepthCurr ( ) == -1 );
		assertTrue ( analyzerTokenString.getIndentDepthPrev ( ) == -1 );
		assertTrue ( analyzerTokenString.hasNext ( ) );
		assertTrue ( analyzerTokenString.getLineNumber ( ) == -1 );
		analyzerTokenString.processIndentToken ( );
		assertTrue ( analyzerTokenString.getIndentToken ( ).equals ( AnalyzerTest.TEST_INDENT_TOKEN ) );
		assertTrue ( analyzerTokenString.getICurr ( ) == AnalyzerTest.TEST_INDENT_TOKEN.length ( ) + 1 );
		assertTrue ( analyzerTokenString.getIndentDepthCurr ( ) == 0 ); // first line
		assertTrue ( analyzerTokenString.getIndentDepthPrev ( ) == -1 );
		assertTrue ( analyzerTokenString.hasNext ( ) );
		assertTrue ( analyzerTokenString.getLineNumber ( ) == 0 );

		assertTrue ( analyzerTokenOnly.getIndentDepthCurr ( ) == -1 );
		assertTrue ( analyzerTokenOnly.getIndentDepthPrev ( ) == -1 );
		assertTrue ( analyzerTokenOnly.hasNext ( ) );
		assertTrue ( analyzerTokenOnly.getLineNumber ( ) == -1 );
		analyzerTokenOnly.processIndentToken ( );
		assertTrue ( analyzerTokenOnly.getIndentToken ( ).equals ( AnalyzerTest.TEST_INDENT_TOKEN ) );
		assertTrue ( analyzerTokenOnly.getICurr ( ) == AnalyzerTest.TEST_INDENT_TOKEN.length ( ) );
		assertTrue ( analyzerTokenOnly.getIndentDepthCurr ( ) == 0 ); // first line
		assertTrue ( analyzerTokenOnly.getIndentDepthPrev ( ) == -1 );
		assertFalse ( analyzerTokenOnly.hasNext ( ) );
		assertTrue ( analyzerTokenOnly.getLineNumber ( ) == 0 );
	}

	@Test
	public void testProcessNode ( ) throws NoSuchFieldException, SecurityException, NoSuchMethodException, IllegalArgumentException, IllegalAccessException, InvocationTargetException
	{
		final PublicAnalyzer analyzerSkipNone = new PublicAnalyzer ( AnalyzerTest.TEST_STRING, Analyzer.SKIP_NONE );
		final PublicAnalyzer analyzerSkipAll = new PublicAnalyzer ( AnalyzerTest.TEST_STRING, Analyzer.SKIP_ALL );

		analyzerSkipNone.setICurr ( 20 ); // "file"
		analyzerSkipNone.processNode ( );
		assertTrue ( analyzerSkipNone.getResult ( ).getNodeNameAsString ( ).equals ( "file" ) );
		assertTrue ( analyzerSkipNone.getResult ( ).getNodeValueAsString ( ).equals ( "" ) );
		assertTrue ( analyzerSkipNone.getICurr ( ) == 25 ); // "\tdirectory"

		analyzerSkipNone.setICurr ( 26 ); // "directory"
		analyzerSkipNone.processNode ( );
		assertTrue ( analyzerSkipNone.getResult ( ).getNodeNameAsString ( ).equals ( "directory" ) );
		assertTrue ( analyzerSkipNone.getResult ( ).getNodeValueAsString ( ).equals ( "./" ) );
		assertTrue ( analyzerSkipNone.getICurr ( ) == 39 ); // "\tname   test.txt"

		analyzerSkipNone.setICurr ( 40 ); // "name   test.txt"
		analyzerSkipNone.processNode ( );
		assertTrue ( analyzerSkipNone.getResult ( ).getNodeNameAsString ( ).equals ( "name" ) );
		assertTrue ( analyzerSkipNone.getResult ( ).getNodeValueAsString ( ).equals ( "  test.txt  " ) );

		analyzerSkipAll.setICurr ( 40 ); // "name   test.txt"
		analyzerSkipAll.processNode ( );
		assertTrue ( analyzerSkipAll.getResult ( ).getNodeNameAsString ( ).equals ( "name" ) );
		assertTrue ( analyzerSkipAll.getResult ( ).getNodeValueAsString ( ).equals ( "test.txt" ) );
	}

	@Test
	public void testProcessLine ( )
	{
		final Analyzer analyzer = new Analyzer ( AnalyzerTest.TEST_STRING );
		final Iterator<Result> iter = analyzer.iterator ( );

		final Result result = iter.next ( );
		assertTrue ( analyzer.getIndentToken ( ) == Analyzer.DEFAULT_INDENT_TOKEN );
		assertTrue ( result.getNodeNameAsString ( ).equals ( "#" ) );
		assertTrue ( result.getNodeValueAsString ( ).equals ( "just a comment" ) );
		assertTrue ( iter.hasNext ( ) );
		iter.next ( );
		assertTrue ( result.getNodeNameAsString ( ).equals ( "file" ) );
		assertTrue ( result.getNodeValueAsString ( ).equals ( "" ) );
		assertTrue ( iter.hasNext ( ) );
		iter.next ( );
		assertTrue ( result.getNodeNameAsString ( ).equals ( "directory" ) );
		assertTrue ( result.getNodeValueAsString ( ).equals ( "./" ) );
		assertTrue ( iter.hasNext ( ) );
		iter.next ( );
		assertTrue ( result.getNodeNameAsString ( ).equals ( "name" ) );
		assertTrue ( result.getNodeValueAsString ( ).equals ( "test.txt" ) );
		assertFalse ( iter.hasNext ( ) );
	}

}
