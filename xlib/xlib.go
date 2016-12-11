
//          Copyright 2016, Vitali Baumtrok.
// Distributed under the Boost Software License, Version 1.0.
//     (See accompanying file LICENSE or copy at
//        http://www.boost.org/LICENSE_1_0.txt)


package xlib


// #cgo LDFLAGS: -lX11
// #include <stdlib.h>
// #include <X11/Xlib.h>
// #include "gotoc.h"
import "C"
import (
	"unsafe"
)


type Display C.Display
type Colormap C.Colormap
type GC C.GC
type Window C.Window
type Screen C.Screen
type Visual C.Visual
type Pixmap C.Pixmap
type Cursor C.Cursor
type VisualID C.VisualID

type XPixmapFormatValues struct {
	Depth int
	BitsPerPixel int
	ScanlinePad int
}

type SetWindowAttributes struct {
	BackgroundPixmap Pixmap
	BackgroundPixel uint64
	BorderPixmap Pixmap
	BorderPixel uint64
	BitGravity int
	WinGravity int
	BackingStore int
	BackingPlanes uint64
	BackingPixel uint64
	SaveUnder bool
	EventMask int64
	DoNotPropagateMask int64
	OverrideRedirect bool
	WinColormap Colormap
	WinCursor Cursor
}

func concat(a []interface{}) string {
	str := ""

	for _, strPart := range a {
		switch s := strPart.(type) {

		case string:
			str += s
		}
	}

	return str
}

func OpenDisplay(displayName ...interface{}) *Display {
	if len(displayName) == 0 {
		display := C.XOpenDisplay(nil)

		return (*Display)(display)

	} else {
		displName := concat(displayName)

		if len(displName) > 0 {
			displNameCStr := C.CString(displName)
			display := C.XOpenDisplay(displNameCStr)

			C.free(unsafe.Pointer(displNameCStr))

			return (*Display)(display)

		} else {
			return nil
		}
	}
}

func AllPlanes() uint64 {
	allPlanes := C.XAllPlanes()

	return uint64(allPlanes)
}

func (this *Display) BlackPixel(screenNumber int) uint {
	displayC := (*C.Display)(this)
	pixel := C.XBlackPixel(displayC, C.int(screenNumber))

	return uint(pixel)
}

func (this *Display) WhitePixel(screenNumber int) uint {
	displayC := (*C.Display)(this)
	pixel := C.XWhitePixel(displayC, C.int(screenNumber))

	return uint(pixel)
}

func (this *Display) ConnectionNumber() int {
	displayC := (*C.Display)(this)
	connectionNumber := C.XConnectionNumber(displayC)

	return int(connectionNumber)
}

func (this *Display) DefaultColormap(screenNumber int) Colormap {
	displayC := (*C.Display)(this)
	colorMap := C.XDefaultColormap(displayC, C.int(screenNumber))

	return Colormap(colorMap)
}

func (this *Display) DefaultDepth(screenNumber int) int {
	displayC := (*C.Display)(this)
	defaultRootWindowDepth := C.XDefaultDepth(displayC, C.int(screenNumber))

	return int(defaultRootWindowDepth)
}

func (this *Display) ListDepths(screenNumber int) []int {
	var listDepths []int
	displayC := (*C.Display)(this)
	countReturn := C.int(0)
	listDepthsC := C.XListDepths(displayC, C.int(screenNumber), &countReturn)

	if countReturn > 0 {
		listDepths = make([]int, int(countReturn))
		listDepthsPtrInt := uintptr(unsafe.Pointer(listDepthsC))

		for i := range listDepths {
			offsetInt := uintptr(C.sizeof_int * i)
			listDepthsPtrIndex := listDepthsPtrInt + offsetInt
			depth := (*C.int)(unsafe.Pointer(listDepthsPtrIndex))
			listDepths[i] = int(*depth)
		}
	}

	return listDepths
}

func (this *Display) DefaultGC(screenNumber int) GC {
	displayC := (*C.Display)(this)
	defaultGraphicsContext := C.XDefaultGC(displayC, C.int(screenNumber))

	return GC(defaultGraphicsContext)
}

