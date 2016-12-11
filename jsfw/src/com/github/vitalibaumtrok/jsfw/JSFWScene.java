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


import javafx.beans.property.ReadOnlyBooleanProperty;
import javafx.beans.property.ReadOnlyObjectProperty;
import javafx.beans.value.ChangeListener;
import javafx.beans.value.ObservableValue;
import javafx.event.EventHandler;
import javafx.scene.Parent;
import javafx.scene.Scene;
import javafx.scene.SceneAntialiasing;
import javafx.scene.input.KeyCode;
import javafx.scene.input.KeyEvent;
import javafx.scene.input.MouseButton;
import javafx.scene.input.MouseEvent;
import javafx.scene.paint.Paint;
import javafx.stage.Stage;
import javafx.stage.Window;


/**
 * @author Vitali Baumtrok
 */
public class JSFWScene extends Scene
{
	protected final JSFWScreens screens;
	protected boolean fullScreenEscape = false;

	public JSFWScene ( final JSFWLayeredPane root )
	{
		super ( root );

		screens = createScreens ( this );

		initListeners ( );
	}

	public JSFWScene ( final JSFWLayeredPane root, final Paint fill )
	{
		super ( root, fill );

		screens = createScreens ( this );

		initListeners ( );
	}

	public JSFWScene ( final JSFWLayeredPane root, final double width, final double height )
	{
		super ( root, width, height );

		screens = createScreens ( this );

		initListeners ( );
	}

	public JSFWScene ( final JSFWLayeredPane root, final double width, final double height, final boolean depthBuffer )
	{
		super ( root, width, height, depthBuffer );

		screens = createScreens ( this );

		initListeners ( );
	}

	public JSFWScene ( final JSFWLayeredPane root, final double width, final double height, final Paint fill )
	{
		super ( root, width, height, fill );

		screens = createScreens ( this );

		initListeners ( );
	}

	public JSFWScene ( final JSFWLayeredPane root, final double width, final double height, final boolean depthBuffer, final SceneAntialiasing antiAliasing )
	{
		super ( root, width, height, depthBuffer, antiAliasing );

		screens = createScreens ( this );

		initListeners ( );
	}

	public boolean isFullScreen ( )
	{
		final Window window = getWindow ( );

		if ( window != null && window instanceof Stage )
		{
			final Stage stage = (Stage) window;

			return stage.isFullScreen ( );
		}

		return false;
	}

	public JSFWLayeredPane getLayeredPaneRoot ( )
	{
		final Parent root = getRoot ( );

		return (JSFWLayeredPane) root;
	}

	protected JSFWScreens createScreens ( final JSFWScene scene )
	{
		final JSFWLayeredPane layeredPane = scene.getLayeredPaneRoot ( );
		final JSFWScreens screens = new JSFWScreens ( layeredPane );

		return screens;
	}

	protected void closeStage ( )
	{
		final javafx.stage.Window window = getWindow ( );
		final Stage stage = (Stage) window;

		stage.close ( );
	}

	protected void initListeners ( )
	{
		final ReadOnlyObjectProperty <Window> windowProp = windowProperty ( );

		addEventHandler ( KeyEvent.KEY_PRESSED, new KeyPressedListener ( ) );
		addEventHandler ( KeyEvent.KEY_RELEASED, new KeyReleasedListener ( ) );
		addEventHandler ( MouseEvent.MOUSE_MOVED, new MouseMovedListener ( ) );
		addEventHandler ( MouseEvent.MOUSE_DRAGGED, new MouseDraggedListener ( ) );
		addEventHandler ( MouseEvent.MOUSE_ENTERED, new MouseEnteredListener ( ) );
		addEventHandler ( MouseEvent.MOUSE_EXITED, new MouseExitedListener ( ) );
		addEventHandler ( MouseEvent.MOUSE_PRESSED, new MouseButtonPressedListener ( ) );
		addEventHandler ( MouseEvent.MOUSE_RELEASED, new MouseButtonReleasedListener ( ) );
		windowProp.addListener ( new WindowListener ( ) );
	}

	protected void keyPressed ( final KeyCode keyCode )
	{
		if ( keyCode == KeyCode.UP && screens.isMenu ( ) )
		{
			final JSFWScreenMenu menu = screens.getMenu ( );

			menu.moveSelectorUp ( );
		}
		else if ( keyCode == KeyCode.DOWN && screens.isMenu ( ) )
		{
			final JSFWScreenMenu menu = screens.getMenu ( );

			menu.moveSelectorDown ( );
		}
		else if ( keyCode == KeyCode.ENTER && screens.isMenu ( ) )
		{
			final JSFWScreenMenu menu = screens.getMenu ( );

			menu.select ( );
		}
		else if ( keyCode.equals ( KeyCode.ESCAPE ) )
		{
			if ( fullScreenEscape == false )
			{
				closeStage ( );
			}
		}
		else if ( keyCode.equals ( KeyCode.F11 ) )
		{
			if ( fullScreenEscape == false )
			{
				final javafx.stage.Window window = getWindow ( );
				final Stage stage = (Stage) window;

				stage.setFullScreen ( true );
			}
		}
	}

