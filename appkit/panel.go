package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "panel.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Panel interface {
	Window
	BecomesKeyOnlyIfNeeded() bool
	SetBecomesKeyOnlyIfNeeded(value bool)
}

type NSPanel struct {
	NSWindow
}

func MakePanel(ptr unsafe.Pointer) *NSPanel {
	if ptr == nil {
		return nil
	}
	return &NSPanel{
		NSWindow: *MakeWindow(ptr),
	}
}

func AllocPanel() *NSPanel {
	return MakePanel(C.C_Panel_Alloc())
}

func (n *NSPanel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) Panel {
	result := C.C_NSPanel_InitWithContentRect_StyleMask_Backing_Defer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag))
	return MakePanel(result)
}

func (n *NSPanel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen Screen) Panel {
	result := C.C_NSPanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag), objc.ExtractPtr(screen))
	return MakePanel(result)
}

func (n *NSPanel) Init() Panel {
	result := C.C_NSPanel_Init(n.Ptr())
	return MakePanel(result)
}

func (n *NSPanel) BecomesKeyOnlyIfNeeded() bool {
	result := C.C_NSPanel_BecomesKeyOnlyIfNeeded(n.Ptr())
	return bool(result)
}

func (n *NSPanel) SetBecomesKeyOnlyIfNeeded(value bool) {
	C.C_NSPanel_SetBecomesKeyOnlyIfNeeded(n.Ptr(), C.bool(value))
}
