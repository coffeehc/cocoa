package webkit

// #include "download.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Download interface {
	objc.Object
	Delegate() objc.Object
	SetDelegate(value objc.Object)
	OriginalRequest() foundation.URLRequest
	WebView() WebView
}

type WKDownload struct {
	objc.NSObject
}

func MakeDownload(ptr unsafe.Pointer) WKDownload {
	return WKDownload{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocDownload() WKDownload {
	return MakeDownload(C.C_Download_Alloc())
}

func (w WKDownload) Init() Download {
	result_ := C.C_WKDownload_Init(w.Ptr())
	return MakeDownload(result_)
}

func (w WKDownload) Delegate() objc.Object {
	result_ := C.C_WKDownload_Delegate(w.Ptr())
	return objc.MakeObject(result_)
}

func (w WKDownload) SetDelegate(value objc.Object) {
	C.C_WKDownload_SetDelegate(w.Ptr(), objc.ExtractPtr(value))
}

func (w WKDownload) OriginalRequest() foundation.URLRequest {
	result_ := C.C_WKDownload_OriginalRequest(w.Ptr())
	return foundation.MakeURLRequest(result_)
}

func (w WKDownload) WebView() WebView {
	result_ := C.C_WKDownload_WebView(w.Ptr())
	return MakeWebView(result_)
}
