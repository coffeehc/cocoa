package appkit

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
	FlushGraphics()
	IsDrawingToScreen() bool
	IsFlipped() bool
	CompositingOperation() CompositingOperation
	SetCompositingOperation(value CompositingOperation)
	ImageInterpolation() ImageInterpolation
	SetImageInterpolation(value ImageInterpolation)
	ShouldAntialias() bool
	SetShouldAntialias(value bool)
	PatternPhase() foundation.Point
	SetPatternPhase(value foundation.Point)
	ColorRenderingIntent() ColorRenderingIntent
	SetColorRenderingIntent(value ColorRenderingIntent)
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
	result_ := C.C_NSGraphicsContext_Init(n.Ptr())
	return MakeGraphicsContext(result_)
}

func GraphicsContextWithBitmapImageRep(bitmapRep BitmapImageRep) GraphicsContext {
	result_ := C.C_NSGraphicsContext_GraphicsContextWithBitmapImageRep(objc.ExtractPtr(bitmapRep))
	return MakeGraphicsContext(result_)
}

func GraphicsContextWithWindow(window Window) GraphicsContext {
	result_ := C.C_NSGraphicsContext_GraphicsContextWithWindow(objc.ExtractPtr(window))
	return MakeGraphicsContext(result_)
}

func GraphicsContext_RestoreGraphicsState() {
	C.C_NSGraphicsContext_GraphicsContext_RestoreGraphicsState()
}

func GraphicsContext_SaveGraphicsState() {
	C.C_NSGraphicsContext_GraphicsContext_SaveGraphicsState()
}

func GraphicsContext_CurrentContextDrawingToScreen() bool {
	result_ := C.C_NSGraphicsContext_GraphicsContext_CurrentContextDrawingToScreen()
	return bool(result_)
}

func (n *NSGraphicsContext) FlushGraphics() {
	C.C_NSGraphicsContext_FlushGraphics(n.Ptr())
}

func GraphicsContext_CurrentContext() GraphicsContext {
	result_ := C.C_NSGraphicsContext_GraphicsContext_CurrentContext()
	return MakeGraphicsContext(result_)
}

func GraphicsContext_SetCurrentContext(value GraphicsContext) {
	C.C_NSGraphicsContext_GraphicsContext_SetCurrentContext(objc.ExtractPtr(value))
}

func (n *NSGraphicsContext) IsDrawingToScreen() bool {
	result_ := C.C_NSGraphicsContext_IsDrawingToScreen(n.Ptr())
	return bool(result_)
}

func (n *NSGraphicsContext) IsFlipped() bool {
	result_ := C.C_NSGraphicsContext_IsFlipped(n.Ptr())
	return bool(result_)
}

func (n *NSGraphicsContext) CompositingOperation() CompositingOperation {
	result_ := C.C_NSGraphicsContext_CompositingOperation(n.Ptr())
	return CompositingOperation(uint(result_))
}

func (n *NSGraphicsContext) SetCompositingOperation(value CompositingOperation) {
	C.C_NSGraphicsContext_SetCompositingOperation(n.Ptr(), C.uint(uint(value)))
}

func (n *NSGraphicsContext) ImageInterpolation() ImageInterpolation {
	result_ := C.C_NSGraphicsContext_ImageInterpolation(n.Ptr())
	return ImageInterpolation(uint(result_))
}

func (n *NSGraphicsContext) SetImageInterpolation(value ImageInterpolation) {
	C.C_NSGraphicsContext_SetImageInterpolation(n.Ptr(), C.uint(uint(value)))
}

func (n *NSGraphicsContext) ShouldAntialias() bool {
	result_ := C.C_NSGraphicsContext_ShouldAntialias(n.Ptr())
	return bool(result_)
}

func (n *NSGraphicsContext) SetShouldAntialias(value bool) {
	C.C_NSGraphicsContext_SetShouldAntialias(n.Ptr(), C.bool(value))
}

func (n *NSGraphicsContext) PatternPhase() foundation.Point {
	result_ := C.C_NSGraphicsContext_PatternPhase(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n *NSGraphicsContext) SetPatternPhase(value foundation.Point) {
	C.C_NSGraphicsContext_SetPatternPhase(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(value))))
}

func (n *NSGraphicsContext) ColorRenderingIntent() ColorRenderingIntent {
	result_ := C.C_NSGraphicsContext_ColorRenderingIntent(n.Ptr())
	return ColorRenderingIntent(int(result_))
}

func (n *NSGraphicsContext) SetColorRenderingIntent(value ColorRenderingIntent) {
	C.C_NSGraphicsContext_SetColorRenderingIntent(n.Ptr(), C.int(int(value)))
}
