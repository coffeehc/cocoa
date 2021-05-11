package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "pasteboard_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type PasteboardItem interface {
	objc.Object
	SetData_ForType(data []byte, _type PasteboardType) bool
	SetString_ForType(_string string, _type PasteboardType) bool
	SetPropertyList_ForType(propertyList objc.Object, _type PasteboardType) bool
	DataForType(_type PasteboardType) []byte
	StringForType(_type PasteboardType) string
	PropertyListForType(_type PasteboardType) objc.Object
}

type NSPasteboardItem struct {
	objc.NSObject
}

func MakePasteboardItem(ptr unsafe.Pointer) *NSPasteboardItem {
	if ptr == nil {
		return nil
	}
	return &NSPasteboardItem{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocPasteboardItem() *NSPasteboardItem {
	return MakePasteboardItem(C.C_PasteboardItem_Alloc())
}

func (n *NSPasteboardItem) Init() PasteboardItem {
	result := C.C_NSPasteboardItem_Init(n.Ptr())
	return MakePasteboardItem(result)
}

func (n *NSPasteboardItem) SetData_ForType(data []byte, _type PasteboardType) bool {
	result := C.C_NSPasteboardItem_SetData_ForType(n.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}, foundation.NewString(string(_type)).Ptr())
	return bool(result)
}

func (n *NSPasteboardItem) SetString_ForType(_string string, _type PasteboardType) bool {
	result := C.C_NSPasteboardItem_SetString_ForType(n.Ptr(), foundation.NewString(_string).Ptr(), foundation.NewString(string(_type)).Ptr())
	return bool(result)
}

func (n *NSPasteboardItem) SetPropertyList_ForType(propertyList objc.Object, _type PasteboardType) bool {
	result := C.C_NSPasteboardItem_SetPropertyList_ForType(n.Ptr(), objc.ExtractPtr(propertyList), foundation.NewString(string(_type)).Ptr())
	return bool(result)
}

func (n *NSPasteboardItem) DataForType(_type PasteboardType) []byte {
	result := C.C_NSPasteboardItem_DataForType(n.Ptr(), foundation.NewString(string(_type)).Ptr())
	resultBuffer := (*[1 << 30]byte)(result.data)[:C.int(result.len)]
	goResult := make([]byte, C.int(result.len))
	copy(goResult, resultBuffer)
	C.free(result.data)
	return goResult
}

func (n *NSPasteboardItem) StringForType(_type PasteboardType) string {
	result := C.C_NSPasteboardItem_StringForType(n.Ptr(), foundation.NewString(string(_type)).Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSPasteboardItem) PropertyListForType(_type PasteboardType) objc.Object {
	result := C.C_NSPasteboardItem_PropertyListForType(n.Ptr(), foundation.NewString(string(_type)).Ptr())
	return objc.MakeObject(result)
}
