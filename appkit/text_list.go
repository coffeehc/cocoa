package appkit

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

func MakeTextList(ptr unsafe.Pointer) NSTextList {
	return NSTextList{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocTextList() NSTextList {
	return MakeTextList(C.C_TextList_Alloc())
}

func (n NSTextList) InitWithMarkerFormat_Options(format TextListMarkerFormat, mask uint) TextList {
	result_ := C.C_NSTextList_InitWithMarkerFormat_Options(n.Ptr(), foundation.NewString(string(format)).Ptr(), C.uint(mask))
	return MakeTextList(result_)
}

func (n NSTextList) Init() TextList {
	result_ := C.C_NSTextList_Init(n.Ptr())
	return MakeTextList(result_)
}

func (n NSTextList) MarkerForItemNumber(itemNum int) string {
	result_ := C.C_NSTextList_MarkerForItemNumber(n.Ptr(), C.int(itemNum))
	return foundation.MakeString(result_).String()
}

func (n NSTextList) MarkerFormat() TextListMarkerFormat {
	result_ := C.C_NSTextList_MarkerFormat(n.Ptr())
	return TextListMarkerFormat(foundation.MakeString(result_).String())
}

func (n NSTextList) ListOptions() TextListOptions {
	result_ := C.C_NSTextList_ListOptions(n.Ptr())
	return TextListOptions(uint(result_))
}

func (n NSTextList) StartingItemNumber() int {
	result_ := C.C_NSTextList_StartingItemNumber(n.Ptr())
	return int(result_)
}

func (n NSTextList) SetStartingItemNumber(value int) {
	C.C_NSTextList_SetStartingItemNumber(n.Ptr(), C.int(value))
}
