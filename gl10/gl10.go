
//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package gl10

// #cgo LDFLAGS: -lGL
// #include <GL/gl.h>
import "C"
import (
	"unsafe"
	//"reflect"
	//"os"
)

// Boolean values
const (
	FALSE = C.GL_FALSE
	TRUE = C.GL_TRUE
)

// Data types
const (
	BYTE = C.GL_BYTE
	UNSIGNED_BYTE = C.GL_UNSIGNED_BYTE
	SHORT = C.GL_SHORT
	UNSIGNED_SHORT = C.GL_UNSIGNED_SHORT
	INT = C.GL_INT
	UNSIGNED_INT = C.GL_UNSIGNED_INT
	FLOAT = C.GL_FLOAT
	TWO_BYTES = C.GL_2_BYTES
	THREE_BYTES = C.GL_3_BYTES
	FOUR_BYTES = C.GL_4_BYTES
	DOUBLE = C.GL_DOUBLE
)

// the Clear parameters
const (
	COLOR_BUFFER_BIT = C.GL_COLOR_BUFFER_BIT
	DEPTH_BUFFER_BIT = C.GL_DEPTH_BUFFER_BIT
	ACCUM_BUFFER_BIT = C.GL_ACCUM_BUFFER_BIT
	STENCIL_BUFFER_BIT = C.GL_STENCIL_BUFFER_BIT
)

// Primitives, the Begin parameters
const (
	POINTS = C.GL_POINTS
	LINES = C.GL_LINES
	LINE_LOOP = C.GL_LINE_LOOP
	LINE_STRIP = C.GL_LINE_STRIP
	TRIANGLES = C.GL_TRIANGLES
	TRIANGLE_STRIP = C.GL_TRIANGLE_STRIP
	TRIANGLE_FAN = C.GL_TRIANGLE_FAN
	QUADS = C.GL_QUADS
	QUAD_STRIP = C.GL_QUAD_STRIP
	POLYGON = C.GL_POLYGON
)

// matrix mode
const (
	MATRIX_MODE = C.GL_MATRIX_MODE
	MODELVIEW = C.GL_MODELVIEW
	PROJECTION = C.GL_PROJECTION
	TEXTURE = C.GL_TEXTURE
)

// blending
const (
	BLEND = C.GL_BLEND
	BLEND_SRC = C.GL_BLEND_SRC
	BLEND_DST = C.GL_BLEND_DST
	ZERO = C.GL_ZERO
	ONE = C.GL_ONE
	SRC_COLOR = C.GL_SRC_COLOR
	ONE_MINUS_SRC_COLOR = C.GL_ONE_MINUS_SRC_COLOR
	SRC_ALPHA = C.GL_SRC_ALPHA
	ONE_MINUS_SRC_ALPHA = C.GL_ONE_MINUS_SRC_ALPHA
	DST_ALPHA = C.GL_DST_ALPHA
	ONE_MINUS_DST_ALPHA = C.GL_ONE_MINUS_DST_ALPHA
	DST_COLOR = C.GL_DST_COLOR
	ONE_MINUS_DST_COLOR = C.GL_ONE_MINUS_DST_COLOR
	SRC_ALPHA_SATURATE = C.GL_SRC_ALPHA_SATURATE
)

// Alpha testing
const (
	ALPHA_TEST = C.GL_ALPHA_TEST
	ALPHA_TEST_REF = C.GL_ALPHA_TEST_REF
	ALPHA_TEST_FUNC = C.GL_ALPHA_TEST_FUNC
)

