
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//      (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


// VIS (Visual Interface System) provides abstraction for graphic output.
package vis

import (
	"github.com/vbsw/events"
)

// Destroy() returns nil.
type VisualSystem interface {
	EventQueue ( ) EventQueue
	Window ( ) Window
	Graphics ( ) Graphics
	Destroy ( ) VisualSystem
}

type EventQueue interface {
	PollEvents ( ) bool
	EventCode ( ) events.Event
	KeyCode ( ) int
	EnableAutoRedraw ( )
	DisableAutoRedraw ( )
}

// Destroy() returns nil.
type Window interface {
	Show ( )
	WidthInPixels ( ) int
	HeightInPixels ( ) int
	SetWidthHeightInPixels ( width, height int )
	Title ( ) string
	SetTitle ( title string )
}

type Graphics interface {
	NewRectangle ( ) Rectangle
	NewTexture ( rgba []byte ) Texture
	SwapBuffers ( )
}

// Destroy() returns nil.
type Rectangle interface {
	X ( ) float32
	Y ( ) float32
	Z ( ) float32
	SetXYZ ( x, y, z float32 )
	Width ( ) float32
	Height ( ) float32
	SetWidthHeight ( width, height float32 )
	NewTextureFill ( texture Texture ) RectangleTextureFill
	NewColorFill ( r, g, b, a float32 ) ColorFill
	Destroy ( ) Rectangle
}

// Destroy() returns nil.
type RectangleTextureFill interface {
	X ( ) float32
	Y ( ) float32
	SetXY ( x, y float32 )
	Width ( ) float32
	Height ( ) float32
	SetWidthHeight ( width, height float32 )
	Destroy ( ) RectangleTextureFill
}

// Destroy() returns nil.
type ColorFill interface {
	R ( ) float32
	G ( ) float32
	B ( ) float32
	A ( ) float32
	SetRGBA ( r, g, b, a float32 )
	Destroy ( ) ColorFill
}

// Destroy() returns nil.
type Texture interface {
	WidthInPixels ( ) int
	HeightInPixels ( ) int
	Destroy ( ) Texture
}

