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


import org.eclipse.jface.text.IAutoEditStrategy;
import org.eclipse.jface.text.presentation.IPresentationReconciler;
import org.eclipse.jface.text.presentation.PresentationReconciler;
import org.eclipse.jface.text.source.ISourceViewer;
import org.eclipse.jface.text.source.SourceViewerConfiguration;


/**
 * @author Vitali Baumtrok
 */
public class GLSLSourceViewerConfiguration extends SourceViewerConfiguration
{

	@Override
	public IAutoEditStrategy[] getAutoEditStrategies ( final ISourceViewer sourceViewer, final String contentType )
	{
		if ( contentType.equals ( GLSLContentType.DEFAULT ) )
		{
			return new IAutoEditStrategy[] { new GLSLAutoEditStrategy ( ) };
		}
		else
		{
			return null;
		}
	}

	@Override
	public IPresentationReconciler getPresentationReconciler ( final ISourceViewer sourceViewer )
	{
		final PresentationReconciler reconciler = new PresentationReconciler ( );

		reconciler.setDocumentPartitioning ( GLSLContentType.DEFAULT );

		return reconciler;
	}

	@Override
	public String[] getConfiguredContentTypes ( final ISourceViewer sourceViewer )
	{
		// must be set, otherwise AutoEditStrategy will not work
		return new String[] { GLSLContentType.DEFAULT };
	}

}
