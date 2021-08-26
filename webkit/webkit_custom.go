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

func (w WKWebView) TakeSnapshotWithConfiguration(configuration SnapshotConfiguration, completionHandler func(image appkit.Image, err foundation.Error)) {
	handle := cgo.NewHandle(completionHandler)
	C.TakeSnapshotWithConfiguration(w.Ptr(), objc.ExtractPtr(configuration), C.uintptr_t(handle))
}

func (w WKWebView) EvaluateJavaScript(javascript string, completionHandler func(value objc.Object, err foundation.Error)) {
	handle := cgo.NewHandle(completionHandler)
	C.EvaluateJavaScript(w.Ptr(), foundation.NewString(javascript).Ptr(), C.uintptr_t(handle))
}

//export callTakeSnapshotWithConfiguration
func callTakeSnapshotWithConfiguration(hp C.uintptr_t, imagePtr unsafe.Pointer, errPtr unsafe.Pointer) {
	handler := cgo.Handle(hp).Value().(func(appkit.Image, foundation.Error))
	handler(appkit.MakeImage(imagePtr), foundation.MakeError(errPtr))
}

//export callEvaluateJavaScript
func callEvaluateJavaScript(hp C.uintptr_t, valuePtr unsafe.Pointer, errPtr unsafe.Pointer) {
	handler := cgo.Handle(hp).Value().(func(objc.Object, foundation.Error))
	handler(objc.MakeObject(valuePtr), foundation.MakeError(errPtr))
}

//export deleteWebKitHandle
func deleteWebKitHandle(hp C.uintptr_t) {
	cgo.Handle(hp).Delete()
}
