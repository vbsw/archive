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


package com.github.vitalibaumtrok.glsled.preferences;


import org.eclipse.jface.preference.FieldEditorPreferencePage;
import org.eclipse.jface.preference.IPreferenceStore;
import org.eclipse.ui.IWorkbench;
import org.eclipse.ui.IWorkbenchPreferencePage;

import com.github.vitalibaumtrok.glsled.Activator;


/**
 * @author Vitali Baumtrok
 */
public class GLSLSyntaxColoringPreferencePage extends FieldEditorPreferencePage implements IWorkbenchPreferencePage
{

	@Override
	protected void createFieldEditors ( )
	{
		// TODO Auto-generated method stub
	}

	@Override
	public void init ( final IWorkbench workbench )
	{
		final Activator activator = Activator.getDefault ( );
		final IPreferenceStore preferenceStore = activator.getPreferenceStore ( );

		setPreferenceStore ( preferenceStore );
		setDescription ( "Font color and font style settings." );
	}

}
