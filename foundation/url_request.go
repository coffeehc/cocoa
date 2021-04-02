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
	CachePolicy() URLRequestCachePolicy
}

var _ URLRequest = (*NSURLRequest)(nil)

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

func (u *NSURLRequest) CachePolicy() URLRequestCachePolicy {
	return URLRequestCachePolicy(C.URLRequest_CachePolicy(u.Ptr()))
}

func NewURLRequest(url URL) URLRequest {
	return MakeURLRequest(C.URLRequest_NewURLRequest(toPointer(url)))
}

func RequestWithURL(url URL) URLRequest {
	return MakeURLRequest(C.URLRequest_RequestWithURL(toPointer(url)))
}
