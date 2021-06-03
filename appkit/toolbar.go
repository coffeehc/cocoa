package appkit

// #include "toolbar.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Toolbar interface {
	objc.Object
	InsertItemWithItemIdentifier_AtIndex(itemIdentifier ToolbarItemIdentifier, index int)
	RemoveItemAtIndex(index int)
	RunCustomizationPalette(sender objc.Object)
	ValidateVisibleItems()
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	Identifier() ToolbarIdentifier
	DisplayMode() ToolbarDisplayMode
	SetDisplayMode(value ToolbarDisplayMode)
	ShowsBaselineSeparator() bool
	SetShowsBaselineSeparator(value bool)
	AllowsUserCustomization() bool
	SetAllowsUserCustomization(value bool)
	AllowsExtensionItems() bool
	SetAllowsExtensionItems(value bool)
	Items() []ToolbarItem
	VisibleItems() []ToolbarItem
	SizeMode() ToolbarSizeMode
	SetSizeMode(value ToolbarSizeMode)
	SelectedItemIdentifier() ToolbarItemIdentifier
	SetSelectedItemIdentifier(value ToolbarItemIdentifier)
	CenteredItemIdentifier() ToolbarItemIdentifier
	SetCenteredItemIdentifier(value ToolbarItemIdentifier)
	IsVisible() bool
	SetVisible(value bool)
	CustomizationPaletteIsRunning() bool
	AutosavesConfiguration() bool
	SetAutosavesConfiguration(value bool)
}

type NSToolbar struct {
	objc.NSObject
}

func MakeToolbar(ptr unsafe.Pointer) NSToolbar {
	return NSToolbar{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocToolbar() NSToolbar {
	return MakeToolbar(C.C_Toolbar_Alloc())
}

func (n NSToolbar) InitWithIdentifier(identifier ToolbarIdentifier) Toolbar {
	result_ := C.C_NSToolbar_InitWithIdentifier(n.Ptr(), foundation.NewString(string(identifier)).Ptr())
	return MakeToolbar(result_)
}

func (n NSToolbar) Init() Toolbar {
	result_ := C.C_NSToolbar_Init(n.Ptr())
	return MakeToolbar(result_)
}

func (n NSToolbar) InsertItemWithItemIdentifier_AtIndex(itemIdentifier ToolbarItemIdentifier, index int) {
	C.C_NSToolbar_InsertItemWithItemIdentifier_AtIndex(n.Ptr(), foundation.NewString(string(itemIdentifier)).Ptr(), C.int(index))
}

func (n NSToolbar) RemoveItemAtIndex(index int) {
	C.C_NSToolbar_RemoveItemAtIndex(n.Ptr(), C.int(index))
}

func (n NSToolbar) RunCustomizationPalette(sender objc.Object) {
	C.C_NSToolbar_RunCustomizationPalette(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSToolbar) ValidateVisibleItems() {
	C.C_NSToolbar_ValidateVisibleItems(n.Ptr())
}

func (n NSToolbar) Delegate() objc.Object {
	result_ := C.C_NSToolbar_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSToolbar) SetDelegate(value objc.Object) {
	C.C_NSToolbar_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSToolbar) Identifier() ToolbarIdentifier {
	result_ := C.C_NSToolbar_Identifier(n.Ptr())
	return ToolbarIdentifier(foundation.MakeString(result_).String())
}

func (n NSToolbar) DisplayMode() ToolbarDisplayMode {
	result_ := C.C_NSToolbar_DisplayMode(n.Ptr())
	return ToolbarDisplayMode(uint(result_))
}

func (n NSToolbar) SetDisplayMode(value ToolbarDisplayMode) {
	C.C_NSToolbar_SetDisplayMode(n.Ptr(), C.uint(uint(value)))
}

func (n NSToolbar) ShowsBaselineSeparator() bool {
	result_ := C.C_NSToolbar_ShowsBaselineSeparator(n.Ptr())
	return bool(result_)
}

func (n NSToolbar) SetShowsBaselineSeparator(value bool) {
	C.C_NSToolbar_SetShowsBaselineSeparator(n.Ptr(), C.bool(value))
}

func (n NSToolbar) AllowsUserCustomization() bool {
	result_ := C.C_NSToolbar_AllowsUserCustomization(n.Ptr())
	return bool(result_)
}

func (n NSToolbar) SetAllowsUserCustomization(value bool) {
	C.C_NSToolbar_SetAllowsUserCustomization(n.Ptr(), C.bool(value))
}

func (n NSToolbar) AllowsExtensionItems() bool {
	result_ := C.C_NSToolbar_AllowsExtensionItems(n.Ptr())
	return bool(result_)
}

func (n NSToolbar) SetAllowsExtensionItems(value bool) {
	C.C_NSToolbar_SetAllowsExtensionItems(n.Ptr(), C.bool(value))
}

func (n NSToolbar) Items() []ToolbarItem {
	result_ := C.C_NSToolbar_Items(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]ToolbarItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeToolbarItem(r)
	}
	return goResult_
}

func (n NSToolbar) VisibleItems() []ToolbarItem {
	result_ := C.C_NSToolbar_VisibleItems(n.Ptr())
	defer C.free(result_.data)
	result_Slice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result_.data))[:result_.len:result_.len]
	var goResult_ = make([]ToolbarItem, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeToolbarItem(r)
	}
	return goResult_
}

func (n NSToolbar) SizeMode() ToolbarSizeMode {
	result_ := C.C_NSToolbar_SizeMode(n.Ptr())
	return ToolbarSizeMode(uint(result_))
}

func (n NSToolbar) SetSizeMode(value ToolbarSizeMode) {
	C.C_NSToolbar_SetSizeMode(n.Ptr(), C.uint(uint(value)))
}

func (n NSToolbar) SelectedItemIdentifier() ToolbarItemIdentifier {
	result_ := C.C_NSToolbar_SelectedItemIdentifier(n.Ptr())
	return ToolbarItemIdentifier(foundation.MakeString(result_).String())
}

func (n NSToolbar) SetSelectedItemIdentifier(value ToolbarItemIdentifier) {
	C.C_NSToolbar_SetSelectedItemIdentifier(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n NSToolbar) CenteredItemIdentifier() ToolbarItemIdentifier {
	result_ := C.C_NSToolbar_CenteredItemIdentifier(n.Ptr())
	return ToolbarItemIdentifier(foundation.MakeString(result_).String())
}

func (n NSToolbar) SetCenteredItemIdentifier(value ToolbarItemIdentifier) {
	C.C_NSToolbar_SetCenteredItemIdentifier(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n NSToolbar) IsVisible() bool {
	result_ := C.C_NSToolbar_IsVisible(n.Ptr())
	return bool(result_)
}

func (n NSToolbar) SetVisible(value bool) {
	C.C_NSToolbar_SetVisible(n.Ptr(), C.bool(value))
}

func (n NSToolbar) CustomizationPaletteIsRunning() bool {
	result_ := C.C_NSToolbar_CustomizationPaletteIsRunning(n.Ptr())
	return bool(result_)
}

func (n NSToolbar) AutosavesConfiguration() bool {
	result_ := C.C_NSToolbar_AutosavesConfiguration(n.Ptr())
	return bool(result_)
}

func (n NSToolbar) SetAutosavesConfiguration(value bool) {
	C.C_NSToolbar_SetAutosavesConfiguration(n.Ptr(), C.bool(value))
}
