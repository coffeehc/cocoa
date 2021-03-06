package appkit

// #include "touch_bar_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TouchBarItem interface {
	objc.Object
	Identifier() TouchBarItemIdentifier
	VisibilityPriority() TouchBarItemPriority
	SetVisibilityPriority(value TouchBarItemPriority)
	IsVisible() bool
	CustomizationLabel() string
	ViewController() ViewController
	View() View
}

type NSTouchBarItem struct {
	objc.NSObject
}

func MakeTouchBarItem(ptr unsafe.Pointer) NSTouchBarItem {
	return NSTouchBarItem{
		NSObject: objc.MakeObject(ptr),
	}
}

func (n NSTouchBarItem) InitWithIdentifier(identifier TouchBarItemIdentifier) NSTouchBarItem {
	result_ := C.C_NSTouchBarItem_InitWithIdentifier(n.Ptr(), foundation.NewString(string(identifier)).Ptr())
	return MakeTouchBarItem(result_)
}

func (n NSTouchBarItem) InitWithCoder(coder foundation.Coder) NSTouchBarItem {
	result_ := C.C_NSTouchBarItem_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeTouchBarItem(result_)
}

func AllocTouchBarItem() NSTouchBarItem {
	result_ := C.C_NSTouchBarItem_AllocTouchBarItem()
	return MakeTouchBarItem(result_)
}

func (n NSTouchBarItem) Autorelease() NSTouchBarItem {
	result_ := C.C_NSTouchBarItem_Autorelease(n.Ptr())
	return MakeTouchBarItem(result_)
}

func (n NSTouchBarItem) Retain() NSTouchBarItem {
	result_ := C.C_NSTouchBarItem_Retain(n.Ptr())
	return MakeTouchBarItem(result_)
}

func (n NSTouchBarItem) Identifier() TouchBarItemIdentifier {
	result_ := C.C_NSTouchBarItem_Identifier(n.Ptr())
	return TouchBarItemIdentifier(foundation.MakeString(result_).String())
}

func (n NSTouchBarItem) VisibilityPriority() TouchBarItemPriority {
	result_ := C.C_NSTouchBarItem_VisibilityPriority(n.Ptr())
	return TouchBarItemPriority(float32(result_))
}

func (n NSTouchBarItem) SetVisibilityPriority(value TouchBarItemPriority) {
	C.C_NSTouchBarItem_SetVisibilityPriority(n.Ptr(), C.float(float32(value)))
}

func (n NSTouchBarItem) IsVisible() bool {
	result_ := C.C_NSTouchBarItem_IsVisible(n.Ptr())
	return bool(result_)
}

func (n NSTouchBarItem) CustomizationLabel() string {
	result_ := C.C_NSTouchBarItem_CustomizationLabel(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSTouchBarItem) ViewController() ViewController {
	result_ := C.C_NSTouchBarItem_ViewController(n.Ptr())
	return MakeViewController(result_)
}

func (n NSTouchBarItem) View() View {
	result_ := C.C_NSTouchBarItem_View(n.Ptr())
	return MakeView(result_)
}
