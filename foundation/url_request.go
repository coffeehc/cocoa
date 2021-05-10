package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "url_request.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type URLRequest interface {
	objc.Object
	ValueForHTTPHeaderField(field string) string
	HTTPMethod() string
	URL() URL
	HTTPBody() []byte
	MainDocumentURL() URL
	TimeoutInterval() TimeInterval
	HTTPShouldHandleCookies() bool
	HTTPShouldUsePipelining() bool
	AllowsCellularAccess() bool
	AllowsConstrainedNetworkAccess() bool
	AllowsExpensiveNetworkAccess() bool
	AssumesHTTP3Capable() bool
}

type NSURLRequest struct {
	objc.NSObject
}

func MakeURLRequest(ptr unsafe.Pointer) *NSURLRequest {
	if ptr == nil {
		return nil
	}
	return &NSURLRequest{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocURLRequest() *NSURLRequest {
	return MakeURLRequest(C.C_URLRequest_Alloc())
}

func (n *NSURLRequest) InitWithURL(URL URL) URLRequest {
	result := C.C_NSURLRequest_InitWithURL(n.Ptr(), objc.ExtractPtr(URL))
	return MakeURLRequest(result)
}

func (n *NSURLRequest) Init() URLRequest {
	result := C.C_NSURLRequest_Init(n.Ptr())
	return MakeURLRequest(result)
}

func RequestWithURL(URL URL) URLRequest {
	result := C.C_NSURLRequest_RequestWithURL(objc.ExtractPtr(URL))
	return MakeURLRequest(result)
}

func (n *NSURLRequest) ValueForHTTPHeaderField(field string) string {
	result := C.C_NSURLRequest_ValueForHTTPHeaderField(n.Ptr(), NewString(field).Ptr())
	return MakeString(result).String()
}

func (n *NSURLRequest) HTTPMethod() string {
	result := C.C_NSURLRequest_HTTPMethod(n.Ptr())
	return MakeString(result).String()
}

func (n *NSURLRequest) URL() URL {
	result := C.C_NSURLRequest_URL(n.Ptr())
	return MakeURL(result)
}

func (n *NSURLRequest) HTTPBody() []byte {
	result := C.C_NSURLRequest_HTTPBody(n.Ptr())
	resultBuffer := (*[1 << 30]byte)(result.data)[:C.int(result.len)]
	goResult := make([]byte, C.int(result.len))
	copy(goResult, resultBuffer)
	C.free(result.data)
	return goResult
}

func (n *NSURLRequest) MainDocumentURL() URL {
	result := C.C_NSURLRequest_MainDocumentURL(n.Ptr())
	return MakeURL(result)
}

func (n *NSURLRequest) TimeoutInterval() TimeInterval {
	result := C.C_NSURLRequest_TimeoutInterval(n.Ptr())
	return TimeInterval(float64(result))
}

func (n *NSURLRequest) HTTPShouldHandleCookies() bool {
	result := C.C_NSURLRequest_HTTPShouldHandleCookies(n.Ptr())
	return bool(result)
}

func (n *NSURLRequest) HTTPShouldUsePipelining() bool {
	result := C.C_NSURLRequest_HTTPShouldUsePipelining(n.Ptr())
	return bool(result)
}

func (n *NSURLRequest) AllowsCellularAccess() bool {
	result := C.C_NSURLRequest_AllowsCellularAccess(n.Ptr())
	return bool(result)
}

func (n *NSURLRequest) AllowsConstrainedNetworkAccess() bool {
	result := C.C_NSURLRequest_AllowsConstrainedNetworkAccess(n.Ptr())
	return bool(result)
}

func (n *NSURLRequest) AllowsExpensiveNetworkAccess() bool {
	result := C.C_NSURLRequest_AllowsExpensiveNetworkAccess(n.Ptr())
	return bool(result)
}

func (n *NSURLRequest) AssumesHTTP3Capable() bool {
	result := C.C_NSURLRequest_AssumesHTTP3Capable(n.Ptr())
	return bool(result)
}
