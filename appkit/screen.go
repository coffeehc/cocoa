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
	result_ := C.C_NSScreen_Init(n.Ptr())
	return MakeScreen(result_)
}

func (n *NSScreen) CanRepresentDisplayGamut(displayGamut DisplayGamut) bool {
	result_ := C.C_NSScreen_CanRepresentDisplayGamut(n.Ptr(), C.int(int(displayGamut)))
	return bool(result_)
}

func (n *NSScreen) ConvertRectFromBacking(rect foundation.Rect) foundation.Rect {
	result_ := C.C_NSScreen_ConvertRectFromBacking(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))))
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSScreen) ConvertRectToBacking(rect foundation.Rect) foundation.Rect {
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
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]Screen, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeScreen(r)
	}
	return goResult_
}

func (n *NSScreen) Depth() WindowDepth {
	result_ := C.C_NSScreen_Depth(n.Ptr())
	return WindowDepth(int32(result_))
}

func (n *NSScreen) Frame() foundation.Rect {
	result_ := C.C_NSScreen_Frame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSScreen) VisibleFrame() foundation.Rect {
	result_ := C.C_NSScreen_VisibleFrame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSScreen) ColorSpace() ColorSpace {
	result_ := C.C_NSScreen_ColorSpace(n.Ptr())
	return MakeColorSpace(result_)
}

func ScreensHaveSeparateSpaces() bool {
	result_ := C.C_NSScreen_ScreensHaveSeparateSpaces()
	return bool(result_)
}

func (n *NSScreen) BackingScaleFactor() coregraphics.Float {
	result_ := C.C_NSScreen_BackingScaleFactor(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSScreen) MaximumPotentialExtendedDynamicRangeColorComponentValue() coregraphics.Float {
	result_ := C.C_NSScreen_MaximumPotentialExtendedDynamicRangeColorComponentValue(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSScreen) MaximumExtendedDynamicRangeColorComponentValue() coregraphics.Float {
	result_ := C.C_NSScreen_MaximumExtendedDynamicRangeColorComponentValue(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSScreen) MaximumReferenceExtendedDynamicRangeColorComponentValue() coregraphics.Float {
	result_ := C.C_NSScreen_MaximumReferenceExtendedDynamicRangeColorComponentValue(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSScreen) LocalizedName() string {
	result_ := C.C_NSScreen_LocalizedName(n.Ptr())
	return foundation.MakeString(result_).String()
}
