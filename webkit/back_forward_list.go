package webkit

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

func MakeBackForwardList(ptr unsafe.Pointer) WKBackForwardList {
	return WKBackForwardList{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocBackForwardList() WKBackForwardList {
	result_ := C.C_WKBackForwardList_AllocBackForwardList()
	return MakeBackForwardList(result_)
}

func (w WKBackForwardList) Autorelease() WKBackForwardList {
	result_ := C.C_WKBackForwardList_Autorelease(w.Ptr())
	return MakeBackForwardList(result_)
}

func (w WKBackForwardList) Retain() WKBackForwardList {
	result_ := C.C_WKBackForwardList_Retain(w.Ptr())
	return MakeBackForwardList(result_)
}

func (w WKBackForwardList) ItemAtIndex(index int) BackForwardListItem {
	result_ := C.C_WKBackForwardList_ItemAtIndex(w.Ptr(), C.int(index))
	return MakeBackForwardListItem(result_)
}

func (w WKBackForwardList) BackItem() BackForwardListItem {
	result_ := C.C_WKBackForwardList_BackItem(w.Ptr())
	return MakeBackForwardListItem(result_)
}

func (w WKBackForwardList) CurrentItem() BackForwardListItem {
	result_ := C.C_WKBackForwardList_CurrentItem(w.Ptr())
	return MakeBackForwardListItem(result_)
}

func (w WKBackForwardList) ForwardItem() BackForwardListItem {
	result_ := C.C_WKBackForwardList_ForwardItem(w.Ptr())
	return MakeBackForwardListItem(result_)
}

func (w WKBackForwardList) BackList() []BackForwardListItem {
	result_ := C.C_WKBackForwardList_BackList(w.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]BackForwardListItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeBackForwardListItem(r)
	}
	return goResult_
}

func (w WKBackForwardList) ForwardList() []BackForwardListItem {
	result_ := C.C_WKBackForwardList_ForwardList(w.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]BackForwardListItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeBackForwardListItem(r)
	}
	return goResult_
}
