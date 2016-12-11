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


import org.eclipse.jface.text.source.ISourceViewer;
import org.eclipse.jface.text.source.IVerticalRuler;
import org.eclipse.swt.custom.StyledText;
import org.eclipse.swt.widgets.Composite;
import org.eclipse.ui.editors.text.TextEditor;


/**
 * @author Vitali Baumtrok
 */
public class GLSLEditor extends TextEditor
{

	public GLSLEditor ( )
	{
		// sets GLSLDocumentPartitioner
		setDocumentProvider ( new GLSLDocumentProvider ( ) );
		// sets GLSLAutoEditStrategy
		setSourceViewerConfiguration ( new GLSLSourceViewerConfiguration ( ) );
	}

	@Override
	protected ISourceViewer createSourceViewer ( final Composite parent, final IVerticalRuler ruler, final int styles )
	{
		final ISourceViewer viewer = super.createSourceViewer ( parent, ruler, styles );
		final StyledText styledText = viewer.getTextWidget ( );

		if ( styledText != null )
		{
			final GLSLFocusListener focusListener = GLSLFocusListener.getInstance ( );

			styledText.addFocusListener ( focusListener );
		}

		return viewer;
	}

}
