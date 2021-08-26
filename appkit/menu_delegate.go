package appkit

// #include "menu_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type MenuDelegate struct {
	Menu_UpdateItem_AtIndex_ShouldCancel func(menu Menu, item MenuItem, index int, shouldCancel bool) bool
	ConfinementRectForMenu_OnScreen      func(menu Menu, screen Screen) foundation.Rect
	Menu_WillHighlightItem               func(menu Menu, item MenuItem)
	MenuWillOpen                         func(menu Menu)
	MenuDidClose                         func(menu Menu)
	NumberOfItemsInMenu                  func(menu Menu) int
	MenuNeedsUpdate                      func(menu Menu)
}

func (delegate *MenuDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapMenuDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export menuDelegate_Menu_UpdateItem_AtIndex_ShouldCancel
func menuDelegate_Menu_UpdateItem_AtIndex_ShouldCancel(hp C.uintptr_t, menu unsafe.Pointer, item unsafe.Pointer, index C.int, shouldCancel C.bool) C.bool {
	delegate := cgo.Handle(hp).Value().(*MenuDelegate)
	result := delegate.Menu_UpdateItem_AtIndex_ShouldCancel(MakeMenu(menu), MakeMenuItem(item), int(index), bool(shouldCancel))
	return C.bool(result)
}

//export menuDelegate_ConfinementRectForMenu_OnScreen
func menuDelegate_ConfinementRectForMenu_OnScreen(hp C.uintptr_t, menu unsafe.Pointer, screen unsafe.Pointer) C.CGRect {
	delegate := cgo.Handle(hp).Value().(*MenuDelegate)
	result := delegate.ConfinementRectForMenu_OnScreen(MakeMenu(menu), MakeScreen(screen))
	return *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(result)))
}

//export menuDelegate_Menu_WillHighlightItem
func menuDelegate_Menu_WillHighlightItem(hp C.uintptr_t, menu unsafe.Pointer, item unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*MenuDelegate)
	delegate.Menu_WillHighlightItem(MakeMenu(menu), MakeMenuItem(item))
}

//export menuDelegate_MenuWillOpen
func menuDelegate_MenuWillOpen(hp C.uintptr_t, menu unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*MenuDelegate)
	delegate.MenuWillOpen(MakeMenu(menu))
}

//export menuDelegate_MenuDidClose
func menuDelegate_MenuDidClose(hp C.uintptr_t, menu unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*MenuDelegate)
	delegate.MenuDidClose(MakeMenu(menu))
}

//export menuDelegate_NumberOfItemsInMenu
func menuDelegate_NumberOfItemsInMenu(hp C.uintptr_t, menu unsafe.Pointer) C.int {
	delegate := cgo.Handle(hp).Value().(*MenuDelegate)
	result := delegate.NumberOfItemsInMenu(MakeMenu(menu))
	return C.int(result)
}

//export menuDelegate_MenuNeedsUpdate
func menuDelegate_MenuNeedsUpdate(hp C.uintptr_t, menu unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*MenuDelegate)
	delegate.MenuNeedsUpdate(MakeMenu(menu))
}

//export MenuDelegate_RespondsTo
func MenuDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*MenuDelegate)
	switch selName {
	case "menu:updateItem:atIndex:shouldCancel:":
		return delegate.Menu_UpdateItem_AtIndex_ShouldCancel != nil
	case "confinementRectForMenu:onScreen:":
		return delegate.ConfinementRectForMenu_OnScreen != nil
	case "menu:willHighlightItem:":
		return delegate.Menu_WillHighlightItem != nil
	case "menuWillOpen:":
		return delegate.MenuWillOpen != nil
	case "menuDidClose:":
		return delegate.MenuDidClose != nil
	case "numberOfItemsInMenu:":
		return delegate.NumberOfItemsInMenu != nil
	case "menuNeedsUpdate:":
		return delegate.MenuNeedsUpdate != nil
	default:
		return false
	}
}
