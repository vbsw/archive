/**
 * This is free and unencumbered software released into the public domain.
 *
 * Anyone is free to copy, modify, publish, use, compile, sell, or
 * distribute this software, either in source code form or as a compiled
 * binary, for any purpose, commercial or non-commercial, and by any
 * means.
 *
 * In jurisdictions that recognize copyright laws, the author or authors
 * of this software dedicate any and all copyright interest in the
 * software to the public domain. We make this dedication for the benefit
 * of the public at large and to the detriment of our heirs and
 * successors. We intend this dedication to be an overt act of
 * relinquishment in perpetuity of all present and future rights to this
 * software under copyright law.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 * IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
 * OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
 * ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
 * OTHER DEALINGS IN THE SOFTWARE.
 *
 * For more information, please refer to <http://unlicense.org>
 */


package de.vbsoftware.log.st;


import java.io.BufferedWriter;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.StandardOpenOption;
import java.text.DateFormat;
import java.text.SimpleDateFormat;
import java.util.Date;


/**
 * Logger.<br>
 * <br>
 * To set output to a file do
 * <pre> Log.output = LogOutput.FILE;
 * Log.file = Paths.get ( System.getProperty ( "user.home" ), "info.log" );</pre>
 * Use methods {@link Log.error}, {@link Log.warning},
 * {@link Log.info} and {@link Log.trace} to log.<br>
 * <br>
 * 
 * @author Vitali Baumtrok
 */
public class Log
{

	private static final StackTracePrintWriter writer = new StackTracePrintWriter ( );

	public static LogOutput output = LogOutput.STD_OUT;
	public static Path file = Paths.get ( "error.log" );
	public static DateFormat timeFormat = new SimpleDateFormat ( "YYYY.MM.dd HH:mm:ss" );
	public static Date timeStamp = null;
	public static String message = null;
	public static String stackTrace = null;
	public static boolean levelError = true;
	public static boolean levelWarning = true;
	public static boolean levelInfo = true;
	public static boolean levelTrace = true;

	/**
	 * If <code>true</code> then the next write to the log file
	 * will overwrite the log file. Afterwards this variable is
	 * set to <code>false</code>.
	 */
	public static boolean newFile = true;

	public static void error ( final String message )
	{
		if ( levelError )
		{
			stackTrace = null;

			createErrorMessage ( message );
			flush ( );
		}
	}

	public static void warning ( final String message )
	{
		if ( levelWarning )
		{
			stackTrace = null;

			createWarningMessage ( message );
			flush ( );
		}
	}

	public static void info ( final String message )
	{
		if ( levelInfo )
		{
			stackTrace = null;

			createInfoMessage ( message );
			flush ( );
		}
	}

	public static void trace ( final String message )
	{
		if ( levelTrace )
		{
			stackTrace = null;

			createTraceMessage ( message );
			flush ( );
		}
	}

	public static void error ( final String message, final Throwable t )
	{
		if ( levelError )
		{
			createErrorMessage ( message );
			createStackTrace ( t );
			flush ( );
		}
	}

	public static void warning ( final String message, final Throwable t )
	{
		if ( levelWarning )
		{
			createWarningMessage ( message );
			createStackTrace ( t );
			flush ( );
		}
	}

	public static void info ( final String message, final Throwable t )
	{
		if ( levelInfo )
		{
			createInfoMessage ( message );
			createStackTrace ( t );
			flush ( );
		}
	}

	public static void trace ( final String message, final Throwable t )
	{
		if ( levelInfo )
		{
			createInfoMessage ( message );
			createStackTrace ( t );
			flush ( );
		}
	}

	public static void createErrorMessage ( final String message )
	{
		timeStamp = new Date ( );
		Log.message = timeFormat.format ( timeStamp ) + " Error: " + message;
	}

	public static void createWarningMessage ( final String message )
	{
		timeStamp = new Date ( );
		Log.message = timeFormat.format ( timeStamp ) + " Warning: " + message;
	}

	public static void createInfoMessage ( final String message )
	{
		timeStamp = new Date ( );
		Log.message = timeFormat.format ( timeStamp ) + " Info: " + message;
	}

	public static void createTraceMessage ( final String message )
	{
		timeStamp = new Date ( );
		Log.message = timeFormat.format ( timeStamp ) + " Trace: " + message;
	}

	public static void createStackTrace ( final Throwable t )
	{
		if ( t != null )
		{
			t.printStackTrace ( writer );

			stackTrace = writer.toString ( );
		}
		else
		{
			stackTrace = "<stack trace missing>\n";
		}
	}

	public static boolean messageIsEmpty ( )
	{
		return stringIsEmpty ( message );
	}

	public static boolean stackTraceIsEmpty ( )
	{
		return stringIsEmpty ( stackTrace );
	}

	private static boolean stringIsEmpty ( final String string )
	{
		if ( string != null )
		{
			for ( int i = 0; i < string.length ( ); i++ )
			{
				final char c = string.charAt ( i );

				if ( c != ' ' && c != '\n' && c != '\t' && c != '\r' )
				{
					return false;
				}
			}
		}
		return true;
	}

	public static void flush ( )
	{
		if ( !messageIsEmpty ( ) )
		{
			switch ( output )
			{
				case STD_OUT:
				System.out.println ( message );

				if ( !stackTraceIsEmpty ( ) )
				{
					System.out.print ( stackTrace );
				}
				break;

				case STD_ERR:
				System.err.println ( message );

				if ( !stackTraceIsEmpty ( ) )
				{
					System.err.print ( stackTrace );
				}
				break;

				case FILE:
				try
				{
					if ( Files.exists ( file ) )
					{
						if ( newFile )
						{
							newFile = false;

							Files.delete ( file );
							Files.createFile ( file );
						}
					}
					else
					{
						Files.createFile ( file );
					}
				}
				catch ( IOException e )
				{
					e.printStackTrace ( );
				}

				try ( BufferedWriter bw = Files.newBufferedWriter ( file, StandardOpenOption.APPEND ) )
				{
					bw.write ( message );
					bw.write ( "\n" );

					if ( !stackTraceIsEmpty ( ) )
					{
						bw.write ( stackTrace );
					}
				}
				catch ( IOException e )
				{
					e.printStackTrace ( );
				}
				break;
			}
		}
	}

}
