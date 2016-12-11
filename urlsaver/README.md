# URL Saver

## About
URL Saver is a program to manage URLs by keywords.

URL Saver is created by Vitali Baumtrok <vbit@mailbox.org> and published at <https://gitlab.com/vbit/urlsaver>.

## Copying
URL Saver is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

URL Saver is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

## Compiling
Download Eclipse: <https://www.eclipse.org>

Clone the repository into your Eclipse workspace

	$ git clone https://gitlab.com/vbit/urlsaver.git URLSaver

Open Eclipse and create a Java Project named URLSaver (like previously created directory). Right click on project, then "export...", then export as "Runnable JAR file".

## Using Git
How to install Git: <https://git-scm.com/book/en/v2/Getting-Started-Installing-Git>

Get the master branch and all refs of this project:

	$ git clone https://gitlab.com/vbit/urlsaver.git
	$ cd window

See all tags:

	$ git tag -l

See local and remote branches:

	$ git branch -a

Checkout other branches than master, for example the development branch:

	$ git branch development origin/development
	$ git checkout development

See tracked remote branches:

	$ git branch -vv

Update all tracked branches and all refs:

	$ git fetch
