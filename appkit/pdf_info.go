package appkit

// #include "pdf_info.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type PDFInfo interface {
	objc.Object
	URL() foundation.URL
	SetURL(value foundation.URL)
	IsFileExtensionHidden() bool
	SetFileExtensionHidden(value bool)
	TagNames() []string
	SetTagNames(value []string)
	Orientation() PaperOrientation
	SetOrientation(value PaperOrientation)
	PaperSize() foundation.Size
	SetPaperSize(value foundation.Size)
	Attributes() map[PrintInfoAttributeKey]objc.Object
}

type NSPDFInfo struct {
	objc.NSObject
}

func MakePDFInfo(ptr unsafe.Pointer) NSPDFInfo {
	return NSPDFInfo{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocPDFInfo() NSPDFInfo {
	result_ := C.C_NSPDFInfo_AllocPDFInfo()
	return MakePDFInfo(result_)
}

func (n NSPDFInfo) Init() NSPDFInfo {
	result_ := C.C_NSPDFInfo_Init(n.Ptr())
	return MakePDFInfo(result_)
}

func NewPDFInfo() NSPDFInfo {
	result_ := C.C_NSPDFInfo_NewPDFInfo()
	return MakePDFInfo(result_)
}

func (n NSPDFInfo) Autorelease() NSPDFInfo {
	result_ := C.C_NSPDFInfo_Autorelease(n.Ptr())
	return MakePDFInfo(result_)
}

func (n NSPDFInfo) Retain() NSPDFInfo {
	result_ := C.C_NSPDFInfo_Retain(n.Ptr())
	return MakePDFInfo(result_)
}

func (n NSPDFInfo) URL() foundation.URL {
	result_ := C.C_NSPDFInfo_URL(n.Ptr())
	return foundation.MakeURL(result_)
}

func (n NSPDFInfo) SetURL(value foundation.URL) {
	C.C_NSPDFInfo_SetURL(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSPDFInfo) IsFileExtensionHidden() bool {
	result_ := C.C_NSPDFInfo_IsFileExtensionHidden(n.Ptr())
	return bool(result_)
}

func (n NSPDFInfo) SetFileExtensionHidden(value bool) {
	C.C_NSPDFInfo_SetFileExtensionHidden(n.Ptr(), C.bool(value))
}

func (n NSPDFInfo) TagNames() []string {
	result_ := C.C_NSPDFInfo_TagNames(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeString(r).String()
	}
	return goResult_
}

func (n NSPDFInfo) SetTagNames(value []string) {
	var cValue C.Array
	if len(value) > 0 {
		cValueData := make([]unsafe.Pointer, len(value))
		for idx, v := range value {
			cValueData[idx] = foundation.NewString(v).Ptr()
		}
		cValue.data = unsafe.Pointer(&cValueData[0])
		cValue.len = C.int(len(value))
	}
	C.C_NSPDFInfo_SetTagNames(n.Ptr(), cValue)
}

func (n NSPDFInfo) Orientation() PaperOrientation {
	result_ := C.C_NSPDFInfo_Orientation(n.Ptr())
	return PaperOrientation(int(result_))
}

func (n NSPDFInfo) SetOrientation(value PaperOrientation) {
	C.C_NSPDFInfo_SetOrientation(n.Ptr(), C.int(int(value)))
}

func (n NSPDFInfo) PaperSize() foundation.Size {
	result_ := C.C_NSPDFInfo_PaperSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func (n NSPDFInfo) SetPaperSize(value foundation.Size) {
	C.C_NSPDFInfo_SetPaperSize(n.Ptr(), *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(value))))
}

func (n NSPDFInfo) Attributes() map[PrintInfoAttributeKey]objc.Object {
	result_ := C.C_NSPDFInfo_Attributes(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[PrintInfoAttributeKey]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[PrintInfoAttributeKey(foundation.MakeString(k).String())] = objc.MakeObject(v)
	}
	return goResult_
}
