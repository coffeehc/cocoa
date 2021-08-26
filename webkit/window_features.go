package webkit

// #include "window_features.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type WindowFeatures interface {
	objc.Object
	AllowsResizing() foundation.Number
	Height() foundation.Number
	Width() foundation.Number
	X() foundation.Number
	Y() foundation.Number
	MenuBarVisibility() foundation.Number
	StatusBarVisibility() foundation.Number
	ToolbarsVisibility() foundation.Number
}

type WKWindowFeatures struct {
	objc.NSObject
}

func MakeWindowFeatures(ptr unsafe.Pointer) WKWindowFeatures {
	return WKWindowFeatures{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocWindowFeatures() WKWindowFeatures {
	return MakeWindowFeatures(C.C_WindowFeatures_Alloc())
}

func (w WKWindowFeatures) Init() WindowFeatures {
	result_ := C.C_WKWindowFeatures_Init(w.Ptr())
	return MakeWindowFeatures(result_)
}

func (w WKWindowFeatures) AllowsResizing() foundation.Number {
	result_ := C.C_WKWindowFeatures_AllowsResizing(w.Ptr())
	return foundation.MakeNumber(result_)
}

func (w WKWindowFeatures) Height() foundation.Number {
	result_ := C.C_WKWindowFeatures_Height(w.Ptr())
	return foundation.MakeNumber(result_)
}

func (w WKWindowFeatures) Width() foundation.Number {
	result_ := C.C_WKWindowFeatures_Width(w.Ptr())
	return foundation.MakeNumber(result_)
}

func (w WKWindowFeatures) X() foundation.Number {
	result_ := C.C_WKWindowFeatures_X(w.Ptr())
	return foundation.MakeNumber(result_)
}

func (w WKWindowFeatures) Y() foundation.Number {
	result_ := C.C_WKWindowFeatures_Y(w.Ptr())
	return foundation.MakeNumber(result_)
}

func (w WKWindowFeatures) MenuBarVisibility() foundation.Number {
	result_ := C.C_WKWindowFeatures_MenuBarVisibility(w.Ptr())
	return foundation.MakeNumber(result_)
}

func (w WKWindowFeatures) StatusBarVisibility() foundation.Number {
	result_ := C.C_WKWindowFeatures_StatusBarVisibility(w.Ptr())
	return foundation.MakeNumber(result_)
}

func (w WKWindowFeatures) ToolbarsVisibility() foundation.Number {
	result_ := C.C_WKWindowFeatures_ToolbarsVisibility(w.Ptr())
	return foundation.MakeNumber(result_)
}
