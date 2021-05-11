package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "graphics_context.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type GraphicsContext interface {
	objc.Object
	RestoreGraphicsState()
	SaveGraphicsState()
	FlushGraphics()
	CGContext() coregraphics.ContextRef
	IsDrawingToScreen() bool
	IsFlipped() bool
	ShouldAntialias() bool
	SetShouldAntialias(value bool)
	PatternPhase() foundation.Point
	SetPatternPhase(value foundation.Point)
}

type NSGraphicsContext struct {
	objc.NSObject
}

func MakeGraphicsContext(ptr unsafe.Pointer) *NSGraphicsContext {
	if ptr == nil {
		return nil
	}
	return &NSGraphicsContext{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocGraphicsContext() *NSGraphicsContext {
	return MakeGraphicsContext(C.C_GraphicsContext_Alloc())
}

func (n *NSGraphicsContext) Init() GraphicsContext {
	result := C.C_NSGraphicsContext_Init(n.Ptr())
	return MakeGraphicsContext(result)
}

func GraphicsContextWithBitmapImageRep(bitmapRep BitmapImageRep) GraphicsContext {
	result := C.C_NSGraphicsContext_GraphicsContextWithBitmapImageRep(objc.ExtractPtr(bitmapRep))
	return MakeGraphicsContext(result)
}

func GraphicsContextWithCGContext_Flipped(graphicsPort coregraphics.ContextRef, initialFlippedState bool) GraphicsContext {
	result := C.C_NSGraphicsContext_GraphicsContextWithCGContext_Flipped(*(*C.CGContextRef)(coregraphics.ToCGContextRefPointer(graphicsPort)), C.bool(initialFlippedState))
	return MakeGraphicsContext(result)
}

func GraphicsContextWithWindow(window Window) GraphicsContext {
	result := C.C_NSGraphicsContext_GraphicsContextWithWindow(objc.ExtractPtr(window))
	return MakeGraphicsContext(result)
}

func GraphicsContextRestoreGraphicsState() {
	C.C_NSGraphicsContext_GraphicsContextRestoreGraphicsState()
}

func (n *NSGraphicsContext) RestoreGraphicsState() {
	C.C_NSGraphicsContext_RestoreGraphicsState(n.Ptr())
}

func GraphicsContextSaveGraphicsState() {
	C.C_NSGraphicsContext_GraphicsContextSaveGraphicsState()
}

func (n *NSGraphicsContext) SaveGraphicsState() {
	C.C_NSGraphicsContext_SaveGraphicsState(n.Ptr())
}

func GraphicsContextCurrentContextDrawingToScreen() bool {
	result := C.C_NSGraphicsContext_GraphicsContextCurrentContextDrawingToScreen()
	return bool(result)
}

func (n *NSGraphicsContext) FlushGraphics() {
	C.C_NSGraphicsContext_FlushGraphics(n.Ptr())
}

func (n *NSGraphicsContext) CGContext() coregraphics.ContextRef {
	result := C.C_NSGraphicsContext_CGContext(n.Ptr())
	return coregraphics.FromCGContextRefPointer(unsafe.Pointer(&result))
}

func (n *NSGraphicsContext) IsDrawingToScreen() bool {
	result := C.C_NSGraphicsContext_IsDrawingToScreen(n.Ptr())
	return bool(result)
}

func (n *NSGraphicsContext) IsFlipped() bool {
	result := C.C_NSGraphicsContext_IsFlipped(n.Ptr())
	return bool(result)
}

func (n *NSGraphicsContext) ShouldAntialias() bool {
	result := C.C_NSGraphicsContext_ShouldAntialias(n.Ptr())
	return bool(result)
}

func (n *NSGraphicsContext) SetShouldAntialias(value bool) {
	C.C_NSGraphicsContext_SetShouldAntialias(n.Ptr(), C.bool(value))
}

func (n *NSGraphicsContext) PatternPhase() foundation.Point {
	result := C.C_NSGraphicsContext_PatternPhase(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result)))
}

func (n *NSGraphicsContext) SetPatternPhase(value foundation.Point) {
	C.C_NSGraphicsContext_SetPatternPhase(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(value))))
}
