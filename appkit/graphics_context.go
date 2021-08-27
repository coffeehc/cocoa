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
	CGContext() coregraphics.ContextRef
	IsDrawingToScreen() bool
	Attributes() map[GraphicsContextAttributeKey]objc.Object
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

func MakeGraphicsContext(ptr unsafe.Pointer) NSGraphicsContext {
	return NSGraphicsContext{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocGraphicsContext() NSGraphicsContext {
	return MakeGraphicsContext(C.C_GraphicsContext_Alloc())
}

func (n NSGraphicsContext) Init() GraphicsContext {
	result_ := C.C_NSGraphicsContext_Init(n.Ptr())
	return MakeGraphicsContext(result_)
}

func GraphicsContextWithAttributes(attributes map[GraphicsContextAttributeKey]objc.Object) GraphicsContext {
	var cAttributes C.Dictionary
	if len(attributes) > 0 {
		cAttributesKeyData := make([]unsafe.Pointer, len(attributes))
		cAttributesValueData := make([]unsafe.Pointer, len(attributes))
		var idx = 0
		for k, v := range attributes {
			cAttributesKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cAttributesValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cAttributes.key_data = unsafe.Pointer(&cAttributesKeyData[0])
		cAttributes.value_data = unsafe.Pointer(&cAttributesValueData[0])
		cAttributes.len = C.int(len(attributes))
	}
	result_ := C.C_NSGraphicsContext_GraphicsContextWithAttributes(cAttributes)
	return MakeGraphicsContext(result_)
}

func GraphicsContextWithBitmapImageRep(bitmapRep BitmapImageRep) GraphicsContext {
	result_ := C.C_NSGraphicsContext_GraphicsContextWithBitmapImageRep(objc.ExtractPtr(bitmapRep))
	return MakeGraphicsContext(result_)
}

func GraphicsContextWithCGContext_Flipped(graphicsPort coregraphics.ContextRef, initialFlippedState bool) GraphicsContext {
	result_ := C.C_NSGraphicsContext_GraphicsContextWithCGContext_Flipped(unsafe.Pointer(graphicsPort), C.bool(initialFlippedState))
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

func (n NSGraphicsContext) FlushGraphics() {
	C.C_NSGraphicsContext_FlushGraphics(n.Ptr())
}

func GraphicsContext_CurrentContext() GraphicsContext {
	result_ := C.C_NSGraphicsContext_GraphicsContext_CurrentContext()
	return MakeGraphicsContext(result_)
}

func GraphicsContext_SetCurrentContext(value GraphicsContext) {
	C.C_NSGraphicsContext_GraphicsContext_SetCurrentContext(objc.ExtractPtr(value))
}

func (n NSGraphicsContext) CGContext() coregraphics.ContextRef {
	result_ := C.C_NSGraphicsContext_CGContext(n.Ptr())
	return coregraphics.ContextRef(result_)
}

func (n NSGraphicsContext) IsDrawingToScreen() bool {
	result_ := C.C_NSGraphicsContext_IsDrawingToScreen(n.Ptr())
	return bool(result_)
}

func (n NSGraphicsContext) Attributes() map[GraphicsContextAttributeKey]objc.Object {
	result_ := C.C_NSGraphicsContext_Attributes(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[GraphicsContextAttributeKey]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[GraphicsContextAttributeKey(foundation.MakeString(k).String())] = objc.MakeObject(v)
	}
	return goResult_
}

func (n NSGraphicsContext) IsFlipped() bool {
	result_ := C.C_NSGraphicsContext_IsFlipped(n.Ptr())
	return bool(result_)
}

func (n NSGraphicsContext) CompositingOperation() CompositingOperation {
	result_ := C.C_NSGraphicsContext_CompositingOperation(n.Ptr())
	return CompositingOperation(uint(result_))
}

func (n NSGraphicsContext) SetCompositingOperation(value CompositingOperation) {
	C.C_NSGraphicsContext_SetCompositingOperation(n.Ptr(), C.uint(uint(value)))
}

func (n NSGraphicsContext) ImageInterpolation() ImageInterpolation {
	result_ := C.C_NSGraphicsContext_ImageInterpolation(n.Ptr())
	return ImageInterpolation(uint(result_))
}

func (n NSGraphicsContext) SetImageInterpolation(value ImageInterpolation) {
	C.C_NSGraphicsContext_SetImageInterpolation(n.Ptr(), C.uint(uint(value)))
}

func (n NSGraphicsContext) ShouldAntialias() bool {
	result_ := C.C_NSGraphicsContext_ShouldAntialias(n.Ptr())
	return bool(result_)
}

func (n NSGraphicsContext) SetShouldAntialias(value bool) {
	C.C_NSGraphicsContext_SetShouldAntialias(n.Ptr(), C.bool(value))
}

func (n NSGraphicsContext) PatternPhase() foundation.Point {
	result_ := C.C_NSGraphicsContext_PatternPhase(n.Ptr())
	return foundation.Point(coregraphics.FromCGPointPointer(unsafe.Pointer(&result_)))
}

func (n NSGraphicsContext) SetPatternPhase(value foundation.Point) {
	C.C_NSGraphicsContext_SetPatternPhase(n.Ptr(), *(*C.CGPoint)(coregraphics.ToCGPointPointer(coregraphics.Point(value))))
}

func (n NSGraphicsContext) ColorRenderingIntent() ColorRenderingIntent {
	result_ := C.C_NSGraphicsContext_ColorRenderingIntent(n.Ptr())
	return ColorRenderingIntent(int(result_))
}

func (n NSGraphicsContext) SetColorRenderingIntent(value ColorRenderingIntent) {
	C.C_NSGraphicsContext_SetColorRenderingIntent(n.Ptr(), C.int(int(value)))
}
