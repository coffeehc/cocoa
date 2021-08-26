package appkit

// #include "color_list.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ColorList interface {
	objc.Object
	ColorWithKey(key ColorName) Color
	InsertColor_Key_AtIndex(color Color, key ColorName, loc uint)
	RemoveColorWithKey(key ColorName)
	SetColor_ForKey(color Color, key ColorName)
	RemoveFile()
	Name() ColorListName
	IsEditable() bool
	AllKeys() []ColorName
}

type NSColorList struct {
	objc.NSObject
}

func MakeColorList(ptr unsafe.Pointer) NSColorList {
	return NSColorList{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocColorList() NSColorList {
	return MakeColorList(C.C_ColorList_Alloc())
}

func (n NSColorList) InitWithName(name ColorListName) ColorList {
	result_ := C.C_NSColorList_InitWithName(n.Ptr(), foundation.NewString(string(name)).Ptr())
	return MakeColorList(result_)
}

func (n NSColorList) InitWithName_FromFile(name ColorListName, path string) ColorList {
	result_ := C.C_NSColorList_InitWithName_FromFile(n.Ptr(), foundation.NewString(string(name)).Ptr(), foundation.NewString(path).Ptr())
	return MakeColorList(result_)
}

func (n NSColorList) Init() ColorList {
	result_ := C.C_NSColorList_Init(n.Ptr())
	return MakeColorList(result_)
}

func ColorListNamed(name ColorListName) ColorList {
	result_ := C.C_NSColorList_ColorListNamed(foundation.NewString(string(name)).Ptr())
	return MakeColorList(result_)
}

func (n NSColorList) ColorWithKey(key ColorName) Color {
	result_ := C.C_NSColorList_ColorWithKey(n.Ptr(), foundation.NewString(string(key)).Ptr())
	return MakeColor(result_)
}

func (n NSColorList) InsertColor_Key_AtIndex(color Color, key ColorName, loc uint) {
	C.C_NSColorList_InsertColor_Key_AtIndex(n.Ptr(), objc.ExtractPtr(color), foundation.NewString(string(key)).Ptr(), C.uint(loc))
}

func (n NSColorList) RemoveColorWithKey(key ColorName) {
	C.C_NSColorList_RemoveColorWithKey(n.Ptr(), foundation.NewString(string(key)).Ptr())
}

func (n NSColorList) SetColor_ForKey(color Color, key ColorName) {
	C.C_NSColorList_SetColor_ForKey(n.Ptr(), objc.ExtractPtr(color), foundation.NewString(string(key)).Ptr())
}

func (n NSColorList) RemoveFile() {
	C.C_NSColorList_RemoveFile(n.Ptr())
}

func ColorList_AvailableColorLists() []ColorList {
	result_ := C.C_NSColorList_ColorList_AvailableColorLists()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]ColorList, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeColorList(r)
	}
	return goResult_
}

func (n NSColorList) Name() ColorListName {
	result_ := C.C_NSColorList_Name(n.Ptr())
	return ColorListName(foundation.MakeString(result_).String())
}

func (n NSColorList) IsEditable() bool {
	result_ := C.C_NSColorList_IsEditable(n.Ptr())
	return bool(result_)
}

func (n NSColorList) AllKeys() []ColorName {
	result_ := C.C_NSColorList_AllKeys(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]ColorName, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = ColorName(foundation.MakeString(r).String())
	}
	return goResult_
}