	protected void keyReleased ( final KeyCode keyCode )
	{
		if ( keyCode.equals ( KeyCode.ESCAPE ) )
		{
			if ( fullScreenEscape )
			{
				fullScreenEscape = false;
			}
		}
	}

	protected void mouseMoved ( final double x, final double y )
	{}

	protected void mouseDragged ( final double x, final double y )
	{}

	protected void mouseEntered ( final double x, final double y )
	{}

	protected void mouseExited ( final double x, final double y )
	{}

	protected void mouseButtonPressed ( final MouseButton button )
	{}

	protected void mouseButtonReleased ( final MouseButton button )
	{}

	protected void fullScreenChanged ( final Boolean fullScreen )
	{
		if ( fullScreen )
		{
			fullScreenEscape = true;
		}
	}

	protected class KeyPressedListener implements EventHandler <KeyEvent>
	{
		@Override
		public void handle ( final KeyEvent event )
		{
			final KeyCode keyCode = event.getCode ( );

			keyPressed ( keyCode );
		}
	}

	protected class KeyReleasedListener implements EventHandler <KeyEvent>
	{
		@Override
		public void handle ( final KeyEvent event )
		{
			final KeyCode keyCode = event.getCode ( );

			keyReleased ( keyCode );
		}
	}

	protected class MouseMovedListener implements EventHandler <MouseEvent>
	{
		@Override
		public void handle ( final MouseEvent event )
		{
			final double x = event.getX ( );
			final double y = event.getY ( );

			mouseMoved ( x, y );
		}
	}

	protected class MouseDraggedListener implements EventHandler <MouseEvent>
	{
		@Override
		public void handle ( final MouseEvent event )
		{
			final double x = event.getX ( );
			final double y = event.getY ( );

			mouseDragged ( x, y );
		}
	}

	protected class MouseEnteredListener implements EventHandler <MouseEvent>
	{
		@Override
		public void handle ( final MouseEvent event )
		{
			final double x = event.getX ( );
			final double y = event.getY ( );

			mouseEntered ( x, y );
		}
	}

	protected class MouseExitedListener implements EventHandler <MouseEvent>
	{
		@Override
		public void handle ( final MouseEvent event )
		{
			final double x = event.getX ( );
			final double y = event.getY ( );

			mouseExited ( x, y );
		}
	}

	protected class MouseButtonPressedListener implements EventHandler <MouseEvent>
	{
		@Override
		public void handle ( final MouseEvent event )
		{
			final MouseButton button = event.getButton ( );

			mouseButtonPressed ( button );
		}
	}

	protected class MouseButtonReleasedListener implements EventHandler <MouseEvent>
	{
		@Override
		public void handle ( final MouseEvent event )
		{
			final MouseButton button = event.getButton ( );

			mouseButtonReleased ( button );
		}
	}

	protected class WindowListener implements ChangeListener <Window>
	{
		private final FullScreenListener fullScreenListener = new FullScreenListener ( );

		@Override
		public void changed ( final ObservableValue <? extends Window> observable, final Window oldValue, final Window newValue )
		{
			if ( oldValue == null )
			{
				if ( newValue != null )
				{
					if ( newValue instanceof Stage )
					{
						final Stage stage = (Stage) newValue;
						final ReadOnlyBooleanProperty fullScreenProp = stage.fullScreenProperty ( );

						fullScreenProp.addListener ( fullScreenListener );
					}
				}
			}
			else
			{
				if ( newValue != oldValue )
				{
					if ( oldValue instanceof Stage )
					{
						final Stage stage = (Stage) oldValue;
						final ReadOnlyBooleanProperty fullScreenProp = stage.fullScreenProperty ( );

						fullScreenProp.removeListener ( fullScreenListener );
					}
					if ( newValue != null && newValue instanceof Stage )
					{
						final Stage stage = (Stage) newValue;
						final ReadOnlyBooleanProperty fullScreenProp = stage.fullScreenProperty ( );

						fullScreenProp.addListener ( fullScreenListener );
					}
				}
			}
		}
	}

	protected class FullScreenListener implements ChangeListener <Boolean>
	{
		@Override
		public void changed ( final ObservableValue <? extends Boolean> observable, final Boolean oldValue, final Boolean newValue )
		{
			fullScreenChanged ( newValue );
		}
	}
}
