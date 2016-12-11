# GLSL Ed

## Abstract
GLSL Ed is an editor for the OpenGL Shading Language. (It is incomplete.)

## Copyright
Copyright 2016, Vitali Baumtrok.

GLSL Ed is distributed under the terms of the Eclipse Public License, Version 1.0. (See accompanying file LICENSE or copy at https://www.eclipse.org/legal/epl-v10.html)

## Features

- None

## Installation
GLSL Ed is installed as a plugin for the Eclipse IDE.

Clone the repository into your Eclipse workspace

	$ git clone https://github.com/vitalibaumtrok/glsled.git GLSLEd

Open Eclipse and create a Plug-in Project named GLSLEd. Uncheck all options on second page
and uncheck "Create a plug-in using one of the templates" on the third page.

Overwrite the newly created files with files from repository.

	$ cd GLSLEd
	$ git checkout master -f

Now you can export the plugin. Right click on project, then select "Export > Plug-in Development > Deployable plug-ins and fragments".
Exporting to a directory will create a jar file. Drop this jar file into your Eclipse plugin directory. Done!

## Using Git
How to install Git: <https://git-scm.com/book/en/v2/Getting-Started-Installing-Git>

Get the master branch and all refs of this project:

	$ git clone https://github.com/vitalibaumtrok/glsled.git
	$ cd glsled

See local and remote branches:

	$ git branch -a

Check out other branches than master, for example the development branch:

	$ git branch development origin/development
	$ git checkout development

See tracked remote branches:

	$ git branch -vv

Update all tracked branches and all refs:

	$ git fetch

## Development

### References

- [Vogella Eclipse Tutorials](http://www.vogella.com/tutorials/eclipse.html)
- [Building an Eclipse Text Editor with JFace Text](http://www.realsolve.co.uk/site/tech/jface-text.php)
- [Create a commercial-quality Eclipse IDE](http://www.ibm.com/developerworks/views/opensource/libraryview.jsp?search_by=Create+commercial-quality+eclipse+ide)

### Tools

To install source code for

	org.eclipse.ui.editors.text

in Eclipse do

1. open in menu "Help" > "Install New Software...",
2. select "The Ecliipse Project Updates",
3. select "Eclipse SDK" and
4. install.
