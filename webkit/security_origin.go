package webkit

// #include "security_origin.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type SecurityOrigin interface {
	objc.Object
	Host() string
	Port() int
	Protocol() string
}

type WKSecurityOrigin struct {
	objc.NSObject
}

func MakeSecurityOrigin(ptr unsafe.Pointer) WKSecurityOrigin {
	return WKSecurityOrigin{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocSecurityOrigin() WKSecurityOrigin {
	return MakeSecurityOrigin(C.C_SecurityOrigin_Alloc())
}

func (w WKSecurityOrigin) Host() string {
	result_ := C.C_WKSecurityOrigin_Host(w.Ptr())
	return foundation.MakeString(result_).String()
}

func (w WKSecurityOrigin) Port() int {
	result_ := C.C_WKSecurityOrigin_Port(w.Ptr())
	return int(result_)
}

func (w WKSecurityOrigin) Protocol() string {
	result_ := C.C_WKSecurityOrigin_Protocol(w.Ptr())
	return foundation.MakeString(result_).String()
}
