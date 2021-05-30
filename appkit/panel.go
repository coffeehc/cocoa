package appkit

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
	SetFloatingPanel(value bool)
	BecomesKeyOnlyIfNeeded() bool
	SetBecomesKeyOnlyIfNeeded(value bool)
	SetWorksWhenModal(value bool)
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
	result_ := C.C_NSPanel_InitWithContentRect_StyleMask_Backing_Defer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag))
	return MakePanel(result_)
}

func (n *NSPanel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen Screen) Panel {
	result_ := C.C_NSPanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag), objc.ExtractPtr(screen))
	return MakePanel(result_)
}

func (n *NSPanel) Init() Panel {
	result_ := C.C_NSPanel_Init(n.Ptr())
	return MakePanel(result_)
}

func (n *NSPanel) SetFloatingPanel(value bool) {
	C.C_NSPanel_SetFloatingPanel(n.Ptr(), C.bool(value))
}

func (n *NSPanel) BecomesKeyOnlyIfNeeded() bool {
	result_ := C.C_NSPanel_BecomesKeyOnlyIfNeeded(n.Ptr())
	return bool(result_)
}

func (n *NSPanel) SetBecomesKeyOnlyIfNeeded(value bool) {
	C.C_NSPanel_SetBecomesKeyOnlyIfNeeded(n.Ptr(), C.bool(value))
}

func (n *NSPanel) SetWorksWhenModal(value bool) {
	C.C_NSPanel_SetWorksWhenModal(n.Ptr(), C.bool(value))
}
