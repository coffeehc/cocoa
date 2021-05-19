package appkit

// #include "text_tab.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type TextTab interface {
	objc.Object
	Location() coregraphics.Float
	Alignment() TextAlignment
}

type NSTextTab struct {
	objc.NSObject
}

func MakeTextTab(ptr unsafe.Pointer) *NSTextTab {
	if ptr == nil {
		return nil
	}
	return &NSTextTab{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocTextTab() *NSTextTab {
	return MakeTextTab(C.C_TextTab_Alloc())
}

func (n *NSTextTab) Init() TextTab {
	result_ := C.C_NSTextTab_Init(n.Ptr())
	return MakeTextTab(result_)
}

func TextTab_ColumnTerminatorsForLocale(aLocale foundation.Locale) foundation.CharacterSet {
	result_ := C.C_NSTextTab_TextTab_ColumnTerminatorsForLocale(objc.ExtractPtr(aLocale))
	return foundation.MakeCharacterSet(result_)
}

func (n *NSTextTab) Location() coregraphics.Float {
	result_ := C.C_NSTextTab_Location(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n *NSTextTab) Alignment() TextAlignment {
	result_ := C.C_NSTextTab_Alignment(n.Ptr())
	return TextAlignment(int(result_))
}
