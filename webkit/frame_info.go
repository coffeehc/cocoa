package webkit

// #include "frame_info.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type FrameInfo interface {
	objc.Object
	IsMainFrame() bool
	Request() foundation.URLRequest
	SecurityOrigin() SecurityOrigin
	WebView() WebView
}

type WKFrameInfo struct {
	objc.NSObject
}

func MakeFrameInfo(ptr unsafe.Pointer) WKFrameInfo {
	return WKFrameInfo{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocFrameInfo() WKFrameInfo {
	result_ := C.C_WKFrameInfo_AllocFrameInfo()
	return MakeFrameInfo(result_)
}

func (w WKFrameInfo) Init() WKFrameInfo {
	result_ := C.C_WKFrameInfo_Init(w.Ptr())
	return MakeFrameInfo(result_)
}

func NewFrameInfo() WKFrameInfo {
	result_ := C.C_WKFrameInfo_NewFrameInfo()
	return MakeFrameInfo(result_)
}

func (w WKFrameInfo) Autorelease() WKFrameInfo {
	result_ := C.C_WKFrameInfo_Autorelease(w.Ptr())
	return MakeFrameInfo(result_)
}

func (w WKFrameInfo) Retain() WKFrameInfo {
	result_ := C.C_WKFrameInfo_Retain(w.Ptr())
	return MakeFrameInfo(result_)
}

func (w WKFrameInfo) IsMainFrame() bool {
	result_ := C.C_WKFrameInfo_IsMainFrame(w.Ptr())
	return bool(result_)
}

func (w WKFrameInfo) Request() foundation.URLRequest {
	result_ := C.C_WKFrameInfo_Request(w.Ptr())
	return foundation.MakeURLRequest(result_)
}

func (w WKFrameInfo) SecurityOrigin() SecurityOrigin {
	result_ := C.C_WKFrameInfo_SecurityOrigin(w.Ptr())
	return MakeSecurityOrigin(result_)
}

func (w WKFrameInfo) WebView() WebView {
	result_ := C.C_WKFrameInfo_WebView(w.Ptr())
	return MakeWebView(result_)
}
