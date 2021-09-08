package webkit

// #include "process_pool.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ProcessPool interface {
	objc.Object
}

type WKProcessPool struct {
	objc.NSObject
}

func MakeProcessPool(ptr unsafe.Pointer) WKProcessPool {
	return WKProcessPool{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocProcessPool() WKProcessPool {
	result_ := C.C_WKProcessPool_AllocProcessPool()
	return MakeProcessPool(result_)
}

func (w WKProcessPool) Init() WKProcessPool {
	result_ := C.C_WKProcessPool_Init(w.Ptr())
	return MakeProcessPool(result_)
}

func NewProcessPool() WKProcessPool {
	result_ := C.C_WKProcessPool_NewProcessPool()
	return MakeProcessPool(result_)
}

func (w WKProcessPool) Autorelease() WKProcessPool {
	result_ := C.C_WKProcessPool_Autorelease(w.Ptr())
	return MakeProcessPool(result_)
}

func (w WKProcessPool) Retain() WKProcessPool {
	result_ := C.C_WKProcessPool_Retain(w.Ptr())
	return MakeProcessPool(result_)
}
