package coregraphics

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework CoreGraphics
// #include "image.h"
import "C"
import "unsafe"

type CGContext struct {
}

// TODO: should be *Image
type ContextRef CGContext

func ToCGContextRefPointer(cs ContextRef) unsafe.Pointer {
	panic("not implemented")
}

func FromCGContextRefPointer(p unsafe.Pointer) ContextRef {
	panic("not implemented")
}