func (this *Display) DefaultRootWindow() Window {
	displayC := (*C.Display)(this)
	defaultRootWindow := C.XDefaultRootWindow(displayC)

	return Window(defaultRootWindow)
}

func (this *Display) DefaultScreenOfDisplay() *Screen {
	displayC := (*C.Display)(this)
	defaultScreen := C.XDefaultScreenOfDisplay(displayC)

	return (*Screen)(defaultScreen)
}

func (this *Display) ScreenOfDisplay(screenNumber int) *Screen {
	displayC := (*C.Display)(this)
	screen := C.XScreenOfDisplay(displayC, C.int(screenNumber))

	return (*Screen)(screen)
}

func (this *Display) DefaultScreen() int {
	displayC := (*C.Display)(this)
	defaultScreenNumber := C.XDefaultScreen(displayC)

	return int(defaultScreenNumber)
}

func (this *Display) DefaultVisual(screenNumber int) *Visual {
	displayC := (*C.Display)(this)
	defaultVisual := C.XDefaultVisual(displayC, C.int(screenNumber))

	return (*Visual)(defaultVisual)
}

func (this *Display) DisplayCells(screenNumber int) int {
	displayC := (*C.Display)(this)
	displayCells := C.XDisplayCells(displayC, C.int(screenNumber))

	return int(displayCells)
}

func (this *Display) DisplayPlanes(screenNumber int) int {
	displayC := (*C.Display)(this)
	displayPlanes := C.XDisplayPlanes(displayC, C.int(screenNumber))

	return int(displayPlanes)
}

func (this *Display) DisplayString() string {
	displayC := (*C.Display)(this)
	displayCStr := C.XDisplayString(displayC)
	displayString := C.GoString(displayCStr)

	C.free(unsafe.Pointer(displayCStr))

	return displayString
}

func (this *Display) ExtendedMaxRequestSize() int64 {
	displayC := (*C.Display)(this)
	extendedMxReqestSize := C.XExtendedMaxRequestSize(displayC)

	return int64(extendedMxReqestSize)
}

func (this *Display) MaxRequestSize() int64 {
	displayC := (*C.Display)(this)
	maxReqestSize := C.XMaxRequestSize(displayC)

	return int64(maxReqestSize)
}

func (this *Display) LastKnownRequestProcessed() uint64 {
	displayC := (*C.Display)(this)
	lastKnownRequestProcessed := C.XLastKnownRequestProcessed(displayC)

	return uint64(lastKnownRequestProcessed)
}

func (this *Display) NextRequest() uint64 {
	displayC := (*C.Display)(this)
	nextRequest := C.XNextRequest(displayC)

	return uint64(nextRequest)
}

func (this *Display) ProtocolVersion() int {
	displayC := (*C.Display)(this)
	protocolVersion := C.XProtocolVersion(displayC)

	return int(protocolVersion)
}

func (this *Display) ProtocolRevision() int {
	displayC := (*C.Display)(this)
	protocolrevision := C.XProtocolRevision(displayC)

	return int(protocolrevision)
}

func (this *Display) QLength() int {
	displayC := (*C.Display)(this)
	qLength := C.XQLength(displayC)

	return int(qLength)
}

func (this *Display) RootWindow(screenNumber int) Window {
	displayC := (*C.Display)(this)
	rootWindow := C.XRootWindow(displayC, C.int(screenNumber))

	return Window(rootWindow)
}

func (this *Display) ScreenCount() int {
	displayC := (*C.Display)(this)
	screenCount := C.XScreenCount(displayC)

	return int(screenCount)
}

func (this *Display) ServerVendor() string {
	displayC := (*C.Display)(this)
	serverVendorCStr := C.XServerVendor(displayC)
	serverVendor := C.GoString(serverVendorCStr)

	C.free(unsafe.Pointer(serverVendorCStr))

	return serverVendor
}

