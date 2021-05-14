package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
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

func MakeTitlebarAccessoryViewController(ptr unsafe.Pointer) *NSTitlebarAccessoryViewController {
	if ptr == nil {
		return nil
	}
	return &NSTitlebarAccessoryViewController{
		NSViewController: *MakeViewController(ptr),
	}
}

func AllocTitlebarAccessoryViewController() *NSTitlebarAccessoryViewController {
	return MakeTitlebarAccessoryViewController(C.C_TitlebarAccessoryViewController_Alloc())
}

func (n *NSTitlebarAccessoryViewController) InitWithNibName_Bundle(nibNameOrNil NibName, nibBundleOrNil foundation.Bundle) TitlebarAccessoryViewController {
	result := C.C_NSTitlebarAccessoryViewController_InitWithNibName_Bundle(n.Ptr(), foundation.NewString(string(nibNameOrNil)).Ptr(), objc.ExtractPtr(nibBundleOrNil))
	return MakeTitlebarAccessoryViewController(result)
}

func (n *NSTitlebarAccessoryViewController) InitWithCoder(coder foundation.Coder) TitlebarAccessoryViewController {
	result := C.C_NSTitlebarAccessoryViewController_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTitlebarAccessoryViewController(result)
}

func (n *NSTitlebarAccessoryViewController) Init() TitlebarAccessoryViewController {
	result := C.C_NSTitlebarAccessoryViewController_Init(n.Ptr())
	return MakeTitlebarAccessoryViewController(result)
}

func (n *NSTitlebarAccessoryViewController) FullScreenMinHeight() coregraphics.Float {
	result := C.C_NSTitlebarAccessoryViewController_FullScreenMinHeight(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSTitlebarAccessoryViewController) SetFullScreenMinHeight(value coregraphics.Float) {
	C.C_NSTitlebarAccessoryViewController_SetFullScreenMinHeight(n.Ptr(), C.double(float64(value)))
}

func (n *NSTitlebarAccessoryViewController) LayoutAttribute() LayoutAttribute {
	result := C.C_NSTitlebarAccessoryViewController_LayoutAttribute(n.Ptr())
	return LayoutAttribute(int(result))
}

func (n *NSTitlebarAccessoryViewController) SetLayoutAttribute(value LayoutAttribute) {
	C.C_NSTitlebarAccessoryViewController_SetLayoutAttribute(n.Ptr(), C.int(int(value)))
}

func (n *NSTitlebarAccessoryViewController) AutomaticallyAdjustsSize() bool {
	result := C.C_NSTitlebarAccessoryViewController_AutomaticallyAdjustsSize(n.Ptr())
	return bool(result)
}

func (n *NSTitlebarAccessoryViewController) SetAutomaticallyAdjustsSize(value bool) {
	C.C_NSTitlebarAccessoryViewController_SetAutomaticallyAdjustsSize(n.Ptr(), C.bool(value))
}

func (n *NSTitlebarAccessoryViewController) IsHidden() bool {
	result := C.C_NSTitlebarAccessoryViewController_IsHidden(n.Ptr())
	return bool(result)
}

func (n *NSTitlebarAccessoryViewController) SetHidden(value bool) {
	C.C_NSTitlebarAccessoryViewController_SetHidden(n.Ptr(), C.bool(value))
}
