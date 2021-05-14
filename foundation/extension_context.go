package foundation

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation
// #include "extension_context.h"
import "C"
import (
	"github.com/hsiafan/cocoa/objc"
	"unsafe"
)

type ExtensionContext interface {
	objc.Object
	CancelRequestWithError(error Error)
	MediaPlayingStarted()
	MediaPlayingPaused()
	DismissNotificationContentExtension()
	PerformNotificationDefaultAction()
}

type NSExtensionContext struct {
	objc.NSObject
}

func MakeExtensionContext(ptr unsafe.Pointer) *NSExtensionContext {
	if ptr == nil {
		return nil
	}
	return &NSExtensionContext{
		NSObject: *objc.MakeObject(ptr),
	}
}

func AllocExtensionContext() *NSExtensionContext {
	return MakeExtensionContext(C.C_ExtensionContext_Alloc())
}

func (n *NSExtensionContext) Init() ExtensionContext {
	result := C.C_NSExtensionContext_Init(n.Ptr())
	return MakeExtensionContext(result)
}

func (n *NSExtensionContext) CancelRequestWithError(error Error) {
	C.C_NSExtensionContext_CancelRequestWithError(n.Ptr(), objc.ExtractPtr(error))
}

func (n *NSExtensionContext) MediaPlayingStarted() {
	C.C_NSExtensionContext_MediaPlayingStarted(n.Ptr())
}

func (n *NSExtensionContext) MediaPlayingPaused() {
	C.C_NSExtensionContext_MediaPlayingPaused(n.Ptr())
}

func (n *NSExtensionContext) DismissNotificationContentExtension() {
	C.C_NSExtensionContext_DismissNotificationContentExtension(n.Ptr())
}

func (n *NSExtensionContext) PerformNotificationDefaultAction() {
	C.C_NSExtensionContext_PerformNotificationDefaultAction(n.Ptr())
}
