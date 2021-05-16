package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
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

func MakeStatusItem(ptr unsafe.Pointer) *NSStatusItem {
	if ptr == nil {
		return nil
	}
	return &NSStatusItem{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocStatusItem() *NSStatusItem {
	return MakeStatusItem(C.C_StatusItem_Alloc())
}

func (n *NSStatusItem) Init() StatusItem {
	result := C.C_NSStatusItem_Init(n.Ptr())
	return MakeStatusItem(result)
}

func (n *NSStatusItem) StatusBar() StatusBar {
	result := C.C_NSStatusItem_StatusBar(n.Ptr())
	return MakeStatusBar(result)
}

func (n *NSStatusItem) Behavior() StatusItemBehavior {
	result := C.C_NSStatusItem_Behavior(n.Ptr())
	return StatusItemBehavior(uint(result))
}

func (n *NSStatusItem) SetBehavior(value StatusItemBehavior) {
	C.C_NSStatusItem_SetBehavior(n.Ptr(), C.uint(uint(value)))
}

func (n *NSStatusItem) Button() StatusBarButton {
	result := C.C_NSStatusItem_Button(n.Ptr())
	return MakeStatusBarButton(result)
}

func (n *NSStatusItem) Menu() Menu {
	result := C.C_NSStatusItem_Menu(n.Ptr())
	return MakeMenu(result)
}

func (n *NSStatusItem) SetMenu(value Menu) {
	C.C_NSStatusItem_SetMenu(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSStatusItem) IsVisible() bool {
	result := C.C_NSStatusItem_IsVisible(n.Ptr())
	return bool(result)
}

func (n *NSStatusItem) SetVisible(value bool) {
	C.C_NSStatusItem_SetVisible(n.Ptr(), C.bool(value))
}

func (n *NSStatusItem) Length() coregraphics.Float {
	result := C.C_NSStatusItem_Length(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSStatusItem) SetLength(value coregraphics.Float) {
	C.C_NSStatusItem_SetLength(n.Ptr(), C.double(float64(value)))
}

func (n *NSStatusItem) AutosaveName() StatusItemAutosaveName {
	result := C.C_NSStatusItem_AutosaveName(n.Ptr())
	return StatusItemAutosaveName(foundation.MakeString(result).String())
}

func (n *NSStatusItem) SetAutosaveName(value StatusItemAutosaveName) {
	C.C_NSStatusItem_SetAutosaveName(n.Ptr(), foundation.NewString(string(value)).Ptr())
}