// Texture mapping
const (
	TEXTURE_ENV = C.GL_TEXTURE_ENV
	TEXTURE_ENV_MODE = C.GL_TEXTURE_ENV_MODE
	TEXTURE_1D = C.GL_TEXTURE_1D
	TEXTURE_2D = C.GL_TEXTURE_2D
	TEXTURE_WRAP_S = C.GL_TEXTURE_WRAP_S
	TEXTURE_WRAP_T = C.GL_TEXTURE_WRAP_T
	TEXTURE_MAG_FILTER = C.GL_TEXTURE_MAG_FILTER
	TEXTURE_MIN_FILTER = C.GL_TEXTURE_MIN_FILTER
	TEXTURE_ENV_COLOR = C.GL_TEXTURE_ENV_COLOR
	TEXTURE_GEN_S = C.GL_TEXTURE_GEN_S
	TEXTURE_GEN_T = C.GL_TEXTURE_GEN_T
	TEXTURE_GEN_R = C.GL_TEXTURE_GEN_R
	TEXTURE_GEN_Q = C.GL_TEXTURE_GEN_Q
	TEXTURE_GEN_MODE = C.GL_TEXTURE_GEN_MODE
	TEXTURE_BORDER_COLOR = C.GL_TEXTURE_BORDER_COLOR
	TEXTURE_WIDTH = C.GL_TEXTURE_WIDTH
	TEXTURE_HEIGHT = C.GL_TEXTURE_HEIGHT
	TEXTURE_BORDER = C.GL_TEXTURE_BORDER
	TEXTURE_COMPONENTS = C.GL_TEXTURE_COMPONENTS
	TEXTURE_RED_SIZE = C.GL_TEXTURE_RED_SIZE
	TEXTURE_GREEN_SIZE = C.GL_TEXTURE_GREEN_SIZE
	TEXTURE_BLUE_SIZE = C.GL_TEXTURE_BLUE_SIZE
	TEXTURE_ALPHA_SIZE = C.GL_TEXTURE_ALPHA_SIZE
	TEXTURE_LUMINANCE_SIZE = C.GL_TEXTURE_LUMINANCE_SIZE
	TEXTURE_INTENSITY_SIZE = C.GL_TEXTURE_INTENSITY_SIZE
	NEAREST_MIPMAP_NEAREST = C.GL_NEAREST_MIPMAP_NEAREST
	NEAREST_MIPMAP_LINEAR = C.GL_NEAREST_MIPMAP_LINEAR
	LINEAR_MIPMAP_NEAREST = C.GL_LINEAR_MIPMAP_NEAREST
	LINEAR_MIPMAP_LINEAR = C.GL_LINEAR_MIPMAP_LINEAR
	OBJECT_LINEAR = C.GL_OBJECT_LINEAR
	OBJECT_PLANE = C.GL_OBJECT_PLANE
	EYE_LINEAR = C.GL_EYE_LINEAR
	EYE_PLANE = C.GL_EYE_PLANE
	SPHERE_MAP = C.GL_SPHERE_MAP
	DECAL = C.GL_DECAL
	MODULATE = C.GL_MODULATE
	NEAREST = C.GL_NEAREST
	REPEAT = C.GL_REPEAT
	CLAMP = C.GL_CLAMP
	S = C.GL_S
	T = C.GL_T
	R = C.GL_R
	Q = C.GL_Q
)

// Stencil
const (
	STENCIL_BITS = C.GL_STENCIL_BITS
	STENCIL_TEST = C.GL_STENCIL_TEST
	STENCIL_CLEAR_VALUE = C.GL_STENCIL_CLEAR_VALUE
	STENCIL_FUNC = C.GL_STENCIL_FUNC
	STENCIL_VALUE_MASK = C.GL_STENCIL_VALUE_MASK
	STENCIL_FAIL = C.GL_STENCIL_FAIL
	STENCIL_PASS_DEPTH_FAIL = C.GL_STENCIL_PASS_DEPTH_FAIL
	STENCIL_PASS_DEPTH_PASS = C.GL_STENCIL_PASS_DEPTH_PASS
	STENCIL_REF = C.GL_STENCIL_REF
	STENCIL_WRITEMASK = C.GL_STENCIL_WRITEMASK
	STENCIL_INDEX = C.GL_STENCIL_INDEX
	KEEP = C.GL_KEEP
	REPLACE = C.GL_REPLACE
	INCR = C.GL_INCR
	DECR = C.GL_DECR
)

