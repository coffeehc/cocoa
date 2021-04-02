package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "status_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type StatusItem interface {
	objc.Object
	StatusBar() StatusBar
	Behavior() StatusItemBehavior
	SetBehavior(behavior StatusItemBehavior)
	Button() StatusBarButton
	Menu() Menu
	SetMenu(menu Menu)
	IsVisible() bool
	SetVisible(visible bool)
	Length() float64
	SetLength(length float64)
	AutosaveName() string
	SetAutosaveName(autosaveName string)
}

var _ StatusItem = (*NSStatusItem)(nil)

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

func (s *NSStatusItem) StatusBar() StatusBar {
	return MakeStatusBar(C.StatusItem_StatusBar(s.Ptr()))
}

func (s *NSStatusItem) Behavior() StatusItemBehavior {
	return StatusItemBehavior(C.StatusItem_Behavior(s.Ptr()))
}

func (s *NSStatusItem) SetBehavior(behavior StatusItemBehavior) {
	C.StatusItem_SetBehavior(s.Ptr(), C.ulong(behavior))
}

func (s *NSStatusItem) Button() StatusBarButton {
	return MakeStatusBarButton(C.StatusItem_Button(s.Ptr()))
}

func (s *NSStatusItem) Menu() Menu {
	return MakeMenu(C.StatusItem_Menu(s.Ptr()))
}

func (s *NSStatusItem) SetMenu(menu Menu) {
	C.StatusItem_SetMenu(s.Ptr(), toPointer(menu))
}

func (s *NSStatusItem) IsVisible() bool {
	return bool(C.StatusItem_IsVisible(s.Ptr()))
}

func (s *NSStatusItem) SetVisible(visible bool) {
	C.StatusItem_SetVisible(s.Ptr(), C.bool(visible))
}

func (s *NSStatusItem) Length() float64 {
	return float64(C.StatusItem_Length(s.Ptr()))
}

func (s *NSStatusItem) SetLength(length float64) {
	C.StatusItem_SetLength(s.Ptr(), C.double(length))
}

func (s *NSStatusItem) AutosaveName() string {
	return C.GoString(C.StatusItem_AutosaveName(s.Ptr()))
}

func (s *NSStatusItem) SetAutosaveName(autosaveName string) {
	cAutosaveName := C.CString(autosaveName)
	defer C.free(unsafe.Pointer(cAutosaveName))
	C.StatusItem_SetAutosaveName(s.Ptr(), cAutosaveName)
}
