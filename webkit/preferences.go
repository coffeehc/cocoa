package webkit

// #include "preferences.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type Preferences interface {
	objc.Object
	MinimumFontSize() coregraphics.Float
	SetMinimumFontSize(value coregraphics.Float)
	TabFocusesLinks() bool
	SetTabFocusesLinks(value bool)
	JavaScriptCanOpenWindowsAutomatically() bool
	SetJavaScriptCanOpenWindowsAutomatically(value bool)
	IsFraudulentWebsiteWarningEnabled() bool
	SetFraudulentWebsiteWarningEnabled(value bool)
	TextInteractionEnabled() bool
	SetTextInteractionEnabled(value bool)
}

type WKPreferences struct {
	objc.NSObject
}

func MakePreferences(ptr unsafe.Pointer) WKPreferences {
	return WKPreferences{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocPreferences() WKPreferences {
	return MakePreferences(C.C_Preferences_Alloc())
}

func (w WKPreferences) Init() Preferences {
	result_ := C.C_WKPreferences_Init(w.Ptr())
	return MakePreferences(result_)
}

func (w WKPreferences) MinimumFontSize() coregraphics.Float {
	result_ := C.C_WKPreferences_MinimumFontSize(w.Ptr())
	return coregraphics.Float(float64(result_))
}

func (w WKPreferences) SetMinimumFontSize(value coregraphics.Float) {
	C.C_WKPreferences_SetMinimumFontSize(w.Ptr(), C.double(float64(value)))
}

func (w WKPreferences) TabFocusesLinks() bool {
	result_ := C.C_WKPreferences_TabFocusesLinks(w.Ptr())
	return bool(result_)
}

func (w WKPreferences) SetTabFocusesLinks(value bool) {
	C.C_WKPreferences_SetTabFocusesLinks(w.Ptr(), C.bool(value))
}

func (w WKPreferences) JavaScriptCanOpenWindowsAutomatically() bool {
	result_ := C.C_WKPreferences_JavaScriptCanOpenWindowsAutomatically(w.Ptr())
	return bool(result_)
}

func (w WKPreferences) SetJavaScriptCanOpenWindowsAutomatically(value bool) {
	C.C_WKPreferences_SetJavaScriptCanOpenWindowsAutomatically(w.Ptr(), C.bool(value))
}

func (w WKPreferences) IsFraudulentWebsiteWarningEnabled() bool {
	result_ := C.C_WKPreferences_IsFraudulentWebsiteWarningEnabled(w.Ptr())
	return bool(result_)
}

func (w WKPreferences) SetFraudulentWebsiteWarningEnabled(value bool) {
	C.C_WKPreferences_SetFraudulentWebsiteWarningEnabled(w.Ptr(), C.bool(value))
}

func (w WKPreferences) TextInteractionEnabled() bool {
	result_ := C.C_WKPreferences_TextInteractionEnabled(w.Ptr())
	return bool(result_)
}

func (w WKPreferences) SetTextInteractionEnabled(value bool) {
	C.C_WKPreferences_SetTextInteractionEnabled(w.Ptr(), C.bool(value))
}
