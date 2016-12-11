
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


// VIS (Visual Interface System) provides abstraction for graphic output. This is version 0.1.0.
package vis

import "gitlab.com/vbit/semver"

const (
	EVENT_NONE = 0
	EVENT_KEY_PRESS = 1
	EVENT_KEY_RELEASE = 2
	EVENT_BUTTON1_PRESS = 3
	EVENT_BUTTON1_RELEASE = 4
	EVENT_BUTTON2_PRESS = 5
	EVENT_BUTTON2_RELEASE = 6
	EVENT_BUTTON3_PRESS = 7
	EVENT_BUTTON3_RELEASE = 8
	EVENT_WHEEL_SPIN = 9
	EVENT_CURSOR_MOVE = 10
	EVENT_REDRAW = 11
	EVENT_WINDOW_CLOSE = 12
)

// Version is 0.1.0.
var Version semver.IVersion
var Queue IQueue
var Graphics IGraphics
var Window IWindow

func init ( ) {
    Version = semver.New(0,1,0)
}

func Destroy ( ) {
	Queue = Queue.Destroy()
	Graphics = Graphics.Destroy()
	Window = Window.Destroy()
}

// Destroy() returns nil.
type IWindow interface {
	Show ( )
	DimensionsInPixels ( ) (float32,float32)
	SetDimensionsInPixels ( width, height float32 )
	Title ( ) string
	SetTitle ( title string )
	Destroy ( ) IWindow
}

// Destroy() returns nil.
type IGraphics interface {
	NewRectangle ( ) IRectangle
	NewTexture ( rgba []byte ) ITexture
	SwapBuffers ( )
	Destroy ( ) IGraphics
}

// Destroy() returns nil.
type IQueue interface {
	Empty ( ) bool
	PushEvent ( eventType int, value1, value2 float32 )
	PopEvent ( ) (int,float32,float32)
	Destroy ( ) IQueue
}

// Destroy() returns nil.
type IRectangle interface {
	Location ( ) (float32,float32,float32)
	SetLocation ( x, y, z float32 )
	Dimensions ( ) (float32,float32)
	SetDimensions ( width, height float32 )
	NewTextureFill ( texture ITexture ) IRectangleTextureFill
	NewColorFill ( r, g, b, a float32 ) IColorFill
	Destroy ( ) IRectangle
}

// Destroy() returns nil.
type IRectangleTextureFill interface {
	Location ( ) (float32,float32)
	SetLocation ( x, y float32 )
	Dimensions ( ) (float32,float32)
	SetDimensions ( width, height float32 )
	Destroy ( ) IRectangleTextureFill
}

// Destroy() returns nil.
type IColorFill interface {
	Color ( ) (float32,float32,float32,float32)
	SetColor ( r, g, b, a float32 )
	Destroy ( ) IColorFill
}

// Destroy() returns nil.
type ITexture interface {
	DimensionsInPixels ( ) (float32,float32)
	Destroy ( ) ITexture
}

