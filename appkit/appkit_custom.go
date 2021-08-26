package appkit

// #include "appkit_custom.h"
import "C"
import (
	"github.com/hsiafan/cocoa/coregraphics"
	"runtime/cgo"
)

func (n NSImage) CGImageForProposedRect_Context_Hints() coregraphics.ImageRef {
	result_ := C.C_NSImage_CGImageForProposedRect_Context_Hints(n.Ptr())
	return coregraphics.ImageRef(result_)
}

//export deleteAppKitHandle
func deleteAppKitHandle(hp C.uintptr_t) {
	cgo.Handle(hp).Delete()
}
