package webkit

// #include "navigation_action.h"
import "C"
import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type NavigationAction interface {
	objc.Object
	NavigationType() NavigationType
	Request() foundation.URLRequest
	SourceFrame() FrameInfo
	TargetFrame() FrameInfo
	ButtonNumber() int
	ModifierFlags() appkit.EventModifierFlags
	ShouldPerformDownload() bool
}

type WKNavigationAction struct {
	objc.NSObject
}

func MakeNavigationAction(ptr unsafe.Pointer) WKNavigationAction {
	return WKNavigationAction{
		NSObject: objc.MakeObject(ptr),
	}
}

func AllocNavigationAction() WKNavigationAction {
	return MakeNavigationAction(C.C_NavigationAction_Alloc())
}

func (w WKNavigationAction) Init() NavigationAction {
	result_ := C.C_WKNavigationAction_Init(w.Ptr())
	return MakeNavigationAction(result_)
}

func (w WKNavigationAction) NavigationType() NavigationType {
	result_ := C.C_WKNavigationAction_NavigationType(w.Ptr())
	return NavigationType(int(result_))
}

func (w WKNavigationAction) Request() foundation.URLRequest {
	result_ := C.C_WKNavigationAction_Request(w.Ptr())
	return foundation.MakeURLRequest(result_)
}

func (w WKNavigationAction) SourceFrame() FrameInfo {
	result_ := C.C_WKNavigationAction_SourceFrame(w.Ptr())
	return MakeFrameInfo(result_)
}

func (w WKNavigationAction) TargetFrame() FrameInfo {
	result_ := C.C_WKNavigationAction_TargetFrame(w.Ptr())
	return MakeFrameInfo(result_)
}

func (w WKNavigationAction) ButtonNumber() int {
	result_ := C.C_WKNavigationAction_ButtonNumber(w.Ptr())
	return int(result_)
}

func (w WKNavigationAction) ModifierFlags() appkit.EventModifierFlags {
	result_ := C.C_WKNavigationAction_ModifierFlags(w.Ptr())
	return appkit.EventModifierFlags(uint(result_))
}

func (w WKNavigationAction) ShouldPerformDownload() bool {
	result_ := C.C_WKNavigationAction_ShouldPerformDownload(w.Ptr())
	return bool(result_)
}
