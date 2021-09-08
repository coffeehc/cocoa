package appkit

// #include "titlebar_accessory_view_controller.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TitlebarAccessoryViewController interface {
	ViewController
	FullScreenMinHeight() coregraphics.Float
	SetFullScreenMinHeight(value coregraphics.Float)
	LayoutAttribute() LayoutAttribute
	SetLayoutAttribute(value LayoutAttribute)
	AutomaticallyAdjustsSize() bool
	SetAutomaticallyAdjustsSize(value bool)
	IsHidden() bool
	SetHidden(value bool)
}

type NSTitlebarAccessoryViewController struct {
	NSViewController
}

func MakeTitlebarAccessoryViewController(ptr unsafe.Pointer) NSTitlebarAccessoryViewController {
	return NSTitlebarAccessoryViewController{
		NSViewController: MakeViewController(ptr),
	}
}

func (n NSTitlebarAccessoryViewController) InitWithNibName_Bundle(nibNameOrNil NibName, nibBundleOrNil foundation.Bundle) NSTitlebarAccessoryViewController {
	result_ := C.C_NSTitlebarAccessoryViewController_InitWithNibName_Bundle(n.Ptr(), foundation.NewString(string(nibNameOrNil)).Ptr(), objc.ExtractPtr(nibBundleOrNil))
	return MakeTitlebarAccessoryViewController(result_)
}

func (n NSTitlebarAccessoryViewController) InitWithCoder(coder foundation.Coder) NSTitlebarAccessoryViewController {
	result_ := C.C_NSTitlebarAccessoryViewController_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTitlebarAccessoryViewController(result_)
}

func (n NSTitlebarAccessoryViewController) Init() NSTitlebarAccessoryViewController {
	result_ := C.C_NSTitlebarAccessoryViewController_Init(n.Ptr())
	return MakeTitlebarAccessoryViewController(result_)
}

func AllocTitlebarAccessoryViewController() NSTitlebarAccessoryViewController {
	result_ := C.C_NSTitlebarAccessoryViewController_AllocTitlebarAccessoryViewController()
	return MakeTitlebarAccessoryViewController(result_)
}

func NewTitlebarAccessoryViewController() NSTitlebarAccessoryViewController {
	result_ := C.C_NSTitlebarAccessoryViewController_NewTitlebarAccessoryViewController()
	return MakeTitlebarAccessoryViewController(result_)
}

func (n NSTitlebarAccessoryViewController) Autorelease() NSTitlebarAccessoryViewController {
	result_ := C.C_NSTitlebarAccessoryViewController_Autorelease(n.Ptr())
	return MakeTitlebarAccessoryViewController(result_)
}

func (n NSTitlebarAccessoryViewController) Retain() NSTitlebarAccessoryViewController {
	result_ := C.C_NSTitlebarAccessoryViewController_Retain(n.Ptr())
	return MakeTitlebarAccessoryViewController(result_)
}

func (n NSTitlebarAccessoryViewController) FullScreenMinHeight() coregraphics.Float {
	result_ := C.C_NSTitlebarAccessoryViewController_FullScreenMinHeight(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSTitlebarAccessoryViewController) SetFullScreenMinHeight(value coregraphics.Float) {
	C.C_NSTitlebarAccessoryViewController_SetFullScreenMinHeight(n.Ptr(), C.double(float64(value)))
}

func (n NSTitlebarAccessoryViewController) LayoutAttribute() LayoutAttribute {
	result_ := C.C_NSTitlebarAccessoryViewController_LayoutAttribute(n.Ptr())
	return LayoutAttribute(int(result_))
}

func (n NSTitlebarAccessoryViewController) SetLayoutAttribute(value LayoutAttribute) {
	C.C_NSTitlebarAccessoryViewController_SetLayoutAttribute(n.Ptr(), C.int(int(value)))
}

func (n NSTitlebarAccessoryViewController) AutomaticallyAdjustsSize() bool {
	result_ := C.C_NSTitlebarAccessoryViewController_AutomaticallyAdjustsSize(n.Ptr())
	return bool(result_)
}

func (n NSTitlebarAccessoryViewController) SetAutomaticallyAdjustsSize(value bool) {
	C.C_NSTitlebarAccessoryViewController_SetAutomaticallyAdjustsSize(n.Ptr(), C.bool(value))
}

func (n NSTitlebarAccessoryViewController) IsHidden() bool {
	result_ := C.C_NSTitlebarAccessoryViewController_IsHidden(n.Ptr())
	return bool(result_)
}

func (n NSTitlebarAccessoryViewController) SetHidden(value bool) {
	C.C_NSTitlebarAccessoryViewController_SetHidden(n.Ptr(), C.bool(value))
}
