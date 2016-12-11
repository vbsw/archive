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


import org.eclipse.jface.text.DefaultIndentLineAutoEditStrategy;
import org.eclipse.jface.text.DocumentCommand;
import org.eclipse.jface.text.IDocument;


/**
 * @author Vitali Baumtrok
 */
public class GLSLAutoEditStrategy extends DefaultIndentLineAutoEditStrategy
{

	@Override
	public void customizeDocumentCommand ( final IDocument d, final DocumentCommand c )
	{}

}
