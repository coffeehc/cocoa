package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "menu_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type MenuItem interface {
	objc.Object
	IsEnabled() bool
	SetEnabled(enabled bool)
	IsHidden() bool
	SetHidden(hidden bool)
	Title() string
	SetTitle(title string)
	Submenu() Menu
	SetSubmenu(submenu Menu)
	HasSubmenu() bool
	IsSeparatorItem() bool
	Menu() Menu
	SetMenu(menu Menu)
	ToolTip() string
	SetToolTip(toolTip string)
	IsHighlighted() bool
	KeyEquivalent() string
	SetKeyEquivalent(keyEquivalent string)
	KeyEquivalentModifierMask() EventModifierFlags
	SetKeyEquivalentModifierMask(keyEquivalentModifierMask EventModifierFlags)
	UserKeyEquivalent() string
	IsAlternate() bool
	SetAlternate(alternate bool)
	IndentationLevel() int
	SetIndentationLevel(indentationLevel int)
	View() View
	SetView(view View)
	AllowsKeyEquivalentWhenHidden() bool
	SetAllowsKeyEquivalentWhenHidden(allowsKeyEquivalentWhenHidden bool)
	State() ControlStateValue
	SetState(state ControlStateValue)
	Tag() int
	SetTag(tag int)
	SetAction(handler ActionHandler)
}

var _ MenuItem = (*NSMenuItem)(nil)

type NSMenuItem struct {
	objc.NSObject
	action ActionHandler
}

func MakeMenuItem(ptr unsafe.Pointer) *NSMenuItem {
	if ptr == nil {
		return nil
	}
	return &NSMenuItem{
		NSObject: *objc.MakeObject(ptr),
	}
}
func (m *NSMenuItem) setDelegate() {
	id := resources.Register(m)
	C.MenuItem_SetDelegate(m.Ptr(), C.long(id))
}

func (m *NSMenuItem) IsEnabled() bool {
	return bool(C.MenuItem_IsEnabled(m.Ptr()))
}

func (m *NSMenuItem) SetEnabled(enabled bool) {
	C.MenuItem_SetEnabled(m.Ptr(), C.bool(enabled))
}

func (m *NSMenuItem) IsHidden() bool {
	return bool(C.MenuItem_IsHidden(m.Ptr()))
}

func (m *NSMenuItem) SetHidden(hidden bool) {
	C.MenuItem_SetHidden(m.Ptr(), C.bool(hidden))
}

func (m *NSMenuItem) Title() string {
	return C.GoString(C.MenuItem_Title(m.Ptr()))
}

func (m *NSMenuItem) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.MenuItem_SetTitle(m.Ptr(), cTitle)
}

func (m *NSMenuItem) Submenu() Menu {
	return MakeMenu(C.MenuItem_Submenu(m.Ptr()))
}

func (m *NSMenuItem) SetSubmenu(submenu Menu) {
	C.MenuItem_SetSubmenu(m.Ptr(), toPointer(submenu))
}

func (m *NSMenuItem) HasSubmenu() bool {
	return bool(C.MenuItem_HasSubmenu(m.Ptr()))
}

func (m *NSMenuItem) IsSeparatorItem() bool {
	return bool(C.MenuItem_IsSeparatorItem(m.Ptr()))
}

func (m *NSMenuItem) Menu() Menu {
	return MakeMenu(C.MenuItem_Menu(m.Ptr()))
}

func (m *NSMenuItem) SetMenu(menu Menu) {
	C.MenuItem_SetMenu(m.Ptr(), toPointer(menu))
}

func (m *NSMenuItem) ToolTip() string {
	return C.GoString(C.MenuItem_ToolTip(m.Ptr()))
}

func (m *NSMenuItem) SetToolTip(toolTip string) {
	cToolTip := C.CString(toolTip)
	defer C.free(unsafe.Pointer(cToolTip))
	C.MenuItem_SetToolTip(m.Ptr(), cToolTip)
}

func (m *NSMenuItem) IsHighlighted() bool {
	return bool(C.MenuItem_IsHighlighted(m.Ptr()))
}

