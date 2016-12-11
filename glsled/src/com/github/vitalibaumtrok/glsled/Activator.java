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


package com.github.vitalibaumtrok.glsled;


import org.eclipse.jface.resource.ImageDescriptor;
import org.eclipse.ui.plugin.AbstractUIPlugin;
import org.osgi.framework.BundleContext;


/**
 * @author Vitali Baumtrok
 */
public class Activator extends AbstractUIPlugin
{

	// Plugin id is not translatable.
	// Silence warning for "National Language Support" (NLS).
	public static final String PLUGIN_ID = "com.github.vitalibaumtrok.glsled"; //$NON-NLS-1$

	private static Activator plugin;

	@Override
	public void start ( final BundleContext context ) throws Exception
	{
		super.start ( context );

		plugin = this;
	}

	@Override
	public void stop ( final BundleContext context ) throws Exception
	{
		plugin = null;

		super.stop ( context );
	}

	/**
	 * Returns the shared instance
	 *
	 * @return the shared instance
	 */
	public static Activator getDefault ( )
	{
		return plugin;
	}

	/**
	 * Returns an image descriptor for the image file at the given
	 * plug-in relative path
	 *
	 * @param path
	 *            the path
	 * @return the image descriptor
	 */
	public static ImageDescriptor getImageDescriptor ( final String path )
	{
		return imageDescriptorFromPlugin ( PLUGIN_ID, path );
	}

}
