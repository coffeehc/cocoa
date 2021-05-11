package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "pasteboard.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Pasteboard interface {
	objc.Object
	ClearContents() int
	SetData_ForType(data []byte, dataType PasteboardType) bool
	SetPropertyList_ForType(plist objc.Object, dataType PasteboardType) bool
	SetString_ForType(_string string, dataType PasteboardType) bool
	IndexOfPasteboardItem(pasteboardItem PasteboardItem) uint
	DataForType(dataType PasteboardType) []byte
	PropertyListForType(dataType PasteboardType) objc.Object
	StringForType(dataType PasteboardType) string
	WriteFileContents(filename string) bool
	WriteFileWrapper(wrapper foundation.FileWrapper) bool
	ReadFileContentsType_ToFile(_type PasteboardType, filename string) string
	ReadFileWrapper() foundation.FileWrapper
	Name() PasteboardName
	ChangeCount() int
}

type NSPasteboard struct {
	objc.NSObject
}

func MakePasteboard(ptr unsafe.Pointer) *NSPasteboard {
	if ptr == nil {
		return nil
	}
	return &NSPasteboard{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocPasteboard() *NSPasteboard {
	return MakePasteboard(C.C_Pasteboard_Alloc())
}

func (n *NSPasteboard) Init() Pasteboard {
	result := C.C_NSPasteboard_Init(n.Ptr())
	return MakePasteboard(result)
}

func PasteboardByFilteringData_OfType(data []byte, _type PasteboardType) Pasteboard {
	result := C.C_NSPasteboard_PasteboardByFilteringData_OfType(C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}, foundation.NewString(string(_type)).Ptr())
	return MakePasteboard(result)
}

func PasteboardByFilteringFile(filename string) Pasteboard {
	result := C.C_NSPasteboard_PasteboardByFilteringFile(foundation.NewString(filename).Ptr())
	return MakePasteboard(result)
}

func PasteboardByFilteringTypesInPasteboard(pboard Pasteboard) Pasteboard {
	result := C.C_NSPasteboard_PasteboardByFilteringTypesInPasteboard(objc.ExtractPtr(pboard))
	return MakePasteboard(result)
}

func PasteboardWithName(name PasteboardName) Pasteboard {
	result := C.C_NSPasteboard_PasteboardWithName(foundation.NewString(string(name)).Ptr())
	return MakePasteboard(result)
}

func PasteboardWithUniqueName() Pasteboard {
	result := C.C_NSPasteboard_PasteboardWithUniqueName()
	return MakePasteboard(result)
}

func (n *NSPasteboard) ClearContents() int {
	result := C.C_NSPasteboard_ClearContents(n.Ptr())
	return int(result)
}

func (n *NSPasteboard) SetData_ForType(data []byte, dataType PasteboardType) bool {
	result := C.C_NSPasteboard_SetData_ForType(n.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}, foundation.NewString(string(dataType)).Ptr())
	return bool(result)
}

func (n *NSPasteboard) SetPropertyList_ForType(plist objc.Object, dataType PasteboardType) bool {
	result := C.C_NSPasteboard_SetPropertyList_ForType(n.Ptr(), objc.ExtractPtr(plist), foundation.NewString(string(dataType)).Ptr())
	return bool(result)
}

func (n *NSPasteboard) SetString_ForType(_string string, dataType PasteboardType) bool {
	result := C.C_NSPasteboard_SetString_ForType(n.Ptr(), foundation.NewString(_string).Ptr(), foundation.NewString(string(dataType)).Ptr())
	return bool(result)
}

func (n *NSPasteboard) IndexOfPasteboardItem(pasteboardItem PasteboardItem) uint {
	result := C.C_NSPasteboard_IndexOfPasteboardItem(n.Ptr(), objc.ExtractPtr(pasteboardItem))
	return uint(result)
}

func (n *NSPasteboard) DataForType(dataType PasteboardType) []byte {
	result := C.C_NSPasteboard_DataForType(n.Ptr(), foundation.NewString(string(dataType)).Ptr())
	resultBuffer := (*[1 << 30]byte)(result.data)[:C.int(result.len)]
	goResult := make([]byte, C.int(result.len))
	copy(goResult, resultBuffer)
	C.free(result.data)
	return goResult
}

func (n *NSPasteboard) PropertyListForType(dataType PasteboardType) objc.Object {
	result := C.C_NSPasteboard_PropertyListForType(n.Ptr(), foundation.NewString(string(dataType)).Ptr())
	return objc.MakeObject(result)
}

func (n *NSPasteboard) StringForType(dataType PasteboardType) string {
	result := C.C_NSPasteboard_StringForType(n.Ptr(), foundation.NewString(string(dataType)).Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSPasteboard) WriteFileContents(filename string) bool {
	result := C.C_NSPasteboard_WriteFileContents(n.Ptr(), foundation.NewString(filename).Ptr())
	return bool(result)
}

func (n *NSPasteboard) WriteFileWrapper(wrapper foundation.FileWrapper) bool {
	result := C.C_NSPasteboard_WriteFileWrapper(n.Ptr(), objc.ExtractPtr(wrapper))
	return bool(result)
}

func (n *NSPasteboard) ReadFileContentsType_ToFile(_type PasteboardType, filename string) string {
	result := C.C_NSPasteboard_ReadFileContentsType_ToFile(n.Ptr(), foundation.NewString(string(_type)).Ptr(), foundation.NewString(filename).Ptr())
	return foundation.MakeString(result).String()
}

func (n *NSPasteboard) ReadFileWrapper() foundation.FileWrapper {
	result := C.C_NSPasteboard_ReadFileWrapper(n.Ptr())
	return foundation.MakeFileWrapper(result)
}

func (n *NSPasteboard) Name() PasteboardName {
	result := C.C_NSPasteboard_Name(n.Ptr())
	return PasteboardName(foundation.MakeString(result).String())
}

func (n *NSPasteboard) ChangeCount() int {
	result := C.C_NSPasteboard_ChangeCount(n.Ptr())
	return int(result)
}
