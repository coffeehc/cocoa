package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "panel.h"
import "C"
import (
	"unsafe"
)

type Panel interface {
	Window
	SetFloatingPanel(floatingPanel bool)
	BecomesKeyOnlyIfNeeded() bool
	SetBecomesKeyOnlyIfNeeded(becomesKeyOnlyIfNeeded bool)
	SetWorksWhenModal(worksWhenModal bool)
}

var _ Panel = (*NSPanel)(nil)

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

func (p *NSPanel) SetFloatingPanel(floatingPanel bool) {
	C.Panel_SetFloatingPanel(p.Ptr(), C.bool(floatingPanel))
}

func (p *NSPanel) BecomesKeyOnlyIfNeeded() bool {
	return bool(C.Panel_BecomesKeyOnlyIfNeeded(p.Ptr()))
}

func (p *NSPanel) SetBecomesKeyOnlyIfNeeded(becomesKeyOnlyIfNeeded bool) {
	C.Panel_SetBecomesKeyOnlyIfNeeded(p.Ptr(), C.bool(becomesKeyOnlyIfNeeded))
}

func (p *NSPanel) SetWorksWhenModal(worksWhenModal bool) {
	C.Panel_SetWorksWhenModal(p.Ptr(), C.bool(worksWhenModal))
}
