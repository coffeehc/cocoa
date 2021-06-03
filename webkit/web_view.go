package webkit

// #include "web_view.h"
import "C"
import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type WebView interface {
	appkit.View
	LoadRequest(request foundation.URLRequest) Navigation
	LoadHTMLString_BaseURL(_string string, baseURL foundation.URL) Navigation
	LoadFileURL_AllowingReadAccessToURL(URL foundation.URL, readAccessURL foundation.URL) Navigation
	LoadData_MIMEType_CharacterEncodingName_BaseURL(data []byte, MIMEType string, characterEncodingName string, baseURL foundation.URL) Navigation
	Reload() Navigation
	ReloadFromOrigin() Navigation
	StopLoading()
	SetMagnification_CenteredAtPoint(magnification coregraphics.Float, point coregraphics.Point)
	GoBack(sender objc.Object)
	GoForward(sender objc.Object)
	GoToBackForwardListItem(item BackForwardListItem) Navigation
	PrintOperationWithPrintInfo(printInfo appkit.PrintInfo) appkit.PrintOperation
	CloseAllMediaPresentations()
	Configuration() WebViewConfiguration
	UIDelegate() objc.Object
	SetUIDelegate(value objc.Object)
	NavigationDelegate() objc.Object
	SetNavigationDelegate(value objc.Object)
	IsLoading() bool
	EstimatedProgress() float64
	Title() string
	URL() foundation.URL
	MediaType() string
	SetMediaType(value string)
	CustomUserAgent() string
	SetCustomUserAgent(value string)
	HasOnlySecureContent() bool
	PageZoom() coregraphics.Float
	SetPageZoom(value coregraphics.Float)
	AllowsMagnification() bool
	SetAllowsMagnification(value bool)
	Magnification() coregraphics.Float
	SetMagnification(value coregraphics.Float)
	AllowsBackForwardNavigationGestures() bool
	SetAllowsBackForwardNavigationGestures(value bool)
	BackForwardList() BackForwardList
	CanGoBack() bool
	CanGoForward() bool
	AllowsLinkPreview() bool
	SetAllowsLinkPreview(value bool)
}

type WKWebView struct {
	appkit.NSView
}

func MakeWebView(ptr unsafe.Pointer) WKWebView {
	return WKWebView{
		NSView: appkit.MakeView(ptr),
	}
}

func AllocWebView() WKWebView {
	return MakeWebView(C.C_WebView_Alloc())
}

func (w WKWebView) InitWithFrame_Configuration(frame coregraphics.Rect, configuration WebViewConfiguration) WebView {
	result_ := C.C_WKWebView_InitWithFrame_Configuration(w.Ptr(), *(*C.CGRect)(coregraphics.ToCGRectPointer(frame)), objc.ExtractPtr(configuration))
	return MakeWebView(result_)
}

func (w WKWebView) InitWithCoder(coder foundation.Coder) WebView {
	result_ := C.C_WKWebView_InitWithCoder(w.Ptr(), objc.ExtractPtr(coder))
	return MakeWebView(result_)
}

func (w WKWebView) Init() WebView {
	result_ := C.C_WKWebView_Init(w.Ptr())
	return MakeWebView(result_)
}

func WebView_HandlesURLScheme(urlScheme string) bool {
	result_ := C.C_WKWebView_WebView_HandlesURLScheme(foundation.NewString(urlScheme).Ptr())
	return bool(result_)
}

func (w WKWebView) LoadRequest(request foundation.URLRequest) Navigation {
	result_ := C.C_WKWebView_LoadRequest(w.Ptr(), objc.ExtractPtr(request))
	return MakeNavigation(result_)
}

func (w WKWebView) LoadHTMLString_BaseURL(_string string, baseURL foundation.URL) Navigation {
	result_ := C.C_WKWebView_LoadHTMLString_BaseURL(w.Ptr(), foundation.NewString(_string).Ptr(), objc.ExtractPtr(baseURL))
	return MakeNavigation(result_)
}

func (w WKWebView) LoadFileURL_AllowingReadAccessToURL(URL foundation.URL, readAccessURL foundation.URL) Navigation {
	result_ := C.C_WKWebView_LoadFileURL_AllowingReadAccessToURL(w.Ptr(), objc.ExtractPtr(URL), objc.ExtractPtr(readAccessURL))
	return MakeNavigation(result_)
}

func (w WKWebView) LoadData_MIMEType_CharacterEncodingName_BaseURL(data []byte, MIMEType string, characterEncodingName string, baseURL foundation.URL) Navigation {
	result_ := C.C_WKWebView_LoadData_MIMEType_CharacterEncodingName_BaseURL(w.Ptr(), C.Array{data: unsafe.Pointer(&data[0]), len: C.int(len(data))}, foundation.NewString(MIMEType).Ptr(), foundation.NewString(characterEncodingName).Ptr(), objc.ExtractPtr(baseURL))
	return MakeNavigation(result_)
}

func (w WKWebView) Reload() Navigation {
	result_ := C.C_WKWebView_Reload(w.Ptr())
	return MakeNavigation(result_)
}

func (w WKWebView) ReloadFromOrigin() Navigation {
	result_ := C.C_WKWebView_ReloadFromOrigin(w.Ptr())
	return MakeNavigation(result_)
}

func (w WKWebView) StopLoading() {
	C.C_WKWebView_StopLoading(w.Ptr())
}

func (w WKWebView) SetMagnification_CenteredAtPoint(magnification coregraphics.Float, point coregraphics.Point) {
	C.C_WKWebView_SetMagnification_CenteredAtPoint(w.Ptr(), C.double(float64(magnification)), *(*C.CGPoint)(coregraphics.ToCGPointPointer(point)))
}

