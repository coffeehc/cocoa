package appkit

// #include "appkit_custom.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"github.com/hsiafan/cocoa/objc"
	"runtime/cgo"
	"unsafe"
)

type extImage interface {
	CGImageForProposedRect_Context_Hints() coregraphics.ImageRef
}

func (n NSImage) CGImageForProposedRect_Context_Hints() coregraphics.ImageRef {
	result_ := C.C_NSImage_CGImageForProposedRect_Context_Hints(n.Ptr())
	return coregraphics.ImageRef(result_)
}

//export deleteAppKitHandle
func deleteAppKitHandle(hp C.uintptr_t) {
	cgo.Handle(hp).Delete()
}

type ModalSession unsafe.Pointer

type extSavePanel interface {
	BeginWithCompletionHandler(handler func(response ModalResponse))
	BeginSheetModalForWindow(window Window, handler func(response ModalResponse))
}

func (n NSSavePanel) BeginWithCompletionHandler(handler func(response ModalResponse)) {
	handle := cgo.NewHandle(handler)
	C.SavePanel_BeginWithCompletionHandler(n.Ptr(), C.uintptr_t(handle))
}

func (n NSSavePanel) BeginSheetModalForWindow(window Window, handler func(response ModalResponse)) {
	handle := cgo.NewHandle(handler)
	C.SavePanel_BeginSheetModalForWindow(n.Ptr(), objc.ExtractPtr(window), C.uintptr_t(handle))
}

//export callModalResponseHandler
func callModalResponseHandler(hp C.uintptr_t, response C.int) {
	handle := cgo.Handle(hp)
	handler := handle.Value().(func(ModalResponse))
	handler(ModalResponse(response))
	handle.Delete()
}
