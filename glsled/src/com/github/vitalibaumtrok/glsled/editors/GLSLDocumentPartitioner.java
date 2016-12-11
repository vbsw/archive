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


import org.eclipse.jface.text.DocumentEvent;
import org.eclipse.jface.text.IDocument;
import org.eclipse.jface.text.IDocumentPartitioner;
import org.eclipse.jface.text.ITypedRegion;
import org.eclipse.jface.text.TypedRegion;


/**
 * @author Vitali Baumtrok
 */
public class GLSLDocumentPartitioner implements IDocumentPartitioner
{

	private IDocument document;

	@Override
	public void connect ( final IDocument document )
	{
		this.document = document;
	}

	@Override
	public void disconnect ( )
	{
		this.document = null;
	}

	@Override
	public void documentAboutToBeChanged ( final DocumentEvent event )
	{}

	@Override
	public boolean documentChanged ( final DocumentEvent event )
	{
		return false;
	}

	@Override
	public String[] getLegalContentTypes ( )
	{
		return GLSLContentType.ALL;
	}

	@Override
	public String getContentType ( final int offset )
	{
		// AutoEditStrategy needs here a valid return
		return GLSLContentType.DEFAULT;
	}

	/**
	 * This is called only if the IPresentationReconciler has been set.
	 */
	@Override
	public ITypedRegion[] computePartitioning ( final int offset, final int length )
	{
		return null;
	}

	@Override
	public ITypedRegion getPartition ( final int offset )
	{
		return new TypedRegion ( 0, document.getLength ( ), GLSLContentType.DEFAULT );
	}

}
