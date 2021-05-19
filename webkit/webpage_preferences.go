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

func MakeWebpagePreferences(ptr unsafe.Pointer) *WKWebpagePreferences {
	if ptr == nil {
		return nil
	}
	return &WKWebpagePreferences{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocWebpagePreferences() *WKWebpagePreferences {
	return MakeWebpagePreferences(C.C_WebpagePreferences_Alloc())
}

func (w *WKWebpagePreferences) Init() WebpagePreferences {
	result_ := C.C_WKWebpagePreferences_Init(w.Ptr())
	return MakeWebpagePreferences(result_)
}

func (w *WKWebpagePreferences) AllowsContentJavaScript() bool {
	result_ := C.C_WKWebpagePreferences_AllowsContentJavaScript(w.Ptr())
	return bool(result_)
}

func (w *WKWebpagePreferences) SetAllowsContentJavaScript(value bool) {
	C.C_WKWebpagePreferences_SetAllowsContentJavaScript(w.Ptr(), C.bool(value))
}

func (w *WKWebpagePreferences) PreferredContentMode() ContentMode {
	result_ := C.C_WKWebpagePreferences_PreferredContentMode(w.Ptr())
	return ContentMode(int(result_))
}

func (w *WKWebpagePreferences) SetPreferredContentMode(value ContentMode) {
	C.C_WKWebpagePreferences_SetPreferredContentMode(w.Ptr(), C.int(int(value)))
}
