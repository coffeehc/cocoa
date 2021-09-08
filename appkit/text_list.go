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

func (n NSTextList) InitWithMarkerFormat_Options(format TextListMarkerFormat, mask uint) NSTextList {
	result_ := C.C_NSTextList_InitWithMarkerFormat_Options(n.Ptr(), foundation.NewString(string(format)).Ptr(), C.uint(mask))
	return MakeTextList(result_)
}

func AllocTextList() NSTextList {
	result_ := C.C_NSTextList_AllocTextList()
	return MakeTextList(result_)
}

func (n NSTextList) Init() NSTextList {
	result_ := C.C_NSTextList_Init(n.Ptr())
	return MakeTextList(result_)
}

func NewTextList() NSTextList {
	result_ := C.C_NSTextList_NewTextList()
	return MakeTextList(result_)
}

func (n NSTextList) Autorelease() NSTextList {
	result_ := C.C_NSTextList_Autorelease(n.Ptr())
	return MakeTextList(result_)
}

func (n NSTextList) Retain() NSTextList {
	result_ := C.C_NSTextList_Retain(n.Ptr())
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
