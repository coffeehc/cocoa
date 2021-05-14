package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
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
	Identifier() ToolbarIdentifier
	ShowsBaselineSeparator() bool
	SetShowsBaselineSeparator(value bool)
	AllowsUserCustomization() bool
	SetAllowsUserCustomization(value bool)
	AllowsExtensionItems() bool
	SetAllowsExtensionItems(value bool)
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

func MakeToolbar(ptr unsafe.Pointer) *NSToolbar {
	if ptr == nil {
		return nil
	}
	return &NSToolbar{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocToolbar() *NSToolbar {
	return MakeToolbar(C.C_Toolbar_Alloc())
}

func (n *NSToolbar) InitWithIdentifier(identifier ToolbarIdentifier) Toolbar {
	result := C.C_NSToolbar_InitWithIdentifier(n.Ptr(), foundation.NewString(string(identifier)).Ptr())
	return MakeToolbar(result)
}

func (n *NSToolbar) Init() Toolbar {
	result := C.C_NSToolbar_Init(n.Ptr())
	return MakeToolbar(result)
}

func (n *NSToolbar) InsertItemWithItemIdentifier_AtIndex(itemIdentifier ToolbarItemIdentifier, index int) {
	C.C_NSToolbar_InsertItemWithItemIdentifier_AtIndex(n.Ptr(), foundation.NewString(string(itemIdentifier)).Ptr(), C.int(index))
}

func (n *NSToolbar) RemoveItemAtIndex(index int) {
	C.C_NSToolbar_RemoveItemAtIndex(n.Ptr(), C.int(index))
}

func (n *NSToolbar) RunCustomizationPalette(sender objc.Object) {
	C.C_NSToolbar_RunCustomizationPalette(n.Ptr(), objc.ExtractPtr(sender))
}

func (n *NSToolbar) ValidateVisibleItems() {
	C.C_NSToolbar_ValidateVisibleItems(n.Ptr())
}

func (n *NSToolbar) Identifier() ToolbarIdentifier {
	result := C.C_NSToolbar_Identifier(n.Ptr())
	return ToolbarIdentifier(foundation.MakeString(result).String())
}

func (n *NSToolbar) ShowsBaselineSeparator() bool {
	result := C.C_NSToolbar_ShowsBaselineSeparator(n.Ptr())
	return bool(result)
}

func (n *NSToolbar) SetShowsBaselineSeparator(value bool) {
	C.C_NSToolbar_SetShowsBaselineSeparator(n.Ptr(), C.bool(value))
}

func (n *NSToolbar) AllowsUserCustomization() bool {
	result := C.C_NSToolbar_AllowsUserCustomization(n.Ptr())
	return bool(result)
}

func (n *NSToolbar) SetAllowsUserCustomization(value bool) {
	C.C_NSToolbar_SetAllowsUserCustomization(n.Ptr(), C.bool(value))
}

func (n *NSToolbar) AllowsExtensionItems() bool {
	result := C.C_NSToolbar_AllowsExtensionItems(n.Ptr())
	return bool(result)
}

func (n *NSToolbar) SetAllowsExtensionItems(value bool) {
	C.C_NSToolbar_SetAllowsExtensionItems(n.Ptr(), C.bool(value))
}

func (n *NSToolbar) SelectedItemIdentifier() ToolbarItemIdentifier {
	result := C.C_NSToolbar_SelectedItemIdentifier(n.Ptr())
	return ToolbarItemIdentifier(foundation.MakeString(result).String())
}

func (n *NSToolbar) SetSelectedItemIdentifier(value ToolbarItemIdentifier) {
	C.C_NSToolbar_SetSelectedItemIdentifier(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n *NSToolbar) CenteredItemIdentifier() ToolbarItemIdentifier {
	result := C.C_NSToolbar_CenteredItemIdentifier(n.Ptr())
	return ToolbarItemIdentifier(foundation.MakeString(result).String())
}

func (n *NSToolbar) SetCenteredItemIdentifier(value ToolbarItemIdentifier) {
	C.C_NSToolbar_SetCenteredItemIdentifier(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n *NSToolbar) IsVisible() bool {
	result := C.C_NSToolbar_IsVisible(n.Ptr())
	return bool(result)
}

func (n *NSToolbar) SetVisible(value bool) {
	C.C_NSToolbar_SetVisible(n.Ptr(), C.bool(value))
}

func (n *NSToolbar) CustomizationPaletteIsRunning() bool {
	result := C.C_NSToolbar_CustomizationPaletteIsRunning(n.Ptr())
	return bool(result)
}

func (n *NSToolbar) AutosavesConfiguration() bool {
	result := C.C_NSToolbar_AutosavesConfiguration(n.Ptr())
	return bool(result)
}

func (n *NSToolbar) SetAutosavesConfiguration(value bool) {
	C.C_NSToolbar_SetAutosavesConfiguration(n.Ptr(), C.bool(value))
}