func (this *Display) VendorRelease() int {
	displayC := (*C.Display)(this)
	vendorRelease := C.XVendorRelease(displayC)

	return int(vendorRelease)
}

func (this *Display) ListPixmapFormats() []XPixmapFormatValues {
	var values []XPixmapFormatValues
	var countReturn C.int
	displayC := (*C.Display)(this)
	valuesPtr := C.XListPixmapFormats(displayC, &countReturn)

	if valuesPtr != nil {
		valuesPtrInt := uintptr(unsafe.Pointer(valuesPtr))
		valuesCSize := C.sizeof_XPixmapFormatValues
		values = make([]XPixmapFormatValues, int(countReturn))

		for i := range values {
			valuesPtrIndex := valuesPtrInt + uintptr(valuesCSize * i)
			valuesPtrValue := (*C.XPixmapFormatValues)(unsafe.Pointer(valuesPtrIndex))
			values[i].Depth = int((*valuesPtrValue).depth)
			values[i].BitsPerPixel = int((*valuesPtrValue).bits_per_pixel)
			values[i].ScanlinePad = int((*valuesPtrValue).scanline_pad)
		}

		C.XFree(unsafe.Pointer(valuesPtr))
	}

	return values
}

func (this *Display) ImageByteOrder() int {
	displayC := (*C.Display)(this)
	imageByteOrder := C.XImageByteOrder(displayC)

	return int(imageByteOrder)
}

func (this *Display) BitmapUnit() int {
	displayC := (*C.Display)(this)
	bitmapUnit := C.XBitmapUnit(displayC)

	return int(bitmapUnit)
}

func (this *Display) BitmapBitOrder() int {
	displayC := (*C.Display)(this)
	bitmapBitOrder := C.XBitmapBitOrder(displayC)

	return int(bitmapBitOrder)
}

func (this *Display) BitmapPad() int {
	displayC := (*C.Display)(this)
	bitmapPad := C.XBitmapPad(displayC)

	return int(bitmapPad)
}

func (this *Display) DisplayHeight(screenNumber int) int {
	displayC := (*C.Display)(this)
	displayHeight := C.XDisplayHeight(displayC, C.int(screenNumber))

	return int(displayHeight)
}

func (this *Display) DisplayHeightMM(screenNumber int) int {
	displayC := (*C.Display)(this)
	displayHeightMM := C.XDisplayHeightMM(displayC, C.int(screenNumber))

	return int(displayHeightMM)
}

func (this *Display) DisplayWidth(screenNumber int) int {
	displayC := (*C.Display)(this)
	displayWidth := C.XDisplayWidth(displayC, C.int(screenNumber))

	return int(displayWidth)
}

func (this *Display) DisplayWidthMM(screenNumber int) int {
	displayC := (*C.Display)(this)
	displayWidthMM := C.XDisplayWidthMM(displayC, C.int(screenNumber))

	return int(displayWidthMM)
}

func (this *Screen) BlackPixelOfScreen() uint64 {
	screenC := (*C.Screen)(this)
	blackPixel := C.XBlackPixelOfScreen(screenC)

	return uint64(blackPixel)
}

func (this *Screen) WhitePixelOfScreen() uint64 {
	screenC := (*C.Screen)(this)
	whitePixel := C.XWhitePixelOfScreen(screenC)

	return uint64(whitePixel)
}

func (this *Screen) CellsOfScreen(screen *Screen) int {
	screenC := (*C.Screen)(this)
	cells := C.XCellsOfScreen(screenC)

	return int(cells)
}

func (this *Screen) DefaultColormapOfScreen() Colormap {
	screenC := (*C.Screen)(this)
	colormap := C.XDefaultColormapOfScreen(screenC)

	return Colormap(colormap)
}

func (this *Screen) DefaultDepthOfScreen() int {
	screenC := (*C.Screen)(this)
	defaultDepth := C.XDefaultDepthOfScreen(screenC)

	return int(defaultDepth)
}

func (this *Screen) DefaultGCOfScreen() GC {
	screenC := (*C.Screen)(this)
	defaultGC := C.XDefaultGCOfScreen(screenC)

	return GC(defaultGC)
}

