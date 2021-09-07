package webkit

// #include "webkit_custom.h"
import "C"
import (
	"github.com/hsiafan/cocoa/appkit"
	"github.com/hsiafan/cocoa/foundation"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type extWebView interface {
	TakeSnapshotWithConfiguration(configuration SnapshotConfiguration, handler func(image appkit.Image, err foundation.Error))
	EvaluateJavaScript(javascript string, handler func(value objc.Object, err foundation.Error))
}

func (w WKWebView) TakeSnapshotWithConfiguration(configuration SnapshotConfiguration, handler func(image appkit.Image, err foundation.Error)) {
	handle := cgo.NewHandle(handler)
	C.WebView_TakeSnapshotWithConfiguration(w.Ptr(), objc.ExtractPtr(configuration), C.uintptr_t(handle))
}

func (w WKWebView) EvaluateJavaScript(javascript string, handler func(value objc.Object, err foundation.Error)) {
	handle := cgo.NewHandle(handler)
	C.WebView_EvaluateJavaScript(w.Ptr(), foundation.NewString(javascript).Ptr(), C.uintptr_t(handle))
}

//export callTakeSnapshotHandler
func callTakeSnapshotHandler(hp C.uintptr_t, imagePtr unsafe.Pointer, errPtr unsafe.Pointer) {
	handle := cgo.Handle(hp)
	handler := handle.Value().(func(appkit.Image, foundation.Error))
	handler(appkit.MakeImage(imagePtr), foundation.MakeError(errPtr))
	handle.Delete()
}

//export callEvaluateJavaScriptHandler
func callEvaluateJavaScriptHandler(hp C.uintptr_t, valuePtr unsafe.Pointer, errPtr unsafe.Pointer) {
	handle := cgo.Handle(hp)
	handler := handle.Value().(func(objc.Object, foundation.Error))
	handler(objc.MakeObject(valuePtr), foundation.MakeError(errPtr))
	handle.Delete()
}

//export deleteWebKitHandle
func deleteWebKitHandle(hp C.uintptr_t) {
	cgo.Handle(hp).Delete()
}
