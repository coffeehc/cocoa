package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "screen.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Screen interface {
	objc.Object
	ConvertRectFromBacking(rect foundation.Rect) foundation.Rect
	ConvertRectToBacking(rect foundation.Rect) foundation.Rect
	Frame() foundation.Rect
	VisibleFrame() foundation.Rect
	ColorSpace() ColorSpace
	BackingScaleFactor() coregraphics.Float
	MaximumPotentialExtendedDynamicRangeColorComponentValue() coregraphics.Float
	MaximumExtendedDynamicRangeColorComponentValue() coregraphics.Float
	MaximumReferenceExtendedDynamicRangeColorComponentValue() coregraphics.Float
	LocalizedName() string
}

type NSScreen struct {
	objc.NSObject
}

func MakeScreen(ptr unsafe.Pointer) *NSScreen {
	if ptr == nil {
		return nil
	}
	return &NSScreen{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocScreen() *NSScreen {
	return MakeScreen(C.C_Screen_Alloc())
}

func (n *NSScreen) Init() Screen {
	result := C.C_NSScreen_Init(n.Ptr())
	return MakeScreen(result)
}

func (n *NSScreen) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	result := C.C_NSScreen_ConvertRectFromBacking(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSScreen) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
	result := C.C_NSScreen_ConvertRectToBacking(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSScreen) Frame() foundation.Rect {
	result := C.C_NSScreen_Frame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSScreen) VisibleFrame() foundation.Rect {
	result := C.C_NSScreen_VisibleFrame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result)))
}

func (n *NSScreen) ColorSpace() ColorSpace {
	result := C.C_NSScreen_ColorSpace(n.Ptr())
	return MakeColorSpace(result)
}

func (n *NSScreen) BackingScaleFactor() coregraphics.Float {
	result := C.C_NSScreen_BackingScaleFactor(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSScreen) MaximumPotentialExtendedDynamicRangeColorComponentValue() coregraphics.Float {
	result := C.C_NSScreen_MaximumPotentialExtendedDynamicRangeColorComponentValue(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSScreen) MaximumExtendedDynamicRangeColorComponentValue() coregraphics.Float {
	result := C.C_NSScreen_MaximumExtendedDynamicRangeColorComponentValue(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSScreen) MaximumReferenceExtendedDynamicRangeColorComponentValue() coregraphics.Float {
	result := C.C_NSScreen_MaximumReferenceExtendedDynamicRangeColorComponentValue(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSScreen) LocalizedName() string {
	result := C.C_NSScreen_LocalizedName(n.Ptr())
	return foundation.MakeString(result).String()
}
