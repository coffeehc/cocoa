package appkit

// #include "menu_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
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

func WrapMenuDelegate(delegate *MenuDelegate) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapMenuDelegate(C.long(id))
	return objc.MakeObject(ptr)
}

//export menuDelegate_Menu_UpdateItem_AtIndex_ShouldCancel
func menuDelegate_Menu_UpdateItem_AtIndex_ShouldCancel(id int64, menu unsafe.Pointer, item unsafe.Pointer, index C.int, shouldCancel C.bool) C.bool {
	delegate := resources.Get(id).(*MenuDelegate)
	result := delegate.Menu_UpdateItem_AtIndex_ShouldCancel(MakeMenu(menu), MakeMenuItem(item), int(index), bool(shouldCancel))
	return C.bool(result)
}

//export menuDelegate_ConfinementRectForMenu_OnScreen
func menuDelegate_ConfinementRectForMenu_OnScreen(id int64, menu unsafe.Pointer, screen unsafe.Pointer) C.CGRect {
	delegate := resources.Get(id).(*MenuDelegate)
	result := delegate.ConfinementRectForMenu_OnScreen(MakeMenu(menu), MakeScreen(screen))
	return *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(result)))
}

//export menuDelegate_Menu_WillHighlightItem
func menuDelegate_Menu_WillHighlightItem(id int64, menu unsafe.Pointer, item unsafe.Pointer) {
	delegate := resources.Get(id).(*MenuDelegate)
	delegate.Menu_WillHighlightItem(MakeMenu(menu), MakeMenuItem(item))
}

//export menuDelegate_MenuWillOpen
func menuDelegate_MenuWillOpen(id int64, menu unsafe.Pointer) {
	delegate := resources.Get(id).(*MenuDelegate)
	delegate.MenuWillOpen(MakeMenu(menu))
}

//export menuDelegate_MenuDidClose
func menuDelegate_MenuDidClose(id int64, menu unsafe.Pointer) {
	delegate := resources.Get(id).(*MenuDelegate)
	delegate.MenuDidClose(MakeMenu(menu))
}

//export menuDelegate_NumberOfItemsInMenu
func menuDelegate_NumberOfItemsInMenu(id int64, menu unsafe.Pointer) C.int {
	delegate := resources.Get(id).(*MenuDelegate)
	result := delegate.NumberOfItemsInMenu(MakeMenu(menu))
	return C.int(result)
}

//export menuDelegate_MenuNeedsUpdate
func menuDelegate_MenuNeedsUpdate(id int64, menu unsafe.Pointer) {
	delegate := resources.Get(id).(*MenuDelegate)
	delegate.MenuNeedsUpdate(MakeMenu(menu))
}

//export MenuDelegate_RespondsTo
func MenuDelegate_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*MenuDelegate)
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

//export deleteMenuDelegate
func deleteMenuDelegate(id int64) {
	resources.Delete(id)
}
