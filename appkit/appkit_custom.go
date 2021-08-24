package appkit

// #include "appkit_custom.h"
import "C"
import "github.com/hsiafan/cocoa/coregraphics"

func (n NSImage) CGImageForProposedRect_Context_Hints() coregraphics.ImageRef {
	result_ := C.C_NSImage_CGImageForProposedRect_Context_Hints(n.Ptr())
	return coregraphics.ImageRef(result_)
}
