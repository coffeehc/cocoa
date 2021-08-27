package appkit

// #include "path_control.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type PathControl interface {
	Control
	SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool)
	PathStyle() PathStyle
	SetPathStyle(value PathStyle)
	BackgroundColor() Color
	SetBackgroundColor(value Color)
	DoubleAction() objc.Selector
	SetDoubleAction(value objc.Selector)
	URL() foundation.URL
	SetURL(value foundation.URL)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	AllowedTypes() []string
	SetAllowedTypes(value []string)
	ClickedPathItem() PathControlItem
	IsEditable() bool
	SetEditable(value bool)
	PathItems() []PathControlItem
	SetPathItems(value []PathControlItem)
	PlaceholderAttributedString() foundation.AttributedString
	SetPlaceholderAttributedString(value foundation.AttributedString)
	PlaceholderString() string
	SetPlaceholderString(value string)
}

type NSPathControl struct {
	NSControl
}

func MakePathControl(ptr unsafe.Pointer) NSPathControl {
	return NSPathControl{
		NSControl: MakeControl(ptr),
	}
}

func AllocPathControl() NSPathControl {
	return MakePathControl(C.C_PathControl_Alloc())
}

func (n NSPathControl) InitWithFrame(frameRect foundation.Rect) PathControl {
	result_ := C.C_NSPathControl_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakePathControl(result_)
}

func (n NSPathControl) InitWithCoder(coder foundation.Coder) PathControl {
	result_ := C.C_NSPathControl_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakePathControl(result_)
}

func (n NSPathControl) Init() PathControl {
	result_ := C.C_NSPathControl_Init(n.Ptr())
	return MakePathControl(result_)
}

func (n NSPathControl) SetDraggingSourceOperationMask_ForLocal(mask DragOperation, isLocal bool) {
	C.C_NSPathControl_SetDraggingSourceOperationMask_ForLocal(n.Ptr(), C.uint(uint(mask)), C.bool(isLocal))
}

func (n NSPathControl) PathStyle() PathStyle {
	result_ := C.C_NSPathControl_PathStyle(n.Ptr())
	return PathStyle(int(result_))
}

func (n NSPathControl) SetPathStyle(value PathStyle) {
	C.C_NSPathControl_SetPathStyle(n.Ptr(), C.int(int(value)))
}

func (n NSPathControl) BackgroundColor() Color {
	result_ := C.C_NSPathControl_BackgroundColor(n.Ptr())
	return MakeColor(result_)
}

func (n NSPathControl) SetBackgroundColor(value Color) {
	C.C_NSPathControl_SetBackgroundColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSPathControl) DoubleAction() objc.Selector {
	result_ := C.C_NSPathControl_DoubleAction(n.Ptr())
	return objc.Selector(result_)
}

func (n NSPathControl) SetDoubleAction(value objc.Selector) {
	C.C_NSPathControl_SetDoubleAction(n.Ptr(), unsafe.Pointer(value))
}

func (n NSPathControl) URL() foundation.URL {
	result_ := C.C_NSPathControl_URL(n.Ptr())
	return foundation.MakeURL(result_)
}

func (n NSPathControl) SetURL(value foundation.URL) {
	C.C_NSPathControl_SetURL(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSPathControl) Delegate() objc.Object {
	result_ := C.C_NSPathControl_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSPathControl) SetDelegate(value objc.Object) {
	C.C_NSPathControl_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSPathControl) AllowedTypes() []string {
	result_ := C.C_NSPathControl_AllowedTypes(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeString(r).String()
	}
	return goResult_
}

func (n NSPathControl) SetAllowedTypes(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = foundation.NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSPathControl_SetAllowedTypes(n.Ptr(), cValue)
}

func (n NSPathControl) ClickedPathItem() PathControlItem {
	result_ := C.C_NSPathControl_ClickedPathItem(n.Ptr())
	return MakePathControlItem(result_)
}

func (n NSPathControl) IsEditable() bool {
	result_ := C.C_NSPathControl_IsEditable(n.Ptr())
	return bool(result_)
}

func (n NSPathControl) SetEditable(value bool) {
	C.C_NSPathControl_SetEditable(n.Ptr(), C.bool(value))
}

func (n NSPathControl) PathItems() []PathControlItem {
	result_ := C.C_NSPathControl_PathItems(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]PathControlItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakePathControlItem(r)
	}
	return goResult_
}

func (n NSPathControl) SetPathItems(value []PathControlItem) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = objc.ExtractPtr(v)
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSPathControl_SetPathItems(n.Ptr(), cValue)
}

func (n NSPathControl) PlaceholderAttributedString() foundation.AttributedString {
	result_ := C.C_NSPathControl_PlaceholderAttributedString(n.Ptr())
	return foundation.MakeAttributedString(result_)
}

func (n NSPathControl) SetPlaceholderAttributedString(value foundation.AttributedString) {
	C.C_NSPathControl_SetPlaceholderAttributedString(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSPathControl) PlaceholderString() string {
	result_ := C.C_NSPathControl_PlaceholderString(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSPathControl) SetPlaceholderString(value string) {
	C.C_NSPathControl_SetPlaceholderString(n.Ptr(), foundation.NewString(value).Ptr())
}
