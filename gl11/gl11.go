
//          Copyright Vitali Baumtrok 2014.
// Distributed under the Boost Software License, Version 1.0.
//    (See accompanying file LICENSE_1_0.txt or copy at
//          http://www.boost.org/LICENSE_1_0.txt)


package gl11

// #cgo LDFLAGS: -lGL
// #include <GL/gl.h>
import "C"
import (
	//"unsafe"
	//"os"
)

func GenTextures(size int, textures []uint) {
	cTex := make([]C.GLuint, size)
	C.glGenTextures(C.GLsizei(size), &cTex[0])
	for i := range textures {
		textures[i] = uint(cTex[i])
	}
}

func DeleteTextures(size int, textures []uint) {
	cTex := make([]C.GLuint, size)
	for i, tex := range textures {
		cTex[i] = C.GLuint(tex)
	}
	C.glDeleteTextures(C.GLsizei(size), &cTex[0])
}

func BindTexture(target, texture uint) {
	C.glBindTexture(C.GLenum(target), C.GLuint(texture))
}














