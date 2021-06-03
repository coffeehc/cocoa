package appkit

// #include "color_picker.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ColorPicker interface {
	objc.Object
	InsertNewButtonImage_In(newButtonImage Image, buttonCell ButtonCell)
	SetMode(mode ColorPanelMode)
	AttachColorList(colorList ColorList)
	DetachColorList(colorList ColorList)
	ViewSizeChanged(sender objc.Object)
	ColorPanel() ColorPanel
	ProvideNewButtonImage() Image
	ButtonToolTip() string
	MinContentSize() foundation.Size
}

type NSColorPicker struct {
	objc.NSObject
}

func MakeColorPicker(ptr unsafe.Pointer) NSColorPicker {
	return NSColorPicker{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocColorPicker() NSColorPicker {
	return MakeColorPicker(C.C_ColorPicker_Alloc())
}

func (n NSColorPicker) InitWithPickerMask_ColorPanel(mask uint, owningColorPanel ColorPanel) ColorPicker {
	result_ := C.C_NSColorPicker_InitWithPickerMask_ColorPanel(n.Ptr(), C.uint(mask), objc.ExtractPtr(owningColorPanel))
	return MakeColorPicker(result_)
}

func (n NSColorPicker) Init() ColorPicker {
	result_ := C.C_NSColorPicker_Init(n.Ptr())
	return MakeColorPicker(result_)
}

func (n NSColorPicker) InsertNewButtonImage_In(newButtonImage Image, buttonCell ButtonCell) {
	C.C_NSColorPicker_InsertNewButtonImage_In(n.Ptr(), objc.ExtractPtr(newButtonImage), objc.ExtractPtr(buttonCell))
}

func (n NSColorPicker) SetMode(mode ColorPanelMode) {
	C.C_NSColorPicker_SetMode(n.Ptr(), C.int(int(mode)))
}

func (n NSColorPicker) AttachColorList(colorList ColorList) {
	C.C_NSColorPicker_AttachColorList(n.Ptr(), objc.ExtractPtr(colorList))
}

func (n NSColorPicker) DetachColorList(colorList ColorList) {
	C.C_NSColorPicker_DetachColorList(n.Ptr(), objc.ExtractPtr(colorList))
}

func (n NSColorPicker) ViewSizeChanged(sender objc.Object) {
	C.C_NSColorPicker_ViewSizeChanged(n.Ptr(), objc.ExtractPtr(sender))
}

func (n NSColorPicker) ColorPanel() ColorPanel {
	result_ := C.C_NSColorPicker_ColorPanel(n.Ptr())
	return MakeColorPanel(result_)
}

func (n NSColorPicker) ProvideNewButtonImage() Image {
	result_ := C.C_NSColorPicker_ProvideNewButtonImage(n.Ptr())
	return MakeImage(result_)
}

func (n NSColorPicker) ButtonToolTip() string {
	result_ := C.C_NSColorPicker_ButtonToolTip(n.Ptr())
	return foundation.MakeString(result_).String()
}

func (n NSColorPicker) MinContentSize() foundation.Size {
	result_ := C.C_NSColorPicker_MinContentSize(n.Ptr())
	return foundation.Size(coregraphics.FromCGSizePointer(unsafe.Pointer(&result_)))
}
