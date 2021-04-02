package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
// #include "text_field.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"unsafe"
)

type TextField interface {
	Control
	IsBezeled() bool
	SetBezeled(bezeled bool)
	DrawsBackground() bool
	SetDrawsBackground(drawsBackground bool)
	IsEditable() bool
	SetEditable(editable bool)
	IsSelectable() bool
	SetSelectable(selectable bool)
	TextColor() Color
	SetTextColor(textColor Color)
	BackgroundColor() Color
	SetBackgroundColor(backgroundColor Color)
	ControlTextDidChange(callback func(notification foundation.Notification))
	_controlTextDidChange() func(notification foundation.Notification)
	ControlTextDidEndEditing(callback func(notification foundation.Notification))
	_controlTextDidEndEditing() func(notification foundation.Notification)
	ControlTextDidBeginEditing(callback func(notification foundation.Notification))
	_controlTextDidBeginEditing() func(notification foundation.Notification)
}

var _ TextField = (*NSTextField)(nil)

type NSTextField struct {
	NSControl
	controlTextDidChange       func(notification foundation.Notification)
	controlTextDidEndEditing   func(notification foundation.Notification)
	controlTextDidBeginEditing func(notification foundation.Notification)
}

func MakeTextField(ptr unsafe.Pointer) *NSTextField {
	if ptr == nil {
		return nil
	}
	return &NSTextField{
		NSControl: *MakeControl(ptr),
	}
}
func (t *NSTextField) setDelegate() {
	id := resources.Register(t)
	C.TextField_SetDelegate(t.Ptr(), C.long(id))
}

func (t *NSTextField) IsBezeled() bool {
	return bool(C.TextField_IsBezeled(t.Ptr()))
}

func (t *NSTextField) SetBezeled(bezeled bool) {
	C.TextField_SetBezeled(t.Ptr(), C.bool(bezeled))
}

func (t *NSTextField) DrawsBackground() bool {
	return bool(C.TextField_DrawsBackground(t.Ptr()))
}

func (t *NSTextField) SetDrawsBackground(drawsBackground bool) {
	C.TextField_SetDrawsBackground(t.Ptr(), C.bool(drawsBackground))
}

func (t *NSTextField) IsEditable() bool {
	return bool(C.TextField_IsEditable(t.Ptr()))
}

func (t *NSTextField) SetEditable(editable bool) {
	C.TextField_SetEditable(t.Ptr(), C.bool(editable))
}

func (t *NSTextField) IsSelectable() bool {
	return bool(C.TextField_IsSelectable(t.Ptr()))
}

func (t *NSTextField) SetSelectable(selectable bool) {
	C.TextField_SetSelectable(t.Ptr(), C.bool(selectable))
}

func (t *NSTextField) TextColor() Color {
	return MakeColor(C.TextField_TextColor(t.Ptr()))
}

func (t *NSTextField) SetTextColor(textColor Color) {
	C.TextField_SetTextColor(t.Ptr(), toPointer(textColor))
}

func (t *NSTextField) BackgroundColor() Color {
	return MakeColor(C.TextField_BackgroundColor(t.Ptr()))
}

func (t *NSTextField) SetBackgroundColor(backgroundColor Color) {
	C.TextField_SetBackgroundColor(t.Ptr(), toPointer(backgroundColor))
}

func NewTextField(frame foundation.Rect) TextField {
	instance := MakeTextField(C.TextField_NewTextField(toNSRect(frame)))
	instance.setDelegate()
	return instance
}

func (t *NSTextField) ControlTextDidChange(callback func(notification foundation.Notification)) {
	t.controlTextDidChange = callback
}

func (t *NSTextField) _controlTextDidChange() func(notification foundation.Notification) {
	return t.controlTextDidChange
}

func (t *NSTextField) ControlTextDidEndEditing(callback func(notification foundation.Notification)) {
	t.controlTextDidEndEditing = callback
}

func (t *NSTextField) _controlTextDidEndEditing() func(notification foundation.Notification) {
	return t.controlTextDidEndEditing
}

func (t *NSTextField) ControlTextDidBeginEditing(callback func(notification foundation.Notification)) {
	t.controlTextDidBeginEditing = callback
}

func (t *NSTextField) _controlTextDidBeginEditing() func(notification foundation.Notification) {
	return t.controlTextDidBeginEditing
}

//export TextField_Delegate_ControlTextDidChange
func TextField_Delegate_ControlTextDidChange(id int64, notification unsafe.Pointer) {
	textField := resources.Get(id).(TextField)
	callback := textField._controlTextDidChange()
	if callback != nil {
		callback(foundation.MakeNotification(notification))
	}
}

//export TextField_Delegate_ControlTextDidEndEditing
func TextField_Delegate_ControlTextDidEndEditing(id int64, notification unsafe.Pointer) {
	textField := resources.Get(id).(TextField)
	callback := textField._controlTextDidEndEditing()
	if callback != nil {
		callback(foundation.MakeNotification(notification))
	}
}

//export TextField_Delegate_ControlTextDidBeginEditing
func TextField_Delegate_ControlTextDidBeginEditing(id int64, notification unsafe.Pointer) {
	textField := resources.Get(id).(TextField)
	callback := textField._controlTextDidBeginEditing()
	if callback != nil {
		callback(foundation.MakeNotification(notification))
	}
}
