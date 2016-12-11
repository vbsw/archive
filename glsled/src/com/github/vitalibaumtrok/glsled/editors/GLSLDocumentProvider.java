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


import org.eclipse.core.runtime.CoreException;
import org.eclipse.jface.text.IDocument;
import org.eclipse.jface.text.IDocumentPartitioner;
import org.eclipse.ui.editors.text.FileDocumentProvider;


/**
 * @author Vitali Baumtrok
 */
public class GLSLDocumentProvider extends FileDocumentProvider
{

	@Override
	protected IDocument createDocument ( final Object element ) throws CoreException
	{
		final IDocument document = super.createDocument ( element );

		if ( null != document )
		{
			final IDocumentPartitioner partitioner = new GLSLDocumentPartitioner ( );

			partitioner.connect ( document );
			document.setDocumentPartitioner ( partitioner );
		}

		return document;
	}

}
