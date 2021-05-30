package appkit

// #include "color_panel.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ColorPanel interface {
	Panel
	SetAction(selector *objc.Selector)
	SetTarget(target objc.Object)
	AttachColorList(colorList ColorList)
	DetachColorList(colorList ColorList)
	Mode() ColorPanelMode
	SetMode(value ColorPanelMode)
	AccessoryView() View
	SetAccessoryView(value View)
	IsContinuous() bool
	SetContinuous(value bool)
	ShowsAlpha() bool
	SetShowsAlpha(value bool)
	Color() Color
	SetColor(value Color)
	Alpha() coregraphics.Float
}

type NSColorPanel struct {
	NSPanel
}

func MakeColorPanel(ptr unsafe.Pointer) *NSColorPanel {
	if ptr == nil {
		return nil
	}
	return &NSColorPanel{
		NSPanel: *MakePanel(ptr),
	}
}

func AllocColorPanel() *NSColorPanel {
	return MakeColorPanel(C.C_ColorPanel_Alloc())
}

func (n *NSColorPanel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) ColorPanel {
	result_ := C.C_NSColorPanel_InitWithContentRect_StyleMask_Backing_Defer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag))
	return MakeColorPanel(result_)
}

func (n *NSColorPanel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen Screen) ColorPanel {
	result_ := C.C_NSColorPanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag), objc.ExtractPtr(screen))
	return MakeColorPanel(result_)
}

func (n *NSColorPanel) Init() ColorPanel {
	result_ := C.C_NSColorPanel_Init(n.Ptr())
	return MakeColorPanel(result_)
}

func (n *NSColorPanel) InitWithCoder(coder foundation.Coder) ColorPanel {
	result_ := C.C_NSColorPanel_InitWithCoder(n.Ptr(), objc.ExtractPtr(coder))
	return MakeColorPanel(result_)
}

func ColorPanel_SetPickerMode(mode ColorPanelMode) {
	C.C_NSColorPanel_ColorPanel_SetPickerMode(C.int(int(mode)))
}

func ColorPanel_SetPickerMask(mask ColorPanelOptions) {
	C.C_NSColorPanel_ColorPanel_SetPickerMask(C.uint(uint(mask)))
}

func (n *NSColorPanel) SetAction(selector *objc.Selector) {
	C.C_NSColorPanel_SetAction(n.Ptr(), objc.ExtractPtr(selector))
}

func (n *NSColorPanel) SetTarget(target objc.Object) {
	C.C_NSColorPanel_SetTarget(n.Ptr(), objc.ExtractPtr(target))
}

func (n *NSColorPanel) AttachColorList(colorList ColorList) {
	C.C_NSColorPanel_AttachColorList(n.Ptr(), objc.ExtractPtr(colorList))
}

func (n *NSColorPanel) DetachColorList(colorList ColorList) {
	C.C_NSColorPanel_DetachColorList(n.Ptr(), objc.ExtractPtr(colorList))
}

func ColorPanel_DragColor_WithEvent_FromView(color Color, event Event, sourceView View) bool {
	result_ := C.C_NSColorPanel_ColorPanel_DragColor_WithEvent_FromView(objc.ExtractPtr(color), objc.ExtractPtr(event), objc.ExtractPtr(sourceView))
	return bool(result_)
}

func SharedColorPanel() ColorPanel {
	result_ := C.C_NSColorPanel_SharedColorPanel()
	return MakeColorPanel(result_)
}

func ColorPanel_SharedColorPanelExists() bool {
	result_ := C.C_NSColorPanel_ColorPanel_SharedColorPanelExists()
	return bool(result_)
}

func (n *NSColorPanel) Mode() ColorPanelMode {
	result_ := C.C_NSColorPanel_Mode(n.Ptr())
	return ColorPanelMode(int(result_))
}

func (n *NSColorPanel) SetMode(value ColorPanelMode) {
	C.C_NSColorPanel_SetMode(n.Ptr(), C.int(int(value)))
}

func (n *NSColorPanel) AccessoryView() View {
	result_ := C.C_NSColorPanel_AccessoryView(n.Ptr())
	return MakeView(result_)
}

func (n *NSColorPanel) SetAccessoryView(value View) {
	C.C_NSColorPanel_SetAccessoryView(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSColorPanel) IsContinuous() bool {
	result_ := C.C_NSColorPanel_IsContinuous(n.Ptr())
	return bool(result_)
}

func (n *NSColorPanel) SetContinuous(value bool) {
	C.C_NSColorPanel_SetContinuous(n.Ptr(), C.bool(value))
}

func (n *NSColorPanel) ShowsAlpha() bool {
	result_ := C.C_NSColorPanel_ShowsAlpha(n.Ptr())
	return bool(result_)
}

func (n *NSColorPanel) SetShowsAlpha(value bool) {
	C.C_NSColorPanel_SetShowsAlpha(n.Ptr(), C.bool(value))
}

func (n *NSColorPanel) Color() Color {
	result_ := C.C_NSColorPanel_Color(n.Ptr())
	return MakeColor(result_)
}

func (n *NSColorPanel) SetColor(value Color) {
	C.C_NSColorPanel_SetColor(n.Ptr(), objc.ExtractPtr(value))
}

func (n *NSColorPanel) Alpha() coregraphics.Float {
	result_ := C.C_NSColorPanel_Alpha(n.Ptr())
	return coregraphics.Float(float64(result_))
}