// Buffers, Pixel Drawing/Reading
const (
	NONE = C.GL_NONE
	LEFT = C.GL_LEFT
	RIGHT = C.GL_RIGHT
	FRONT_LEFT = C.GL_FRONT_LEFT
	FRONT_RIGHT = C.GL_FRONT_RIGHT
	BACK_LEFT = C.GL_BACK_LEFT
	BACK_RIGHT = C.GL_BACK_RIGHT
	AUX0 = C.GL_AUX0
	AUX1 = C.GL_AUX1
	AUX2 = C.GL_AUX2
	AUX3 = C.GL_AUX3
	COLOR_INDEX = C.GL_COLOR_INDEX
	RED = C.GL_RED
	GREEN = C.GL_GREEN
	BLUE = C.GL_BLUE
	ALPHA = C.GL_ALPHA
	LUMINANCE = C.GL_LUMINANCE
	LUMINANCE_ALPHA = C.GL_LUMINANCE_ALPHA
	ALPHA_BITS = C.GL_ALPHA_BITS
	RED_BITS = C.GL_RED_BITS
	GREEN_BITS = C.GL_GREEN_BITS
	BLUE_BITS = C.GL_BLUE_BITS
	INDEX_BITS = C.GL_INDEX_BITS
	SUBPIXEL_BITS = C.GL_SUBPIXEL_BITS
	AUX_BUFFERS = C.GL_AUX_BUFFERS
	READ_BUFFER = C.GL_READ_BUFFER
	DRAW_BUFFER = C.GL_DRAW_BUFFER
	DOUBLEBUFFER = C.GL_DOUBLEBUFFER
	STEREO = C.GL_STEREO
	BITMAP = C.GL_BITMAP
	COLOR = C.GL_COLOR
	DEPTH = C.GL_DEPTH
	STENCIL = C.GL_STENCIL
	DITHER = C.GL_DITHER
	RGB = C.GL_RGB
	RGBA = C.GL_RGBA
)


func ClearColor(red, green, blue, alpha float32) {
	C.glClearColor(C.GLclampf(red), C.GLclampf(green), C.GLclampf(blue), C.GLclampf(alpha))
}

func Clear(mask uint) {
	C.glClear(C.GLbitfield(mask))
}

func Begin(mode uint) {
	C.glBegin(C.GLenum(mode))
}

func End() {
	C.glEnd()
}

func Flush() {
	C.glFlush()
}

func LineWidth(width float32) {
	C.glLineWidth(C.GLfloat(width))
}

func Viewport(x, y, width, height int) {
	C.glViewport(C.GLint(x), C.GLint(y), C.GLsizei(width), C.GLsizei(height))
}

func MatrixMode(mode uint) {
	C.glMatrixMode(C.GLenum(mode))
}

func LoadIdentity() {
	C.glLoadIdentity()
}

func MultMatrixd(m []float64) {
	C.glMultMatrixd( (*C.GLdouble)(&m[0]) )
}

func MultMatrixf(m []float32) {
	C.glMultMatrixf( (*C.GLfloat)(&m[0]) )
}

func BlendFunc(sfactor, dfactor uint) {
	C.glBlendFunc(C.GLenum(sfactor), C.GLenum(dfactor))
}

func Enable(capability uint) {
	C.glEnable(C.GLenum(capability))
}

func TexImage1D(target uint, level, internalFormat, width, border int, format, pixelType uint, pixels []byte) {
	C.glTexImage1D( C.GLenum(target), C.GLint(level),
	                C.GLint(internalFormat),
	                C.GLsizei(width), C.GLint(border),
	                C.GLenum(format), C.GLenum(pixelType),
	                unsafe.Pointer(&pixels) )
}

// use unsafe.Pointer(&arr[0]) or similar for pixels
func TexImage2D(target uint, level, internalFormat, width, height, border int, format, pixelType uint, pixels unsafe.Pointer) {
	C.glTexImage2D( C.GLenum(target), C.GLint(level),
	                C.GLint(internalFormat),
	                C.GLsizei(width), C.GLsizei(height),
	                C.GLint(border), C.GLenum(format), C.GLenum(pixelType),
	                pixels )
}

