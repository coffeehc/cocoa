package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "formatter.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Formatter interface {
	objc.Object
	StringForObjectValue(obj objc.Object) string
	EditingStringForObjectValue(obj objc.Object) string
}

type NSFormatter struct {
	objc.NSObject
}

func MakeFormatter(ptr unsafe.Pointer) *NSFormatter {
	if ptr == nil {
		return nil
	}
	return &NSFormatter{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocFormatter() *NSFormatter {
	return MakeFormatter(C.C_Formatter_Alloc())
}

func (n *NSFormatter) Init() Formatter {
	result_ := C.C_NSFormatter_Init(n.Ptr())
	return MakeFormatter(result_)
}

func (n *NSFormatter) StringForObjectValue(obj objc.Object) string {
	result_ := C.C_NSFormatter_StringForObjectValue(n.Ptr(), objc.ExtractPtr(obj))
	return MakeString(result_).String()
}

func (n *NSFormatter) EditingStringForObjectValue(obj objc.Object) string {
	result_ := C.C_NSFormatter_EditingStringForObjectValue(n.Ptr(), objc.ExtractPtr(obj))
	return MakeString(result_).String()
}
