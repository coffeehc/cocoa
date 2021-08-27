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
	Options() map[TextTabOptionKey]objc.Object
}

type NSTextTab struct {
	objc.NSObject
}

func MakeTextTab(ptr unsafe.Pointer) NSTextTab {
	return NSTextTab{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocTextTab() NSTextTab {
	return MakeTextTab(C.C_TextTab_Alloc())
}

func (n NSTextTab) InitWithTextAlignment_Location_Options(alignment TextAlignment, loc coregraphics.Float, options map[TextTabOptionKey]objc.Object) TextTab {
	var cOptions C.Dictionary
	if len(options) > 0 {
		cOptionsKeyData := make([]unsafe.Pointer, len(options))
		cOptionsValueData := make([]unsafe.Pointer, len(options))
		var idx = 0
		for k, v := range options {
			cOptionsKeyData[idx] = foundation.NewString(string(k)).Ptr()
			cOptionsValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cOptions.key_data = unsafe.Pointer(&cOptionsKeyData[0])
		cOptions.value_data = unsafe.Pointer(&cOptionsValueData[0])
		cOptions.len = C.int(len(options))
	}
	result_ := C.C_NSTextTab_InitWithTextAlignment_Location_Options(n.Ptr(), C.int(int(alignment)), C.double(float64(loc)), cOptions)
	return MakeTextTab(result_)
}

func (n NSTextTab) Init() TextTab {
	result_ := C.C_NSTextTab_Init(n.Ptr())
	return MakeTextTab(result_)
}

func TextTab_ColumnTerminatorsForLocale(aLocale foundation.Locale) foundation.CharacterSet {
	result_ := C.C_NSTextTab_TextTab_ColumnTerminatorsForLocale(objc.ExtractPtr(aLocale))
	return foundation.MakeCharacterSet(result_)
}

func (n NSTextTab) Location() coregraphics.Float {
	result_ := C.C_NSTextTab_Location(n.Ptr())
	return coregraphics.Float(float64(result_))
}

func (n NSTextTab) Alignment() TextAlignment {
	result_ := C.C_NSTextTab_Alignment(n.Ptr())
	return TextAlignment(int(result_))
}

func (n NSTextTab) Options() map[TextTabOptionKey]objc.Object {
	result_ := C.C_NSTextTab_Options(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[TextTabOptionKey]objc.Object)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[TextTabOptionKey(foundation.MakeString(k).String())] = objc.MakeObject(v)
	}
	return goResult_
}
