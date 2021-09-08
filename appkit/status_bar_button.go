package appkit

// #include "status_bar_button.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type StatusBarButton interface {
	Button
	AppearsDisabled() bool
	SetAppearsDisabled(value bool)
}

type NSStatusBarButton struct {
	NSButton
}

func MakeStatusBarButton(ptr unsafe.Pointer) NSStatusBarButton {
	return NSStatusBarButton{
		NSButton: MakeButton(ptr),
	}
}

func StatusBarButton_CheckboxWithTitle_Target_Action(title string, target objc.Object, action objc.Selector) NSStatusBarButton {
	result_ := C.C_NSStatusBarButton_StatusBarButton_CheckboxWithTitle_Target_Action(foundation.NewString(title).Ptr(), objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakeStatusBarButton(result_)
}

func StatusBarButton_ButtonWithImage_Target_Action(image Image, target objc.Object, action objc.Selector) NSStatusBarButton {
	result_ := C.C_NSStatusBarButton_StatusBarButton_ButtonWithImage_Target_Action(objc.ExtractPtr(image), objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakeStatusBarButton(result_)
}

func StatusBarButton_RadioButtonWithTitle_Target_Action(title string, target objc.Object, action objc.Selector) NSStatusBarButton {
	result_ := C.C_NSStatusBarButton_StatusBarButton_RadioButtonWithTitle_Target_Action(foundation.NewString(title).Ptr(), objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakeStatusBarButton(result_)
}

func StatusBarButton_ButtonWithTitle_Image_Target_Action(title string, image Image, target objc.Object, action objc.Selector) NSStatusBarButton {
	result_ := C.C_NSStatusBarButton_StatusBarButton_ButtonWithTitle_Image_Target_Action(foundation.NewString(title).Ptr(), objc.ExtractPtr(image), objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakeStatusBarButton(result_)
}

func StatusBarButton_ButtonWithTitle_Target_Action(title string, target objc.Object, action objc.Selector) NSStatusBarButton {
	result_ := C.C_NSStatusBarButton_StatusBarButton_ButtonWithTitle_Target_Action(foundation.NewString(title).Ptr(), objc.ExtractPtr(target), unsafe.Pointer(action))
	return MakeStatusBarButton(result_)
}

func (n NSStatusBarButton) InitWithFrame(frameRect foundation.Rect) NSStatusBarButton {
	result_ := C.C_NSStatusBarButton_InitWithFrame(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(frameRect))))
	return MakeStatusBarButton(result_)
}

func (n NSStatusBarButton) InitWithCoder(coder foundation.Coder) NSStatusBarButton {
	result_ := C.C_NSStatusBarButton_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeStatusBarButton(result_)
}

func (n NSStatusBarButton) Init() NSStatusBarButton {
	result_ := C.C_NSStatusBarButton_Init(n.Ptr())
	return MakeStatusBarButton(result_)
}

func AllocStatusBarButton() NSStatusBarButton {
	result_ := C.C_NSStatusBarButton_AllocStatusBarButton()
	return MakeStatusBarButton(result_)
}

func NewStatusBarButton() NSStatusBarButton {
	result_ := C.C_NSStatusBarButton_NewStatusBarButton()
	return MakeStatusBarButton(result_)
}

func (n NSStatusBarButton) Autorelease() NSStatusBarButton {
	result_ := C.C_NSStatusBarButton_Autorelease(n.Ptr())
	return MakeStatusBarButton(result_)
}

func (n NSStatusBarButton) Retain() NSStatusBarButton {
	result_ := C.C_NSStatusBarButton_Retain(n.Ptr())
	return MakeStatusBarButton(result_)
}

func (n NSStatusBarButton) AppearsDisabled() bool {
	result_ := C.C_NSStatusBarButton_AppearsDisabled(n.Ptr())
	return bool(result_)
}

func (n NSStatusBarButton) SetAppearsDisabled(value bool) {
	C.C_NSStatusBarButton_SetAppearsDisabled(n.Ptr(), C.bool(value))
}
