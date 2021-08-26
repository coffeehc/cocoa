package webkit

// #include "ui_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type UIDelegate struct {
	WebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures func(webView WebView, configuration WebViewConfiguration, navigationAction NavigationAction, windowFeatures WindowFeatures) WebView
	WebViewDidClose                                                           func(webView WebView)
}

func (delegate *UIDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapUIDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export uIDelegate_WebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures
func uIDelegate_WebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures(hp C.uintptr_t, webView unsafe.Pointer, configuration unsafe.Pointer, navigationAction unsafe.Pointer, windowFeatures unsafe.Pointer) unsafe.Pointer {
	delegate := cgo.Handle(hp).Value().(*UIDelegate)
	result := delegate.WebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures(MakeWebView(webView), MakeWebViewConfiguration(configuration), MakeNavigationAction(navigationAction), MakeWindowFeatures(windowFeatures))
	return objc.ExtractPtr(result)
}

//export uIDelegate_WebViewDidClose
func uIDelegate_WebViewDidClose(hp C.uintptr_t, webView unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*UIDelegate)
	delegate.WebViewDidClose(MakeWebView(webView))
}

//export UIDelegate_RespondsTo
func UIDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*UIDelegate)
	switch selName {
	case "webView:createWebViewWithConfiguration:forNavigationAction:windowFeatures:":
		return delegate.WebView_CreateWebViewWithConfiguration_ForNavigationAction_WindowFeatures != nil
	case "webViewDidClose:":
		return delegate.WebViewDidClose != nil
	default:
		return false
	}
}

//export deleteUIDelegate
func deleteUIDelegate(hp C.uintptr_t) {
	cgo.Handle(hp).Delete()
}
