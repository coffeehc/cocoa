package appkit

// #include "color_picking_custom.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ColorPickingCustom struct {
	SetColor                      func(newColor Color)                                     // required
	CurrentMode                   func() ColorPanelMode                                    // required
	SupportsMode                  func(mode ColorPanelMode) bool                           // required
	ProvideNewView                func(initialRequest bool) View                           // required
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

func WrapColorPickingCustom(delegate *ColorPickingCustom) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapColorPickingCustom(C.long(id))
	return objc.MakeObject(ptr)
}

//export colorPickingCustom_SetColor
func colorPickingCustom_SetColor(id int64, newColor unsafe.Pointer) {
	delegate := resources.Get(id).(*ColorPickingCustom)
	delegate.SetColor(MakeColor(newColor))
}

//export colorPickingCustom_CurrentMode
func colorPickingCustom_CurrentMode(id int64) C.int {
	delegate := resources.Get(id).(*ColorPickingCustom)
	result := delegate.CurrentMode()
	return C.int(int(result))
}

//export colorPickingCustom_SupportsMode
func colorPickingCustom_SupportsMode(id int64, mode C.int) C.bool {
	delegate := resources.Get(id).(*ColorPickingCustom)
	result := delegate.SupportsMode(ColorPanelMode(int(mode)))
	return C.bool(result)
}

//export colorPickingCustom_ProvideNewView
func colorPickingCustom_ProvideNewView(id int64, initialRequest C.bool) unsafe.Pointer {
	delegate := resources.Get(id).(*ColorPickingCustom)
	result := delegate.ProvideNewView(bool(initialRequest))
	return objc.ExtractPtr(result)
}

//export colorPickingCustom_InitWithPickerMask_ColorPanel
func colorPickingCustom_InitWithPickerMask_ColorPanel(id int64, mask C.uint, owningColorPanel unsafe.Pointer) unsafe.Pointer {
	delegate := resources.Get(id).(*ColorPickingCustom)
	result := delegate.InitWithPickerMask_ColorPanel(uint(mask), MakeColorPanel(owningColorPanel))
	return objc.ExtractPtr(result)
}

//export colorPickingCustom_SetMode
func colorPickingCustom_SetMode(id int64, mode C.int) {
	delegate := resources.Get(id).(*ColorPickingCustom)
	delegate.SetMode(ColorPanelMode(int(mode)))
}

//export colorPickingCustom_InsertNewButtonImage_In
func colorPickingCustom_InsertNewButtonImage_In(id int64, newButtonImage unsafe.Pointer, buttonCell unsafe.Pointer) {
	delegate := resources.Get(id).(*ColorPickingCustom)
	delegate.InsertNewButtonImage_In(MakeImage(newButtonImage), MakeButtonCell(buttonCell))
}

//export colorPickingCustom_ProvideNewButtonImage
func colorPickingCustom_ProvideNewButtonImage(id int64) unsafe.Pointer {
	delegate := resources.Get(id).(*ColorPickingCustom)
	result := delegate.ProvideNewButtonImage()
	return objc.ExtractPtr(result)
}

//export colorPickingCustom_MinContentSize
func colorPickingCustom_MinContentSize(id int64) C.CGSize {
	delegate := resources.Get(id).(*ColorPickingCustom)
	result := delegate.MinContentSize()
	return *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(result)))
}

//export colorPickingCustom_ButtonToolTip
func colorPickingCustom_ButtonToolTip(id int64) unsafe.Pointer {
	delegate := resources.Get(id).(*ColorPickingCustom)
	result := delegate.ButtonToolTip()
	return foundation.NewString(result).Ptr()
}

//export colorPickingCustom_AlphaControlAddedOrRemoved
func colorPickingCustom_AlphaControlAddedOrRemoved(id int64, sender unsafe.Pointer) {
	delegate := resources.Get(id).(*ColorPickingCustom)
	delegate.AlphaControlAddedOrRemoved(objc.MakeObject(sender))
}

//export colorPickingCustom_ViewSizeChanged
func colorPickingCustom_ViewSizeChanged(id int64, sender unsafe.Pointer) {
	delegate := resources.Get(id).(*ColorPickingCustom)
	delegate.ViewSizeChanged(objc.MakeObject(sender))
}

//export colorPickingCustom_AttachColorList
func colorPickingCustom_AttachColorList(id int64, colorList unsafe.Pointer) {
	delegate := resources.Get(id).(*ColorPickingCustom)
	delegate.AttachColorList(MakeColorList(colorList))
}

//export colorPickingCustom_DetachColorList
func colorPickingCustom_DetachColorList(id int64, colorList unsafe.Pointer) {
	delegate := resources.Get(id).(*ColorPickingCustom)
	delegate.DetachColorList(MakeColorList(colorList))
}

//export ColorPickingCustom_RespondsTo
func ColorPickingCustom_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*ColorPickingCustom)
	switch selName {
	case "setColor:":
		return delegate.SetColor != nil
	case "currentMode:":
		return delegate.CurrentMode != nil
	case "supportsMode:":
		return delegate.SupportsMode != nil
	case "provideNewView:":
		return delegate.ProvideNewView != nil
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

//export deleteColorPickingCustom
func deleteColorPickingCustom(id int64) {
	resources.Delete(id)
}
