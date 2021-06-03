package webkit

// #include "web_view_configuration.h"
import "C"
import (
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type WebViewConfiguration interface {
	objc.Object
	SetURLSchemeHandler_ForURLScheme(urlSchemeHandler objc.Object, urlScheme string)
	UrlSchemeHandlerForURLScheme(urlScheme string) objc.Object
	WebsiteDataStore() WebsiteDataStore
	SetWebsiteDataStore(value WebsiteDataStore)
	UserContentController() UserContentController
	SetUserContentController(value UserContentController)
	ProcessPool() ProcessPool
	SetProcessPool(value ProcessPool)
	ApplicationNameForUserAgent() string
	SetApplicationNameForUserAgent(value string)
	LimitsNavigationsToAppBoundDomains() bool
	SetLimitsNavigationsToAppBoundDomains(value bool)
	Preferences() Preferences
	SetPreferences(value Preferences)
	DefaultWebpagePreferences() WebpagePreferences
	SetDefaultWebpagePreferences(value WebpagePreferences)
	SuppressesIncrementalRendering() bool
	SetSuppressesIncrementalRendering(value bool)
	AllowsAirPlayForMediaPlayback() bool
	SetAllowsAirPlayForMediaPlayback(value bool)
	MediaTypesRequiringUserActionForPlayback() AudiovisualMediaTypes
	SetMediaTypesRequiringUserActionForPlayback(value AudiovisualMediaTypes)
	UserInterfaceDirectionPolicy() UserInterfaceDirectionPolicy
	SetUserInterfaceDirectionPolicy(value UserInterfaceDirectionPolicy)
}

type WKWebViewConfiguration struct {
	objc.NSObject
}

func MakeWebViewConfiguration(ptr unsafe.Pointer) WKWebViewConfiguration {
	return WKWebViewConfiguration{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocWebViewConfiguration() WKWebViewConfiguration {
	return MakeWebViewConfiguration(C.C_WebViewConfiguration_Alloc())
}

func (w WKWebViewConfiguration) Init() WebViewConfiguration {
	result_ := C.C_WKWebViewConfiguration_Init(w.Ptr())
	return MakeWebViewConfiguration(result_)
}

func (w WKWebViewConfiguration) SetURLSchemeHandler_ForURLScheme(urlSchemeHandler objc.Object, urlScheme string) {
	C.C_WKWebViewConfiguration_SetURLSchemeHandler_ForURLScheme(w.Ptr(), objc.ExtractPtr(urlSchemeHandler), foundation.NewString(urlScheme).Ptr())
}

func (w WKWebViewConfiguration) UrlSchemeHandlerForURLScheme(urlScheme string) objc.Object {
	result_ := C.C_WKWebViewConfiguration_UrlSchemeHandlerForURLScheme(w.Ptr(), foundation.NewString(urlScheme).Ptr())
	return objc.MakeObject(result_)
}

func (w WKWebViewConfiguration) WebsiteDataStore() WebsiteDataStore {
	result_ := C.C_WKWebViewConfiguration_WebsiteDataStore(w.Ptr())
	return MakeWebsiteDataStore(result_)
}

func (w WKWebViewConfiguration) SetWebsiteDataStore(value WebsiteDataStore) {
	C.C_WKWebViewConfiguration_SetWebsiteDataStore(w.Ptr(), objc.ExtractPtr(value))
}

func (w WKWebViewConfiguration) UserContentController() UserContentController {
	result_ := C.C_WKWebViewConfiguration_UserContentController(w.Ptr())
	return MakeUserContentController(result_)
}

func (w WKWebViewConfiguration) SetUserContentController(value UserContentController) {
	C.C_WKWebViewConfiguration_SetUserContentController(w.Ptr(), objc.ExtractPtr(value))
}

func (w WKWebViewConfiguration) ProcessPool() ProcessPool {
	result_ := C.C_WKWebViewConfiguration_ProcessPool(w.Ptr())
	return MakeProcessPool(result_)
}

func (w WKWebViewConfiguration) SetProcessPool(value ProcessPool) {
	C.C_WKWebViewConfiguration_SetProcessPool(w.Ptr(), objc.ExtractPtr(value))
}

func (w WKWebViewConfiguration) ApplicationNameForUserAgent() string {
	result_ := C.C_WKWebViewConfiguration_ApplicationNameForUserAgent(w.Ptr())
	return foundation.MakeString(result_).String()
}

func (w WKWebViewConfiguration) SetApplicationNameForUserAgent(value string) {
	C.C_WKWebViewConfiguration_SetApplicationNameForUserAgent(w.Ptr(), foundation.NewString(value).Ptr())
}

func (w WKWebViewConfiguration) LimitsNavigationsToAppBoundDomains() bool {
	result_ := C.C_WKWebViewConfiguration_LimitsNavigationsToAppBoundDomains(w.Ptr())
	return bool(result_)
}

func (w WKWebViewConfiguration) SetLimitsNavigationsToAppBoundDomains(value bool) {
	C.C_WKWebViewConfiguration_SetLimitsNavigationsToAppBoundDomains(w.Ptr(), C.bool(value))
}

func (w WKWebViewConfiguration) Preferences() Preferences {
	result_ := C.C_WKWebViewConfiguration_Preferences(w.Ptr())
	return MakePreferences(result_)
}

func (w WKWebViewConfiguration) SetPreferences(value Preferences) {
	C.C_WKWebViewConfiguration_SetPreferences(w.Ptr(), objc.ExtractPtr(value))
}

func (w WKWebViewConfiguration) DefaultWebpagePreferences() WebpagePreferences {
	result_ := C.C_WKWebViewConfiguration_DefaultWebpagePreferences(w.Ptr())
	return MakeWebpagePreferences(result_)
}

func (w WKWebViewConfiguration) SetDefaultWebpagePreferences(value WebpagePreferences) {
	C.C_WKWebViewConfiguration_SetDefaultWebpagePreferences(w.Ptr(), objc.ExtractPtr(value))
}

func (w WKWebViewConfiguration) SuppressesIncrementalRendering() bool {
	result_ := C.C_WKWebViewConfiguration_SuppressesIncrementalRendering(w.Ptr())
	return bool(result_)
}

func (w WKWebViewConfiguration) SetSuppressesIncrementalRendering(value bool) {
	C.C_WKWebViewConfiguration_SetSuppressesIncrementalRendering(w.Ptr(), C.bool(value))
}

func (w WKWebViewConfiguration) AllowsAirPlayForMediaPlayback() bool {
	result_ := C.C_WKWebViewConfiguration_AllowsAirPlayForMediaPlayback(w.Ptr())
	return bool(result_)
}

func (w WKWebViewConfiguration) SetAllowsAirPlayForMediaPlayback(value bool) {
	C.C_WKWebViewConfiguration_SetAllowsAirPlayForMediaPlayback(w.Ptr(), C.bool(value))
}

func (w WKWebViewConfiguration) MediaTypesRequiringUserActionForPlayback() AudiovisualMediaTypes {
	result_ := C.C_WKWebViewConfiguration_MediaTypesRequiringUserActionForPlayback(w.Ptr())
	return AudiovisualMediaTypes(uint(result_))
}

func (w WKWebViewConfiguration) SetMediaTypesRequiringUserActionForPlayback(value AudiovisualMediaTypes) {
	C.C_WKWebViewConfiguration_SetMediaTypesRequiringUserActionForPlayback(w.Ptr(), C.uint(uint(value)))
}

func (w WKWebViewConfiguration) UserInterfaceDirectionPolicy() UserInterfaceDirectionPolicy {
	result_ := C.C_WKWebViewConfiguration_UserInterfaceDirectionPolicy(w.Ptr())
	return UserInterfaceDirectionPolicy(int(result_))
}

func (w WKWebViewConfiguration) SetUserInterfaceDirectionPolicy(value UserInterfaceDirectionPolicy) {
	C.C_WKWebViewConfiguration_SetUserInterfaceDirectionPolicy(w.Ptr(), C.int(int(value)))
}
