package appkit

// #include "font_panel.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
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

func AllocFontPanel() NSFontPanel {
	return MakeFontPanel(C.C_FontPanel_Alloc())
}

func (n NSFontPanel) InitWithContentRect_StyleMask_Backing_Defer(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) FontPanel {
	result_ := C.C_NSFontPanel_InitWithContentRect_StyleMask_Backing_Defer(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag))
	return MakeFontPanel(result_)
}

func (n NSFontPanel) InitWithContentRect_StyleMask_Backing_Defer_Screen(contentRect foundation.Rect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen Screen) FontPanel {
	result_ := C.C_NSFontPanel_InitWithContentRect_StyleMask_Backing_Defer_Screen(n.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(coregraphics.Rect(contentRect))), C.uint(uint(style)), C.uint(uint(backingStoreType)), C.bool(flag), objc.ExtractPtr(screen))
	return MakeFontPanel(result_)
}

func (n NSFontPanel) Init() FontPanel {
	result_ := C.C_NSFontPanel_Init(n.Ptr())
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
