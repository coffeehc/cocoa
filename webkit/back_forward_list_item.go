package webkit

// #include "back_forward_list_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type BackForwardListItem interface {
	objc.Object
	Title() string
	URL() foundation.URL
	InitialURL() foundation.URL
}

type WKBackForwardListItem struct {
	objc.NSObject
}

func MakeBackForwardListItem(ptr unsafe.Pointer) *WKBackForwardListItem {
	if ptr == nil {
		return nil
	}
	return &WKBackForwardListItem{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocBackForwardListItem() *WKBackForwardListItem {
	return MakeBackForwardListItem(C.C_BackForwardListItem_Alloc())
}

func (w *WKBackForwardListItem) Title() string {
	result_ := C.C_WKBackForwardListItem_Title(w.Ptr())
	return foundation.MakeString(result_).String()
}

func (w *WKBackForwardListItem) URL() foundation.URL {
	result_ := C.C_WKBackForwardListItem_URL(w.Ptr())
	return foundation.MakeURL(result_)
}

func (w *WKBackForwardListItem) InitialURL() foundation.URL {
	result_ := C.C_WKBackForwardListItem_InitialURL(w.Ptr())
	return foundation.MakeURL(result_)
}
