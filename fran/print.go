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


package main

import "fmt"




func printVersion() {
	fmt.Println("fran (version 0.3.0)")
	fmt.Println("This is free software; see the source for copying conditions. There is NO")
	fmt.Println("warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.")
	fmt.Println()
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("    fran help")
	fmt.Println("    fran version")
	fmt.Println("    fran controls")
	fmt.Println("    fran examples")
	fmt.Println("    fran FILE WxH[xC] DELAY [REPLAYDELAY]")
	fmt.Println()
	fmt.Println("Arguments:")
	fmt.Println()
	fmt.Println("    FILE         path to a PNG file")
	fmt.Println("    W            number of images in horizontal")
	fmt.Println("    H            number of images in vertical")
	fmt.Println("    C            number of images in total")
	fmt.Println("    DELAY        delay between images in milliseconds")
	fmt.Println("    REPLAYDELAY  delay between loops in milliseconds")
	fmt.Println()
}

func printControls() {
	fmt.Println("Keys to control animation:")
	fmt.Println()
	fmt.Println("    o           set window to fit original size of image")
	fmt.Println("    p           pause/play animation")
	fmt.Println("    r           reload image")
	fmt.Println("    +           zoom in")
	fmt.Println("    -           zoom out")
	fmt.Println("    left        pause animation and show previous frame")
	fmt.Println("    right       pause animation and show next frame")
	fmt.Println("    up          increase the delay between images by 1 millisecond")
	fmt.Println("    down        decrease the delay between images by 1 millisecond")
	fmt.Println("    ctrl+arrow  like left/right/up/down, but by 10")
	fmt.Println()
}

func printExamples() {
	fmt.Println("Examples:")
	fmt.Println()
	fmt.Println("   fran help               prints help")
	fmt.Println("   fran img.png 8x8 30     plays animation stored in img.png")
	fmt.Println("   fran img.png 8x8x60 30  plays animation consisting of 60 frames")
	fmt.Println()
}

func printStartUpError() {
	fmt.Println("Error: Unknown or not enough arguments, or wrong format.")
	fmt.Println("Use \"fran help\" for more information.")
	fmt.Println()
}

func printAnimationNotImplemented() {
	fmt.Println("Animation not implemented, yet.")
	fmt.Println()
}

func printCanNotInterpretImage() {
	fmt.Println("Error: Couldn't decode image")
}

func printErrorCouldNotLoadFile() {
	fmt.Printf("Error: Couldn't load file \"%s\".\n", argFile)
	fmt.Println()
}

func printNotImplemented(funcName string) {
	fmt.Printf("Error: Function \"%s\" not implemented.", funcName)
	fmt.Println()
}

func printFrameDelay() {
	fmt.Println("New frame delay:", frameDelay)
}

