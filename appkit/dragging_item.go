package appkit

// #include "dragging_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type DraggingItem interface {
	objc.Object
	SetDraggingFrame_Contents(frame foundation.Rect, contents objc.Object)
	DraggingFrame() foundation.Rect
	SetDraggingFrame(value foundation.Rect)
	ImageComponents() []DraggingImageComponent
	Item() objc.Object
}

type NSDraggingItem struct {
	objc.NSObject
}

func MakeDraggingItem(ptr unsafe.Pointer) NSDraggingItem {
	return NSDraggingItem{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocDraggingItem() NSDraggingItem {
	return MakeDraggingItem(C.C_DraggingItem_Alloc())
}

func (n NSDraggingItem) InitWithPasteboardWriter(pasteboardWriter objc.Object) DraggingItem {
	result_ := C.C_NSDraggingItem_InitWithPasteboardWriter(n.Ptr(), objc.ExtractPtr(pasteboardWriter))
	return MakeDraggingItem(result_)
}

func (n NSDraggingItem) SetDraggingFrame_Contents(frame foundation.Rect, contents objc.Object) {
	C.C_NSDraggingItem_SetDraggingFrame_Contents(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frame))), objc.ExtractPtr(contents))
}

func (n NSDraggingItem) DraggingFrame() foundation.Rect {
	result_ := C.C_NSDraggingItem_DraggingFrame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSDraggingItem) SetDraggingFrame(value foundation.Rect) {
	C.C_NSDraggingItem_SetDraggingFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(value))))
}

func (n NSDraggingItem) ImageComponents() []DraggingImageComponent {
	result_ := C.C_NSDraggingItem_ImageComponents(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]DraggingImageComponent, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeDraggingImageComponent(r)
	}
	return goResult_
}

func (n NSDraggingItem) Item() objc.Object {
	result_ := C.C_NSDraggingItem_Item(n.Ptr())
	return objc.MakeObject(result_)
}
