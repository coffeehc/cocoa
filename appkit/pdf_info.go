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
	return MakePDFInfo(C.C_PDFInfo_Alloc())
}

func (n NSPDFInfo) Init() PDFInfo {
	result_ := C.C_NSPDFInfo_Init(n.Ptr())
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
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeString(r).String()
	}
	return goResult_
}

func (n NSPDFInfo) SetTagNames(value []string) {
	cValueData := make([]unsafe.Pointer, len(value))
	for idx, v := range value {
		cValueData[idx] = foundation.NewString(v).Ptr()
	}
	cValue := C.Array{data: unsafe.Pointer(&cValueData[0]), len: C.int(len(value))}
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
