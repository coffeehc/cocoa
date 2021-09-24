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

func (n NSDraggingItem) InitWithPasteboardWriter(pasteboardWriter objc.Object) NSDraggingItem {
	result_ := C.C_NSDraggingItem_InitWithPasteboardWriter(n.Ptr(), objc.ExtractPtr(pasteboardWriter))
	return MakeDraggingItem(result_)
}

func AllocDraggingItem() NSDraggingItem {
	result_ := C.C_NSDraggingItem_AllocDraggingItem()
	return MakeDraggingItem(result_)
}

func (n NSDraggingItem) Autorelease() NSDraggingItem {
	result_ := C.C_NSDraggingItem_Autorelease(n.Ptr())
	return MakeDraggingItem(result_)
}

func (n NSDraggingItem) Retain() NSDraggingItem {
	result_ := C.C_NSDraggingItem_Retain(n.Ptr())
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
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
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

type DraggingImageComponent interface {
	objc.Object
	Key() DraggingImageComponentKey
	SetKey(value DraggingImageComponentKey)
	Contents() objc.Object
	SetContents(value objc.Object)
	Frame() foundation.Rect
	SetFrame(value foundation.Rect)
}

type NSDraggingImageComponent struct {
	objc.NSObject
}

func MakeDraggingImageComponent(ptr unsafe.Pointer) NSDraggingImageComponent {
	return NSDraggingImageComponent{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSDraggingImageComponent) InitWithKey(key DraggingImageComponentKey) NSDraggingImageComponent {
	result_ := C.C_NSDraggingImageComponent_InitWithKey(n.Ptr(), foundation.NewString(string(key)).Ptr())
	return MakeDraggingImageComponent(result_)
}

func AllocDraggingImageComponent() NSDraggingImageComponent {
	result_ := C.C_NSDraggingImageComponent_AllocDraggingImageComponent()
	return MakeDraggingImageComponent(result_)
}

func (n NSDraggingImageComponent) Autorelease() NSDraggingImageComponent {
	result_ := C.C_NSDraggingImageComponent_Autorelease(n.Ptr())
	return MakeDraggingImageComponent(result_)
}

func (n NSDraggingImageComponent) Retain() NSDraggingImageComponent {
	result_ := C.C_NSDraggingImageComponent_Retain(n.Ptr())
	return MakeDraggingImageComponent(result_)
}

func DraggingImageComponentWithKey(key DraggingImageComponentKey) DraggingImageComponent {
	result_ := C.C_NSDraggingImageComponent_DraggingImageComponentWithKey(foundation.NewString(string(key)).Ptr())
	return MakeDraggingImageComponent(result_)
}

func (n NSDraggingImageComponent) Key() DraggingImageComponentKey {
	result_ := C.C_NSDraggingImageComponent_Key(n.Ptr())
	return DraggingImageComponentKey(foundation.MakeString(result_).String())
}

func (n NSDraggingImageComponent) SetKey(value DraggingImageComponentKey) {
	C.C_NSDraggingImageComponent_SetKey(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n NSDraggingImageComponent) Contents() objc.Object {
	result_ := C.C_NSDraggingImageComponent_Contents(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSDraggingImageComponent) SetContents(value objc.Object) {
	C.C_NSDraggingImageComponent_SetContents(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSDraggingImageComponent) Frame() foundation.Rect {
	result_ := C.C_NSDraggingImageComponent_Frame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n NSDraggingImageComponent) SetFrame(value foundation.Rect) {
	C.C_NSDraggingImageComponent_SetFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(value))))
}
