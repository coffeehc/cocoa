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
	return MakeProcessPool(C.C_ProcessPool_Alloc())
}

func (w WKProcessPool) Init() ProcessPool {
	result_ := C.C_WKProcessPool_Init(w.Ptr())
	return MakeProcessPool(result_)
}
