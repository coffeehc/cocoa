package webkit

// #include "navigation_delegate.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type NavigationDelegate struct {
	WebView_DidStartProvisionalNavigation                    func(webView WebView, navigation Navigation)
	WebView_DidReceiveServerRedirectForProvisionalNavigation func(webView WebView, navigation Navigation)
	WebView_DidCommitNavigation                              func(webView WebView, navigation Navigation)
	WebView_DidFinishNavigation                              func(webView WebView, navigation Navigation)
	WebView_DidFailNavigation_WithError                      func(webView WebView, navigation Navigation, error foundation.Error)
	WebView_DidFailProvisionalNavigation_WithError           func(webView WebView, navigation Navigation, error foundation.Error)
	WebViewWebContentProcessDidTerminate                     func(webView WebView)
	WebView_NavigationAction_DidBecomeDownload               func(webView WebView, navigationAction NavigationAction, download Download)
	WebView_NavigationResponse_DidBecomeDownload             func(webView WebView, navigationResponse NavigationResponse, download Download)
}

func (delegate *NavigationDelegate) ToObjc() objc.Object {
	h := cgo.NewHandle(delegate)
	ptr := C.WrapNavigationDelegate(C.uintptr_t(h))
	return objc.MakeObject(ptr)
}

//export navigationDelegate_WebView_DidStartProvisionalNavigation
func navigationDelegate_WebView_DidStartProvisionalNavigation(hp C.uintptr_t, webView unsafe.Pointer, navigation unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*NavigationDelegate)
	delegate.WebView_DidStartProvisionalNavigation(MakeWebView(webView), MakeNavigation(navigation))
}

//export navigationDelegate_WebView_DidReceiveServerRedirectForProvisionalNavigation
func navigationDelegate_WebView_DidReceiveServerRedirectForProvisionalNavigation(hp C.uintptr_t, webView unsafe.Pointer, navigation unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*NavigationDelegate)
	delegate.WebView_DidReceiveServerRedirectForProvisionalNavigation(MakeWebView(webView), MakeNavigation(navigation))
}

//export navigationDelegate_WebView_DidCommitNavigation
func navigationDelegate_WebView_DidCommitNavigation(hp C.uintptr_t, webView unsafe.Pointer, navigation unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*NavigationDelegate)
	delegate.WebView_DidCommitNavigation(MakeWebView(webView), MakeNavigation(navigation))
}

//export navigationDelegate_WebView_DidFinishNavigation
func navigationDelegate_WebView_DidFinishNavigation(hp C.uintptr_t, webView unsafe.Pointer, navigation unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*NavigationDelegate)
	delegate.WebView_DidFinishNavigation(MakeWebView(webView), MakeNavigation(navigation))
}

//export navigationDelegate_WebView_DidFailNavigation_WithError
func navigationDelegate_WebView_DidFailNavigation_WithError(hp C.uintptr_t, webView unsafe.Pointer, navigation unsafe.Pointer, error unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*NavigationDelegate)
	delegate.WebView_DidFailNavigation_WithError(MakeWebView(webView), MakeNavigation(navigation), foundation.MakeError(error))
}

//export navigationDelegate_WebView_DidFailProvisionalNavigation_WithError
func navigationDelegate_WebView_DidFailProvisionalNavigation_WithError(hp C.uintptr_t, webView unsafe.Pointer, navigation unsafe.Pointer, error unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*NavigationDelegate)
	delegate.WebView_DidFailProvisionalNavigation_WithError(MakeWebView(webView), MakeNavigation(navigation), foundation.MakeError(error))
}

//export navigationDelegate_WebViewWebContentProcessDidTerminate
func navigationDelegate_WebViewWebContentProcessDidTerminate(hp C.uintptr_t, webView unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*NavigationDelegate)
	delegate.WebViewWebContentProcessDidTerminate(MakeWebView(webView))
}

//export navigationDelegate_WebView_NavigationAction_DidBecomeDownload
func navigationDelegate_WebView_NavigationAction_DidBecomeDownload(hp C.uintptr_t, webView unsafe.Pointer, navigationAction unsafe.Pointer, download unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*NavigationDelegate)
	delegate.WebView_NavigationAction_DidBecomeDownload(MakeWebView(webView), MakeNavigationAction(navigationAction), MakeDownload(download))
}

//export navigationDelegate_WebView_NavigationResponse_DidBecomeDownload
func navigationDelegate_WebView_NavigationResponse_DidBecomeDownload(hp C.uintptr_t, webView unsafe.Pointer, navigationResponse unsafe.Pointer, download unsafe.Pointer) {
	delegate := cgo.Handle(hp).Value().(*NavigationDelegate)
	delegate.WebView_NavigationResponse_DidBecomeDownload(MakeWebView(webView), MakeNavigationResponse(navigationResponse), MakeDownload(download))
}

//export NavigationDelegate_RespondsTo
func NavigationDelegate_RespondsTo(hp C.uintptr_t, selectorPtr unsafe.Pointer) bool {
	sel := objc.Selector(selectorPtr)
	selName := objc.Sel_GetName(sel)
	delegate := cgo.Handle(hp).Value().(*NavigationDelegate)
	switch selName {
	case "webView:didStartProvisionalNavigation:":
		return delegate.WebView_DidStartProvisionalNavigation != nil
	case "webView:didReceiveServerRedirectForProvisionalNavigation:":
		return delegate.WebView_DidReceiveServerRedirectForProvisionalNavigation != nil
	case "webView:didCommitNavigation:":
		return delegate.WebView_DidCommitNavigation != nil
	case "webView:didFinishNavigation:":
		return delegate.WebView_DidFinishNavigation != nil
	case "webView:didFailNavigation:withError:":
		return delegate.WebView_DidFailNavigation_WithError != nil
	case "webView:didFailProvisionalNavigation:withError:":
		return delegate.WebView_DidFailProvisionalNavigation_WithError != nil
	case "webViewWebContentProcessDidTerminate:":
		return delegate.WebViewWebContentProcessDidTerminate != nil
	case "webView:navigationAction:didBecomeDownload:":
		return delegate.WebView_NavigationAction_DidBecomeDownload != nil
	case "webView:navigationResponse:didBecomeDownload:":
		return delegate.WebView_NavigationResponse_DidBecomeDownload != nil
	default:
		return false
	}
}
