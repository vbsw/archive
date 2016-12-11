# Lovely Space

## Abstract
Lovely Space is a shoot 'em up game.

Lovely Space is created by Vitali Baumtrok <vbsoftware@mailbox.org> and published at <https://github.com/vitalibaumtrok/lovelyspace>.

## Copying
Lovely Space is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

Lovely Space is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program. If not, see <http://www.gnu.org/licenses/>.

## Controls

	f11       full screen
	escape    quit

## Development
The programming language Java and the framework JavaFX is used.

Dependencies:

	com.github.vitalibaumtrok.jsfw

## Using Eclipse
- Download Eclipse from <http://www.eclipse.org/downloads/>.
- Copy Lovely Space (or clone with Git) into workspace of Eclipse (for example: workspace/LovelySpace/).
- Create in Eclipse a "Java Project" with name "LovelySpace".
- Refresh the newly create project "LovelySpace".

## Using Git
How to install Git: <https://git-scm.com/book/en/v2/Getting-Started-Installing-Git>

Get the master branch and all refs of this project:

	$ git clone https://github.com/vitalibaumtrok/lovelyspace.git
	$ cd lovelyspace

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
