package appkit

// #include "font_panel.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type FontPanel interface {
	Panel
	ReloadDefaultFontFamilies()
	SetPanelFont_IsMultiple(fontObj Font, flag bool)
	PanelConvertFont(fontObj Font) Font
	IsEnabled() bool
	SetEnabled(value bool)
	AccessoryView() View
	SetAccessoryView(value View)
}

type NSFontPanel struct {
	NSPanel
}

func MakeFontPanel(ptr unsafe.Pointer) NSFontPanel {
	return NSFontPanel{
		NSPanel: MakePanel(ptr),
	}
}

func FontPanel_WindowWithContentViewController(contentViewController ViewController) NSFontPanel {
	result_ := C.C_NSFontPanel_FontPanel_WindowWithContentViewController(objc.ExtractPtr(contentViewController))
	return MakeFontPanel(result_)
}

func (n NSFontPanel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) NSFontPanel {
	result_ := C.C_NSFontPanel_InitWithContentRect_StyleMask_Backing_Defer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag))
	return MakeFontPanel(result_)
}

func (n NSFontPanel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen Screen) NSFontPanel {
	result_ := C.C_NSFontPanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag), objc.ExtractPtr(screen))
	return MakeFontPanel(result_)
}

func (n NSFontPanel) Init() NSFontPanel {
	result_ := C.C_NSFontPanel_Init(n.Ptr())
	return MakeFontPanel(result_)
}

func AllocFontPanel() NSFontPanel {
	result_ := C.C_NSFontPanel_AllocFontPanel()
	return MakeFontPanel(result_)
}

func NewFontPanel() NSFontPanel {
	result_ := C.C_NSFontPanel_NewFontPanel()
	return MakeFontPanel(result_)
}

func (n NSFontPanel) Autorelease() NSFontPanel {
	result_ := C.C_NSFontPanel_Autorelease(n.Ptr())
	return MakeFontPanel(result_)
}

func (n NSFontPanel) Retain() NSFontPanel {
	result_ := C.C_NSFontPanel_Retain(n.Ptr())
	return MakeFontPanel(result_)
}

func (n NSFontPanel) ReloadDefaultFontFamilies() {
	C.C_NSFontPanel_ReloadDefaultFontFamilies(n.Ptr())
}

func (n NSFontPanel) SetPanelFont_IsMultiple(fontObj Font, flag bool) {
	C.C_NSFontPanel_SetPanelFont_IsMultiple(n.Ptr(), objc.ExtractPtr(fontObj), C.bool(flag))
}

func (n NSFontPanel) PanelConvertFont(fontObj Font) Font {
	result_ := C.C_NSFontPanel_PanelConvertFont(n.Ptr(), objc.ExtractPtr(fontObj))
	return MakeFont(result_)
}

func SharedFontPanel() FontPanel {
	result_ := C.C_NSFontPanel_SharedFontPanel()
	return MakeFontPanel(result_)
}

func SharedFontPanelExists() bool {
	result_ := C.C_NSFontPanel_SharedFontPanelExists()
	return bool(result_)
}

func (n NSFontPanel) IsEnabled() bool {
	result_ := C.C_NSFontPanel_IsEnabled(n.Ptr())
	return bool(result_)
}

func (n NSFontPanel) SetEnabled(value bool) {
	C.C_NSFontPanel_SetEnabled(n.Ptr(), C.bool(value))
}

func (n NSFontPanel) AccessoryView() View {
	result_ := C.C_NSFontPanel_AccessoryView(n.Ptr())
	return MakeView(result_)
}

func (n NSFontPanel) SetAccessoryView(value View) {
	C.C_NSFontPanel_SetAccessoryView(n.Ptr(), objc.ExtractPtr(value))
}

type FontChanging struct {
	ChangeFont             func(sender FontManager)
	ValidModesForFontPanel func(fontPanel FontPanel) FontPanelModeMask
}

func (delegate *FontChanging) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapFontChanging(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export fontChanging_ChangeFont
func fontChanging_ChangeFont(hp C.uintptr_t, sender unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*FontChanging)
	delegate.ChangeFont(MakeFontManager(sender))
}

//export fontChanging_ValidModesForFontPanel
func fontChanging_ValidModesForFontPanel(hp C.uintptr_t, fontPanel unsafe.Pointer) C.uint {
	delegate := cgo.Handle(hp).Value().(*FontChanging)
	result := delegate.ValidModesForFontPanel(MakeFontPanel(fontPanel))
	return C.uint(uint(result))
}

//export FontChanging_RespondsTo
func FontChanging_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*FontChanging)
	switch selName {
	case "changeFont:":
		return delegate.ChangeFont != nil
	case "validModesForFontPanel:":
		return delegate.ValidModesForFontPanel != nil
	default:
		return false
	}
}
