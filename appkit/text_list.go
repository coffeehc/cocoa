package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "text_list.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TextList interface {
	objc.Object
	MarkerForItemNumber(itemNum int) string
	MarkerFormat() TextListMarkerFormat
	ListOptions() TextListOptions
	StartingItemNumber() int
	SetStartingItemNumber(value int)
}

type NSTextList struct {
	objc.NSObject
}

func MakeTextList(ptr unsafe.Pointer) *NSTextList {
	if ptr == nil {
		return nil
	}
	return &NSTextList{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocTextList() *NSTextList {
	return MakeTextList(C.C_TextList_Alloc())
}

func (n *NSTextList) InitWithMarkerFormat_Options(format TextListMarkerFormat, mask uint) TextList {
	result := C.C_NSTextList_InitWithMarkerFormat_Options(n.Ptr(), foundation.NewString(string(format)).Ptr(), C.uint(mask))
	return MakeTextList(result)
}

func (n *NSTextList) Init() TextList {
	result := C.C_NSTextList_Init(n.Ptr())
	return MakeTextList(result)
}

func (n *NSTextList) MarkerForItemNumber(itemNum int) string {
	result := C.C_NSTextList_MarkerForItemNumber(n.Ptr(), C.int(itemNum))
	return foundation.MakeString(result).String()
}

func (n *NSTextList) MarkerFormat() TextListMarkerFormat {
	result := C.C_NSTextList_MarkerFormat(n.Ptr())
	return TextListMarkerFormat(foundation.MakeString(result).String())
}

func (n *NSTextList) ListOptions() TextListOptions {
	result := C.C_NSTextList_ListOptions(n.Ptr())
	return TextListOptions(uint(result))
}

func (n *NSTextList) StartingItemNumber() int {
	result := C.C_NSTextList_StartingItemNumber(n.Ptr())
	return int(result)
}

func (n *NSTextList) SetStartingItemNumber(value int) {
	C.C_NSTextList_SetStartingItemNumber(n.Ptr(), C.int(value))
}
