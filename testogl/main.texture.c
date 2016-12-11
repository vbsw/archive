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

#define TEXTURE_WIDTH 128
#define TEXTURE_HEIGHT 128
#define TEXTURE_PIXEL_COUNT TEXTURE_WIDTH*TEXTURE_HEIGHT
#define SQUARE_SIZE 16
GLuint texture[TEXTURE_PIXEL_COUNT];
GLuint textureid;


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

	glBindTexture ( GL_TEXTURE_2D, textureid );
	glBegin ( GL_QUADS );
		glTexCoord2f ( 0.0f, 0.0f ); glVertex3f ( -0.5f, 0.5f, 0.0f );
		glTexCoord2f ( 1.0f, 0.0f ); glVertex3f ( 0.5f, 0.5f, 0.0f );
		glTexCoord2f ( 1.0f, 1.0f ); glVertex3f ( 0.5f, -0.5f, 0.0f );
		glTexCoord2f ( 0.0f, 1.0f ); glVertex3f ( -0.5f, -0.5f, 0.0f );
	glEnd ();
	glBindTexture ( GL_TEXTURE_2D, 0 );

	glutSwapBuffers ();
}


void keyboardFunc ( unsigned char key, int x, int y ) {
	if (key==27)
		glutDestroyWindow ( glutGetWindow () );
}


void main ( int argc, char **argv ) {
	glutInit ( &argc, argv );
	glutInitDisplayMode ( GLUT_DOUBLE | GLUT_RGBA );
	glutInitWindowSize ( INITIAL_WINDOW_WIDTH, INITIAL_WINDOW_HEIGHT );
	glutCreateWindow ( "Test OpenGL: texture" );
	glutDisplayFunc ( &displayFunc );
	glutKeyboardFunc ( &keyboardFunc );

	glClearColor ( 0, 0, 0, 0 );
	glBlendFunc ( GL_SRC_ALPHA, GL_ONE_MINUS_SRC_ALPHA );
	glEnable ( GL_BLEND );
	glEnable ( GL_TEXTURE_2D );

	// init texture
	int i;
	for (i=0; i<TEXTURE_PIXEL_COUNT; ++i) {
		GLubyte *const colors = (GLubyte*) &texture[i];

		if ( ((i/TEXTURE_WIDTH) & SQUARE_SIZE) ^ (i & SQUARE_SIZE) ) {
			colors[0] = 0xFF;
			colors[1] = 0xFF;
			colors[2] = 0xFF;
			colors[3] = 0x00;

		} else {
			colors[0] = 0x00;
			colors[1] = 0xFF;
			colors[2] = 0x00;
			colors[3] = 0xFF;
		}
	}
	glGenTextures ( 1, &textureid );
	glBindTexture ( GL_TEXTURE_2D, textureid );
	glTexImage2D ( GL_TEXTURE_2D, 0, GL_RGBA, TEXTURE_WIDTH, TEXTURE_HEIGHT, 0, GL_RGBA, GL_UNSIGNED_BYTE, texture );
	glTexParameteri ( GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_NEAREST );
	glTexParameteri ( GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_NEAREST );
	glTexEnvf ( GL_TEXTURE_ENV, GL_TEXTURE_ENV_MODE, GL_REPLACE );
	glBindTexture ( GL_TEXTURE_2D, 0 );

	glutMainLoop ();
}





































