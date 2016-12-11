/**
 * GLSL Ed (text editor for the OpenGL Shading Language)
 * Copyright (C) 2016 Vitali Baumtrok
 *
 * This file is part of GLSL Ed.
 *
 * GLSL Ed is distributed under the terms of the Eclipse Public License,
 * Version 1.0. (See accompanying file LICENSE or copy at
 * https://www.eclipse.org/legal/epl-v10.html)
 */


package com.github.vitalibaumtrok.glsled.editors;


import org.eclipse.swt.events.FocusEvent;
import org.eclipse.swt.events.FocusListener;


/**
 * @author Vitali Baumtrok
 */
public class GLSLFocusListener implements FocusListener
{

	private static final GLSLFocusListener instance = new GLSLFocusListener ( );

	private GLSLFocusListener ( )
	{}

	public static GLSLFocusListener getInstance ( )
	{
		return instance;
	}

	@Override
	public void focusGained ( final FocusEvent e )
	{
		GLSLContext.activate ( );
	}

	@Override
	public void focusLost ( final FocusEvent e )
	{
		GLSLContext.deactivate ( );
	}

}
