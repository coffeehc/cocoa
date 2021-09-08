package foundation

// #include "formatter.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Formatter interface {
	objc.Object
	StringForObjectValue(obj objc.Object) string
	AttributedStringForObjectValue_WithDefaultAttributes(obj objc.Object, attrs map[AttributedStringKey]objc.Object) AttributedString
	EditingStringForObjectValue(obj objc.Object) string
}

type NSFormatter struct {
	objc.NSObject
}

func MakeFormatter(ptr unsafe.Pointer) NSFormatter {
	return NSFormatter{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocFormatter() NSFormatter {
	result_ := C.C_NSFormatter_AllocFormatter()
	return MakeFormatter(result_)
}

func (n NSFormatter) Init() NSFormatter {
	result_ := C.C_NSFormatter_Init(n.Ptr())
	return MakeFormatter(result_)
}

func NewFormatter() NSFormatter {
	result_ := C.C_NSFormatter_NewFormatter()
	return MakeFormatter(result_)
}

func (n NSFormatter) Autorelease() NSFormatter {
	result_ := C.C_NSFormatter_Autorelease(n.Ptr())
	return MakeFormatter(result_)
}

func (n NSFormatter) Retain() NSFormatter {
	result_ := C.C_NSFormatter_Retain(n.Ptr())
	return MakeFormatter(result_)
}

func (n NSFormatter) StringForObjectValue(obj objc.Object) string {
	result_ := C.C_NSFormatter_StringForObjectValue(n.Ptr(), objc.ExtractPtr(obj))
	return MakeString(result_).String()
}

func (n NSFormatter) AttributedStringForObjectValue_WithDefaultAttributes(obj objc.Object, attrs map[AttributedStringKey]objc.Object) AttributedString {
	var cAttrs C.Dictionary
	if len(attrs) > 0 {
		cAttrsKeyData := make([]unsafe.Pointer, len(attrs))
		cAttrsValueData := make([]unsafe.Pointer, len(attrs))
		var idx = 0
		for k, v := range attrs {
			cAttrsKeyData[idx] = NewString(string(k)).Ptr()
			cAttrsValueData[idx] = objc.ExtractPtr(v)
			idx++
		}
		cAttrs.key_data = unsafe.Pointer(&cAttrsKeyData[0])
		cAttrs.value_data = unsafe.Pointer(&cAttrsValueData[0])
		cAttrs.len = C.int(len(attrs))
	}
	result_ := C.C_NSFormatter_AttributedStringForObjectValue_WithDefaultAttributes(n.Ptr(), objc.ExtractPtr(obj), cAttrs)
	return MakeAttributedString(result_)
}

func (n NSFormatter) EditingStringForObjectValue(obj objc.Object) string {
	result_ := C.C_NSFormatter_EditingStringForObjectValue(n.Ptr(), objc.ExtractPtr(obj))
	return MakeString(result_).String()
}
