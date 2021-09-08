package appkit

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
	CanRepresentDisplayGamut(displayGamut DisplayGamut) bool
	ConvertRectFromBacking(rect foundation.Rect) foundation.Rect
	ConvertRectToBacking(rect foundation.Rect) foundation.Rect
	Depth() WindowDepth
	Frame() foundation.Rect
	DeviceDescription() map[DeviceDescriptionKey]objc.Object
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

func MakeScreen(ptr unsafe.Pointer) NSScreen {
	return NSScreen{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocScreen() NSScreen {
	result_ := C.C_NSScreen_AllocScreen()
	return MakeScreen(result_)
}

func (n NSScreen) Init() NSScreen {
	result_ := C.C_NSScreen_Init(n.Ptr())
	return MakeScreen(result_)
}

func NewScreen() NSScreen {
	result_ := C.C_NSScreen_NewScreen()
	return MakeScreen(result_)
}

func (n NSScreen) Autorelease() NSScreen {
	result_ := C.C_NSScreen_Autorelease(n.Ptr())
	return MakeScreen(result_)
}

func (n NSScreen) Retain() NSScreen {
	result_ := C.C_NSScreen_Retain(n.Ptr())
	return MakeScreen(result_)
}

func (n NSScreen) CanRepresentDisplayGamut(displayGamut DisplayGamut) bool {
	result_ := C.C_NSScreen_CanRepresentDisplayGamut(n.Ptr(), C.int(int(displayGamut)))
	return bool(result_)
}

func (n NSScreen) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSScreen_ConvertRectFromBacking(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSScreen) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSScreen_ConvertRectToBacking(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func MainScreen() Screen {
	result_ := C.C_NSScreen_MainScreen()
	return MakeScreen(result_)
}

func DeepestScreen() Screen {
	result_ := C.C_NSScreen_DeepestScreen()
	return MakeScreen(result_)
}

func Screens() []Screen {
	result_ := C.C_NSScreen_Screens()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]Screen, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeScreen(r)
	}
	return goResult_
}

func (n NSScreen) Depth() WindowDepth {
	result_ := C.C_NSScreen_Depth(n.Ptr())
	return WindowDepth(int32(result_))
}

func (n NSScreen) Frame() foundation.Rect {
	result_ := C.C_NSScreen_Frame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSScreen) DeviceDescription() map[DeviceDescriptionKey]objc.Object {
	result_ := C.C_NSScreen_DeviceDescription(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[DeviceDescriptionKey]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[DeviceDescriptionKey(foundation.MakeString(k).String())] = objc.MakeObject(v)
	}
	return goResult_
}

func (n NSScreen) VisibleFrame() foundation.Rect {
	result_ := C.C_NSScreen_VisibleFrame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSScreen) ColorSpace() ColorSpace {
	result_ := C.C_NSScreen_ColorSpace(n.Ptr())
	return MakeColorSpace(result_)
}

func ScreensHaveSeparateSpaces() bool {
	result_ := C.C_NSScreen_ScreensHaveSeparateSpaces()
	return bool(result_)
}

func (n NSScreen) BackingScaleFactor() coregraphics.Float {
	result_ := C.C_NSScreen_BackingScaleFactor(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSScreen) MaximumPotentialExtendedDynamicRangeColorComponentValue() coregraphics.Float {
	result_ := C.C_NSScreen_MaximumPotentialExtendedDynamicRangeColorComponentValue(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSScreen) MaximumExtendedDynamicRangeColorComponentValue() coregraphics.Float {
	result_ := C.C_NSScreen_MaximumExtendedDynamicRangeColorComponentValue(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSScreen) MaximumReferenceExtendedDynamicRangeColorComponentValue() coregraphics.Float {
	result_ := C.C_NSScreen_MaximumReferenceExtendedDynamicRangeColorComponentValue(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSScreen) LocalizedName() string {
	result_ := C.C_NSScreen_LocalizedName(n.Ptr())
	return foundation.MakeString(result_).String()
}