// use unsafe.Pointer(&arr[0]) or similar for pixels
func GetTexImage(target uint, level int, format, pixelType uint, pixels unsafe.Pointer) {
	C.glGetTexImage( C.GLenum(target), C.GLint(level),
	                 C.GLenum(format), C.GLenum(pixelType),
	                 pixels )
}

func TexEnvf(target, pname uint, param float32) {
	C.glTexEnvf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}

func TexEnvi(target, pname uint, param int) {
	C.glTexEnvi(C.GLenum(target), C.GLenum(pname), C.GLint(param))
}

func TexParameterf(target, pname uint, param float32) {
	C.glTexParameterf(C.GLenum(target), C.GLenum(pname), C.GLfloat(param))
}

func TexParameteri(target, pname uint, param int) {
	C.glTexParameteri(C.GLenum(target), C.GLenum(pname), C.GLint(param))
}

func TexCoord1d(s float64) { C.glTexCoord1d(C.GLdouble(s)) }
func TexCoord1f(s float32) { C.glTexCoord1f(C.GLfloat(s)) }
func TexCoord1i(s int) { C.glTexCoord1i(C.GLint(s)) }
func TexCoord1s(s int16) { C.glTexCoord1s(C.GLshort(s)) }

func TexCoord2d(s, t float64) { C.glTexCoord2d(C.GLdouble(s), C.GLdouble(t)) }
func TexCoord2f(s, t float32) { C.glTexCoord2f(C.GLfloat(s), C.GLfloat(t)) }
func TexCoord2i(s, t int) { C.glTexCoord2i(C.GLint(s), C.GLint(t)) }
func TexCoord2s(s, t int16) { C.glTexCoord2s(C.GLshort(s), C.GLshort(t)) }

func TexCoord3d(s, t, r float64) { C.glTexCoord3d(C.GLdouble(s), C.GLdouble(t), C.GLdouble(r)) }
func TexCoord3f(s, t, r float32) { C.glTexCoord3f(C.GLfloat(s), C.GLfloat(t), C.GLfloat(r)) }
func TexCoord3i(s, t, r int) { C.glTexCoord3i(C.GLint(s), C.GLint(t), C.GLint(r)) }
func TexCoord3s(s, t, r int16) { C.glTexCoord3s(C.GLshort(s), C.GLshort(t), C.GLshort(r)) }

func TexCoord4d(s, t, r, q float64) { C.glTexCoord4d(C.GLdouble(s), C.GLdouble(t), C.GLdouble(r), C.GLdouble(q)) }
func TexCoord4f(s, t, r, q float32) { C.glTexCoord4f(C.GLfloat(s), C.GLfloat(t), C.GLfloat(r), C.GLfloat(q)) }
func TexCoord4i(s, t, r, q int) { C.glTexCoord4i(C.GLint(s), C.GLint(t), C.GLint(r), C.GLint(q)) }
func TexCoord4s(s, t, r, q int16) { C.glTexCoord4s(C.GLshort(s), C.GLshort(t), C.GLshort(r), C.GLshort(q)) }

/*
GLAPI void GLAPIENTRY glTexCoord1dv( const GLdouble *v );
GLAPI void GLAPIENTRY glTexCoord1fv( const GLfloat *v );
GLAPI void GLAPIENTRY glTexCoord1iv( const GLint *v );
GLAPI void GLAPIENTRY glTexCoord1sv( const GLshort *v );

GLAPI void GLAPIENTRY glTexCoord2dv( const GLdouble *v );
GLAPI void GLAPIENTRY glTexCoord2fv( const GLfloat *v );
GLAPI void GLAPIENTRY glTexCoord2iv( const GLint *v );
GLAPI void GLAPIENTRY glTexCoord2sv( const GLshort *v );

GLAPI void GLAPIENTRY glTexCoord3dv( const GLdouble *v );
GLAPI void GLAPIENTRY glTexCoord3fv( const GLfloat *v );
GLAPI void GLAPIENTRY glTexCoord3iv( const GLint *v );
GLAPI void GLAPIENTRY glTexCoord3sv( const GLshort *v );

GLAPI void GLAPIENTRY glTexCoord4dv( const GLdouble *v );
GLAPI void GLAPIENTRY glTexCoord4fv( const GLfloat *v );
GLAPI void GLAPIENTRY glTexCoord4iv( const GLint *v );
GLAPI void GLAPIENTRY glTexCoord4sv( const GLshort *v );
*/

