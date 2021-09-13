package foundation

// #include "extension_context.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ExtensionContext interface {
	objc.Object
	CancelRequestWithError(error Error)
	InputItems() []objc.Object
}

type NSExtensionContext struct {
	objc.NSObject
}

func MakeExtensionContext(ptr unsafe.Pointer) NSExtensionContext {
	return NSExtensionContext{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocExtensionContext() NSExtensionContext {
	result_ := C.C_NSExtensionContext_AllocExtensionContext()
	return MakeExtensionContext(result_)
}

func (n NSExtensionContext) Init() NSExtensionContext {
	result_ := C.C_NSExtensionContext_Init(n.Ptr())
	return MakeExtensionContext(result_)
}

func NewExtensionContext() NSExtensionContext {
	result_ := C.C_NSExtensionContext_NewExtensionContext()
	return MakeExtensionContext(result_)
}

func (n NSExtensionContext) Autorelease() NSExtensionContext {
	result_ := C.C_NSExtensionContext_Autorelease(n.Ptr())
	return MakeExtensionContext(result_)
}

func (n NSExtensionContext) Retain() NSExtensionContext {
	result_ := C.C_NSExtensionContext_Retain(n.Ptr())
	return MakeExtensionContext(result_)
}

func (n NSExtensionContext) CancelRequestWithError(error Error) {
	C.C_NSExtensionContext_CancelRequestWithError(n.Ptr(), objc.ExtractPtr(error))
}

func (n NSExtensionContext) InputItems() []objc.Object {
	result_ := C.C_NSExtensionContext_InputItems(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.data)
	}
	result_Slice := unsafe.Slice((*unsafe.Pointer)(result_.data), int(result_.len))
	var goResult_ = make([]objc.Object, len(result_Slice))
	for idx, r := range result_Slice {
		goResult_[idx] = objc.MakeObject(r)
	}
	return goResult_
}