func (m *NSMenuItem) KeyEquivalent() string {
	return C.GoString(C.MenuItem_KeyEquivalent(m.Ptr()))
}

func (m *NSMenuItem) SetKeyEquivalent(keyEquivalent string) {
	cKeyEquivalent := C.CString(keyEquivalent)
	defer C.free(unsafe.Pointer(cKeyEquivalent))
	C.MenuItem_SetKeyEquivalent(m.Ptr(), cKeyEquivalent)
}

func (m *NSMenuItem) KeyEquivalentModifierMask() EventModifierFlags {
	return EventModifierFlags(C.MenuItem_KeyEquivalentModifierMask(m.Ptr()))
}

func (m *NSMenuItem) SetKeyEquivalentModifierMask(keyEquivalentModifierMask EventModifierFlags) {
	C.MenuItem_SetKeyEquivalentModifierMask(m.Ptr(), C.ulong(keyEquivalentModifierMask))
}

func (m *NSMenuItem) UserKeyEquivalent() string {
	return C.GoString(C.MenuItem_UserKeyEquivalent(m.Ptr()))
}

func (m *NSMenuItem) IsAlternate() bool {
	return bool(C.MenuItem_IsAlternate(m.Ptr()))
}

func (m *NSMenuItem) SetAlternate(alternate bool) {
	C.MenuItem_SetAlternate(m.Ptr(), C.bool(alternate))
}

func (m *NSMenuItem) IndentationLevel() int {
	return int(C.MenuItem_IndentationLevel(m.Ptr()))
}

func (m *NSMenuItem) SetIndentationLevel(indentationLevel int) {
	C.MenuItem_SetIndentationLevel(m.Ptr(), C.long(indentationLevel))
}

func (m *NSMenuItem) View() View {
	return MakeView(C.MenuItem_View(m.Ptr()))
}

func (m *NSMenuItem) SetView(view View) {
	C.MenuItem_SetView(m.Ptr(), toPointer(view))
}

func (m *NSMenuItem) AllowsKeyEquivalentWhenHidden() bool {
	return bool(C.MenuItem_AllowsKeyEquivalentWhenHidden(m.Ptr()))
}

func (m *NSMenuItem) SetAllowsKeyEquivalentWhenHidden(allowsKeyEquivalentWhenHidden bool) {
	C.MenuItem_SetAllowsKeyEquivalentWhenHidden(m.Ptr(), C.bool(allowsKeyEquivalentWhenHidden))
}

func UsesUserKeyEquivalents() bool {
	return bool(C.MenuItem_UsesUserKeyEquivalents())
}

func SetUsesUserKeyEquivalents(usesUserKeyEquivalents bool) {
	C.MenuItem_SetUsesUserKeyEquivalents(C.bool(usesUserKeyEquivalents))
}

func (m *NSMenuItem) State() ControlStateValue {
	return ControlStateValue(C.MenuItem_State(m.Ptr()))
}

func (m *NSMenuItem) SetState(state ControlStateValue) {
	C.MenuItem_SetState(m.Ptr(), C.long(state))
}

func (m *NSMenuItem) Tag() int {
	return int(C.MenuItem_Tag(m.Ptr()))
}

func (m *NSMenuItem) SetTag(tag int) {
	C.MenuItem_SetTag(m.Ptr(), C.long(tag))
}

func NewMenuItem(title string, selector *objc.Selector, charCode string) MenuItem {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	cCharCode := C.CString(charCode)
	defer C.free(unsafe.Pointer(cCharCode))
	return MakeMenuItem(C.MenuItem_NewMenuItem(cTitle, toPointer(selector), cCharCode))
}

func NewSeparatorItem() MenuItem {
	return MakeMenuItem(C.MenuItem_NewSeparatorItem())
}

func (m *NSMenuItem) SetAction(handler ActionHandler) {
	m.action = handler
}

//export MenuItem_Target_Action
func MenuItem_Target_Action(id int64, sender unsafe.Pointer) {
	menuItem := resources.Get(id).(*NSMenuItem)
	if menuItem.action != nil {
		menuItem.action(objc.MakeObject(sender))
	}
}