func Vertex2d(x, y float64) { C.glVertex2d(C.GLdouble(x), C.GLdouble(y)) }
func Vertex2f(x, y float32) { C.glVertex2f(C.GLfloat(x), C.GLfloat(y)) }
func Vertex2i(x, y int) { C.glVertex2i(C.GLint(x), C.GLint(y)) }
func Vertex2s(x, y int16) { C.glVertex2s(C.GLshort(x), C.GLshort(y)) }

func Vertex3d(x, y, z float64) { C.glVertex3d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z)) }
func Vertex3f(x, y, z float32) { C.glVertex3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z)) }
func Vertex3i(x, y, z int) { C.glVertex3i(C.GLint(x), C.GLint(y), C.GLint(z)) }
func Vertex3s(x, y, z int16) { C.glVertex3s(C.GLshort(x), C.GLshort(y), C.GLshort(z)) }

func Vertex4d(x, y, z, w float64) { C.glVertex4d(C.GLdouble(x), C.GLdouble(y), C.GLdouble(z), C.GLdouble(w)) }
func Vertex4f(x, y, z, w float32) { C.glVertex4f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z), C.GLfloat(w)) }
func Vertex4i(x, y, z, w int) { C.glVertex4i(C.GLint(x), C.GLint(y), C.GLint(z), C.GLint(w)) }
func Vertex4s(x, y, z, w int16) { C.glVertex4s(C.GLshort(x), C.GLshort(y), C.GLshort(z), C.GLshort(w)) }

/*
GLAPI void GLAPIENTRY glVertex2dv( const GLdouble *v );
GLAPI void GLAPIENTRY glVertex2fv( const GLfloat *v );
GLAPI void GLAPIENTRY glVertex2iv( const GLint *v );
GLAPI void GLAPIENTRY glVertex2sv( const GLshort *v );

GLAPI void GLAPIENTRY glVertex3dv( const GLdouble *v );
GLAPI void GLAPIENTRY glVertex3fv( const GLfloat *v );
GLAPI void GLAPIENTRY glVertex3iv( const GLint *v );
GLAPI void GLAPIENTRY glVertex3sv( const GLshort *v );

GLAPI void GLAPIENTRY glVertex4dv( const GLdouble *v );
GLAPI void GLAPIENTRY glVertex4fv( const GLfloat *v );
GLAPI void GLAPIENTRY glVertex4iv( const GLint *v );
GLAPI void GLAPIENTRY glVertex4sv( const GLshort *v );




GLAPI void GLAPIENTRY glNormal3b( GLbyte nx, GLbyte ny, GLbyte nz );
GLAPI void GLAPIENTRY glNormal3d( GLdouble nx, GLdouble ny, GLdouble nz );
GLAPI void GLAPIENTRY glNormal3f( GLfloat nx, GLfloat ny, GLfloat nz );
GLAPI void GLAPIENTRY glNormal3i( GLint nx, GLint ny, GLint nz );
GLAPI void GLAPIENTRY glNormal3s( GLshort nx, GLshort ny, GLshort nz );

GLAPI void GLAPIENTRY glNormal3bv( const GLbyte *v );
GLAPI void GLAPIENTRY glNormal3dv( const GLdouble *v );
GLAPI void GLAPIENTRY glNormal3fv( const GLfloat *v );
GLAPI void GLAPIENTRY glNormal3iv( const GLint *v );
GLAPI void GLAPIENTRY glNormal3sv( const GLshort *v );


GLAPI void GLAPIENTRY glIndexd( GLdouble c );
GLAPI void GLAPIENTRY glIndexf( GLfloat c );
GLAPI void GLAPIENTRY glIndexi( GLint c );
GLAPI void GLAPIENTRY glIndexs( GLshort c );
GLAPI void GLAPIENTRY glIndexub( GLubyte c );

GLAPI void GLAPIENTRY glIndexdv( const GLdouble *c );
GLAPI void GLAPIENTRY glIndexfv( const GLfloat *c );
GLAPI void GLAPIENTRY glIndexiv( const GLint *c );
GLAPI void GLAPIENTRY glIndexsv( const GLshort *c );
GLAPI void GLAPIENTRY glIndexubv( const GLubyte *c );
*/