func (this *Screen) DefaultVisualOfScreen() *Visual {
	screenC := (*C.Screen)(this)
	visual := C.XDefaultVisualOfScreen(screenC)

	return (*Visual)(visual)
}

func (this *Screen) DoesBackingStore() int {
	screenC := (*C.Screen)(this)
	doesBackingStore := C.XDoesBackingStore(screenC)

	return int(doesBackingStore)
}

func (this *Screen) DoesSaveUnders() bool {
	screenC := (*C.Screen)(this)
	doesSaveUnders := C.XDoesSaveUnders(screenC)

	return bool(doesSaveUnders != 0)
}

func (this *Screen) DisplayOfScreen() *Display {
	screenC := (*C.Screen)(this)
	display := C.XDisplayOfScreen(screenC)

	return (*Display)(display)
}

func (this *Screen) ScreenNumberOfScreen() int {
	screenC := (*C.Screen)(this)
	screenNumber := C.XScreenNumberOfScreen(screenC)

	return int(screenNumber)
}

func (this *Screen) EventMaskOfScreen() int64 {
	screenC := (*C.Screen)(this)
	eventMask := C.XEventMaskOfScreen(screenC)

	return int64(eventMask)
}

func (this *Screen) WidthOfScreen() int {
	screenC := (*C.Screen)(this)
	width := C.XWidthOfScreen(screenC)

	return int(width)
}

func (this *Screen) HeightOfScreen() int {
	screenC := (*C.Screen)(this)
	height := C.XHeightOfScreen(screenC)

	return int(height)
}

func (this *Screen) WidthMMOfScreen() int {
	screenC := (*C.Screen)(this)
	widthMM := C.XWidthMMOfScreen(screenC)

	return int(widthMM)
}

func (this *Screen) HeightMMOfScreen() int {
	screenC := (*C.Screen)(this)
	heightMM := C.XHeightMMOfScreen(screenC)

	return int(heightMM)
}

func (this *Screen) MaxCmapsOfScreen() int {
	screenC := (*C.Screen)(this)
	maxCmaps := C.XMaxCmapsOfScreen(screenC)

	return int(maxCmaps)
}

func (this *Screen) MinCmapsOfScreen() int {
	screenC := (*C.Screen)(this)
	minCmaps := C.XMinCmapsOfScreen(screenC)

	return int(minCmaps)
}

func (this *Screen) PlanesOfScreen() int {
	screenC := (*C.Screen)(this)
	planes := C.XPlanesOfScreen(screenC)

	return int(planes)
}

func (this *Screen) RootWindowOfScreen() Window {
	screenC := (*C.Screen)(this)
	rootWindow := C.XRootWindowOfScreen(screenC)

	return Window(rootWindow)
}

func (this *Display) NoOp() {
	displayC := (*C.Display)(this)

	C.XNoOp(displayC)
}

func (this *Display) CloseDisplay() {
	displayC := (*C.Display)(this)

	C.XCloseDisplay(displayC)
}

func (this *Display) SetCloseDownMode(closeMode int) {
	displayC := (*C.Display)(this)

	C.XSetCloseDownMode(displayC, C.int(closeMode))
}

func InitThreads() {
	C.XInitThreads()
}

func (this *Display) LockDisplay() {
	displayC := (*C.Display)(this)

	C.XLockDisplay(displayC)
}

func (this *Display) UnlockDisplay() {
	displayC := (*C.Display)(this)

	C.XUnlockDisplay(displayC)
}

func (this *Visual) VisualIDFromVisual() VisualID {
	visualC := (*C.Visual)(this)
	visualIDC := C.XVisualIDFromVisual(visualC)

	return VisualID(visualIDC)
}

