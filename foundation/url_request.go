package foundation

// #include "url_request.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type URLRequest interface {
	objc.Object
	ValueForHTTPHeaderField(field string) string
	CachePolicy() URLRequestCachePolicy
	HTTPMethod() string
	URL() URL
	HTTPBody() []byte
	MainDocumentURL() URL
	AllHTTPHeaderFields() map[string]string
	TimeoutInterval() TimeInterval
	HTTPShouldHandleCookies() bool
	HTTPShouldUsePipelining() bool
	AllowsCellularAccess() bool
	AllowsConstrainedNetworkAccess() bool
	AllowsExpensiveNetworkAccess() bool
	NetworkServiceType() URLRequestNetworkServiceType
	AssumesHTTP3Capable() bool
}

type NSURLRequest struct {
	objc.NSObject
}

func MakeURLRequest(ptr unsafe.Pointer) NSURLRequest {
	return NSURLRequest{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocURLRequest() NSURLRequest {
	return MakeURLRequest(C.C_URLRequest_Alloc())
}

func (n NSURLRequest) InitWithURL(URL URL) URLRequest {
	result_ := C.C_NSURLRequest_InitWithURL(n.Ptr(), objc.ExtractPtr(URL))
	return MakeURLRequest(result_)
}

func (n NSURLRequest) InitWithURL_CachePolicy_TimeoutInterval(URL URL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) URLRequest {
	result_ := C.C_NSURLRequest_InitWithURL_CachePolicy_TimeoutInterval(n.Ptr(), objc.ExtractPtr(URL), C.uint(uint(cachePolicy)), C.double(float64(timeoutInterval)))
	return MakeURLRequest(result_)
}

func (n NSURLRequest) Init() URLRequest {
	result_ := C.C_NSURLRequest_Init(n.Ptr())
	return MakeURLRequest(result_)
}

func URLRequest_RequestWithURL(URL URL) URLRequest {
	result_ := C.C_NSURLRequest_URLRequest_RequestWithURL(objc.ExtractPtr(URL))
	return MakeURLRequest(result_)
}

func URLRequest_RequestWithURL_CachePolicy_TimeoutInterval(URL URL, cachePolicy URLRequestCachePolicy, timeoutInterval TimeInterval) URLRequest {
	result_ := C.C_NSURLRequest_URLRequest_RequestWithURL_CachePolicy_TimeoutInterval(objc.ExtractPtr(URL), C.uint(uint(cachePolicy)), C.double(float64(timeoutInterval)))
	return MakeURLRequest(result_)
}

func (n NSURLRequest) ValueForHTTPHeaderField(field string) string {
	result_ := C.C_NSURLRequest_ValueForHTTPHeaderField(n.Ptr(), NewString(field).Ptr())
	return MakeString(result_).String()
}

func (n NSURLRequest) CachePolicy() URLRequestCachePolicy {
	result_ := C.C_NSURLRequest_CachePolicy(n.Ptr())
	return URLRequestCachePolicy(uint(result_))
}

func (n NSURLRequest) HTTPMethod() string {
	result_ := C.C_NSURLRequest_HTTPMethod(n.Ptr())
	return MakeString(result_).String()
}

func (n NSURLRequest) URL() URL {
	result_ := C.C_NSURLRequest_URL(n.Ptr())
	return MakeURL(result_)
}

func (n NSURLRequest) HTTPBody() []byte {
	result_ := C.C_NSURLRequest_HTTPBody(n.Ptr())
	var goResult_ []byte
	if result_.len > 0 {
		result_Buffer := unsafe.Slice((*byte)(result_.data), int(result_.len))
		goResult_ = make([]byte, C.int(result_.len))
		copy(goResult_, result_Buffer)
	}
	return goResult_
}

func (n NSURLRequest) MainDocumentURL() URL {
	result_ := C.C_NSURLRequest_MainDocumentURL(n.Ptr())
	return MakeURL(result_)
}

func (n NSURLRequest) AllHTTPHeaderFields() map[string]string {
	result_ := C.C_NSURLRequest_AllHTTPHeaderFields(n.Ptr())
	if result_.len > 0 {
		defer C.free(result_.key_data)
		defer C.free(result_.value_data)
	}
	result_KeySlice := unsafe.Slice((*unsafe.Pointer)(result_.key_data), int(result_.len))
	result_ValueSlice := unsafe.Slice((*unsafe.Pointer)(result_.value_data), int(result_.len))
	var goResult_ = make(map[string]string)
	for idx, k := range result_KeySlice {
		v := result_ValueSlice[idx]
		goResult_[MakeString(k).String()] = MakeString(v).String()
	}
	return goResult_
}

func (n NSURLRequest) TimeoutInterval() TimeInterval {
	result_ := C.C_NSURLRequest_TimeoutInterval(n.Ptr())
	return TimeInterval(float64(result_))
}

func (n NSURLRequest) HTTPShouldHandleCookies() bool {
	result_ := C.C_NSURLRequest_HTTPShouldHandleCookies(n.Ptr())
	return bool(result_)
}

func (n NSURLRequest) HTTPShouldUsePipelining() bool {
	result_ := C.C_NSURLRequest_HTTPShouldUsePipelining(n.Ptr())
	return bool(result_)
}

func (n NSURLRequest) AllowsCellularAccess() bool {
	result_ := C.C_NSURLRequest_AllowsCellularAccess(n.Ptr())
	return bool(result_)
}

func (n NSURLRequest) AllowsConstrainedNetworkAccess() bool {
	result_ := C.C_NSURLRequest_AllowsConstrainedNetworkAccess(n.Ptr())
	return bool(result_)
}

func (n NSURLRequest) AllowsExpensiveNetworkAccess() bool {
	result_ := C.C_NSURLRequest_AllowsExpensiveNetworkAccess(n.Ptr())
	return bool(result_)
}

func (n NSURLRequest) NetworkServiceType() URLRequestNetworkServiceType {
	result_ := C.C_NSURLRequest_NetworkServiceType(n.Ptr())
	return URLRequestNetworkServiceType(uint(result_))
}

func URLRequest_SupportsSecureCoding() bool {
	result_ := C.C_NSURLRequest_URLRequest_SupportsSecureCoding()
	return bool(result_)
}

func (n NSURLRequest) AssumesHTTP3Capable() bool {
	result_ := C.C_NSURLRequest_AssumesHTTP3Capable(n.Ptr())
	return bool(result_)
}