func Color3d(red, green, blue float64) { C.glColor3d(C.GLdouble(red), C.GLdouble(green), C.GLdouble(blue)) }
func Color3f(red, green, blue float32) { C.glColor3f(C.GLfloat(red), C.GLfloat(green), C.GLfloat(blue)) }
func Color3b(red, green, blue int8) { C.glColor3b(C.GLbyte(red), C.GLbyte(green), C.GLbyte(blue)) }
func Color3s(red, green, blue int16) { C.glColor3s(C.GLshort(red), C.GLshort(green), C.GLshort(blue)) }
func Color3i(red, green, blue int) { C.glColor3i(C.GLint(red), C.GLint(green), C.GLint(blue)) }
func Color3ub(red, green, blue uint8) { C.glColor3ub(C.GLubyte(red), C.GLubyte(green), C.GLubyte(blue)) }
func Color3us(red, green, blue uint16) { C.glColor3us(C.GLushort(red), C.GLushort(green), C.GLushort(blue)) }
func Color3ui(red, green, blue uint) { C.glColor3ui(C.GLuint(red), C.GLuint(green), C.GLuint(blue)) }

/*
GLAPI void GLAPIENTRY glColor4b( GLbyte red, GLbyte green,
                                   GLbyte blue, GLbyte alpha );
GLAPI void GLAPIENTRY glColor4d( GLdouble red, GLdouble green,
                                   GLdouble blue, GLdouble alpha );
GLAPI void GLAPIENTRY glColor4f( GLfloat red, GLfloat green,
                                   GLfloat blue, GLfloat alpha );
GLAPI void GLAPIENTRY glColor4i( GLint red, GLint green,
                                   GLint blue, GLint alpha );
GLAPI void GLAPIENTRY glColor4s( GLshort red, GLshort green,
                                   GLshort blue, GLshort alpha );
GLAPI void GLAPIENTRY glColor4ub( GLubyte red, GLubyte green,
                                    GLubyte blue, GLubyte alpha );
GLAPI void GLAPIENTRY glColor4ui( GLuint red, GLuint green,
                                    GLuint blue, GLuint alpha );
GLAPI void GLAPIENTRY glColor4us( GLushort red, GLushort green,
                                    GLushort blue, GLushort alpha );


GLAPI void GLAPIENTRY glColor3bv( const GLbyte *v );
GLAPI void GLAPIENTRY glColor3dv( const GLdouble *v );
GLAPI void GLAPIENTRY glColor3fv( const GLfloat *v );
GLAPI void GLAPIENTRY glColor3iv( const GLint *v );
GLAPI void GLAPIENTRY glColor3sv( const GLshort *v );
GLAPI void GLAPIENTRY glColor3ubv( const GLubyte *v );
GLAPI void GLAPIENTRY glColor3uiv( const GLuint *v );
GLAPI void GLAPIENTRY glColor3usv( const GLushort *v );

GLAPI void GLAPIENTRY glColor4bv( const GLbyte *v );
GLAPI void GLAPIENTRY glColor4dv( const GLdouble *v );
GLAPI void GLAPIENTRY glColor4fv( const GLfloat *v );
GLAPI void GLAPIENTRY glColor4iv( const GLint *v );
GLAPI void GLAPIENTRY glColor4sv( const GLshort *v );
GLAPI void GLAPIENTRY glColor4ubv( const GLubyte *v );
GLAPI void GLAPIENTRY glColor4uiv( const GLuint *v );
GLAPI void GLAPIENTRY glColor4usv( const GLushort *v );
*/














