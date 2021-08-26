package appkit

// #include "sharing_service_picker.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type SharingServicePicker interface {
	objc.Object
	ShowRelativeToRect_OfView_PreferredEdge(rect foundation.Rect, view View, preferredEdge foundation.RectEdge)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
}

type NSSharingServicePicker struct {
	objc.NSObject
}

func MakeSharingServicePicker(ptr unsafe.Pointer) NSSharingServicePicker {
	return NSSharingServicePicker{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocSharingServicePicker() NSSharingServicePicker {
	return MakeSharingServicePicker(C.C_SharingServicePicker_Alloc())
}

func (n NSSharingServicePicker) InitWithItems(items []objc.Object) SharingServicePicker {
	var cItems C.Array
	if len(items) > 0 {
		cItemsData := make([]unsafe.Pointer, len(items))
		for idx, v := range items {
			cItemsData[idx] = objc.ExtractPtr(v)
		}
		cItems.data = unsafe.Pointer(&cItemsData[0])
		cItems.len = C.int(len(items))
	}
	result_ := C.C_NSSharingServicePicker_InitWithItems(n.Ptr(), cItems)
	return MakeSharingServicePicker(result_)
}

func (n NSSharingServicePicker) ShowRelativeToRect_OfView_PreferredEdge(rect foundation.Rect, view View, preferredEdge foundation.RectEdge) {
	C.C_NSSharingServicePicker_ShowRelativeToRect_OfView_PreferredEdge(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(rect))), objc.ExtractPtr(view), C.uint(uint(preferredEdge)))
}

func (n NSSharingServicePicker) Delegate() objc.Object {
	result_ := C.C_NSSharingServicePicker_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSSharingServicePicker) SetDelegate(value objc.Object) {
	C.C_NSSharingServicePicker_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}
