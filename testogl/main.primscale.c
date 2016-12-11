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
const float INITIAL_WINDOW_RATIO = 300.0f / 300.0f;
float m[16] = { 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1 };


void reshapeFunc ( int width, int height ) {
	const float ratio = (float)width / (float)height;
	// broader new window
	if ( INITIAL_WINDOW_RATIO<ratio ) {
		m[0] = INITIAL_WINDOW_RATIO/ratio;
		m[5] = 1;
	// narrower new window
	} else if ( INITIAL_WINDOW_RATIO>ratio ) {
		m[0] = 1;
		m[5] = ratio/INITIAL_WINDOW_RATIO;
	} else {
		m[0] = 1;
		m[5] = 1;
	}

	glViewport ( 0, 0, width, height );
	glLoadIdentity ();
	glMultMatrixf ( m );
}


void displayFunc () {
	glClear ( GL_COLOR_BUFFER_BIT );
	glColor4f ( 1, 0, 0, 1 );
	glBegin ( GL_TRIANGLES );
		glVertex3f ( 1.0f, 1.0f, 0.0f );
		glVertex3f ( -1.0f, -0.5f, 0.0f );
		glVertex3f ( -0.5f, -1.0f, 0.0f );
	glEnd ();
	glColor4f ( 0, 0, 1, 1 );
	glBegin ( GL_TRIANGLES );
		glVertex3f ( 1.0f, -1.0f, 0.0f );
		glVertex3f ( -0.5f, 1.0f, 0.0f );
		glVertex3f ( -1.0f, 0.5f, 0.0f );
	glEnd ();

	glutSwapBuffers ();
}


void keyboardFunc ( unsigned char key, int x, int y ) {
	if ( key==27 ) // escape key
		glutDestroyWindow ( glutGetWindow () );
}


void main ( int argc, char **argv ) {
	glutInit ( &argc, argv );
	glutInitDisplayMode ( GLUT_DOUBLE | GLUT_RGBA );
	glutInitWindowSize ( INITIAL_WINDOW_WIDTH, INITIAL_WINDOW_HEIGHT );
	glutCreateWindow ( "Test OpenGL: primscale" );
	glutReshapeFunc ( &reshapeFunc );
	glutDisplayFunc ( &displayFunc );
	glutKeyboardFunc ( &keyboardFunc );

	glClearColor ( 0, 0, 0, 0 );
	glBlendFunc ( GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA );
	glEnable ( GL_BLEND );

	glutMainLoop ();
}





































