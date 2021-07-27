package appkit

// #include "color_picking_default.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type ColorPickingDefault struct {
	InitWithPickerMask_ColorPanel func(mask uint, owningColorPanel ColorPanel) objc.Object // required
	SetMode                       func(mode ColorPanelMode)                                // required
	InsertNewButtonImage_In       func(newButtonImage Image, buttonCell ButtonCell)        // required
	ProvideNewButtonImage         func() Image                                             // required
	MinContentSize                func() foundation.Size                                   // required
	ButtonToolTip                 func() string                                            // required
	AlphaControlAddedOrRemoved    func(sender objc.Object)                                 // required
	ViewSizeChanged               func(sender objc.Object)                                 // required
	AttachColorList               func(colorList ColorList)                                // required
	DetachColorList               func(colorList ColorList)                                // required
}

func (delegate *ColorPickingDefault) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapColorPickingDefault(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export colorPickingDefault_InitWithPickerMask_ColorPanel
func colorPickingDefault_InitWithPickerMask_ColorPanel(hp C.uintptr_t, mask C.uint, owningColorPanel unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*ColorPickingDefault)
	result := delegate.InitWithPickerMask_ColorPanel(uint(mask), MakeColorPanel(owningColorPanel))
	return objc.ExtractPtr(result)
}

//export colorPickingDefault_SetMode
func colorPickingDefault_SetMode(hp C.uintptr_t, mode C.int) {
	delegate := cgo.Handle(hp).Value().(*ColorPickingDefault)
	delegate.SetMode(ColorPanelMode(int(mode)))
}

//export colorPickingDefault_InsertNewButtonImage_In
func colorPickingDefault_InsertNewButtonImage_In(hp C.uintptr_t, newButtonImage unsafe.Pointer, buttonCell unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ColorPickingDefault)
	delegate.InsertNewButtonImage_In(MakeImage(newButtonImage), MakeButtonCell(buttonCell))
}

//export colorPickingDefault_ProvideNewButtonImage
func colorPickingDefault_ProvideNewButtonImage(hp C.uintptr_t) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*ColorPickingDefault)
	result := delegate.ProvideNewButtonImage()
	return objc.ExtractPtr(result)
}

//export colorPickingDefault_MinContentSize
func colorPickingDefault_MinContentSize(hp C.uintptr_t) C.CGSize {
	delegate := cgo.Handle(hp).Value().(*ColorPickingDefault)
	result := delegate.MinContentSize()
	return *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(result)))
}

//export colorPickingDefault_ButtonToolTip
func colorPickingDefault_ButtonToolTip(hp C.uintptr_t) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*ColorPickingDefault)
	result := delegate.ButtonToolTip()
	return foundation.NewString(result).Ptr()
}

//export colorPickingDefault_AlphaControlAddedOrRemoved
func colorPickingDefault_AlphaControlAddedOrRemoved(hp C.uintptr_t, sender unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ColorPickingDefault)
	delegate.AlphaControlAddedOrRemoved(objc.MakeObject(sender))
}

//export colorPickingDefault_ViewSizeChanged
func colorPickingDefault_ViewSizeChanged(hp C.uintptr_t, sender unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ColorPickingDefault)
	delegate.ViewSizeChanged(objc.MakeObject(sender))
}

//export colorPickingDefault_AttachColorList
func colorPickingDefault_AttachColorList(hp C.uintptr_t, colorList unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ColorPickingDefault)
	delegate.AttachColorList(MakeColorList(colorList))
}

//export colorPickingDefault_DetachColorList
func colorPickingDefault_DetachColorList(hp C.uintptr_t, colorList unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ColorPickingDefault)
	delegate.DetachColorList(MakeColorList(colorList))
}

//export ColorPickingDefault_RespondsTo
func ColorPickingDefault_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*ColorPickingDefault)
	switch selName {
	case "initWithPickerMask:colorPanel:":
		return delegate.InitWithPickerMask_ColorPanel != nil
	case "setMode:":
		return delegate.SetMode != nil
	case "insertNewButtonImage:in:":
		return delegate.InsertNewButtonImage_In != nil
	case "provideNewButtonImage:":
		return delegate.ProvideNewButtonImage != nil
	case "minContentSize:":
		return delegate.MinContentSize != nil
	case "buttonToolTip:":
		return delegate.ButtonToolTip != nil
	case "alphaControlAddedOrRemoved:":
		return delegate.AlphaControlAddedOrRemoved != nil
	case "viewSizeChanged:":
		return delegate.ViewSizeChanged != nil
	case "attachColorList:":
		return delegate.AttachColorList != nil
	case "detachColorList:":
		return delegate.DetachColorList != nil
	default:
		return false
	}
}

//export deleteColorPickingDefault
func deleteColorPickingDefault(hp C.uintptr_t) {
	cgo.Handle(hp).Delete()
}
