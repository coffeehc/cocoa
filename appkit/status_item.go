package appkit

// #include "status_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type StatusItem interface {
	objc.Object
	StatusBar() StatusBar
	Behavior() StatusItemBehavior
	SetBehavior(value StatusItemBehavior)
	Button() StatusBarButton
	Menu() Menu
	SetMenu(value Menu)
	IsVisible() bool
	SetVisible(value bool)
	Length() coregraphics.Float
	SetLength(value coregraphics.Float)
	AutosaveName() StatusItemAutosaveName
	SetAutosaveName(value StatusItemAutosaveName)
}

type NSStatusItem struct {
	objc.NSObject
}

func MakeStatusItem(ptr unsafe.Pointer) NSStatusItem {
	return NSStatusItem{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocStatusItem() NSStatusItem {
	return MakeStatusItem(C.C_StatusItem_Alloc())
}

func (n NSStatusItem) Init() StatusItem {
	result_ := C.C_NSStatusItem_Init(n.Ptr())
	return MakeStatusItem(result_)
}

func (n NSStatusItem) StatusBar() StatusBar {
	result_ := C.C_NSStatusItem_StatusBar(n.Ptr())
	return MakeStatusBar(result_)
}

func (n NSStatusItem) Behavior() StatusItemBehavior {
	result_ := C.C_NSStatusItem_Behavior(n.Ptr())
	return StatusItemBehavior(uint(result_))
}

func (n NSStatusItem) SetBehavior(value StatusItemBehavior) {
	C.C_NSStatusItem_SetBehavior(n.Ptr(), C.uint(uint(value)))
}

func (n NSStatusItem) Button() StatusBarButton {
	result_ := C.C_NSStatusItem_Button(n.Ptr())
	return MakeStatusBarButton(result_)
}

func (n NSStatusItem) Menu() Menu {
	result_ := C.C_NSStatusItem_Menu(n.Ptr())
	return MakeMenu(result_)
}

func (n NSStatusItem) SetMenu(value Menu) {
	C.C_NSStatusItem_SetMenu(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSStatusItem) IsVisible() bool {
	result_ := C.C_NSStatusItem_IsVisible(n.Ptr())
	return bool(result_)
}

func (n NSStatusItem) SetVisible(value bool) {
	C.C_NSStatusItem_SetVisible(n.Ptr(), C.bool(value))
}

func (n NSStatusItem) Length() coregraphics.Float {
	result_ := C.C_NSStatusItem_Length(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSStatusItem) SetLength(value coregraphics.Float) {
	C.C_NSStatusItem_SetLength(n.Ptr(), C.double(float64(value)))
}

func (n NSStatusItem) AutosaveName() StatusItemAutosaveName {
	result_ := C.C_NSStatusItem_AutosaveName(n.Ptr())
	return StatusItemAutosaveName(foundation.MakeString(result_).String())
}

func (n NSStatusItem) SetAutosaveName(value StatusItemAutosaveName) {
	C.C_NSStatusItem_SetAutosaveName(n.Ptr(), foundation.NewString(string(value)).Ptr())
}
