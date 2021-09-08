package webkit

// #include "navigation_response.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type NavigationResponse interface {
	objc.Object
	Response() foundation.URLResponse
	CanShowMIMEType() bool
	IsForMainFrame() bool
}

type WKNavigationResponse struct {
	objc.NSObject
}

func MakeNavigationResponse(ptr unsafe.Pointer) WKNavigationResponse {
	return WKNavigationResponse{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocNavigationResponse() WKNavigationResponse {
	result_ := C.C_WKNavigationResponse_AllocNavigationResponse()
	return MakeNavigationResponse(result_)
}

func (w WKNavigationResponse) Init() WKNavigationResponse {
	result_ := C.C_WKNavigationResponse_Init(w.Ptr())
	return MakeNavigationResponse(result_)
}

func NewNavigationResponse() WKNavigationResponse {
	result_ := C.C_WKNavigationResponse_NewNavigationResponse()
	return MakeNavigationResponse(result_)
}

func (w WKNavigationResponse) Autorelease() WKNavigationResponse {
	result_ := C.C_WKNavigationResponse_Autorelease(w.Ptr())
	return MakeNavigationResponse(result_)
}

func (w WKNavigationResponse) Retain() WKNavigationResponse {
	result_ := C.C_WKNavigationResponse_Retain(w.Ptr())
	return MakeNavigationResponse(result_)
}

func (w WKNavigationResponse) Response() foundation.URLResponse {
	result_ := C.C_WKNavigationResponse_Response(w.Ptr())
	return foundation.MakeURLResponse(result_)
}

func (w WKNavigationResponse) CanShowMIMEType() bool {
	result_ := C.C_WKNavigationResponse_CanShowMIMEType(w.Ptr())
	return bool(result_)
}

func (w WKNavigationResponse) IsForMainFrame() bool {
	result_ := C.C_WKNavigationResponse_IsForMainFrame(w.Ptr())
	return bool(result_)
}
