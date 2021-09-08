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

func MakePanel(ptr unsafe.Pointer) NSPanel {
	return NSPanel{
		NSWindow: MakeWindow(ptr),
	}
}

func Panel_WindowWithContentViewController(contentViewController ViewController) NSPanel {
	result_ := C.C_NSPanel_Panel_WindowWithContentViewController(objc.ExtractPtr(contentViewController))
	return MakePanel(result_)
}

func (n NSPanel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) NSPanel {
	result_ := C.C_NSPanel_InitWithContentRect_StyleMask_Backing_Defer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag))
	return MakePanel(result_)
}

func (n NSPanel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen Screen) NSPanel {
	result_ := C.C_NSPanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag), objc.ExtractPtr(screen))
	return MakePanel(result_)
}

func (n NSPanel) Init() NSPanel {
	result_ := C.C_NSPanel_Init(n.Ptr())
	return MakePanel(result_)
}

func AllocPanel() NSPanel {
	result_ := C.C_NSPanel_AllocPanel()
	return MakePanel(result_)
}

func NewPanel() NSPanel {
	result_ := C.C_NSPanel_NewPanel()
	return MakePanel(result_)
}

func (n NSPanel) Autorelease() NSPanel {
	result_ := C.C_NSPanel_Autorelease(n.Ptr())
	return MakePanel(result_)
}

func (n NSPanel) Retain() NSPanel {
	result_ := C.C_NSPanel_Retain(n.Ptr())
	return MakePanel(result_)
}

func (n NSPanel) SetFloatingPanel(value bool) {
	C.C_NSPanel_SetFloatingPanel(n.Ptr(), C.bool(value))
}

func (n NSPanel) BecomesKeyOnlyIfNeeded() bool {
	result_ := C.C_NSPanel_BecomesKeyOnlyIfNeeded(n.Ptr())
	return bool(result_)
}

func (n NSPanel) SetBecomesKeyOnlyIfNeeded(value bool) {
	C.C_NSPanel_SetBecomesKeyOnlyIfNeeded(n.Ptr(), C.bool(value))
}

func (n NSPanel) SetWorksWhenModal(value bool) {
	C.C_NSPanel_SetWorksWhenModal(n.Ptr(), C.bool(value))
}
