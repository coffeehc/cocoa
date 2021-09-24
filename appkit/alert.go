package appkit

// #include "alert.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type Alert interface {
	objc.Object
	Layout()
	RunModal() ModalResponse
	AddButtonWithTitle(title string) Button
	AlertStyle() AlertStyle
	SetAlertStyle(value AlertStyle)
	AccessoryView() View
	SetAccessoryView(value View)
	ShowsHelp() bool
	SetShowsHelp(value bool)
	HelpAnchor() HelpAnchorName
	SetHelpAnchor(value HelpAnchorName)
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	SuppressionButton() Button
	ShowsSuppressionButton() bool
	SetShowsSuppressionButton(value bool)
	InformativeText() string
	SetInformativeText(value string)
	MessageText() string
	SetMessageText(value string)
	Icon() Image
	SetIcon(value Image)
	Buttons() []Button
	Window() Window
}

type NSAlert struct {
	objc.NSObject
}

func MakeAlert(ptr unsafe.Pointer) NSAlert {
	return NSAlert{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocAlert() NSAlert {
	result_ := C.C_NSAlert_AllocAlert()
	return MakeAlert(result_)
}

func (n NSAlert) Init() NSAlert {
	result_ := C.C_NSAlert_Init(n.Ptr())
	return MakeAlert(result_)
}

func NewAlert() NSAlert {
	result_ := C.C_NSAlert_NewAlert()
	return MakeAlert(result_)
}

func (n NSAlert) Autorelease() NSAlert {
	result_ := C.C_NSAlert_Autorelease(n.Ptr())
	return MakeAlert(result_)
}

func (n NSAlert) Retain() NSAlert {
	result_ := C.C_NSAlert_Retain(n.Ptr())
	return MakeAlert(result_)
}

func AlertWithError(error foundation.Error) Alert {
	result_ := C.C_NSAlert_AlertWithError(objc.ExtractPtr(error))
	return MakeAlert(result_)
}

func (n NSAlert) Layout() {
	C.C_NSAlert_Layout(n.Ptr())
}

func (n NSAlert) RunModal() ModalResponse {
	result_ := C.C_NSAlert_RunModal(n.Ptr())
	return ModalResponse(int(result_))
}

func (n NSAlert) AddButtonWithTitle(title string) Button {
	result_ := C.C_NSAlert_AddButtonWithTitle(n.Ptr(), foundation.NewString(title).Ptr())
	return MakeButton(result_)
}

func (n NSAlert) AlertStyle() AlertStyle {
	result_ := C.C_NSAlert_AlertStyle(n.Ptr())
	return AlertStyle(uint(result_))
}

func (n NSAlert) SetAlertStyle(value AlertStyle) {
	C.C_NSAlert_SetAlertStyle(n.Ptr(), C.uint(uint(value)))
}

func (n NSAlert) AccessoryView() View {
	result_ := C.C_NSAlert_AccessoryView(n.Ptr())
	return MakeView(result_)
}

func (n NSAlert) SetAccessoryView(value View) {
	C.C_NSAlert_SetAccessoryView(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSAlert) ShowsHelp() bool {
	result_ := C.C_NSAlert_ShowsHelp(n.Ptr())
	return bool(result_)
}

func (n NSAlert) SetShowsHelp(value bool) {
	C.C_NSAlert_SetShowsHelp(n.Ptr(), C.bool(value))
}

func (n NSAlert) HelpAnchor() HelpAnchorName {
	result_ := C.C_NSAlert_HelpAnchor(n.Ptr())
	return HelpAnchorName(foundation.MakeString(result_).String())
}

func (n NSAlert) SetHelpAnchor(value HelpAnchorName) {
	C.C_NSAlert_SetHelpAnchor(n.Ptr(), foundation.NewString(string(value)).Ptr())
}

func (n NSAlert) Delegate() objc.Object {
	result_ := C.C_NSAlert_Delegate(n.Ptr())
	return objc.MakeObject(result_)
}

func (n NSAlert) SetDelegate(value objc.Object) {
	C.C_NSAlert_SetDelegate(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSAlert) SuppressionButton() Button {
	result_ := C.C_NSAlert_SuppressionButton(n.Ptr())
	return MakeButton(result_)
}

func (n NSAlert) ShowsSuppressionButton() bool {
	result_ := C.C_NSAlert_ShowsSuppressionButton(n.Ptr())
	return bool(result_)
}

func (n NSAlert) SetShowsSuppressionButton(value bool) {
	C.C_NSAlert_SetShowsSuppressionButton(n.Ptr(), C.bool(value))
}

func (n NSAlert) InformativeText() string {
	result_ := C.C_NSAlert_InformativeText(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSAlert) SetInformativeText(value string) {
	C.C_NSAlert_SetInformativeText(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSAlert) MessageText() string {
	result_ := C.C_NSAlert_MessageText(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSAlert) SetMessageText(value string) {
	C.C_NSAlert_SetMessageText(n.Ptr(), foundation.NewString(value).Ptr())
}

func (n NSAlert) Icon() Image {
	result_ := C.C_NSAlert_Icon(n.Ptr())
	return MakeImage(result_)
}

func (n NSAlert) SetIcon(value Image) {
	C.C_NSAlert_SetIcon(n.Ptr(), objc.ExtractPtr(value))
}

func (n NSAlert) Buttons() []Button {
	result_ := C.C_NSAlert_Buttons(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]Button, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = MakeButton(r)
	}
	return goResult_
}

func (n NSAlert) Window() Window {
	result_ := C.C_NSAlert_Window(n.Ptr())
	return MakeWindow(result_)
}

type AlertDelegate struct {
	AlertShowHelp func(alert Alert) bool
}

func (delegate *AlertDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapAlertDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export alertDelegate_AlertShowHelp
func alertDelegate_AlertShowHelp(hp C.uintptr_t, alert unsafe.Pointer) C.bool {
	delegate := cgo.Handle(hp).Value().(*AlertDelegate)
	result := delegate.AlertShowHelp(MakeAlert(alert))
	return C.bool(result)
}

//export AlertDelegate_RespondsTo
func AlertDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*AlertDelegate)
	switch selName {
	case "alertShowHelp:":
		return delegate.AlertShowHelp != nil
	default:
		return false
	}
}
