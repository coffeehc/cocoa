package appkit

// #include "color_picking.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type ColorPickingCustom interface {
	SetColor(newColor Color)
	CurrentMode() ColorPanelMode
	SupportsMode(mode ColorPanelMode) bool
	ProvideNewView(initialRequest bool) View
	InitWithPickerMask_ColorPanel(mask uint, owningColorPanel ColorPanel) objc.Object
	SetMode(mode ColorPanelMode)
	InsertNewButtonImage_In(newButtonImage Image, buttonCell ButtonCell)
	ProvideNewButtonImage() Image
	MinContentSize() foundation.Size
	ButtonToolTip() string
	AlphaControlAddedOrRemoved(sender objc.Object)
	ViewSizeChanged(sender objc.Object)
	AttachColorList(colorList ColorList)
	DetachColorList(colorList ColorList)
}

func ColorPickingCustomToObjc(protocol ColorPickingCustom) objc.Object {
	h := cgo.NewHandle(protocol)
	ptr := C.WrapColorPickingCustom(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export colorPickingCustom_SetColor
func colorPickingCustom_SetColor(hp C.uintptr_t, newColor unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(ColorPickingCustom)
	protocol.SetColor(MakeColor(newColor))
}

//export colorPickingCustom_CurrentMode
func colorPickingCustom_CurrentMode(hp C.uintptr_t) C.int {
	protocol := cgo.Handle(hp).Value().(ColorPickingCustom)
	result := protocol.CurrentMode()
	return C.int(int(result))
}

//export colorPickingCustom_SupportsMode
func colorPickingCustom_SupportsMode(hp C.uintptr_t, mode C.int) C.bool {
	protocol := cgo.Handle(hp).Value().(ColorPickingCustom)
	result := protocol.SupportsMode(ColorPanelMode(int(mode)))
	return C.bool(result)
}

//export colorPickingCustom_ProvideNewView
func colorPickingCustom_ProvideNewView(hp C.uintptr_t, initialRequest C.bool) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(ColorPickingCustom)
	result := protocol.ProvideNewView(bool(initialRequest))
	return objc.ExtractPtr(result)
}

//export colorPickingCustom_InitWithPickerMask_ColorPanel
func colorPickingCustom_InitWithPickerMask_ColorPanel(hp C.uintptr_t, mask C.uint, owningColorPanel unsafe.Pointer) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(ColorPickingCustom)
	result := protocol.InitWithPickerMask_ColorPanel(uint(mask), MakeColorPanel(owningColorPanel))
	return objc.ExtractPtr(result)
}

//export colorPickingCustom_SetMode
func colorPickingCustom_SetMode(hp C.uintptr_t, mode C.int) {
	protocol := cgo.Handle(hp).Value().(ColorPickingCustom)
	protocol.SetMode(ColorPanelMode(int(mode)))
}

//export colorPickingCustom_InsertNewButtonImage_In
func colorPickingCustom_InsertNewButtonImage_In(hp C.uintptr_t, newButtonImage unsafe.Pointer, buttonCell unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(ColorPickingCustom)
	protocol.InsertNewButtonImage_In(MakeImage(newButtonImage), MakeButtonCell(buttonCell))
}

//export colorPickingCustom_ProvideNewButtonImage
func colorPickingCustom_ProvideNewButtonImage(hp C.uintptr_t) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(ColorPickingCustom)
	result := protocol.ProvideNewButtonImage()
	return objc.ExtractPtr(result)
}

//export colorPickingCustom_MinContentSize
func colorPickingCustom_MinContentSize(hp C.uintptr_t) C.CGSize {
	protocol := cgo.Handle(hp).Value().(ColorPickingCustom)
	result := protocol.MinContentSize()
	return *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(result)))
}

//export colorPickingCustom_ButtonToolTip
func colorPickingCustom_ButtonToolTip(hp C.uintptr_t) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(ColorPickingCustom)
	result := protocol.ButtonToolTip()
	return foundation.NewString(result).Ptr()
}

//export colorPickingCustom_AlphaControlAddedOrRemoved
func colorPickingCustom_AlphaControlAddedOrRemoved(hp C.uintptr_t, sender unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(ColorPickingCustom)
	protocol.AlphaControlAddedOrRemoved(objc.MakeObject(sender))
}

//export colorPickingCustom_ViewSizeChanged
func colorPickingCustom_ViewSizeChanged(hp C.uintptr_t, sender unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(ColorPickingCustom)
	protocol.ViewSizeChanged(objc.MakeObject(sender))
}

//export colorPickingCustom_AttachColorList
func colorPickingCustom_AttachColorList(hp C.uintptr_t, colorList unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(ColorPickingCustom)
	protocol.AttachColorList(MakeColorList(colorList))
}

//export colorPickingCustom_DetachColorList
func colorPickingCustom_DetachColorList(hp C.uintptr_t, colorList unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(ColorPickingCustom)
	protocol.DetachColorList(MakeColorList(colorList))
}

