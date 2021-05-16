package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "open_panel.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type OpenPanel interface {
	SavePanel
	CanChooseFiles() bool
	SetCanChooseFiles(value bool)
	CanChooseDirectories() bool
	SetCanChooseDirectories(value bool)
	ResolvesAliases() bool
	SetResolvesAliases(value bool)
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	IsAccessoryViewDisclosed() bool
	SetAccessoryViewDisclosed(value bool)
	URLs() []foundation.URL
	CanDownloadUbiquitousContents() bool
	SetCanDownloadUbiquitousContents(value bool)
	CanResolveUbiquitousConflicts() bool
	SetCanResolveUbiquitousConflicts(value bool)
}

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

func AllocOpenPanel() *NSOpenPanel {
	return MakeOpenPanel(C.C_OpenPanel_Alloc())
}

func (n *NSOpenPanel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) OpenPanel {
	result := C.C_NSOpenPanel_InitWithContentRect_StyleMask_Backing_Defer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag))
	return MakeOpenPanel(result)
}

func (n *NSOpenPanel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen Screen) OpenPanel {
	result := C.C_NSOpenPanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag), objc.ExtractPtr(screen))
	return MakeOpenPanel(result)
}

func (n *NSOpenPanel) Init() OpenPanel {
	result := C.C_NSOpenPanel_Init(n.Ptr())
	return MakeOpenPanel(result)
}

func OpenPanel_() OpenPanel {
	result := C.C_NSOpenPanel_OpenPanel_()
	return MakeOpenPanel(result)
}

func (n *NSOpenPanel) CanChooseFiles() bool {
	result := C.C_NSOpenPanel_CanChooseFiles(n.Ptr())
	return bool(result)
}

func (n *NSOpenPanel) SetCanChooseFiles(value bool) {
	C.C_NSOpenPanel_SetCanChooseFiles(n.Ptr(), C.bool(value))
}

func (n *NSOpenPanel) CanChooseDirectories() bool {
	result := C.C_NSOpenPanel_CanChooseDirectories(n.Ptr())
	return bool(result)
}

func (n *NSOpenPanel) SetCanChooseDirectories(value bool) {
	C.C_NSOpenPanel_SetCanChooseDirectories(n.Ptr(), C.bool(value))
}

func (n *NSOpenPanel) ResolvesAliases() bool {
	result := C.C_NSOpenPanel_ResolvesAliases(n.Ptr())
	return bool(result)
}

func (n *NSOpenPanel) SetResolvesAliases(value bool) {
	C.C_NSOpenPanel_SetResolvesAliases(n.Ptr(), C.bool(value))
}

func (n *NSOpenPanel) AllowsMultipleSelection() bool {
	result := C.C_NSOpenPanel_AllowsMultipleSelection(n.Ptr())
	return bool(result)
}

func (n *NSOpenPanel) SetAllowsMultipleSelection(value bool) {
	C.C_NSOpenPanel_SetAllowsMultipleSelection(n.Ptr(), C.bool(value))
}

func (n *NSOpenPanel) IsAccessoryViewDisclosed() bool {
	result := C.C_NSOpenPanel_IsAccessoryViewDisclosed(n.Ptr())
	return bool(result)
}

func (n *NSOpenPanel) SetAccessoryViewDisclosed(value bool) {
	C.C_NSOpenPanel_SetAccessoryViewDisclosed(n.Ptr(), C.bool(value))
}

func (n *NSOpenPanel) URLs() []foundation.URL {
	result := C.C_NSOpenPanel_URLs(n.Ptr())
	defer C.free(result.data)
	resultSlice := (*[1 << 28]unsafe.Pointer)(unsafe.Pointer(result.data))[:result.len:result.len]
	var goResult = make([]foundation.URL, len(resultSlice))
	for idx, r := range resultSlice {
		goResult[idx] = foundation.MakeURL(r)
	}
	return goResult
}

func (n *NSOpenPanel) CanDownloadUbiquitousContents() bool {
	result := C.C_NSOpenPanel_CanDownloadUbiquitousContents(n.Ptr())
	return bool(result)
}

func (n *NSOpenPanel) SetCanDownloadUbiquitousContents(value bool) {
	C.C_NSOpenPanel_SetCanDownloadUbiquitousContents(n.Ptr(), C.bool(value))
}

func (n *NSOpenPanel) CanResolveUbiquitousConflicts() bool {
	result := C.C_NSOpenPanel_CanResolveUbiquitousConflicts(n.Ptr())
	return bool(result)
}

func (n *NSOpenPanel) SetCanResolveUbiquitousConflicts(value bool) {
	C.C_NSOpenPanel_SetCanResolveUbiquitousConflicts(n.Ptr(), C.bool(value))
}
