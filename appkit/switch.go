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

func MakeSwitch(ptr unsafe.Pointer) NSSwitch {
	return NSSwitch{
		NSControl: MakeControl(ptr),
	}
}

func (n NSSwitch) InitWithFrame(frameRect foundation.Rect) NSSwitch {
	result_ := C.C_NSSwitch_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeSwitch(result_)
}

func (n NSSwitch) InitWithCoder(coder foundation.Coder) NSSwitch {
	result_ := C.C_NSSwitch_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeSwitch(result_)
}

func (n NSSwitch) Init() NSSwitch {
	result_ := C.C_NSSwitch_Init(n.Ptr())
	return MakeSwitch(result_)
}

func AllocSwitch() NSSwitch {
	result_ := C.C_NSSwitch_AllocSwitch()
	return MakeSwitch(result_)
}

func NewSwitch() NSSwitch {
	result_ := C.C_NSSwitch_NewSwitch()
	return MakeSwitch(result_)
}

func (n NSSwitch) Autorelease() NSSwitch {
	result_ := C.C_NSSwitch_Autorelease(n.Ptr())
	return MakeSwitch(result_)
}

func (n NSSwitch) Retain() NSSwitch {
	result_ := C.C_NSSwitch_Retain(n.Ptr())
	return MakeSwitch(result_)
}

func (n NSSwitch) State() ControlStateValue {
	result_ := C.C_NSSwitch_State(n.Ptr())
	return ControlStateValue(int(result_))
}

func (n NSSwitch) SetState(value ControlStateValue) {
	C.C_NSSwitch_SetState(n.Ptr(), C.int(int(value)))
}
