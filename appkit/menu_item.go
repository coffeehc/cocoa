package appkit

// #include "menu_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type MenuItem interface {
	objc.Object
	IsEnabled() bool
	SetEnabled(value bool)
	IsHidden() bool
	SetHidden(value bool)
	IsHiddenOrHasHiddenAncestor() bool
	Target() objc.Object
	SetTarget(value objc.Object)
	Action() *objc.Selector
	SetAction(value *objc.Selector)
	Title() string
	SetTitle(value string)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.AttributedString)
	Tag() int
	SetTag(value int)
	State() ControlStateValue
	SetState(value ControlStateValue)
	Image() Image
	SetImage(value Image)
	OnStateImage() Image
	SetOnStateImage(value Image)
	OffStateImage() Image
	SetOffStateImage(value Image)
	MixedStateImage() Image
	SetMixedStateImage(value Image)
	Submenu() Menu
	SetSubmenu(value Menu)
	HasSubmenu() bool
	ParentItem() MenuItem
	IsSeparatorItem() bool
	Menu() Menu
	SetMenu(value Menu)
	KeyEquivalent() string
	SetKeyEquivalent(value string)
	KeyEquivalentModifierMask() EventModifierFlags
	SetKeyEquivalentModifierMask(value EventModifierFlags)
	UserKeyEquivalent() string
	IsAlternate() bool
	SetAlternate(value bool)
	IndentationLevel() int
	SetIndentationLevel(value int)
	ToolTip() string
	SetToolTip(value string)
	RepresentedObject() objc.Object
	SetRepresentedObject(value objc.Object)
	View() View
	SetView(value View)
	IsHighlighted() bool
	AllowsKeyEquivalentWhenHidden() bool
	SetAllowsKeyEquivalentWhenHidden(value bool)
}

type NSMenuItem struct {
	objc.NSObject
}