func (w WKWebView) GoBack(sender objc.Object) {
	C.C_WKWebView_GoBack(w.Ptr(), objc.ExtractPtr(sender))
}

func (w WKWebView) GoForward(sender objc.Object) {
	C.C_WKWebView_GoForward(w.Ptr(), objc.ExtractPtr(sender))
}

func (w WKWebView) GoToBackForwardListItem(item BackForwardListItem) Navigation {
	result_ := C.C_WKWebView_GoToBackForwardListItem(w.Ptr(), objc.ExtractPtr(item))
	return MakeNavigation(result_)
}

func (w WKWebView) PrintOperationWithPrintInfo(printInfo appkit.PrintInfo) appkit.PrintOperation {
	result_ := C.C_WKWebView_PrintOperationWithPrintInfo(w.Ptr(), objc.ExtractPtr(printInfo))
	return appkit.MakePrintOperation(result_)
}

func (w WKWebView) CloseAllMediaPresentations() {
	C.C_WKWebView_CloseAllMediaPresentations(w.Ptr())
}

func (w WKWebView) Configuration() WebViewConfiguration {
	result_ := C.C_WKWebView_Configuration(w.Ptr())
	return MakeWebViewConfiguration(result_)
}

func (w WKWebView) UIDelegate() objc.Object {
	result_ := C.C_WKWebView_UIDelegate(w.Ptr())
	return objc.MakeObject(result_)
}

func (w WKWebView) SetUIDelegate(value objc.Object) {
	C.C_WKWebView_SetUIDelegate(w.Ptr(), objc.ExtractPtr(value))
}

func (w WKWebView) NavigationDelegate() objc.Object {
	result_ := C.C_WKWebView_NavigationDelegate(w.Ptr())
	return objc.MakeObject(result_)
}

func (w WKWebView) SetNavigationDelegate(value objc.Object) {
	C.C_WKWebView_SetNavigationDelegate(w.Ptr(), objc.ExtractPtr(value))
}

func (w WKWebView) IsLoading() bool {
	result_ := C.C_WKWebView_IsLoading(w.Ptr())
	return bool(result_)
}

func (w WKWebView) EstimatedProgress() float64 {
	result_ := C.C_WKWebView_EstimatedProgress(w.Ptr())
	return float64(result_)
}

func (w WKWebView) Title() string {
	result_ := C.C_WKWebView_Title(w.Ptr())
	return foundation.MakeString(result_).String()
}

func (w WKWebView) URL() foundation.URL {
	result_ := C.C_WKWebView_URL(w.Ptr())
	return foundation.MakeURL(result_)
}

func (w WKWebView) MediaType() string {
	result_ := C.C_WKWebView_MediaType(w.Ptr())
	return foundation.MakeString(result_).String()
}

func (w WKWebView) SetMediaType(value string) {
	C.C_WKWebView_SetMediaType(w.Ptr(), foundation.NewString(value).Ptr())
}

func (w WKWebView) CustomUserAgent() string {
	result_ := C.C_WKWebView_CustomUserAgent(w.Ptr())
	return foundation.MakeString(result_).String()
}

func (w WKWebView) SetCustomUserAgent(value string) {
	C.C_WKWebView_SetCustomUserAgent(w.Ptr(), foundation.NewString(value).Ptr())
}

func (w WKWebView) HasOnlySecureContent() bool {
	result_ := C.C_WKWebView_HasOnlySecureContent(w.Ptr())
	return bool(result_)
}

func (w WKWebView) PageZoom() coregraphics.Float {
	result_ := C.C_WKWebView_PageZoom(w.Ptr())
	return coregraphics.Float(float64(result_))
}

func (w WKWebView) SetPageZoom(value coregraphics.Float) {
	C.C_WKWebView_SetPageZoom(w.Ptr(), C.double(float64(value)))
}

func (w WKWebView) AllowsMagnification() bool {
	result_ := C.C_WKWebView_AllowsMagnification(w.Ptr())
	return bool(result_)
}

func (w WKWebView) SetAllowsMagnification(value bool) {
	C.C_WKWebView_SetAllowsMagnification(w.Ptr(), C.bool(value))
}

func (w WKWebView) Magnification() coregraphics.Float {
	result_ := C.C_WKWebView_Magnification(w.Ptr())
	return coregraphics.Float(float64(result_))
}

func (w WKWebView) SetMagnification(value coregraphics.Float) {
	C.C_WKWebView_SetMagnification(w.Ptr(), C.double(float64(value)))
}

func (w WKWebView) AllowsBackForwardNavigationGestures() bool {
	result_ := C.C_WKWebView_AllowsBackForwardNavigationGestures(w.Ptr())
	return bool(result_)
}

func (w WKWebView) SetAllowsBackForwardNavigationGestures(value bool) {
	C.C_WKWebView_SetAllowsBackForwardNavigationGestures(w.Ptr(), C.bool(value))
}

func (w WKWebView) BackForwardList() BackForwardList {
	result_ := C.C_WKWebView_BackForwardList(w.Ptr())
	return MakeBackForwardList(result_)
}

func (w WKWebView) CanGoBack() bool {
	result_ := C.C_WKWebView_CanGoBack(w.Ptr())
	return bool(result_)
}

func (w WKWebView) CanGoForward() bool {
	result_ := C.C_WKWebView_CanGoForward(w.Ptr())
	return bool(result_)
}

func (w WKWebView) AllowsLinkPreview() bool {
	result_ := C.C_WKWebView_AllowsLinkPreview(w.Ptr())
	return bool(result_)
}

func (w WKWebView) SetAllowsLinkPreview(value bool) {
	C.C_WKWebView_SetAllowsLinkPreview(w.Ptr(), C.bool(value))
}