func (this *Display) CreateWindow(parent Window, x, y int, width, height, borderWidth uint, depth int, class uint, visual *Visual, valuemask uint64, attributes *SetWindowAttributes) Window {
	displayC := (*C.Display)(this)
	windowC := (C.Window)(parent)
	xC := C.int(x)
	yC := C.int(y)
	widthC := C.uint(width)
	heightC := C.uint(height)
	borderWidthC := C.uint(borderWidth)
	depthC := C.int(depth)
	classC := C.uint(class)
	visualC := (*C.Visual)(visual)
	valuemaskC := C.ulong(valuemask)
	attributesC := newXSetWindowAttributes(attributes)
	window := C.XCreateWindow(displayC, windowC, xC, yC, widthC, heightC, borderWidthC, depthC, classC, visualC, valuemaskC, attributesC)

	return Window(window)
}

func (this *Display) CreateSimpleWindow(parent Window, x, y int, width, height, borderWidth uint, border, background uint64) Window {
	displayC := (*C.Display)(this)
	windowC := (C.Window)(parent)
	xC := C.int(x)
	yC := C.int(y)
	widthC := C.uint(width)
	heightC := C.uint(height)
	borderWidthC := C.uint(borderWidth)
	borderC := C.ulong(border)
	backgroundC := C.ulong(background)
	window := C.XCreateSimpleWindow(displayC, windowC, xC, yC, widthC, heightC, borderWidthC, borderC, backgroundC)

	return Window(window)
}

func (this *Display) DestroyWindow(window Window) {
	displayC := (*C.Display)(this)
	windowC := (C.Window)(window)

	C.XDestroyWindow(displayC, windowC)
}

func (this *Display) DestroySubwindows(window Window) {
	displayC := (*C.Display)(this)
	windowC := (C.Window)(window)

	C.XDestroySubwindows(displayC, windowC)
}

func (this *Display) MapWindow(window Window) {
	displayC := (*C.Display)(this)
	windowC := (C.Window)(window)

	C.XMapWindow(displayC, windowC)
}

func (this *Display) MapRaised(window Window) {
	displayC := (*C.Display)(this)
	windowC := (C.Window)(window)

	C.XMapRaised(displayC, windowC)
}

func (this *Display) MapSubwindows(window Window) {
	displayC := (*C.Display)(this)
	windowC := (C.Window)(window)

	C.XMapSubwindows(displayC, windowC)
}

func (this *Display) UnmapWindow(window Window) {
	displayC := (*C.Display)(this)
	windowC := (C.Window)(window)

	C.XUnmapWindow(displayC, windowC)
}

func (this *Display) UnmapSubwindows(window Window) {
	displayC := (*C.Display)(this)
	windowC := (C.Window)(window)

	C.XUnmapSubwindows(displayC, windowC)
}

func newXSetWindowAttributes(attributes *SetWindowAttributes) *C.XSetWindowAttributes {
	attributesC := new(C.XSetWindowAttributes)
	var saveUnderC C.int = 0
	var overrideRedirectC C.int = 0

	if attributes.SaveUnder {
		saveUnderC = 1
	}

	if attributes.OverrideRedirect {
		overrideRedirectC = 1
	}

	attributesC.background_pixmap = C.Pixmap(attributes.BackgroundPixmap)
	attributesC.background_pixel = C.ulong(attributes.BackgroundPixel)
	attributesC.border_pixmap = C.Pixmap(attributes.BorderPixmap)
	attributesC.border_pixel = C.ulong(attributes.BorderPixel)
	attributesC.bit_gravity = C.int(attributes.BitGravity)
	attributesC.win_gravity = C.int(attributes.WinGravity)
	attributesC.backing_store = C.int(attributes.BackingStore)
	attributesC.backing_planes = C.ulong(attributes.BackingPlanes)
	attributesC.backing_pixel = C.ulong(attributes.BackingPixel)
	attributesC.save_under = C.Bool(saveUnderC)
	attributesC.event_mask = C.long(attributes.EventMask)
	attributesC.do_not_propagate_mask = C.long(attributes.DoNotPropagateMask)
	attributesC.override_redirect = C.Bool(overrideRedirectC)
	attributesC.colormap = C.Colormap(attributes.WinColormap)
	attributesC.cursor = C.Cursor(attributes.WinCursor)

	return attributesC
}