func MakeMenuItem(ptr unsafe.Pointer) *NSMenuItem {
	if ptr == nil {
		return nil
	}
	return &NSMenuItem{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocMenuItem() *NSMenuItem {
	return MakeMenuItem(C.C_MenuItem_Alloc())
}

func (n *NSMenuItem) InitWithTitle_Action_KeyEquivalent(_string string, selector *objc.Selector, charCode string) MenuItem {
	result_ := C.C_NSMenuItem_InitWithTitle_Action_KeyEquivalent(n.Ptr(), foundation.NewString(_string).Ptr(), objc.ExtractPtr(selector), foundation.NewString(charCode).Ptr())
	return MakeMenuItem(result_)
}

func (n *NSMenuItem) InitWithCoder(coder foundation.Coder) MenuItem {
	result_ := C.C_NSMenuItem_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeMenuItem(result_)
}

func (n *NSMenuItem) Init() MenuItem {
	result_ := C.C_NSMenuItem_Init(n.Ptr())
	return MakeMenuItem(result_)
}

func MenuItem_SeparatorItem() MenuItem {
	result_ := C.C_NSMenuItem_MenuItem_SeparatorItem()
	return MakeMenuItem(result_)
}

func (n *NSMenuItem) IsEnabled() bool {
	result_ := C.C_NSMenuItem_IsEnabled(n.Ptr())
	return bool(result_)
}

func (n *NSMenuItem) SetEnabled(value bool) {
	C.C_NSMenuItem_SetEnabled(n.Ptr(), C.bool(value))
}

func (n *NSMenuItem) IsHidden() bool {
	result_ := C.C_NSMenuItem_IsHidden(n.Ptr())
	return bool(result_)
}

func (n *NSMenuItem) SetHidden(value bool) {
	C.C_NSMenuItem_SetHidden(n.Ptr(), C.bool(value))
}

func (n *NSMenuItem) IsHiddenOrHasHiddenAncestor() bool {
	result_ := C.C_NSMenuItem_IsHiddenOrHasHiddenAncestor(n.Ptr())
	return bool(result_)
}

func (n *NSMenuItem) Target() objc.Object {
	result_ := C.C_NSMenuItem_Target(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSMenuItem) SetTarget(value objc.Object) {
	C.C_NSMenuItem_SetTarget(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSMenuItem) Action() *objc.Selector {
	result_ := C.C_NSMenuItem_Action(n.Ptr())
	return objc.MakeSelector(result_)
}

func (n *NSMenuItem) SetAction(value *objc.Selector) {
	C.C_NSMenuItem_SetAction(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSMenuItem) Title() string {
	result_ := C.C_NSMenuItem_Title(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSMenuItem) SetTitle(value string) {
	C.C_NSMenuItem_SetTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSMenuItem) AttributedTitle() foundation.AttributedString {
	result_ := C.C_NSMenuItem_AttributedTitle(n.Ptr())
	return foundation.MakeAttributedString(result_)
}

func (n *NSMenuItem) SetAttributedTitle(value foundation.AttributedString) {
	C.C_NSMenuItem_SetAttributedTitle(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSMenuItem) Tag() int {
	result_ := C.C_NSMenuItem_Tag(n.Ptr())
	return int(result_)
}

func (n *NSMenuItem) SetTag(value int) {
	C.C_NSMenuItem_SetTag(n.Ptr(), C.int(value))
}

func (n *NSMenuItem) State() ControlStateValue {
	result_ := C.C_NSMenuItem_State(n.Ptr())
	return ControlStateValue(int(result_))
}

func (n *NSMenuItem) SetState(value ControlStateValue) {
	C.C_NSMenuItem_SetState(n.Ptr(), C.int(int(value)))
}

func (n *NSMenuItem) Image() Image {
	result_ := C.C_NSMenuItem_Image(n.Ptr())
	return MakeImage(result_)
}

func (n *NSMenuItem) SetImage(value Image) {
	C.C_NSMenuItem_SetImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSMenuItem) OnStateImage() Image {
	result_ := C.C_NSMenuItem_OnStateImage(n.Ptr())
	return MakeImage(result_)
}

func (n *NSMenuItem) SetOnStateImage(value Image) {
	C.C_NSMenuItem_SetOnStateImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSMenuItem) OffStateImage() Image {
	result_ := C.C_NSMenuItem_OffStateImage(n.Ptr())
	return MakeImage(result_)
}

func (n *NSMenuItem) SetOffStateImage(value Image) {
	C.C_NSMenuItem_SetOffStateImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSMenuItem) MixedStateImage() Image {
	result_ := C.C_NSMenuItem_MixedStateImage(n.Ptr())
	return MakeImage(result_)
}

func (n *NSMenuItem) SetMixedStateImage(value Image) {
	C.C_NSMenuItem_SetMixedStateImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSMenuItem) Submenu() Menu {
	result_ := C.C_NSMenuItem_Submenu(n.Ptr())
	return MakeMenu(result_)
}

func (n *NSMenuItem) SetSubmenu(value Menu) {
	C.C_NSMenuItem_SetSubmenu(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSMenuItem) HasSubmenu() bool {
	result_ := C.C_NSMenuItem_HasSubmenu(n.Ptr())
	return bool(result_)
}

func (n *NSMenuItem) ParentItem() MenuItem {
	result_ := C.C_NSMenuItem_ParentItem(n.Ptr())
	return MakeMenuItem(result_)
}

func (n *NSMenuItem) IsSeparatorItem() bool {
	result_ := C.C_NSMenuItem_IsSeparatorItem(n.Ptr())
	return bool(result_)
}

func (n *NSMenuItem) Menu() Menu {
	result_ := C.C_NSMenuItem_Menu(n.Ptr())
	return MakeMenu(result_)
}

func (n *NSMenuItem) SetMenu(value Menu) {
	C.C_NSMenuItem_SetMenu(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSMenuItem) KeyEquivalent() string {
	result_ := C.C_NSMenuItem_KeyEquivalent(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSMenuItem) SetKeyEquivalent(value string) {
	C.C_NSMenuItem_SetKeyEquivalent(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSMenuItem) KeyEquivalentModifierMask() EventModifierFlags {
	result_ := C.C_NSMenuItem_KeyEquivalentModifierMask(n.Ptr())
	return EventModifierFlags(uint(result_))
}

func (n *NSMenuItem) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	C.C_NSMenuItem_SetKeyEquivalentModifierMask(n.Ptr(), C.uint(uint(value)))
}

func MenuItem_UsesUserKeyEquivalents() bool {
	result_ := C.C_NSMenuItem_MenuItem_UsesUserKeyEquivalents()
	return bool(result_)
}

func MenuItem_SetUsesUserKeyEquivalents(value bool) {
	C.C_NSMenuItem_MenuItem_SetUsesUserKeyEquivalents(C.bool(value))
}

func (n *NSMenuItem) UserKeyEquivalent() string {
	result_ := C.C_NSMenuItem_UserKeyEquivalent(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSMenuItem) IsAlternate() bool {
	result_ := C.C_NSMenuItem_IsAlternate(n.Ptr())
	return bool(result_)
}

func (n *NSMenuItem) SetAlternate(value bool) {
	C.C_NSMenuItem_SetAlternate(n.Ptr(), C.bool(value))
}

func (n *NSMenuItem) IndentationLevel() int {
	result_ := C.C_NSMenuItem_IndentationLevel(n.Ptr())
	return int(result_)
}

func (n *NSMenuItem) SetIndentationLevel(value int) {
	C.C_NSMenuItem_SetIndentationLevel(n.Ptr(), C.int(value))
}

func (n *NSMenuItem) ToolTip() string {
	result_ := C.C_NSMenuItem_ToolTip(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSMenuItem) SetToolTip(value string) {
	C.C_NSMenuItem_SetToolTip(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSMenuItem) RepresentedObject() objc.Object {
	result_ := C.C_NSMenuItem_RepresentedObject(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSMenuItem) SetRepresentedObject(value objc.Object) {
	C.C_NSMenuItem_SetRepresentedObject(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSMenuItem) View() View {
	result_ := C.C_NSMenuItem_View(n.Ptr())
	return MakeView(result_)
}

func (n *NSMenuItem) SetView(value View) {
	C.C_NSMenuItem_SetView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSMenuItem) IsHighlighted() bool {
	result_ := C.C_NSMenuItem_IsHighlighted(n.Ptr())
	return bool(result_)
}

func (n *NSMenuItem) AllowsKeyEquivalentWhenHidden() bool {
	result_ := C.C_NSMenuItem_AllowsKeyEquivalentWhenHidden(n.Ptr())
	return bool(result_)
}

func (n *NSMenuItem) SetAllowsKeyEquivalentWhenHidden(value bool) {
	C.C_NSMenuItem_SetAllowsKeyEquivalentWhenHidden(n.Ptr(), C.bool(value))
}
