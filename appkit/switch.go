package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "switch.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type Switch interface {
	Control
	State() ControlStateValue
	SetState(state ControlStateValue)
}

var _ Switch = (*NSSwitch)(nil)

type NSSwitch struct {
	NSControl
}

func MakeSwitch(ptr unsafe.Pointer) *NSSwitch {
	if ptr == nil {
		return nil
	}
	return &NSSwitch{
		NSControl: *MakeControl(ptr),
	}
}

func (s *NSSwitch) State() ControlStateValue {
	return ControlStateValue(C.Switch_State(s.Ptr()))
}

func (s *NSSwitch) SetState(state ControlStateValue) {
	C.Switch_SetState(s.Ptr(), C.long(state))
}

func NewSwitch(frame foundation.Rect) Switch {
	return MakeSwitch(C.Switch_NewSwitch(toNSRect(frame)))
}
