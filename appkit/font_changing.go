package appkit

// #include "font_changing.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type FontChanging struct {
	ChangeFont             func(sender FontManager)
	ValidModesForFontPanel func(fontPanel FontPanel) FontPanelModeMask
}

func WrapFontChanging(delegate *FontChanging) objc.Object {
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

//export deleteFontChanging
func deleteFontChanging(hp C.uintptr_t) {
	cgo.Handle(hp).Delete()
}
