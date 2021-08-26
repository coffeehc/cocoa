package appkit

// #include "color_picking_custom.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
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

func (delegate *ColorPickingCustom) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapColorPickingCustom(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export colorPickingCustom_SetColor
func colorPickingCustom_SetColor(hp C.uintptr_t, newColor unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ColorPickingCustom)
	delegate.SetColor(MakeColor(newColor))
}

//export colorPickingCustom_CurrentMode
func colorPickingCustom_CurrentMode(hp C.uintptr_t) C.int {
	delegate := cgo.Handle(hp).Value().(*ColorPickingCustom)
	result := delegate.CurrentMode()
	return C.int(int(result))
}

//export colorPickingCustom_SupportsMode
func colorPickingCustom_SupportsMode(hp C.uintptr_t, mode C.int) C.bool {
	delegate := cgo.Handle(hp).Value().(*ColorPickingCustom)
	result := delegate.SupportsMode(ColorPanelMode(int(mode)))
	return C.bool(result)
}

//export colorPickingCustom_ProvideNewView
func colorPickingCustom_ProvideNewView(hp C.uintptr_t, initialRequest C.bool) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*ColorPickingCustom)
	result := delegate.ProvideNewView(bool(initialRequest))
	return objc.ExtractPtr(result)
}

//export colorPickingCustom_InitWithPickerMask_ColorPanel
func colorPickingCustom_InitWithPickerMask_ColorPanel(hp C.uintptr_t, mask C.uint, owningColorPanel unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*ColorPickingCustom)
	result := delegate.InitWithPickerMask_ColorPanel(uint(mask), MakeColorPanel(owningColorPanel))
	return objc.ExtractPtr(result)
}

//export colorPickingCustom_SetMode
func colorPickingCustom_SetMode(hp C.uintptr_t, mode C.int) {
	delegate := cgo.Handle(hp).Value().(*ColorPickingCustom)
	delegate.SetMode(ColorPanelMode(int(mode)))
}

//export colorPickingCustom_InsertNewButtonImage_In
func colorPickingCustom_InsertNewButtonImage_In(hp C.uintptr_t, newButtonImage unsafe.Pointer, buttonCell unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ColorPickingCustom)
	delegate.InsertNewButtonImage_In(MakeImage(newButtonImage), MakeButtonCell(buttonCell))
}

//export colorPickingCustom_ProvideNewButtonImage
func colorPickingCustom_ProvideNewButtonImage(hp C.uintptr_t) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*ColorPickingCustom)
	result := delegate.ProvideNewButtonImage()
	return objc.ExtractPtr(result)
}

//export colorPickingCustom_MinContentSize
func colorPickingCustom_MinContentSize(hp C.uintptr_t) C.CGSize {
	delegate := cgo.Handle(hp).Value().(*ColorPickingCustom)
	result := delegate.MinContentSize()
	return *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(result)))
}

//export colorPickingCustom_ButtonToolTip
func colorPickingCustom_ButtonToolTip(hp C.uintptr_t) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*ColorPickingCustom)
	result := delegate.ButtonToolTip()
	return foundation.NewString(result).Ptr()
}

//export colorPickingCustom_AlphaControlAddedOrRemoved
func colorPickingCustom_AlphaControlAddedOrRemoved(hp C.uintptr_t, sender unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ColorPickingCustom)
	delegate.AlphaControlAddedOrRemoved(objc.MakeObject(sender))
}

//export colorPickingCustom_ViewSizeChanged
func colorPickingCustom_ViewSizeChanged(hp C.uintptr_t, sender unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ColorPickingCustom)
	delegate.ViewSizeChanged(objc.MakeObject(sender))
}

//export colorPickingCustom_AttachColorList
func colorPickingCustom_AttachColorList(hp C.uintptr_t, colorList unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ColorPickingCustom)
	delegate.AttachColorList(MakeColorList(colorList))
}

//export colorPickingCustom_DetachColorList
func colorPickingCustom_DetachColorList(hp C.uintptr_t, colorList unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*ColorPickingCustom)
	delegate.DetachColorList(MakeColorList(colorList))
}

//export ColorPickingCustom_RespondsTo
func ColorPickingCustom_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*ColorPickingCustom)
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
