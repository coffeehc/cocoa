package appkit

// #include "pasteboard_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type PasteboardItem interface {
	objc.Object
	AvailableTypeFromArray(types []PasteboardType) PasteboardType
	SetDataProvider_ForTypes(dataProvider objc.Object, types []PasteboardType) bool
	SetData_ForType(data []byte, _type PasteboardType) bool
	SetString_ForType(_string string, _type PasteboardType) bool
	SetPropertyList_ForType(propertyList objc.Object, _type PasteboardType) bool
	DataForType(_type PasteboardType) []byte
	StringForType(_type PasteboardType) string
	PropertyListForType(_type PasteboardType) objc.Object
	Types() []PasteboardType
}

type NSPasteboardItem struct {
	objc.NSObject
}

func MakePasteboardItem(ptr unsafe.Pointer) NSPasteboardItem {
	return NSPasteboardItem{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocPasteboardItem() NSPasteboardItem {
	result_ := C.C_NSPasteboardItem_AllocPasteboardItem()
	return MakePasteboardItem(result_)
}

func (n NSPasteboardItem) Init() NSPasteboardItem {
	result_ := C.C_NSPasteboardItem_Init(n.Ptr())
	return MakePasteboardItem(result_)
}

func NewPasteboardItem() NSPasteboardItem {
	result_ := C.C_NSPasteboardItem_NewPasteboardItem()
	return MakePasteboardItem(result_)
}

func (n NSPasteboardItem) Autorelease() NSPasteboardItem {
	result_ := C.C_NSPasteboardItem_Autorelease(n.Ptr())
	return MakePasteboardItem(result_)
}

func (n NSPasteboardItem) Retain() NSPasteboardItem {
	result_ := C.C_NSPasteboardItem_Retain(n.Ptr())
	return MakePasteboardItem(result_)
}

func (n NSPasteboardItem) AvailableTypeFromArray(types []PasteboardType) PasteboardType {
	var cTypes C.Array
	if len(types) > 0 {
		cTypesData := make([]unsafe.Pointer, len(types))
		for idx, v := range types {
			cTypesData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cTypes.data = unsafe.Pointer(&cTypesData[0])
		cTypes.len = C.int(len(types))
	}
	result_ := C.C_NSPasteboardItem_AvailableTypeFromArray(n.Ptr(), cTypes)
	return PasteboardType(foundation.MakeString(result_).String())
}

func (n NSPasteboardItem) SetDataProvider_ForTypes(dataProvider objc.Object, types []PasteboardType) bool {
	var cTypes C.Array
	if len(types) > 0 {
		cTypesData := make([]unsafe.Pointer, len(types))
		for idx, v := range types {
			cTypesData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cTypes.data = unsafe.Pointer(&cTypesData[0])
		cTypes.len = C.int(len(types))
	}
	result_ := C.C_NSPasteboardItem_SetDataProvider_ForTypes(n.Ptr(), objc.ExtractPtr(dataProvider), cTypes)
	return bool(result_)
}

func (n NSPasteboardItem) SetData_ForType(data []byte, _type PasteboardType) bool {
	result_ := C.C_NSPasteboardItem_SetData_ForType(n.Ptr(), foundation.NewData(data).Ptr(), foundation.NewString(string(_type)).Ptr())
	return bool(result_)
}

func (n NSPasteboardItem) SetString_ForType(_string string, _type PasteboardType) bool {
	result_ := C.C_NSPasteboardItem_SetString_ForType(n.Ptr(), foundation.NewString(_string).Ptr(), foundation.NewString(string(_type)).Ptr())
	return bool(result_)
}

func (n NSPasteboardItem) SetPropertyList_ForType(propertyList objc.Object, _type PasteboardType) bool {
	result_ := C.C_NSPasteboardItem_SetPropertyList_ForType(n.Ptr(), objc.ExtractPtr(propertyList), foundation.NewString(string(_type)).Ptr())
	return bool(result_)
}

func (n NSPasteboardItem) DataForType(_type PasteboardType) []byte {
	result_ := C.C_NSPasteboardItem_DataForType(n.Ptr(), foundation.NewString(string(_type)).Ptr())
	return foundation.MakeData(result_).ToBytes()
}

func (n NSPasteboardItem) StringForType(_type PasteboardType) string {
	result_ := C.C_NSPasteboardItem_StringForType(n.Ptr(), foundation.NewString(string(_type)).Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSPasteboardItem) PropertyListForType(_type PasteboardType) objc.Object {
	result_ := C.C_NSPasteboardItem_PropertyListForType(n.Ptr(), foundation.NewString(string(_type)).Ptr())
	return objc.MakeObject(result_)
}

func (n NSPasteboardItem) Types() []PasteboardType {
	result_ := C.C_NSPasteboardItem_Types(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]PasteboardType, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = PasteboardType(foundation.MakeString(r).String())
	}
	return goResult_
}