//export ColorPickingCustom_RespondsTo
func ColorPickingCustom_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	protocol := cgo.Handle(hp).Value().(ColorPickingCustom)
	_ = protocol
	switch selName {
	case "setColor:":
		return true
	case "currentMode":
		return true
	case "supportsMode:":
		return true
	case "provideNewView:":
		return true
	case "initWithPickerMask:colorPanel:":
		return true
	case "setMode:":
		return true
	case "insertNewButtonImage:in:":
		return true
	case "provideNewButtonImage":
		return true
	case "minContentSize":
		return true
	case "buttonToolTip":
		return true
	case "alphaControlAddedOrRemoved:":
		return true
	case "viewSizeChanged:":
		return true
	case "attachColorList:":
		return true
	case "detachColorList:":
		return true
	default:
		return false
	}
}

type ColorPickingDefault interface {
	InitWithPickerMask_ColorPanel(mask uint, owningColorPanel ColorPanel) objc.Object
	SetMode(mode ColorPanelMode)
	InsertNewButtonImage_In(newButtonImage Image, buttonCell ButtonCell)
	ProvideNewButtonImage() Image
	MinContentSize() foundation.Size
	ButtonToolTip() string
	AlphaControlAddedOrRemoved(sender objc.Object)
	ViewSizeChanged(sender objc.Object)
	AttachColorList(colorList ColorList)
	DetachColorList(colorList ColorList)
}

func ColorPickingDefaultToObjc(protocol ColorPickingDefault) objc.Object {
	h := cgo.NewHandle(protocol)
	ptr := C.WrapColorPickingDefault(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export colorPickingDefault_InitWithPickerMask_ColorPanel
func colorPickingDefault_InitWithPickerMask_ColorPanel(hp C.uintptr_t, mask C.uint, owningColorPanel unsafe.Pointer) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(ColorPickingDefault)
	result := protocol.InitWithPickerMask_ColorPanel(uint(mask), MakeColorPanel(owningColorPanel))
	return objc.ExtractPtr(result)
}

//export colorPickingDefault_SetMode
func colorPickingDefault_SetMode(hp C.uintptr_t, mode C.int) {
	protocol := cgo.Handle(hp).Value().(ColorPickingDefault)
	protocol.SetMode(ColorPanelMode(int(mode)))
}

//export colorPickingDefault_InsertNewButtonImage_In
func colorPickingDefault_InsertNewButtonImage_In(hp C.uintptr_t, newButtonImage unsafe.Pointer, buttonCell unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(ColorPickingDefault)
	protocol.InsertNewButtonImage_In(MakeImage(newButtonImage), MakeButtonCell(buttonCell))
}

//export colorPickingDefault_ProvideNewButtonImage
func colorPickingDefault_ProvideNewButtonImage(hp C.uintptr_t) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(ColorPickingDefault)
	result := protocol.ProvideNewButtonImage()
	return objc.ExtractPtr(result)
}

//export colorPickingDefault_MinContentSize
func colorPickingDefault_MinContentSize(hp C.uintptr_t) C.CGSize {
	protocol := cgo.Handle(hp).Value().(ColorPickingDefault)
	result := protocol.MinContentSize()
	return *(*C.CGSize)(coregraphics.ToCGSizePointer(coregraphics.Size(result)))
}

//export colorPickingDefault_ButtonToolTip
func colorPickingDefault_ButtonToolTip(hp C.uintptr_t) unsafe.Pointer {
	protocol := cgo.Handle(hp).Value().(ColorPickingDefault)
	result := protocol.ButtonToolTip()
	return foundation.NewString(result).Ptr()
}

//export colorPickingDefault_AlphaControlAddedOrRemoved
func colorPickingDefault_AlphaControlAddedOrRemoved(hp C.uintptr_t, sender unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(ColorPickingDefault)
	protocol.AlphaControlAddedOrRemoved(objc.MakeObject(sender))
}

//export colorPickingDefault_ViewSizeChanged
func colorPickingDefault_ViewSizeChanged(hp C.uintptr_t, sender unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(ColorPickingDefault)
	protocol.ViewSizeChanged(objc.MakeObject(sender))
}

//export colorPickingDefault_AttachColorList
func colorPickingDefault_AttachColorList(hp C.uintptr_t, colorList unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(ColorPickingDefault)
	protocol.AttachColorList(MakeColorList(colorList))
}

//export colorPickingDefault_DetachColorList
func colorPickingDefault_DetachColorList(hp C.uintptr_t, colorList unsafe.Pointer) {
	protocol := cgo.Handle(hp).Value().(ColorPickingDefault)
	protocol.DetachColorList(MakeColorList(colorList))
}

//export ColorPickingDefault_RespondsTo
func ColorPickingDefault_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	protocol := cgo.Handle(hp).Value().(ColorPickingDefault)
	_ = protocol
	switch selName {
	case "initWithPickerMask:colorPanel:":
		return true
	case "setMode:":
		return true
	case "insertNewButtonImage:in:":
		return true
	case "provideNewButtonImage":
		return true
	case "minContentSize":
		return true
	case "buttonToolTip":
		return true
	case "alphaControlAddedOrRemoved:":
		return true
	case "viewSizeChanged:":
		return true
	case "attachColorList:":
		return true
	case "detachColorList:":
		return true
	default:
		return false
	}
}
