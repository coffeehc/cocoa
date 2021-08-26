package appkit

// #include "print_panel.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type PrintPanel interface {
	objc.Object
	DefaultButtonTitle() string
	SetDefaultButtonTitle(defaultButtonTitle string)
	AddAccessoryController(accessoryController ViewController)
	RemoveAccessoryController(accessoryController ViewController)
	RunModal() int
	RunModalWithPrintInfo(printInfo PrintInfo) int
	JobStyleHint() PrintPanelJobStyleHint
	SetJobStyleHint(value PrintPanelJobStyleHint)
	Options() PrintPanelOptions
	SetOptions(value PrintPanelOptions)
	HelpAnchor() HelpAnchorName
	SetHelpAnchor(value HelpAnchorName)
	AccessoryControllers() []ViewController
	PrintInfo() PrintInfo
}

type NSPrintPanel struct {
	objc.NSObject
}

func MakePrintPanel(ptr unsafe.Pointer) NSPrintPanel {
	return NSPrintPanel{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocPrintPanel() NSPrintPanel {
	return MakePrintPanel(C.C_PrintPanel_Alloc())
}

func (n NSPrintPanel) Init() PrintPanel {
	result_ := C.C_NSPrintPanel_Init(n.Ptr())
	return MakePrintPanel(result_)
}

func PrintPanel_() PrintPanel {
	result_ := C.C_NSPrintPanel_PrintPanel_()
	return MakePrintPanel(result_)
}

func (n NSPrintPanel) DefaultButtonTitle() string {
	result_ := C.C_NSPrintPanel_DefaultButtonTitle(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSPrintPanel) SetDefaultButtonTitle(defaultButtonTitle string) {
	C.C_NSPrintPanel_SetDefaultButtonTitle(n.Ptr(), foundation.NewString(defaultButtonTitle).Ptr())
}

func (n NSPrintPanel) AddAccessoryController(accessoryController ViewController) {
	C.C_NSPrintPanel_AddAccessoryController(n.Ptr(), objc.ExtractPtr(accessoryController))
}

func (n NSPrintPanel) RemoveAccessoryController(accessoryController ViewController) {
	C.C_NSPrintPanel_RemoveAccessoryController(n.Ptr(), objc.ExtractPtr(accessoryController))
}

func (n NSPrintPanel) RunModal() int {
	result_ := C.C_NSPrintPanel_RunModal(n.Ptr())
	return int(result_)
}

func (n NSPrintPanel) RunModalWithPrintInfo(printInfo PrintInfo) int {
	result_ := C.C_NSPrintPanel_RunModalWithPrintInfo(n.Ptr(), objc.ExtractPtr(printInfo))
	return int(result_)
}

func (n NSPrintPanel) JobStyleHint() PrintPanelJobStyleHint {
	result_ := C.C_NSPrintPanel_JobStyleHint(n.Ptr())
	return PrintPanelJobStyleHint(foundation.MakeString(result_).String())
}

func (n NSPrintPanel) SetJobStyleHint(value PrintPanelJobStyleHint) {
	C.C_NSPrintPanel_SetJobStyleHint(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n NSPrintPanel) Options() PrintPanelOptions {
	result_ := C.C_NSPrintPanel_Options(n.Ptr())
	return PrintPanelOptions(uint(result_))
}

func (n NSPrintPanel) SetOptions(value PrintPanelOptions) {
	C.C_NSPrintPanel_SetOptions(n.Ptr(), C.uint(uint(value)))
}

func (n NSPrintPanel) HelpAnchor() HelpAnchorName {
	result_ := C.C_NSPrintPanel_HelpAnchor(n.Ptr())
	return HelpAnchorName(foundation.MakeString(result_).String())
}

func (n NSPrintPanel) SetHelpAnchor(value HelpAnchorName) {
	C.C_NSPrintPanel_SetHelpAnchor(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n NSPrintPanel) AccessoryControllers() []ViewController {
	result_ := C.C_NSPrintPanel_AccessoryControllers(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]ViewController, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeViewController(r)
	}
	return goResult_
}

func (n NSPrintPanel) PrintInfo() PrintInfo {
	result_ := C.C_NSPrintPanel_PrintInfo(n.Ptr())
	return MakePrintInfo(result_)
}
