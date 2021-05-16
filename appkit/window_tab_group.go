package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "window_tab_group.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type WindowTabGroup interface {
	objc.Object
	AddWindow(window Window)
	InsertWindow_AtIndex(window Window, index int)
	RemoveWindow(window Window)
	Identifier() WindowTabbingIdentifier
	IsOverviewVisible() bool
	SetOverviewVisible(value bool)
	IsTabBarVisible() bool
	SelectedWindow() Window
	SetSelectedWindow(value Window)
}

type NSWindowTabGroup struct {
	objc.NSObject
}

func MakeWindowTabGroup(ptr unsafe.Pointer) *NSWindowTabGroup {
	if ptr == nil {
		return nil
	}
	return &NSWindowTabGroup{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocWindowTabGroup() *NSWindowTabGroup {
	return MakeWindowTabGroup(C.C_WindowTabGroup_Alloc())
}

func (n *NSWindowTabGroup) Init() WindowTabGroup {
	result := C.C_NSWindowTabGroup_Init(n.Ptr())
	return MakeWindowTabGroup(result)
}

func (n *NSWindowTabGroup) AddWindow(window Window) {
	C.C_NSWindowTabGroup_AddWindow(n.Ptr(), objc.ExtractPtr(window))
}

func (n *NSWindowTabGroup) InsertWindow_AtIndex(window Window, index int) {
	C.C_NSWindowTabGroup_InsertWindow_AtIndex(n.Ptr(), objc.ExtractPtr(window), C.int(index))
}

func (n *NSWindowTabGroup) RemoveWindow(window Window) {
	C.C_NSWindowTabGroup_RemoveWindow(n.Ptr(), objc.ExtractPtr(window))
}

func (n *NSWindowTabGroup) Identifier() WindowTabbingIdentifier {
	result := C.C_NSWindowTabGroup_Identifier(n.Ptr())
	return WindowTabbingIdentifier(foundation.MakeString(result).String())
}

func (n *NSWindowTabGroup) IsOverviewVisible() bool {
	result := C.C_NSWindowTabGroup_IsOverviewVisible(n.Ptr())
	return bool(result)
}

func (n *NSWindowTabGroup) SetOverviewVisible(value bool) {
	C.C_NSWindowTabGroup_SetOverviewVisible(n.Ptr(), C.bool(value))
}

func (n *NSWindowTabGroup) IsTabBarVisible() bool {
	result := C.C_NSWindowTabGroup_IsTabBarVisible(n.Ptr())
	return bool(result)
}

func (n *NSWindowTabGroup) SelectedWindow() Window {
	result := C.C_NSWindowTabGroup_SelectedWindow(n.Ptr())
	return MakeWindow(result)
}

func (n *NSWindowTabGroup) SetSelectedWindow(value Window) {
	C.C_NSWindowTabGroup_SetSelectedWindow(n.Ptr(), objc.ExtractPtr(value))
}