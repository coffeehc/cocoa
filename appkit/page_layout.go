package appkit

// #include "page_layout.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type PageLayout interface {
	objc.Object
	RunModal() int
	RunModalWithPrintInfo(printInfo PrintInfo) int
	AddAccessoryController(accessoryController ViewController)
	RemoveAccessoryController(accessoryController ViewController)
	AccessoryControllers() []ViewController
	PrintInfo() PrintInfo
}

type NSPageLayout struct {
	objc.NSObject
}

func MakePageLayout(ptr unsafe.Pointer) *NSPageLayout {
	if ptr == nil {
		return nil
	}
	return &NSPageLayout{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocPageLayout() *NSPageLayout {
	return MakePageLayout(C.C_PageLayout_Alloc())
}

func (n *NSPageLayout) Init() PageLayout {
	result_ := C.C_NSPageLayout_Init(n.Ptr())
	return MakePageLayout(result_)
}

func PageLayout_() PageLayout {
	result_ := C.C_NSPageLayout_PageLayout_()
	return MakePageLayout(result_)
}

func (n *NSPageLayout) RunModal() int {
	result_ := C.C_NSPageLayout_RunModal(n.Ptr())
	return int(result_)
}

func (n *NSPageLayout) RunModalWithPrintInfo(printInfo PrintInfo) int {
	result_ := C.C_NSPageLayout_RunModalWithPrintInfo(n.Ptr(), objc.ExtractPtr(printInfo))
	return int(result_)
}

func (n *NSPageLayout) AddAccessoryController(accessoryController ViewController) {
	C.C_NSPageLayout_AddAccessoryController(n.Ptr(), objc.ExtractPtr(accessoryController))
}

func (n *NSPageLayout) RemoveAccessoryController(accessoryController ViewController) {
	C.C_NSPageLayout_RemoveAccessoryController(n.Ptr(), objc.ExtractPtr(accessoryController))
}

func (n *NSPageLayout) AccessoryControllers() []ViewController {
	result_ := C.C_NSPageLayout_AccessoryControllers(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]ViewController, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeViewController(r)
	}
	return goResult_
}

func (n *NSPageLayout) PrintInfo() PrintInfo {
	result_ := C.C_NSPageLayout_PrintInfo(n.Ptr())
	return MakePrintInfo(result_)
}
