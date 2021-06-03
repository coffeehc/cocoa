package appkit

// #include "pdf_panel.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type PDFPanel interface {
	objc.Object
	AccessoryController() ViewController
	SetAccessoryController(value ViewController)
	Options() PDFPanelOptions
	SetOptions(value PDFPanelOptions)
	DefaultFileName() string
	SetDefaultFileName(value string)
}

type NSPDFPanel struct {
	objc.NSObject
}

func MakePDFPanel(ptr unsafe.Pointer) NSPDFPanel {
	return NSPDFPanel{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocPDFPanel() NSPDFPanel {
	return MakePDFPanel(C.C_PDFPanel_Alloc())
}

func (n NSPDFPanel) Init() PDFPanel {
	result_ := C.C_NSPDFPanel_Init(n.Ptr())
	return MakePDFPanel(result_)
}

func PDFPanel_Panel() PDFPanel {
	result_ := C.C_NSPDFPanel_PDFPanel_Panel()
	return MakePDFPanel(result_)
}

func (n NSPDFPanel) AccessoryController() ViewController {
	result_ := C.C_NSPDFPanel_AccessoryController(n.Ptr())
	return MakeViewController(result_)
}

func (n NSPDFPanel) SetAccessoryController(value ViewController) {
	C.C_NSPDFPanel_SetAccessoryController(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSPDFPanel) Options() PDFPanelOptions {
	result_ := C.C_NSPDFPanel_Options(n.Ptr())
	return PDFPanelOptions(int(result_))
}

func (n NSPDFPanel) SetOptions(value PDFPanelOptions) {
	C.C_NSPDFPanel_SetOptions(n.Ptr(), C.int(int(value)))
}

func (n NSPDFPanel) DefaultFileName() string {
	result_ := C.C_NSPDFPanel_DefaultFileName(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSPDFPanel) SetDefaultFileName(value string) {
	C.C_NSPDFPanel_SetDefaultFileName(n.Ptr(), foundation.NewString(value).Ptr())
}
