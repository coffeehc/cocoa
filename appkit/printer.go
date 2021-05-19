package appkit

// #include "printer.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Printer interface {
	objc.Object
	PageSizeForPaper(paperName PrinterPaperName) foundation.Size
	Name() string
	Type() PrinterTypeName
	LanguageLevel() int
}

type NSPrinter struct {
	objc.NSObject
}

func MakePrinter(ptr unsafe.Pointer) *NSPrinter {
	if ptr == nil {
		return nil
	}
	return &NSPrinter{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocPrinter() *NSPrinter {
	return MakePrinter(C.C_Printer_Alloc())
}

func (n *NSPrinter) Init() Printer {
	result_ := C.C_NSPrinter_Init(n.Ptr())
	return MakePrinter(result_)
}

func PrinterWithName(name string) Printer {
	result_ := C.C_NSPrinter_PrinterWithName(foundation.NewString(name).Ptr())
	return MakePrinter(result_)
}

func PrinterWithType(_type PrinterTypeName) Printer {
	result_ := C.C_NSPrinter_PrinterWithType(foundation.NewString(string(_type)).Ptr())
	return MakePrinter(result_)
}

func (n *NSPrinter) PageSizeForPaper(paperName PrinterPaperName) foundation.Size {
	result_ := C.C_NSPrinter_PageSizeForPaper(n.Ptr(), foundation.NewString(string(paperName)).Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func PrinterNames() []string {
	result_ := C.C_NSPrinter_PrinterNames()
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]string, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = foundation.MakeString(r).String()
	}
	return goResult_
}

func PrinterTypes() []PrinterTypeName {
	result_ := C.C_NSPrinter_PrinterTypes()
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]PrinterTypeName, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = PrinterTypeName(foundation.MakeString(r).String())
	}
	return goResult_
}

func (n *NSPrinter) Name() string {
	result_ := C.C_NSPrinter_Name(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSPrinter) Type() PrinterTypeName {
	result_ := C.C_NSPrinter_Type(n.Ptr())
	return PrinterTypeName(foundation.MakeString(result_).String())
}

func (n *NSPrinter) LanguageLevel() int {
	result_ := C.C_NSPrinter_LanguageLevel(n.Ptr())
	return int(result_)
}
