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
	result_ := C.C_WKSecurityOrigin_AllocSecurityOrigin()
	return MakeSecurityOrigin(result_)
}

func (w WKSecurityOrigin) Autorelease() WKSecurityOrigin {
	result_ := C.C_WKSecurityOrigin_Autorelease(w.Ptr())
	return MakeSecurityOrigin(result_)
}

func (w WKSecurityOrigin) Retain() WKSecurityOrigin {
	result_ := C.C_WKSecurityOrigin_Retain(w.Ptr())
	return MakeSecurityOrigin(result_)
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
