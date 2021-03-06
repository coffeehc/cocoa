package appkit

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
	AvailableTypeFromArray(types []PasteboardType) PasteboardType
	CanReadItemWithDataConformingToTypes(types []string) bool
	PrepareForNewContentsWithOptions(options PasteboardContentsOptions) int
	DeclareTypes_Owner(newTypes []PasteboardType, newOwner objc.Object) int
	AddTypes_Owner(newTypes []PasteboardType, newOwner objc.Object) int
	WriteFileContents(filename string) bool
	WriteFileWrapper(wrapper foundation.FileWrapper) bool
	ReadFileContentsType_ToFile(_type PasteboardType, filename string) string
	ReadFileWrapper() foundation.FileWrapper
	PasteboardItems() []PasteboardItem
	Types() []PasteboardType
	Name() PasteboardName
	ChangeCount() int
}

type NSPasteboard struct {
	objc.NSObject
}

func MakePasteboard(ptr unsafe.Pointer) NSPasteboard {
	return NSPasteboard{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocPasteboard() NSPasteboard {
	result_ := C.C_NSPasteboard_AllocPasteboard()
	return MakePasteboard(result_)
}

func (n NSPasteboard) Init() NSPasteboard {
	result_ := C.C_NSPasteboard_Init(n.Ptr())
	return MakePasteboard(result_)
}

func NewPasteboard() NSPasteboard {
	result_ := C.C_NSPasteboard_NewPasteboard()
	return MakePasteboard(result_)
}

func (n NSPasteboard) Autorelease() NSPasteboard {
	result_ := C.C_NSPasteboard_Autorelease(n.Ptr())
	return MakePasteboard(result_)
}

func (n NSPasteboard) Retain() NSPasteboard {
	result_ := C.C_NSPasteboard_Retain(n.Ptr())
	return MakePasteboard(result_)
}

func PasteboardByFilteringData_OfType(data []byte, _type PasteboardType) Pasteboard {
	result_ := C.C_NSPasteboard_PasteboardByFilteringData_OfType(foundation.NewData(data).Ptr(), foundation.NewString(string(_type)).Ptr())
	return MakePasteboard(result_)
}

func PasteboardByFilteringFile(filename string) Pasteboard {
	result_ := C.C_NSPasteboard_PasteboardByFilteringFile(foundation.NewString(filename).Ptr())
	return MakePasteboard(result_)
}

func PasteboardByFilteringTypesInPasteboard(pboard Pasteboard) Pasteboard {
	result_ := C.C_NSPasteboard_PasteboardByFilteringTypesInPasteboard(objc.ExtractPtr(pboard))
	return MakePasteboard(result_)
}

func PasteboardWithName(name PasteboardName) Pasteboard {
	result_ := C.C_NSPasteboard_PasteboardWithName(foundation.NewString(string(name)).Ptr())
	return MakePasteboard(result_)
}

func PasteboardWithUniqueName() Pasteboard {
	result_ := C.C_NSPasteboard_PasteboardWithUniqueName()
	return MakePasteboard(result_)
}

func (n NSPasteboard) ClearContents() int {
	result_ := C.C_NSPasteboard_ClearContents(n.Ptr())
	return int(result_)
}

func (n NSPasteboard) SetData_ForType(data []byte, dataType PasteboardType) bool {
	result_ := C.C_NSPasteboard_SetData_ForType(n.Ptr(), foundation.NewData(data).Ptr(), foundation.NewString(string(dataType)).Ptr())
	return bool(result_)
}

func (n NSPasteboard) SetPropertyList_ForType(plist objc.Object, dataType PasteboardType) bool {
	result_ := C.C_NSPasteboard_SetPropertyList_ForType(n.Ptr(), objc.ExtractPtr(plist), foundation.NewString(string(dataType)).Ptr())
	return bool(result_)
}

func (n NSPasteboard) SetString_ForType(_string string, dataType PasteboardType) bool {
	result_ := C.C_NSPasteboard_SetString_ForType(n.Ptr(), foundation.NewString(_string).Ptr(), foundation.NewString(string(dataType)).Ptr())
	return bool(result_)
}

func (n NSPasteboard) IndexOfPasteboardItem(pasteboardItem PasteboardItem) uint {
	result_ := C.C_NSPasteboard_IndexOfPasteboardItem(n.Ptr(), objc.ExtractPtr(pasteboardItem))
	return uint(result_)
}

func (n NSPasteboard) DataForType(dataType PasteboardType) []byte {
	result_ := C.C_NSPasteboard_DataForType(n.Ptr(), foundation.NewString(string(dataType)).Ptr())
	return foundation.MakeData(result_).ToBytes()
}

func (n NSPasteboard) PropertyListForType(dataType PasteboardType) objc.Object {
	result_ := C.C_NSPasteboard_PropertyListForType(n.Ptr(), foundation.NewString(string(dataType)).Ptr())
	return objc.MakeObject(result_)
}

func (n NSPasteboard) StringForType(dataType PasteboardType) string {
	result_ := C.C_NSPasteboard_StringForType(n.Ptr(), foundation.NewString(string(dataType)).Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSPasteboard) AvailableTypeFromArray(types []PasteboardType) PasteboardType {
	var cTypes C.Array
	if len(types) > 0 {
		cTypesData := make([]unsafe.Pointer, len(types))
		for idx, v := range types {
			cTypesData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cTypes.data = unsafe.Pointer(&cTypesData[0])
		cTypes.len = C.int(len(types))
	}
	result_ := C.C_NSPasteboard_AvailableTypeFromArray(n.Ptr(), cTypes)
	return PasteboardType(foundation.MakeString(result_).String())
}

func (n NSPasteboard) CanReadItemWithDataConformingToTypes(types []string) bool {
	var cTypes C.Array
	if len(types) > 0 {
		cTypesData := make([]unsafe.Pointer, len(types))
		for idx, v := range types {
			cTypesData[idx] = foundation.NewString(v).Ptr()
		}
		cTypes.data = unsafe.Pointer(&cTypesData[0])
		cTypes.len = C.int(len(types))
	}
	result_ := C.C_NSPasteboard_CanReadItemWithDataConformingToTypes(n.Ptr(), cTypes)
	return bool(result_)
}

func Pasteboard_TypesFilterableTo(_type PasteboardType) []PasteboardType {
	result_ := C.C_NSPasteboard_Pasteboard_TypesFilterableTo(foundation.NewString(string(_type)).Ptr())
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

func (n NSPasteboard) PrepareForNewContentsWithOptions(options PasteboardContentsOptions) int {
	result_ := C.C_NSPasteboard_PrepareForNewContentsWithOptions(n.Ptr(), C.uint(uint(options)))
	return int(result_)
}

func (n NSPasteboard) DeclareTypes_Owner(newTypes []PasteboardType, newOwner objc.Object) int {
	var cNewTypes C.Array
	if len(newTypes) > 0 {
		cNewTypesData := make([]unsafe.Pointer, len(newTypes))
		for idx, v := range newTypes {
			cNewTypesData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cNewTypes.data = unsafe.Pointer(&cNewTypesData[0])
		cNewTypes.len = C.int(len(newTypes))
	}
	result_ := C.C_NSPasteboard_DeclareTypes_Owner(n.Ptr(), cNewTypes, objc.ExtractPtr(newOwner))
	return int(result_)
}

func (n NSPasteboard) AddTypes_Owner(newTypes []PasteboardType, newOwner objc.Object) int {
	var cNewTypes C.Array
	if len(newTypes) > 0 {
		cNewTypesData := make([]unsafe.Pointer, len(newTypes))
		for idx, v := range newTypes {
			cNewTypesData[idx] = foundation.NewString(string(v)).Ptr()
		}
		cNewTypes.data = unsafe.Pointer(&cNewTypesData[0])
		cNewTypes.len = C.int(len(newTypes))
	}
	result_ := C.C_NSPasteboard_AddTypes_Owner(n.Ptr(), cNewTypes, objc.ExtractPtr(newOwner))
	return int(result_)
}

func (n NSPasteboard) WriteFileContents(filename string) bool {
	result_ := C.C_NSPasteboard_WriteFileContents(n.Ptr(), foundation.NewString(filename).Ptr())
	return bool(result_)
}

func (n NSPasteboard) WriteFileWrapper(wrapper foundation.FileWrapper) bool {
	result_ := C.C_NSPasteboard_WriteFileWrapper(n.Ptr(), objc.ExtractPtr(wrapper))
	return bool(result_)
}

func (n NSPasteboard) ReadFileContentsType_ToFile(_type PasteboardType, filename string) string {
	result_ := C.C_NSPasteboard_ReadFileContentsType_ToFile(n.Ptr(), foundation.NewString(string(_type)).Ptr(), foundation.NewString(filename).Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSPasteboard) ReadFileWrapper() foundation.FileWrapper {
	result_ := C.C_NSPasteboard_ReadFileWrapper(n.Ptr())
	return foundation.MakeFileWrapper(result_)
}

func GeneralPasteboard() Pasteboard {
	result_ := C.C_NSPasteboard_GeneralPasteboard()
	return MakePasteboard(result_)
}

func (n NSPasteboard) PasteboardItems() []PasteboardItem {
	result_ := C.C_NSPasteboard_PasteboardItems(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]PasteboardItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakePasteboardItem(r)
	}
	return goResult_
}

func (n NSPasteboard) Types() []PasteboardType {
	result_ := C.C_NSPasteboard_Types(n.Ptr())
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

func (n NSPasteboard) Name() PasteboardName {
	result_ := C.C_NSPasteboard_Name(n.Ptr())
	return PasteboardName(foundation.MakeString(result_).String())
}

func (n NSPasteboard) ChangeCount() int {
	result_ := C.C_NSPasteboard_ChangeCount(n.Ptr())
	return int(result_)
}
