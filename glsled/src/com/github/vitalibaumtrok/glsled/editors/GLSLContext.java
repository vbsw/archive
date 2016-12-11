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


import org.eclipse.ui.IWorkbench;
import org.eclipse.ui.PlatformUI;
import org.eclipse.ui.contexts.IContextActivation;
import org.eclipse.ui.contexts.IContextService;


/**
 * @author Vitali Baumtrok
 */
public class GLSLContext
{

	public static final String ID = "com.github.vitalibaumtrok.glsled.editors.context"; //$NON-NLS-1$

	private static IContextActivation contextActivation;

	public static void activate ( )
	{
		if ( contextActivation == null )
		{
			final IWorkbench workbench = PlatformUI.getWorkbench ( );
			final IContextService contextService = workbench.getService ( IContextService.class );

			contextActivation = contextService.activateContext ( ID );
		}
	}

	public static void deactivate ( )
	{
		if ( contextActivation != null )
		{
			final IWorkbench workbench = PlatformUI.getWorkbench ( );
			final IContextService contextService = workbench.getService ( IContextService.class );

			contextService.deactivateContext ( contextActivation );

			contextActivation = null;
		}
	}

}
