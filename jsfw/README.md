# JSFW

## Abstract
JSFW (JavaFX Shoot 'em up FrameWork) is a framework to ease the creation of shoot 'em up games.

JSFW is created by Vitali Baumtrok <vbsoftware@mailbox.org> and published at <https://github.com/vitalibaumtrok/jsfw>.

## Copying
JSFW is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

JSFW is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program. If not, see <http://www.gnu.org/licenses/>.

## Development
The programming language Java and the framework JavaFX is used.

## Using Eclipse
- Download Eclipse from <http://www.eclipse.org/downloads/>.
- Copy JSFW (or clone with Git) into workspace of Eclipse (for example: workspace/jsfw/).
- Create in Eclipse a "Java Project" with name "jsfw".
- Refresh the newly create project "jsfw".

## Using Git
How to install Git: <https://git-scm.com/book/en/v2/Getting-Started-Installing-Git>

Get the master branch and all refs of this project:

	$ git clone https://github.com/vitalibaumtrok/jsfw.git
	$ cd jsfw

See all tags:

	$ git tag -l

See local and remote branches:

	$ git branch -a

Check out other branches than master, for example the development branch:

	$ git branch development origin/development
	$ git checkout development

See tracked remote branches:

	$ git branch -vv

Update all tracked branches and all refs:

	$ git fetch
