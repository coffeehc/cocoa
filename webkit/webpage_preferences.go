package webkit

// #include "webpage_preferences.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type WebpagePreferences interface {
	objc.Object
	AllowsContentJavaScript() bool
	SetAllowsContentJavaScript(value bool)
	PreferredContentMode() ContentMode
	SetPreferredContentMode(value ContentMode)
}

type WKWebpagePreferences struct {
	objc.NSObject
}

func MakeWebpagePreferences(ptr unsafe.Pointer) WKWebpagePreferences {
	return WKWebpagePreferences{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocWebpagePreferences() WKWebpagePreferences {
	result_ := C.C_WKWebpagePreferences_AllocWebpagePreferences()
	return MakeWebpagePreferences(result_)
}

func (w WKWebpagePreferences) Init() WKWebpagePreferences {
	result_ := C.C_WKWebpagePreferences_Init(w.Ptr())
	return MakeWebpagePreferences(result_)
}

func NewWebpagePreferences() WKWebpagePreferences {
	result_ := C.C_WKWebpagePreferences_NewWebpagePreferences()
	return MakeWebpagePreferences(result_)
}

func (w WKWebpagePreferences) Autorelease() WKWebpagePreferences {
	result_ := C.C_WKWebpagePreferences_Autorelease(w.Ptr())
	return MakeWebpagePreferences(result_)
}

func (w WKWebpagePreferences) Retain() WKWebpagePreferences {
	result_ := C.C_WKWebpagePreferences_Retain(w.Ptr())
	return MakeWebpagePreferences(result_)
}

func (w WKWebpagePreferences) AllowsContentJavaScript() bool {
	result_ := C.C_WKWebpagePreferences_AllowsContentJavaScript(w.Ptr())
	return bool(result_)
}

func (w WKWebpagePreferences) SetAllowsContentJavaScript(value bool) {
	C.C_WKWebpagePreferences_SetAllowsContentJavaScript(w.Ptr(), C.bool(value))
}

func (w WKWebpagePreferences) PreferredContentMode() ContentMode {
	result_ := C.C_WKWebpagePreferences_PreferredContentMode(w.Ptr())
	return ContentMode(int(result_))
}

func (w WKWebpagePreferences) SetPreferredContentMode(value ContentMode) {
	C.C_WKWebpagePreferences_SetPreferredContentMode(w.Ptr(), C.int(int(value)))
}
