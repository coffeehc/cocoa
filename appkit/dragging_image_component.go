package appkit

// #include "dragging_image_component.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

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

func MakeDraggingImageComponent(ptr unsafe.Pointer) *NSDraggingImageComponent {
	if ptr == nil {
		return nil
	}
	return &NSDraggingImageComponent{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocDraggingImageComponent() *NSDraggingImageComponent {
	return MakeDraggingImageComponent(C.C_DraggingImageComponent_Alloc())
}

func (n *NSDraggingImageComponent) InitWithKey(key DraggingImageComponentKey) DraggingImageComponent {
	result_ := C.C_NSDraggingImageComponent_InitWithKey(n.Ptr(), foundation.NewString(string(key)).Ptr())
	return MakeDraggingImageComponent(result_)
}

func DraggingImageComponentWithKey(key DraggingImageComponentKey) DraggingImageComponent {
	result_ := C.C_NSDraggingImageComponent_DraggingImageComponentWithKey(foundation.NewString(string(key)).Ptr())
	return MakeDraggingImageComponent(result_)
}

func (n *NSDraggingImageComponent) Key() DraggingImageComponentKey {
	result_ := C.C_NSDraggingImageComponent_Key(n.Ptr())
	return DraggingImageComponentKey(foundation.MakeString(result_).String())
}

func (n *NSDraggingImageComponent) SetKey(value DraggingImageComponentKey) {
	C.C_NSDraggingImageComponent_SetKey(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n *NSDraggingImageComponent) Contents() objc.Object {
	result_ := C.C_NSDraggingImageComponent_Contents(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSDraggingImageComponent) SetContents(value objc.Object) {
	C.C_NSDraggingImageComponent_SetContents(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSDraggingImageComponent) Frame() foundation.Rect {
	result_ := C.C_NSDraggingImageComponent_Frame(n.Ptr())
	return foundation.Rect(coregraphics.FromCGRectPointer(unsafe.Pointer(&result_)))
}

func (n *NSDraggingImageComponent) SetFrame(value foundation.Rect) {
	C.C_NSDraggingImageComponent_SetFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(value))))
}
