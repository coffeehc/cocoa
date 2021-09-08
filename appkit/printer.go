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
	DeviceDescription() map[DeviceDescriptionKey]objc.Object
}

type NSPrinter struct {
	objc.NSObject
}

func MakePrinter(ptr unsafe.Pointer) NSPrinter {
	return NSPrinter{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocPrinter() NSPrinter {
	result_ := C.C_NSPrinter_AllocPrinter()
	return MakePrinter(result_)
}

func (n NSPrinter) Init() NSPrinter {
	result_ := C.C_NSPrinter_Init(n.Ptr())
	return MakePrinter(result_)
}

func NewPrinter() NSPrinter {
	result_ := C.C_NSPrinter_NewPrinter()
	return MakePrinter(result_)
}

func (n NSPrinter) Autorelease() NSPrinter {
	result_ := C.C_NSPrinter_Autorelease(n.Ptr())
	return MakePrinter(result_)
}

func (n NSPrinter) Retain() NSPrinter {
	result_ := C.C_NSPrinter_Retain(n.Ptr())
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

func (n NSPrinter) PageSizeForPaper(paperName PrinterPaperName) foundation.Size {
	result_ := C.C_NSPrinter_PageSizeForPaper(n.Ptr(), foundation.NewString(string(paperName)).Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}

func PrinterNames() []string {
	result_ := C.C_NSPrinter_PrinterNames()
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

func PrinterTypes() []PrinterTypeName {
	result_ := C.C_NSPrinter_PrinterTypes()
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]PrinterTypeName, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = PrinterTypeName(foundation.MakeString(r).String())
	}
	return goResult_
}

func (n NSPrinter) Name() string {
	result_ := C.C_NSPrinter_Name(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSPrinter) Type() PrinterTypeName {
	result_ := C.C_NSPrinter_Type(n.Ptr())
	return PrinterTypeName(foundation.MakeString(result_).String())
}

func (n NSPrinter) LanguageLevel() int {
	result_ := C.C_NSPrinter_LanguageLevel(n.Ptr())
	return int(result_)
}

func (n NSPrinter) DeviceDescription() map[DeviceDescriptionKey]objc.Object {
	result_ := C.C_NSPrinter_DeviceDescription(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[DeviceDescriptionKey]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[DeviceDescriptionKey(foundation.MakeString(k).String())] = objc.MakeObject(v)
	}
	return goResult_
}
