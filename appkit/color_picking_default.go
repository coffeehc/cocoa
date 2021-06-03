package appkit

// #include "color_picking_default.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
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

func WrapColorPickingDefault(delegate *ColorPickingDefault) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapColorPickingDefault(C.long(id))
	return objc.MakeObject(ptr)
}

//export colorPickingDefault_InitWithPickerMask_ColorPanel
func colorPickingDefault_InitWithPickerMask_ColorPanel(id int64, mask C.uint, owningColorPanel unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*ColorPickingDefault)
	result := delegate.InitWithPickerMask_ColorPanel(uint(mask), MakeColorPanel(owningColorPanel))
	return objc.ExtractPtr(result)
}

//export colorPickingDefault_SetMode
func colorPickingDefault_SetMode(id int64, mode C.int) {
	delegate := resources.Get(id).(*ColorPickingDefault)
	delegate.SetMode(ColorPanelMode(int(mode)))
}

//export colorPickingDefault_InsertNewButtonImage_In
func colorPickingDefault_InsertNewButtonImage_In(id int64, newButtonImage unsafe.Pointer, buttonCell unsafe.Pointer) {
	delegate := resources.Get(id).(*ColorPickingDefault)
	delegate.InsertNewButtonImage_In(MakeImage(newButtonImage), MakeButtonCell(buttonCell))
}

//export colorPickingDefault_ProvideNewButtonImage
func colorPickingDefault_ProvideNewButtonImage(id int64) unsafe.Pointer {
	delegate := resources.Get(id).(*ColorPickingDefault)
	result := delegate.ProvideNewButtonImage()
	return objc.ExtractPtr(result)
}

//export colorPickingDefault_MinContentSize
func colorPickingDefault_MinContentSize(id int64) C.CGSize {
	delegate := resources.Get(id).(*ColorPickingDefault)
	result := delegate.MinContentSize()
	return *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(result)))
}

//export colorPickingDefault_ButtonToolTip
func colorPickingDefault_ButtonToolTip(id int64) unsafe.Pointer {
	delegate := resources.Get(id).(*ColorPickingDefault)
	result := delegate.ButtonToolTip()
	return foundation.NewString(result).Ptr()
}

//export colorPickingDefault_AlphaControlAddedOrRemoved
func colorPickingDefault_AlphaControlAddedOrRemoved(id int64, sender unsafe.Pointer) {
	delegate := resources.Get(id).(*ColorPickingDefault)
	delegate.AlphaControlAddedOrRemoved(objc.MakeObject(sender))
}

//export colorPickingDefault_ViewSizeChanged
func colorPickingDefault_ViewSizeChanged(id int64, sender unsafe.Pointer) {
	delegate := resources.Get(id).(*ColorPickingDefault)
	delegate.ViewSizeChanged(objc.MakeObject(sender))
}

//export colorPickingDefault_AttachColorList
func colorPickingDefault_AttachColorList(id int64, colorList unsafe.Pointer) {
	delegate := resources.Get(id).(*ColorPickingDefault)
	delegate.AttachColorList(MakeColorList(colorList))
}

//export colorPickingDefault_DetachColorList
func colorPickingDefault_DetachColorList(id int64, colorList unsafe.Pointer) {
	delegate := resources.Get(id).(*ColorPickingDefault)
	delegate.DetachColorList(MakeColorList(colorList))
}

//export ColorPickingDefault_RespondsTo
func ColorPickingDefault_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*ColorPickingDefault)
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
func deleteColorPickingDefault(id int64) {
	resources.Delete(id)
}
