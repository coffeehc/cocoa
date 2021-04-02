package webkit

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework WebKit
// #include "wk_web_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type WKWebView interface {
	appkit.View
	IsLoading() bool
	EstimatedProgress() float64
	StopLoading(sender objc.Object)
	Reload(sender objc.Object)
	ReloadFromOrigin(sender objc.Object)
	GoBack(sender objc.Object)
	GoForward(sender objc.Object)
	LoadRequest(request foundation.URLRequest)
	LoadHTMLString(string string, baseURL foundation.URL)
}

var _ WKWebView = (*NSWKWebView)(nil)

type NSWKWebView struct {
	appkit.NSView
}

func MakeWKWebView(ptr unsafe.Pointer) *NSWKWebView {
	if ptr == nil {
		return nil
	}
	return &NSWKWebView{
		NSView: *appkit.MakeView(ptr),
	}
}

func (w *NSWKWebView) IsLoading() bool {
	return bool(C.WKWebView_IsLoading(w.Ptr()))
}

func (w *NSWKWebView) EstimatedProgress() float64 {
	return float64(C.WKWebView_EstimatedProgress(w.Ptr()))
}

func NewWKWebView(frameRect foundation.Rect) WKWebView {
	return MakeWKWebView(C.WKWebView_NewWKWebView(toNSRect(frameRect)))
}

func (w *NSWKWebView) StopLoading(sender objc.Object) {
	C.WKWebView_StopLoading(w.Ptr(), toPointer(sender))
}

func (w *NSWKWebView) Reload(sender objc.Object) {
	C.WKWebView_Reload(w.Ptr(), toPointer(sender))
}

func (w *NSWKWebView) ReloadFromOrigin(sender objc.Object) {
	C.WKWebView_ReloadFromOrigin(w.Ptr(), toPointer(sender))
}

func (w *NSWKWebView) GoBack(sender objc.Object) {
	C.WKWebView_GoBack(w.Ptr(), toPointer(sender))
}

func (w *NSWKWebView) GoForward(sender objc.Object) {
	C.WKWebView_GoForward(w.Ptr(), toPointer(sender))
}

func (w *NSWKWebView) LoadRequest(request foundation.URLRequest) {
	C.WKWebView_LoadRequest(w.Ptr(), toPointer(request))
}

func (w *NSWKWebView) LoadHTMLString(string string, baseURL foundation.URL) {
	cString := C.CString(string)
	defer C.free(unsafe.Pointer(cString))
	C.WKWebView_LoadHTMLString(w.Ptr(), cString, toPointer(baseURL))
}
