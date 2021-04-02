package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "open_panel.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type OpenPanel interface {
	SavePanel
	CanChooseFiles() bool
	SetCanChooseFiles(canChooseFiles bool)
	CanChooseDirectories() bool
	SetCanChooseDirectories(canChooseDirectories bool)
	ResolvesAliases() bool
	SetResolvesAliases(resolvesAliases bool)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(allowsMultipleSelection bool)
	IsAccessoryViewDisclosed() bool
	SetAccessoryViewDisclosed(accessoryViewDisclosed bool)
	URLs() []foundation.URL
	CanDownloadUbiquitousContents() bool
	SetCanDownloadUbiquitousContents(canDownloadUbiquitousContents bool)
	CanResolveUbiquitousConflicts() bool
	SetCanResolveUbiquitousConflicts(canResolveUbiquitousConflicts bool)
}

var _ OpenPanel = (*NSOpenPanel)(nil)

type NSOpenPanel struct {
	NSSavePanel
}

func MakeOpenPanel(ptr unsafe.Pointer) *NSOpenPanel {
	if ptr == nil {
		return nil
	}
	return &NSOpenPanel{
		NSSavePanel: *MakeSavePanel(ptr),
	}
}

func (o *NSOpenPanel) CanChooseFiles() bool {
	return bool(C.OpenPanel_CanChooseFiles(o.Ptr()))
}

func (o *NSOpenPanel) SetCanChooseFiles(canChooseFiles bool) {
	C.OpenPanel_SetCanChooseFiles(o.Ptr(), C.bool(canChooseFiles))
}

func (o *NSOpenPanel) CanChooseDirectories() bool {
	return bool(C.OpenPanel_CanChooseDirectories(o.Ptr()))
}

func (o *NSOpenPanel) SetCanChooseDirectories(canChooseDirectories bool) {
	C.OpenPanel_SetCanChooseDirectories(o.Ptr(), C.bool(canChooseDirectories))
}

func (o *NSOpenPanel) ResolvesAliases() bool {
	return bool(C.OpenPanel_ResolvesAliases(o.Ptr()))
}

func (o *NSOpenPanel) SetResolvesAliases(resolvesAliases bool) {
	C.OpenPanel_SetResolvesAliases(o.Ptr(), C.bool(resolvesAliases))
}

func (o *NSOpenPanel) AllowsMultipleSelection() bool {
	return bool(C.OpenPanel_AllowsMultipleSelection(o.Ptr()))
}

func (o *NSOpenPanel) SetAllowsMultipleSelection(allowsMultipleSelection bool) {
	C.OpenPanel_SetAllowsMultipleSelection(o.Ptr(), C.bool(allowsMultipleSelection))
}

func (o *NSOpenPanel) IsAccessoryViewDisclosed() bool {
	return bool(C.OpenPanel_IsAccessoryViewDisclosed(o.Ptr()))
}

func (o *NSOpenPanel) SetAccessoryViewDisclosed(accessoryViewDisclosed bool) {
	C.OpenPanel_SetAccessoryViewDisclosed(o.Ptr(), C.bool(accessoryViewDisclosed))
}

func (o *NSOpenPanel) URLs() []foundation.URL {
	var cArray C.Array = C.OpenPanel_URLs(o.Ptr())
	defer C.free(cArray.data)
	result := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(cArray.data))[:cArray.len:cArray.len]
	var goArray = make([]foundation.URL, len(result))
	for idx, r := range result {
		goArray[idx] = foundation.MakeURL(r)
	}
	return goArray
}

func (o *NSOpenPanel) CanDownloadUbiquitousContents() bool {
	return bool(C.OpenPanel_CanDownloadUbiquitousContents(o.Ptr()))
}

func (o *NSOpenPanel) SetCanDownloadUbiquitousContents(canDownloadUbiquitousContents bool) {
	C.OpenPanel_SetCanDownloadUbiquitousContents(o.Ptr(), C.bool(canDownloadUbiquitousContents))
}

func (o *NSOpenPanel) CanResolveUbiquitousConflicts() bool {
	return bool(C.OpenPanel_CanResolveUbiquitousConflicts(o.Ptr()))
}

func (o *NSOpenPanel) SetCanResolveUbiquitousConflicts(canResolveUbiquitousConflicts bool) {
	C.OpenPanel_SetCanResolveUbiquitousConflicts(o.Ptr(), C.bool(canResolveUbiquitousConflicts))
}

func NewOpenPanel() OpenPanel {
	return MakeOpenPanel(C.OpenPanel_NewOpenPanel())
}
