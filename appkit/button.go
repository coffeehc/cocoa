package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "button.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Button interface {
	Control
	Title() string
	SetTitle(title string)
	BezelStyle() BezelStyle
	SetBezelStyle(bezelStyle BezelStyle)
	State() ControlStateValue
	SetState(state ControlStateValue)
	SetButtonType(buttonType ButtonType)
	SetAction(handler ActionHandler)
}

var _ Button = (*NSButton)(nil)

type NSButton struct {
	NSControl
	action ActionHandler
}

func MakeButton(ptr unsafe.Pointer) *NSButton {
	if ptr == nil {
		return nil
	}
	return &NSButton{
		NSControl: *MakeControl(ptr),
	}
}
func (b *NSButton) setDelegate() {
	id := resources.Register(b)
	C.Button_SetDelegate(b.Ptr(), C.long(id))
}

func (b *NSButton) Title() string {
	return C.GoString(C.Button_Title(b.Ptr()))
}

func (b *NSButton) SetTitle(title string) {
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))
	C.Button_SetTitle(b.Ptr(), cTitle)
}

func (b *NSButton) BezelStyle() BezelStyle {
	return BezelStyle(C.Button_BezelStyle(b.Ptr()))
}

func (b *NSButton) SetBezelStyle(bezelStyle BezelStyle) {
	C.Button_SetBezelStyle(b.Ptr(), C.ulong(bezelStyle))
}

func (b *NSButton) State() ControlStateValue {
	return ControlStateValue(C.Button_State(b.Ptr()))
}

func (b *NSButton) SetState(state ControlStateValue) {
	C.Button_SetState(b.Ptr(), C.long(state))
}

func NewButton(frame foundation.Rect) Button {
	instance := MakeButton(C.Button_NewButton(toNSRect(frame)))
	instance.setDelegate()
	return instance
}

func (b *NSButton) SetButtonType(buttonType ButtonType) {
	C.Button_SetButtonType(b.Ptr(), C.ulong(buttonType))
}

func (b *NSButton) SetAction(handler ActionHandler) {
	b.action = handler
}

//export Button_Target_Action
func Button_Target_Action(id int64, sender unsafe.Pointer) {
	button := resources.Get(id).(*NSButton)
	if button.action != nil {
		button.action(objc.MakeObject(sender))
	}
}
