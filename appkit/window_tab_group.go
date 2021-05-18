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
	Windows() []Window
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
	result_ := C.C_NSWindowTabGroup_Init(n.Ptr())
	return MakeWindowTabGroup(result_)
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
	result_ := C.C_NSWindowTabGroup_Identifier(n.Ptr())
	return WindowTabbingIdentifier(foundation.MakeString(result_).String())
}

func (n *NSWindowTabGroup) IsOverviewVisible() bool {
	result_ := C.C_NSWindowTabGroup_IsOverviewVisible(n.Ptr())
	return bool(result_)
}

func (n *NSWindowTabGroup) SetOverviewVisible(value bool) {
	C.C_NSWindowTabGroup_SetOverviewVisible(n.Ptr(), C.bool(value))
}

func (n *NSWindowTabGroup) IsTabBarVisible() bool {
	result_ := C.C_NSWindowTabGroup_IsTabBarVisible(n.Ptr())
	return bool(result_)
}

func (n *NSWindowTabGroup) Windows() []Window {
	result_ := C.C_NSWindowTabGroup_Windows(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]Window, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeWindow(r)
	}
	return goResult_
}

func (n *NSWindowTabGroup) SelectedWindow() Window {
	result_ := C.C_NSWindowTabGroup_SelectedWindow(n.Ptr())
	return MakeWindow(result_)
}

func (n *NSWindowTabGroup) SetSelectedWindow(value Window) {
	C.C_NSWindowTabGroup_SetSelectedWindow(n.Ptr(), objc.ExtractPtr(value))
}
