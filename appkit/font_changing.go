package appkit

// #include "font_changing.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type FontChanging struct {
	ChangeFont             func(sender FontManager)
	ValidModesForFontPanel func(fontPanel FontPanel) FontPanelModeMask
}

func WrapFontChanging(delegate *FontChanging) objc.Object {
	id := resources.NextId()
	resources.Store(id, delegate)
	ptr := C.WrapFontChanging(C.long(id))
	return objc.MakeObject(ptr)
}

//export fontChanging_ChangeFont
func fontChanging_ChangeFont(id int64, sender unsafe.Pointer) {
	delegate := resources.Get(id).(*FontChanging)
	delegate.ChangeFont(MakeFontManager(sender))
}

//export fontChanging_ValidModesForFontPanel
func fontChanging_ValidModesForFontPanel(id int64, fontPanel unsafe.Pointer) C.uint {
	delegate := resources.Get(id).(*FontChanging)
	result := delegate.ValidModesForFontPanel(MakeFontPanel(fontPanel))
	return C.uint(uint(result))
}

//export FontChanging_RespondsTo
func FontChanging_RespondsTo(id int64, selectorPtr unsafe.Pointer) bool {
	sel := objc.MakeSelector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := resources.Get(id).(*FontChanging)
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
func deleteFontChanging(id int64) {
	resources.Delete(id)
}
