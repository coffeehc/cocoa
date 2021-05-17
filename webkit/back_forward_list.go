package webkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework WebKit
// #include "back_forward_list.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type BackForwardList interface {
	objc.Object
	ItemAtIndex(index int) BackForwardListItem
	BackItem() BackForwardListItem
	CurrentItem() BackForwardListItem
	ForwardItem() BackForwardListItem
	BackList() []BackForwardListItem
	ForwardList() []BackForwardListItem
}

type WKBackForwardList struct {
	objc.NSObject
}

func MakeBackForwardList(ptr unsafe.Pointer) *WKBackForwardList {
	if ptr == nil {
		return nil
	}
	return &WKBackForwardList{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocBackForwardList() *WKBackForwardList {
	return MakeBackForwardList(C.C_BackForwardList_Alloc())
}

func (w *WKBackForwardList) ItemAtIndex(index int) BackForwardListItem {
	result_ := C.C_WKBackForwardList_ItemAtIndex(w.Ptr(), C.int(index))
	return MakeBackForwardListItem(result_)
}

func (w *WKBackForwardList) BackItem() BackForwardListItem {
	result_ := C.C_WKBackForwardList_BackItem(w.Ptr())
	return MakeBackForwardListItem(result_)
}

func (w *WKBackForwardList) CurrentItem() BackForwardListItem {
	result_ := C.C_WKBackForwardList_CurrentItem(w.Ptr())
	return MakeBackForwardListItem(result_)
}

func (w *WKBackForwardList) ForwardItem() BackForwardListItem {
	result_ := C.C_WKBackForwardList_ForwardItem(w.Ptr())
	return MakeBackForwardListItem(result_)
}

func (w *WKBackForwardList) BackList() []BackForwardListItem {
	result_ := C.C_WKBackForwardList_BackList(w.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]BackForwardListItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeBackForwardListItem(r)
	}
	return goResult_
}

func (w *WKBackForwardList) ForwardList() []BackForwardListItem {
	result_ := C.C_WKBackForwardList_ForwardList(w.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]BackForwardListItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeBackForwardListItem(r)
	}
	return goResult_
}
