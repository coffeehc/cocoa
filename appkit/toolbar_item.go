package appkit

// #include "toolbar_item.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ToolbarItem interface {
	objc.Object
	Validate()
	ItemIdentifier() ToolbarItemIdentifier
	Toolbar() Toolbar
	Label() string
	SetLabel(value string)
	PaletteLabel() string
	SetPaletteLabel(value string)
	ToolTip() string
	SetToolTip(value string)
	Title() string
	SetTitle(value string)
	MenuFormRepresentation() MenuItem
	SetMenuFormRepresentation(value MenuItem)
	Tag() int
	SetTag(value int)
	Target() objc.Object
	SetTarget(value objc.Object)
	Action() *objc.Selector
	SetAction(value *objc.Selector)
	IsBordered() bool
	SetBordered(value bool)
	IsNavigational() bool
	SetNavigational(value bool)
	IsEnabled() bool
	SetEnabled(value bool)
	Image() Image
	SetImage(value Image)
	View() View
	SetView(value View)
	VisibilityPriority() ToolbarItemVisibilityPriority
	SetVisibilityPriority(value ToolbarItemVisibilityPriority)
	Autovalidates() bool
	SetAutovalidates(value bool)
	AllowsDuplicatesInToolbar() bool
}

type NSToolbarItem struct {
	objc.NSObject
}

func MakeToolbarItem(ptr unsafe.Pointer) *NSToolbarItem {
	if ptr == nil {
		return nil
	}
	return &NSToolbarItem{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocToolbarItem() *NSToolbarItem {
	return MakeToolbarItem(C.C_ToolbarItem_Alloc())
}

func (n *NSToolbarItem) InitWithItemIdentifier(itemIdentifier ToolbarItemIdentifier) ToolbarItem {
	result_ := C.C_NSToolbarItem_InitWithItemIdentifier(n.Ptr(), foundation.NewString(string(itemIdentifier)).Ptr())
	return MakeToolbarItem(result_)
}

func (n *NSToolbarItem) Validate() {
	C.C_NSToolbarItem_Validate(n.Ptr())
}

func (n *NSToolbarItem) ItemIdentifier() ToolbarItemIdentifier {
	result_ := C.C_NSToolbarItem_ItemIdentifier(n.Ptr())
	return ToolbarItemIdentifier(foundation.MakeString(result_).String())
}

func (n *NSToolbarItem) Toolbar() Toolbar {
	result_ := C.C_NSToolbarItem_Toolbar(n.Ptr())
	return MakeToolbar(result_)
}

func (n *NSToolbarItem) Label() string {
	result_ := C.C_NSToolbarItem_Label(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSToolbarItem) SetLabel(value string) {
	C.C_NSToolbarItem_SetLabel(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSToolbarItem) PaletteLabel() string {
	result_ := C.C_NSToolbarItem_PaletteLabel(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSToolbarItem) SetPaletteLabel(value string) {
	C.C_NSToolbarItem_SetPaletteLabel(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSToolbarItem) ToolTip() string {
	result_ := C.C_NSToolbarItem_ToolTip(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSToolbarItem) SetToolTip(value string) {
	C.C_NSToolbarItem_SetToolTip(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSToolbarItem) Title() string {
	result_ := C.C_NSToolbarItem_Title(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n *NSToolbarItem) SetTitle(value string) {
	C.C_NSToolbarItem_SetTitle(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n *NSToolbarItem) MenuFormRepresentation() MenuItem {
	result_ := C.C_NSToolbarItem_MenuFormRepresentation(n.Ptr())
	return MakeMenuItem(result_)
}

func (n *NSToolbarItem) SetMenuFormRepresentation(value MenuItem) {
	C.C_NSToolbarItem_SetMenuFormRepresentation(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSToolbarItem) Tag() int {
	result_ := C.C_NSToolbarItem_Tag(n.Ptr())
	return int(result_)
}

func (n *NSToolbarItem) SetTag(value int) {
	C.C_NSToolbarItem_SetTag(n.Ptr(), C.int(value))
}

func (n *NSToolbarItem) Target() objc.Object {
	result_ := C.C_NSToolbarItem_Target(n.Ptr())
	return objc.MakeObject(result_)
}

func (n *NSToolbarItem) SetTarget(value objc.Object) {
	C.C_NSToolbarItem_SetTarget(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSToolbarItem) Action() *objc.Selector {
	result_ := C.C_NSToolbarItem_Action(n.Ptr())
	return objc.MakeSelector(result_)
}

func (n *NSToolbarItem) SetAction(value *objc.Selector) {
	C.C_NSToolbarItem_SetAction(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSToolbarItem) IsBordered() bool {
	result_ := C.C_NSToolbarItem_IsBordered(n.Ptr())
	return bool(result_)
}

func (n *NSToolbarItem) SetBordered(value bool) {
	C.C_NSToolbarItem_SetBordered(n.Ptr(), C.bool(value))
}

func (n *NSToolbarItem) IsNavigational() bool {
	result_ := C.C_NSToolbarItem_IsNavigational(n.Ptr())
	return bool(result_)
}

func (n *NSToolbarItem) SetNavigational(value bool) {
	C.C_NSToolbarItem_SetNavigational(n.Ptr(), C.bool(value))
}

func (n *NSToolbarItem) IsEnabled() bool {
	result_ := C.C_NSToolbarItem_IsEnabled(n.Ptr())
	return bool(result_)
}

func (n *NSToolbarItem) SetEnabled(value bool) {
	C.C_NSToolbarItem_SetEnabled(n.Ptr(), C.bool(value))
}

func (n *NSToolbarItem) Image() Image {
	result_ := C.C_NSToolbarItem_Image(n.Ptr())
	return MakeImage(result_)
}

func (n *NSToolbarItem) SetImage(value Image) {
	C.C_NSToolbarItem_SetImage(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSToolbarItem) View() View {
	result_ := C.C_NSToolbarItem_View(n.Ptr())
	return MakeView(result_)
}

func (n *NSToolbarItem) SetView(value View) {
	C.C_NSToolbarItem_SetView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSToolbarItem) VisibilityPriority() ToolbarItemVisibilityPriority {
	result_ := C.C_NSToolbarItem_VisibilityPriority(n.Ptr())
	return ToolbarItemVisibilityPriority(int(result_))
}

func (n *NSToolbarItem) SetVisibilityPriority(value ToolbarItemVisibilityPriority) {
	C.C_NSToolbarItem_SetVisibilityPriority(n.Ptr(), C.int(int(value)))
}

func (n *NSToolbarItem) Autovalidates() bool {
	result_ := C.C_NSToolbarItem_Autovalidates(n.Ptr())
	return bool(result_)
}

func (n *NSToolbarItem) SetAutovalidates(value bool) {
	C.C_NSToolbarItem_SetAutovalidates(n.Ptr(), C.bool(value))
}

func (n *NSToolbarItem) AllowsDuplicatesInToolbar() bool {
	result_ := C.C_NSToolbarItem_AllowsDuplicatesInToolbar(n.Ptr())
	return bool(result_)
}
