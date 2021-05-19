package appkit

// #include "switch.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Switch interface {
	Control
	State() ControlStateValue
	SetState(value ControlStateValue)
}

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

func AllocSwitch() *NSSwitch {
	return MakeSwitch(C.C_Switch_Alloc())
}

func (n *NSSwitch) InitWithFrame(frameRect foundation.Rect) Switch {
	result_ := C.C_NSSwitch_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeSwitch(result_)
}

func (n *NSSwitch) InitWithCoder(coder foundation.Coder) Switch {
	result_ := C.C_NSSwitch_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeSwitch(result_)
}

func (n *NSSwitch) Init() Switch {
	result_ := C.C_NSSwitch_Init(n.Ptr())
	return MakeSwitch(result_)
}

func (n *NSSwitch) State() ControlStateValue {
	result_ := C.C_NSSwitch_State(n.Ptr())
	return ControlStateValue(int(result_))
}

func (n *NSSwitch) SetState(value ControlStateValue) {
	C.C_NSSwitch_SetState(n.Ptr(), C.int(int(value)))
}
