package foundation

// #include "url_response.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type URLResponse interface {
	objc.Object
	SuggestedFilename() string
	MIMEType() string
	TextEncodingName() string
	URL() URL
}

type NSURLResponse struct {
	objc.NSObject
}

func MakeURLResponse(ptr unsafe.Pointer) NSURLResponse {
	return NSURLResponse{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocURLResponse() NSURLResponse {
	return MakeURLResponse(C.C_URLResponse_Alloc())
}

func (n NSURLResponse) InitWithURL_MIMEType_ExpectedContentLength_TextEncodingName(URL URL, MIMEType string, length int, name string) URLResponse {
	result_ := C.C_NSURLResponse_InitWithURL_MIMEType_ExpectedContentLength_TextEncodingName(n.Ptr(), objc.ExtractPtr(URL), NewString(MIMEType).Ptr(), C.int(length), NewString(name).Ptr())
	return MakeURLResponse(result_)
}

func (n NSURLResponse) Init() URLResponse {
	result_ := C.C_NSURLResponse_Init(n.Ptr())
	return MakeURLResponse(result_)
}

func (n NSURLResponse) SuggestedFilename() string {
	result_ := C.C_NSURLResponse_SuggestedFilename(n.Ptr())
	return MakeString(result_).String()
}

func (n NSURLResponse) MIMEType() string {
	result_ := C.C_NSURLResponse_MIMEType(n.Ptr())
	return MakeString(result_).String()
}

func (n NSURLResponse) TextEncodingName() string {
	result_ := C.C_NSURLResponse_TextEncodingName(n.Ptr())
	return MakeString(result_).String()
}

func (n NSURLResponse) URL() URL {
	result_ := C.C_NSURLResponse_URL(n.Ptr())
	return MakeURL(result_)
}
