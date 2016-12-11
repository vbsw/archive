/*
 * This is free and unencumbered software released into the public domain.
 * 
 * Anyone is free to copy, modify, publish, use, compile, sell, or
 * distribute this software, either in source code form or as a compiled
 * binary, for any purpose, commercial or non-commercial, and by any
 * means.
 * 
 * In jurisdictions that recognize copyright laws, the author or authors
 * of this software dedicate any and all copyright interest in the
 * software to the public domain. We make this dedication for the benefit
 * of the public at large and to the detriment of our heirs and
 * successors. We intend this dedication to be an overt act of
 * relinquishment in perpetuity of all present and future rights to this
 * software under copyright law.
 * 
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 * IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
 * OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
 * ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
 * OTHER DEALINGS IN THE SOFTWARE.
 * 
 * For more information, please refer to <http://unlicense.org>
 */


#include <stdio.h>

#include <GL/glut.h>


const int INITIAL_WINDOW_WIDTH = 300;
const int INITIAL_WINDOW_HEIGHT = 300;

int mousestartx = 0;
int mousestarty = 0;

int clientoffsetsinitialized = 0;
int clientprevx = -1;
int clientprevy = -1;
int dragging = 0;
int clientoffsetx = 0;
int clientoffsety = 0;


void displayFunc () {
	glClear ( GL_COLOR_BUFFER_BIT );
}


void keyboardFunc ( unsigned char key, int x, int y ) {
	if ( key==27 ) { // escape key
		glutDestroyWindow ( glutGetWindow () );
	} else if ( key=='o' ) {
		glutReshapeWindow ( INITIAL_WINDOW_WIDTH, INITIAL_WINDOW_HEIGHT );
		glutPostRedisplay ();
	}
}


void mouseFunc ( int button, int state, int x, int y ) {
	if ( button==GLUT_LEFT_BUTTON ) {
		if ( state==GLUT_DOWN ) {
			dragging = 1;
			mousestartx = x;
			mousestarty = y;
		} else {
			dragging = 0;
		}
	}
}


void motionFunc ( int x, int y ) {
	if ( dragging ) {
		const int clientx = glutGet ( GLUT_WINDOW_X );
		const int clienty = glutGet ( GLUT_WINDOW_Y );
		const int mousedx = x-mousestartx;
		const int mousedy = y-mousestarty;
		
		if ( clientoffsetsinitialized ) {
			glutPositionWindow ( clientx+mousedx-clientoffsetx, clienty+mousedy-clientoffsety );

		} else if ( dragging==1 ) {
			dragging = 2;
			clientprevx = clientx+mousedx;
			clientprevy = clienty+mousedy;
			glutPositionWindow ( clientprevx, clientprevy );

		} else {
			clientoffsetx = clientx-clientprevx;
			clientoffsety = clienty-clientprevy;
			glutPositionWindow ( clientx+mousedx-clientoffsetx, clienty+mousedy-clientoffsety );
			clientoffsetsinitialized = 1;
		}
	}
}


void main ( int argc, char **argv ) {
	glutInit ( &argc, argv );
	glutInitDisplayMode ( GLUT_SINGLE | GLUT_RGBA );
	glutInitWindowSize ( INITIAL_WINDOW_WIDTH, INITIAL_WINDOW_HEIGHT );
	glutCreateWindow ( "Test OpenGL: winmod" );
	glutDisplayFunc ( &displayFunc );
	glutKeyboardFunc ( &keyboardFunc );
	glutMouseFunc ( &mouseFunc );
	glutMotionFunc ( &motionFunc );

	glClearColor ( 0, 0, 0, 0 );
	glBlendFunc ( GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA );
	glEnable ( GL_BLEND );
	glutPositionWindow ( 0, 0 );

	glutMainLoop ();
}





































