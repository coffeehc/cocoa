package appkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework AppKit
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
	result := C.C_NSTextTab_Init(n.Ptr())
	return MakeTextTab(result)
}

func TextTab_ColumnTerminatorsForLocale(aLocale foundation.Locale) foundation.CharacterSet {
	result := C.C_NSTextTab_TextTab_ColumnTerminatorsForLocale(objc.ExtractPtr(aLocale))
	return foundation.MakeCharacterSet(result)
}

func (n *NSTextTab) Location() coregraphics.Float {
	result := C.C_NSTextTab_Location(n.Ptr())
	return coregraphics.Float(float64(result))
}

func (n *NSTextTab) Alignment() TextAlignment {
	result := C.C_NSTextTab_Alignment(n.Ptr())
	return TextAlignment(int(result))
}
